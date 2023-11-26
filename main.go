//go:generate go install -v github.com/josephspurrier/goversioninfo/cmd/goversioninfo
//go:generate goversioninfo -icon=res/papp.ico -manifest=res/papp.manifest
package main

import (
	"os"

	"github.com/portapps/portapps/v3"
	"github.com/portapps/portapps/v3/pkg/log"
	"github.com/portapps/portapps/v3/pkg/utl"
)

var (
	app *portapps.App
)

func init() {
	var err error

	// Init app
	if app, err = portapps.New("stm32cubemx-portable", "STM32CubeMX"); err != nil {
		log.Fatal().Err(err).Msg("Cannot initialize application. See log file for more info.")
	}
}

func main() {
	utl.CreateFolder(app.DataPath)
	app.Process = utl.PathJoin(app.AppPath, "STM32CubeMX.exe")

	os.Setenv("STM32CubeMX_PATH", app.AppPath)
	os.Setenv("_JAVA_OPTIONS", "-Duser.home=" + app.DataPath)

	updater_file := utl.PathJoin(app.DataPath, ".stm32cubemx/plugins/updater/updater.ini")
	if utl.Exists(updater_file) {
		utl.ReplaceByPrefix(updater_file, "SoftwarePath=", "SoftwarePath="+app.AppPath)
		utl.ReplaceByPrefix(updater_file, "RepositoryPath=", "RepositoryPath="+utl.PathJoin(app.DataPath, "STM32Cube/Repository/"))
		utl.ReplaceByPrefix(updater_file, "UpdaterPath=", "UpdaterPath="+utl.PathJoin(app.DataPath, ".stm32cubemx/plugins/updater/"))
	}

	defer app.Close()
	app.Launch(os.Args[1:])
}
