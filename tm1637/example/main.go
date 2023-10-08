package main

import (
	"machine"
	"time"

	"github.com/pavelanni/tinygo-drivers/tm1637"
)

func main() {

	tm := tm1637.New(machine.GP2, machine.GP3, 7) // clk, dio, brightness
	tm.Configure()

	tm.ClearDisplay()

	tm.DisplayText([]byte("Tiny"))
	time.Sleep(time.Millisecond * 1000)

	tm.ClearDisplay()

	tm.DisplayChr(byte('G'), 1)
	tm.DisplayDigit(0, 2) // looks like O
	time.Sleep(time.Millisecond * 1000)

	tm.DisplayClock(12, 59, true)

	for i := uint8(0); i < 8; i++ {
		tm.Brightness(i)
		time.Sleep(time.Millisecond * 200)
	}

	tm.DisplayPoint()
	time.Sleep(time.Second * 2)

	i := int16(0)
	tStart := time.Now()
	for {
		tm.DisplayNumber(i)
		//i++
		time.Sleep(time.Millisecond * 80)
		if i > 88 {
			break
		}

	}
	duration := time.Since(tStart)
	println("duration:", duration.Seconds())

}
