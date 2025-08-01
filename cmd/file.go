/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/spf13/cobra"
)

// fileCmd represents the file command
var fileCmd = &cobra.Command{
	Use:   "file",
	Short: "文件操作工具",
	Long: `文件操作工具集合，包括:
- info: 显示文件信息
- exists: 检查文件是否存在
- size: 显示文件大小`,
}

// infoCmd represents the info command
var infoCmd = &cobra.Command{
	Use:   "info [文件路径]",
	Short: "显示文件信息",
	Long:  `显示文件的详细信息，包括大小、权限和修改时间。`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		filePath := args[0]
		info, err := os.Stat(filePath)
		if err != nil {
			fmt.Printf("错误: %v\n", err)
			return
		}

		fmt.Printf("文件: %s\n", filePath)
		fmt.Printf("大小: %d 字节\n", info.Size())
		fmt.Printf("权限: %s\n", info.Mode())
		fmt.Printf("修改时间: %s\n", info.ModTime().Format(time.RFC3339))
		fmt.Printf("是否目录: %t\n", info.IsDir())
	},
}

// existsCmd represents the exists command
var existsCmd = &cobra.Command{
	Use:   "exists [文件路径]",
	Short: "检查文件是否存在",
	Long:  `检查文件或目录是否存在。`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		filePath := args[0]
		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			fmt.Printf("文件 '%s' 不存在\n", filePath)
		} else {
			fmt.Printf("文件 '%s' 存在\n", filePath)
		}
	},
}

// sizeCmd represents the size command
var sizeCmd = &cobra.Command{
	Use:   "size [文件路径]",
	Short: "显示文件大小",
	Long:  `显示文件的大小（字节数）。`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		filePath := args[0]
		info, err := os.Stat(filePath)
		if err != nil {
			fmt.Printf("错误: %v\n", err)
			return
		}

		if info.IsDir() {
			// Calculate directory size
			var size int64
			err := filepath.Walk(filePath, func(_ string, info os.FileInfo, err error) error {
				if err != nil {
					return err
				}
				if !info.IsDir() {
					size += info.Size()
				}
				return nil
			})
			if err != nil {
				fmt.Printf("计算目录大小时发生错误: %v\n", err)
				return
			}
			fmt.Printf("%d 字节\n", size)
		} else {
			fmt.Printf("%d 字节\n", info.Size())
		}
	},
}

func init() {
	rootCmd.AddCommand(fileCmd)
	fileCmd.AddCommand(infoCmd)
	fileCmd.AddCommand(existsCmd)
	fileCmd.AddCommand(sizeCmd)
}
