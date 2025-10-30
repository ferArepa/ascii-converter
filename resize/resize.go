package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"log"
	"os"
	"strconv"

	"github.com/nfnt/resize"
)

func resizeImage(inputPath, outputPath string, newWidth uint, newHeight uint) error {

	file, err := os.Open(inputPath)
	if err != nil {
		return fmt.Errorf("error al abrir el archivo: %w", err)
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return fmt.Errorf("error al decodificar la imagen: %w", err)
	}

	m := resize.Resize(newWidth, newHeight, img, resize.NearestNeighbor)

	out, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("error al crear el archivo de salida: %w", err)
	}
	defer out.Close()

	err = jpeg.Encode(out, m, &jpeg.Options{Quality: 90})
	if err != nil {
		return fmt.Errorf("error al codificar la imagen: %w", err)
	}

	return nil
}

func main() {

	i := 1

	for {
		err := resizeImage("ezgif-frame-"+strconv.FormatInt(int64(i), 10)+".jpg", "frame-"+strconv.FormatInt(int64(i), 10)+".jpg", 120, 72)
		if err != nil {
			log.Fatal(err)
		}
		i++
	}

}
