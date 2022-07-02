package parsinglogfiles

import "regexp"

func IsValidLine(text string) bool {
	re := regexp.MustCompile(`^\[(?:TRC|DBG|INF|WRN|ERR|FTL)\]`)
	return re.MatchString(text)
}

func SplitLogLine(text string) []string {
	re := regexp.MustCompile(`<\W*>`)
	return re.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	re := regexp.MustCompile(`(?i)".*\bpassword\b.*"`)
	count := 0
	for _, line := range lines {
		sl := re.FindStringSubmatch(line)
		count += len(sl)
	}

	return count
}

func RemoveEndOfLineText(text string) string {
	re := regexp.MustCompile(`end-of-line\d*`)
	return re.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	re := regexp.MustCompile(`User\s+(\w+)`)
	for i, line := range lines {
		sl := re.FindStringSubmatch(line)
		if sl != nil {
			lines[i] = "[USR] " + sl[1] + " " + line
		}
	}

	return lines
}
