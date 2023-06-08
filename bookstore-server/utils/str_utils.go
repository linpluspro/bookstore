package utils

// 截取字符串的前几个字符，并在末尾添加省略号以表示字符串被截断
func SubStrLen(str string, length int) string {
	nameRune := []rune(str)
	if len(str) > length {
		return string(nameRune[:length-1]) + "..."
	}
	return string(nameRune)
}
