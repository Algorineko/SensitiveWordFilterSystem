# 敏感词过滤系统 / Sensitive Word Filter System

这是一个基于字典树（Trie）的高性能敏感词过滤系统，支持中文、英文、数字及混合内容的敏感词检测和替换。

This is a high-performance sensitive word filter system based on Trie tree, supporting Chinese, English, numbers, and mixed content detection and replacement.

#### 📁 项目结构 / Project Structure


sensitive_text/
├── main.go                 # 主程序入口 / Main program entry
├── trie_tree.go            # 字典树核心实现 / Trie tree core implementation
├── sensitive_dict.txt      # 敏感词词典文件 / Sensitive word dictionary file
└── README.md              # 项目说明文档 / Project documentation

#### 🚀 快速开始 / Quick Start

1. 安装Go环境 / Install Go Environment

确保已安装Go 1.16+版本 / Ensure Go 1.16+ is installed
```sh
go version
```


2. 运行程序 / Run the Program

```go
// 进入项目目录 / Navigate to project directory
cd sensitive_text

// 运行程序 / Run the program
go run .
```

#### 📋 使用说明 / Usage Instructions

1. 配置敏感词库 / Configure Sensitive Words

编辑 sensitive_dict.txt 文件，每行一个敏感词：
Edit the sensitive_dict.txt file, one sensitive word per line:

```txt
内部
哈气
龇牙
```


2. 修改检测内容 / Modify Test Content

在 main.go 文件的 matchContents 数组中添加要检测的文本：
Add test text in the matchContents array in main.go:

```go
matchContents := []string{
    "内#部资料",
    "猫哈*气",
    "狗龇@&牙",
    "正常文本",
}
```


3. 运行检测 / Run Detection

```go
go run .
```


4. 查看结果 / View Results

程序会输出每条文本的检测结果：
The program will output detection results for each text:

```sh
成功从 sensitive_dict.txt 读取敏感词: [内部 哈气 龇牙]

--------- 前缀树匹配敏感词 ---------
srcText        ->  内#部资料
isSensitive    ->  true
sensitiveWords ->  [内部]

srcText        ->  猫哈*气
isSensitive    ->  true
sensitiveWords ->  [哈气]

srcText        ->  狗龇@&牙
isSensitive    ->  true
sensitiveWords ->  [龇牙]

srcText        ->  正常文本
isSensitive    ->  false
sensitiveWords ->  []
```

#### 🧩 代码示例 / Code Examples

```go
// 从文件读取敏感词 / Read sensitive words from file
sensitiveWords, err := readSensitiveWordsFromFile("sensitive_dict.txt")

// 创建字典树实例 / Create Trie instance
trie := NewSensitiveTrie()

// 添加敏感词 / Add sensitive words
trie.AddWords(sensitiveWords)

// 检测文本 / Detect text
isSensitive, foundWords := trie.Match("测试文本包含敏感词")
```

#### 🔧 算法原理 / Algorithm Principles

字典树结构 / Trie Structure

├── 傻 (傻)
│   ├── 宝 (傻宝) [END]
│   └── 叉 (傻叉) [END]
├── 杂
│   └── 鱼 (杂鱼) [END]
└── f
    └── u
        └── c
            └── k (fuck) [END]


匹配流程 / Matching Process

1. 预处理文本：过滤特殊字符，转为小写
2. 逐字符匹配：从每个字符开始尝试匹配
3. 最长匹配：寻找最长的敏感词匹配
4. 结果收集：记录所有匹配到的敏感词
5. 敏感词替换：用指定字符替换敏感词部分

#### 📊 性能特点 / Performance Characteristics

时间复杂度 / Time Complexity

• 构建字典树：O(n×m)，n为敏感词数量，m为平均长度

• 文本匹配：O(k)，k为文本长度

空间复杂度 / Space Complexity

• 字典树存储：优化共享前缀，节省空间

• 内存占用：与敏感词数量和长度成正比

#### 🎯 应用场景 / Application Scenarios

• 🔒 内容审核：论坛、评论区敏感词过滤

• 💬 即时通讯：聊天内容安全检测

• 📝 文档处理：批量文档内容审查

• 🎮 游戏聊天：游戏内聊天内容监控

• 📱 社交媒体：用户生成内容审核

🐛 故障排除 / Troubleshooting



_最后更新 / Last Updated: 2026-01-14_