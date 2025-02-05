package lib

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

func GenerateConfig(nodes []Node, outputFile string, port int, target string) error {
	// 去重節點名稱
	ResolveDuplicateNames(nodes)

	proxies := make([]string, 0)
	proxyNames := make([]string, 0)
	nameCount := make(map[string]int)
	for _, node := range nodes {
		nodeProxyIn := node.ToClashProxy()
		if nodeProxyIn != nil {
			name := nodeProxyIn.GetName()
			if _, ok := nameCount[name]; ok {
				nameCount[name]++
				name = fmt.Sprintf("%s_%d", name, nameCount[name])
			} else {
				nameCount[name] = 1
			}
			proxyNames = append(proxyNames, name)
			proxies = append(proxies, FormatProxyConfig(nodeProxyIn))
		}

	}
	// 获取当前程序运行的目录，并构建完整的输出路径
	currentDir, err := os.Getwd()
	if err != nil {
		return err
	}
	if !Exists(currentDir) {
		if err := os.Mkdir(currentDir, os.ModePerm); err != nil {
			return err
		}
	}
	outputPath := filepath.Join(currentDir, outputFile)
	fd, err := os.OpenFile(outputPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer fd.Close()

	w := bufio.NewWriter(fd)
	w.WriteString("allowLan: false\n")
	w.WriteString(fmt.Sprintf("mixed-port: %d\n", port))
	w.WriteString("rules:\n")
	w.WriteString("  - MATCH, proxy_pool\n")
	w.WriteString("proxy-groups:\n")
	w.WriteString("  - name: proxy_pool\n")
	w.WriteString("    type: load-balance\n")
	w.WriteString("    proxies:\n")
	for _, proxyName := range proxyNames {
		w.WriteString(fmt.Sprintf("      - %s\n", proxyName))
	}
	w.WriteString(fmt.Sprintf("    url: %s\n", target))
	w.WriteString("    interval: 5\n")
	w.WriteString("    strategy: round-robin\n")
	w.WriteString("proxies:\n")
	for _, proxy := range proxies {
		w.WriteString(fmt.Sprintf("%s\n", proxy))
	}
	w.Flush()
	fmt.Printf(" [+] http & socks 监听端口：%d\n", port)
	return nil
}

func ResolveDuplicateNames(nodes []Node) {
	tmp := make(map[string]int)
	for _, node := range nodes {
		if _, ok := tmp[node.GetName()]; ok {
			tmp[node.GetName()]++
			node.SetName(fmt.Sprintf("%s__%d", node.GetName(), tmp[node.GetName()]))
		} else {
			tmp[node.GetName()] = 1
		}
	}
}
