package day19

import (
	"aoc2023/common"
	"fmt"
	"regexp"
)

type Condition interface {
	Apply(part Part) bool
}

var reCondition *regexp.Regexp = regexp.MustCompile(`([xmas])([<>])(\d+)`)

func ParseCondition(str string) Condition {
	tokens := reCondition.FindStringSubmatch(str)
	if len(tokens) != 4 {
		panic(fmt.Sprintf("wrong condition, len=%d, str=%s", len(tokens), str))
	}
	return NewSmallerLarger(tokens[1], tokens[2], common.StringToInt(tokens[3]))
}

// ****

type SmallerLarger struct {
	category string // [xmas]
	sign     string // [<>]
	value    int
}

var _ Condition = SmallerLarger{}

func NewSmallerLarger(cat, sign string, value int) SmallerLarger {
	return SmallerLarger{
		category: cat,
		sign:     sign,
		value:    value,
	}
}

func (sl SmallerLarger) Apply(part Part) bool {
	partValue := part.xmas[sl.category]

	switch sl.sign {
	case "<":
		return partValue < sl.value
	case ">":
		return partValue > sl.value
	default:
		panic("wrong sign")
	}
}
