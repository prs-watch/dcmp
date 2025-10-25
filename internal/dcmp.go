// Core module of `dcmp`.
package internal

import "errors"

/*
printResult determines and prints C/A/D (Change/Add/Delete) based on LCS.
Parameters:
  - bf: Before file lines
  - af: After file lines
  - briefFlag: If true, only report if files differ
  - identicalFlag: If true, only report if files are identical
  - colorMode: Color output mode (auto/always/never)

Returns an error if there's an issue with color mode or if files differ with briefFlag.
*/
func printResult(bf []string, af []string, briefFlag bool, identicalFlag bool, colorMode string) error {
	// Apply colorMode for printing
	err := ApplyColorMode(colorMode)
	if err != nil {
		return err
	}

	// Get LCS pairs
	pairs := GetLcsPairs(bf, af)

	// Check if files are identical
	if len(pairs) == len(bf) && len(pairs) == len(af) {
		if identicalFlag {
			handleIdentical()
		}
		return nil
	} else {
		// For -q, --brief, output and return error
		if briefFlag {
			handleBrief()
			return errors.New("")
		}
	}

	// Unprocessed lines. Start from line 1 initially.
	bui, aui := 1, 1

	// Scan matching line pairs and check C/A/D.
	for _, p := range pairs {
		// Matching line pair
		bmi, ami := p[0], p[1]

		// Check position of unprocessed and matching lines. If before matching line, treat as unprocessed.
		isbu := bui <= bmi-1
		isau := aui <= ami-1

		// Determine C/A/D.
		switch {
		case isbu && isau:
			handleChange(bui, bmi, bf, aui, ami, af)
		case isbu:
			handleDelete(bui, bmi, bf)
		case isau:
			handleAdd(aui, ami, af)
		default:
			// No processing for matching lines.
		}

		// Increment by scanned lines.
		bui = bmi + 1
		aui = ami + 1
	}

	// Handle last lines separately as they may be missed.
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
		// No processing for matching lines.
	}

	return nil
}

/*
handleChange extracts line information for Change locations and prints them.
Parameters:
  - bui: Before unprocessed index (start)
  - bmi: Before match index (end)
  - bf: Before file lines
  - aui: After unprocessed index (start)
  - ami: After match index (end)
  - af: After file lines
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
handleDelete extracts line information for Delete locations and prints them.
Parameters:
  - bui: Before unprocessed index (start)
  - bmi: Before match index (end)
  - bf: Before file lines
*/
func handleDelete(bui int, bmi int, bf []string) {
	var bct []string
	for i := bui - 1; i < bmi-1; i++ {
		bct = append(bct, "<"+bf[i])
	}
	PrintDelete(bui, bmi-1, bct)
}

/*
handleAdd extracts line information for Add locations and prints them.
Parameters:
  - aui: After unprocessed index (start)
  - ami: After match index (end)
  - af: After file lines
*/
func handleAdd(aui int, ami int, af []string) {
	var act []string
	for i := aui - 1; i < ami-1; i++ {
		act = append(act, ">"+af[i])
	}
	PrintAdd(aui, ami-1, act)
}

/*
handleBrief prints message when difference is detected with -q, --brief option.
*/
func handleBrief() {
	PrintBrief()
}

/*
handleIdentical prints message when files are identical with -s, --report-identical-files option.
*/
func handleIdentical() {
	PrintIdentical()
}

/*
Execute is the module entry point that compares two files and prints difference information.
Parameters:
  - bfpath: Path to the before file
  - afpath: Path to the after file
  - briefFlag: If true, only report if files differ (-q, --brief)
  - identicalFlag: If true, only report if files are identical (-s, --report-identical-files)
  - ignoreBlankFlag: If true, ignore blank lines (-B, --ignore-blank-lines)
  - ignoreCaseFlag: If true, ignore case differences (-i, --ignore-case)
  - ignoreSpaceFlag: If true, ignore whitespace changes (-b, --ignore-space-change)
  - ignoreAllSpaceFlag: If true, ignore all whitespace (-w, --ignore-all-space)
  - colorMode: Color output mode: auto/always/never (--color)
  - ignoreCrFlag: If true, ignore trailing CR (--strip-trailing-cr)
  - ignoreMatchingLines: Regular expressions for lines to ignore (-I, --ignore-matching-lines)
  - expandTabsFlag: If true, replace tabs with spaces before comparison (-t, --expand-tabs)

Returns an error if file reading fails or if differences are found with briefFlag.
*/
func Execute(
	bfpath string,
	afpath string,
	briefFlag bool,
	identicalFlag bool,
	ignoreBlankFlag bool,
	ignoreCaseFlag bool,
	ignoreSpaceFlag bool,
	ignoreAllSpaceFlag bool,
	colorMode string,
	ignoreCrFlag bool,
	ignoreMatchingLines []string,
	expandTabsFlag bool,
) error {
	bflines, err := GetLines(
		bfpath,
		ignoreBlankFlag,
		ignoreCaseFlag,
		ignoreSpaceFlag,
		ignoreAllSpaceFlag,
		ignoreCrFlag,
		ignoreMatchingLines,
		expandTabsFlag,
	)
	if err != nil {
		return err
	}
	aflines, err := GetLines(
		afpath,
		ignoreBlankFlag,
		ignoreCaseFlag,
		ignoreSpaceFlag,
		ignoreAllSpaceFlag,
		ignoreCrFlag,
		ignoreMatchingLines,
		expandTabsFlag,
	)
	if err != nil {
		return err
	}

	err = printResult(bflines, aflines, briefFlag, identicalFlag, colorMode)
	if err != nil {
		return err
	}

	return nil
}
