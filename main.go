package main

import (
	"fmt"
	"os/exec"

	"github.com/pterm/pterm"
)

func shutdown() {
	if err := exec.Command("cmd", "/C", "shutdown", "/s").Run(); err != nil {
		fmt.Println("Failed to shutdown:", err)
	}

}
func restart() {
	if err := exec.Command("cmd", "/C", "shutdown", "/r"); err != nil {
		fmt.Println(err)
	}
}

func main() {
	//	options := []string{"Shutdown", "Restart"}

	//switch options{
	//case "Shutdown":

	//case "Restart":
	//	exec.Command("cmd", "/C", "shutdown", "/r")
	//}

	a, _ := pterm.DefaultInteractiveMultiselect.
		WithOptions([]string{"shutdown", "restart"}).
		Show()
	//confirm for leaving a, _ := pterm.DefaultInteractiveConfirm.Show()

	pterm.DefaultSpinner.Start()

	pterm.Info.Println("Updating the system..")
	pterm.Info.Println("Waiting for a answer..")

	pterm.DefaultSpinner.Stop()
	pterm.Info.Println("%T", a)
	shutdown()
}

