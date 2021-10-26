package main

import (
	"io"
	"log"
	"os"
)

type IFileOpener interface {
	OpenReader(filename string) io.Reader
	OpenWriter(filename string) io.Writer
	CloseReader()
	CloseWriter()
}

type File struct {
	reader *os.File
	writer *os.File
}

func NewFileOpener() *File {
	return &File{}
}

func (f *File) OpenReader(filename string) io.Reader {
	reader, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Unable to open %q: %v", filename, err)
	}
	f.reader = reader
	return f.reader
}

func (f *File) OpenWriter(filename string) io.Writer {
	writer, err := os.Create(filename)
	if err != nil {
		log.Fatalf("Unable to open %q: %v", filename, err)
	}
	f.writer = writer
	return f.writer
}

func (f *File) CloseReader() {
	f.reader.Close()
}

func (f *File) CloseWriter() {
	f.writer.Close()
}
