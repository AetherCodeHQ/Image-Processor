package main

import (
	"encoding/binary"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: image-processor <file>")
		os.Exit(1)
	}
	data, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	if len(data) < 8 {
		fmt.Println("not a valid image")
		os.Exit(1)
	}
	if data[0] == 0x89 && data[1] == 0x50 && data[2] == 0x4e && data[3] == 0x47 {
		w := binary.BigEndian.Uint32(data[16:20])
		h := binary.BigEndian.Uint32(data[20:24])
		fmt.Printf("format=PNG width=%d height=%d size=%d\n", w, h, len(data))
		return
	}
	if data[0] == 0xff && data[1] == 0xd8 {
		fmt.Printf("format=JPEG size=%d bytes\n", len(data))
		return
	}
	if data[0] == 0x47 && data[1] == 0x49 && data[2] == 0x46 {
		w := binary.LittleEndian.Uint16(data[6:8])
		h := binary.LittleEndian.Uint16(data[8:10])
		fmt.Printf("format=GIF width=%d height=%d size=%d\n", w, h, len(data))
		return
	}
	fmt.Println("unknown image format")
	os.Exit(1)
}
