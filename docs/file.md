# 文件操作命令

## 功能特性

- `info` - 显示文件信息
- `exists` - 检查文件是否存在
- `size` - 显示文件或目录大小

## 使用示例

### 查看文件信息
```bash
./toolbox file info go.mod
```

### 检查文件是否存在
```bash
./toolbox file exists myfile.txt
```

### 查看文件大小
```bash
./toolbox file size go.mod
```

## 参数说明

- `path` - 文件或目录路径（必需参数）

## 输出格式

### file info 输出
```
文件名: go.mod
大小: 1234 bytes
修改时间: 2024-01-01 12:00:00
权限: -rw-r--r--
```

### file exists 输出
- 文件存在时输出: `true`
- 文件不存在时输出: `false`

### file size 输出
- 文件: `1234 bytes`
- 目录: `Total size: 5678 bytes (5 files)`

## 注意事项

- 支持相对路径和绝对路径
- 目录大小计算会递归统计所有子文件和子目录
- 文件信息包含基本属性如大小、修改时间、权限等 