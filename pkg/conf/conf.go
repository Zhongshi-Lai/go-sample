package conf

import (
	"errors"

	"github.com/spf13/viper"
)

func ReadFromFile(confPath string, confName string) {
	viper.AddConfigPath(confPath)
	viper.SetConfigName(confName)
	viper.SetConfigType("toml")
	if err := viper.ReadInConfig(); err != nil {
		var configFileNotFoundError viper.ConfigFileNotFoundError
		if errors.As(err, &configFileNotFoundError) {
			panic("conf file not found")
		} else {
			panic(err)
		}
	}
}

func MergeFromFile(confName string) {
	// second file must use viper.MergeInConfig(); otherwise it will lose first conf content;
	viper.SetConfigName(confName)
	viper.SetConfigType("toml")
	if err := viper.MergeInConfig(); err != nil {
		var configFileNotFoundError viper.ConfigFileNotFoundError
		if errors.As(err, &configFileNotFoundError) {
			panic("conf file not found")
		} else {
			panic(err)
		}
	}
}

func AppConfInit(confPath string) {
	ReadFromFile(confPath, "server")
	MergeFromFile("log")
}

func ScriptConfInit() {

}
