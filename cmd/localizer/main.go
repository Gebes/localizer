package main

import (
	"github.com/Gebes/localizer/pkg/localization"
	"github.com/Gebes/localizer/pkg/logger"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		logger.Error.Fatalln("Provide one localizations folder")
	}
	err := localization.Update(args[0])
	if err != nil {
		logger.Error.Fatalln("Could not update languages:", err)
	}
}
