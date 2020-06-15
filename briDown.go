package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func main() {
	guardar(leer())
}

func guardar(cantidad string) {
	c, _ := strconv.Atoi(cantidad)
	c -= 25
	println(strconv.Itoa(c), "Cadena")
	file, err := os.OpenFile("/sys/class/backlight/intel_backlight/brightness", os.O_WRONLY, 0644)
	if err != nil {
		log.Println("Funciona", err.Error())
	} else {
		file.WriteString(strconv.Itoa(c))
	}
	file.Close()
}

func leer() string {
	var reg string
	file, err := os.Open("/sys/class/backlight/intel_backlight/actual_brightness")
	if err != nil {
		log.Fatal("Error al abrir el archivo  " + err.Error())
	} else {
		sc := bufio.NewScanner(file)

		var contador int
		for sc.Scan() {
			contador++
			reg = sc.Text()
		}
	}

	return reg
}
