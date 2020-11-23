package mdavatar

import (
	"image"
	"image/color"
)

// Option set field value for Config
type Option func(*Config)

// WithAvatarTextHandle with option for Config.avatarTextHandle field
func WithAvatarTextHandle(avatarTextHandle AvatarTextHandler) Option {
	return func(config *Config) {
		config.avatarTextHandle = avatarTextHandle
	}
}

// WithAvatarSize with option for Config.avatarSize field
func WithAvatarSize(avatarSize int) Option {
	return func(config *Config) {
		config.avatarSize = avatarSize
	}
}

// WithPadding with option for Config.padding field
func WithPadding(padding int) Option {
	return func(config *Config) {
		config.padding = padding
	}
}

// WithLetterFont with option for Config.letterFont field
func WithLetterFont(letterFont string) Option {
	return func(config *Config) {
		config.letterFont = letterFont
	}
}

// WithAsianFont with option for Config.asianFont field
func WithAsianFont(asianFont string) Option {
	return func(config *Config) {
		config.asianFont = asianFont
	}
}

// DisableAsianFontChar set Config.enableAsianFontChar is false and reset Config.asianFont is empty
func DisableAsianFontChar() Option {
	return func(config *Config) {
		config.enableAsianFontChar = false
		config.asianFont = ""
	}
}

// WithColors with option for Config.colors field
func WithColors(colors []color.RGBA) Option {
	return func(config *Config) {
		config.colors = colors
	}
}

// WithBackground with option for Config.background field
func WithBackground(background *image.RGBA) Option {
	return func(config *Config) {
		config.background = background
	}
}
