package internal

import (
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

var Config = &AppConfig{}

type AppConfig struct {
	FilePathConfig FilePathConfig `yaml:"filePathConfig"`
	BasicConfig    BasicConfig    `yaml:"basicConfig"`
}

type BasicConfig struct {
	MaxDataFileSize uint64 `yaml:"maxDataFileSize int64"`
	MaxKeySize      uint32 `yaml:"maxKeySize"`
	MaxValueSize    uint64 `yaml:"maxKeySize"`
	Sync            bool   `yaml:"maxKeySize"`
}
type FilePathConfig struct {
	BasePath     string `yaml:"basePath"`
	IndexPath    string `yaml:"indexFile"`
	DataFilePath string `yaml:"dataFilePath"`
}

func loadConfig(filename string, conf interface{}) error {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Printf("read config file failure: %v", err)
		return err
	}
	if err = yaml.Unmarshal(data, conf); err != nil {
		log.Printf("parse config file failure: %v", err)
		return err
	}
	return nil
}

func init() {
	// defaults.SetDefaults(Config)
	loadConfig("../config.yml", Config)
	fmt.Println("--------------EasyDB LoadConfig-------------")
	initFileErr := initFilePath(Config.FilePathConfig)
	if initFileErr != nil {
		fmt.Errorf("init file error %v", initFileErr)
	}
	fmt.Println("--------------EasyDB LoadConfig Success-------------")
}

func initFilePath(config FilePathConfig) error {
	basePath := config.BasePath
	indexPath := config.IndexPath
	dataFilepath := config.DataFilePath
	if err := CreatePathIfNotExists(basePath); err != nil {
		return err
	}
	if err := CreateFileIfNotExists(indexPath); err != nil {
		return err
	}
	if err := CreatePathIfNotExists(dataFilepath); err != nil {
		return err
	}
	fmt.Println("init file path success")
	return nil
}
