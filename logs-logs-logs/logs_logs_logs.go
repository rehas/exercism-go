package logs



// Application identifies the application emitting the given log.
func Application(log string) string {
	for _, c := range log{
        switch{
            case c == '\u2757':
        		return "recommendation"
            case c== 128269:
        		return "search"
            case c == '\u2600':
        		return "weather"
        }
    }
	return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	logt := []rune(log)
    for i, c := range logt{
        if c == oldRune{
            logt[i] = newRune
        }
    }
	return string(logt)
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	var len int
    for i, _ := range log{
        if i > len{
            len = i
        }
    }
	return len+1 <= limit
}
