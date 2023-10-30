package main

import (
	"machine"
	"time"

	"github.com/pavelanni/tinygo-drivers/dfplayermini"
)

func main() {
	time.Sleep(2 * time.Second)
	d := dfplayermini.New(machine.UART1, machine.GP8, machine.GP9)
	d.Configure()
	time.Sleep(2 * time.Second)
	d.QueryStatus()
	time.Sleep(1 * time.Second)
	d.Read()
	d.Volume(20)
	time.Sleep(2 * time.Second)
	d.Play(1)
	d.Read()
	time.Sleep(2 * time.Second)
	d.Play(2)
	d.Read()
	time.Sleep(2 * time.Second)
	d.Play(3)
	d.Read()
	time.Sleep(2 * time.Second)
	d.Read()
	for i := 0; i < 3; i++ {
		d.Next()
		d.Read()
	}
	time.Sleep(2 * time.Second)
	for i := 0; i < 6; i++ {
		d.Play(5)
		time.Sleep(1000 * time.Millisecond)
	}
	time.Sleep(2 * time.Second)
}
