# 字符串处理命令

## 功能特性

- `reverse` - 反转字符串
- `upper` - 转换为大写
- `lower` - 转换为小写  
- `trim` - 去除首尾空白字符

## 使用示例

### 反转字符串
```bash
./toolbox string reverse "hello world"
# 输出: dlrow olleh
```

### 转换为大写
```bash
./toolbox string upper "hello world"  
# 输出: HELLO WORLD
```

### 转换为小写
```bash
./toolbox string lower "HELLO WORLD"
# 输出: hello world
```

### 去除空白字符
```bash
./toolbox string trim "  hello world  "
# 输出: hello world
```

## 参数说明

- `text` - 要处理的文本字符串（必需参数）

## 注意事项

- 所有命令都支持从标准输入读取数据
- 处理结果会输出到标准输出
- 支持Unicode字符的处理 