package slicebinpack

import (
	"fmt"
	"time"
	"strconv"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"math/rand"
	"os"
	"os/exec"

	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/fixed"
)

// create dir for new set of network slice requests
func Mkdir(dir string) {
	sh_cmd    := "mkdir -p " + dir
	input_cmd := sh_cmd
	err       := exec.Command("/bin/sh", "-c", input_cmd).Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("create dir for bin packing picture : logs/binapck/", dir)
}

// add label to Network Slice Requests
func addLabel(img *image.RGBA, x, y int, label string) {
	col := color.RGBA{255, 255, 255, 255}
	point := fixed.Point26_6{fixed.Int26_6((x - 7) * 64), fixed.Int26_6((y - 13) * 64)}

	d := &font.Drawer{
		Dst:  img,
		Src:  image.NewUniform(col),
		Face: basicfont.Face7x13,
		Dot:  point,
	}
	d.DrawString(label)
}

// draw deployment of network slice requests
func DrawBinPackResult(dir string, id string, drawinfos []DrawBlock, width int, height int, scale_ratio int) {
	img           := image.NewRGBA(image.Rect(0, height / scale_ratio * (-1), width / scale_ratio, 0))
	new_png_file  := dir + "/timewindow-" + id +".png"
	for i :=0; i< len(drawinfos); i++ {
		t          := time.Now().UnixNano()
		r          := rand.New(rand.NewSource(t))
		color_r    := uint8(r.Intn(255))
		color_g    := uint8(r.Intn(255))
		color_b    := uint8(r.Intn(255))
		color      := color.RGBA{color_r, color_g, color_b, 255}
		topleftx   := drawinfos[i].TopLeftX / scale_ratio
		toplefty   := drawinfos[i].TopLeftY * (-1) / scale_ratio
		downrightx := drawinfos[i].DownRightX / scale_ratio
		downrighty := drawinfos[i].DownRightY * (-1) / scale_ratio
		labelx     := int((topleftx + downrightx) / 2)
		labely     := int((toplefty + downrighty) / 2)
		label      := strconv.Itoa(i + 1)
		rectangle  := image.Rect(topleftx, toplefty, downrightx, downrighty)
		draw.Draw(img, rectangle, &image.Uniform{color}, image.ZP, draw.Src)
		addLabel(img, labelx, labely, label)
	}
    myfile, err   := os.Create(new_png_file)
    if err != nil {
        panic(err)
    }
    png.Encode(myfile, img)
}
