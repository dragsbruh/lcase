package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func lowercaseFile(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	tempPath := path + ".temp"
	tempFile, err := os.Create(tempPath)
	if err != nil {
		return err
	}
	defer tempFile.Close()

	scanner := bufio.NewScanner(file)
	writer := bufio.NewWriter(tempFile)

	for scanner.Scan() {
		line := strings.ToLower(scanner.Text())
		_, err := writer.WriteString(line + "\n")
		if err != nil {
			return err
		}
	}

	err = scanner.Err()
	if err != nil {
		return err
	}

	err = writer.Flush()
	if err != nil {
		return err
	}

	file.Close()
	tempFile.Close()

	err = os.Rename(tempPath, path)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: lcase <path-to-file> [...<more-path-if-needed>]")
		os.Exit(1)
	}

	for i, path := range os.Args[1:] {
		err := lowercaseFile(path)
		if err != nil {
			fmt.Printf("lowercased only %d file(s) out of %d given file(s)\n", i, len(os.Args)-1)
			fmt.Printf("(%s) error: %v\n", path, err)
			os.Exit(1)
		}
	}

	fmt.Printf("lowercased all (%d) file(s)\n", len(os.Args)-1)
}
