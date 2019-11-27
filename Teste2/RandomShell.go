package main

import (
	"fmt"
	"os/exec"
)

func main() {
	out, _ := exec.Command(test).Output()
	output := string(out[:])
	fmt.Println("qass", string(out))
}
