package lib

import (
	"fmt"
	"io"
	"os"
)

func FileExists(place string) bool {
	f, err := os.Stat(place)
	if err != nil {
		return false
	}
	return !f.IsDir()
}

func moveFile(src, dst string) error {
	// 检查目标文件是否存在
	_, err := os.Stat(dst)
	if err == nil {
		// 如果目标文件存在，则删除目标文件
		err = os.Remove(dst)
		if err != nil {
			return fmt.Errorf("failed to remove existing destination file: %w", err)
		}
	} else if !os.IsNotExist(err) {
		// 如果检查目标文件时发生错误
		return fmt.Errorf("failed to check destination file: %w", err)
	}

	// 打开源文件
	srcFile, err := os.Open(src)
	if err != nil {
		return fmt.Errorf("failed to open source file: %w", err)
	}
	defer srcFile.Close()

	// 创建目标文件
	dstFile, err := os.Create(dst)
	if err != nil {
		return fmt.Errorf("failed to create destination file: %w", err)
	}
	defer dstFile.Close()

	// 复制内容
	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		return fmt.Errorf("failed to copy file: %w", err)
	}

	return nil
}
