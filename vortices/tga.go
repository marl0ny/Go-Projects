/* TGA file creation management.

Reference:

Wikipedia - Truevision TGA
https://en.wikipedia.org/wiki/Truevision_TGA

*/
package main

import (
	"log"
	"os"
)

func WriteTGA(filename string,
			  image_data []byte,
			  image_w, image_h int) {
	// This follows the example given for opening a file found
	// in the os package documentation for the OpenFile function:
	// https://pkg.go.dev/os
	f, err := os.OpenFile(filename, os.O_WRONLY | os.O_CREATE, 0644)
	if err == nil {
		// TODO: Use header structs instead of using a bytes
		// array.
		tga_header_bytes := [18]byte {
			0, // id length
			0, // color map type
			2, // image type
			// TGA color map spec
			0, 0, // Color map offset
			0, 0, // Color map entry count
			0, // Bits per entry
			// TGA image spec
			0, 0, // x origin
			0, 0, // y origin
			byte(image_w % 256), byte(image_w/256), // pixel width
			byte(image_h % 256), byte(image_h/256), // pixel height
			24, // bits per pixel
			0, // image descriptor
		}
		_, err2 := f.WriteAt(tga_header_bytes[0:18], 0)
		if err2 != nil {
			log.Fatal(err2)
		}
		_, err3 := f.WriteAt(image_data[0:3*(image_w*image_h)], 18)
		if err3 != nil {
			log.Fatal(err3)
		}
		err4 := f.Close()
		if err4 != nil {
			log.Fatal(err4)
		}
	} else {
		log.Fatal(err)
	}
}
