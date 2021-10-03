package parser

import (
	"io"
	"strings"

	html "golang.org/x/net/html"
)

type Link struct {
	Href, Text string
}

func GetLinks(r io.Reader) ([]Link, error) {

	doc, err := html.Parse(r)
	if err != nil {
		return nil, err
	}

	var links []Link
	aList := getListOfLink(doc)
	for _, l := range aList {
		links = append(links, getLinkDetails(l))
	}

	return links, nil

}

func getListOfLink(parent *html.Node) []*html.Node {

	// type Node struct {
	// 	Parent, FirstChild, LastChild, PrevSibling, NextSibling *Node

	// 	Type      NodeType
	// 	DataAtom  atom.Atom
	// 	Data      string
	// 	Namespace string
	// 	Attr      []Attribute
	// }

	if parent.Type == html.ElementNode && parent.Data == "a" {
		return []*html.Node{parent}
	}

	var ret []*html.Node
	for c := parent.FirstChild; c != nil; c = c.NextSibling {
		ret = append(ret, getListOfLink(c)...)
	}
	return ret
}

func getLinkDetails(parent *html.Node) Link {
	var href string
	for _, a := range parent.Attr {
		if a.Key == "href" {
			href = a.Val
			break
		}
	}

	linkText := ""
	getLinkStrings(parent, &linkText)
	return Link{href, cleanUpString(linkText)}
}

func getLinkStrings(parent *html.Node, s *string) {
	for c := parent.FirstChild; c != nil; c = c.NextSibling {
		if c.Type == html.TextNode && c.Data != "" {
			*s = *s + c.Data
		}

		getLinkStrings(c, s)
	}
}

func cleanUpString(s string) string {
	return strings.Join(strings.Fields(s), " ")
}
