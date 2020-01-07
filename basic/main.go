package main

func main() {
	constDemo()
	bitDemo()
	boolDemo()
	intDemo()
	byteDemo()
	stringDemo()
	convertDemo()
	arrayDemo()
	sliceDemo()
	mapDemo()
	objectDemo()

	if result, err := testFuncWithMultiReturn(41, 6, "/"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("result: %d\n", result)
	}

	// 命令行参数
	fmt.Println(os.Args)

	// 返回值
	// 不会调用 defer
	os.Exit(-1)
}
