package luhn

import "strings"

func Valid(id string) bool {
    // besides spaces, there should be at least 1 digit
	if len(strings.TrimSpace(id)) <=1 {
        return false
    }

    var oddSum int
	var evenSum int

    var even bool

    // begin from the end
    for i := len(id)-1; i>=0; i-- {
    	c := id[i] // this is problematic for variable sized characters
        // spaces are valid but they don't count
        if c == ' ' {
            continue
        }
		// values that are not numbers are invalid 
        // can be checked via regex before the loop
    	if !(c >= '0' && c <= '9'){
            return false
        }    	
    	
    	value := int(c - '0') // convert to int from string

        if even{
            value *=2
            // this is different than value [mod] 9
            if value > 9{
                value -=9
            }
            evenSum += value
            even = false
            continue 
        }
        oddSum += value
        even = true
    }

    return (oddSum + evenSum) % 10 == 0
}
