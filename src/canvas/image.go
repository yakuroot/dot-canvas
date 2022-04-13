package canvas

import (
	"bytes"
	"image/png"
	"os"
)

func GetImageBuffer() (*bytes.Buffer, error) {
	buf := new(bytes.Buffer)
	if err := png.Encode(buf, canvas.Image()); err != nil {
		return nil, err
	}

	return buf, nil
}

func SaveImage() error { return SaveNamedImage("canvas") }

func SaveNamedImage(fileName string) error {
	buf, err := GetImageBuffer()
	if err != nil {
		return err
	}

	path, err := os.Getwd()
	if err != nil {
		return err
	}

	if err := os.WriteFile(path+"/images/"+fileName+".png", buf.Bytes(), 0644); err != nil {
		return err
	}

	return nil
}

func RemoveImage(fileName string) error {
	path, err := os.Getwd()
	if err != nil {
		return err
	}

	if err := os.Remove(path + "/images/" + fileName + ".png"); err != nil {
		return err
	}

	return nil
}
