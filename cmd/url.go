/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"io"
	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// urlCmd represents the url command
var urlCmd = &cobra.Command{
	Use:   "url",
	Short: "URL编码与解码工具",
	Long: `URL编码与解码工具集合，包括:
- encode: 将字符串编码为URL格式
- decode: 将URL格式的字符串解码为原始格式`,
}

// urlEncodeCmd represents the url encode command
var urlEncodeCmd = &cobra.Command{
	Use:   "encode [字符串]",
	Short: "编码为URL格式",
	Long:  `将输入字符串编码为URL格式。如果未指定字符串，则从标准输入读取。`,
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var input string

		if len(args) == 0 {
			// 从标准输入读取
			data, err := io.ReadAll(os.Stdin)
			if err != nil {
				fmt.Printf("错误: %v\n", err)
				return
			}
			input = string(data)
		} else {
			// 从命令行参数读取
			input = args[0]
		}

		encoded := urlEncode(input)
		fmt.Println(encoded)
	},
}

// urlDecodeCmd represents the url decode command
var urlDecodeCmd = &cobra.Command{
	Use:   "decode [字符串]",
	Short: "从URL格式解码",
	Long:  `将URL格式的字符串解码为原始格式。如果未指定字符串，则从标准输入读取。`,
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var input string
		var err error

		if len(args) == 0 {
			// 从标准输入读取
			data, err := io.ReadAll(os.Stdin)
			if err != nil {
				fmt.Printf("错误: %v\n", err)
				return
			}
			input = string(data)
		} else {
			// 从命令行参数读取
			input = args[0]
		}

		decoded, err := urlDecode(input)
		if err != nil {
			fmt.Printf("解码错误: %v\n", err)
			return
		}

		fmt.Print(decoded)
	},
}

// urlEncode encodes string to URL format
func urlEncode(input string) string {
	return url.QueryEscape(input)
}

// urlDecode decodes URL format string to original string
func urlDecode(encoded string) (string, error) {
	return url.QueryUnescape(encoded)
}

func init() {
	rootCmd.AddCommand(urlCmd)
	urlCmd.AddCommand(urlEncodeCmd)
	urlCmd.AddCommand(urlDecodeCmd)
}
