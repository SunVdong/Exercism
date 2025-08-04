package railfence

// generateZigzagPattern 生成zigzag模式的轨道索引序列
func generateZigzagPattern(length, rails int) []int {
	if rails == 1 || rails >= length {
		pattern := make([]int, length)
		for i := range pattern {
			pattern[i] = 0
		}
		return pattern
	}

	pattern := make([]int, length)
	rail, dir := 0, 1
	for i := 0; i < length; i++ {
		pattern[i] = rail
		if rail == 0 {
			dir = 1
		} else if rail == rails-1 {
			dir = -1
		}
		rail += dir
	}
	return pattern
}

func Encode(message string, rails int) string {
	length := len(message)
	if rails == 1 || rails >= length {
		return message
	}

	pattern := generateZigzagPattern(length, rails)

	// 预计算每个轨道的字符数量
	counts := make([]int, rails)
	for _, rail := range pattern {
		counts[rail]++
	}

	// 计算每个轨道在结果中的起始位置
	positions := make([]int, rails+1)
	for i := 0; i < rails; i++ {
		positions[i+1] = positions[i] + counts[i]
	}

	// 直接构建结果，避免创建中间切片
	result := make([]rune, length)
	railIndices := make([]int, rails) // 每个轨道当前写入的位置

	for i, char := range message {
		rail := pattern[i]
		pos := positions[rail] + railIndices[rail]
		result[pos] = char
		railIndices[rail]++
	}

	return string(result)
}

func Decode(message string, rails int) string {
	length := len(message)
	if rails == 1 || rails >= length {
		return message
	}

	pattern := generateZigzagPattern(length, rails)

	// 计算每个轨道的字符数量
	counts := make([]int, rails)
	for _, rail := range pattern {
		counts[rail]++
	}

	// 计算每个轨道在编码结果中的起始位置
	positions := make([]int, rails+1)
	for i := 0; i < rails; i++ {
		positions[i+1] = positions[i] + counts[i]
	}

	// 直接构建解码结果
	result := make([]rune, length)
	railIndices := make([]int, rails) // 每个轨道当前读取的位置

	for i := 0; i < length; i++ {
		rail := pattern[i]
		railStart := positions[rail]
		result[i] = rune(message[railStart+railIndices[rail]])
		railIndices[rail]++
	}

	return string(result)
}
