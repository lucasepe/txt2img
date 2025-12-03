package cmd

import (
	"flag"
	"image/color"

	"github.com/lucasepe/t2i/internal/image/text"
	"github.com/lucasepe/x/text/conv"
)

type Options struct {
	text.RenderOptions
	TabSize int
	Outfile string
}

func Configure(fs *flag.FlagSet, vals *FlagValues, args []string) Options {
	// Parse command-line args
	if err := fs.Parse(args); err != nil {
		return Options{}
	}

	// Parse colors
	fr, fg, fb, fa := conv.RGBA(*vals.TextColor)
	br, bg, bb, ba := conv.RGBA(*vals.BgColor)

	// Build the final render options
	return Options{
		RenderOptions: text.RenderOptions{
			Margin:          *vals.Margin,
			AutoSize:        true,
			Square:          *vals.Square,
			TextColor:       color.RGBA{R: fr, G: fg, B: fb, A: fa},
			BackgroundColor: color.RGBA{R: br, G: bg, B: bb, A: ba},
			LineSpacing:     *vals.LineSpacing,
		},
		TabSize: *vals.TabSize,
		Outfile: *vals.Outfile,
	}
}
