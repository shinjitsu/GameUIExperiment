package main

import (
	"embed"
	"fmt"
	"github.com/blizzy78/ebitenui"
	uiImage "github.com/blizzy78/ebitenui/image"
	"github.com/blizzy78/ebitenui/widget"
	"golang.org/x/image/font/basicfont"
	"image"
	"image/color"
)

func MakeUI() *ebitenui.UI {
	rootContainer := widget.NewContainer()
	gameUI := &ebitenui.UI{Container: rootContainer}
	topButtonsContainer := widget.NewContainer(
		widget.ContainerOpts.Layout(widget.NewRowLayout(
			widget.RowLayoutOpts.Direction(widget.DirectionHorizontal),
			widget.RowLayoutOpts.Padding(widget.Insets{
				Top:  20,
				Left: 20,
			}),
		)))
	topButtonsContainer.SetLocation(image.Rectangle{
		Min: image.Point{
			X: 0,
			Y: 0,
		},
	})

	makeButtonImg := &widget.ButtonImage{
		Idle:     loadButtonImage("buttons_63.png", UIgraphics),
		Hover:    loadButtonImage("buttons_64.png", UIgraphics),
		Pressed:  loadButtonImage("buttons_61.png", UIgraphics),
		Disabled: nil,
	}
	makeButton := widget.NewButton(
		// specify the images to use
		widget.ButtonOpts.Image(makeButtonImg),
		widget.ButtonOpts.Text("Press Me to get another one", basicfont.Face7x13,
			&widget.ButtonTextColor{
				Idle: color.RGBA{0xdf, 0x04, 0x0f, 0xff},
			}),
		widget.ButtonOpts.TextPadding(widget.Insets{
			Left:   10,
			Right:  10,
			Top:    20,
			Bottom: 20,
		}),
		widget.ButtonOpts.ClickedHandler(makeButtonPressed),
	)
	topButtonsContainer.AddChild(makeButton)
	rootContainer.AddChild(topButtonsContainer)
	return gameUI
}

func loadButtonImage(imagePict string, embeddedLoc embed.FS) *uiImage.NineSlice {
	pict, err := loadImageNineSlice(imagePict, "UIGraphics", embeddedLoc, 14, 13)
	if err != nil {
		fmt.Errorf("ERROR LOADING IMAGE %s", err)
	}
	return pict
}
