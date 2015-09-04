package heron

import (
    "os"
    "encoding/json"
)

type ConfigFile struct {
    Debug bool `json:"debug"`
    LogPath string `json:"log_path"`
}

var Config ConfigFile

func ParseJsonFile(path string, v interface{}) {
    file, err := os.Open(path)
    if err != nil {
        logger.Fatal("Fail to open file:", err)
    }
    defer file.Close()
    decoder := json.NewDecoder(file)
    err = decoder.Decode(v)
    if err != nil {
        logger.Fatal("Fail to parse file:", err)
    }
}
