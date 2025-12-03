package cmd

import (
	"flag"
	"os"
)

type FlagValues struct {
	Margin      *int
	TabSize     *int
	Square      *bool
	TextColor   *string
	BgColor     *string
	LineSpacing *float64
	Outfile     *string
	ShowHelp    *bool
}

// NewFlagSet creates a FlagSet with all supported CLI options.
func NewFlagSet() (*flag.FlagSet, *FlagValues) {
	fs := flag.NewFlagSet("txt2img", flag.ContinueOnError)
	fs.SetOutput(os.Stdout)

	vals := &FlagValues{
		Margin:      fs.Int("m", 24, "Margin around the text (pixels)"),
		TabSize:     fs.Int("t", 4, "Number of spaces to replace each tab"),
		Square:      fs.Bool("s", false, "Force the image to be square"),
		TextColor:   fs.String("c", "#000", "Text color in HEX format"),
		BgColor:     fs.String("b", "#fff", "Background color in HEX format"),
		LineSpacing: fs.Float64("l", 1.1, "Space between lines."),
		Outfile:     fs.String("o", "out.png", "Output PNG file path"),
		ShowHelp:    fs.Bool("h", false, "Show help"),
	}

	return fs, vals
}
