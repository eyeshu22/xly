package main

import (
	"bytes"
	"fmt"
	"github.com/spf13/cobra"
	"strings"
)

func executeCommand(root *cobra.Command, args ...string) (output string, err error) {
	_, output, err = executeCommandC(root, args...)
	return output, err
}

func executeCommandC(root *cobra.Command, args ...string) (c *cobra.Command, output string, err error) {
	buf := new(bytes.Buffer)
	root.SetOut(buf)
	root.SetErr(buf)
	root.SetArgs(args)

	c, err = root.ExecuteC()

	return c, buf.String(), err
}

// 创建根命令
var rootCmd = &cobra.Command{
	Use:   "tools.exe",
	Short: "A simple multi-layer command-line tool",
	//Long:  "This is a multi-layer command-line tool using Cobra in Go.",
}

// task 命令
var taskCmd = &cobra.Command{
	Use:   "task",
	Short: "Configuring Task Scheduling",
	Run: func(cmd *cobra.Command, args []string) {
		folders, _ := cmd.Flags().GetString("folders")
		if folders == "" {
			//cmd.Help()
			cmd.Help()
			return
		}
		fmt.Printf("Hello, %s !\n", folders)
		cmd.Flags().GetString("test")
		//TOOD:待执行代码
	},
}

// task 命令
var affectCmd = &cobra.Command{
	Use:   "affect",
	Short: "Impact Analysis",
	Run: func(cmd *cobra.Command, args []string) {
		keywords, _ := cmd.Flags().GetString("keywords")
		if keywords == "" {
			cmd.Help()
			return
		}
		keywordsArr := strings.Split(keywords, ",")
		for i := 0; i < len(keywordsArr); i++ {
			keywordsArr[i] = strings.TrimSpace(keywordsArr[i])
		}
		fmt.Printf("Hello, %s !\n", keywordsArr)
		//TOOD:待执行代码
	},
}

func main() {
	// 将子命令添加到根命令
	rootCmd.AddCommand(taskCmd, affectCmd)
	// 为 greetCmd 添加 --name 和 -n 标志
	taskCmd.Flags().StringP("folders", "f", "", "D:\\list.xlsx")
	taskCmd.Flags().StringP("test", "c", "false", "true or false")
	affectCmd.Flags().StringP("keywords", "k", "", "t_user name age")
	//output, _ := executeCommand(rootCmd, "task")
	//fmt.Printf(output)
	output, _ := executeCommand(rootCmd, "affect", "-k", "item1,item2,item3")
	fmt.Printf(output)
}
