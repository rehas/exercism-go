package thefarm
import (
    "errors"
    "fmt"
    )
// See types.go for the types defined for this exercise.

// TODO: Define the SillyNephewError type here.
type SillyNephewError struct{
    cows int
    }

func (s *SillyNephewError) Error() string{
    return fmt.Sprintf("silly nephew, there cannot be %d cows", s.cows)
}

// DivideFood computes the fodder amount per cow for the given cows.
func DivideFood(weightFodder WeightFodder, cows int) (float64, error) {
	amount, err := weightFodder.FodderAmount()

    if cows == 0{
        return 0.0, errors.New("division by zero")
    }
    
    if err != ErrScaleMalfunction && err!=nil{
        return 0.0, err
    }

    if amount <0{
        if err == ErrScaleMalfunction || err == nil{
            return 0.0, errors.New("negative fodder")
        }
    }

    if cows < 0 {
        return 0.0, &SillyNephewError{cows:cows}
    }

    if err == ErrScaleMalfunction {
        return amount * 2.0 / float64(cows) , nil
    }
    

    return amount / float64(cows), nil
}
