package main

import (
	"fmt"
	"os"

	"github.com/lucasepe/txt2img/internal/cmd"
	"github.com/lucasepe/txt2img/internal/image/text"
	cmdutil "github.com/lucasepe/txt2img/internal/util/cmd"
	ioutil "github.com/lucasepe/txt2img/internal/util/io"
	imageio "github.com/lucasepe/x/image/io"
	textutil "github.com/lucasepe/x/text"
)

func main() {
	fs, fv := cmd.NewFlagSet()

	// Custom usage
	fs.Usage = cmd.Usage(fs)

	opts := cmd.Configure(fs, fv, os.Args[1:])

	// If help flag was supplied: print help and exit
	if *fv.ShowHelp {
		fs.Usage()
		return
	}

	input, err := ioutil.ReadInput(fs.Args())
	cmdutil.CheckErr("error reading input", err)
	input = textutil.Clean(input, opts.TabSize)

	img, err := text.RenderGG(string(input), opts.RenderOptions)
	cmdutil.CheckErr("rendering text to image failed", err)

	f, err := os.Create(opts.Outfile)
	cmdutil.CheckErr("cannot write output", err)
	defer f.Close()

	err = imageio.WriteToFile(img.Image(), opts.Outfile, imageio.PNG)
	cmdutil.CheckErr("failed to write image", err)

	fmt.Println("Image written to", opts.Outfile)
}
