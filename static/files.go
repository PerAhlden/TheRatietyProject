package static

import (
	_ "embed"
	"go.wdy.de/nago/application"
)

//go:embed vote-for-blog.jpg
var Logo application.StaticBytes
