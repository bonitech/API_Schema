package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

const (
	// DebugMode is const for debug
	DebugMode string = "debug"
	// ReleaseMode is const for release
	ReleaseMode string = "release"
)

// Info is golbal variable
var Info *configStruct

// ConfigStruct is JSON structure
type configStruct struct {
	Env    envStruct    `json:"Env"`
	Db     dbStruct     `json:"Db"`
	Resdis resdisStruct `json:"Resdis"`
}

// env info
type envStruct struct {
	Port string `json:"port"`
}

// db info
type dbStruct struct {
	URL  string `json:"url"`
	Port string `json:"port"`
	User string `json:"user"`
	Pass string `json:"pass"`
}

// resdis info
type resdisStruct struct {
	URL  string `json:"url"`
	Port string `json:"port"`
	User string `json:"user"`
	Pass string `json:"pass"`
}

// Init Config struct
func Init(mode string) {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)
	// set default path for dev
	jsonPath := exPath + "/../../config/dev.json"
	// set path for prod
	if mode == ReleaseMode {
		jsonPath = exPath + "/config/prod.json"
	}

	jsonFile, err := os.Open(jsonPath)

	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal(byteValue, &Info)
}
