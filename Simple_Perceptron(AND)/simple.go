package main

import (
	"fmt"
	"time"
	"math/rand"
)
/*
	AND E1 E2
	1		1		1
	0		1		0
	0		0		1
	0		0		0
*/

func calcular(entrada1 float64,entrada2 float64,peso1 float64,peso2 float64,umbral float64) int{
	valor := umbral + entrada1*peso1 + entrada2*peso2
	valor_1 := sigmoide(valor)
	return valor_1
}

func sigmoide(res float64) int {
	if res > 0 {
		return 1
	}else {
		return 0
	}
	return 2
}

func main() {
		var peso1, peso2, umbral float64
		var res, neuronas, nuevo int
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		neuronas = 0

		for {
				nuevo = 0
				peso1 = (r.Float64() - r.Float64())
				peso2 = (r.Float64() - r.Float64())
				umbral = (r.Float64() - r.Float64())
				fmt.Println("----------------------------------------------------")
				fmt.Println("peso1: ",peso1, "peso1: ", peso2,"umbral: ",umbral)
				res = calcular(1,1, peso1, peso2, umbral)
				fmt.Println("entrada1: 1 entrada2: 1 Res: ", res)
				if res == 1{
					nuevo = nuevo + 1
				}
				res = calcular(1,0, peso1, peso2, umbral)
				fmt.Println("entrada1: 1 entrada2: 0 Res: ", res)

				if res == 0{
					nuevo = nuevo + 1
				}
				res = calcular(0,1, peso1, peso2, umbral)
				fmt.Println("entrada1: 0 entrada2: 1 Res: ", res)
				if res == 0{
						nuevo = nuevo + 1
				}
				res = calcular(0,0, peso1, peso2, umbral)
				fmt.Println("entrada1: 0 entrada2: 0 Res: ", res)
				if res == 0{
					nuevo = nuevo + 1
				}
				fmt.Println("nuevo: ", nuevo)
				if nuevo == 4 {
					fmt.Println("Neuronas: ",neuronas)
					break
				}
				neuronas = neuronas + 1
		}
}
