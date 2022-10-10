package main

import (
	"fmt"
	"math"
	"sync"
)

func child(wg *sync.WaitGroup) {
        defer wg.Done()
}

func main() {
        var peso, altura, imc float64
        var wg sync.WaitGroup
        wg.Add(1)
        go child(&wg)

        fmt.Println("Digite seu peso e sua altura : ")
        fmt.Scanf("%f", &peso)
        fmt.Scanf("%f", &altura)

        fmt.Println("Você pesa", peso, "kilos")
        wg.Wait()
        fmt.Println("Você tem", altura, "metros")

        imc = peso / math.Pow(altura, 2)

        fmt.Println("Seu IMC é: ", imc)
        fmt.Print("Sua categoria de risco é: ")

        if imc < float64(18.5) {
                fmt.Println("Abaixo do peso")
        } else if imc < 25 {
                fmt.Println("Peso normal")
        } else if imc < 30 {
                fmt.Println("Acima do Peso")
         } else {
                fmt.Println("Obeso")
        }

        pesoNormal := 25 * math.Pow(altura, 2)
        delta := peso - pesoNormal

        fmt.Printf("O peso normal para sua altura é de : %0.2v kilograms.\n", pesoNormal)

        if (delta > 0) && (imc > 30) {
                fmt.Printf("Você precisa reduzir %0.2v kilos.\n", math.Abs(delta))
        }

        if (delta < 0) && (imc < float64(18.5)) {
                fmt.Printf("Você precisa aumentar %0.2v kilos.\n", math.Abs(delta))
        }
}