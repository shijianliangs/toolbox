# Toolbox - Golang 工具集合

一个基于 Cobra 框架的 Golang 常用工具集合，提供字符串处理、文件操作、Base64编码解码等实用功能。

## 可用命令

| 命令 | 说明 |
|------|------|
| [string](docs/string.md) | 字符串处理工具，支持反转、大小写转换、去除空白等操作 |
| [file](docs/file.md) | 文件操作工具，支持查看文件信息、检查存在性、计算大小等 |
| [base64](docs/base64.md) | Base64编码解码工具，支持标准编码和URL安全编码 |

## 快速开始

```bash
# 字符串处理
./toolbox string reverse "hello world"
./toolbox string upper "hello world"

# 文件操作
./toolbox file info go.mod
./toolbox file exists myfile.txt

# Base64编码解码
echo "Hello World" | ./toolbox base64 encode
echo "SGVsbG8gV29ybGQK" | ./toolbox base64 decode
```

## 项目结构

```
toolbox/
├── cmd/           # Cobra 命令定义
│   ├── root.go    # 根命令
│   ├── string.go  # 字符串处理命令
│   ├── file.go    # 文件操作命令
│   └── base64.go  # Base64编码解码命令
├── pkg/utils/     # 工具函数包
│   ├── string.go  # 字符串工具
│   ├── file.go    # 文件工具
│   └── base64.go  # Base64工具
├── docs/          # 详细文档
│   ├── string.md  # 字符串处理文档
│   ├── file.md    # 文件操作文档
│   └── base64.md  # Base64文档
├── main.go        # 程序入口
├── go.mod         # Go 模块定义
├── Makefile       # 构建脚本
└── README.md      # 项目文档
```

## 依赖

- [github.com/spf13/cobra](https://github.com/spf13/cobra) - CLI 框架

## 许可证

MIT License