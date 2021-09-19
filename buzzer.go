package laser

import (
	"github.com/stianeikeland/go-rpio/v4"
	"time"
)

type Buzzer struct {
	pin rpio.Pin
}

func NewBuzzer(pin rpio.Pin) *Buzzer {
	return &Buzzer{pin}
}

func (b *Buzzer) Buzz() error {
	err := rpio.Open()
	if err != nil {
		return err
	}
	defer func() {
		err := rpio.Close()
		if err != nil {
			panic(err)
		}
	}()

	b.pin.Output()
	b.pin.High()
	time.Sleep(time.Second * 2)
	b.pin.Low()

	return nil
}
