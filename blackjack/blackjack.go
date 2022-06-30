package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch card {
	case "ace":
		return 11
	case "king", "queen", "jack", "ten":
		return 10
	case "nine":
		return 9
	case "eight":
		return 8
	case "seven":
		return 7
	case "six":
		return 6
	case "five":
		return 5
	case "four":
		return 4
	case "three":
		return 3
	case "two":
		return 2
	case "one":
		return 1
	}

	return 0
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	stand := "S"
	hit := "H"
	split := "P"
	win := "W"

	sum := ParseCard(card1) + ParseCard(card2)
	dealerScore := ParseCard(dealerCard)

	isPairAce := card1 == "ace" && card2 == "ace"
	isDealerPotentialWin := dealerScore >= 10
	isBJ := sum == 21

	switch {
	case isPairAce:
		return split
	case isBJ && !isDealerPotentialWin:
		return win
	case isBJ && isDealerPotentialWin:
		return stand
	case sum >= 17 && sum <= 20:
		return stand
	case sum >= 12 && sum <= 16 && dealerScore >= 7:
		return hit
	case sum >= 12 && sum <= 16 && dealerScore < 7:
		return stand
	}

	return hit
}
