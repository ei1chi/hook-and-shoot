package main

import (
	"math"
	"math/rand"
	"time"

	don "github.com/hajimehoshi/ebiten"
)

type Player struct {
	x, y      float64
	angle     float64
	charge    float64
	chargeSpd float64
	recast    int
}

type Bullet struct {
	x, y  float64
	speed float64
	angle float64
}

type Hook struct {
	x, y float64
}

type Target struct {
	x, y float64
}

var (
	player  *Player
	bullets []*Bullet
	hooks   []*Hook
	targets []*Target

	triggerd bool
)

const (
	bulMax = 3
)

func initGame() {
	rand.Seed(time.Now().UnixNano())

	targets = make([]*Target, 10)
	for i, _ := range targets {
		targets[i] = &Target{
			x: float64(rand.Intn(screenWidth)),
			y: float64(rand.Intn(screenHeight)),
		}
	}

	player = &Player{
		x:         screenWidth / 2,
		y:         screenHeight / 2,
		chargeSpd: 1,
	}

	hooks = make([]*Hook, 1, 10)
	for i, _ := range hooks {
		hooks[i] = &Hook{
			x: float64(rand.Intn(screenWidth)),
			y: float64(rand.Intn(screenHeight)),
		}
	}
}

func updateGame() {
	px, py := don.CursorPosition()
	dx, dy := player.x-float64(px), player.y-float64(py)
	nextAngle := math.Atan2(dy, dx)

	for _, bul := range bullets {
		bul.x += math.Cos(bul.angle) * bul.speed
		bul.y += math.Sin(bul.angle) * bul.speed
	}

	triggers := don.IsMouseButtonPressed(don.MouseButtonLeft)
	if triggers {
		if len(bullets) < bulMax {
			player.charge += player.chargeSpd
		}
	} else {
		player.charge = 0.0
		player.angle = nextAngle
	}

	if player.charge > 100.0 {
		player.charge = 0.0
		bullets = append(bullets, &Bullet{
			x:     player.x,
			y:     player.y,
			speed: 3.0,
			angle: player.angle,
		})
	}
	triggerd = triggers
}
