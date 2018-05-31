package bing

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/fatih/color"
	"github.com/schollz/closestmatch"
)

const api = "http://xtk.azurewebsites.net/BingDictService.aspx?Word="

// Translate 查词翻译
func Translate(word string) *WordModel {
	resp, err := http.Get(api + word)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	var Word WordModel
	if err := json.Unmarshal(body, &Word); err != nil {
		return nil
	}

	return &Word
}

// Print 打印单词
func Print(w *WordModel) {
	fmt.Println()

	color.Cyan("%s: \n", w.Word)
	fmt.Println()
	color.Yellow("美: [%s]   英: [%s]", w.Pronunciation.AmE, w.Pronunciation.BrE)

	fmt.Println()
	for _, w := range w.Defs {
		if w.Pos == "Web" {
			color.Green("	- 网络	%s\n", w.Def)
			continue
		}
		color.Green("	- %s	%s\n", w.Pos, w.Def)
	}
	fmt.Println()
	color.Cyan("例句: \n\n")

	for i, s := range w.Sams {
		ws := strings.Split(s.Eng, " ")
		cm := closestmatch.New(ws, []int{2})
		keyword := cm.Closest(w.Word)

		fmt.Printf("	%d. ", i+1)
		for _, k := range ws {
			if keyword == k {
				fmt.Printf("%c[1;0;31m%s%c[0m ", 0x1B, k, 0x1B)
			} else {
				fmt.Printf("%s ", k)
			}
		}
		fmt.Printf("\n")

		color.Green("	   %s", s.Chn)
		fmt.Println()
	}
}
