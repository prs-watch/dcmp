package internal

import (
	"bufio"
	"os"
	"regexp"
	"strings"
)

/*
ファイル内容をロードし[]stringとして返却.
*/
func GetLines(path string, ignoreBlankFlag bool, ignoreCaseFlag bool, ignoreSpaceFlag bool, ignoreAllSpaceFlag bool, ignoreCrFlag bool, ignoreMatchingLines []string) ([]string, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	// prepare for scanning
	var lines []string
	fs := bufio.NewScanner(f)
	fs.Buffer(make([]byte, 1024), 1024*1024)
	patterns := []*regexp.Regexp{}
	for _, pattern := range ignoreMatchingLines {
		re, err := regexp.Compile(pattern)
		if err != nil {
			return nil, err
		}
		patterns = append(patterns, re)
	}

	for fs.Scan() {
		line := fs.Text()
		ignoreLine := false

		// ignore specific regex
		for _, pattern := range patterns {
			if pattern.MatchString(line) {
				ignoreLine = true
				break
			}
		}
		if ignoreLine {
			continue
		}

		// ignore blank
		if ignoreBlankFlag && line == "" {
			continue
		}

		// ignore case diff
		if ignoreCaseFlag {
			line = strings.ToUpper(line)
		}

		// ignore space
		re := regexp.MustCompile(`\s+`)
		if ignoreSpaceFlag {
			line = strings.TrimSpace(line)
			line = re.ReplaceAllString(line, " ")
		}

		// ignore all spaces
		if ignoreAllSpaceFlag {
			line = re.ReplaceAllString(line, "")
		}

		// ignore CR
		if ignoreCrFlag {
			line = strings.TrimRight(line, "\r")
		}

		lines = append(lines, line)
	}

	return lines, nil
}
