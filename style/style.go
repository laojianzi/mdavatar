package style

import (
	"image"
)

// MDAvatarBuildStyle mdavatar style type
type MDAvatarBuildStyle image.Image

// MDAvatarBuildStyleNewFunc func for new a mdavatar style
type MDAvatarBuildStyleNewFunc func(image.Image) MDAvatarBuildStyle
