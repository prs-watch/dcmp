package internal

import (
	"bufio"
	"os"
	"strings"
)

/*
ファイル内容をロードし[]stringとして返却.
*/
func GetLines(path string, ignoreBlankFlag bool, ignoreCaseFlag bool) ([]string, error) {
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
		// case check
		if ignoreCaseFlag {
			lines = append(lines, strings.ToUpper(fs.Text()))
		} else {
			lines = append(lines, fs.Text())
		}
	}

	return lines, nil
}
