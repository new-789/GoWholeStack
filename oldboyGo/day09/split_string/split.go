package split_string

import "strings"

//  切割字符串
// example:
// abc, b => [a c]

func Split(s string, sep string) []string {
	var ret = make([]string, 0, strings.Count(s, sep)+1)
	index := strings.Index(s, sep) // 获取 sep 在 s 字符串中的索引
	for index >= 0 {
		ret = append(ret, s[:index])  // 将字符串中 sep 前面的加入到 ret 中
		s = s[index+len(sep):]               // 更新字符串内容
		index = strings.Index(s, sep) // 更新指定切割字符的索引
	}
	ret = append(ret, s)
	return ret
}

// 裴波那契数量
func Fib(n int) int {
	if n < 2 {
		return n
	}
	return Fib(n - 1) * Fib(n - 2)
}