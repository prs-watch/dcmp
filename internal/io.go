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
func GetLines(path string, ignoreBlankFlag bool, ignoreCaseFlag bool, ignoreSpaceFlag bool) ([]string, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var lines []string
	fs := bufio.NewScanner(f)
	fs.Buffer(make([]byte, 1024), 1024*1024)
	for fs.Scan() {
		// blank check
		if ignoreBlankFlag && fs.Text() == "" {
			continue
		}

		// appendする文字列.
		line := fs.Text()
		// case check
		if ignoreCaseFlag {
			line = strings.ToUpper(line)
		}

		// check space
		if ignoreSpaceFlag {
			line = strings.TrimSpace(line)
			re := regexp.MustCompile(`\s+`)
			line = re.ReplaceAllString(line, " ")
		}

		lines = append(lines, line)
	}

	return lines, nil
}
