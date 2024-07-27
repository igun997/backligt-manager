package main

import (
	"fmt"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"os/exec"
)

var device string

func setBacklight(level string) {
	if device == "" {
		fmt.Println("Error: Backlight device not set")
		return
	}

	cmd := exec.Command("sudo", "brightnessctl", "--device="+device, "s", level)
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error:", err)
		fmt.Println("Output:", string(output))
	}
}

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Backlight Manager with BPM Detection")

	deviceInput := widget.NewEntry()
	deviceInput.SetPlaceHolder("Enter backlight device (e.g., tpacpi::kbd_backlight)")

	label := widget.NewLabel("Set Keyboard Backlight Level")
	levelZero := widget.NewButton("Off", func() { setBacklight("0") })
	levelLow := widget.NewButton("Low", func() { setBacklight("1") })
	levelHigh := widget.NewButton("High", func() { setBacklight("2") })

	startButton := widget.NewButton("Set Device", func() {
		device = deviceInput.Text
		fmt.Println("Device set to:", device)
	})

	myWindow.SetContent(container.NewVBox(
		widget.NewLabel("Enter Backlight Device"),
		deviceInput,
		startButton,
		label,
		levelZero,
		levelLow,
		levelHigh,
	))

	myWindow.ShowAndRun()
}
