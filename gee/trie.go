package gee

import (
	"fmt"
	"strings"
)

/*
	pattern    待匹配路由  例如 /p/:lang/doc
	part       路由中的一部分 例如:lang
	children   子节点 例如[doc]
	isWild     是否精确匹配, part中含有 : * 时为true
*/

type node struct {
	pattern  string
	part     string
	children []*node
	isWild   bool
}

func (n *node) matchChild(part string) *node {
	/*
		range return 索引 & 元素
		_ 就是不需要这个索引，空出来了
	*/
	for _, child := range n.children {
		fmt.Println(child)
		if child.part == part || child.isWild {
			return child
		}
	}
	return nil
}

func (n *node) matchChildren(part string) []*node {
	nodes := make([]*node, 0)
	for _, child := range n.children {
		if child.part == part || child.isWild {
			nodes = append(nodes, child)
		}
	}
	return nodes
}

/*
	RUL 中的 : 表示后面的是一个变量名，需要一个具体数值
	eg：/lalala/:lang/doc
	意思是，lang是个变量名，具体值可能不同，可以匹配
	/lalala/c/doc
	/lalala/go/doc

	RUL 中的 * 表示后面的是一个通配内容,只有头部相同，后面跟什么都行
	eg：/lalala/*filenema
	可以匹配
	/lalala/hehe.txt
	/lalala/go/nono.go

	这里正是动态路由功能的需求原因
*/
func (n *node) insert(pattern string, parts []string, height int) {
	if len(parts) == height {
		n.pattern = pattern
		return
	}

	part := parts[height]
	child := n.matchChild(part)
	if child == nil {
		child = &node{part: part, isWild: part[0] == ':' || part[0] == '*'}
		n.children = append(n.children, child)
	}
	child.insert(pattern, parts, height+1)
}

func (n *node) search(parts []string, height int) *node {
	if len(parts) == height || strings.HasPrefix(n.part, "*") {
		if n.pattern == "" {
			return nil
		}
		return n
	}

	part := parts[height]
	children := n.matchChildren(part)

	for _, child := range children {
		result := child.search(parts, height+1)
		if result != nil {
			return result
		}
	}

	return nil
}
