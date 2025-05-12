package service

import (
	"database/sql"
	"fmt"
	"log"
	"todo-api/config"
	"todo-api/model"

	_ "github.com/denisenkom/go-mssqldb" // 註冊 "sqlserver" driver
)

var todos = []model.Todo{
	{ID: 1, Title: "Learn Go", Done: false},
	{ID: 2, Title: "Build API", Done: true},
}

func GetAllTodos() []model.Todo {
	return todos
}

func GetTodoByID(id int) (model.Todo, bool) {
	for _, t := range todos {
		if t.ID == id {
			return t, true
		}
	}
	return model.Todo{}, false
}

func ConnString() string {
	return fmt.Sprintf(
		"sqlserver://%s:%s@%s?database=%s&encrypt=disable",
		config.GetEnv("MSSQL_USER", "sa"),
		config.GetEnv("MSSQL_PASSWORD", "password"),
		config.GetEnv("MSSQL_SERVER", "localhost"),
		config.GetEnv("MSSQL_DATABASE", "todo"),
	)
}

func GetConnString() string {
	return ConnString()
}

func MssqlTest() string {
	connString := ConnString()

	db, err := sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("無法連線到資料庫，請檢查連線字串或資料庫狀態：", err)
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		log.Fatal("無法 ping 資料庫，請檢查資料庫狀態：", err)
	}

	msg := "✅ 資料庫連線成功"
	log.Println(msg)
	return msg
}

// 獲取所有表名
func GetAllTableNames() ([]string, error) {
	connString := ConnString()

	db, err := sql.Open("sqlserver", connString)
	if err != nil {
		return nil, fmt.Errorf("無法連線到資料庫，請檢查連線字串或資料庫狀態：%v", err)
	}
	defer db.Close()

	var cmd string = `
            SELECT
				'[' + TABLE_CATALOG + '].[' + TABLE_SCHEMA + '].[' + TABLE_NAME + ']' 
				AS FullTableName
			FROM INFORMATION_SCHEMA.TABLES
			WHERE TABLE_TYPE = 'BASE TABLE'
			ORDER BY TABLE_NAME;`

	rows, err := db.Query(cmd)
	if err != nil {
		return nil, fmt.Errorf("查詢資料庫表名失敗：%v", err)
	}
	defer rows.Close()

	var tableNames []string
	for rows.Next() {
		var tableName string
		if err := rows.Scan(&tableName); err != nil {
			return nil, fmt.Errorf("掃描表名失敗：%v", err)
		}
		tableNames = append(tableNames, tableName)
	}

	return tableNames, nil
}

// 獲取表格數據
func GetTableData(tableName string) ([]map[string]interface{}, error) {
	connString := ConnString()

	db, err := sql.Open("sqlserver", connString)
	if err != nil {
		return nil, fmt.Errorf("無法連線到資料庫，請檢查連線字串或資料庫狀態：%v", err)
	}
	defer db.Close()

	var cmd string = fmt.Sprintf("SELECT * FROM %s", tableName)

	rows, err := db.Query(cmd)
	if err != nil {
		return nil, fmt.Errorf("查詢資料庫表名失敗：%v", err)
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		return nil, fmt.Errorf("獲取列名失敗：%v", err)
	}

	var results []map[string]interface{}
	for rows.Next() {
		values := make([]interface{}, len(columns))
		for i := range values {
			values[i] = new(interface{})
		}

		if err := rows.Scan(values...); err != nil {
			return nil, fmt.Errorf("掃描行數據失敗：%v", err)
		}

		rowData := make(map[string]interface{})
		for i, colName := range columns {
			val := *(values[i].(*interface{}))
			rowData[colName] = val
		}
		results = append(results, rowData)
	}

	return results, nil
}
