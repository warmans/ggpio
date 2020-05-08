package main

import (
	"flag"
	"github.com/jacobsa/go-serial/serial"
	"github.com/warmans/ggpio"
	"github.com/warmans/go-rtk"
	"log"
	"time"
)

func main() {

	adapter := flag.String("adapter", "rtk", "Defines which adapter to use for GPIO")
	flag.Parse()

	if adapter == nil {
		log.Fatal("Adapter must be set")
	}

	var gpio ggpio.GPIO
	switch *adapter {
	case "rtk":
		port, err := serial.Open(rtk.SerialOptions("/dev/serial/by-path/pci-0000:00:06.0-usb-0:2:1.0-port0"))
		if err != nil {
			log.Fatalf("serial.Open: %v", err)
		}
		defer port.Close()

		gpio = ggpio.NewRTk(rtk.NewGPIOClient(port))
	case "rpio":
		var err error
		gpio, err = ggpio.NewRPIO()
		if err != nil {
			log.Fatalf("failed to create rpio instance: %s" + err.Error())
		}
	default:
		panic("unknown adapter: " + *adapter)
	}
	defer gpio.Close()

	const pin = 10

	if err := gpio.Configure(
		pin,
		ggpio.SetPinMode(ggpio.PinModeWritable),
		ggpio.SetPull(ggpio.PullDown),
		ggpio.SetInitialState(ggpio.PinStateLow),
	); err != nil {
		log.Fatalf("setup failed %v", err)
	}

	state := ggpio.PinStateLow
	for i := 0; i < 6; i++ {
		if i%2 == 0 {
			state = ggpio.PinStateLow
		} else {
			state = ggpio.PinStateHigh
		}
		log.Printf("Turning on: %v\n", state)
		if err := gpio.Write(pin, state); err != nil {
			log.Fatalf("output failed %v", err)
		}
		time.Sleep(time.Second)
	}
}
