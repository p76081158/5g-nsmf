package slicebinpack

import (
	"time"
	"math/rand"
	"image"
	"image/color"
	"image/draw"
    "image/png"
	"os"
)

func DrawBinPackResult(drawinfos []DrawBlock, width int, height int, scale_ratio int) {
	img := image.NewRGBA(image.Rect(0, height / scale_ratio * (-1), width / scale_ratio, 0))
	new_png_file := "logs/binpack/result.png"
	for i :=0; i< len(drawinfos); i++ {
		t := time.Now().UnixNano()
		r := rand.New(rand.NewSource(t))
		color_r := uint8(r.Intn(255))
		color_g := uint8(r.Intn(255))
		color_b := uint8(r.Intn(255))
		color   := color.RGBA{color_r, color_g, color_b, 255}
		topleftx  := drawinfos[i].TopLeftX / scale_ratio
		toplefty  := drawinfos[i].TopLeftY * (-1) / scale_ratio
		dowrightx := drawinfos[i].DownRightX / scale_ratio
		dowrighty := drawinfos[i].DownRightY * (-1) / scale_ratio
		rectangle := image.Rect(topleftx, toplefty, dowrightx, dowrighty)
		draw.Draw(img, rectangle, &image.Uniform{color}, image.ZP, draw.Src)
	}
    myfile, err := os.Create(new_png_file)
    if err != nil {
        panic(err)
    }
    png.Encode(myfile, img)
}