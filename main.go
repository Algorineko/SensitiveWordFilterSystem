package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// 前缀树匹配敏感词
func conductMatch(sensitiveWords []string, matchContents []string) {

	trie := NewSensitiveTrie()
	trie.AddWords(sensitiveWords)

	for _, srcText := range matchContents {
		isSensitive, sensitiveWords := trie.Match(srcText)
		fmt.Println("srcText        -> ", srcText)
		fmt.Println("isSensitive    -> ", isSensitive)
		fmt.Println("sensitiveWords -> ", sensitiveWords)
		fmt.Println()
	}
}

// 从文件读取敏感词
func readSensitiveWordsFromFile(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("无法打开敏感词文件: %v", err)
	}
	defer file.Close()

	var sensitiveWords []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		word := strings.TrimSpace(scanner.Text())
		if word != "" { // 跳过空行
			sensitiveWords = append(sensitiveWords, word)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("读取敏感词文件时出错: %v", err)
	}

	return sensitiveWords, nil
}

func main() {
	// 从文件读取敏感词
	filename := "sensitive_dict.txt"
	sensitiveWords, err := readSensitiveWordsFromFile(filename)
	if err != nil {
		fmt.Printf("错误: %v\n", err)
		return
	} else {
		fmt.Printf("成功从 %s 读取敏感词: %v\n", filename, sensitiveWords)
	}

	matchContents := []string{
		"内#部资料",
		"猫哈*气",
		"狗龇@&牙",
		"正常文本",
	}

	fmt.Println("\n--------- 前缀树匹配敏感词 ---------")
	conductMatch(sensitiveWords, matchContents)

}
