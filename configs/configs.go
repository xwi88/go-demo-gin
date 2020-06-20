// Package configs store global config
package configs

import (
	"errors"
	"fmt"
	"path/filepath"
	"time"

	"github.com/spf13/viper"
	"github.com/xwi88/kit4go/datetime"
)

var cfg *Config

// Init global configs init, must be invoke when server pre run or resources init
func Init(file string) error {
	if len(file) == 0 {
		return errors.New("blank file")
	}
	// use New, can avoid global viper obj!!!
	// here we shall use default, so we can export all configs to extern files
	// v := viper.New()
	// v.SetConfigFile(file)
	// err := v.ReadInConfig()

	viper.SetConfigFile(file)
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}
	err = viper.Unmarshal(&cfg)
	return err
}

// GetCfg get global config
func GetCfg() *Config {
	return cfg
}

// New returns an initialized Config instance.
func New() *Config {
	cfg := new(Config)
	return cfg
}

// Reset is intended for testing, will reset all to default settings.
// In the public interface for the configs package so applications
// can use it in their testing as well.
func Reset() {
	cfg = New()
}

// WriteConfig export config, maybe overwrite the exist file
func WriteConfig(filename string) error {
	return writeConfig(filepath.Dir(filename), filepath.Base(filename), true)
}

// SafeWriteConfig export config
func SafeWriteConfig(filename string) error {
	return writeConfig(filepath.Dir(filename), filepath.Base(filename), false)
}

// WriteConfigWithPath export config file with special dir, maybe overwrite the exist file
func WriteConfigWithPath(pathname, filename string) error {
	return writeConfig(pathname, filename, true)
}

// SafeWriteConfigWithPath export config file with special dir
func SafeWriteConfigWithPath(pathname, filename string) error {
	return writeConfig(pathname, filename, false)
}

func writeConfig(pathname, filename string, force bool) (err error) {
	_dir := filepath.Dir(pathname)
	_filename := filepath.Base(filename)
	newFilename := fmt.Sprintf("%v%c%v", _dir, filepath.Separator, _filename)
	if force {
		err = viper.WriteConfigAs(newFilename)
	} else {
		err = viper.SafeWriteConfigAs(newFilename)
	}
	if err != nil {
		_newFilename := fmt.Sprintf("%v%c%v_%v", _dir, filepath.Separator,
			time.Now().Format(datetime.LayoutDateTimeShort), _filename)
		err = viper.WriteConfigAs(_newFilename)
	}
	return err
}
