package main

import (
	"fmt"
	"testing"
)

func TestLoadConfig(t *testing.T) {

	var menuOptions []MenuOption
	var menuItens []MenuIten
	// var programPath string
	
	type MenuItenType int64
	
	const (
		Choice    MenuItenType = 0
		Separator              = 1
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
	

   menuOptions,menuItens= loadConfig("my-tray-menu.yaml")
   fmt.Println(menuOptions)
   fmt.Println(menuItens)
}

