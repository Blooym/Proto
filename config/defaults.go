/*
Copyright © 2022 BitsOfAByte

GPLv3 License, see the LICENSE file for more information.
*/
package config

import (
	"os"

	"github.com/BitsOfAByte/proto/core"
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
	viper.SetDefault("storage.tmp", core.UsePath(os.TempDir(), true))

	// Configure app defaults
	viper.SetDefault("app.force", "false")
	viper.SetDefault("app.sources", []string{
		"GloriousEggroll/proton-ge-custom",
		"GloriousEggroll/wine-ge-custom",
	})
	viper.SetDefault("app.customlocations", map[string]string{
		"steam":  "~/.steam/root/compatibilitytools.d/",
		"lutris": "~/.local/share/lutris/runners/wine/",
	})

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			os.MkdirAll(configDir+"/proto", os.ModePerm)
			viper.SafeWriteConfig()
		} else {
			panic(err)
		}
	}

}
