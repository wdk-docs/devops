---
title: '语言功能'
linkTitle: ''
weight: 2
description: >
  智能感知,代码导航,代码编辑,诊断,测试,调试
type: 'docs'
---

## 智能感知

- 当您输入符号的自动完成 (using language server or `gocode`)
- Signature Help for functions as you type (using language server or `gogetdoc` or `godef`+`go doc`)
- Quick Info on the symbol as you hover over it (using language server or `gogetdoc` or `godef`+`go doc`)

## 代码导航

- Go to or Peek Definition of symbols (using language server or `gogetdoc` or `godef`+`go doc`)
- Find References of symbols and Implementations of interfaces (using language server or `guru`)
- Go to symbol in file or see the file outline (using `go-outline`)
- Go to symbol in workspace (using language server or `go-symbols`)
- Toggle between a Go program and the corresponding test file.

## 代码编辑

- [Code Snippets](https://github.com/microsoft/vscode-go/blob/master/snippets/go.json) for quick coding
- Format code on file save as well as format manually (using `goreturns` or `goimports` which also remove unused imports
  or `gofmt`). To disable the format on save feature, add `"[go]": {"editor.formatOnSave": false}` to your settings.
- Symbol Rename (using `gorename`. Note: For Undo after rename to work in Windows you need to have `diff` tool in your
  path)
- Add Imports to current file (using `gopkgs`)
- Add/Remove Tags on struct fields (using `gomodifytags`)
- Generate method stubs for interfaces (using `impl`)
- Fill struct literals with default values (using `fillstruct`)

## 诊断

- Build-on-save to compile code and show build errors. (using `go build` and `go test`)
- Vet-on-save to run `go vet` and show errors as warnings
- Lint-on-save to show linting errors as warnings (using `golint`, `staticcheck`, `golangci-lint` or `revive`)
- Semantic/Syntactic error reporting as you type (using `gotype-live`)

## 测试

- Run Tests under the cursor, in current file, in current package, in the whole workspace using either commands or
  codelens
- Run Benchmarks under the cursor using either commands or codelens
- Show code coverage either on demand or after running tests in the package.
- Generate unit tests skeleton (using `gotests`)

## 调试

- 调试代码，二进制文件或测试 (using `delve`)

## 其他

- Install/Update all dependent Go tools
- Upload to the Go Playground (using `goplay`)
