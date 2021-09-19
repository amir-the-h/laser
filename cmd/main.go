package main

import "github.com/amir-the-h/raspberry-pi/laser"

func main() {
	buzzer := laser.NewBuzzer(17)
	if err := buzzer.Buzz(); err != nil {
		panic(err)
	}
}
