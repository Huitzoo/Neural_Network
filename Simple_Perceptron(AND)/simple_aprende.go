package main


import (
  "fmt"
  "time"
	"math/rand"
)

type neurona struct{
  pesos []float32
  pesos_a []float32
  entradas []float32
  umbral float32
  umbral_a float32
  sum float32
  valor float32
  tasa float32
}

func (n *neurona) calcular() (int,float32) {
  var i int = 0
  var sum float32 = 0

  for i < len(n.pesos) {
    sum = n.entradas[i]*n.pesos[i] + sum
    i++
  }
  sum = sum + n.umbral
  val := sigmoide(sum)
  return val, sum
}
func (n *neurona) aprender() {
  if  len(n.pesos_a) != 0{
    n.pesos[0] = n.pesos_a[0] + n.tasa * ((n.valor-n.sum)*n.entradas[0])
    n.pesos[1] = n.pesos_a[1] + n.tasa * ((n.valor-n.sum)*n.entradas[1])
    n.umbral = n.umbral_a + n.tasa * (n.valor-n.sum)
  }
}

func sigmoide(res float32) int {
  if res > 0 {
    return 1
  }else {
    return 0
  }
  return 2
}

//Error = lo que necesitamos - lo que tenemos
//peso = pesoanterior + tasa aprendizaje *error *entradas
//tasa de aprendizaje

func main() {
  var res, n_neuronas, nuevo int
  r := rand.New(rand.NewSource(time.Now().UnixNano()))
  n := neurona{pesos:[]float32{0,0},pesos_a:[]float32{0,0},entradas:[]float32{0,0}, tasa:1}
  n_neuronas = 0
  for {
      nuevo = 0
      if n.pesos[0]==0 && n.pesos[1] == 0{
        n.pesos[0] = (r.Float32() - r.Float32())
        n.pesos[1] = (r.Float32() - r.Float32())
        n.umbral = (r.Float32() - r.Float32())
        n.pesos_a = n.pesos
        n.umbral_a = n.umbral
      }
      fmt.Println("----------------------------------------------------")
      fmt.Println("peso1: ",n.pesos[0], "peso2: ", n.pesos[1],"umbral: ",n.umbral)
      n.entradas[0] = 1
      n.entradas[1] = 1
      res,n.sum = n.calcular()
      fmt.Println("entrada1: 1 entrada2: 1 Res: ", res)
      if res != 1{
          n.valor = 1
          n.aprender()
      }else{
        nuevo = nuevo + 1
      }
      n.entradas[0] = 1
      n.entradas[1] = 0
      res,n.sum = n.calcular()
      fmt.Println("entrada1: 1 entrada2: 0 Res: ", res)
      if res != 0{
          n.valor = 0
          n.aprender()
      }else{
          nuevo = nuevo + 1
      }
      n.entradas[0] = 0
      n.entradas[1] = 1
      res,n.sum = n.calcular()
      fmt.Println("entrada1: 0 entrada2: 1 Res: ", res)
      if res != 0{
        n.valor = 0
        n.aprender()
      }else{
        nuevo = nuevo + 1
      }
      n.entradas[0] = 0
      n.entradas[1] = 0
      res,n.sum = n.calcular()
      fmt.Println("entrada1: 0 entrada2: 0 Res: ", res)
      if res != 0{
        n.valor = 0
        n.aprender()
      }else{
        nuevo = nuevo + 1
      }
      fmt.Println("nuevo: ", nuevo)
      if nuevo == 4 {
        fmt.Println("Neuronas: ",n_neuronas)
        break
      }else{
      }
      n.pesos_a = n.pesos
      n.umbral_a = n.umbral
      n_neuronas = n_neuronas + 1
  }
}
