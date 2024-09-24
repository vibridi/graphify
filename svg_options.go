package graphify

import "github.com/vibridi/graphify/internal/svg"

var defaultSvgOptions = svg.Options{
	CanvasPadding:          60,
	DrawVirtualNodes:       true,
	PrintNodePosition:      false,
	HighlightReversedEdges: false,
}

type SvgOption func(*svg.Options)

func WithCanvasPadding(padding int) SvgOption {
	return func(o *svg.Options) {
		o.CanvasPadding = padding
	}
}

// Currently this option does nothing because autog's output layout doesn't distinguish virtual nodes from regular nodes.
func WithSkipVirtualNodes() SvgOption {
	return func(o *svg.Options) {
		o.DrawVirtualNodes = false
	}
}

func WithPrintNodePosition() SvgOption {
	return func(o *svg.Options) {
		o.PrintNodePosition = true
	}
}

func WithHighlightReversedEdges() SvgOption {
	return func(o *svg.Options) {
		o.HighlightReversedEdges = true
	}
}
