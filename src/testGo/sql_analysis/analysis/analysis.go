package main

import (
	"fmt"
	"github.com/pingcap/tidb/parser"
	_ "github.com/pingcap/tidb/parser/test_driver"
)

func main() {
	p := parser.New()
	reservedKws := []string{
		"add", "all", "alter", "analyze", "and", "as", "asc", "between", "bigint",
		"binary", "blob", "both", "by", "call", "cascade", "case", "change", "character", "check", "collate",
		"column", "constraint", "convert", "create", "cross", "current_date", "current_time",
		"current_timestamp", "current_user", "database", "databases", "day_hour", "day_microsecond",
		"day_minute", "day_second", "decimal", "default", "delete", "desc", "describe",
		"distinct", "distinctRow", "div", "double", "drop", "dual", "else", "enclosed", "escaped",
		"exists", "explain", "false", "float", "fetch", "for", "force", "foreign", "from",
		"fulltext", "grant", "group", "having", "hour_microsecond", "hour_minute",
		"hour_second", "if", "ignore", "in", "index", "infile", "inner", "insert", "int", "into", "integer",
		"interval", "is", "join", "key", "keys", "kill", "leading", "left", "like", "ilike", "limit", "lines", "load",
		"localtime", "localtimestamp", "lock", "longblob", "longtext", "mediumblob", "maxvalue", "mediumint", "mediumtext",
		"minute_microsecond", "minute_second", "mod", "not", "no_write_to_binlog", "null", "numeric",
		"on", "option", "optionally", "or", "order", "outer", "partition", "precision", "primary", "procedure", "range", "read", "real", "recursive",
		"references", "regexp", "rename", "repeat", "replace", "revoke", "restrict", "right", "rlike",
		"schema", "schemas", "second_microsecond", "select", "set", "show", "smallint",
		"starting", "table", "terminated", "then", "tinyblob", "tinyint", "tinytext", "to",
		"trailing", "true", "union", "unique", "unlock", "unsigned",
		"update", "use", "using", "utc_date", "values", "varbinary", "varchar",
		"when", "where", "write", "xor", "year_month", "zerofill",
		"generated", "virtual", "stored", "usage",
		"delayed", "high_priority", "low_priority",
		"cumeDist", "denseRank", "firstValue", "lag", "lastValue", "lead", "nthValue", "ntile",
		"over", "percentRank", "rank", "row", "rows", "rowNumber", "window", "linear",
		"match", "until", "placement", "tablesample", "failedLoginAttempts", "passwordLockTime",
		// TODO: support the following keywords
		// "with",
	}
	for _, kw := range reservedKws {
		src := fmt.Sprintf("SELECT * FROM db.%s;", kw)
		stmtNode0, err := p.ParseOneStmt(src, "", "")
		if err != nil {
			println(err)
		}
		//println(stmtNode0)
		println(stmtNode0.Text())
		src = fmt.Sprintf("SELECT * FROM %s.desc", kw)
		stmtNode1, err := p.ParseOneStmt(src, "", "")
		if err != nil {
			println(err)
		}
		//println(stmtNode1)
		println(stmtNode1.Text())
		src = fmt.Sprintf("SELECT t.%s FROM t", kw)
		stmtNode2, err := p.ParseOneStmt(src, "", "")
		if err != nil {
			println(err)
		}
		//println(stmtNode2)
		println(stmtNode2.Text())

	}
}

func main2() {

	i := new(string)
	*i = "123123123"
	println(i)
	println(*i)
	println(&i)

	b := *i
	println(b)
	println(*&b)
	println(&b)

	c := &b

	println(c)
	println(&c)
	println(*&c)
	println(*c)
}

//defer 常用于 捕获 panic 并恢复程序，避免程序崩溃。

func main1() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("捕获异常:", r)
		}
	}()

	fmt.Println("程序开始")
	panic("发生错误")       // 触发 panic
	fmt.Println("程序结束") // 不会执行

}
