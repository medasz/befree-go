package lib

import "fmt"

func GenerateConfig(nodes []Node) {
	// 去重節點名稱

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
