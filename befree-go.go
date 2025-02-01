package main

import (
	"flag"
	"fmt"
	"os"

	"befree-go/lib"
)

var (
	InputFile string
	Port      int
	Target    string
	YamlFile  string
	ClashPath string
)

const banner = "\n _______             ________                                       ______            \n|       \\           |        \\                                     /      \\           \n| $$$$$$$\\  ______  | $$$$$$$$______    ______    ______          |  $$$$$$\\  ______  \n| $$__/ $$ /      \\ | $$__   /      \\  /      \\  /      \\  ______ | $$ __\\$$ /      \\ \n| $$    $$|  $$$$$$\\| $$  \\ |  $$$$$$\\|  $$$$$$\\|  $$$$$$\\|      \\| $$|    \\|  $$$$$$\\\n| $$$$$$$\\| $$    $$| $$$$$ | $$   \\$$| $$    $$| $$    $$ \\$$$$$$| $$ \\$$$$| $$  | $$\n| $$__/ $$| $$$$$$$$| $$    | $$      | $$$$$$$$| $$$$$$$$        | $$__| $$| $$__/ $$\n| $$    $$ \\$$     \\| $$    | $$       \\$$     \\ \\$$     \\         \\$$    $$ \\$$    $$\n \\$$$$$$$   \\$$$$$$$ \\$$     \\$$        \\$$$$$$$  \\$$$$$$$          \\$$$$$$   \\$$$$$$ \n"

func init() {
	flag.StringVar(&InputFile, "f", "./aaa.txt", "Specify a contain subscribe file path")
	flag.IntVar(&Port, "p", 59981, "Specify a port number(http&socks5)")
	flag.StringVar(&Target, "t", "", "Specify a link for speed testing(default:https://www.google.com)")
	flag.StringVar(&YamlFile, "y", "sectest.yaml", "Specify a yourself clash yaml file")
	flag.StringVar(&ClashPath, "c", "", "Specify your custom clash.exe path")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", banner)
		flag.PrintDefaults()
	}
	flag.Parse()
}

func main() {
	if YamlFile != "" {
		if lib.FileExists(YamlFile) {
			fmt.Printf(" [+] 检测到 %s 文件，程序正在启动...\n", YamlFile)
			if err := lib.Runner(YamlFile, ClashPath); err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		}
	}
}
