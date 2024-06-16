package utils

import (
	"regexp"
)

// 定义邮箱的正则表达式
// 1. 邮箱必须包含@符号
// 2. 邮箱前缀必须是字母或数字
// 3. 邮箱后缀必须是字母或数字
// 4. 邮箱长度不能超过255字节
// 5. 邮箱不能包含空格
// 6. 邮箱不能包含特殊字符
// 7. 邮箱不能包含连续的点号
// 8. 邮箱不能包含连续的下划线
// 9. 邮箱不能包含连续的@符号

func IsEmailValid(email string) bool {

	const emailRegex = `^[a-zA-Z0-9]([a-zA-Z0-9._%+-]*)?@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	// 编译正则表达式
	re := regexp.MustCompile(emailRegex)

	// 使用正则表达式进行匹配
	if !re.MatchString(email) {
		return false
	}

	// 检查邮箱长度是否超过255字节
	if len(email) > 255 {
		return false
	}

	// 检查是否包含空格
	if containsWhitespace(email) {
		return false
	}

	// 检查是否包含连续的点号、下划线或@符号
	if containsConsecutiveCharacters(email, ".", "_", "@") {
		return false
	}

	return true
}

// containsWhitespace 检查字符串中是否包含空格
func containsWhitespace(s string) bool {
	for _, r := range s {
		if r == ' ' {
			return true
		}
	}
	return false
}

// containsConsecutiveCharacters 检查字符串中是否包含连续的特定字符
func containsConsecutiveCharacters(s string, chars ...string) bool {
	for _, char := range chars {
		if containsSubstring(s, char+char) {
			return true
		}
	}
	return false
}

// containsSubstring 检查字符串中是否包含特定子串
func containsSubstring(s, substr string) bool {
	return len(s) > 0 && len(substr) > 0 && len(s) >= len(substr) && s[0:len(substr)] == substr || (len(s) > 0 && containsSubstring(s[1:], substr))
}

func main() {
	// 测试邮箱地址
	// testEmails := []string{
	// 	"valid.email@example.com",
	// 	"invalid..email@example.com",
	// 	"invalid__email@example.com",
	// 	"invalid@@example.com",
	// 	"invalid email@example.com",
	// 	"invalid@exa_mple.com",
	// 	"invalid@",
	// 	"@example.com",
	// }

	// for _, email := range testEmails {
	// 	fmt.Printf("Is '%s' a valid email? %v\n", email, IsEmailValid(email))
	// }
}
