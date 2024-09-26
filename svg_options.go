package graphify

import "github.com/vibridi/graphify/internal/svg"

var defaultSvgOptions = svg.Options{
	CanvasPadding:          60,
	ShowTimestamp:          false,
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

func WithShowTimestamp() SvgOption {
	return func(o *svg.Options) {
		o.ShowTimestamp = true
	}
}

// Currently this option does nothing because autog's output layout doesn't distinguish virtual nodes from regular nodes.
func WithSkipVirtualNodes() SvgOption {
	return func(o *svg.Options) {
		o.DrawVirtualNodes = false
	}
}

func WithDrawSplines() SvgOption {
	return func(o *svg.Options) {
		o.DrawSplines = true
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
