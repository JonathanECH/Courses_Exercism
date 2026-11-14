package blackjack
import "strings"
// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	card = strings.ToLower(card)
	switch card {
	case "two":
		return 2
	case "three":
		return 3
	case "four":
		return 4
	case "five":
		return 5
	case "six":
		return 6
	case "seven":
		return 7
	case "eight":
		return 8
	case "nine":
		return 9
	case "ten", "jack",
		"queen", "king":
		return 10
	case "ace":
		return 11
	default:
		return 0
	}
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	myScore := ParseCard(card1) + ParseCard(card2)
	dealearScore := ParseCard(dealerCard)
	switch {
	case myScore == 22:
		return "P"
	case myScore == 21:
		if dealearScore < 10 {
			return "W"
		}
		return "S"
	case myScore >= 17 && myScore <= 20:
		return "S"
	case myScore >= 12 && myScore <= 16:
		if dealearScore >= 7{
			return "H"
		}
		return "S"
	default:
		return "H"
	}
}