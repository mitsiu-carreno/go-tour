package ziputil

import (
	"archive/zip"
	"io"
	"os"
)

// ZipFiles compress one or many files into a single zip archive file (name, array of files to zip)
func ZipFiles(filename string, files[]string) error {
	newFile, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer newFile.Close()

	zipWriter := zip.NewWriter(newFile)
	defer zipWriter.Close()

	// Add files to zip
	for _, file := range files{
		zipfile, err := os.Open(file)
		if err != nil {
			return err
		}
		defer zipfile.Close()

		// Get the file information
		info, err := zipfile.Stat()
		if err != nil {
			return err
		}

		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}

		// Change to deflate to gain better compression
		header.Method = zip.Deflate

		writer, err := zipWriter.CreateHeader(header)
		if err != nil {
			return err
		}

		_, err = io.Copy(writer, zipfile)
		if err != nil{
			return err
		}
		zipfile.Close()
	}
	return nil
}