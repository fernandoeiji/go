// declarar pacote principal
package main

//importar função fmt
import "fmt"

//declaração da variável const
const ebuliçãoF float64 = 212.0

//função principal
func main() {

	//var tempF float64 = ebuliçãoF // variavel do valor da temperatura em Fahrenheit
	tempF := ebuliçãoF // variavel do valor da temperatura em Fahrenheit
	//	var tempC float64 = (tempF - 32) * 5 / 9
	tempC := (tempF - 32) * 5 / 9

	//resultado
	fmt.Println("A temperatura da ebulição da água em °C = ", tempC)

}
