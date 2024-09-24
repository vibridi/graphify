# graphify

This repository contains visualization tools for [`github.com/nulab/autog`](https://github.com/nulab/autog).

## SVG

Draws autog's `graph.Layout` to an SVG file. Example usage:

```go
package main

import (
	"os"

	"github.com/nulab/autog"
	"github.com/vibridi/graphify"
)

func main() {
	layout := autog.Layout(/* args */)

	f, err := os.Create("output.svg")
	if err != nil {
		// handle error
    }
	graphify.DrawSVG(layout, f)
}
```
