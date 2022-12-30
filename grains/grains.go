package grains

import (
    "errors"
    )

func Square(number int) (uint64, error) {
	if number <1  || number > 64 {
        return 0, errors.New("please provide a number between 1-64. ")
    }
	return uint64(1) << (number-1), nil
}

func Total() uint64 {
    // nasty hack since chess board size is fixed
	return ^uint64(0) 
}
