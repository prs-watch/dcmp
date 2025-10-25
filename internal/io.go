// File I/O operations module.
package internal

import (
	"bufio"
	"os"
	"regexp"
	"strings"
)

/*
GetLines loads file contents and returns them as a slice of strings with optional filters applied.
Parameters:
  - path: File path to read
  - ignoreBlankFlag: If true, skip blank lines
  - ignoreCaseFlag: If true, convert all text to uppercase
  - ignoreSpaceFlag: If true, normalize whitespace (trim and collapse to single space)
  - ignoreAllSpaceFlag: If true, remove all whitespace
  - ignoreCrFlag: If true, strip trailing carriage returns
  - ignoreMatchingLines: Regex patterns for lines to skip
  - expandTabsFlag: If true, replace tabs with 8 spaces

Returns a slice of processed lines and any error encountered.
*/
func GetLines(
	path string,
	ignoreBlankFlag bool,
	ignoreCaseFlag bool,
	ignoreSpaceFlag bool,
	ignoreAllSpaceFlag bool,
	ignoreCrFlag bool,
	ignoreMatchingLines []string,
	expandTabsFlag bool,
) ([]string, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	// Prepare for scanning
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

		// Ignore specific regex
		for _, pattern := range patterns {
			if pattern.MatchString(line) {
				ignoreLine = true
				break
			}
		}
		if ignoreLine {
			continue
		}

		// Ignore blank
		if ignoreBlankFlag && line == "" {
			continue
		}

		// Ignore case diff
		if ignoreCaseFlag {
			line = strings.ToUpper(line)
		}

		// Ignore space
		re := regexp.MustCompile(`\s+`)
		if ignoreSpaceFlag {
			line = strings.TrimSpace(line)
			line = re.ReplaceAllString(line, " ")
		}

		// Ignore all spaces
		if ignoreAllSpaceFlag {
			line = re.ReplaceAllString(line, "")
		}

		// Ignore CR
		if ignoreCrFlag {
			line = strings.TrimRight(line, "\r")
		}

		// Replace tabs to spaces
		if expandTabsFlag {
			line = strings.ReplaceAll(line, "\t", "        ")
		}

		lines = append(lines, line)
	}

	return lines, nil
}
