package main

import (
	"fmt"
	"log"
	"os/exec"
	"time"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/gordonklaus/portaudio"
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

func detectBPM() (float64, error) {
	portaudio.Initialize()
	defer portaudio.Terminate()

	in := make([]float32, 64)
	stream, err := portaudio.OpenDefaultStream(1, 0, 44100, len(in), in)
	if err != nil {
		return 0, err
	}
	defer stream.Close()

	err = stream.Start()
	if err != nil {
		return 0, err
	}
	defer stream.Stop()

	var bpm float64
	for i := 0; i < 10; i++ {
		err = stream.Read()
		if err != nil {
			return 0, err
		}

		// Here, you would process the `in` buffer to detect the BPM
		// For demonstration purposes, we'll just simulate a BPM value
		bpm = 120.0
		time.Sleep(1 * time.Second)
	}

	return bpm, nil
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

	bpmLabel := widget.NewLabel("BPM: 0")
	bpmButton := widget.NewButton("Detect BPM", func() {
		bpm, err := detectBPM()
		if err != nil {
			log.Println("Error detecting BPM:", err)
			bpmLabel.SetText("BPM: Error")
			return
		}
		bpmLabel.SetText(fmt.Sprintf("BPM: %.2f", bpm))
	})

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
		bpmButton,
		bpmLabel,
	))

	myWindow.ShowAndRun()
}
