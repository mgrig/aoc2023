package day15

type Lens struct {
	label string
	focal int
}

func NewLens(label string, focus int) Lens {
	return Lens{
		label: label,
		focal: focus,
	}
}

// ****

type Box struct {
	lenses []Lens
}

func NewBox() *Box {
	return &Box{
		lenses: make([]Lens, 0),
	}
}

func (b *Box) Append(newLens Lens) (oldLens Lens, found bool) {
	for i, lens := range b.lenses {
		if lens.label == newLens.label {
			oldLens = lens
			b.lenses[i] = newLens
			return oldLens, true
		}
	}
	b.lenses = append(b.lenses, newLens)
	return oldLens, false
}

func (b *Box) Remove(targetLabel string) (oldLens Lens, found bool) {
	for i, lens := range b.lenses {
		if lens.label == targetLabel {
			oldLens = lens
			b.lenses = append(b.lenses[:i], b.lenses[i+1:]...)
			return oldLens, true
		}
	}
	return oldLens, false
}
