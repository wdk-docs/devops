---
title: '如何使用这个扩展？'
linkTitle: '如何使用'
weight: 3
description: >

type: 'docs'
---

安装并打开 [Visual Studio Code](https://code.visualstudio.com). 按 `Ctrl+Shift+X` 或 `Cmd+Shift+X` 打开扩展窗格. 查找并
安装`Go`扩展. 您还可以从[市场扩展](https://marketplace.visualstudio.com/items?itemName=ms-vscode.Go)安装. 打开 VS Code
中的任何`.go`文件。扩展现在已启动。

This extension uses a set of Go tools to provide the various rich features. These tools are installed in your GOPATH by
default. If you wish to have these tools in a separate location, provide the desired location in the setting
`go.toolsGopath`. Read more about this and the tools at
[Go tools that the Go extension depends on](https://github.com/Microsoft/vscode-go/wiki/Go-tools-that-the-Go-extension-depends-on).

You will see `Analysis Tools Missing` in the bottom right, clicking this will offer to install all of the dependent Go
tools. You can also run the [command](https://code.visualstudio.com/docs/getstarted/userinterface#_command-palette)
`Go: Install/Update tools` to install/update the same. You need to have git installed for these tool installations to
work.

**Note 1**: Read
[GOPATH in the VS Code Go extension](https://github.com/Microsoft/vscode-go/wiki/GOPATH-in-the-VS-Code-Go-extension) to
learn about the different ways you can get the extension to set GOPATH.

**Note 2**: `Format on save` 在格式化被中止后功能有 750 毫秒的超时. 您可以使用设置`editor.formatOnSaveTimeout`更改此超时
. 当您在 Visual Studio 启用代码的`Auto Save`功能时该功能被停用。

**Note 3**: Unless `go.useLanguageServer` is set to `true`, this extension uses `gocode` to provide completion lists as
you type. If you have disabled the `go.buildOnSave` setting, then you may not get fresh results from not-yet-built
dependencies. Therefore, ensure you have built your dependencies manually in such cases.

## 定制`Go`扩展功能

The Go extension is ready to use on the get go. If you want to customize the features, you can edit the settings in your
User or Workspace settings. Read
[All Settings & Commands in Visual Studio Code Go extension](https://github.com/Microsoft/vscode-go/wiki/All-Settings-&-Commands-in-Visual-Studio-Code-Go-extension)
for the full list of options and their descriptions.

## `Go`语言服务器

The Go extension uses a host of
[Go tools](https://github.com/Microsoft/vscode-go/wiki/Go-tools-that-the-Go-extension-depends-on) to provide the various
language features. An alternative is to use a single language server that provides the same features using the
[Language Server Protocol](https://microsoft.github.io/language-server-protocol/)

Previously, we added support to use `go-langserver`, the
[language server from Sourcegraph](https://github.com/sourcegraph/go-langserver). There is no active development for it
anymore and it doesn't support Go modules. Therefore, we are now switching to use `gopls`, the
[language server from Google](https://github.com/golang/go/wiki/gopls) which is currently in active development.

- If you are already using the language server from Sourcegraph, you can continue to use it as long as you are not using
  Go modules. We do suggest you to move to using `gopls` though.
  - To do so, delete the `go-langserver` binary/executable in your machine and this extension will prompt you to install
    `gopls` after a reload of the VS Code window.
- If you are working on a project that uses Go modules, you will be prompted to use the language server from Google as
  it provides much better support for Go modules.
- If you have never used language server before, and now opt to use it, you will be prompted to install and use the
  language server from Google as long as you are using a Go version > 1.10.

> 注意: 从谷歌的语言服务器支持`Go`版>仅 1.10

### 安装/更新`Go`语言服务器

Ideally, you would see prompts to use/install/update the language server. Follow the prompts and the language server
should get set up correctly. If you want to manually install/update the language server,

- 确保您在您的设置中设置`go.useLanguageServer`到`true`;
- 使用`Go: Install/Update Tools` 命令, 从列表中选择`gopls`，然后按确定。

### 设置来控制使用`Go`语言服务器

下面是你可以用它来控制使用的语言服务器的设置。你需要重新加载 VS Code 窗口中更改这些设置生效。

- Set `go.useLanguageServer` to `true` to enable the use of language server.
- 当使用`gopls`, 参见[推荐设置](https://github.com/golang/tools/blob/master/gopls/doc/vscode.md).
- 如果使用设置`go.languageServerExperimentalFeatures`需要一些从语言服务器的功能可以被禁用。下面是这样，你可以控制的功能
  。默认情况下，所有被设置为`true` 即启用。

```json
  "go.languageServerExperimentalFeatures": {
    "format": true,
    "diagnostics": true,
    "documentLink": true
  }
```

- 设置 `"go.languageServerFlags": ["-logfile", "path to a text file that exists"]` 在日志文件中收集日志.
- 设置 `"go.languageServerFlags": ["-rpc.trace"]` 来查看完整的 RPC 跟踪输出面板 (`View` -> `Output` -> `gopls`)

### 提供 gopls 反馈

如果您发现使用`gopls`语言服务器的任何问题, 请首先查
看[现有的 gopls 问题清单](https://github.com/golang/go/issues?q=is%3Aissue+is%3Aopen+label%3Agopls) and update the
relevant ones with your case before logging a new one at https://github.com/golang/go/issues

### 对于 gopls 有用的链接

- [维基 gopls](https://github.com/golang/tools/blob/master/gopls/doc/user.md)
- [对于 VSCode 推荐设置使用 gopls 时](https://github.com/golang/tools/blob/master/gopls/doc/vscode.md)
- [gopls 疑难解答](https://github.com/golang/go/wiki/gopls#troubleshooting)
- [已知的 gopls 的 bug](https://github.com/golang/go/wiki/gopls#known-issues)
- [Github 上的 gopls 问题](https://github.com/golang/go/issues?q=is%3Aissue+is%3Aopen+label%3Agopls)

## Linter

Linter 是给编码风格的反馈和建议的工具。默认情况下，此扩展使用官方[golint](https://github.com/golang/lint)为 Linter。

您可以更改默认的 linter 并使用更先进[golangci-lint](https://github.com/golangci/golangci-lint) ,通过在您的设置里设置
`go.lintTool` 至 "golangci-lint" .

It shares some of the performance characteristics of megacheck, but supports a broader range of tools. You can configure
golangci-lint with `go.lintFlags`, for example to show issues only in new code and to enable all linters:

```javascript
  "go.lintFlags": ["--enable-all", "--new"],
```

You can also use [staticcheck](https://github.com/dominikh/go-tools/tree/master/cmd/staticcheck).

Another alternative of golint is [revive](https://github.com/mgechev/revive). It is extensible, configurable, provides
superset of the rules of golint, and has significantly better performance.

To configure revive, use:

```javascript
  "go.lintFlags": ["-exclude=vendor/...", "-config=${workspaceFolder}/config.toml"]
```

Finally, the result of those linters will show right in the code (locations with suggestions will be underlined), as
well as in the output pane.

## 命令

除了综合编辑功能，扩展还提供了命令面板多个命令与`Go`文件工作:

- `Go: Add Import` to add an import from the list of packages in your `Go` context
- `Go: Current GOPATH` 看到您的当前配置 GOPATH
- `Go: Test at cursor` to run a test at the current cursor position in the active document
- `Go: Test Package` to run all tests in the package containing the active document
- `Go: Test File` to run all tests in the current active document
- `Go: Test Previous` to run the previously run test command
- `Go: Test All Packages in Workspace` to run all tests in the current workspace
- `Go: Generate Unit Tests For Package` Generates unit tests for the current package
- `Go: Generate Unit Tests For File` Generates unit tests for the current file
- `Go: Generate Unit Tests For Function` Generates unit tests for the selected function in the current file
- `Go: Install Tools` Installs/updates all the Go tools that the extension depends on
- `Go: Add Tags` Adds configured tags to selected struct fields.
- `Go: Remove Tags` Removes configured tags from selected struct fields.
- `Go: Generate Interface Stubs` Generates method stubs for given interface
- `Go: Fill Struct` Fills struct literal with default values
- `Go: Run on Go Playground` Upload the current selection or file to the Go Playground

您可以在命令面板访问所有上述命令的 (`Cmd+Shift+P` or `Ctrl+Shift+P`).

一些这些都是在编辑器上下文菜单作为一个实验性的功能可用。为了控制这些命令在编辑器上下文菜单中显示出来, 更新设
置`go.editorContextMenuCommands`.

## _可选的_: 调试

使用调试器, 你目前必须手动[安装 delve](https://github.com/derekparker/delve/tree/master/Documentation/installation). 更
多阅读[使用 VS Code 调试 Go 代码](https://github.com/Microsoft/vscode-go/wiki/Debugging-Go-code-using-VS-Code).

### 调试 WSL

如果在 Windows 上使用 WSL, 你将需要 WSL 2 Linux 内核.

See [WSL 2 Installation](https://docs.microsoft.com/en-us/windows/wsl/wsl2-install) and note the Window 10 build version
requirements.

### 远程调试

要使用 VS Code 远程调试,
读[远程调试](https://github.com/Microsoft/vscode-go/wiki/Debugging-Go-code-using-VS-Code#remote-debugging).
