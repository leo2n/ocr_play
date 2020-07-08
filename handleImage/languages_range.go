package handleImage

import "strings"

var support_languages = []string{"eng", "chi_sim", "jpn"}

// 确认传入的语言标识符是否在支持的范围之内
func isInSupportLanguagesRange(input string) bool {
	inputLanguages := strings.Split(input, ";")
	for _, v := range inputLanguages {
		if isInList(v, support_languages) == false {
			return false
		}
	}
	return true
}

// 单个字符串是否在列表支持范围之内
func isInList(x string, list []string) bool {
	for _, v := range list {
		if x == v {
			return true
		}
	}
	return false
}