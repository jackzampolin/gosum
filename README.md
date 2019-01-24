# `gosum`

A very simple utility to calculate the sha256sum of `|`ed input.

### Install

```bash
$ go get github.com/jackzampolin/gosum
```

or

```bash
$ git clone git@github.com:jackzampolin/gosum
$ go install main.go
```

### Usage

```bash
$ tar -c dir/ | gosum
262f2dbc2ca359c0e985ccc2e8e19078597eecaaa6857b71fe0b68c06dd59c05  -
```
