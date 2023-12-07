package day07

func CardToInt(r rune) int {
	switch r {
	case 'A':
		return 14
	case 'K':
		return 13
	case 'Q':
		return 12
	case 'J':
		return 11
	case 'T':
		return 10
	default:
		return int(r - '0')
	}
}

func CardToIntWithJoker(r rune) int {
	i := CardToInt(r)
	if i == 11 {
		return 0
	}
	return i
}

/*
five of a kind - 7
four of a kind - 6
full house - 5
3 of a kind - 4
two pair - 3
one pair - 2
high card - 1
*/
func CardsToType(cards []int) int {
	hist := make(map[int]int, 0)
	for _, c := range cards {
		hist[c] += 1
	}
	// fmt.Println(hist)
	if len(hist) == 1 {
		return 7 // 5 of a kind
	}
	if len(hist) == 5 {
		return 1 // high card
	}
	if len(hist) == 4 {
		return 2 // one pair
	}
	if len(hist) == 2 {
		// 4 of a kind OR full house
		// if hist[0] == 1 || hist[0] == 4 {
		for _, v := range hist {
			if v == 1 || v == 4 {
				return 6 // 4 of a kind
			} else {
				return 5 // full house
			}
		}
	}

	// 3 of a kind OR 2 pair
	for _, v := range hist {
		if v == 2 {
			return 3 // 2 pair
		}
		if v == 3 {
			return 4 // 3 of a kind
		}
	}

	panic("cannot find a type?")
}

/*
five of a kind - 7
four of a kind - 6
full house - 5
3 of a kind - 4
two pair - 3
one pair - 2
high card - 1
*/
func CardsToTypeWithJoker(cards []int) int {
	nrJokers := 0
	hist := make(map[int]int, 0) // for non jokers
	for _, c := range cards {
		if c == 0 {
			nrJokers++
		} else {
			hist[c]++
		}
	}

	if nrJokers == 0 {
		return CardsToType(cards)
	}

	if nrJokers == 1 {
		switch len(hist) {
		case 1:
			return 7 // 5 of a kind
		case 4:
			return 2 // 1 pair
		case 3:
			return 4 // 3 of a kind
		}
		// case 2
		for _, v := range hist {
			if v == 1 || v == 3 {
				return 6 // 4 of a kind
			} else {
				return 5 // full house
			}
		}
	}

	if nrJokers == 2 {
		switch len(hist) {
		case 1:
			return 7
		case 2:
			return 6
		case 3:
			return 4
		default:
			panic("on 2 Jokers")
		}
	}

	if nrJokers == 3 {
		switch len(hist) {
		case 1:
			return 7
		case 2:
			return 6
		default:
			panic(" on 3 Jokers")
		}
	}

	return 7
}
