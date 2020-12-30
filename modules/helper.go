package modules

import "strings"

const summaryPrefix = "summary: "

//ExtractSummary Extract a line "^summary: (.*)" from the jira output lines.
func ExtractSummary(lines []string) (result string) {
	result = ""
	for _, line := range lines {
		if strings.HasPrefix(line, summaryPrefix) {
			result = line[len(summaryPrefix):]
		}
	}
	return
}
