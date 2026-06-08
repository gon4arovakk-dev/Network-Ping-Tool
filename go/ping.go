package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run ping.go <host>")
		os.Exit(1)
	}
	host := os.Args[1]
	cmd := exec.Command("ping", "-c", "4", host)
	out, err := cmd.Output()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
	fmt.Print(string(out))
}
