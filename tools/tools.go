package tools

import (
	"bufio"
	"log"
	"os"
	"path"
	"path/filepath"

	"github.com/BurntSushi/toml"
	uuid "github.com/satori/go.uuid"
)

type Default struct {
	TimeoutUpdate int  // default: 60 sec
	Debug         bool // default: true
	Token         string
}

// Contains config data
type Config struct {
	Default Default
}

// Return config data
func ParseConfigToml(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return &Config{}, err
	}

	var conf Config

	if _, err := toml.Decode(string(data), &conf); err != nil {
		return &Config{}, err
	}

	return &conf, nil
}

func GetNewUuid() string {
	return uuid.NewV4().String()
}

func ReadFileFromUploads() ([]byte, string) {
	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatalf("%v", err)
	}
	uploadsDir := filepath.Join(currentDir, "uploads")
	file, err := os.Open(path.Join(uploadsDir, "example.jpeg"))
	if err != nil {
		log.Fatalf("%v", err)
	}
	defer file.Close()

	fileInfo, _ := file.Stat()
	bytes := make([]byte, fileInfo.Size())
	buffer := bufio.NewReader(file)
	_, err = buffer.Read(bytes)
	if err != nil {
		log.Fatalf("%v", err)
	}

	return bytes, GetNewUuid()
}

func GetTempToken(userID int64) string {
	return "TOKEN"
}

func GetPath(userID int64) string {
	return "PATH"
}
