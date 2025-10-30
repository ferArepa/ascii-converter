package main

import (
	"fmt"
	"image"
	"image/color"
	_ "image/jpeg"
	_ "image/png"
	"os"
	"os/exec"
	"strconv"
	"time"
)

func limpiarConsola() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

// var mapaRotadores = make(map[int]string)
var charset = make(map[int]byte)

func obtenerBrillo(r uint32, g uint32, b uint32) float32 {
	brillo := (0.3 * float32(r)) + (0.587 * float32(g)) + (0.114 * float32(b))
	return brillo
}

// func mostrarRotador() {

// 	mapaRotadores[0] = "|"
// 	mapaRotadores[1] = "/"
// 	mapaRotadores[2] = "-"
// 	mapaRotadores[3] = "\\"

// 	var contador int = 0

// 	for {
// 		fmt.Println(mapaRotadores[contador] + " cargando por favor espere...")
// 		time.Sleep(time.Millisecond * 100)
// 		limpiarConsola()
// 		contador++
// 		if contador >= 4 {
// 			contador = 0
// 		}

// 	}

// }

func obtenerCaracterSegunBrillo(brillo uint32) byte {

	charset[0] = '.'
	charset[1] = ','
	charset[2] = '-'
	charset[3] = '~'
	charset[4] = ':'
	charset[5] = ';'
	charset[6] = '='
	charset[7] = '!'
	charset[8] = '*'
	charset[9] = '#'
	charset[10] = '$'
	charset[11] = '@'

	charsetCounter := 0

	if int(brillo) <= 22 {
		return charset[0]
	}

	for i := 0; i < int(brillo); i += 22 {
		charsetCounter++
	}
	if charsetCounter > 11 {

		charsetCounter = 11
	}
	return charset[charsetCounter]
}

func renderizarImagen(matriz [][]byte) {
	for i := 0; i < len(matriz); i++ {
		linea := string(matriz[i])
		fmt.Println(linea)
	}
	time.Sleep(time.Millisecond * 100)

	limpiarConsola()
}

func cargarImagen(path string) image.Image {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	image, _, err := image.Decode(file)
	if err != nil {
		panic(err)
	}

	return image
}

func main() {

	var outputImage [][]byte
	var row []byte

	for i := 1; i < 32; {

		image := cargarImagen("resize/frame-" + strconv.FormatInt(int64(i), 10) + ".jpg")
		i++
		imageRectangle := image.Bounds()

		for i := imageRectangle.Min.Y; i < imageRectangle.Max.Y; i += 1 {

			if i > imageRectangle.Max.Y {
				break
			}

			for j := imageRectangle.Min.X; j < imageRectangle.Max.X; j++ {

				colore := image.At(j, i)

				brillo := color.GrayModel.Convert(colore)
				tono, _, _, _ := brillo.RGBA()

				caracter := obtenerCaracterSegunBrillo(tono / 256)
				row = append(row, caracter)

			}
			outputImage = append(outputImage, row)
			row = []byte{}

		}

		renderizarImagen(outputImage)

		outputImage = [][]byte{}
		row = []byte{}
		if i == 32 {
			i = 1
		}
	}

}
