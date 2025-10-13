// Print時のスタイル制御モジュール.
package internal

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
)

var (
	HEADER = color.New(color.FgBlack).Add(color.Bold) // ヘッダスタイル
	BEFORE = color.New(color.FgRed)                   // Beforeスタイル
	AFTER  = color.New(color.FgGreen)                 // Afterスタイル
)

/*
C/A/DでChangeとして検知された情報のPrint定義.
*/
func PrintChange(bls int, ble int, bct []string, als int, ale int, act []string) {
	HEADER.Printf("%d-%dc%d-%d\n", bls, ble, als, ale)
	BEFORE.Printf("%s\n", strings.Join(bct, "\n"))
	fmt.Printf("------------\n")
	AFTER.Printf("%s\n", strings.Join(act, "\n"))
	fmt.Printf("\n")
}

/*
C/A/DでAddとして検知された情報のPrint定義.
*/
func PrintAdd(auis int, auie int, act []string) {
	HEADER.Printf("0a%d-%d\n", auis, auie)
	AFTER.Printf("%s\n", strings.Join(act, "\n"))
	fmt.Printf("\n")
}

/*
C/A/DでDeleteとして検知された情報のPrint定義.
*/
func PrintDelete(buis int, buie int, bct []string) {
	HEADER.Printf("%d-%dd0\n", buis, buie)
	BEFORE.Printf("%s\n", strings.Join(bct, "\n"))
	fmt.Printf("\n")
}

/*
-q, --briefオプションで差分が検知出来た場合のPrint定義.
*/
func PrintBrief() {
	fmt.Printf("Files differ\n")
}

/*
-s, --report-identical-filesオプションでファイルが同一の場合のPrint定義.
*/
func PrintIdentical() {
	fmt.Printf("Files are identical\n")
}
