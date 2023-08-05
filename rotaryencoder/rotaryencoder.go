// Package rotaryencoder provides a simple interface to a rotary encoder.
package rotaryencoder

import (
	"machine"
)

var (
	states = []int8{0, -1, 1, 0, 1, 0, 0, -1, -1, 0, 0, 1, 0, 1, -1, 0}
)

// New creates a new rotary encoder.
func New(pinA, pinB machine.Pin) *Device {
	return &Device{pinA: pinA, pinB: pinB, oldAB: 0b00000011, value: 0, Dir: make(chan int, 8), Switch: make(chan bool)}
}

// Device represents a rotary encoder.
type Device struct {
	pinA machine.Pin
	pinB machine.Pin

	oldAB  int
	value  int
	Dir    chan int
	Switch chan bool
}

// Configure configures the rotary encoder.
func (enc *Device) Configure() {
	enc.pinA.Configure(machine.PinConfig{Mode: machine.PinInputPullup})
	enc.pinA.SetInterrupt(machine.PinRising|machine.PinFalling, enc.interrupt)

	enc.pinB.Configure(machine.PinConfig{Mode: machine.PinInputPullup})
	enc.pinB.SetInterrupt(machine.PinRising|machine.PinFalling, enc.interrupt)
}

func (enc *Device) interrupt(pin machine.Pin) {
	aHigh, bHigh := enc.pinA.Get(), enc.pinB.Get()
	enc.oldAB <<= 2
	if aHigh {
		enc.oldAB |= 1 << 1
	}
	if bHigh {
		enc.oldAB |= 1
	}

	enc.value += int(states[enc.oldAB&0x0f])

	// Each full click of the encoder generates 4 interupts.
	// Each interrupt add 1 or -1 to the value.
	// We send the direction to the channel only when we have a full click, i.e. 4 interrupts.
	if enc.value%4 == 0 {
		direction := enc.value / 4
		if direction != 0 {
			select { // non-blocking way of sending to channel
			case enc.Dir <- direction:
			default:
			}
		}
		enc.value = 0
	}
}

// Value returns the value of the rotary encoder.
func (enc *Device) Value() int {
	return enc.value / 4
}

// SetValue sets the value of the rotary encoder.
func (enc *Device) SetValue(v int) {
	enc.value = v * 4
}
