/*unsafe pointer work*/


import (
	"fmt"
	"unsafe"
)

func main() {
  var test_arr [4]byte
	test_arr[0] = 10
	test_arr[1] = 100
	test_arr[2] = 101
	test_arr[3] = 102

	start_adr := unsafe.Pointer(&test_arr[0]) // создаем небезопасный указатель на 1 элемент массива
	adr := uintptr(start_adr) // преобразуем указатель в число с которым можно производить арифмитические действия

	p_arr := (*byte)(start_adr) // преобразуем к стандартному указателю на байт
	fmt.Println(*p_arr, p_arr) // выводим значение по этому адресу и сам адрес

	adr = adr + 1 // инкрементируем переменную идентичную адресу начала массива
	start_adr = unsafe.Pointer(adr) // преобразуем новый адрес в небезопасный указатель
	p_arr = (*byte)(start_adr) // преобразуем в стандартный указатель на байт
	fmt.Println(*p_arr, p_arr) // вывод

	adr = adr + 1
	start_adr = unsafe.Pointer(adr)
	p_arr = (*byte)(start_adr)
	fmt.Println(*p_arr, p_arr)

	*p_arr = 100 // изменяем значение по адресу p_arr
	fmt.Println(*p_arr, p_arr) // выводим

	fmt.Println(test_arr) // отличается от первоначального буфера
}