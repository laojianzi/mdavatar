/*
 * Material-Design-Avatars
 * https://github.com/lincanbin/Material-Design-Avatars
 *
 * Copyright 2015 Canbin Lin (lincanbin@hotmail.com)
 * http://www.94cb.com/
 *
 * Licensed under the Apache License, Version 2.0:
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Create material deisgn avatars for users just like Gmail or Messager in Android.
 */

// Ref: https://blog.logrocket.com/working-with-go-images/
// Ref: https://github.com/zdhxiong/mdclub/tree/master/src/Vendor
package mdavatar

import (
	"image"
	"image/color"
	"image/draw"
	"io/ioutil"

	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"golang.org/x/image/math/fixed"

	"github.com/laojianzi/mdavatar/util"
)

// Config on mdavatar build
type Config struct {
	avatarText          string
	avatarTextHandle    AvatarTextHandler
	avatarSize          int
	padding             int
	letterFont          string
	letterFontContent   []byte
	asianFont           string
	asianFontContent    []byte
	enableAsianFontChar bool
	colors              []color.RGBA
	background          *image.RGBA
}

// AvatarTextHandler handle for avatar text
// allow custom !!!
// e.g:
//    text: laojianzi => L
//    text: 老健仔 => L
type AvatarTextHandler func(s string, enableAsianFontChar bool) string

// New return a mdavatar config (*Config)
func New(text string, opts ...Option) *Config {
	defaultConfig := *DefaultConfig
	config := &defaultConfig
	for _, opt := range opts {
		opt(config)
	}

	if config.asianFont != "" {
		config.enableAsianFontChar = util.IsFile(config.asianFont)
	}

	config.avatarText = config.avatarTextHandle(text, config.enableAsianFontChar)
	return config
}

// Build generate a image from config
func (config *Config) Build() (*image.RGBA, error) {
	var err error
	if config.background == nil {
		if _, err = config.MakeBackground(); err != nil {
			return nil, err
		}
	}

	return config.drawText()
}

// Background new a background image from Config.avatarSize and Config.colors
func (config *Config) MakeBackground() (*image.RGBA, error) {
	bgColorIndex, err := util.MtRand(0, int64(len(config.colors)-1))
	if err != nil {
		return nil, err
	}

	bgColor := image.Uniform{C: config.colors[bgColorIndex]}
	config.background = image.NewRGBA(image.Rect(0, 0, config.avatarSize, config.avatarSize))
	draw.Draw(config.background, config.background.Bounds(), &image.Uniform{C: &bgColor},
		image.Point{}, draw.Src)

	return config.background, nil
}

func (config *Config) drawText() (*image.RGBA, error) {
	var (
		fgColor  image.Image
		fontSize = 128.0
		err      error
	)
	fgColor = image.White

	fontContent, err := config.fontContent()
	if err != nil {
		return nil, err
	}

	fontFace, err := opentype.Parse(fontContent)
	if err != nil {
		return nil, err
	}

	face, err := opentype.NewFace(fontFace, &opentype.FaceOptions{
		Size:    fontSize,
		DPI:     72.0,
		Hinting: font.HintingFull,
	})
	if err != nil {
		return nil, err
	}

	fontDrawer := &font.Drawer{
		Dst:  config.background,
		Src:  fgColor,
		Face: face,
	}
	textBounds, _ := fontDrawer.BoundString(config.avatarText)
	xPosition := (fixed.I(config.background.Rect.Max.X) - fontDrawer.MeasureString(config.avatarText)) / 2
	textHeight := textBounds.Max.Y - textBounds.Min.Y
	yPosition := fixed.I((config.background.Rect.Max.Y)-textHeight.Ceil())/2 + fixed.I(textHeight.Ceil())
	fontDrawer.Dot = fixed.Point26_6{
		X: xPosition,
		Y: yPosition,
	}
	fontDrawer.DrawString(config.avatarText)
	return config.background, err
}

func (config *Config) fontContent() ([]byte, error) {
	var content []byte
	var err error

	if config.enableAsianFontChar {
		if config.asianFont != "" {
			content, err = ioutil.ReadFile(config.asianFont)
			if err != nil {
				return nil, err
			}

			config.asianFontContent = content
		}

		content = []byte(config.asianFontContent)
	} else {
		if config.letterFont != "" {
			content, err = ioutil.ReadFile(config.letterFont)
			if err != nil {
				return nil, err
			}

			config.letterFontContent = content
		}

		content = []byte(config.asianFontContent)
	}

	return content, nil
}
