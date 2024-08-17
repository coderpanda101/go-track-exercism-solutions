package parsinglogfiles

import (
	"fmt"
	"regexp"
)

func IsValidLine(text string) bool {
	validRegex := regexp.MustCompile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)\]`)
	return validRegex.MatchString(text)
}

func SplitLogLine(text string) []string {
	regex := regexp.MustCompile(`<[-~=*]*>`)
	return regex.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	validRegex := regexp.MustCompile(`"(?i).*password.*"`)
	count := 0
	for _, line := range lines {
		if validRegex.MatchString(line) {
			count++
		}
	}
	return count
}

func RemoveEndOfLineText(text string) string {
	regex := regexp.MustCompile(`end-of-line\d*`)
	return regex.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	regex := regexp.MustCompile(`(?:User\s+([A-z0-9]+))`)
	for i, line := range lines {
		if regex.MatchString(line) {
			submatch := regex.FindStringSubmatch(line)
			lines[i] = fmt.Sprintf("[USR] %s %s", submatch[1], line)
		}
	}
	return lines
}
