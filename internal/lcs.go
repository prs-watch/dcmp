// LCSによるC/A/D検知モジュール.
package internal

/*
テキストを比較しLCSテーブルを取得する.
*/
func getLcsTable(bf []string, af []string) [][]int {
	bfLen, afLen := len(bf), len(af)

	// LCSテーブル
	lt := make([][]int, bfLen+1)
	for i := range lt {
		lt[i] = make([]int, afLen+1)
	}

	// テキスト走査、一致行をカウントしてLCSテーブルに格納.
	for i := 1; i <= bfLen; i++ {
		for j := 1; j <= afLen; j++ {
			if bf[i-1] == af[j-1] {
				// 行一致の場合
				lt[i][j] = lt[i-1][j-1] + 1
			} else if lt[i-1][j] >= lt[i][j-1] {
				// 行不一致の場合
				lt[i][j] = lt[i-1][j]
			} else {
				// 行不一致の場合
				lt[i][j] = lt[i][j-1]
			}
		}
	}

	return lt
}

/*
LCSテーブルから一致行のペアを取得する. 後続のC/A/D判定ロジックにて利用.
*/
func GetLcsPairs(bf []string, af []string) [][2]int {
	lt := getLcsTable(bf, af)
	i, j := len(bf), len(af)

	// 一致行ペア
	var pairs [][2]int

	// 一致行ペア確認. 一致行数が上/左で一緒の場合は上へ移動.
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

	// テキスト末からペアを格納しているため、先頭から並ぶ様ソート.
	for ti, di := 0, len(pairs)-1; ti < di; ti, di = ti+1, di-1 {
		pairs[ti], pairs[di] = pairs[di], pairs[ti]
	}

	return pairs
}
