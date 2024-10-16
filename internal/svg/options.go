package svg

type Options struct {
	CanvasPadding          int
	ShowTimestamp          bool
	DrawVirtualNodes       bool
	DrawSplines            bool
	PrintNodePosition      bool
	HighlightReversedEdges bool
	FontSize               int16
	Elements               []Stringer
}
