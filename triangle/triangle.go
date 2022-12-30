// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package triangle should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package triangle


// Notice KindFromSides() returns this type. Pick a suitable data type.
type Kind int

const (
    // Pick values for the following identifiers used by the test program.
    NaT Kind = 0// not a triangle
    Equ Kind = 1// equilateral
    Iso Kind = 2// isosceles
    Sca Kind = 3// scalene
)

// KindFromSides should have a comment documenting it.
func KindFromSides(a, b, c float64) Kind {
	// Write some code here to pass the test suite.
	// Then remove all the stock comments.
	// They're here to help you get started but they only clutter a finished solution.
	// If you leave them in, reviewers may protest!
	if !isTriangle(a,b,c) {
        return NaT
    }

    if a == b && b == c {
        return Equ
    }

    if a == b || b == c || a == c{
        return Iso
    }
    
	return Sca
}

func isTriangle (a,b,c float64) bool{
    if a <=0 || b<=0 || c<= 0{
        return false
    }
    if a + b < c {
        return false
    }

    if a + c < b {
        return false
    }

    if b + c < a{
        return false
    }


    return true
}
