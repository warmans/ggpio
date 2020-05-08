package ggpio

import "github.com/warmans/go-rtk"

func NewRTk(client *rtk.GPIOClient) *Rtk {
	return &Rtk{client: client}
}

type Rtk struct {
	client *rtk.GPIOClient
}

func (r *Rtk) Configure(pin uint8, opts ...SetupOpt) error {
	return r.client.Setup(pin, rtkSetupOpts(opts)...)
}

func (r *Rtk) Read(pin uint8) (PinState, error) {
	state, err := r.client.Input(pin)
	return pinStateFromRtk(state), err
}

func (r *Rtk) Write(pin uint8, state PinState) error {
	return r.client.Output(pin, rtkPinState(state))
}

func (r *Rtk) Close() error {
	r.client.Close()
	return nil
}

func rtkSetupOpts(opts []SetupOpt) []rtk.SetupOpt {

	setup := applyOpts(opts...)

	var rtkOpts []rtk.SetupOpt
	if setup.pinMode != nil {
		rtkOpts = append(rtkOpts, rtk.InitialPinMode(rtkPinMode(*setup.pinMode)))
	}
	if setup.initialState != nil {
		rtkOpts = append(rtkOpts, rtk.InitialState(rtkPinState(*setup.initialState)))
	}
	if setup.pull != nil {
		rtkOpts = append(rtkOpts, rtk.Pull(rtkPull(*setup.pull)))
	}
	return rtkOpts
}

func rtkPinMode(mode PinMode) rtk.PinMode {
	switch mode {
	case PinModeReadable:
		return rtk.PinModeInput
	case PinModeWritable:
		return rtk.PinModeOutput
	}
	panic("unknown pin mode " + string(mode))
}

func rtkPinState(state PinState) rtk.PinState {
	switch state {
	case PinStateLow:
		return rtk.PinStateLow
	case PinStateHigh:
		return rtk.PinStateHigh
	}
	panic("unknown pin state " + string(state))
}

func rtkPull(pull Pull) rtk.PullMode {
	switch pull {
	case PullUp:
		return rtk.PullUp
	case PullDown:
		return rtk.PullDown
	case PullNone:
		return rtk.PullNone
	}
	panic("unknown pull " + string(pull))
}

func pinStateFromRtk(state rtk.PinState) PinState {
	switch state {
	case rtk.PinStateHigh:
		return PinStateHigh
	case rtk.PinStateLow:
		return PinStateLow
	}
	panic("unknown pin state " + string(state))
}
