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

## Testing

### Test Structure

`dcmp` uses black-box testing to verify command-line option behavior. Tests are located in `internal/dcmp_test.go` and use data-driven test cases under `internal/testdata/`.

#### Directory Structure

```
internal/
├── dcmp_test.go           # Main test file
└── testdata/              # Test data for each option
    ├── no-options/        # Basic diff without options
    │   ├── a.txt          # Before file
    │   ├── b.txt          # After file
    │   └── expected.txt   # Expected output
    ├── brief/             # -q, --brief
    ├── report-identical-files/  # -s, --report-identical-files
    ├── ignore-blank-lines/      # -B, --ignore-blank-lines
    ├── ignore-case/             # -i, --ignore-case
    ├── ignore-space-change/     # -b, --ignore-space-change
    ├── ignore-all-space/        # -w, --ignore-all-space
    ├── strip-trailing-cr/       # --strip-trailing-cr
    ├── ignore-matching-lines/   # -I, --ignore-matching-lines
    └── expand-tabs/             # -t, --expand-tabs
```

Each test case directory contains:
- `a.txt`: Before file (input)
- `b.txt`: After file (input)
- `expected.txt`: Expected output when comparing with the specific option

### Running Tests

```bash
# Run all tests
go test ./internal

# Run tests with verbose output
go test ./internal -v

# Run a specific test case
go test ./internal -v -run TestExecute/brief
```

### Adding New Tests

When adding a new command-line option or modifying existing behavior:

1. **Create a test directory** under `internal/testdata/[option-name]/`

2. **Create test files**:
   - `a.txt`: Before file with content that demonstrates the option
   - `b.txt`: After file with appropriate differences
   - `expected.txt`: Expected output (use `--color=never` when generating)

3. **Generate expected output**:
   ```bash
   go run main.go internal/testdata/[option-name]/a.txt \
                  internal/testdata/[option-name]/b.txt \
                  [your-option] --color=never > internal/testdata/[option-name]/expected.txt
   ```

4. **Verify line endings**: Ensure `expected.txt` uses LF (`\n`), not CRLF (`\r\n`)
   ```bash
   dos2unix internal/testdata/[option-name]/expected.txt
   ```

5. **Add the test case** to the `cases` array in `internal/dcmp_test.go`:
   ```go
   var cases = []string{
       // ... existing cases ...
       "your-new-option",
   }
   ```

6. **Add option flag logic** in `TestExecute`:
   ```go
   if tc == "your-new-option" {
       yourNewOptionFlag = true
   }
   ```

7. **Run the test** to verify:
   ```bash
   go test ./internal -v -run TestExecute/your-new-option
   ```

### Test Guidelines

- **Black-box testing**: Tests verify external behavior only, not internal implementation
- **Data-driven**: Each option has dedicated test data to ensure isolation
- **Reproducible**: Expected outputs are version-controlled for consistency
- **No color codes**: Always use `--color=never` when generating expected outputs to avoid ANSI escape sequences in test files

## License

See [LICENSE](LICENSE) file for details.

## Contributing

Contributions are welcome! Please feel free to submit issues or pull requests.