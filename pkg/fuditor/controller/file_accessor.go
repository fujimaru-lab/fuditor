package controller

import (
	"io/ioutil"
	"os"
)

// ReadFromFile ファイルからテキスト情報を読み込む
func ReadFromFile(filepath string) ([]byte, error) {
	reader, err := os.Open(filepath)
	defer reader.Close()

	filedata, err := ioutil.ReadAll(reader)
	return filedata, err
}
