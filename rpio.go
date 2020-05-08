package ggpio

import (
	"fmt"
	"github.com/stianeikeland/go-rpio/v4"
)

func NewRPIO() (*RPIO, error) {
	if err := rpio.Open(); err != nil {
		return nil, err
	}
	return &RPIO{}, nil
}

type RPIO struct {
}

func (R *RPIO) Configure(pin uint8, opts ...SetupOpt) error {

	cfg := applyOpts(opts...)
	if cfg.pull != nil {
		rpio.PullMode(rpioPin(pin), rpioPull(*cfg.pull))
	}
	if cfg.initialState != nil {
		if err := R.Write(pin, *cfg.initialState); err != nil {
			return err
		}
	}
	if cfg.pinMode != nil {
		rpio.PinMode(rpioPin(pin), rpioPinMode(*cfg.pinMode))
	}
	return nil
}

func (R *RPIO) Read(pin uint8) (PinState, error) {
	return stateFromRpio(rpio.ReadPin(rpioPin(pin))), nil
}

func (R *RPIO) Write(pin uint8, state PinState) error {
	rpio.WritePin(rpioPin(pin), rpioState(state))
	return nil
}

func (R *RPIO) Close() error {
	return rpio.Close()
}

func rpioPull(pull Pull) rpio.Pull {
	switch pull {
	case PullNone:
		return rpio.PullOff
	case PullUp:
		return rpio.PullUp
	case PullDown:
		return rpio.PullDown
	}
	panic("unknown pull: " + string(pull))
}

func rpioPinMode(mode PinMode) rpio.Mode {
	switch mode {
	case PinModeReadable:
		return rpio.Input
	case PinModeWritable:
		return rpio.Output
	}
	panic("unknown pin mode: " + string(mode))
}

func rpioPin(pin uint8) rpio.Pin {
	return rpio.Pin(pin)
}

func rpioState(state PinState) rpio.State {
	switch state {
	case PinStateLow:
		return rpio.Low
	case PinStateHigh:
		return rpio.High
	}
	panic("unknown pin state: " + string(state))
}

func stateFromRpio(state rpio.State) PinState {
	switch state {
	case rpio.Low:
		return PinStateLow
	case rpio.High:
		return PinStateHigh
	}
	panic(fmt.Sprintf("unknown pin state: %v", state))
}
