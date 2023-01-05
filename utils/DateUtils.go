package utils

import (
	"fmt"
	"time"
)

func GetDateStr(timeT int64) string {
	currentTime := time.Now() //获取当前时间，类型是Go的时间类型Time
	fmt.Println(currentTime)
	return currentTime.Format("yyyy-mm-dd")
}
