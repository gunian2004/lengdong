package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func main() {
	// // 读取 FrozenProduct.json 文件
	// fileContent, err := ioutil.ReadFile("./artifacts/contracts/FrozenProduct.sol/FrozenProduct.json")
	// if err != nil {
	// 	fmt.Printf("读取文件失败: %v\n", err)
	// 	return
	// }

	// // 定义一个结构体来解析 JSON 数据
	// var contractData struct {
	// 	Abi      interface{} `json:"abi"`
	// 	Bytecode string      `json:"bytecode"`
	// }

	// // 解析 JSON 数据
	// err = json.Unmarshal(fileContent, &contractData)
	// if err != nil {
	// 	fmt.Printf("解析 JSON 数据失败: %v\n", err)
	// 	return
	// }

	// // 提取 ABI 和字节码
	// abiBytes, err := json.Marshal(contractData.Abi)
	// if err != nil {
	// 	fmt.Printf("转换 ABI 为 JSON 字符串失败: %v\n", err)
	// 	return
	// }
	// abi := string(abiBytes)
	// bytecode := contractData.Bytecode

	// // 将 ABI 和字节码保存到文件
	// err = ioutil.WriteFile("FrozenProduct.abi", []byte(abi), 0644)
	// if err != nil {
	// 	fmt.Printf("保存 ABI 到文件失败: %v\n", err)
	// 	return
	// }
	// err = ioutil.WriteFile("FrozenProduct.bin", []byte(bytecode), 0644)
	// if err != nil {
	// 	fmt.Printf("保存字节码到文件失败: %v\n", err)
	// 	return
	// }

	// fmt.Println("ABI 和字节码已成功保存到文件")
		// 读取 FrozenProduct.json 文件
		fileContent, err := ioutil.ReadFile("./artifacts/contracts/FrozenProduct.sol/FrozenProduct.json")
		if err != nil {
			fmt.Printf("读取文件失败: %v\n", err)
			return
		}
	
		// 定义一个结构体来解析 JSON 数据
		var contractData struct {
			Abi      interface{} `json:"abi"`
			Bytecode string      `json:"bytecode"`
		}
	
		// 解析 JSON 数据
		err = json.Unmarshal(fileContent, &contractData)
		if err != nil {
			fmt.Printf("解析 JSON 数据失败: %v\n", err)
			return
		}
	
		// 提取 ABI 和字节码
		abiBytes, err := json.Marshal(contractData.Abi)
		if err != nil {
			fmt.Printf("转换 ABI 为 JSON 字符串失败: %v\n", err)
			return
		}
		abi := string(abiBytes)
		bytecode := contractData.Bytecode
	
		// 将 ABI 和字节码保存到文件
		err = ioutil.WriteFile("FrozenProduct.abi", []byte(abi), 0644)
		if err != nil {
			fmt.Printf("保存 ABI 到文件失败: %v\n", err)
			return
		}
		err = ioutil.WriteFile("FrozenProduct.bin", []byte(bytecode), 0644)
		if err != nil {
			fmt.Printf("保存字节码到文件失败: %v\n", err)
			return
		}
	
		fmt.Println("ABI 和字节码已成功保存到文件")
}
