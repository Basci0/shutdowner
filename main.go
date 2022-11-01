package main

import (
	"os/exec"
	"time"
	"github.com/pterm/pterm"
)

func main() {

	optionsSelected, _ := pterm.DefaultInteractiveMultiselect.
		WithOptions([]string{"shutdown", "restart"}).
		Show()

	pterm.DefaultSpinner.Start()

	switch optionsSelected[0] {
	case "shutdown":
		time.Sleep(1 * time.Second)
		pterm.Info.Println("Updating the drivers..")

		time.Sleep(1 * time.Second)
		pterm.Info.Println("Searching for errors ..")
		time.Sleep(1 * time.Second)
		pterm.Info.Println("Shutting down..")

		pterm.DefaultSpinner.Stop()

		if err := exec.Command("cmd", "/C", "shutdown", "/s").Run(); err != nil {
			pterm.Error.Println("Failed to shutdown:", err)
		}
		pterm.DefaultSpinner.Start()

	case "restart":
		time.Sleep(1 * time.Second)
		pterm.Info.Println("Updating the drivers..")

		time.Sleep(1 * time.Second)
		pterm.Info.Println("Searching for errors ..")
		time.Sleep(1 * time.Second)
		pterm.Info.Println("Restarting..")
		pterm.DefaultSpinner.Stop()

		if err := exec.Command("cmd", "/C", "shutdown", "/r").Run(); err != nil {
			pterm.Error.Println("Failed to restart:", err)
		}
	}

}

