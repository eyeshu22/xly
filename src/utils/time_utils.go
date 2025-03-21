package utils

import (
	"fmt"
	"time"
)

func Convert() string {
	// 获取当前时间
	currentTime := time.Now()
	// 格式化时间为字符串
	layout := "2006-01-02 15:04:05" // 使用 Go 特定的时间格式
	timeStr := currentTime.Format(layout)

	fmt.Println("转换后的时间:", timeStr)
	return timeStr
}
