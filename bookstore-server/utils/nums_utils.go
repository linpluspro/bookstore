package utils

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

// 数值是否存在
func NumsInList(num int, nums []int) bool {
	for _, s := range nums {
		if s == num {
			return true
		}
	}
	return false
}

// 随机数
func GenValidateCode(width int) string {
	numeric := [10]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	r := len(numeric)
	rand.Seed(time.Now().UnixNano())

	var sb strings.Builder
	for i := 0; i < width; i++ {
		_, err := fmt.Fprintf(&sb, "%d", numeric[rand.Intn(r)])
		if err != nil {
			return ""
		}
	}
	return sb.String()
}

// 生成一个包含当前时间戳和 4 个随机数的字符串
func GenOrderNo() string {
	numeric := [10]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	r := len(numeric)
	rand.Seed(time.Now().UnixNano())

	var sb strings.Builder
	for i := 0; i < 4; i++ {
		_, err := fmt.Fprintf(&sb, "%d", numeric[rand.Intn(r)])
		if err != nil {
			return ""
		}
	}
	timestamp := strconv.FormatInt(time.Now().UnixNano()/1e6, 10)
	return timestamp + sb.String()
}

// 字符串转换为切片
func StrToInt(strNum string) (nums []int) {
	strNums := strings.Split(strNum, ",")
	for _, s := range strNums {
		i, _ := strconv.Atoi(s)
		nums = append(nums, i)
	}
	return
}
