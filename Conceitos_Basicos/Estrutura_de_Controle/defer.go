package main

import (
	"io"
	"os"
)

func main() {
	err := createTxtAndWrite("hello.txt", "Hello, World!")
	if err != nil {
		println("Error creating file:", err)
		return
	}
	println("File created and written successfully")
}

func createTxtAndWrite(fileName, text string) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close() // O defer faz com que independente de qualquer ramificação no código, por exemplo, com if, o arquivo será fechado e irá libera o recurso de volta pr osistema.

	_, err = io.WriteString(file, text)
	if err != nil {
		return err
	}
	return nil
}
