package sql2struct

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type DBModel struct {
	DBEngine *sql.DB
	DBInfo   *DBInfo
}

type DBInfo struct {
	DBType   string
	Host     string
	Username string
	Password string
	Charset  string
}

type TableColumn struct {
	ColumnName    string
	ColumnType    string
	ColumnKey     string
	ColumnComment string
	IsNullable    string
	DataType      string
}

var DBTypeToStructType = map[string]string{
	"int":        "int32",
	"tinyint":    "int8",
	"smallint":   "int",
	"mediumint":  "int64",
	"bigint":     "int64",
	"bit":        "int",
	"bool":       "bool",
	"enum":       "string",
	"set":        "string",
	"varchar":    "string",
	"char":       "string",
	"tinytext":   "string",
	"mediumtext": "string",
	"text":       "string",
	"longtext":   "string",
	"blob":       "string",
	"tinyblob":   "string",
	"mediumblob": "string",
	"longblob":   "string",
	"date":       "time.Time",
	"datetime":   "time.Time",
	"timestamp":  "time.Time",
	"time":       "time.Time",
	"float":      "float64",
	"double":     "float64",
}

func NewDBModel(info *DBInfo) *DBModel {
	return &DBModel{
		DBInfo: info,
	}
}

func (m *DBModel) Connect() error {
	var err error
	s := "%s:%s@tcp(%s)/information_schema?charset=%s&parseTime=True&loc=Local"
	dsn := fmt.Sprintf(s, m.DBInfo.Username, m.DBInfo.Password, m.DBInfo.Host, m.DBInfo.Charset)
	m.DBEngine, err = sql.Open(m.DBInfo.DBType, dsn)
	if err != nil {
		return err
	}
	return nil
}

func (m *DBModel) GetColumns(dbName, tableName string) ([]*TableColumn, error) {
	query := `
	SELECT COLUMN_NAME, DATA_TYPE, COLUMN_KEY, IS_NULLABLE, COLUMN_TYPE, COLUMN_COMMENT FROM
	COLUMNS WHERE TABLE_SCHEMA = ? AND TABLE_NAME = ?
	`
	rows, err := m.DBEngine.Query(query, dbName, tableName)
	defer rows.Close()

	if err != nil {
		return nil, err
	}

	if rows == nil {
		return nil, errors.New(fmt.Sprintf("数据库 %s 中的 %s 没有元信息", dbName, tableName))
	}

	var cols []*TableColumn
	for rows.Next() {
		var col TableColumn
		err := rows.Scan(&col.ColumnName, &col.DataType, &col.ColumnKey, &col.IsNullable, &col.ColumnType, &col.ColumnComment)
		if err != nil {
			return nil, err
		}
		cols = append(cols, &col)
	}

	return cols, nil
}
