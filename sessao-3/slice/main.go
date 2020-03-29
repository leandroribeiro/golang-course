package main

import "fmt"

func main() {
	var numeros []int
	fmt.Println(numeros, len(numeros), cap(numeros))
	numeros = make([]int, 5)
	fmt.Println(numeros, len(numeros), cap(numeros))

	mundos := []string{"Tatoine"}
	fmt.Println(mundos, len(mundos), cap(mundos))
	mundos = append(mundos, "Datomir")
	fmt.Println(mundos, len(mundos), cap(mundos))

	jedis := make([]string, 4)
	jedis[0] = "Anakin"
	jedis[1] = "Yoda"
	jedis[2] = "Obiwan"
	jedis[3] = "Zayfo Wias"
	fmt.Println(jedis, len(jedis), cap(jedis))

	for _, jedi := range jedis {
		fmt.Println(jedi)
	}
}
