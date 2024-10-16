package graphify

import "github.com/vibridi/graphify/internal/svg"

type SvgStringer = svg.Stringer

var defaultSvgOptions = svg.Options{
	CanvasPadding:          60,
	ShowTimestamp:          false,
	DrawVirtualNodes:       true,
	PrintNodePosition:      false,
	HighlightReversedEdges: false,
	FontSize:               25, // pixels
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

func WithSVGElements(elements []SvgStringer) SvgOption {
	return func(o *svg.Options) {
		o.Elements = elements
	}
}

func WithFontSize(size int16) SvgOption {
	return func(o *svg.Options) {
		o.FontSize = size
	}
}
