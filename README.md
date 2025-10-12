# dcmp

LCSベースでDiff抽出の上、PrintするCLIツール.

## Install

```bash
$ go install github.com/prs-watch/dcmp@latest
```

## Example

`hoge.md` と `fuga.md` がある場合

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

`dcmp` コマンドを実行することでDiffを取得出来ます.

```bash
$ dcmp hoge.md fuga.md

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