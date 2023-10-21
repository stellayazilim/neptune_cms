package utils

import (
	"os"
	"path/filepath"
	"strings"
)

func GetRootDir() string {
	wd, _ := os.Getwd()
	for !strings.HasSuffix(wd, "neptune_cms") {
		wd = filepath.Dir(wd)
	}

	return wd
}
