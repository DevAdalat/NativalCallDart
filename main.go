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
func helloWorld() *C.char {
	return C.CString("Hello World")
}

//export getUid
func getUid() int{
	return net.IPv4len
}

//export decode
func decode(path *C.char) *C.char  {
	file, err := os.Open(C.GoString(path))
	if err != nil {
		img, _, _ := image.Decode(file)
		bmp, _ := gozxing.NewBinaryBitmapFromImage(img)
		qrReader := qrcode.NewQRCodeReader()
		result, _ := qrReader.Decode(bmp, nil)
		fmt.Println(result)
		fmt.Println(result.GetText())
		return C.CString(result.GetText())
	} else {
		fmt.Println(err)
		return C.CString(err.Error())
	}
}


func main() {}
