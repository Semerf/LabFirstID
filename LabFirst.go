package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Добро пожаловать в программу задач по защите информации!")
	fmt.Println("Введите номер задачи")
	var number int
	_, err := fmt.Scan(&number)
	if err != nil {
		fmt.Println("Ошибка чтения!")
	}

	TaskSwitch(number)
}

func TaskSwitch(number int) {
	switch number {
	case 1:
		Task1()
		break
	case 2:
		TaskCrypt()
		break
	case 3:
		TaskUnCrypt()
		break
	case 4:
		TaskLab1()
		break
	case 5:
		TaskFA() //frequance analysis
		break
	}

}
func TaskFA() { //частотный анализ
	m := make(map[rune]int)
	var AllSymbols uint32
	fileText, err := os.ReadFile("Text.txt")
	if err != nil {
		fmt.Println("Ошибка чтения из файла! ", err)
	}
	str := string(fileText)
	for _, symb := range str {
		m[symb] += 1
		AllSymbols++
	}

	for key, value := range m {
		fmt.Println("Key:", string(key), "Value:", value, "\tPopularity:", float64(value)/float64(AllSymbols)*100.)
	}

}
func TaskLab1() {
	fmt.Println("Введите кодовую фразу...")
	var input string
	var inputData []byte

	fmt.Scan(&input)
	file, err := os.Open("Text.txt")
	crypted, err2 := os.Create("crypted.txt")
	uncrypted, err3 := os.Create("uncrypted.txt")
	if err != nil || err2 != nil || err3 != nil {
		panic("Ошибка чтения файла!")
	}
	for _, valueInp := range input {
		inputData = append(inputData, byte(valueInp))
	}
	var i byte
	for i = 0; i < 3; i++ {
		for ind := range input {
			inputData = append(inputData, (inputData[ind] - 1 - i))
		}
		fmt.Print(string(inputData[:]), "\n\n")
	}
	data := make([]byte, len(inputData))
	for {
		_, err := file.Read(data)
		for indJ, valJ := range inputData {
			data[indJ] = (data[indJ] | valJ) & (^(data[indJ] & valJ))
		}
		if err == io.EOF { // если конец файла
			break // выходим из цикла
		}
		fmt.Print(string(data[:]))
		crypted.Write(data)
	}
	crypted.Seek(0, 0)
	fmt.Print("\n\n")
	for {
		_, err := crypted.Read(data)
		for indJ, valJ := range inputData {
			data[indJ] = (data[indJ] | valJ) & (^(data[indJ] & valJ))
		}
		if err == io.EOF { // если конец файла
			break // выходим из цикла
		}
		fmt.Print(string(data[:]))
		uncrypted.Write(data)
	}
}

func GetArg(n int) []uint8 {
	fmt.Println("Введите ", n, " значения...")
	var Variable uint8
	var a []uint8 = []uint8{}
	for i := 0; i < n; i++ {
		fmt.Scan(&Variable)
		a = append(a, Variable)
	}
	return a
}
func Task1() {
	ch := GetArg(3)
	a, k, n := ch[0], ch[1], ch[2]
	var result uint8
	result = ((a & (^(1 << n))) & (^(1 << k))) | (((a >> k) & 1) << n) | (((a >> n) & 1) << k)
	fmt.Println("Результат равен = ", result)
}
func TaskCrypt() {
	fileText, err := os.ReadFile("Text.txt")
	var step int
	fmt.Println("Введите значение сдвига")
	fmt.Scan(&step)
	//crypted, err2 := os.Create("crypted.txt")
	if err != nil /*&& err2 != nil*/ {
		fmt.Println("Ошибка чтения из файла! ", err)
	}
	str := string(fileText)

	for _, symb := range str {
		symb += rune(step)
		if err == io.EOF { // если конец файла
			break // выходим из цикла
		}
		fmt.Print(string(symb))
		//crypted.Write(symb)

	}
}

func TaskUnCrypt() {
	var exitVar int = 1
	var CryptKey byte = 0
	for {
		file, err := os.Open("crypted.txt")
		uncrypted, err2 := os.Create("uncrypted.txt")
		if err != nil && err2 != nil {
			fmt.Println("Ошибка чтения из файла! ", err)
		}
		data := make([]byte, 1)

		for {
			n, err := file.Read(data)
			data[0] -= CryptKey
			if err == io.EOF { // если конец файла
				break // выходим из цикла
			}
			fmt.Print(string(data[:n]))
			uncrypted.Write(data)

		}
		fmt.Println("\nХотите сдвинуть еще на символ? 0 - нет")
		fmt.Scanln(&exitVar)
		if exitVar == 0 {
			break
		}
		CryptKey++
	}
}
