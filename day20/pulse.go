package day20

const (
	LOW  = 0
	HIGH = 1
)

type PulseType int

type Pulse struct {
	from, to string
	typ      PulseType
}

func NewPulse(from, to string, typ PulseType) Pulse {
	return Pulse{
		from: from,
		to:   to,
		typ:  typ,
	}
}
