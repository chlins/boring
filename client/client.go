package client

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"os"
	"time"

	"github.com/chlins/boring/common"
	"github.com/fatih/color"
)

// Peer client
type Peer struct {
	conn net.Conn
}

// New client
func New(server string) *Peer {
	c, err := net.Dial("tcp", server)
	if err != nil {
		fmt.Printf("Connect to server fail, %s\n", err.Error())
		os.Exit(1)
	}

	return &Peer{
		conn: c,
	}
}

// RWLoop read and write loop
func (p *Peer) RWLoop() {
	go p.readLoop()

	// write msg from stdin
	for {
		stdReader := bufio.NewReader(os.Stdin)

		c := color.New(color.FgCyan)
		c.Print("> ")
		data, err := stdReader.ReadString('\n')
		if err != nil {
			fmt.Printf("%s\n", err.Error())
			return
		}

		_, err = p.conn.Write([]byte(data))
		if err != nil {
			fmt.Printf("%s\n", err.Error())
			return
		}

		time.Sleep(200 * time.Millisecond)
	}
}

func (p *Peer) readLoop() {
	for {
		buf := make([]byte, 1024)
		n, err := p.conn.Read(buf)
		if err != nil {
			fmt.Printf("Read Loop error, %s\n", err.Error())
			os.Exit(1)
		}

		var msg common.Msg
		err = json.Unmarshal(buf[:n], &msg)
		if err != nil {
			fmt.Printf("Json unmarshal error: %s\n", err.Error())
			continue
		}

		color.Yellow("------------- New Message -------------")
		color.Blue("%s [%s] say", msg.Time, msg.Sender)
		color.Magenta("%s", msg.Content)
		color.Yellow("----------------  END -----------------")
	}
}
