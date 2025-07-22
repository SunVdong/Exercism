package transpose

import "strings"

func Transpose(input []string) []string {
	if len(input) == 0 {
		return []string{}
	}

	// 找到最长的字符串长度
	maxLen := 0
	for _, s := range input {
		if len(s) > maxLen {
			maxLen = len(s)
		}
	}

	// 创建结果切片
	result := make([]string, maxLen)

	// 对每一列进行处理
	for col := 0; col < maxLen; col++ {
		var sb strings.Builder

		// 对每一行进行处理
		for row := 0; row < len(input); row++ {
			if col < len(input[row]) {
				// 如果该位置有字符，直接添加
				sb.WriteByte(input[row][col])
			} else {
				// 如果该位置没有字符，需要检查是否需要添加空格
				// 只有当后面的行在这一列有字符时，才需要添加空格
				needSpace := false
				for nextRow := row + 1; nextRow < len(input); nextRow++ {
					if col < len(input[nextRow]) {
						needSpace = true
						break
					}
				}
				if needSpace {
					sb.WriteByte(' ')
				}
			}
		}
		result[col] = sb.String()
	}

	return result
}
