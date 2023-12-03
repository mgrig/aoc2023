package day03

import (
	"fmt"
	"math"
)

func Part1(lines []string) int {
	numbers := make([]PosNumber, 0)
	// symbols := make([]Coord, 0)
	symbolGrid := NewSymbolGrid(len(lines))

	var currentNumber *PosNumber = nil
	for r, line := range lines {
		for c, cell := range line {
			if currentNumber == nil {
				if isEmptyCell(cell) {
					// nop
				} else if isDigitCell(cell) {
					currentNumber = NewPosNumber(cellToInt(cell), NewCoord(r, c))
				} else {
					// some symbol
					// symbols = append(symbols, NewCoord(r, c))
					symbolGrid.SetSymbolAt(NewCoord(r, c))
				}
			} else { // already parsing a number
				if isEmptyCell(cell) {
					// ending current number
					numbers = append(numbers, *currentNumber)
					currentNumber = nil
				} else if isDigitCell(cell) {
					// update current number
					currentNumber.AppendDigit(cellToInt(cell))
				} else {
					// some symbol
					// symbols = append(symbols, NewCoord(r, c))
					symbolGrid.SetSymbolAt(NewCoord(r, c))

					// ending current number
					numbers = append(numbers, *currentNumber)
					currentNumber = nil
				}
			}
		}
		if currentNumber != nil {
			numbers = append(numbers, *currentNumber)
			currentNumber = nil
		}
	}

	// fmt.Println(len(numbers), numbers)
	// fmt.Println(len(symbols), symbols)
	// fmt.Println(symbolGrid)

	sum := 0
	for _, pn := range numbers {
		if isAdjacentToSymbol(pn, symbolGrid) {
			fmt.Println(pn.value, nrDigitsOfInt(pn.value))
			sum += pn.value
			continue
		} else {
			fmt.Println("ignoring", pn.value)
		}
	}

	return sum
}

func Part2(lines []string) int {
	numbers := make([]PosNumber, 0)
	symbols := make([]Coord, 0)
	// symbolGrid := NewSymbolGrid(len(lines))

	var currentNumber *PosNumber = nil
	for r, line := range lines {
		for c, cell := range line {
			if currentNumber == nil {
				if isEmptyCell(cell) {
					// nop
				} else if isDigitCell(cell) {
					currentNumber = NewPosNumber(cellToInt(cell), NewCoord(r, c))
				} else if isStarCell(cell) {
					// * symbol
					symbols = append(symbols, NewCoord(r, c))
					// symbolGrid.SetSymbolAt(NewCoord(r, c))
				} else {
					// nop
				}
			} else { // already parsing a number
				if isEmptyCell(cell) {
					// ending current number
					numbers = append(numbers, *currentNumber)
					currentNumber = nil
				} else if isDigitCell(cell) {
					// update current number
					currentNumber.AppendDigit(cellToInt(cell))
				} else if isStarCell(cell) {
					// some symbol
					symbols = append(symbols, NewCoord(r, c))
					// symbolGrid.SetSymbolAt(NewCoord(r, c))

					// ending current number
					numbers = append(numbers, *currentNumber)
					currentNumber = nil
				} else {
					// ending current number
					numbers = append(numbers, *currentNumber)
					currentNumber = nil
				}
			}
		}
		if currentNumber != nil {
			numbers = append(numbers, *currentNumber)
			currentNumber = nil
		}
	}

	// fmt.Println(len(numbers), numbers)
	// fmt.Println(len(symbols), symbols)
	// fmt.Println(symbolGrid)

	// N := len(lines)
	sum := 0
	for _, symbolCoord := range symbols {
		has2, matchNumbers := symbolHas2AdjNumbers(symbolCoord, &numbers)
		if has2 {
			gearRatio := matchNumbers[0].value * matchNumbers[1].value
			sum += gearRatio
		}
	}

	return sum
}

func symbolHas2AdjNumbers(symbolCoord Coord, numbers *[]PosNumber) (match bool, retNumbers []PosNumber) {
	for _, pn := range *numbers {
		if isAdjacentToNumber(symbolCoord, pn) {
			retNumbers = append(retNumbers, pn)
			if len(retNumbers) > 2 {
				return false, nil
			}
		}
	}
	if len(retNumbers) == 2 {
		return true, retNumbers
	}
	return false, nil
}

func isAdjacentToNumber(coord Coord, pn PosNumber) bool {
	nrDigits := nrDigitsOfInt(pn.value)

	if coord.Equals(pn.startPos.r, pn.startPos.c-1) || coord.Equals(pn.startPos.r, pn.startPos.c+nrDigits) {
		return true
	}

	if int(math.Abs(float64(coord.r-pn.startPos.r))) > 1 {
		return false
	}

	for c := pn.startPos.c - 1; c <= pn.startPos.c+nrDigits; c++ {
		if coord.Equals(pn.startPos.r-1, c) || coord.Equals(pn.startPos.r+1, c) {
			return true
		}
	}

	return false
}

func isAdjacentToSymbol(pn PosNumber, sg *SymbolGrid) bool {
	nrDigits := nrDigitsOfInt(pn.value)

	if sg.IsSymbolAt(pn.startPos.r, pn.startPos.c-1) {
		return true
	}

	if sg.IsSymbolAt(pn.startPos.r, pn.startPos.c+nrDigits) {
		return true
	}

	for c := pn.startPos.c - 1; c <= pn.startPos.c+nrDigits; c++ {
		if sg.IsSymbolAt(pn.startPos.r-1, c) || sg.IsSymbolAt(pn.startPos.r+1, c) {
			return true
		}
	}

	return false
}

func nrDigitsOfInt(n int) int {
	return int(math.Floor(math.Log10(float64(n)))) + 1
}

func isEmptyCell(r rune) bool {
	return r == '.'
}

func isDigitCell(r rune) bool {
	return r >= '0' && r <= '9'
}

func isStarCell(r rune) bool {
	return r == '*'
}

func cellToInt(r rune) int {
	return int(r - '0')
}
