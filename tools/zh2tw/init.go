package main

import "sort"

func init() {
	// 作为map键的繁体没有重复
	// 但一个繁体可能对应多个简体, 需要按照key字典顺序导入
	// 只保留根据繁体key字典顺序第一个出现的简体
	kkMap := make([]int, 0, len(_TSCharactersMap))
	for k, _ := range _TSCharactersMap {
		kkMap = append(kkMap, int(k))
	}
	sort.Ints(kkMap)

	// 导入初始转换表
	for _, k := range kkMap {
		k := rune(k)
		v := _TSCharactersMap[k]
		tw2zhMap[k] = v
		zh2twMap[v] = k
	}

	// 修正错误的转换(仅简体到繁体)
	for k, v := range zh2twMapPatch {
		zh2twMap[k] = v
	}
}
