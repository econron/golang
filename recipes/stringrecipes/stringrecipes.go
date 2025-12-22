package stringrecipes

import (
	"fmt"
	"strings"
	"strconv"
	"unicode"
)
func Basic1() {
	s := "こんにちは"
	runes := []rune(s)
	fmt.Printf("runes: %v\n", runes)
	fmt.Printf("value: %v\n", runes[0]) // 12371
	fmt.Printf("value: %v\n", string(runes[0])) // こ
}

func Basic2() {
	s := "こんにちは"

	// 文字列を含んでいるか
	ret := strings.Contains(s, "こん")
	fmt.Printf("ret: %v\n", ret)

	// 文字列を分割する
	parts := strings.Split(s, "に")
	fmt.Printf("parts: %v\n", parts)

	// 文字列のスライスを結合する
	// こん、ちは　の間を「に」で接着する
	joined := strings.Join(parts, "に")
	fmt.Printf("joined: %v\n", joined)

	// 置換
	replaced := strings.Replace(s, "こん", "さく", 1)
	fmt.Printf("replaced: %v\n", replaced)

	s2 := "myakumyakumyaku"
	// -1はすべてのmyaをwaに置換する
	// replaceAllというメソッドもある
	replaced2 := strings.Replace(s2, "mya", "wa", -1)
	fmt.Printf("replaced2: %v\n", replaced2)
}

// strings.Builderを使うと、文字列を効率的に結合できる
func Basic3() {
	var builder strings.Builder
	for i := 0; i < 1000; i++ {
		builder.WriteString("a")
	}
	result := builder.String()
	fmt.Printf("result: %v\n", result)
}

// 文字列の変換
func Basic4() {
	i, _ := strconv.Atoi("100")
	fmt.Printf("i: %v\n", i)
	
	s := strconv.Itoa(100)
	fmt.Printf("s: %v\n", s)
}

// ランレングス符号
func RLE(s string) string {
	if len(s) == 0 {
		return ""
	}
	runes := []rune(s)
	var builder strings.Builder
	i := 0
	for i < len(runes) {
		char := runes[i]
		i++

		digitStart := i
		// もし数値ならばカーソルを1つ前に進めることで、スライスの切り出し処理に利用できる
		// 例えばa4b5とあってa->4と見つかったらbにカーソルを合わせると、
		// countStr := string(runes[digitStart:i])で4を取り出せる
		for i < len(runes) && unicode.IsDigit(runes[i]) {
			i++
		}
		countStr := string(runes[digitStart:i])
		// 文字列を数値に変換
		count, _ := strconv.Atoi(countStr)
		for j := 0; j < count; j++ {
			// 文字をcount回繰り返してbuilderに追加
			builder.WriteRune(char)
		}
	}
	return builder.String()
}

