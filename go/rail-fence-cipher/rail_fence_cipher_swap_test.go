package railfence

import (
	"fmt"
	"testing"
)

// action specifies 0 for decode or 1 for encode
type action int

const (
	decode action = iota
	encode
)

// 可视化zigzag模式的数学规律
func visualizeZigzagPattern(message string, rails int) {
	fmt.Printf("\n=== Zigzag模式数学规律分析 ===\n")
	fmt.Printf("消息: %s (长度: %d)\n", message, len(message))
	fmt.Printf("轨道数: %d\n", rails)

	// 创建可视化矩阵
	matrix := make([][]string, rails)
	for i := range matrix {
		matrix[i] = make([]string, len(message))
		for j := range matrix[i] {
			matrix[i][j] = " "
		}
	}

	// 填充zigzag模式
	rail, dir := 0, 1
	for i, char := range message {
		matrix[rail][i] = string(char)
		if rail == 0 {
			dir = 1
		} else if rail == rails-1 {
			dir = -1
		}
		rail += dir
	}

	// 显示矩阵
	fmt.Println("\nZigzag模式矩阵:")
	for i, row := range matrix {
		fmt.Printf("轨道%d: ", i)
		for _, cell := range row {
			if cell != " " {
				fmt.Printf("[%s] ", cell)
			} else {
				fmt.Printf("[ ] ")
			}
		}
		fmt.Println()
	}

	// 分析每个轨道的字符位置
	fmt.Println("\n每个轨道的字符位置分析:")
	for rail := 0; rail < rails; rail++ {
		fmt.Printf("轨道%d的字符位置: ", rail)
		positions := []int{}
		for i := 0; i < len(message); i++ {
			if matrix[rail][i] != " " {
				positions = append(positions, i)
			}
		}
		fmt.Printf("%v\n", positions)

		// 分析位置间隔
		if len(positions) > 1 {
			fmt.Printf("  位置间隔: ")
			for i := 1; i < len(positions); i++ {
				fmt.Printf("%d ", positions[i]-positions[i-1])
			}
			fmt.Println()
		}
	}

	// 计算周期
	cycle := (rails - 1) * 2
	fmt.Printf("\n周期长度: %d (计算公式: (rails-1)*2)\n", cycle)

	// 分析每个轨道的步长模式
	fmt.Println("\n每个轨道的步长模式:")
	for rail := 0; rail < rails; rail++ {
		fmt.Printf("轨道%d: ", rail)
		if rail == 0 || rail == rails-1 {
			fmt.Printf("固定步长 %d\n", cycle)
		} else {
			fmt.Printf("交替步长: %d 和 %d\n", rail*2, cycle-rail*2)
		}
	}
}

// Swap算法实现（带调试信息）
func SwapDebug(s string, numRails int, a action) string {
	var i, index, delta int

	// moving pointers which track and response position and string position
	var rpos, spos *int
	if a == encode {
		rpos = &index
		spos = &i
		fmt.Printf("编码模式: rpos指向index, spos指向i\n")
	} else { // decoding is just the opposite of encoding, swap the pointers
		rpos = &i
		spos = &index
		fmt.Printf("解码模式: rpos指向i, spos指向index\n")
	}

	cycle := (numRails - 1) * 2
	fmt.Printf("周期长度: %d\n", cycle)
	resp := []byte(s)

	for rail := 0; rail < numRails; rail++ {
		fmt.Printf("\n处理轨道 %d:\n", rail)
		delta = rail * 2
		fmt.Printf("初始delta: %d\n", delta)

		for i = rail; i < len(s); i += delta {
			fmt.Printf("  i=%d, index=%d, delta=%d: 将s[%d]='%c' 复制到 resp[%d]\n",
				i, index, delta, *spos, s[*spos], *rpos)
			resp[*rpos] = s[*spos]
			index++
			if delta == cycle {
				fmt.Printf("    delta等于cycle，继续\n")
				continue
			}
			delta = cycle - delta
			fmt.Printf("    更新delta: %d\n", delta)
		}
	}

	fmt.Printf("最终结果: %s\n", string(resp))
	return string(resp)
}

// Swap算法实现
func Swap(s string, numRails int, a action) string {
	var i, index, delta int

	// moving pointers which track and response position and string position
	var rpos, spos *int
	if a == encode {
		rpos = &index
		spos = &i
	} else { // decoding is just the opposite of encoding, swap the pointers
		rpos = &i
		spos = &index
	}

	cycle := (numRails - 1) * 2
	resp := []byte(s)
	for rail := 0; rail < numRails; rail++ {
		delta = rail * 2
		for i = rail; i < len(s); i += delta {
			resp[*rpos] = s[*spos]
			index++
			if delta == cycle {
				continue
			}
			delta = cycle - delta
		}
	}
	return string(resp)
}

// Swap版本的Encode和Decode
func EncodeSwap(s string, numRails int) string {
	return Swap(s, numRails, encode)
}

func DecodeSwap(s string, numRails int) string {
	return Swap(s, numRails, decode)
}

// 可视化编码和解码的对称性
func visualizeEncodeDecodeSymmetry(message string, rails int) {
	fmt.Printf("\n=== 编码和解码对称性分析 ===\n")
	fmt.Printf("消息: %s\n", message)
	fmt.Printf("轨道数: %d\n", rails)

	// 编码过程
	fmt.Println("\n--- 编码过程 ---")
	encoded := Swap(message, rails, encode)
	fmt.Printf("编码结果: %s\n", encoded)

	// 解码过程
	fmt.Println("\n--- 解码过程 ---")
	decoded := Swap(encoded, rails, decode)
	fmt.Printf("解码结果: %s\n", decoded)

	// 分析对称性
	fmt.Println("\n--- 对称性分析 ---")
	fmt.Println("编码: 从原始消息读取，写入编码结果")
	fmt.Println("解码: 从编码结果读取，写入原始消息")
	fmt.Println("关键洞察: 编码和解码只是'源'和'目标'的交换！")

	// 可视化指针交换
	fmt.Println("\n--- 指针交换机制 ---")
	fmt.Println("编码模式:")
	fmt.Println("  rpos = &index  (指向结果位置)")
	fmt.Println("  spos = &i      (指向源位置)")
	fmt.Println("  操作: resp[rpos] = s[spos]")

	fmt.Println("\n解码模式:")
	fmt.Println("  rpos = &i      (指向结果位置)")
	fmt.Println("  spos = &index  (指向源位置)")
	fmt.Println("  操作: resp[rpos] = s[spos]")

	fmt.Println("\n注意: 只是交换了指针，算法逻辑完全相同！")
}

// 测试编码解码对称性
func TestEncodeDecodeSymmetry(t *testing.T) {
	testCases := []struct {
		name    string
		message string
		rails   int
	}{
		{"two_rails", "XOXOXOXOXOXOXOXOXO", 2},
		{"three_rails", "WEAREDISCOVEREDFLEEATONCE", 3},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			visualizeEncodeDecodeSymmetry(tc.message, tc.rails)
		})
	}
}

// 测试Swap算法的工作原理
func TestSwapAlgorithm(t *testing.T) {
	fmt.Println("=== Swap算法工作原理演示 ===")

	message := "XOXOXOXOXOXOXOXOXO"
	rails := 2

	fmt.Printf("消息: %s\n", message)
	fmt.Printf("轨道数: %d\n", rails)

	fmt.Println("\n--- 编码过程 ---")
	result := SwapDebug(message, rails, encode)

	fmt.Println("\n--- 解码过程 ---")
	decoded := SwapDebug(result, rails, decode)

	if decoded != message {
		t.Errorf("解码失败: 期望 %s, 得到 %s", message, decoded)
	}
}

// 测试数学规律分析
func TestZigzagPatternAnalysis(t *testing.T) {
	testCases := []struct {
		name    string
		message string
		rails   int
	}{
		{"two_rails", "XOXOXOXOXOXOXOXOXO", 2},
		{"three_rails", "WEAREDISCOVEREDFLEEATONCE", 3},
		{"four_rails", "EXERCISES", 4},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			visualizeZigzagPattern(tc.message, tc.rails)
		})
	}
}

// 测试两种算法的正确性
func TestAlgorithmComparison(t *testing.T) {
	testCases := []struct {
		name     string
		message  string
		rails    int
		expected string
	}{
		{
			name:     "two rails",
			message:  "XOXOXOXOXOXOXOXOXO",
			rails:    2,
			expected: "XXXXXXXXXOOOOOOOOO",
		},
		{
			name:     "three rails",
			message:  "WEAREDISCOVEREDFLEEATONCE",
			rails:    3,
			expected: "WECRLTEERDSOEEFEAOCAIVDEN",
		},
		{
			name:     "four rails",
			message:  "EXERCISES",
			rails:    4,
			expected: "ESXIEECSR",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// 测试原始算法
			result1 := Encode(tc.message, tc.rails)
			if result1 != tc.expected {
				t.Errorf("原始算法 Encode(%q, %d) = %q, want %q", tc.message, tc.rails, result1, tc.expected)
			}

			// 测试Swap算法
			result2 := EncodeSwap(tc.message, tc.rails)
			if result2 != tc.expected {
				t.Errorf("Swap算法 Encode(%q, %d) = %q, want %q", tc.message, tc.rails, result2, tc.expected)
			}

			// 测试解码
			decoded1 := Decode(result1, tc.rails)
			if decoded1 != tc.message {
				t.Errorf("原始算法 Decode(%q, %d) = %q, want %q", result1, tc.rails, decoded1, tc.message)
			}

			decoded2 := DecodeSwap(result2, tc.rails)
			if decoded2 != tc.message {
				t.Errorf("Swap算法 Decode(%q, %d) = %q, want %q", result2, tc.rails, decoded2, tc.message)
			}
		})
	}
}

// 性能对比基准测试
func BenchmarkEncodeOriginal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range encodeTests {
			Encode(test.message, test.rails)
		}
	}
}

func BenchmarkEncodeSwap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range encodeTests {
			EncodeSwap(test.message, test.rails)
		}
	}
}

func BenchmarkDecodeOriginal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range decodeTests {
			Decode(test.message, test.rails)
		}
	}
}

func BenchmarkDecodeSwap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range decodeTests {
			DecodeSwap(test.message, test.rails)
		}
	}
}

// 展示Swap算法的设计过程
func demonstrateSwapDesignProcess(message string, rails int) {
	fmt.Printf("\n=== Swap算法设计过程演示 ===\n")
	fmt.Printf("消息: %s\n", message)
	fmt.Printf("轨道数: %d\n", rails)

	// 步骤1：观察数学模式
	fmt.Println("\n步骤1: 观察数学模式")
	cycle := (rails - 1) * 2
	fmt.Printf("周期长度: %d = (%d-1)*2\n", cycle, rails)

	fmt.Println("每个轨道的步长模式:")
	for rail := 0; rail < rails; rail++ {
		if rail == 0 || rail == rails-1 {
			fmt.Printf("  轨道%d: 固定步长 %d\n", rail, cycle)
		} else {
			fmt.Printf("  轨道%d: 交替步长 %d 和 %d\n", rail, rail*2, cycle-rail*2)
		}
	}

	// 步骤2：发现对称性
	fmt.Println("\n步骤2: 发现编码解码对称性")
	fmt.Println("编码: 从原始消息读取 → 写入编码结果")
	fmt.Println("解码: 从编码结果读取 → 写入原始消息")
	fmt.Println("关键洞察: 只是'源'和'目标'的交换！")

	// 步骤3：设计指针交换机制
	fmt.Println("\n步骤3: 设计指针交换机制")
	fmt.Println("统一算法框架:")
	fmt.Println("  resp[rpos] = s[spos]")
	fmt.Println("  其中 rpos 和 spos 根据编码/解码模式设置")

	fmt.Println("\n编码模式:")
	fmt.Println("  rpos = &index  (指向结果位置)")
	fmt.Println("  spos = &i      (指向源位置)")

	fmt.Println("\n解码模式:")
	fmt.Println("  rpos = &i      (指向结果位置)")
	fmt.Println("  spos = &index  (指向源位置)")

	// 步骤4：实现步长计算
	fmt.Println("\n步骤4: 实现步长计算")
	fmt.Println("初始步长: delta = rail * 2")
	fmt.Println("交替步长: delta = cycle - delta")
	fmt.Println("固定步长: 当 delta == cycle 时保持不变")

	// 步骤5：完整算法
	fmt.Println("\n步骤5: 完整算法")
	fmt.Println("for rail := 0; rail < numRails; rail++ {")
	fmt.Println("  delta = rail * 2")
	fmt.Println("  for i = rail; i < len(s); i += delta {")
	fmt.Println("    resp[rpos] = s[spos]")
	fmt.Println("    if delta != cycle {")
	fmt.Println("      delta = cycle - delta")
	fmt.Println("    }")
	fmt.Println("  }")
	fmt.Println("}")
}

// 测试设计过程演示
func TestSwapDesignProcess(t *testing.T) {
	testCases := []struct {
		name    string
		message string
		rails   int
	}{
		{"two_rails", "XOXOXOXOXOXOXOXOXO", 2},
		{"three_rails", "WEAREDISCOVEREDFLEEATONCE", 3},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			demonstrateSwapDesignProcess(tc.message, tc.rails)
		})
	}
}
