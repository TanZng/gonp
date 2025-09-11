package main

import "image/color"

type Config struct {
	ScreenWidth, ScreenHeight int

	LayoutWidth, LayoutHeight int
	Paint                     color.Color
}

func NewConfig() *Config {
	screenWidth := 1080
	screenHeight := 720
	return &Config{
		ScreenWidth:  screenWidth,
		ScreenHeight: screenHeight,

		LayoutWidth:  screenWidth / 4,
		LayoutHeight: screenHeight / 4,

		Paint: color.RGBA{140, 211, 134, 255},
	}
}
