package main

import (
	"encoding/json"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"log"

	"github.com/syndtr/goleveldb/leveldb"
)

// 读取 Excel 数据
func readExcel(filePath string) ([][]string, error) {
	f, err := excelize.OpenFile(filePath)
	if err != nil {
		return nil, err
	}

	// 获取第一个 Sheet 名称
	sheetName := f.GetSheetName(0)
	rows := f.GetRows(sheetName)
	if err != nil {
		return nil, err
	}

	return rows, nil
}

// 存储数据到 LevelDB
func saveToLevelDB(dbPath string, data [][]string) error {
	db, err := leveldb.OpenFile(dbPath, nil)
	if err != nil {
		return err
	}
	defer db.Close()

	for i, row := range data {
		// 转换为 JSON 存储
		rowData, _ := json.Marshal(row)
		key := fmt.Sprintf("row_%d", i) // key 设为 row_xxx 形式
		err := db.Put([]byte(key), rowData, nil)
		if err != nil {
			return err
		}
	}

	return nil
}

// 读取 LevelDB 数据
func readFromLevelDB(dbPath string) {
	db, err := leveldb.OpenFile(dbPath, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	iter := db.NewIterator(nil, nil)
	for iter.Next() {
		fmt.Printf("Key: %s, Value: %s\n", iter.Key(), iter.Value())
	}
	iter.Release()
}

func main() {
	excelPath := "data.xlsx"
	dbPath := "leveldb_data"

	// 读取 Excel
	rows, err := readExcel(excelPath)
	if err != nil {
		log.Fatal("Excel 读取失败:", err)
	}

	fmt.Println("Excel 数据:", rows)

	// 存储到 LevelDB
	err = saveToLevelDB(dbPath, rows)
	if err != nil {
		log.Fatal("LevelDB 存储失败:", err)
	}

	fmt.Println("数据成功存入 LevelDB!")

	// 读取 LevelDB 数据
	readFromLevelDB(dbPath)
}
