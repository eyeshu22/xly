package excel

import (
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/extrame/xls"
	"github.com/tealeg/xlsx"
)

func readExcel() {
	excelize.OpenFile("")
	xls.Open("", "")
	xlsx.ColLettersToIndex("")
}
