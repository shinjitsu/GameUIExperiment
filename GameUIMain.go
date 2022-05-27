package main

import (
	"fmt"
	"github.com/blizzy78/ebitenui"
	"github.com/blizzy78/ebitenui/widget"
	"github.com/hajimehoshi/ebiten/v2"
)

type EbitenGame struct {
	gameUI          *ebitenui.UI
	drawOptions     ebiten.DrawImageOptions
	firstDrawnStuff []Sprite
}

type Sprite struct {
	XLoc int
	YLoc int
	pict *ebiten.Image
}

func main() {
	ebiten.SetWindowSize(1200, 1200)
	ebiten.SetWindowTitle("Test Ebiten UI with New UI Images and Sprites")
	theGame := &EbitenGame{
		gameUI:          MakeUI(),
		drawOptions:     ebiten.DrawImageOptions{},
		firstDrawnStuff: nil,
	}
	ebiten.RunGame(theGame)
}

func (e *EbitenGame) Update() error {
	e.gameUI.Update()
	return nil
}

func (e EbitenGame) Draw(screen *ebiten.Image) {
	e.gameUI.Draw(screen)
}

func (e EbitenGame) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func makeButtonPressed(args *widget.ButtonClickedEventArgs) {
	fmt.Println("MAKE BUTTON PRESSED!!!!!!!!!")
}
