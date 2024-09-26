package svg

import (
	"fmt"
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

	canvas.Start(int(right-left+float64(padding)*2), int(bottom-top+float64(padding)*2))
	if opts.ShowTimestamp {
		canvas.Text(10, 10+padding, time.Now().String(), "font-size:20px;fill:black")
	}
	canvas.Def()
	canvas.Marker("arrowhead", 7, 2, 8, 8, `orient="auto-start-reverse"`)
	canvas.Path("M0,0 V4 L8,2 Z", "fill:black")
	canvas.MarkerEnd()
	canvas.DefEnd()

	for _, n := range g.Nodes {
		canvas.Rect(int(n.X)+padding, int(n.Y)+padding, int(n.W), int(n.H), "fill:white;stroke:black")

		text := n.ID
		if opts.PrintNodePosition {
			// text += "-" + strconv.Itoa(n.LayerPos)
		}
		canvas.Text(int(n.X)+int(n.W)/2+padding, int(n.Y)+int(n.H)/2+padding, text, "text-anchor:middle;font-size:30px;fill:black")
	}

	for _, e := range g.Edges {
		if opts.DrawSplines {
			drawCubicBezier(canvas, e, opts)
		} else {
			drawPolyline(canvas, e, opts)
		}
	}
	canvas.End()
}

func drawPolyline(canvas *svgo.SVG, e graph.Edge, opts Options) {
	if len(e.Points) == 0 {
		return
	}
	var xs, ys []int
	for _, p := range e.Points {
		xs = append(xs, int(p[0])+opts.CanvasPadding)
		ys = append(ys, int(p[1])+opts.CanvasPadding)
	}
	canvas.Polyline(xs, ys, lineStyle(lineParams(e, opts)))
}

func drawCubicBezier(canvas *svgo.SVG, e graph.Edge, opts Options) {
	if len(e.Points) == 0 {
		return
	}
	for i := 0; i < len(e.Points); i += 4 {
		p1, p2, p3, p4 := e.Points[i], e.Points[i+1], e.Points[i+2], e.Points[i+3]
		d := fmt.Sprintf("M%.2f,%.2f C%.2f,%.2f %.2f,%.2f %.2f,%.2f", p1[0], p1[1], p2[0], p2[1], p3[0], p3[1], p4[0], p4[1])
		canvas.Path(d, lineStyle(lineParams(e, opts)))
	}
}

func lineParams(e graph.Edge, opts Options) (stroke, marker string) {
	marker = "marker-end"
	stroke = "black"
	if e.ArrowHeadStart {
		if opts.HighlightReversedEdges {
			stroke = "red"
		}
		marker = "marker-start"
	}
	return
}

func lineStyle(stroke, marker string) string {
	return "stroke-width:2;fill:none;stroke:" + stroke + ";" + marker + ":url(#arrowhead)"
}
