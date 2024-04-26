package ahoCorasick

import (
	"unicode/utf8"
)

type trieNode struct {
	childNode  map[rune]*trieNode
	fairNode   *trieNode
	outputList []string
}

type AhoCorasick struct {
	root *trieNode
}

type word struct {
	start int
	end   int
}

func NewAhoCorasick() *AhoCorasick {
	return &AhoCorasick{root: &trieNode{childNode: make(map[rune]*trieNode)}}
}

func (c *AhoCorasick) insert(pattern string) {
	node := c.root
	for _, char := range pattern {
		if node.childNode[char] == nil {
			node.childNode[char] = &trieNode{childNode: make(map[rune]*trieNode)}
		}
		node = node.childNode[char]
	}
	node.outputList = append(node.outputList, pattern)
}

func (c *AhoCorasick) buildFailurePointers() {
	queue := make([]*trieNode, 0)
	for _, node := range c.root.childNode {
		node.fairNode = c.root
		queue = append(queue, node)
	}
	for len(queue) > 0 {
		currentNode := queue[0]
		queue = queue[1:]
		for char, node := range currentNode.childNode {
			failNode := currentNode.fairNode
			for failNode != nil && failNode.childNode[char] == nil {
				failNode = failNode.fairNode
			}
			if failNode == nil {
				node.fairNode = c.root
			} else {
				node.fairNode = failNode
			}
			queue = append(queue, node)
			node.outputList = append(node.outputList, node.fairNode.outputList...)
		}
	}
}

func (c *AhoCorasick) Match(text []rune) []*word {
	result := make([]*word, 0)
	node := c.root
	for i, char := range text {
		for node.childNode[char] == nil && node != c.root {
			node = node.fairNode
			c.collectMatches(node, i, &result)
		}
		if node.childNode[char] != nil {
			node = node.childNode[char]
		}
		c.collectMatches(node, i, &result)
	}
	return result
}

// 收集匹配结果
func (c *AhoCorasick) collectMatches(node *trieNode, index int, result *[]*word) {
	for _, pattern := range node.outputList {
		*result = append(*result, &word{start: index - utf8.RuneCountInString(pattern) + 1, end: index})
	}
}

func (c *AhoCorasick) Run(patterns []string) {
	for _, pattern := range patterns {
		c.insert(pattern)
	}
	c.buildFailurePointers()
}

func (c *AhoCorasick) MatchAndRewrite(strRune []rune) (string, bool) {
	replaceWords := c.Match(strRune)
	for _, w := range replaceWords {
		for i := w.start; i <= w.end; i++ {
			strRune[i] = '*'
		}
	}
	return string(strRune), len(replaceWords) > 0
}
