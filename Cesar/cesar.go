/*
CIFRADO DEL CESAR

IKER PACHECO & MIGUEL MINANA

Packages empleados:
- fmt -> Mensajes por consola :D
- os -> Punteros a in/out
- strconv -> ParseInt -> strconv.Atoi()
- unicode -> Pasar las runas a mayúscula

INSTRUCCIONES:

	[1] go run cesar.go
	[2] go run cesar.go 3
	[3] go run cesar.go entrada.txt salida.txt
	[4] go run cesar.go 3 entrada.txt salida.txt

*/

package main

import (
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func main() {

	var fin, fout *os.File // punteros a ficheros de entrada y salida
	var err error          // receptor de error

	var key int // clave de desplazamiento

	// alfabeto con el que vamos a trabajar
	alfabeto := map[rune]int{'A': 0, 'B': 1, 'C': 2, 'D': 3, 'E': 4, 'F': 5, 'G': 6, 'H': 7, 'I': 8, 'J': 9,
		'K': 10, 'L': 11, 'M': 12, 'N': 13, 'Ñ': 14, 'O': 15, 'P': 16, 'Q': 17, 'R': 18, 'S': 19,
		'T': 20, 'U': 21, 'V': 22, 'W': 23, 'X': 24, 'Y': 25, 'Z': 26}

	// APLICAMOS LA LOGICA DE SWITCH DE GO
	// DETERMINAMOS QUE HACER SEGUN LOS PARAMETROS PASADOS POR CONSOLA

	switch len(os.Args) {
	case 1: // no hay parámetros, usamos entrada (teclado) y salida estándar (pantalla)
		key = 3
		fin = os.Stdin
		fout = os.Stdout
	case 2: // tenemos la clave de desplazamiento
		key, err = strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Println("La clave debe ser integer")
			os.Exit(1)
		}
		fin = os.Stdin
		fout = os.Stdout
	case 3: // tenemos la clave por defecto (3), el archivo de entrada y el archivo de salida
		key = 3
		fin, err = os.Open(os.Args[1])
		if err != nil {
			panic(err)
		}
		defer fin.Close()
		fout, err = os.Create(os.Args[2])
		if err != nil {
			panic(err)
		}
		defer fout.Close()
	case 4: // tenemos la clave de desplazamiento, el archivo de entrada y el archivo de salida

		// CONVERTIMOS EL PARAMETRO (KEY) DE STRING A ENTERO
		key, err = strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Println("La clave debe ser integer")
			os.Exit(1)
		}

		// ABRIMOS EL FICHERO DE ENTRADA
		fin, err = os.Open(os.Args[2])
		if err != nil {
			panic(err)
		}
		defer fin.Close()

		// CREAMOS EL FICHERO DE SALIDA
		fout, err = os.Create(os.Args[3])
		if err != nil {
			panic(err)
		}
		defer fout.Close()
	default:
		fmt.Println("Numero incorrecto de parámetros :(")
		os.Exit(1)
	}

	for { // bucle infinito, rompo con BREAK
		var c rune // carácter a leer: Runa es un tipo de dato de Go que representa un carácter Unicode

		_, err = fmt.Fscanf(fin, "%c", &c) // lectura de la entrada, pillo un caracter "%c" y lo paso a valor en Runa c

		if err != nil { // si hay error (fin de fichero)
			break //salimos del bucle
		}

		C := unicode.ToUpper(c) // pasamos a mayúsculas

		posicion, ok := alfabeto[C] // comprobamos que está en el alfabeto y obtenemos su posición
		if ok {                     // Si en alfabeto -> Cifro!

			// Desplazamiento del cesar
			// Sumo clave y aplico módulo
			nuevaPos := (posicion + key) % 27

			// Si la posición es negativa (key negativa), le sumo 27 para mantener en [0,26]
			if nuevaPos < 0 {
				nuevaPos += 27
			}

			// Busco en mapa alfabeto letra de la nueva posicion
			for letra, pos := range alfabeto {
				if pos == nuevaPos {
					fmt.Fprintf(fout, "%c", letra)
					break
				}
			}
		}
	}

}
