package solution

import (
	"strings"

	"golang.org/x/net/html"
)

type Link struct {
	Href string
	Text string
}

func getLinkNodeText(node *html.Node, href string) string {
	if node == nil || (node.Type != html.TextNode && node.Type != html.ElementNode) {
		return ""
	}
	if node.Type == html.TextNode {
		return node.Data
	}

	var textValue string
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		textValue += getLinkNodeText(c, href)
	}
	return textValue
}

func extractLinkData(node *html.Node) Link {
	attributes := node.Attr
	var linkValue string
	for _, attr := range attributes {
		if attr.Key == "href" {
			linkValue = attr.Val
			break
		}
	}

	textValue := getLinkNodeText(node, linkValue)
	textSlice := strings.Fields(textValue)
	textValue = strings.Join(textSlice, " ")

	return Link{
		linkValue,
		textValue,
	}
}

func GetLinks(node *html.Node) []Link {
	var myLinks []Link
	if node == nil {
		return myLinks
	}

	// link node
	if node.Data == "a" {
		myLinks = append(myLinks, extractLinkData(node))
		return myLinks
	}

	/*
		TODO:
			here in this loop what i made a mistake is that i wrote it like below
			for c := node.FirstChild; c != nil; c = node.NextSibling{}
			and i was thinking that it also should work but it doesn't work.
			Let's find out why it was not working in recursion.
	*/
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		myLinks = append(myLinks, GetLinks(c)...)
	}

	return myLinks
}
