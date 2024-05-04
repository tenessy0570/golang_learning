package internal

import (
	"os"
	"path"
)

var RootPath, _ = os.Getwd()

var dbFileName = "test.db"
var DbPath = path.Join(RootPath, "internal", "database", dbFileName)
