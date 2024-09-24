package graphify

import (
	"io"

	"github.com/nulab/autog/graph"
	"github.com/vibridi/graphify/internal/svg"
)

func DrawSVG(layout graph.Layout, w io.Writer, options ...SvgOption) {
	opts := defaultSvgOptions
	for _, o := range options {
		o(&opts)
	}
	svg.Draw(layout, w, opts)
}
