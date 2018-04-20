package server

import (
	"encoding/json"
	"log"
	"net"
	"sync"
	"time"

	"github.com/chlins/boring/common"
)

// Controller model
type Controller struct {
	// listener
	listener net.Listener
	// all clients
	clients []net.Conn
	// message channel is gate of message, default buffer 10
	msgChannel chan *common.Msg

	sync.Mutex
}

// New controller
func New() *Controller {
	lsn, err := net.Listen("tcp", "0.0.0.0:0")
	if err != nil {
		log.Fatalf("Start Listen fail, %s\n", err.Error())
		return nil
	}

	return &Controller{
		listener:   lsn,
		clients:    make([]net.Conn, 0),
		msgChannel: make(chan *common.Msg, 10),
	}
}

// Addr get listener addr
func (c *Controller) Addr() string {
	return c.listener.Addr().String()
}

// Accept handle out connection
func (c *Controller) Accept() {
	go c.dispatch()

	for {
		conn, err := c.listener.Accept()
		if err != nil {
			log.Printf("Accept out conn fail, %s\n", err.Error())
			continue
		}

		go c.handle(conn)
	}
}

// handle conn
func (c *Controller) handle(conn net.Conn) {
	// add client
	c.Lock()
	c.clients = append(c.clients, conn)
	c.Unlock()

	defer func() {
		// delete client
		conn.Close()

		for i, co := range c.clients {
			if co == conn {
				c.Lock()
				c.clients = append(c.clients[:i], c.clients[i+1:]...)
				c.Unlock()
			}
		}
	}()

	for {
		buf := make([]byte, 1024)
		n, err := conn.Read(buf)

		if err != nil {
			return
		}

		msg := &common.Msg{
			Time:    time.Now().Format("2006-01-02 15:04:05"),
			Sender:  conn.RemoteAddr().String(),
			Content: string(buf[:n]),
		}

		c.msgChannel <- msg
	}

}

// dispatch msg
func (c *Controller) dispatch() {
	for {
		select {
		case msg := <-c.msgChannel:
			m, err := json.Marshal(msg)
			if err != nil {
				log.Printf("Json marshal error: %s\n", err.Error())
				continue
			}

			log.Println("Receive client msg: ", string(m))

			for _, conn := range c.clients {
				conn.Write(m)
			}
		}
	}
}
