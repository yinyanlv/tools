package cmd

import (
	"log"
	"tools/internal/sql2struct"

	"github.com/spf13/cobra"
)

var username string
var password string
var host string
var charset string
var dbType string
var dbName string
var tableName string

var sqlCmd = &cobra.Command{
	Use:   "sql",
	Short: "sql转换处理",
	Long:  "sql转换处理",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

var sql2StructCmd = &cobra.Command{
	Use:   "struct",
	Short: "sql转struct model",
	Long:  "sql转struct model",
	Run: func(cmd *cobra.Command, args []string) {
		dbInfo := &sql2struct.DBInfo{
			DBType:   dbType,
			Host:     host,
			Username: username,
			Password: password,
			Charset:  charset,
		}
		dbModel := sql2struct.NewDBModel(dbInfo)
		err := dbModel.Connect()
		if err != nil {
			log.Fatalf("dbModel.Connect err: %v", err)
		}
		cols, err := dbModel.GetColumns(dbName, tableName)
		if err != nil {
			log.Fatalf("dbModel.GetColumns err: %v", err)
		}
		tpl := sql2struct.NewStructTemplate()
		tplCols := tpl.MapColumns(cols)
		err = tpl.Generate(tableName, tplCols)
		if err != nil {
			log.Fatalf("template.Genereate err: %v", err)
		}
	},
}

func init() {
	sqlCmd.AddCommand(sql2StructCmd)
	sql2StructCmd.Flags().StringVarP(&username, "username", "u", "", "请输入数据库账号")
	sql2StructCmd.Flags().StringVarP(&password, "password", "p", "", "请输入数据库账号密码")
	sql2StructCmd.Flags().StringVarP(&host, "host", "", "127.0.0.1:3306", "请输入数据库host")
	sql2StructCmd.Flags().StringVarP(&charset, "charset", "c", "utf8mb4", "请输入数据库编码")
	sql2StructCmd.Flags().StringVarP(&dbType, "type", "", "mysql", "请输入数据库类型")
	sql2StructCmd.Flags().StringVarP(&dbName, "database", "d", "", "请输入数据库名称")
	sql2StructCmd.Flags().StringVarP(&tableName, "tableName", "t", "", "请输入表名称")
}
