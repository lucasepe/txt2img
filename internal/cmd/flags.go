package cmd

import (
	"flag"
	"os"
)

type FlagValues struct {
	Margin        *int
	TabSize       *int
	TextColor     *string
	BgColor       *string
	LineSpacing   *float64
	Outfile       *string
	DebugBaseline *bool
	DebugGrid     *bool
	ShowHelp      *bool
}

// NewFlagSet creates a FlagSet with all supported CLI options.
func NewFlagSet() (*flag.FlagSet, *FlagValues) {
	fs := flag.NewFlagSet("txt2img", flag.ContinueOnError)
	fs.SetOutput(os.Stdout)

	vals := &FlagValues{
		Margin:        fs.Int("m", 24, "Margin around the text (pixels)"),
		TabSize:       fs.Int("t", 4, "Number of spaces to replace each tab"),
		TextColor:     fs.String("c", "#000", "Text color in HEX format"),
		BgColor:       fs.String("b", "#fff", "Background color in HEX format"),
		LineSpacing:   fs.Float64("l", 1.0, "Space between lines."),
		Outfile:       fs.String("o", "out.png", "Output PNG file path"),
		DebugBaseline: fs.Bool("B", false, "Draws the text baseline for each row"),
		DebugGrid:     fs.Bool("G", false, "Draws the layout grid (cell boundaries)"),
		ShowHelp:      fs.Bool("h", false, "Show help"),
	}

	return fs, vals
}
