package grep

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func Search(pattern string, flags, files []string) []string {
	var nFlag, lFlag, iFlag, vFlag, xFlag, mfileFlag bool

	for _, flag := range flags {
		switch flag {
		case "-n":
			nFlag = true
		case "-l":
			lFlag = true
		case "-i":
			iFlag = true
		case "-v":
			vFlag = true
		case "-x":
			xFlag = true
		}
	}
	if len(files) > 1 {
		mfileFlag = true
	}

	res := []string{}
	for _, file := range files {
		fileHandle, _ := os.Open(file)
		defer fileHandle.Close()

		// 创建一个新的 Scanner
		scanner := bufio.NewScanner(fileHandle)
		var match bool
		var lineCount = 0
		// 按行读取文件
		for scanner.Scan() {
			lineCount++
			lineOriginal := scanner.Text()
			line := lineOriginal
			if iFlag {
				line = strings.ToLower(lineOriginal)
				pattern = strings.ToLower(pattern)
			}
			if xFlag {
				match = line == pattern
			} else {
				match = strings.Contains(line, pattern)
			}
			if vFlag {
				match = !match
			}
			if match {
				var linePre, filePre string
				if nFlag {
					linePre = strconv.Itoa(lineCount) + ":"
				}
				if mfileFlag {
					filePre = file + ":"
				}

				if lFlag {
					res = append(res, file)
					break
				}

				res = append(res, filePre+linePre+lineOriginal)
			}
		}
	}

	return res
}
