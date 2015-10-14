package config

import (
	"errors"
	"io/ioutil"
)

type ConfigFileManager struct {
	Filename string
	*MultiAccountConfig
}

func NewConfigFileManager(configFile string) *ConfigFileManager {
	if len(configFile) == 0 {
		configFile = FindConfigFile()
	}

	return &ConfigFileManager{
		Filename: configFile,
	}
}

func (configFileManager *ConfigFileManager) ParseConfigFile() error {
	var err error

	configFileManager.MultiAccountConfig, err = ParseConfig(configFileManager.Filename)
	if err != nil {
		return errInvalidConfigFile
	}

	return nil
}

func (configFileManager *ConfigFileManager) Add(conf Config) error {
	if configFileManager.keepXmppClientCompat {
		return errors.New("Cant add accounts while in compat mode")
	}

	configFileManager.MultiAccountConfig.Add(conf)

	return nil
}

func (configFileManager *ConfigFileManager) Save() error {
	if configFileManager.keepXmppClientCompat {
		account := configFileManager.MultiAccountConfig.Accounts[0]
		account.Filename = configFileManager.Filename
		return account.Save()
	}

	contents, err := configFileManager.Serialize()
	if err != nil {
		return err
	}

	return ioutil.WriteFile(configFileManager.Filename, contents, 0600)
}
