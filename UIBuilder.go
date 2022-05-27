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
	rootContainer := widget.NewContainer(widget.ContainerOpts.WidgetOpts(widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{
		StretchHorizontal: true,
	})),
		widget.ContainerOpts.Layout(widget.NewRowLayout(
			widget.RowLayoutOpts.Direction(widget.DirectionVertical),
			widget.RowLayoutOpts.Spacing(10),
		)),
	)
	gameUI := &ebitenui.UI{Container: rootContainer}
	theRowLAyout := widget.NewRowLayout(
		widget.RowLayoutOpts.Direction(widget.DirectionHorizontal),
		widget.RowLayoutOpts.Padding(widget.Insets{
			Top:  20,
			Left: 20,
		}),
		widget.RowLayoutOpts.Spacing(20),
	)
	topButtonsContainer := widget.NewContainer(
		widget.ContainerOpts.Layout(theRowLAyout))
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
	topBarWidgets := make([]widget.PreferredSizeLocateableWidget, 3, 10)
	topBarWidgets[0] = makeButton
	//topButtonsContainer.AddChild(makeButton)
	xEntryField := makeTextField(UIgraphics, "Enter X location")
	yEntryField := makeTextField(UIgraphics, "Enter Y location")
	topBarWidgets[1] = xEntryField
	topBarWidgets[2] = yEntryField
	topButtonsContainer.AddChild(xEntryField)
	topButtonsContainer.AddChild(yEntryField)
	topButtonsContainer.AddChild(makeButton)
	theRowLAyout.Layout(topBarWidgets, image.Rectangle{
		Min: image.Point{0, 0},
		Max: image.Point{1200, 200},
	})
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

func makeTextField(embeddedLoc embed.FS, prompt string) *widget.TextInput {
	tboxrealImg := loadNineSliceSimple("TextBox2.png", "UIGraphics", embeddedLoc, 22, 24)

	tboxImg := &widget.TextInputImage{
		Idle:     tboxrealImg,
		Disabled: nil,
	}
	textColor := &widget.TextInputColor{
		Idle: color.RGBA{
			R: 0x0f,
			G: 0x04,
			B: 0x0f,
			A: 0xff,
		},
		Disabled: color.RGBA{
			R: 0x5a,
			G: 0x7a,
			B: 0x91,
			A: 0xff,
		},
		Caret: color.RGBA{
			R: 0x0f,
			G: 0x04,
			B: 0x0f,
			A: 0xff,
		},
	}
	return widget.NewTextInput(
		widget.TextInputOpts.WidgetOpts(widget.WidgetOpts.LayoutData(widget.RowLayoutData{
			Stretch: true,
		})),
		widget.TextInputOpts.Image(tboxImg),
		widget.TextInputOpts.Color(textColor),
		widget.TextInputOpts.Face(basicfont.Face7x13),
		widget.TextInputOpts.CaretOpts(
			widget.CaretOpts.Size(basicfont.Face7x13, 2),
		),
		widget.TextInputOpts.Padding(widget.Insets{
			Left:   13,
			Right:  13,
			Top:    7,
			Bottom: 7,
		}),
		widget.TextInputOpts.Placeholder(prompt),
	)

}
