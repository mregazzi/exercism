package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch card {
        case "ace":
        	return 11
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
        case "ten":
        	return 10
        case "jack":
        	return 10
        case "queen":
        	return 10
        case "king":
        	return 10
    }
    return 0
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	var myValue int = ParseCard(card1) + ParseCard(card2)
    var dealerValue int = ParseCard(dealerCard)
    if(myValue == 22) {
        return "P"
    }
    if(myValue == 21 && dealerValue < 10) {
        return "W"
    }
    if(myValue >= 17) {
        return "S"
    }
    if(myValue >=12 && dealerValue < 7){
    	return "S"
    }
    return "H"
}
