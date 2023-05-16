package utility

import (
	"fmt"
	"log"
	"os/exec"
)

func ExecShell(dep string, updateFlag bool){
	log.Printf("Attempted to get %v", dep)
	var args []string
	if updateFlag {
		args = []string{"get", "-u", dep}
	} else {
		args = []string{"get", dep}
	}
	
	cmd := exec.Command("go", args...)
	// cmd := exec.Command("go", "get", dep)
	out, err := cmd.Output()
	if err != nil {
		log.Fatalf("Failed to exec command: %v", err)
	}
	if string(out) != ""{
		fmt.Println("Command Output:", string(out))
	}
	log.Printf("Successfully added %v", dep)
}