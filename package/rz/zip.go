package rz

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"sync"
)

func (rz *Rz) Unzip(zipFile, destDir string) error {
	reader, err := zip.OpenReader(zipFile)
	if err != nil {
		return err
	}
	defer reader.Close()

	var wg sync.WaitGroup
	fileChan := make(chan *zip.File)

	// 启动一定数量的 goroutine 进行并发处理
	concurrency := runtime.NumCPU()
	for i := 0; i < concurrency; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for file := range fileChan {
				err := extractFile(file, destDir)
				if err != nil {
					fmt.Println("Error extracting file:", err)
				}
			}
		}()
	}

	// 将 ZIP 文件中的文件发送到通道
	for _, file := range reader.File {
		fileChan <- file
	}
	close(fileChan)

	// 等待所有处理完成
	wg.Wait()

	return nil
}

func extractFile(file *zip.File, destDir string) error {
	path := filepath.Join(destDir, file.Name)

	if file.FileInfo().IsDir() {
		os.MkdirAll(path, os.ModePerm)
		return nil
	}

	writer, err := os.Create(path)
	if err != nil {
		return err
	}
	defer writer.Close()

	reader, err := file.Open()
	if err != nil {
		return err
	}
	defer reader.Close()

	_, err = io.Copy(writer, reader)
	if err != nil {
		return err
	}

	return nil
}
