package testfiles

import "embed"

//go:embed *.txt
var TestFiles embed.FS
