package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/getlantern/systray"
	"gopkg.in/yaml.v3"
)

	var menuItensPtr []*systray.MenuItem
	var config map[string]string
	var commands []string
	var labels []string
	var programPath string

	func main() {
		setProgramPath()
		config = readconfig()
		time.Sleep(1 * time.Second)
		systray.Run(onReady, onExit)
	}

	func onReady() {
		systray.SetIcon(getIcon(filepath.Join(programPath,"assets/menu.ico")))
		menuItensPtr = make([]*systray.MenuItem, 0)
		i := 0
		op0 := systray.AddMenuItem(labels[i], commands[i])
		i++
		op1 := systray.AddMenuItem(labels[i], commands[i])
		i++
		op2 := systray.AddMenuItem(labels[i], commands[i])
		i++
		op3 := systray.AddMenuItem(labels[i], commands[i])
		i++

		systray.AddSeparator()
		mQuit := systray.AddMenuItem("Quit", "Quits this app")
		go func() {
			for {
				systray.SetTitle("My tray menu")
				systray.SetTooltip("https://github.com/evandrojr/my-tray-menu")
				time.Sleep(1 * time.Second)
			}
		}()

		go func() {
			for {
				select {
				case <-op0.ClickedCh:
					execute(commands[0])
				case <-op1.ClickedCh:
					execute(commands[1])
				case <-op2.ClickedCh:
					execute(commands[2])
				case <-op3.ClickedCh:
					execute(commands[3])
				case <-mQuit.ClickedCh:
					systray.Quit()
					return
				}
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

	func setProgramPath(){
		ex, err := os.Executable()
		if err != nil {
			panic(err)
		}
		programPath = filepath.Dir(ex)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}

	func execute(commands string) {
		command_array := strings.Split(commands, " ")
		command := ""
		command, command_array = command_array[0], command_array[1:]
		cmd := exec.Command(command, command_array...)
		var out bytes.Buffer
		cmd.Stdout = &out
		err := cmd.Run()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Output %s\n", out.String())
	}

	func readconfig() map[string]string {
		yfile, err := ioutil.ReadFile(filepath.Join(programPath,"my-tray-menu.yaml"))
		if err != nil {
			log.Fatal(err)
		}

		data := make(map[string]string)
		err2 := yaml.Unmarshal(yfile, &data)
		if err2 != nil {
			log.Fatal(err2)
		}

		labels = make([]string, 0)
		commands = make([]string, 0)

		for k, v := range data {
			labels = append(labels, k)
			commands = append(commands, v)
			fmt.Printf("%s -> %s\n", k, v)
		}
		fmt.Print(len(labels))
		return data
	}
