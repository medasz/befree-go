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
	flag.StringVar(&Target, "t", "https://www.google.com", "Specify a link for speed testing(default:https://www.google.com)")
	flag.StringVar(&YamlFile, "y", "", "Specify a yourself clash yaml file")
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
	} else {
		fmt.Println("我的天空！ Befree v0.4")

		// 1.加載訂閲
		subscriptionUrls, err := lib.LoadSubscriptionUrls(InputFile)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Printf(" [+] %s文件中，发现%d 个订阅地址\n", InputFile, len(subscriptionUrls))

		// 2.請求訂閲並解析
		allNodes := make([]lib.Node, 0)
		for _, url := range subscriptionUrls {
			fmt.Printf(" [+] 正在处理订阅地址：%s\n", url)
			nodes, err := lib.FetchAndParseSubscription(url)
			if err != nil {
				fmt.Println(err)
				continue
			}
			allNodes = append(allNodes, nodes...)
		}

		fmt.Printf(" [+] 总共解析到%d 个正常转换节点\n", len(allNodes))
		// 输出所有协议类型节点的总数
		fmt.Printf(" [+] 其中包含vmess节点数量为: %d\n", lib.TotalVmessCount)
		fmt.Printf(" [+] 其中包含ss节点数量为: %d\n", lib.TotalSsCount)
		fmt.Printf(" [+] 其中包含trojan节点数量为: %d", lib.TotalTrojanCount)

		if len(allNodes) > 0 {
			outputFile := "sectest.yaml"
			lib.GenerateConfig(allNodes)
			//4.运行clash
			if err := lib.Runner(outputFile, ClashPath); err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		} else {
			fmt.Println(" [-] 未获取到可用节点，无法启动befree")
		}
	}
}
