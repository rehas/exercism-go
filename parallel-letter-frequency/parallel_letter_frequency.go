package letter

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

// ConcurrentFrequency counts the frequency of each rune in the given strings,
// by making use of concurrency.
func ConcurrentFrequency(l []string) FreqMap {
	// for every string in strings
    // fire a go routine that creates a local freqMap and fills it out. (1)
    // this local freqMap should be passed on to an out channel when complete. 
	// listen to the out channel on main thread. (2)
    // combine the incoming freqMaps into the res FreqMap
    out := make(chan FreqMap)

	for _, s := range l{
        go count(s, out) // (1)
    }

    res := make(FreqMap)
    
    for range l{
        fm := <- out // (2)
        for k, v := range fm{
            if _, ok := res[k]; ok{
                res[k]+=v
            }else{
            	res[k] = v
            }
        }
    }
	return res
}

func count(s string, out chan FreqMap){
    freqMap := make(FreqMap)
    for _, c := range s{
        if _, ok := freqMap[c]; ok{
            freqMap[c]++
        }else{
        	freqMap[c] = 1
        }
    }
    out <- freqMap
}
