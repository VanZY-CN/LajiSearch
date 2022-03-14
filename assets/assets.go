package assets

import (
	"embed"
)

//go:embed dict.txt minwpdict.txt ua.txt
var Dict embed.FS