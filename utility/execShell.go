package utility

import (
	"fmt"
	"log"
	"os/exec"
)

func ExecShell(args string){
	cmd := exec.Command("go", "get", args)
	out, err := cmd.Output()
	if err != nil {
		log.Fatalf("Failed to exec command: %v", err)
	}
	if string(out) != ""{
		fmt.Println("Command Output:", string(out))
	}
	
}