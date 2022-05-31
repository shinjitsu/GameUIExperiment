package main

import (
	"embed"
	"log"

	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
)

//This file was originally from the EbitenUI demo applications but modiries to handle embedded files
const (
	fontFaceRegular = "fonts/NotoSans-Regular.ttf"
	fontFaceBold    = "NotoSans-Bold.ttf"
)

type fonts struct {
	face         font.Face
	titleFace    font.Face
	bigTitleFace font.Face
	toolTipFace  font.Face
}

//func loadFonts() (*fonts, error) {
//	fontFace, err := loadFont(fontFaceRegular, 20)
//	if err != nil {
//		return nil, err
//	}
//
//	titleFontFace, err := loadFont(fontFaceBold, 24)
//	if err != nil {
//		return nil, err
//	}
//
//	bigTitleFontFace, err := loadFont(fontFaceBold, 28)
//	if err != nil {
//		return nil, err
//	}
//
//	toolTipFace, err := loadFont(fontFaceRegular, 15)
//	if err != nil {
//		return nil, err
//	}
//
//	return &fonts{
//		face:         fontFace,
//		titleFace:    titleFontFace,
//		bigTitleFace: bigTitleFontFace,
//		toolTipFace:  toolTipFace,
//	}, nil
//}

func (f *fonts) close() {
	if f.face != nil {
		_ = f.face.Close()
	}

	if f.titleFace != nil {
		_ = f.titleFace.Close()
	}

	if f.bigTitleFace != nil {
		_ = f.bigTitleFace.Close()
	}
}

func loadFontFromEmbedded(path string, size float64, embedLoc embed.FS, folder string) (font.Face, error) {
	uiFolder, err := embedLoc.ReadDir(folder)
	if err != nil {
		log.Fatal("failed to read embedded dir ", uiFolder, " ", err)
	}
	embededFontFile, err := embedLoc.ReadFile(folder + "/" + path)
	if err != nil {
		log.Fatal("failed to open embedded font", embededFontFile, err)
	}

	ttfFont, err := truetype.Parse(embededFontFile)
	if err != nil {
		return nil, err
	}

	return truetype.NewFace(ttfFont, &truetype.Options{
		Size:    size,
		DPI:     72,
		Hinting: font.HintingFull,
	}), nil
}
