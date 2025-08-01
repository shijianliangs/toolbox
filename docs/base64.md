# Base64编码解码命令

## 功能特性

- `encode` - 将数据编码为base64格式
- `decode` - 将base64数据解码为原始格式

## 使用示例

### 编码字符串
```bash
echo "Hello World" | ./toolbox base64 encode
# 输出: SGVsbG8gV29ybGQK
```

### 解码字符串
```bash
echo "SGVsbG8gV29ybGQK" | ./toolbox base64 decode
# 输出: Hello World
```

### 编码文件
```bash
./toolbox base64 encode myfile.txt
```

### 解码文件
```bash
./toolbox base64 decode encoded.txt
```

### 使用URL安全编码
```bash
echo "Hello+World/Test=" | ./toolbox base64 encode --url
```

## 参数说明

- `data` - 要编码或解码的数据（可选，默认从标准输入读取）
- `--url` - 使用URL安全的Base64编码（仅encode命令）

## 选项说明

### encode 命令选项
- `--url` - 使用URL安全的Base64编码，将 `+` 替换为 `-`，`/` 替换为 `_`

### decode 命令选项
- 自动检测是否为URL安全的Base64编码

## 输入输出

### 输入方式
- 命令行参数: `./toolbox base64 encode "Hello World"`
- 标准输入: `echo "Hello World" | ./toolbox base64 encode`
- 文件输入: `./toolbox base64 encode myfile.txt`

### 输出方式
- 标准输出: 编码/解码结果直接输出到终端
- 文件输出: 可通过重定向保存到文件

## 注意事项

- 支持从标准输入读取数据
- 支持文件输入和输出
- URL安全编码适用于在URL中传输Base64数据
- 自动处理换行符和特殊字符 