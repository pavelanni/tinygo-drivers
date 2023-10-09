package dfplayermini

import (
	"encoding/binary"
	"machine"
	"strconv"
)

type Device struct {
	uart  *machine.UART
	txPin machine.Pin
	rxPin machine.Pin
}

func printBytes(data []byte) {
	for _, b := range data {
		print(strconv.FormatInt(int64(b), 16), " ")
	}
	println()
}

func New(uart *machine.UART, txPin machine.Pin, rxPin machine.Pin) *Device {
	return &Device{
		uart:  uart,
		txPin: txPin,
		rxPin: rxPin,
	}
}

func (d *Device) Configure() {
	println("configuring uart")
	err := d.uart.Configure(machine.UARTConfig{
		BaudRate: 9600,
		TX:       d.txPin,
		RX:       d.rxPin,
	})
	if err != nil {
		panic(err)
	}
	println("uart configured")
}

func (d *Device) Read() ([]byte, error) {
	data := make([]byte, 0)

	for {
		if d.uart.Buffered() > 0 {
			inByte, err := d.uart.ReadByte()
			if err != nil {
				return []byte{}, err
			}
			data = append(data, inByte)
			if inByte == 0xef {
				break
			}
		}
	}
	print("received message: ")
	printBytes(data)
	return data, nil
}

func (d *Device) QueryStatus() {
	message := []byte{0x7e, 0xff, 0x06, 0x42, 0x00, 0x00, 0x00, 0xef}
	print("sending message: ")
	printBytes(message)
	_, err := d.uart.Write(message)
	if err != nil {
		panic(err)
	}
}

func (d *Device) Volume(volume uint8) {
	message := []byte{0x7e, 0xff, 0x06, 0x06, 0x00, 0x00}
	message = append(message, volume)
	message = append(message, 0xef)
	print("sending message: ")
	printBytes(message)
	_, err := d.uart.Write(message)
	if err != nil {
		panic(err)
	}
}

func (d *Device) VolumeUp() {
	d.uart.Write([]byte{0x7e, 0xff, 0x06, 0x04, 0x00, 0x00, 0x00, 0xef})
}

func (d *Device) VolumeDown() {
	d.uart.Write([]byte{0x7e, 0xff, 0x06, 0x05, 0x00, 0x00, 0x00, 0xef})
}

func (d *Device) Next() {
	d.uart.Write([]byte{0x7e, 0xff, 0x06, 0x01, 0x00, 0x00, 0x00, 0xef})
}

func (d *Device) Previous() {
	d.uart.Write([]byte{0x7e, 0xff, 0x06, 0x02, 0x00, 0x00, 0x00, 0xef})
}

func (d *Device) Play(track uint16) {
	bs := make([]byte, 2)
	binary.BigEndian.PutUint16(bs, track)
	message := []byte{0x7e, 0xff, 0x06, 0x03, 0x00}
	message = append(message, bs...)
	message = append(message, 0xef)
	print("sending message: ")
	printBytes(message)
	_, err := d.uart.Write(message)
	if err != nil {
		panic(err)
	}
}
