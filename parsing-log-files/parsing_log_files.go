package parsinglogfiles

import (
    "regexp"
    "fmt"
    )

func IsValidLine(text string) bool {
	re := regexp.MustCompile(`^(\[TRC\]|\[DBG\]|\[INF\]|\[WRN\]|\[ERR\]|\[FTL\])`)
    return re.MatchString(text)
}

func SplitLogLine(text string) []string {
	re := regexp.MustCompile(`<[~*-=]*>`)
	return re.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
    re := regexp.MustCompile(`(?i)".*password.*"`)
    var count int
    for _, l := range lines{
        if re.MatchString(l){
            count++
        }
    }

    return count
}

func RemoveEndOfLineText(text string) string {
	re := regexp.MustCompile(`end-of-line[0-9]*`)
    return re.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	re := regexp.MustCompile(`User[\s]*[a-zA-Z]+[0-9]+`)
	for i, line := range lines{
    	match := re.FindString(line)
        if match == ""{
            continue
        }
    	userNameRe := regexp.MustCompile(`[a-zA-Z]+[0-9]+`)
        userName := userNameRe.FindString(match)
        lines[i] = fmt.Sprintf("[USR] %s %s", userName, line)
    }

    return lines
}
