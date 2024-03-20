package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	for i := 0; i < 1000; i++ {
		bytesContent := bytes.NewBuffer([]byte("lá ele!"))

		fileName := fmt.Sprintf("la-ele-%d.txt", i)
		file, err := os.Create(fileName)
		if err != nil {
			fmt.Println("Error creating file:", err)
			return
		}

		_, err = io.Copy(file, bytesContent)
		if err != nil {
			fmt.Println("Error copying content to file:", err)
			return
		}

		file.Close()
	}
}
