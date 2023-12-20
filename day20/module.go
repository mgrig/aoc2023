package day20

const (
	BROADCAST int = 0
	FLIPFLOP  int = 1
	CONJ      int = 2
	UNTYPED   int = 3
)

type ModuleType int

type Module interface {
	ProcessPulse(pulse Pulse) (emitted []Pulse)
	GetName() string
	GetDest() []string
	AddIn(from string)
}

func emit(pulseType PulseType, from string, dest []string) []Pulse {
	ret := make([]Pulse, len(dest))
	for i, dst := range dest {
		ret[i] = NewPulse(from, dst, pulseType)
	}
	return ret
}

// ****
const (
	FF_OFF int = 0
	FF_ON  int = 1
)

type FlipFlop struct {
	name  string
	dest  []string
	state int
}

var _ Module = &FlipFlop{}

func NewFlipFlop(name string, dest []string) Module {
	return &FlipFlop{
		name:  name,
		dest:  dest,
		state: FF_OFF,
	}
}

func (ff *FlipFlop) GetName() string {
	return ff.name
}

func (ff *FlipFlop) GetDest() []string {
	return ff.dest
}

func (ff *FlipFlop) AddIn(from string) {}

func (ff *FlipFlop) ProcessPulse(pulse Pulse) (emitted []Pulse) {
	if pulse.typ == LOW {
		switch ff.state {
		case FF_OFF:
			emitted = emit(HIGH, ff.name, ff.dest)
		case FF_ON:
			emitted = emit(LOW, ff.name, ff.dest)
		}
		ff.state = 1 - ff.state
	}
	return
}

// ****

type Conj struct {
	name string
	dest []string
	in   map[string]PulseType // remember the type of the most recent pulse received from each of their connected input modules
}

var _ Module = &Conj{}

func NewConj(name string, dest []string) Module {
	return &Conj{
		name: name,
		dest: dest,
		in:   make(map[string]PulseType),
	}
}

func (c *Conj) AddIn(from string) {
	c.in[from] = LOW
}

func (c *Conj) GetName() string {
	return c.name
}

func (c *Conj) GetDest() []string {
	return c.dest
}

func (c *Conj) ProcessPulse(pulse Pulse) (emitted []Pulse) {
	c.in[pulse.from] = pulse.typ

	allHigh := true
	for _, v := range c.in {
		if v == LOW {
			allHigh = false
			break
		}
	}

	if allHigh {
		emitted = emit(LOW, c.name, c.dest)
	} else {
		emitted = emit(HIGH, c.name, c.dest)
	}
	return
}

// ****

type Broadcast struct {
	name string
	dest []string
}

var _ Module = &Broadcast{}

func NewBroadcast(name string, dest []string) Module {
	return &Broadcast{
		name: name,
		dest: dest,
	}
}

func (b *Broadcast) GetName() string {
	return b.name
}

func (b *Broadcast) GetDest() []string {
	return b.dest
}

func (b *Broadcast) AddIn(from string) {}

func (b *Broadcast) ProcessPulse(pulse Pulse) (emitted []Pulse) {
	return emit(pulse.typ, b.name, b.dest)
}

// ****

type Untyped struct {
	name string
}

var _ Module = &Untyped{}

func NewUntyped(name string) Module {
	return &Untyped{
		name: name,
	}
}

func (u *Untyped) GetName() string {
	return u.name
}

func (u *Untyped) GetDest() []string {
	return make([]string, 0)
}

func (u *Untyped) AddIn(from string) {}

func (u *Untyped) ProcessPulse(pulse Pulse) (emitted []Pulse) {
	return
}
