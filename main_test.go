package main

import (
	"fmt"
	"testing"
)

func TestLoadConfig(t *testing.T) {

   config:= loadConfig("my-tray-menu.yaml")
   fmt.Println(config)
}

