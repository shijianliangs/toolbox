# Toolbox - Golang 工具集合

一个基于 Cobra 框架的 Golang 常用工具集合，提供字符串处理、文件操作、Base64编码解码等实用功能。

## 可用命令

### 字符串处理 (string)
- `reverse` - 反转字符串
- `upper` - 转换为大写
- `lower` - 转换为小写  
- `trim` - 去除首尾空白字符

**详细文档**: [字符串处理命令](docs/string.md)

### 文件操作 (file)
- `info` - 显示文件信息
- `exists` - 检查文件是否存在
- `size` - 显示文件或目录大小

**详细文档**: [文件操作命令](docs/file.md)

### Base64编码 (base64)
- `encode` - 将数据编码为base64格式
- `decode` - 将base64数据解码为原始格式

**详细文档**: [Base64编码解码命令](docs/base64.md)

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