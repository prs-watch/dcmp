package internal

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
)

var HEADER = color.New(color.FgBlack).Add(color.Bold)
var DELETED = color.New(color.FgRed)
var ADDED = color.New(color.FgGreen)

func getLcsTable(bf []string, af []string) [][]int {
	bfLen, afLen := len(bf), len(af)
	lt := make([][]int, bfLen+1)
	for i := range lt {
		lt[i] = make([]int, afLen+1)
	}
	// check # of LCS parts
	for i := 1; i <= bfLen; i++ {
		for j := 1; j <= afLen; j++ {
			if bf[i-1] == af[j-1] {
				// matched
				lt[i][j] = lt[i-1][j-1] + 1
			} else if lt[i-1][j] >= lt[i][j-1] {
				// unmatched
				lt[i][j] = lt[i-1][j]
			} else {
				// unmatched
				lt[i][j] = lt[i][j-1]
			}
		}
	}

	return lt
}

func getLcsPairs(bf []string, af []string) [][2]int {
	lt := getLcsTable(bf, af)
	i, j := len(bf), len(af)
	var pairs [][2]int

	// get pairs
	for i > 0 && j > 0 {
		if bf[i-1] == af[j-1] {
			pairs = append(pairs, [2]int{i, j})
			i--
			j--
		} else if lt[i-1][j] >= lt[i][j-1] {
			i--
		} else {
			j--
		}
	}

	// reverse pairs
	for ti, di := 0, len(pairs)-1; ti < di; ti, di = ti+1, di-1 {
		pairs[ti], pairs[di] = pairs[di], pairs[ti]
	}

	return pairs
}

func printChange(bui int, bmi int, bf []string, aui int, ami int, af []string) {
	var bct []string
	for i := bui - 1; i < bmi-1; i++ {
		bct = append(bct, "<"+bf[i])
	}
	var act []string
	for i := aui - 1; i < ami-1; i++ {
		act = append(act, ">"+af[i])
	}
	HEADER.Printf("%d-%dc%d-%d\n", bui, bmi-1, aui, ami-1)
	DELETED.Printf("%s\n", strings.Join(bct, "\n"))
	fmt.Printf("------------\n")
	ADDED.Printf("%s\n", strings.Join(act, "\n"))
	fmt.Printf("\n")
}

func printDelete(bui int, bmi int, bf []string) {
	var bct []string
	for i := bui - 1; i < bmi-1; i++ {
		bct = append(bct, "<"+bf[i])
	}
	HEADER.Printf("%d-%dd0\n", bui, bmi-1)
	DELETED.Printf("%s\n", strings.Join(bct, "\n"))
	fmt.Printf("\n")
}

func printAdd(aui int, ami int, af []string) {
	var act []string
	for i := aui - 1; i < ami-1; i++ {
		act = append(act, ">"+af[i])
	}
	HEADER.Printf("0a%d-%d\n", aui, ami-1)
	ADDED.Printf("%s\n", strings.Join(act, "\n"))
	fmt.Printf("\n")
}

func getDiffInfo(bf []string, af []string) {
	pairs := getLcsPairs(bf, af)

	// unprocessed row
	bui, aui := 1, 1

	// pairs-based check
	for _, p := range pairs {
		bmi, ami := p[0], p[1]
		isbu := bui <= bmi-1
		isau := aui <= ami-1
		switch {
		case isbu && isau:
			printChange(bui, bmi, bf, aui, ami, af)
		case isbu:
			printDelete(bui, bmi, bf)
		case isau:
			printAdd(aui, ami, af)
		default:
			// nope
		}
		bui = bmi + 1
		aui = ami + 1
	}

	// check last line
	bEnd, aEnd := len(bf), len(af)
	isbu := bui <= bEnd
	isau := aui <= aEnd
	switch {
	case isbu && isau:
		printChange(bui, bEnd+1, bf, aui, aEnd+1, af)
	case isbu:
		printDelete(bui, bEnd+1, bf)
	case isau:
		printAdd(aui, aEnd+1, af)
	default:
		// nope
	}
}

func Execute(bfpath string, afpath string) error {
	bff, err := os.Open(bfpath)
	if err != nil {
		return err
	}
	defer bff.Close()
	aff, err := os.Open(afpath)
	if err != nil {
		return err
	}
	defer aff.Close()

	var bflines []string
	bfs := bufio.NewScanner(bff)
	for bfs.Scan() {
		bflines = append(bflines, bfs.Text())
	}
	var aflines []string
	afs := bufio.NewScanner(aff)
	for afs.Scan() {
		aflines = append(aflines, afs.Text())
	}

	getDiffInfo(bflines, aflines)

	return nil
}
