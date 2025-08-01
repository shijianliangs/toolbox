# URL编码解码命令

## 功能特性

- `encode` - 将字符串编码为URL格式
- `decode` - 将URL格式的字符串解码为原始格式

## 使用示例

### 编码字符串
```bash
./toolbox url encode "Hello World"
# 输出: Hello%20World

./toolbox url encode "https://example.com/path?param=value"
# 输出: https%3A%2F%2Fexample.com%2Fpath%3Fparam%3Dvalue
```

### 解码字符串
```bash
./toolbox url decode "Hello%20World"
# 输出: Hello World

./toolbox url decode "https%3A%2F%2Fexample.com%2Fpath%3Fparam%3Dvalue"
# 输出: https://example.com/path?param=value
```

### 从标准输入读取
```bash
echo "Hello World" | ./toolbox url encode
# 输出: Hello%20World

echo "Hello%20World" | ./toolbox url decode
# 输出: Hello World
```

## 参数说明

- `字符串` - 要编码或解码的字符串（可选，默认从标准输入读取）

## 编码规则

### URL编码规则
- 空格编码为 `%20`
- 特殊字符编码为对应的百分号编码
- 字母数字字符保持不变
- 中文字符编码为UTF-8字节序列的百分号编码

### 常见编码对照
| 字符 | 编码 |
|------|------|
| 空格 | `%20` |
| ! | `%21` |
| " | `%22` |
| # | `%23` |
| $ | `%24` |
| % | `%25` |
| & | `%26` |
| ' | `%27` |
| ( | `%28` |
| ) | `%29` |
| * | `%2A` |
| + | `%2B` |
| , | `%2C` |
| / | `%2F` |
| : | `%3A` |
| ; | `%3B` |
| = | `%3D` |
| ? | `%3F` |
| @ | `%40` |
| [ | `%5B` |
| ] | `%5D` |

## 输入输出

### 输入方式
- 命令行参数: `./toolbox url encode "Hello World"`
- 标准输入: `echo "Hello World" | ./toolbox url encode`

### 输出方式
- 标准输出: 编码/解码结果直接输出到终端
- 文件输出: 可通过重定向保存到文件

## 注意事项

- 使用Go标准库的 `url.QueryEscape` 和 `url.QueryUnescape` 函数
- 支持UTF-8编码的字符串
- 自动处理特殊字符和空格
- 适用于URL参数、查询字符串等场景
- 解码时会自动处理各种百分号编码格式 