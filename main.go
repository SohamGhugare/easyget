/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"easyget/cmd"
	"log"
)

func init(){
	log.SetPrefix("[EASYGET] ")
	log.SetFlags(0)
}

func main() {
	cmd.Execute()
}
