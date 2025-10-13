package internal

import (
	"bufio"
	"os"
)

/*
ファイル内容をロードし[]stringとして返却.
*/
func GetLines(path string, ignoreBlankFlag bool) ([]string, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var lines []string
	fs := bufio.NewScanner(f)
	fs.Buffer(make([]byte, 1024), 1024*1024)
	for fs.Scan() {
		if ignoreBlankFlag && fs.Text() == "" {
			continue
		}
		lines = append(lines, fs.Text())
	}

	return lines, nil
}
