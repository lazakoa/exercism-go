package parsinglogfiles

import "regexp"

func IsValidLine(text string) bool {
	re := regexp.MustCompile(`\A(\[TRC\]|\[DBG\]|\[INF\]|\[WRN\]|\[ERR\]|\[FTL\]).+`)
	return re.MatchString(text)

}

func SplitLogLine(text string) []string {
	re := regexp.MustCompile(`(<[=\*\~\-]*>)`)
	return re.Split(text, -1)

}

func CountQuotedPasswords(lines []string) int {
	re := regexp.MustCompile(`((?i)".*password.*")`)
	total := 0
	for _, line := range lines {
		total += len(re.FindAllString(line, -1))
	}
	return total
}

func RemoveEndOfLineText(text string) string {
	re := regexp.MustCompile(`((?i)end-of-line\d*)`)
	return re.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	var newLines []string
	re := regexp.MustCompile(`User\s+[a-zA-Z\d]*`)
	usere := regexp.MustCompile(`(User\s+)?`)
	for _, line := range lines {
		if re.MatchString(line) {
			user := re.FindString(line)
			newLine := "[USR] " + usere.ReplaceAllString(user, "") + " " + line
			newLines = append(newLines, newLine)
		} else {
			newLines = append(newLines, line)
		}
	}
	return newLines
}
