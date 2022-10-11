/*
Copyright © 2022 BitsOfAByte

GPLv3 License, see the LICENSE file for more information.
*/
package config

import (
	"BitsOfAByte/proto/shared"
	"os"

	"github.com/spf13/viper"
)

/*
	SetDefaults sets the default values for the configuration file.
*/
func SetDefaults() {
	configDir, _ := os.UserConfigDir()
	viper.SetConfigName("config")
	viper.AddConfigPath(configDir + "/proto")
	viper.SetConfigType("toml")

	// Configure CLI Defaults
	viper.SetDefault("cli.verbose", "false")

	// Configure storage defaults
	viper.SetDefault("storage.tmp", shared.UsePath("/tmp/proto/", true))

	// Configure app defaults
	viper.SetDefault("app.force", "false")
	viper.SetDefault("app.sources", []string{
		"GloriousEggroll/proton-ge-custom",
		"GloriousEggroll/wine-ge-custom",
		"bottlesdevs/wine/",
		"lutris/wine/",
	})
	viper.SetDefault("app.customlocations", map[string]string{
		"steam":  "~/.steam/root/compatibilitytools.d/",
		"lutris": "~/.local/share/lutris/runners/wine/",
	})
}
