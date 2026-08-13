package main

import (
	"fmt"
	"strconv"
	"strings"
)

type IPAddr [4]byte

// TODO: Add a "String() string" method to IPAddr.
func (ip IPAddr) String() string {
	// Convert IP address bytes to dotted decimal notation
	return fmt.Sprintf("%d.%d.%d.%d", ip[0], ip[1], ip[2], ip[3])
}

// DynamicByteArray - пример для массива байтов произвольного размера
type DynamicByteArray []byte

// String() для динамического массива байтов
func (dba DynamicByteArray) String() string {
	if len(dba) == 0 {
		return ""
	}

	// Способ 1: Используя strings.Builder (эффективный)
	var builder strings.Builder
	for i, b := range dba {
		if i > 0 {
			builder.WriteString(".")
		}
		builder.WriteString(strconv.Itoa(int(b)))
	}
	return builder.String()
}

// StringWithFmt - альтернативный способ с fmt.Sprintf
func (dba DynamicByteArray) StringWithFmt() string {
	if len(dba) == 0 {
		return ""
	}

	// Создаем слайс интерфейсов для fmt.Sprintf
	args := make([]interface{}, len(dba))
	for i, b := range dba {
		args[i] = int(b)
	}

	// Создаем формат строку динамически
	format := strings.Repeat("%d.", len(dba)-1) + "%d"
	return fmt.Sprintf(format, args...)
}

// StringSimple - простой способ с конкатенацией строк
func (dba DynamicByteArray) StringSimple() string {
	if len(dba) == 0 {
		return ""
	}

	result := strconv.Itoa(int(dba[0]))
	for i := 1; i < len(dba); i++ {
		result += "." + strconv.Itoa(int(dba[i]))
	}
	return result
}

func main() {
	// Оригинальный код с IPAddr
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	fmt.Println("=== IP Addresses ===")
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}

	// Примеры с динамическими массивами байтов
	fmt.Println("\n=== Dynamic Byte Arrays ===")

	// IPv4 адрес
	ipv4 := DynamicByteArray{192, 168, 1, 1}
	fmt.Printf("IPv4: %s\n", ipv4.String())

	// IPv6 адрес (первые 8 байт)
	ipv6 := DynamicByteArray{0x20, 0x01, 0x0d, 0xb8, 0x85, 0xa3, 0x08, 0xd3}
	fmt.Printf("IPv6 (first 8 bytes): %s\n", ipv6.String())

	// Произвольный массив
	random := DynamicByteArray{10, 20, 30, 40, 50}
	fmt.Printf("Random bytes: %s\n", random.String())

	// Сравнение методов
	fmt.Println("\n=== Method Comparison ===")
	testBytes := DynamicByteArray{1, 2, 3, 4, 5}
	fmt.Printf("String(): %s\n", testBytes.String())
	fmt.Printf("StringWithFmt(): %s\n", testBytes.StringWithFmt())
	fmt.Printf("StringSimple(): %s\n", testBytes.StringSimple())
}
