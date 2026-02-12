/*
Ejemplo 1

Este programa copia de la entrada a la salida carácter a carácter,
restringiéndose a un alfabeto limitado y pasando a mayúsculas.
Permite leer de la entrada y salida estándar o usar ficheros


-lectura y escritura
-entrada y salida estándar
-ficheros
-parámetros en línea de comandos (os.Args)

ENUNCIADO DEL CESAR:
• Leer el texto en claro de la entrada estándar
• Escribir el texto cifrado en la salida estándar
• El texto estará compuesto por las letras del alfabeto castellano (de la A a la Z incluyendo la Ñ).
• La entrada podrá ser en minúsculas o mayúsculas, la salida será siempre en mayúsculas. Los
espacios serán ignorados.
• Aceptar un parámetro entero en la línea de comandos que sea la clave (desplazamiento), puede aceptar negativos y por defecto es 3.
• [opcional] Aceptar parámetros en la línea de comandos que permitan elegir un fichero de
entrada y otro de salida en lugar de la entrada y salida estándar.

INSTRUCCIONES:
	- PEDIMOS PARAMETRO DE ENTRADA PODIENDO SER: [1] SCAN y por DEFECTO (3 CORRIDAS), [2] SCAN, CORRIDA, [4] SCAN, CORRIDA, IN, OUT

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

	var fin *os.File  // fichero de entrada
	var fout *os.File // fichero de salida
	var err error     // receptor de error

	var key int // clave de desplazamiento

	// alfabeto con el que vamos a trabajar
	alfabeto := map[rune]int{'A': 0, 'B': 1, 'C': 2, 'D': 3, 'E': 4, 'F': 5, 'G': 6, 'H': 7, 'I': 8, 'J': 9,
		'K': 10, 'L': 11, 'M': 12, 'N': 13, 'Ñ': 14, 'O': 15, 'P': 16, 'Q': 17, 'R': 18, 'S': 19,
		'T': 20, 'U': 21, 'V': 22, 'W': 23, 'X': 24, 'Y': 25, 'Z': 26}

	if len(os.Args) == 1 { // no hay parámetros, usamos entrada (teclado) y salida estándar (pantalla)

		key = 3 // por defecto es 3

		fin = os.Stdin
		fout = os.Stdout

	} else if len(os.Args) == 2 { // tenemos la clave de desplazamiento

		// Convertimos el parámetro string a entero
		key, err = strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Println("Error: la clave debe ser un número entero")
			os.Exit(1)
		}

		fin = os.Stdin
		fout = os.Stdout

	} else if len(os.Args) == 3 { // tenemos clave como 3 (default), archivo de entrada y archivo de salida

		key = 3 // por defecto es 3

		// Abrimos el fichero de entrada
		fin, err = os.Open(os.Args[1])
		if err != nil {
			panic(err)
		}
		defer fin.Close()

		// Creamos el fichero de salida
		fout, err = os.Create(os.Args[2])
		if err != nil {
			panic(err)
		}
		defer fout.Close()

	} else if len(os.Args) == 4 { // tenemos clave, archivo de entrada y archivo de salida

		// Convertimos el parámetro string a entero
		key, err = strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Println("Error: la clave debe ser un número entero")
			os.Exit(1)
		}

		// Abrimos el fichero de entrada
		fin, err = os.Open(os.Args[2])
		if err != nil {
			panic(err)
		}
		defer fin.Close()

		// Creamos el fichero de salida
		fout, err = os.Create(os.Args[3])
		if err != nil {
			panic(err)
		}
		defer fout.Close()

	} else if len(os.Args) == 4 { // tenemos clave, archivo de entrada y archivo de salida

		// Convertimos el parámetro string a entero
		key, err = strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Println("Error: la clave debe ser un número entero")
			os.Exit(1)
		}

		// Abrimos el fichero de entrada
		fin, err = os.Open(os.Args[2])
		if err != nil {
			panic(err)
		}
		defer fin.Close()

		// Creamos el fichero de salida
		fout, err = os.Create(os.Args[3])
		if err != nil {
			panic(err)
		}
		defer fout.Close()

	} else { // error de parámetros
		fmt.Println("Uso: cesar [clave] [archivo_entrada] [archivo_salida]")
		fmt.Println("  - Sin parámetros: lee de stdin, escribe a stdout, clave=3")
		fmt.Println("  - 1 parámetro: clave personalizada, stdin/stdout")
		fmt.Println("  - 3 parámetros: clave, archivo entrada, archivo salida")
		os.Exit(1)
	}

	for { // bucle infinito
		var c rune // carácter a leer

		_, err = fmt.Fscanf(fin, "%c", &c) // lectura de la entrada

		if err != nil { // si hay error (fin de fichero)
			break //salimos del bucle
		}

		C := unicode.ToUpper(c) // pasamos a mayúsculas

		posicion, ok := alfabeto[C] // comprobamos que está en el alfabeto y obtenemos su posición
		if ok {                     // si está en el alfabeto, aplicamos el cifrado César

			// Aplicamos el desplazamiento (cifrado César)
			// Sumamos la clave y usamos módulo para que sea circular
			nuevaPosicion := (posicion + key) % 27

			// Si la posición es negativa (clave negativa), ajustamos
			if nuevaPosicion < 0 {
				nuevaPosicion += 27
			}

			// Buscamos la letra que corresponde a la nueva posición
			for letra, pos := range alfabeto {
				if pos == nuevaPosicion {
					fmt.Fprintf(fout, "%c", letra)
					break
				}
			}
		}
	}

}
