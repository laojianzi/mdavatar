package mdavatar

import (
	"strings"

	"github.com/laojianzi/mdavatar/util"
)

// DefaultConfig config on mdavatar default
// can be used directly
var DefaultConfig = &Config{
	avatarTextHandle: DefaultAvatarTextHandle,
	avatarSize:       256,
	letterFont:       "static/Roboto-Light.ttf",
	asianFont:        "static/NotoSansSC-Regular.otf",
	colors:           DefaultColors,
}

// DefaultColors colors on mdavatar default
var DefaultColors = mdColors

// DefaultAvatarTextHandle mdavatar default handler for avatar text
func DefaultAvatarTextHandle(text string, enableAsianFontChar bool) string {
	char := string([]rune(strings.ToUpper(text))[:1])
	if !enableAsianFontChar {
		char, _ = util.FirstLetter(char)
	}

	return char
}
