package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/getlantern/systray"
	"github.com/tawesoft/golib/v2/dialog"
)

type MenuItenType int64

// Types of item
const (
	Choice    MenuItenType = iota
	Separator MenuItenType = iota
)

type MenuOption struct {
	label   string
	command string
}

type MenuIten struct {
	menuItenType MenuItenType
	label        string
	command      string
}

var menuItensPtr []*systray.MenuItem
var menuOptions []MenuOption
var menuItens []MenuIten
var programPath string

func main() {
	menuItens = make([]MenuIten, 0)
	setProgramPath()
	loadConfig(filepath.Join(programPath, "my-tray-menu.yaml"))
	parsePaths()
	time.Sleep(1 * time.Second)
	systray.Run(onReady, onExit)
}

func parsePaths() {
	for i := range menuOptions {
		menuOptions[i].command = strings.ReplaceAll(menuOptions[i].command, "{PROGRAMPATH}", programPath)
	}
}

func onReady() {
	systray.SetIcon(getIcon(filepath.Join(programPath, "assets/menu.ico")))
	menuItensPtr = make([]*systray.MenuItem, 0)

	indexOption := 0
	for i := range menuItens {
		if menuItens[i].menuItenType == Separator {
			systray.AddSeparator()
			continue
		}
		menuItemPtr := systray.AddMenuItem(menuOptions[indexOption].label, menuOptions[indexOption].label)
		menuItensPtr = append(menuItensPtr, menuItemPtr)
		indexOption++
	}
	systray.AddSeparator()
	mQuit := systray.AddMenuItem("Quit", "Quits this app")
	cmdChan := make(chan string)

	for i, menuItenPtr := range menuItensPtr {
		go func(c chan struct{}, cmd string) {
			for range c {
				cmdChan <- cmd
			}
		}(menuItenPtr.ClickedCh, menuOptions[i].command)
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
	b, err := os.ReadFile(s)
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
}

func execute(commands string) {
	println(commands)
	command_array := strings.Split(commands, " ")
	command := ""
	command, command_array = command_array[0], command_array[1:]
	cmd := exec.Command(command, command_array...)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		errMsg := fmt.Sprintf("Error executing command: %s", err)
		fmt.Println(errMsg)
		dialog.Error(errMsg)
		// log.Fatal(err)
	}
	fmt.Printf("Output: %s\n", out.String())
}

func loadConfig(path string) {

	file, err := os.Open(path)
	if err != nil {
		dialog.Error("Erro loading my-tray-menu.yaml %s", err)
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		i := strings.Index(line, ":")
		if i == -1 {
			errMsg := "Erro loading my-tray-menu.yaml, invalid format"
			dialog.Error(errMsg)
			log.Fatal(errMsg)
		}
		label := strings.TrimSpace(line[0:i])
		command := strings.TrimSpace(line[i+1:])
		var menuItenType MenuItenType = Choice
		if strings.ToLower(label) == "[separator]" {
			menuItenType = Separator
		}

		menuIten := MenuIten{
			label:        label,
			command:      command,
			menuItenType: menuItenType,
		}
		menuItens = append(menuItens, menuIten)

		if menuItenType == Choice {
			option := MenuOption{
				label:   label,
				command: command,
			}
			menuOptions = append(menuOptions, option)
		}
	}

	if err := scanner.Err(); err != nil {
		dialog.Error("Erro loading my-tray-menu.yaml %s", err)
		log.Fatal(err)
	}

	// return options, menuItens
}
