package main

import (
	"github.com/blizzy78/ebitenui"
	"github.com/blizzy78/ebitenui/widget"
	"image"
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
	makeButton := widget.Button{
		Image:             nil,
		KeepPressedOnExit: false,
		GraphicImage:      nil,
		TextColor:         nil,
		PressedEvent:      nil,
		ReleasedEvent:     nil,
		ClickedEvent:      nil,
	}
	rootContainer.AddChild(topButtonsContainer)
	return gameUI
}
