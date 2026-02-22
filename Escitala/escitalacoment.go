/*
CIFRADO ESCÍTALA (Scytale Cipher)
==================================

DESCRIPCIÓN DEL ALGORITMO:
--------------------------
La escítala es un método de cifrado por transposición usado en la antigua Grecia.
Consiste en enrollar una tira de cuero alrededor de un bastón (vara) de un grosor específico.

FUNCIONAMIENTO:
--------------
1. Se crea una matriz de FILAS x COLUMNAS (donde FILAS = capacidad de la vara)
2. El texto se escribe en la matriz COLUMNA POR COLUMNA (de arriba a abajo, de izquierda a derecha)
3. El texto cifrado se lee FILA POR FILA (de izquierda a derecha, de arriba a abajo)

EJEMPLO:
--------
Texto: "HOLA MUNDO ES PAÑA"
Filas: 4, Columnas: 4

Escritura por COLUMNAS:     Lectura por FILAS (cifrado):
[H, M, O, A]                HMOA
[O, U, E, Ñ]                OUEÑ
[L, N, S, A]                LNSA
[A, D, P, _]                ADPX

Resultado cifrado: "HMOAOUEÑLNSADPX"

IMPLEMENTACIÓN:
--------------
- Se usa make([][]rune, filas) para crear la matriz bidimensional
- Se rellena por columnas con los caracteres del texto
- Se lee por filas para obtener el texto cifrado
- Si sobran espacios, se rellenan con 'X'

REQUISITOS:
-----------
• Leer el texto en claro de la entrada estándar o archivo
• Escribir el texto cifrado en la salida estándar o archivo
• El texto estará compuesto por las letras del alfabeto castellano (A-Z incluyendo Ñ)
• La entrada puede ser minúsculas o mayúsculas, la salida será siempre en MAYÚSCULAS
• Los espacios serán ignorados
• Aceptar parámetros enteros en la línea de comandos (filas y columnas)
• [opcional] Aceptar parámetros para elegir ficheros de entrada y salida

USO:
----
go run escitala.go filas columnas                          # stdin/stdout
go run escitala.go entrada.txt filas columnas              # archivo entrada, stdout
go run escitala.go entrada.txt salida.txt filas columnas   # archivo entrada y salida
*/

package main

import (
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func main() {
	// Variables para manejo de archivos
	var fin *os.File  // fichero de entrada (por defecto stdin)
	var fout *os.File // fichero de salida (por defecto stdout)
	var err error     // receptor de errores

	// Variables para las dimensiones de la matriz (clave del cifrado)
	var filas, columnas int

	// Alfabeto castellano con el que vamos a trabajar (A-Z + Ñ)
	alfabeto := map[rune]bool{
		'A': true, 'B': true, 'C': true, 'D': true, 'E': true, 'F': true, 'G': true,
		'H': true, 'I': true, 'J': true, 'K': true, 'L': true, 'M': true, 'N': true,
		'Ñ': true, 'O': true, 'P': true, 'Q': true, 'R': true, 'S': true, 'T': true,
		'U': true, 'V': true, 'W': true, 'X': true, 'Y': true, 'Z': true,
	}

	// Por defecto, usamos stdin y stdout
	fin = os.Stdin
	fout = os.Stdout

	// PROCESAMIENTO DE ARGUMENTOS DE LÍNEA DE COMANDOS
	// =================================================

	if len(os.Args) == 3 {
		// Caso 1: go run escitala.go filas columnas
		// Lee de stdin, escribe a stdout

		filas, err = strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error: el número de filas debe ser un número entero")
			os.Exit(1)
		}

		columnas, err = strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error: el número de columnas debe ser un número entero")
			os.Exit(1)
		}

	} else if len(os.Args) == 4 {
		// Caso 2: go run escitala.go entrada.txt filas columnas
		// Lee de archivo, escribe a stdout

		// Abrimos el fichero de entrada
		fin, err = os.Open(os.Args[1])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error al abrir el archivo de entrada: %v\n", err)
			os.Exit(1)
		}
		defer fin.Close()

		filas, err = strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error: el número de filas debe ser un número entero")
			os.Exit(1)
		}

		columnas, err = strconv.Atoi(os.Args[3])
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error: el número de columnas debe ser un número entero")
			os.Exit(1)
		}

	} else if len(os.Args) == 5 {
		// Caso 3: go run escitala.go entrada.txt salida.txt filas columnas
		// Lee de archivo, escribe a archivo

		// Abrimos el fichero de entrada
		fin, err = os.Open(os.Args[1])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error al abrir el archivo de entrada: %v\n", err)
			os.Exit(1)
		}
		defer fin.Close()

		// Creamos el fichero de salida
		fout, err = os.Create(os.Args[2])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error al crear el archivo de salida: %v\n", err)
			os.Exit(1)
		}
		defer fout.Close()

		filas, err = strconv.Atoi(os.Args[3])
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error: el número de filas debe ser un número entero")
			os.Exit(1)
		}

		columnas, err = strconv.Atoi(os.Args[4])
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error: el número de columnas debe ser un número entero")
			os.Exit(1)
		}

	} else {
		// Error: número incorrecto de parámetros
		fmt.Fprintln(os.Stderr, "Uso: escitala [entrada.txt] [salida.txt] filas columnas")
		fmt.Fprintln(os.Stderr, "")
		fmt.Fprintln(os.Stderr, "Opciones:")
		fmt.Fprintln(os.Stderr, "  escitala filas columnas")
		fmt.Fprintln(os.Stderr, "    Lee de stdin, escribe a stdout")
		fmt.Fprintln(os.Stderr, "")
		fmt.Fprintln(os.Stderr, "  escitala entrada.txt filas columnas")
		fmt.Fprintln(os.Stderr, "    Lee de archivo, escribe a stdout")
		fmt.Fprintln(os.Stderr, "")
		fmt.Fprintln(os.Stderr, "  escitala entrada.txt salida.txt filas columnas")
		fmt.Fprintln(os.Stderr, "    Lee de archivo, escribe a archivo")
		os.Exit(1)
	}

	// Validación de parámetros
	if filas <= 0 || columnas <= 0 {
		fmt.Fprintln(os.Stderr, "Error: filas y columnas deben ser números positivos")
		os.Exit(1)
	}

	// LECTURA DEL TEXTO DE ENTRADA
	// =============================
	// Leemos todo el texto y filtramos solo las letras válidas del alfabeto

	var textoLimpio []rune // slice para almacenar solo las letras válidas

	for {
		var c rune
		_, err = fmt.Fscanf(fin, "%c", &c) // lectura carácter por carácter

		if err != nil { // fin de fichero o error
			break
		}

		C := unicode.ToUpper(c) // convertimos a mayúsculas

		// Solo añadimos si está en el alfabeto (ignoramos espacios y otros caracteres)
		if alfabeto[C] {
			textoLimpio = append(textoLimpio, C)
		}
	}

	// CREACIÓN Y LLENADO DE LA MATRIZ
	// ================================
	// Calculamos cuántos caracteres necesitamos en total
	totalCeldas := filas * columnas

	// Si el texto es más corto que la matriz, rellenamos con 'X'
	for len(textoLimpio) < totalCeldas {
		textoLimpio = append(textoLimpio, 'X')
	}

	// Si el texto es más largo, lo truncamos (o podríamos procesarlo en bloques)
	if len(textoLimpio) > totalCeldas {
		textoLimpio = textoLimpio[:totalCeldas]
	}

	// Creamos la matriz bidimensional
	matriz := make([][]rune, filas)
	for i := range matriz {
		matriz[i] = make([]rune, columnas)
	}

	// LLENAMOS LA MATRIZ POR COLUMNAS
	// ================================
	// Recorremos columna por columna, de arriba a abajo
	indice := 0
	for col := 0; col < columnas; col++ {
		for fila := 0; fila < filas; fila++ {
			if indice < len(textoLimpio) {
				matriz[fila][col] = textoLimpio[indice]
				indice++
			}
		}
	}

	// LECTURA DE LA MATRIZ POR FILAS (CIFRADO)
	// =========================================
	// Recorremos fila por fila, de izquierda a derecha
	for fila := 0; fila < filas; fila++ {
		for col := 0; col < columnas; col++ {
			fmt.Fprintf(fout, "%c", matriz[fila][col])
		}
	}

	// Añadimos un salto de línea al final si estamos escribiendo a stdout
	if fout == os.Stdout {
		fmt.Fprintln(fout)
	}
}
