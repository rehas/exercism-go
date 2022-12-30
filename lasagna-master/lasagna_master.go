package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, avgTimePerLayer int) int{
    var time int
    if avgTimePerLayer == 0{
        time = 2
    }else{
    	time = avgTimePerLayer
    }
	return time * len(layers)
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (noodles int, sauce float64){
    for _, l := range layers{
        if l == "sauce"{
            sauce += 0.2
        }
    	if l == "noodles"{
            noodles +=50
        }
    }
	return
}

// TODO: define the 'AddSecretIngredient()' function

func AddSecretIngredient(friends, mine []string){
    mine[len(mine)-1] = friends[len(friends)-1]
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(amounts []float64, portions int) []float64{
    res := make([]float64, len(amounts))

    for i, a := range amounts{
        res[i] = a * (float64(portions) /2.0)
    }
	return res
}
