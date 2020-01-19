package test

import (
	"fmt"
	"golang.org/x/image/draw"
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
	"testing"
)


const (
	maxEsc = 100
	rMin   = -2.
	rMax   = .5
	iMin   = -1.
	iMax   = 1.
	width  = 750
	red    = 230
	green  = 235
	blue   = 255
)

func mandelbrot(a complex128) float64 {
	i := 0
	for z := a; cmplx.Abs(z) < 2 && i < maxEsc; i++ {
		z = z*z + a
	}
	return float64(maxEsc-i) / maxEsc
}
func TestTestfa(t *testing.T) {
	//MainLogger := logger.NewLogger("logs/test.log", zapcore.DebugLevel, 1, 2, 7, true, "Main").Sugar()
	//mainLogger := logger.NewLogger("logs/test.log", zapcore.DebugLevel, 1, 3, 7, true, "Main").Sugar()
	//for {
	//	mainLogger.Info("PanicError==》》--" + fmt.Sprint("aaaaaaaaaaaaaaaaaaaa"))
	//	mainLogger.Debug("i am debug",zap.String("key","debug"))
	//	mainLogger.Debug("aaaaaaa","12312312")
	//}

	scale := width / (rMax - rMin)
	height := int(scale * (iMax - iMin))
	bounds := image.Rect(0, 0, width, height)
	b := image.NewNRGBA(bounds)
	draw.Draw(b, bounds, image.NewUniform(color.Black), image.ZP, draw.Src)
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			fEsc := mandelbrot(complex(
				float64(x)/scale+rMin,
				float64(y)/scale+iMin))
			b.Set(x, y, color.NRGBA{uint8(red * fEsc),
				uint8(green * fEsc), uint8(blue * fEsc), 255})

		}
	}
	f, err := os.Create("mandelbrot.png")
	if err != nil {
		fmt.Println(err)
		return
	}
	if err = png.Encode(f, b); err != nil {
		fmt.Println(err)
	}
	if err = f.Close(); err != nil {
		fmt.Println(err)
	}
}
