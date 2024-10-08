package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
	"sort"
)

func main(){
	var slice = []int{}
	fmt.Println("Bienvenido al creador de y organizador de slices")
        for {
                Lector := bufio.NewScanner(os.Stdin)
				
                fmt.Print("Ingrese el numero que desee añadir, o escriba 'salir' para terminar el programa y visualizar el slice : ")
                Lector.Scan()

                Texto := Lector.Text()

                if strings.ToLower(Texto) == "salir" {
                        fmt.Println("Visualizacion")
                        break
                }

                Numero, err := strconv.Atoi(Texto)

				if err != nil {
					fmt.Println("Error no se permiten letras, a excepcion de la palabra 'salir'", err)
					continue
				} else{
					slice = append(slice, Numero)
					fmt.Println(Numero, " Añadido")
				}

        }

		var visulizar=1
				fmt.Println("¿Como desea visulizar el slice?\n1) de menor a mayor\n2) de mayor a menor")
				fmt.Scan(&visulizar)

				visulizarconvertido := strconv.Itoa(visulizar)

				
				if visulizarconvertido == "1"{
					sort.Ints(slice)
					fmt.Println("Ordenados de menor a mayor: ", slice)
				}else if visulizarconvertido == "2"{
					sort.Sort(sort.Reverse(sort.IntSlice(slice)))
					fmt.Println("Ordenados de mayor a menor: ",slice)}
			

				
				

}