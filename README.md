# dcmp

A CLI tool for extracting and displaying file differences based on the LCS (Longest Common Subsequence) algorithm.

## Overview

`dcmp` is a file comparison utility that uses the LCS algorithm to detect and display differences between two files. It provides various options for customizing the comparison behavior, including ignoring whitespace, case differences, blank lines, and more.

## Features

- **LCS-based difference detection**: Accurately identifies changes, additions, and deletions
- **Colored output**: Highlights differences with customizable color modes
- **Flexible comparison options**: Multiple flags for ignoring specific content
- **Regex pattern matching**: Exclude lines matching specific patterns
- **Cross-platform**: Written in Go for portability

## Installation

```bash
go install github.com/prs-watch/dcmp@latest
```

## Usage

### Basic Usage

```bash
dcmp [file1] [file2] [flags]
```

### Example

Given two files `hoge.md` and `fuga.md`:

**hoge.md:**
```md
# hoge

- 1
- 2
- 3
- 4
- 5
- 6

堂本光一
```

**fuga.md:**
```md
# fuga

- 1
- 2
- 3
- 5
- 6
- 7

堂本剛
```

Running `dcmp`:

```bash
dcmp hoge.md fuga.md
```

**Output:**
```
1-1c1-1
<# hoge
------------
># fuga

6-6d0
<- 4

0a8-8
>- 7

10-10c10-10
<堂本光一
------------
>堂本剛
```

## Command-Line Flags

### Output Control

| Flag | Short | Description |
|------|-------|-------------|
| `--brief` | `-q` | Output to stdout only when file differences exist |
| `--report-identical-files` | `-s` | Output to stdout only when files are identical |
| `--color` | | Control colored output: `auto`, `always`, or `never` (default: `auto`) |

### Comparison Options

| Flag | Short | Description |
|------|-------|-------------|
| `--ignore-blank-lines` | `-B` | Ignore blank lines when comparing files |
| `--ignore-case` | `-i` | Ignore case differences when comparing files |
| `--ignore-space-change` | `-b` | Ignore whitespace changes when comparing files |
| `--ignore-all-space` | `-w` | Ignore all whitespace when comparing files |
| `--strip-trailing-cr` | | Ignore trailing CR (carriage return) when comparing files |
| `--ignore-matching-lines` | `-I` | Ignore lines matching the specified regex pattern (can be used multiple times) |
| `--expand-tabs` | `-t` | Replace tabs with spaces (8 spaces) before comparison |

### Examples

**Ignore case differences:**
```bash
dcmp -i file1.txt file2.txt
```

**Ignore all whitespace:**
```bash
dcmp -w file1.txt file2.txt
```

**Only report if files differ:**
```bash
dcmp -q file1.txt file2.txt
```

**Ignore lines matching a pattern:**
```bash
dcmp -I "^#.*" file1.txt file2.txt
```

**Combine multiple flags:**
```bash
dcmp -i -B -w file1.txt file2.txt
```

**Force colored output:**
```bash
dcmp --color always file1.txt file2.txt
```

## Output Format

The output follows this format for detected changes:

- **Change (C)**: `[before-line-range]c[after-line-range]`
  - Lines prefixed with `<` indicate removed content
  - Lines prefixed with `>` indicate added content
  - Separator: `------------`

- **Delete (D)**: `[before-line-range]d0`
  - Lines prefixed with `<` indicate deleted content

- **Add (A)**: `0a[after-line-range]`
  - Lines prefixed with `>` indicate added content

## Architecture

`dcmp` is built with the following modules:

- **LCS Module** (`internal/lcs.go`): Implements the Longest Common Subsequence algorithm
- **Core Module** (`internal/dcmp.go`): Handles C/A/D detection logic
- **Print Module** (`internal/print.go`): Manages styled output with color support
- **I/O Module** (`internal/io.go`): Handles file reading with various filters
- **Command Module** (`cmd/root.go`): Cobra-based CLI interface

## License

See [LICENSE](LICENSE) file for details.

## Contributing

Contributions are welcome! Please feel free to submit issues or pull requests.