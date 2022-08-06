package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
	"strings"
	"time"

	"github.com/getlantern/systray"
	"gopkg.in/yaml.v3"
)

var menuItensPtr []*systray.MenuItem
var config map[string]string
var commands []string

func main() {
	config = readconfig()
	systray.Run(onReady, onExit)
}

func onReady() {
	systray.SetIcon(getIcon("assets/menu.ico"))
	menuItensPtr = make([]*systray.MenuItem,0)
	commands = make([]string,0)
	for k, v := range config {
		menuItemPtr := systray.AddMenuItem(k, k)
		menuItensPtr = append(menuItensPtr, menuItemPtr)
		commands = append(commands, v)
   }
   systray.AddSeparator()
	// mQuit := systray.AddMenuItem("Quit", "Quits this app")
	go func() {
		for {
			systray.SetTitle("My tray menu")
			systray.SetTooltip("https://github.com/evandrojr/my-tray-menu")
			time.Sleep(1 * time.Second)
		}
	}()

	go func() {
		for{
			for i, menuItenPtr := range menuItensPtr {
				select { 
				case<-menuItenPtr.ClickedCh:
					execute(commands[i])
					break
				}
			}	
			// select {
			// case <-mQuit.ClickedCh:
			// 	systray.Quit()
			// 	return
			// // default:
			// }
		}
	}()
}

func onExit() {
	// Cleaning stuff will go here. 
}

func getIcon(s string) []byte {
	b, err := ioutil.ReadFile(s)
	if err != nil {
		fmt.Print(err)
	}
	return b
}

func execute(commands string){
	command_array:= strings.Split(commands, " ")
	command:="" 
	command, command_array = command_array[0], command_array[1:]
	cmd := exec.Command(command, command_array ...)
    var out bytes.Buffer
    cmd.Stdout = &out
    err := cmd.Run()
    if err != nil {
        log.Fatal(err)
    }
    // fmt.Printf("Output %s\n", out.String())
}

func readconfig()  map[string]string{
	yfile, err := ioutil.ReadFile("my-tray-menu.yaml")
	if err != nil {
		 log.Fatal(err)
	}
	data := make(map[string]string)
	err2 := yaml.Unmarshal(yfile, &data)
	if err2 != nil {
		 log.Fatal(err2)
	}
	for k, v := range data {
		 fmt.Printf("%s -> %s\n", k, v)
	}
	return data
}
