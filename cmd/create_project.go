package cmd

import (
	"fmt"
	createproj "tools/internal/create_project"

	"github.com/spf13/cobra"
)

var repositoryURL string
var projectName string
var configPath string

var createProjectCmd = &cobra.Command{
	Use:   "create",
	Short: "创建新工程",
	Long:  "创建从远程仓库创建新工程",
	Run: func(cmd *cobra.Command, args []string) {
		err := createproj.Create(repositoryURL, projectName)
		if err != nil {
			fmt.Println(err)
			return
		}
	},
}

func init() {
	createProjectCmd.Flags().StringVarP(&repositoryURL, "repository", "r", "", "请输入远程仓库地址")
	createProjectCmd.Flags().StringVarP(&projectName, "name", "n", "", "请输入远程仓库地址")
	createProjectCmd.Flags().StringVarP(&configPath, "config", "c", "", "请输入配置文件地址")
}
