package chapter8

import (
	"fmt"
	"sort"
)

// 构造一个将英文饮料名映射为法语（或者任意你的母语）的集合；先打印所有的饮料，
// 然后打印原名和翻译后的名字。接下来按照英文名排序后再打印出来。

var Drinks =  map[string]string {
	"Coca Cola" : "可口可乐",
	"Vodka"	: "伏特加",
	"Fanta" : "芬达",
}

func SortDrinks (drinks map[string]string) {
	fmt.Printf("drinks map: %v", drinks)		// 先打印所有的饮料

	drinksKeys := make([]string, 0, len(drinks))
	drinksVals := make([]string, 0, len(drinks))
	for drink, drinkInChinese := range drinks {
		drinksKeys = append(drinksKeys, drink)
		drinksVals = append(drinksVals, drinkInChinese)
	}

	fmt.Printf("English name of drinks: %v", drinksKeys) // 打印原名
	fmt.Printf("Chinese name of drinks: %v", drinksVals) // 打印翻译名

	sort.Strings(drinksKeys)

	for _, k := range drinksKeys {
		fmt.Printf("drinks key: %s, drinks val: %s", k, drinks[k])	// 按照英文名排序后再打印出来
	}
}



