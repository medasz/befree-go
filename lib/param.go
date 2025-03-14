package lib

import (
	"bufio"
	"crypto/tls"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

var (
	// 全局计数器,统计各类型节点总共获取数量
	TotalVmessCount  int
	TotalSsCount     int
	TotalSsrCount    int
	TotalTrojanCount int
)

// Exists 判断所给路径文件/文件夹是否存在
func Exists(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

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

func LoadSubscriptionUrls(inputFile string) ([]string, error) {
	var subscriptionUrls []string
	if !FileExists(inputFile) {
		return subscriptionUrls, fmt.Errorf("订阅文件未找到：%s", inputFile)
	}
	file, err := os.Open(inputFile)
	if err != nil {
		return subscriptionUrls, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		subscriptionUrls = append(subscriptionUrls, scanner.Text())
	}
	return subscriptionUrls, nil
}

func FetchAndParseSubscription(url string) ([]Node, error) {
	nodes := make([]Node, 0)
	httpClient := &http.Client{
		Timeout: time.Second * 10,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}
	resp, err := httpClient.Get(url)
	if err != nil {
		return nodes, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nodes, nil
	}
	fmt.Printf(" [+] 订阅获取成功： %s\n", url)
	response, err := io.ReadAll(resp.Body)
	if err != nil {
		return nodes, err
	}
	if strings.Contains(string(response), "proxy-groups") {
		return nodes, nil
	}
	decData, err := base64.StdEncoding.DecodeString(cleanBase64String(string(response)))
	if err != nil {
		return nodes, err
	}
	return ParseNodes(string(decData))
}

func cleanBase64String(base64String string) string {
	base64String = strings.ReplaceAll(base64String, "\n", "")
	base64String = strings.ReplaceAll(base64String, "\r", "")
	base64String = strings.TrimSpace(base64String)
	if len(base64String)%4 != 0 {
		base64String += strings.Repeat("=", 4-len(base64String)%4)
	}
	return base64String
}

func ParseNodes(rawData string) ([]Node, error) {
	nodes := make([]Node, 0)
	for _, line := range strings.Split(rawData, "\n") {
		var node Node
		var err error
		if strings.HasPrefix(line, "vmess://") {
			node, err = NewVMessNode(strings.TrimPrefix(line, "vmess://"))
			if err != nil {
				return nil, err
			}
			TotalVmessCount++
		} else if strings.HasPrefix(line, "ss://") {
			node, err = NewShadowsocksNode(strings.TrimPrefix(line, "ss://"))
			if err != nil {
				return nil, err
			}
			TotalSsCount++
		} else if strings.HasPrefix(line, "ssr://") {

		} else if strings.HasPrefix(line, "trojan://") {
			node, err = NewTrojanNode(strings.TrimPrefix(line, "trojan://"))
			if err != nil {
				return nil, err
			}
			TotalTrojanCount++
		}
		if node != nil {
			nodes = append(nodes, node)
		}
	}
	return nodes, nil
}

func FormatProxyConfig(proxy NodeProxy) string {
	res := ""
	res += fmt.Sprintf("  - name: %s\n", proxy.GetName())
	res += fmt.Sprintf("    type: %s\n", proxy.GetType())
	res += fmt.Sprintf("    server: %s\n", proxy.GetServer())
	res += fmt.Sprintf("    port: %d\n", proxy.GetPort())
	if proxy.GetPassword() != nil {
		res += fmt.Sprintf("    password: %s\n", proxy.GetPassword())
	}
	if proxy.GetCipher() != nil {
		res += fmt.Sprintf("    cipher: %s\n", proxy.GetCipher())
	}
	if proxy.GetUuid() != nil {
		res += fmt.Sprintf("    uuid: %s\n", proxy.GetUuid())
	}
	if proxy.GetAlterId() != nil {
		res += fmt.Sprintf("    alterId: %d\n", proxy.GetAlterId())
	}
	if proxy.GetSkipCertVerify() != nil {
		res += fmt.Sprintf("    skipCertVerify: %t\n", proxy.GetSkipCertVerify())
	}
	return res
}
