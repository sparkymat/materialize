package materialize

import (
	"github.com/sparkymat/html"
	"github.com/sparkymat/html/viewport"
)

const ViewportDefaultScale = 1.0

type Template struct {
	title          string
	bodyNode       *html.BodyNode
	extraHeadNodes []html.HeadNode
	styleNodes     []html.HeadNode
	scriptNodes    []html.BodyNode
}

func NewTemplate(title string) Template {
	return Template{
		title: title,
	}
}

func (t Template) CSS(styleLinks []string) Template {
	copiedTemplate := t

	for _, styleLink := range styleLinks {
		copiedTemplate.styleNodes = append(copiedTemplate.styleNodes, html.Link(styleLink))
	}

	return copiedTemplate
}

func (t Template) ExtraHeadNodes(nodes []html.HeadNode) Template {
	copiedTemplate := t
	copiedTemplate.extraHeadNodes = nodes

	return copiedTemplate
}

func (t Template) Body(bodyNode html.BodyNode) Template {
	copiedTemplate := t
	copiedTemplate.bodyNode = &bodyNode

	return copiedTemplate
}

func (t Template) HTML() html.HTMLNode {
	headNodes := []html.HeadNode{
		html.MetaViewport(viewport.WidthDeviceWidth(), viewport.InitialScale(ViewportDefaultScale)),
		html.Title(t.title),
	}
	headNodes = append(headNodes, t.styleNodes...)
	headNodes = append(headNodes, t.extraHeadNodes...)

	bodyNodes := []html.BodyNode{}

	if t.bodyNode != nil {
		bodyNodes = append(bodyNodes, *t.bodyNode)
	}

	bodyNodes = append(bodyNodes, t.scriptNodes...)

	return html.HTML(
		html.Head(headNodes...),
		html.Body(bodyNodes...),
	)
}
