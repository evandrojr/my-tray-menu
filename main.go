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
	systray.SetIcon(getIcon("assets/clock.ico"))
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
			systray.SetTitle("My menu")
			systray.SetTooltip("My menu")
			time.Sleep(1 * time.Second)
		}
	}()

	go func() {

		// chans = make()
		// for  menuItenPtr := range menuItensPtr {
		
		// }

		// agg := make(chan string)

		// for _, ch := range menuItensPtr.Channels {
		//   go func(c chan string) {
		// 	for msg := range c {
		// 	  agg <- msg
		// 	}
		//   }(ch)
		// }

		tick := time.After(200 * time.Millisecond)
	

		for{

			// fmt.Printf("Loop")
			for i, menuItenPtr := range menuItensPtr {

				// fmt.Printf("Loop menu %d", i)
				select { 
				
				case <-tick:
					fmt.Println("tick.")

					
				case<-menuItenPtr.ClickedCh:
					execute(commands[i])
					break
				
						
				// default:
					
				}

				time.Sleep(1 * time.Millisecond)
				// fmt.Printf("Name: %s Age: %d\n", dog.Name, dog.Age)
				// fmt.Printf("Addr: %p\n", &dog)
		
				// fmt.Println("")
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
	// Cleaning stuff here.
}

func getIcon(s string) []byte {
	b, err := ioutil.ReadFile(s)
	if err != nil {
		fmt.Print(err)
	}
	return b
}

func execute(commands string){
	// args := []string{"what", "ever", "you", "like"}
	
	command_array:= strings.Split(commands, " ")
	command:="" 
	command, command_array = command_array[0], command_array[1:]
	cmd := exec.Command(command, command_array ...)
	
	// fmt.Printf("Output %q\n", x)

    var out bytes.Buffer
    cmd.Stdout = &out

    err := cmd.Run()

    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Output %q\n", out.String())
}

func readconfig()  map[string]string{
	yfile, err := ioutil.ReadFile("my-menu.yaml")

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
		//  fmt.Printf( k, string(v))
		 
	}
	return data
}
