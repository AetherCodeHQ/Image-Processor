package main

import (
	"fmt"
	"os"
)

// image_processor - Image processing tool
func image_processor(path string) {
	fmt.Println("========================================")
	fmt.Println("  Image-Processor")
	fmt.Println("  Image processing tool")
	fmt.Println("========================================")
	fmt.Println()
	fmt.Println("Target:", path)
	fmt.Println("Processing...")
	fmt.Println("Done!")
}

func main() {
	path := "."
	if len(os.Args) > 1 {
		path = os.Args[1]
	}
	image_processor(path)
}
