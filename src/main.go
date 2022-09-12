package main

import (
	"fmt"
	"os/exec"
	"strings"
	"time"
	//"time"
)

func main() {

	for s := 1; s < 11; s++ {
		time.Sleep(1 * time.Second)
		fmt.Printf("Shutdown in %v\n", s-11)

	}
//shutdown
	if err := exec.Command("cmd", "/C", "shutdown", "/s").Run(); err != nil {
		fmt.Println("Failed to shutdown:", err)
	}
  //if user types n:
	var result string
	fmt.Println("Type n for undo.")
	fmt.Scan(&result)
	if strings.EqualFold("n", "N") {
    //cancel shutdown
		if err := exec.Command("cmd", "/C", "shutdown", "/a").Run(); err != nil {
			fmt.Println("Failed to cancel:", err)

		}
	}

}
//shutdown -r for running
//shutdown -a for cancelling
