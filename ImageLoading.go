package main

import (
	"embed"
	"github.com/blizzy78/ebitenui/image"
	"github.com/hajimehoshi/ebiten/v2"
	"image/png"
	"log"
)

//go:embed UIGraphics/*
var UIgraphics embed.FS

//go:embed SpriteImages/*
var SpritePicts embed.FS

func loadPNGImageFromEmbedded(name string, embedLoc embed.FS, folder string) *ebiten.Image {
	pictNames, err := embedLoc.ReadDir(folder)
	if err != nil {
		log.Fatal("failed to read embedded dir ", pictNames, " ", err)
	}
	embeddedFile, err := embedLoc.Open(folder + "/" + name)
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

func loadImageNineSlice(path string, folder string, embedLoc embed.FS, centerWidth int, centerHeight int) (*image.NineSlice, error) {
	img := loadPNGImageFromEmbedded(path, embedLoc, folder)

	w, h := img.Size()
	return image.NewNineSlice(img,
			[3]int{(w - centerWidth) / 2, centerWidth, w - (w-centerWidth)/2 - centerWidth},
			[3]int{(h - centerHeight) / 2, centerHeight, h - (h-centerHeight)/2 - centerHeight}),
		nil
}

//func loadButtonResources() (*widget.ButtonImage, error) {
//	idle, err := loadImageNineSlice("button-idle.png", 12, 0)
//	if err != nil {
//		return nil, err
//	}
//
//	hover, err := loadImageNineSlice("button-hover.png", 12, 0)
//	if err != nil {
//		return nil, err
//	}
//
//	pressed, err := loadImageNineSlice("button-pressed.png", 12, 0)
//	if err != nil {
//		return nil, err
//	}
//
//	disabled, err := loadImageNineSlice("button-disabled.png", 12, 0)
//	if err != nil {
//		return nil, err
//	}
//
//	buttonImages := &widget.ButtonImage{
//		Idle:     idle,
//		Hover:    hover,
//		Pressed:  pressed,
//		Disabled: disabled,
//	}
//	return buttonImages, nil
//}
