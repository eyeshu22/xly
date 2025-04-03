package main

import (
	"fmt"
	"github.com/xwb1989/sqlparser"
	"log"
	"strings"
)

// 移除 SQL 语句中的注释
func removeComments(sql string) string {
	lines := strings.Split(sql, "\n")
	var result []string
	for _, line := range lines {
		// 去除单行注释 `--`
		if strings.HasPrefix(strings.TrimSpace(line), "--") {
			continue
		}
		// 去除 `#` 开头的注释（MySQL 兼容）
		if strings.HasPrefix(strings.TrimSpace(line), "#") {
			continue
		}
		// 处理 `/* ... */` 形式的多行注释
		if strings.Contains(line, "/*") && strings.Contains(line, "*/") {
			line = line[:strings.Index(line, "/*")] + line[strings.Index(line, "*/")+2:]
		}
		result = append(result, line)
	}
	return strings.Join(result, "\n")
}

// 处理 SQL 解析
func parseSQL(sql string) {
	// 预处理：移除 SQL 语句中的注释
	cleanSQL := removeComments(sql)

	// 解析 SQL
	stmt, err := sqlparser.Parse(cleanSQL)
	if err != nil {
		log.Fatalf("SQL 解析失败: %v", err)
	}

	// 根据解析结果处理不同 SQL 语句
	switch stmt := stmt.(type) {
	case *sqlparser.Select:
		fmt.Println("解析到 SELECT 语句:", sqlparser.String(stmt))
	case *sqlparser.Insert:
		fmt.Println("解析到 INSERT 语句:", sqlparser.String(stmt))
	case *sqlparser.Update:
		fmt.Println("解析到 UPDATE 语句:", sqlparser.String(stmt))
	case *sqlparser.Delete:
		fmt.Println("解析到 DELETE 语句:", sqlparser.String(stmt))
	case *sqlparser.DDL:
		fmt.Println("解析到 CREATE TABLE 语句:", sqlparser.String(stmt))
	//case *sqlparser.DropTable:
	//	fmt.Println("解析到 DROP TABLE 语句:", sqlparser.String(stmt))
	default:
		fmt.Println("未识别的 SQL 语句:", sqlparser.String(stmt))
	}
}
func main() {
	sql := "SELECT name, age FROM users WHERE age > 30 AND city = 'Beijing'"
	parseSQL(sql)
	stmt, err := sqlparser.Parse(sql)
	if err != nil {
		fmt.Println("SQL 解析失败:", err)
		return
	}

	fmt.Println("解析成功:", sqlparser.String(stmt))
}
