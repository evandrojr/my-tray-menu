package main

import (
	"bufio"
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
)

var menuItensPtr []*systray.MenuItem
var options []Option
var programPath string

type Option struct {
	label   string
	command string
}

func main() {
	setProgramPath()
	options = loadConfig(filepath.Join(programPath, "my-tray-menu.yaml"))
	time.Sleep(1 * time.Second)
	systray.Run(onReady, onExit)
}

func onReady() {
	systray.SetIcon(getIcon(filepath.Join(programPath, "assets/menu.ico")))
	menuItensPtr = make([]*systray.MenuItem, 0)

	for _, v := range options {
		if(strings.ToLower(v.label)=="separator"){
			systray.AddSeparator()
			continue
		}
		menuItemPtr := systray.AddMenuItem(v.label, v.label)
		menuItensPtr = append(menuItensPtr, menuItemPtr)
	}
	systray.AddSeparator()
	mQuit := systray.AddMenuItem("Quit", "Quits this app")

	cmdChan := make(chan string)

	for i, menuItenPtr := range menuItensPtr {
		go func(c chan struct{}, cmd string) {
			for range c {
				cmdChan <- cmd
			}
		}(menuItenPtr.ClickedCh, options[i].command)
	}

	go func() {
		for {
			select {
			case cmd := <-cmdChan:
				execute(cmd) // Handle command
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

func setProgramPath() {
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

func loadConfig(path string) []Option {

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	options := make([]Option, 0)
	for scanner.Scan() {
		line := scanner.Text()
		i := strings.Index(line, ":")
		label := strings.TrimSpace(line[0:i])
		command := strings.TrimSpace(line[i+1:])

		option := Option{
			label:   label,
			command: command,
		}
		options = append(options, option)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return options
}
