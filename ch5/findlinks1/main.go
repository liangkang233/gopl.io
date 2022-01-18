// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 122.
//!+main

// Findlinks1 prints the links in an HTML document read from standard input.
package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

var freq = make(map[string]int)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
	// 练习5.2： 编写函数，记录在HTML树中出现的同名元素的次数。 仿照outline
	// cat .\fetch\index.html |  go run .\findlinks1\
	tagFreq(doc)
	for tag, count := range freq {
		fmt.Printf("%4d %s\n", count, tag)
	}
}

//!-main

//!+visit
// visit appends to links each link found in n and returns the result.
func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	// 练习5.4： 扩展visit函数，使其能够处理其他类型的结点，如images、scripts和style sheets。
	if n.Type == html.ElementNode && (n.Data == "img" || n.Data == "script") {
		for _, a := range n.Attr {
			if a.Key == "src" {
				links = append(links, a.Val)
			}
		}
	}
	// for c := n.FirstChild; c != nil; c = c.NextSibling {
	// 	links = visit(links, c)
	// }

	// 练习5.1： 修改findlinks代码中遍历n.FirstChild链表的部分，将循环调用visit，改成递归调用。
	if n.FirstChild != nil {
		links = visit(links, n.FirstChild)
	}
	if n.NextSibling != nil {
		links = visit(links, n.NextSibling)
	}
	return links
}

//!-visit

func tagFreq(n *html.Node) {
	if n.Type == html.ElementNode {
		freq[n.Data]++
		// fmt.Printf("%4d %s\n", freq[n.Data], n.Data)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		tagFreq(c)
	}
}

// 练习5.3 https://github.com/torbiak/gopl/blob/master/ex5.3/main.go
// 网友直接使用 z := html.NewTokenizer(os.Stdin) 来做的解析tag
// z := html.NewTokenizer(r)  ,大部分值有偏差
// 5.2 例子: https://github.com/torbiak/gopl/blob/master/ex5.2/main.go
/* PS E:\Task\Study_Go\gopl.io\ch5>  cat .\fetch\index.html |  go run .\test.go
  15 input
   2 head
   8 link
   2 title
   2 body
   6 textarea
 105 div
  16 img
   4 ul
   2 b
  20 li
  18 p
   5 meta
  52 script
   2 map
   2 form
   2 html
  16 style
   2 noscript
 100 a
  12 i
  49 span
   1 area
PS E:\Task\Study_Go\gopl.io\ch5>  cat .\fetch\index.html |  go run .\findlinks1\
   6 i
  16 img
   1 area
   2 ul
   9 p
   1 html
   8 link
  28 span
  10 li
   1 head
   1 title
   3 textarea
  53 div
   1 form
  15 input
   5 meta
  26 script
   1 body
  56 a
   1 map
   1 b
   8 style
   1 noscript */

/*
//!+html
package html

type Node struct {
	Type                    NodeType
	Data                    string
	Attr                    []Attribute
	FirstChild, NextSibling *Node
}

type NodeType int32

const (
	ErrorNode NodeType = iota
	TextNode
	DocumentNode
	ElementNode
	CommentNode
	DoctypeNode
)

type Attribute struct {
	Key, Val string
}

func Parse(r io.Reader) (*Node, error)
//!-html
*/
