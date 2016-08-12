package main

import (
	"github.com/cloudson/pizzeria/graph"

	"io"
)

type Graph interface {
	Height() int
	Width() int
	Values() []*graph.Value
}

type GraphService interface {
	Build(g *Graph, w io.Writer) string
}
