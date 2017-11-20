package main

import (
	"fmt"
	"image/color"
	"log"

	don "github.com/hajimehoshi/ebiten"
	donutil "github.com/hajimehoshi/ebiten/ebitenutil"
)

const (
	screenWidth  = 640
	screenHeight = 480
)

type imgInfo struct {
	img          *don.Image
	halfx, halfy float64
}

var (
	playerImage, bulletImage imgInfo
	hookImage                imgInfo
	gaugeImage, gFrameImage  imgInfo
	targetImages             []imgInfo
)

func update(screen *don.Image) error {
	if don.IsRunningSlowly() {
		return nil
	}
	updateGame()
	screen.Fill(color.NRGBA{0x00, 0x00, 0x00, 0xff})
	draw(screen)
	return nil
}

func main() {

	targetImages = make([]imgInfo, 1)

	imgPath := "docs/_resources/"
	var err error

	imgFiles := map[string]*imgInfo{
		"player.png":       &playerImage,
		"bullet.png":       &bulletImage,
		"hook.png":         &hookImage,
		"gauge.png":        &gaugeImage,
		"gauge_frame.png":  &gFrameImage,
		"target_green.png": &targetImages[0],
	}

	for f, i := range imgFiles {
		i.img, _, err = donutil.NewImageFromFile(imgPath+f, don.FilterNearest)
		if err != nil {
			err = fmt.Errorf("%s\nloading failed: %s", err, f)
		}
		sx, sy := i.img.Size()
		i.halfx, i.halfy = float64(sx/2), float64(sy/2)
	}

	if err != nil {
		log.Fatal(err)
	}

	initGame()

	if err := don.Run(update, screenWidth, screenHeight, 2, "hook-and-shoot game"); err != nil {
		log.Fatal(err)
	}
}
