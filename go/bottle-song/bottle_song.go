package bottlesong

import (
	"fmt"
	"strings"
)

// numberToWords 将数字转换为对应的英语单词
func numberToWords(n int) string {
	words := []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten"}
	if n >= 0 && n < len(words) {
		return words[n]
	}
	return ""
}

// capitalizeFirstLetter 将字符串的第一个字母大写
func capitalizeFirstLetter(s string) string {
	if len(s) == 0 {
		return s
	}
	return strings.ToUpper(string(s[0])) + s[1:]
}

// Recite 生成 "Ten Green Bottles" 歌曲的歌词
func Recite(startBottles, takeDown int) []string {
	var lyrics []string

	for i := startBottles; i > startBottles-takeDown; i-- {
		currentWord := numberToWords(i)
		nextWord := numberToWords(i - 1)

		if i > 1 {
			lyrics = append(lyrics, fmt.Sprintf("%s green bottles hanging on the wall,", capitalizeFirstLetter(currentWord)))
			lyrics = append(lyrics, fmt.Sprintf("%s green bottles hanging on the wall,", capitalizeFirstLetter(currentWord)))
			lyrics = append(lyrics, "And if one green bottle should accidentally fall,")
			if i-1 == 1 {
				lyrics = append(lyrics, "There'll be one green bottle hanging on the wall.")
			} else {
				lyrics = append(lyrics, fmt.Sprintf("There'll be %s green bottles hanging on the wall.", nextWord))
			}
		} else {
			lyrics = append(lyrics, "One green bottle hanging on the wall,")
			lyrics = append(lyrics, "One green bottle hanging on the wall,")
			lyrics = append(lyrics, "And if one green bottle should accidentally fall,")
			lyrics = append(lyrics, "There'll be no green bottles hanging on the wall.")
		}
		lyrics = append(lyrics, "") // 每节歌词之间空一行
	}

	if len(lyrics) > 0 {
		lyrics = lyrics[:len(lyrics)-1] // 去掉最后的空行
	}

	return lyrics
}
