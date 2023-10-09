package main

import (
	"machine"
	"time"

	"github.com/pavelanni/tinygo-drivers/dfplayermini"
)

func main() {
	time.Sleep(2 * time.Second)
	d := dfplayermini.New(machine.UART0, machine.UART0_TX_PIN, machine.UART0_RX_PIN)
	d.Configure()
	time.Sleep(2 * time.Second)
	d.QueryStatus()
	time.Sleep(1 * time.Second)
	d.Read()
	d.Volume(20)
	time.Sleep(2 * time.Second)
	//d.VolumeUp()
	//d.VolumeDown()
	//d.Next()
	//d.Previous()
	d.Play(1)
	d.Read()
	//time.Sleep(2 * time.Second)
	d.Play(2)
	d.Read()
	//time.Sleep(2 * time.Second)
	d.Play(3)
	d.Read()
}
