/*slice in go*/

arr := []string{"Hello world", "Test"}

fmt.Println(len(arr), cap(arr), arr)
fmt.Println(arr[0:2])

************
2, 2, Hello world Test
Hello world Test