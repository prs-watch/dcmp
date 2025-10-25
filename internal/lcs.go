// LCS module.
package internal

/*
getLcsTable compares texts and returns the LCS (Longest Common Subsequence) table.
Parameters:
  - bf: Before file lines
  - af: After file lines

Returns a 2D slice representing the LCS table.
*/
func getLcsTable(bf []string, af []string) [][]int {
	bfLen, afLen := len(bf), len(af)

	// LCS table
	lt := make([][]int, bfLen+1)
	for i := range lt {
		lt[i] = make([]int, afLen+1)
	}

	// Scan texts, count matching lines and store in LCS table.
	for i := 1; i <= bfLen; i++ {
		for j := 1; j <= afLen; j++ {
			if bf[i-1] == af[j-1] {
				// Lines match
				lt[i][j] = lt[i-1][j-1] + 1
			} else if lt[i-1][j] >= lt[i][j-1] {
				// Lines don't match
				lt[i][j] = lt[i-1][j]
			} else {
				// Lines don't match
				lt[i][j] = lt[i][j-1]
			}
		}
	}

	return lt
}

/*
GetLcsPairs extracts matching line pairs from the LCS table.
Used in subsequent C/A/D (Change/Add/Delete) detection logic.
Parameters:
  - bf: Before file lines
  - af: After file lines

Returns an array of [2]int pairs where each pair represents matching line indices.
*/
func GetLcsPairs(bf []string, af []string) [][2]int {
	lt := getLcsTable(bf, af)
	i, j := len(bf), len(af)

	// Matching line pairs
	var pairs [][2]int

	// Check matching line pairs. Move up if match count is same for up/left.
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

	// Sort so pairs are arranged from the beginning since they're stored from the end of text.
	for ti, di := 0, len(pairs)-1; ti < di; ti, di = ti+1, di-1 {
		pairs[ti], pairs[di] = pairs[di], pairs[ti]
	}

	return pairs
}
