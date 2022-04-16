package main

import "C"
import (
	"fmt"
	"image"
	"net"
	"os"

	"github.com/makiuchi-d/gozxing"
	"github.com/makiuchi-d/gozxing/qrcode"
)

//export helloWorld
func helloWorld() string {
	return "Hello World"
}

//export getUid
func getUid() int{
	return net.IPv4len
}

//export decode
func decode(path string) string {
	file, err := os.Open(path)
	if err != nil {
		img, _, _ := image.Decode(file)
		bmp, _ := gozxing.NewBinaryBitmapFromImage(img)
		qrReader := qrcode.NewQRCodeReader()
		result, _ := qrReader.Decode(bmp, nil)
		fmt.Println(result)
		return ""
	} else {
		fmt.Println(err)
		return err.Error()
	}
}


func main() {}
