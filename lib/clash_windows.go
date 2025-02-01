//go:build windows

package lib

import (
	_ "embed"
	"fmt"
	"os"
	"os/exec"
)

//go:embed clash-windows-amd64.exe
var ClashWinBin []byte

func Runner(yamlFile string, clashPath string) error {
	// 判斷指定文件存不存在
	if clashPath != "" && !FileExists(clashPath) {
		return fmt.Errorf("clash path %s does not exist", clashPath)
	}
	// 如果爲空設置默認值
	if clashPath == "" {
		clashPath = "clash-windows-amd64.exe"
	}
	// 默認值不存在就自動創建
	if !FileExists(clashPath) {
		err := os.WriteFile(clashPath, ClashWinBin, 0644)
		if err != nil {
			return err
		}
		fmt.Println(" [+] start!!!")
	}
	Execute(yamlFile, clashPath)

	return nil
}

func Execute(yamlFile string, clashPath string) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in Execute", r)
		}
	}()
	fmt.Println(" [+] running...")
	cmd := exec.Command("cmd", "/C", clashPath, "-f", yamlFile)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	if err := cmd.Run(); err != nil {
		panic(err)
	}
	fmt.Println(" [-] stop...")
}
