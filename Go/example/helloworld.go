package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func main() {
	args := os.Args
	fmt.Println("Hello, World! " + args[0])

	// 適当な数字のスライス nums を作成
	nums := []int{1, 2, 3, 4, 5}
	// nums を標準出力する
	fmt.Println(nums)
	// nums に 106 を追加
	nums = append(nums, 106)
	// nums を標準出力する
	fmt.Println(nums)
	// 適当な文字列のスライス strs を作成
	strs := []string{"a", "b", "c", "d", "e"}
	// strs を標準出力する
	fmt.Println(strs)
	// strs をJSON文字列に変更する
	jsonStrs := fmt.Sprintf("%v", strs)
	// jsonStrs を標準出力する
	fmt.Println(jsonStrs)

	bytes, err := json.Marshal(strs)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(bytes))
}
