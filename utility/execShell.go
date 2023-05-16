package utility

import (
	"fmt"
	"log"
	"os/exec"
)

func ExecShell(args string){
	log.Printf("Attempted to get %v", args)
	cmd := exec.Command("go", "get", args)
	out, err := cmd.Output()
	if err != nil {
		log.Fatalf("Failed to exec command: %v", err)
	}
	if string(out) != ""{
		fmt.Println("Command Output:", string(out))
	}
	log.Printf("Successfully added %v", args)
}