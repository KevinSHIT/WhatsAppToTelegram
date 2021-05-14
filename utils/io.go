package utils

import (
	"bufio"
	"bytes"
	"image"
	"image/jpeg"
	"os"
)

func Exists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

func IsFile(path string) bool {
	return !IsDir(path)
}

func SavePic(imgByte []byte, path string) error {

	img, _, err := image.Decode(bytes.NewReader(imgByte))
	if err != nil {
		return err
	}

	out, _ := os.Create(path)
	defer out.Close()

	var opts jpeg.Options
	opts.Quality = 100

	err = jpeg.Encode(out, img, &opts)
	return err
}

func FileToBytes(path string) []byte {
	file, err := os.Open(path)

	if err != nil {
		return nil
	}

	defer file.Close()

	fileInfo, _ := file.Stat()
	size := fileInfo.Size()
	bytesObj := make([]byte, size)

	buffer := bufio.NewReader(file)
	_, err = buffer.Read(bytesObj)

	if err != nil {
		return nil
	}

	return bytesObj

}
