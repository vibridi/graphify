package svg

import (
	"io"
	"time"

	svgo "github.com/ajstarks/svgo"
	"github.com/nulab/autog/graph"
)

func Draw(g graph.Layout, w io.Writer, opts Options) {
	canvas := svgo.New(w)

	var left, right, top, bottom float64
	for _, n := range g.Nodes {
		left = min(left, n.X)
		right = max(right, n.X+n.W)
		top = min(top, n.Y)
		bottom = max(bottom, n.Y+n.H)
	}
	padding := opts.CanvasPadding

	canvas.Start(int(right-left+float64(padding)*2), int(bottom-top+float64(padding)))
	if opts.ShowTimestamp {
		canvas.Text(10, 10+padding, time.Now().String(), "font-size:20px;fill:black")
	}
	canvas.Def()
	canvas.Marker("arrowhead", 7, 2, 8, 8, `orient="auto-start-reverse"`)
	canvas.Path("M0,0 V4 L8,2 Z", "fill:black")
	canvas.MarkerEnd()
	canvas.DefEnd()

	for _, n := range g.Nodes {
		canvas.Rect(int(n.X)+padding, int(n.Y)+padding*2, int(n.W), int(n.H), "fill:white;stroke:black")

		text := n.ID
		if opts.PrintNodePosition {
			// text += "-" + strconv.Itoa(n.LayerPos)
		}
		canvas.Text(int(n.X)+int(n.W)/2+padding, int(n.Y)+int(n.H)/2+padding*2, text, "text-anchor:middle;font-size:30px;fill:black")
	}

	for _, e := range g.Edges {
		// if splines {
		// 	splineEdges(canvas, e, padding)
		// } else {
		// 	regularEdges(canvas, e, padding)
		// }
		regularEdges(canvas, e, opts)
	}
	canvas.End()
}

func regularEdges(canvas *svgo.SVG, e graph.Edge, opts Options) {
	if len(e.Points) == 0 {
		return
	}
	var xs, ys []int
	for _, p := range e.Points {
		xs = append(xs, int(p[0])+opts.CanvasPadding)
		ys = append(ys, int(p[1])+opts.CanvasPadding*2)
	}

	marker := "marker-end"
	stroke := "black"
	if e.ArrowHeadStart {
		if opts.HighlightReversedEdges {
			stroke = "red"
		}
		marker = "marker-start"
	}
	canvas.Polyline(xs, ys, "stroke-width:2;fill:none;stroke:"+stroke+";"+marker+":url(#arrowhead)")
}
