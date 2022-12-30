package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
values := map[string]int{
"ace":	11,	 
"eight":	8,
"two":	2,	 
"nine":	9,
"three":	3,
"ten":	10,
"four":	4,
"jack":	10,
"five":	5,
"queen":	10,
"six":	6,	 
"king":	10,
"seven":	7,
"other":	0,}

    v, ok := values[card]
    if ok{
        return v
    }else{
    	return 0
    }
    
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
    cardTotal := ParseCard(card1) + ParseCard(card2)
    dealerValue := ParseCard(dealerCard)
	switch{
        case card1 == card2 && card1 == "ace":
    		return "P"
        case cardTotal == 21 && dealerValue < 10:
    		return "W"
        case cardTotal == 21 && dealerValue >= 10:
    		return "S"
        case cardTotal >=17 && cardTotal <=20 :
    		return "S"
        case cardTotal >=12 && cardTotal <=16 && dealerValue >=7:
    		return "H"
        case cardTotal >=12 && cardTotal <=16 && dealerValue <7:
    		return "S"
        case cardTotal <=11:
    		return "H"
    	default:
    		return "X"
    }
}
