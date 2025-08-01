/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/spf13/cobra"
)

// base64Cmd represents the base64 command
var base64Cmd = &cobra.Command{
	Use:   "base64",
	Short: "Base64编码与解码工具",
	Long: `Base64编码与解码工具集合，包括:
- encode: 将数据编码为base64格式
- decode: 将base64数据解码为原始格式`,
}

// encodeCmd represents the encode command
var encodeCmd = &cobra.Command{
	Use:   "encode [文件]",
	Short: "编码为base64格式",
	Long:  `将输入数据编码为base64格式。如果未指定文件，则从标准输入读取。`,
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var data []byte
		var err error

		if len(args) == 0 {
			// 从标准输入读取
			data, err = ioutil.ReadAll(os.Stdin)
		} else {
			// 从文件读取
			data, err = ioutil.ReadFile(args[0])
		}

		if err != nil {
			fmt.Printf("错误: %v\n", err)
			return
		}

		// 获取 URL 安全选项
		urlSafe, _ := cmd.Flags().GetBool("url")

		var encoded string
		if urlSafe {
			encoded = base64URLEncode(data)
		} else {
			encoded = base64Encode(data)
		}

		fmt.Println(encoded)
	},
}

// decodeCmd represents the decode command
var decodeCmd = &cobra.Command{
	Use:   "decode [文件]",
	Short: "从base64格式解码",
	Long:  `将base64格式的数据解码为原始格式。如果未指定文件，则从标准输入读取。`,
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var data []byte
		var err error

		if len(args) == 0 {
			// 从标准输入读取
			data, err = ioutil.ReadAll(os.Stdin)
		} else {
			// 从文件读取
			data, err = ioutil.ReadFile(args[0])
		}

		if err != nil {
			fmt.Printf("错误: %v\n", err)
			return
		}

		// 获取 URL 安全选项
		urlSafe, _ := cmd.Flags().GetBool("url")

		var decoded []byte
		if urlSafe {
			decoded, err = base64URLDecode(string(data))
		} else {
			decoded, err = base64Decode(string(data))
		}

		if err != nil {
			fmt.Printf("解码错误: %v\n", err)
			return
		}

		// 输出解码后的数据
		fmt.Print(string(decoded))
	},
}

// base64Encode encodes data to base64 string
func base64Encode(data []byte) string {
	return base64.StdEncoding.EncodeToString(data)
}

// base64Decode decodes base64 string to data
func base64Decode(encoded string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(encoded)
}

// base64URLEncode encodes data to URL-safe base64 string
func base64URLEncode(data []byte) string {
	return base64.URLEncoding.EncodeToString(data)
}

// base64URLDecode decodes URL-safe base64 string to data
func base64URLDecode(encoded string) ([]byte, error) {
	return base64.URLEncoding.DecodeString(encoded)
}

func init() {
	rootCmd.AddCommand(base64Cmd)
	base64Cmd.AddCommand(encodeCmd)
	base64Cmd.AddCommand(decodeCmd)

	// 添加 URL 安全选项
	encodeCmd.Flags().BoolP("url", "u", false, "使用URL安全的base64编码")
	decodeCmd.Flags().BoolP("url", "u", false, "使用URL安全的base64解码")
}
