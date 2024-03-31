package util

import (
	"fmt"
	"regexp"
	"strconv"
	"time"
)

var closeLog = true

func ConvertHexToColor(hexColor string) string {
	// 匹配十六进制颜色值的正则表达式
	regex := regexp.MustCompile(`^#(?:[0-9a-fA-F]{3}){1,2}$`)

	// 检查输入的颜色值是否合法
	if !regex.MatchString(hexColor) {
		return fmt.Sprintf("Invalid color: %s", hexColor)
	}

	// 提取颜色的 R、G、B 分量
	r, _ := strconv.ParseInt(hexColor[1:3], 16, 64)
	g, _ := strconv.ParseInt(hexColor[3:5], 16, 64)
	b, _ := strconv.ParseInt(hexColor[5:7], 16, 64)

	// 根据颜色分量生成 ANSI 转义码
	return fmt.Sprintf("\033[38;2;%d;%d;%dm", r, g, b)
}

var defaultHexColor = "#00FF00" // 默认的十六进制颜色值

func PrintLogWithColor(message string, hexColor ...string) {

	if closeLog {
		_hexColor := defaultHexColor
		if len(hexColor) > 0 {
			_hexColor = hexColor[0]
		}

		colorCode := ConvertHexToColor(_hexColor)
		currentTime := time.Now().Format("2006-01-02 15:04:05")
		// 在终端中输出彩色文本
		//+ message
		fmt.Println(colorCode + currentTime + " " + message)
	}

}

// func _setLogTime(showTime bool) {
// 	if showTime {
// 		log.SetFlags(log.Ldate | log.Ltime)
// 	} else {
// 		log.SetFlags(0) // 不显示时间
// 	}
// }

// func PrintLog(message string, showTime bool) {
// 	_setLogTime(showTime)
// 	log.Println(message)
// }
