package main

import (
	"image"

	don "github.com/hajimehoshi/ebiten"
)

func draw(screen *don.Image) {
	op := &don.DrawImageOptions{}

	// draw targets
	for _, t := range targets {
		op.GeoM.Reset()
		op.GeoM.Translate(t.x-targetImages[0].halfx, t.y-targetImages[0].halfy)
		screen.DrawImage(targetImages[0].img, op)
	}

	// draw hooks
	for _, h := range hooks {
		op.GeoM.Reset()
		op.GeoM.Translate(h.x-hookImage.halfx, h.y-hookImage.halfy)
		screen.DrawImage(hookImage.img, op)
	}

	// draw bullets
	for _, b := range bullets {
		op.GeoM.Reset()
		op.GeoM.Translate(-bulletImage.halfx, -bulletImage.halfy)
		op.GeoM.Rotate(b.angle)
		op.GeoM.Translate(b.x, b.y)
		screen.DrawImage(bulletImage.img, op)
	}

	// draw player
	op.GeoM.Reset()
	op.GeoM.Translate(-playerImage.halfx, -playerImage.halfy)
	op.GeoM.Rotate(player.angle)
	op.GeoM.Translate(player.x, player.y)
	screen.DrawImage(playerImage.img, op)

	// draw gauges
	ofsx, ofsy := 0.0, 48.0
	gaugex, gaugey := player.x+ofsx, player.y+ofsy

	// draw gauge
	rate := player.charge / 100.0
	sx, sy := gaugeImage.halfx*2, gaugeImage.halfy*2
	rect := &image.Rectangle{
		Min: image.Point{0, 0},
		Max: image.Point{int(sx * rate), int(sy)},
	}
	op.SourceRect = rect
	op.GeoM.Reset()
	op.GeoM.Translate(gaugex-gaugeImage.halfx, gaugey-gaugeImage.halfy)
	screen.DrawImage(gaugeImage.img, op)

	// draw gauge's frame
	op.SourceRect = nil
	op.GeoM.Reset()
	op.GeoM.Translate(gaugex-gFrameImage.halfx, gaugey-gFrameImage.halfy)
	screen.DrawImage(gFrameImage.img, op)

}
