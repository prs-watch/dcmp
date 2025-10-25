// `dcmp` のコアモジュール.
package internal

import "errors"

/*
LCSベースに検出したC/A/Dを判定の上Printする.
*/
func printResult(bf []string, af []string, briefFlag bool, identicalFlag bool, colorMode string) error {
	// Print時のcolorModeの適用
	err := ApplyColorMode(colorMode)
	if err != nil {
		return err
	}

	// LCSペアの取得
	pairs := GetLcsPairs(bf, af)

	// 完全一致か判定
	if len(pairs) == len(bf) && len(pairs) == len(af) {
		if identicalFlag {
			handleIdentical()
		}
		return nil
	} else {
		// -q, --briefの場合出力の上エラー返却
		if briefFlag {
			handleBrief()
			return errors.New("")
		}
	}

	// 未処理行. 初期は1行目からスタート.
	bui, aui := 1, 1

	// 一致行ペアを走査してC/A/Dをチェック.
	for _, p := range pairs {
		// 一致行ペア
		bmi, ami := p[0], p[1]

		// 未処理行と一致行のポジション確認. 一致行より手前にあれば未処理として判定.
		isbu := bui <= bmi-1
		isau := aui <= ami-1

		// C/A/Dを判定.
		switch {
		case isbu && isau:
			handleChange(bui, bmi, bf, aui, ami, af)
		case isbu:
			handleDelete(bui, bmi, bf)
		case isau:
			handleAdd(aui, ami, af)
		default:
			// 一致行は処理無し.
		}

		// 走査済行分だけincrementする.
		bui = bmi + 1
		aui = ami + 1
	}

	// 末行だけ処理が漏れるため個別にハンドリング.
	bEnd, aEnd := len(bf), len(af)
	isbu := bui <= bEnd
	isau := aui <= aEnd

	switch {
	case isbu && isau:
		handleChange(bui, bEnd+1, bf, aui, aEnd+1, af)
	case isbu:
		handleDelete(bui, bEnd+1, bf)
	case isau:
		handleAdd(aui, aEnd+1, af)
	default:
		// 一致行は処理無し.
	}

	return nil
}

/*
Change箇所の行情報取得とPrint.
*/
func handleChange(bui int, bmi int, bf []string, aui int, ami int, af []string) {
	var bct []string
	for i := bui - 1; i < bmi-1; i++ {
		bct = append(bct, "<"+bf[i])
	}
	var act []string
	for i := aui - 1; i < ami-1; i++ {
		act = append(act, ">"+af[i])
	}
	PrintChange(bui, bmi-1, bct, aui, ami-1, act)
}

/*
Delete箇所の行情報取得とPrint.
*/
func handleDelete(bui int, bmi int, bf []string) {
	var bct []string
	for i := bui - 1; i < bmi-1; i++ {
		bct = append(bct, "<"+bf[i])
	}
	PrintDelete(bui, bmi-1, bct)
}

/*
Add箇所の行情報取得とPrint.
*/
func handleAdd(aui int, ami int, af []string) {
	var act []string
	for i := aui - 1; i < ami-1; i++ {
		act = append(act, ">"+af[i])
	}
	PrintAdd(aui, ami-1, act)
}

/*
-q, --brief時に差分が検知された場合のPrint.
*/
func handleBrief() {
	PrintBrief()
}

/*
-s, --report-identical-files時に同一ファイルであった場合のPrint.
*/
func handleIdentical() {
	PrintIdentical()
}

/*
モジュールエントリポイント. 2ファイルパスをinputに差分情報をPrintする.
*/
func Execute(bfpath string, afpath string, briefFlag bool, identicalFlag bool, ignoreBlankFlag bool, ignoreCaseFlag bool, ignoreSpaceFlag bool, ignoreAllSpaceFlag bool, colorMode string, ignoreCrFlag bool, ignoreMatchingLines []string) error {
	bflines, err := GetLines(bfpath, ignoreBlankFlag, ignoreCaseFlag, ignoreSpaceFlag, ignoreAllSpaceFlag, ignoreCrFlag, ignoreMatchingLines)
	if err != nil {
		return err
	}
	aflines, err := GetLines(afpath, ignoreBlankFlag, ignoreCaseFlag, ignoreSpaceFlag, ignoreAllSpaceFlag, ignoreCrFlag, ignoreMatchingLines)
	if err != nil {
		return err
	}

	err = printResult(bflines, aflines, briefFlag, identicalFlag, colorMode)
	if err != nil {
		return err
	}

	return nil
}
