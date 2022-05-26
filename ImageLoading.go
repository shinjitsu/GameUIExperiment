package main

import (
	"embed"
	"github.com/hajimehoshi/ebiten/v2"
	"image/png"
	"log"
)

//go:embed UIGraphics/*
var UIgraphics embed.FS

//go:embed SpriteImages/*
var SpritePicts embed.FS

func loadPNGImageFromEmbedded(name string, embedLoc embed.FS) *ebiten.Image {
	pictNames, err := embedLoc.ReadDir("picts")
	if err != nil {
		log.Fatal("failed to read embedded dir ", pictNames, " ", err)
	}
	embeddedFile, err := embedLoc.Open("picts/" + name)
	if err != nil {
		log.Fatal("failed to load embedded image ", embeddedFile, err)
	}
	rawImage, err := png.Decode(embeddedFile)
	if err != nil {
		log.Fatal("failed to load embedded image ", name, err)
	}
	gameImage := ebiten.NewImageFromImage(rawImage)
	return gameImage
}
