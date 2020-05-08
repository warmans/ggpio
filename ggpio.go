package ggpio

type Pull string

func (m Pull) Ptr() *Pull {
	return &m
}

const (
	PullUp   Pull = "U"
	PullDown Pull = "D"
	PullNone Pull = "N"
)

type PinMode string

func (m PinMode) Ptr() *PinMode {
	return &m
}

const (
	PinModeReadable PinMode = "I"
	PinModeWritable PinMode = "O"
)

type PinState string

func (s PinState) Ptr() *PinState {
	return &s
}

const (
	PinStateHigh PinState = "1"
	PinStateLow  PinState = "0"
)

type setupOpts struct {
	pull         *Pull
	initialState *PinState
	pinMode      *PinMode
}

type SetupOpt func(opts *setupOpts)

func SetPull(mode Pull) SetupOpt {
	return func(opts *setupOpts) {
		opts.pull = mode.Ptr()
	}
}

func SetInitialState(state PinState) SetupOpt {
	return func(opts *setupOpts) {
		opts.initialState = state.Ptr()
	}
}

func SetPinMode(mode PinMode) SetupOpt {
	return func(opts *setupOpts) {
		opts.pinMode = mode.Ptr()
	}
}

func applyOpts(opts ...SetupOpt) *setupOpts {
	setup := &setupOpts{}
	for _, opt := range opts {
		opt(setup)
	}
	return setup
}

type GPIO interface {
	Configure(pin uint8, opts ...SetupOpt) error
	Read(pin uint8) (PinState, error)
	Write(pin uint8, state PinState) error
	Close() error
}
