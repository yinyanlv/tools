package cmd

import (
	"fmt"
	"log"
	createproj "tools/internal/create_project"

	"github.com/spf13/cobra"
)

var projectName string
var repositoryURL string

var createProjectCmd = &cobra.Command{
	Use:   "create",
	Short: "创建新工程",
	Long:  "创建从远程仓库创建新工程",
	Run: func(cmd *cobra.Command, args []string) {
		if repositoryURL == "" {
			log.Fatalf("远程仓库地址不可为空！")
			return
		}
		err := createproj.Create(repositoryURL, createproj.TemplateVar{
			ProjectName: projectName,
		})
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("---")
		fmt.Println("项目创建成功！")
	},
}

func init() {
	createProjectCmd.Flags().StringVarP(&projectName, "project", "p", "", "请输入远程仓库地址")
	createProjectCmd.Flags().StringVarP(&repositoryURL, "repository", "r", "", "请输入远程仓库地址")
}
