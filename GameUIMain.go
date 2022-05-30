package main

import (
	"github.com/blizzy78/ebitenui"
	"github.com/blizzy78/ebitenui/widget"
	"github.com/hajimehoshi/ebiten/v2"
	"log"
	"strconv"
)

var theGame *EbitenGame

type EbitenGame struct {
	gameUI          *daddyGUI
	drawOptions     ebiten.DrawImageOptions
	firstDrawnStuff []Sprite
}

type daddyGUI struct {
	xEntryField   *widget.TextInput
	yEntryField   *widget.TextInput
	mainContainer *ebitenui.UI
}

type Sprite struct {
	XLoc int
	YLoc int
	pict *ebiten.Image
}

func main() {
	ebiten.SetWindowSize(1200, 1200)
	ebiten.SetWindowTitle("Test Ebiten UI with New UI Images and Sprites")
	theGame = &EbitenGame{
		gameUI:          MakeUI(),
		drawOptions:     ebiten.DrawImageOptions{},
		firstDrawnStuff: nil,
	}
	ebiten.RunGame(theGame)
}

func (e *EbitenGame) Update() error {
	e.gameUI.mainContainer.Update()
	for loc, _ := range e.firstDrawnStuff {
		e.firstDrawnStuff[loc].XLoc += 3
		if e.firstDrawnStuff[loc].XLoc > 1200 {
			e.firstDrawnStuff[loc].XLoc = -e.firstDrawnStuff[loc].pict.Bounds().Size().X
		}
	}
	return nil
}

func (e EbitenGame) Draw(screen *ebiten.Image) {
	e.gameUI.mainContainer.Draw(screen)
	for _, sprite := range e.firstDrawnStuff {
		e.drawOptions.GeoM.Reset()
		e.drawOptions.GeoM.Translate(float64(sprite.XLoc), float64(sprite.YLoc))
		screen.DrawImage(sprite.pict, &e.drawOptions)
	}
}

func (e EbitenGame) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func makeButtonPressed(args *widget.ButtonClickedEventArgs) {
	xString := theGame.gameUI.xEntryField.InputText
	yString := theGame.gameUI.yEntryField.InputText
	yLoc := 0
	xLoc := 0
	var err error = nil
	if xString != "" {
		xLoc, err = strconv.Atoi(xString)
		if err != nil {
			log.Fatal("error xloc not an int", err)
		}
	}
	if yString != "" {
		yLoc, err = strconv.Atoi(yString)
		if err != nil {
			log.Fatal("error yloc not an int", err)
		}
	}
	newOne := Sprite{
		XLoc: xLoc,
		YLoc: yLoc,
		pict: loadPNGImageFromEmbedded("Kill Bot.png", SpritePicts, "SpriteImages"),
	}
	theGame.firstDrawnStuff = append(theGame.firstDrawnStuff, newOne)
	theGame.gameUI.xEntryField.InputText = ""
	theGame.gameUI.yEntryField.InputText = ""
}
