package main

import (
	"fmt"
	"os/exec"
	"strings"
	"time"
)

type Message struct {
	Type      string `json:"type"`
	DeviceID  string `json:"device_id"`
	MessageID string `json:"message_id"`
	Content   string `json:"content"`
}

func getClipBoard() (string, error) {
	cmd := exec.Command("wl-paste", "--no-newline")
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return string(out), nil
}
func setClipBoard(text string) error {
	cmd := exec.Command("wl-copy")
	cmd.Stdin = strings.NewReader(text)
	return cmd.Run()
}

func main() {
	var previous string
	for {
		current, err := getClipBoard()
		if err != nil {
			fmt.Println("clipboard error", err)
			time.Sleep(time.Second)
			continue
		}
		if current != previous {
			fmt.Println("Clipboard changed:")
			fmt.Printf("%q\n", current)
			previous = current

		}
		time.Sleep(200 * time.Millisecond)
	}
}
