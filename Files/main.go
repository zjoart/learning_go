package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Hello!, Files in Golang")

	//creating a new file
	var newfile *os.File
	fmt.Println(newfile)

	nfile, err := os.Create("a.txt")
	if err != nil {
		// fmt.Println(err)
		// os.Exit(1)
		log.Fatal(err)
	}
	nfile.Close()

	// os.OpenFile() to open file, FileINfo for fileinfo  , os.Stat()

	// to rename file os.Rename(oldPath, newPath)
	fmt.Println("Wrting bytes to file")
	// opening the file in write-only mode if the file exists and then it truncates the file.
	// if the file doesn't exist it creates the file with 0644 permissions
	file, err := os.OpenFile(
		"b.txt",
		os.O_WRONLY|os.O_TRUNC|os.O_CREATE,
		0644,
	)
	// error handling
	if err != nil {
		log.Fatal(err)
	}
	// defer closing the file
	defer file.Close()

	// WRITING BYTES TO FILE

	byteSlice := []byte("I learn Golang! 传")   // converting a string to a bytes slice
	bytesWritten, err := file.Write(byteSlice) // writing bytes to file.
	// It returns the no. of bytes written and an error value
	// error handling
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Bytes written: %d\n", bytesWritten) // => 2019/10/21 16:26:16 Bytes written: 19

	//riting to file using Buffered writer
	newFile, err := os.OpenFile("ayo.txt", os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer newFile.Close()

	bufferedWriter := bufio.NewWriter(newFile)
	bs := []byte{97, 98, 90}
	bytesWrittens, err := bufferedWriter.Write(bs)

	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%d", bytesWrittens)
}
