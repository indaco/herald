// Using herald's Figure with image2ascii to render images in the terminal.
//
// This example is a separate Go module with its own go.mod to keep
// github.com/qeesung/image2ascii out of herald's core dependencies.
//
// Run:
//
//	cd examples/208_figure-with-image && go run . [path-to-image.png]
//
// If no image path is provided, it renders the bundled sample.png (Go logo).
package main

import (
	"fmt"
	"image"
	_ "image/png"
	"os"
	"path/filepath"
	"runtime"

	"github.com/indaco/herald"
	"github.com/qeesung/image2ascii/convert"
)

func main() {
	ty := herald.New()

	path := imgPath()
	f, err := os.Open(path)
	if err != nil {
		os.Stderr.WriteString("error: " + err.Error() + "\n")
		os.Exit(1)
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		os.Stderr.WriteString("error decoding image: " + err.Error() + "\n")
		os.Exit(1)
	}

	converter := convert.NewImageConverter()
	opts := convert.DefaultOptions
	opts.FixedWidth = 60
	opts.FixedHeight = 20
	ascii := converter.Image2ASCIIString(img, &opts)

	fmt.Println(ty.Compose(
		ty.H1("Figure + Image Demo"),
		ty.P("Rendering an image as ASCII art inside a herald Figure."),
		ty.Figure(ascii, "Figure 1: "+filepath.Base(path)),
		ty.HR(),
		ty.FigureTop(
			ty.Fieldset("Image Info", ty.KVGroup([][2]string{
				{"Source", filepath.Base(path)},
				{"Width", fmt.Sprintf("%d px", img.Bounds().Dx())},
				{"Height", fmt.Sprintf("%d px", img.Bounds().Dy())},
				{"Format", "ASCII art via image2ascii"},
			})),
			"Figure 2: Image metadata",
		),
	))
}

// imgPath returns the image file to render: the first CLI argument,
// or the bundled sample.png next to this source file.
func imgPath() string {
	if len(os.Args) > 1 {
		return os.Args[1]
	}
	_, src, _, _ := runtime.Caller(0)
	return filepath.Join(filepath.Dir(src), "sample.png")
}
