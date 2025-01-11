package lib

import "math/rand"

// randBySource 根据提供的字符源生成指定长度的随机字符串。
func randBySource(source []rune, length int) string {
	target := make([]rune, length)
	for i := range target {
		target[i] = source[rand.Intn(len(source))]
	}
	return string(target)
}

// RandNumbers 生成指定长度的随机数字字符串。
func RandNumbers(length int) string {
	source := []rune("0123456789")
	return randBySource(source, length)
}

// RandNumbersLetters 生成指定长度的随机字母和数字组合字符串（包含大小写字母）。
func RandNumbersLetters(length int) string {
	source := []rune("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")
	return randBySource(source, length)
}
