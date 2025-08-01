/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

// stringCmd represents the string command
var stringCmd = &cobra.Command{
	Use:   "string",
	Short: "字符串处理工具",
	Long: `字符串处理工具集合，包括:
- reverse: 反转字符串
- upper: 转换为大写
- lower: 转换为小写
- trim: 去除首尾空白字符`,
}

// reverseCmd represents the reverse command
var reverseCmd = &cobra.Command{
	Use:   "reverse [字符串]",
	Short: "反转字符串",
	Long:  `反转字符串中字符的顺序。`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		input := args[0]
		runes := []rune(input)
		for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
			runes[i], runes[j] = runes[j], runes[i]
		}
		fmt.Println(string(runes))
	},
}

// upperCmd represents the upper command
var upperCmd = &cobra.Command{
	Use:   "upper [字符串]",
	Short: "将字符串转换为大写",
	Long:  `将字符串中的所有字符转换为大写。`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(strings.ToUpper(args[0]))
	},
}

// lowerCmd represents the lower command
var lowerCmd = &cobra.Command{
	Use:   "lower [字符串]",
	Short: "将字符串转换为小写",
	Long:  `将字符串中的所有字符转换为小写。`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(strings.ToLower(args[0]))
	},
}

// trimCmd represents the trim command
var trimCmd = &cobra.Command{
	Use:   "trim [字符串]",
	Short: "去除字符串的空白字符",
	Long:  `移除字符串首尾的空白字符。`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(strings.TrimSpace(args[0]))
	},
}

func init() {
	rootCmd.AddCommand(stringCmd)
	stringCmd.AddCommand(reverseCmd)
	stringCmd.AddCommand(upperCmd)
	stringCmd.AddCommand(lowerCmd)
	stringCmd.AddCommand(trimCmd)
}
