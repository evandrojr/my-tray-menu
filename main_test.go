package main

import (
	"fmt"
	"testing"
)

// var programPath string

func TestLoadConfig(t *testing.T) {

	loadConfig("my-tray-menu.yaml")
	fmt.Println(menuOptions)
	fmt.Println(menuItens)
}

type z struct {
	label   string
	command string
}

func TestSlice(t *testing.T) {

	a := make([]string, 1)
	z := z{}
	alteraElemento(a, &z)
	fmt.Println(a, z)
}

func alteraElemento(e []string, z *z) {
	e[0] = "junior"
	z.command = "Ola"
}
