/*
ESCITALA ESPARTACA

IKER PACHECO % MIGUEL MINANA

INSTRUCCIONES:

	[3] go run escitala.go filas columnas
	[4] go run escitala.go entrada.txt filas columnas
	[5] go run escitala.go entrada.txt salida.txt filas columnas


NOS HEMOS TOMADO LAS LIBERTADES DE DISEÑO DE ...
	- SI EL TEXTO ES MAS CORTO QUE LA MATRIZ, SE RELLENA CON 'X'
	- SI EL TEXTO ES MAS LARGO QUE LA MATRIZ, SE TRUNCA
	- PARA LA ENTRADA POR TECLADO, SE USA '_' COMO CARACTER DE ESCAPE
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

	// Variables para las dimensiones de la matriz (claves del cifrado)
	var filas, columnas int

	// Mapa alfabeto
	alfabeto := map[rune]bool{
		'A': true, 'B': true, 'C': true, 'D': true, 'E': true, 'F': true, 'G': true,
		'H': true, 'I': true, 'J': true, 'K': true, 'L': true, 'M': true, 'N': true,
		'Ñ': true, 'O': true, 'P': true, 'Q': true, 'R': true, 'S': true, 'T': true,
		'U': true, 'V': true, 'W': true, 'X': true, 'Y': true, 'Z': true,
	}

	// SWITCH PARA MANEJAR FUNCIONAMIENTO PROGRAMA

	switch len(os.Args) {
	case 3:
		fmt.Println("Escribe el texto a cifrar terminado en '_', como caracter de escape!")
		// Leo de teclado y escribo en pantalla
		fin = os.Stdin
		fout = os.Stdout
		// CONVERTIMOS LOS PARAMETROS DE STRING A ENTERO FILAS Y COLUMNAS
		filas, err = strconv.Atoi(os.Args[1])
		if err != nil || filas <= 0 {
			fmt.Println("La fila debe ser un número natural")
			os.Exit(1)
		}

		columnas, err = strconv.Atoi(os.Args[2])
		if err != nil || columnas <= 0 {
			fmt.Println("La columna debe ser un número natural")
			os.Exit(1)
		}

	case 4:
		// Caso 2: go run escitala.go entrada.txt filas columnas
		// Lee de archivo, escribe a stdout
		fout = os.Stdout
		// Abrimos el fichero de entrada
		fin, err = os.Open(os.Args[1])
		if err != nil {
			fmt.Println("Error al abrir el archivo de entrada")
			os.Exit(1)
		}
		defer fin.Close()

		filas, err = strconv.Atoi(os.Args[2])
		if err != nil || filas <= 0 {
			fmt.Println("La fila debe ser un número natural")
			os.Exit(1)
		}

		columnas, err = strconv.Atoi(os.Args[3])
		if err != nil || columnas <= 0 {
			fmt.Println("La columna debe ser un número natural")
			os.Exit(1)
		}

	case 5:
		// go run escitala.go entrada.txt salida.txt filas columnas
		// Lee de archivo, escribe a archivo

		// Abrimos el fichero de entrada
		fin, err = os.Open(os.Args[1])
		if err != nil {
			fmt.Println("Error al abrir el archivo de entrada")
			os.Exit(1)
		}
		defer fin.Close()

		// Creamos el fichero de salida
		fout, err = os.Create(os.Args[2])
		if err != nil {
			fmt.Println("Error al crear el archivo de salida")
			os.Exit(1)
		}
		defer fout.Close()

		filas, err = strconv.Atoi(os.Args[3])
		if err != nil || filas <= 0 {
			fmt.Println("La fila debe ser un número natural")
			os.Exit(1)
		}

		columnas, err = strconv.Atoi(os.Args[4])
		if err != nil || columnas <= 0 {
			fmt.Println("La columna debe ser un número natural")
			os.Exit(1)
		}

	default:
		fmt.Println("Numero incorrecto de parámetros :(")
		os.Exit(1)
	}

	// Lectura del texto de entrada y filtrado de letras en mapa

	var cleanTxt []rune // slice de runas validas

	for {
		var c rune
		_, err = fmt.Fscanf(fin, "%c", &c) // lectura carácter por carácter

		if err != nil { // fin de fichero o error
			break
		}

		if c == '_' {
			break
		}

		C := unicode.ToUpper(c) // convertimos a mayúsculas

		// Solo añadimos si está en el alfabeto (ignoramos espacios y otros caracteres)
		if alfabeto[C] {
			cleanTxt = append(cleanTxt, C)
		}
	}

	// Slots totales de la matriz
	totalCeldas := filas * columnas

	// Si el texto es mas pequeño que la matriz, relleno cadena con X
	// antes de meterla a la matriz

	for i := len(cleanTxt); i < totalCeldas; i++ {
		cleanTxt = append(cleanTxt, 'X')
	}

	// Si el texto es mas largo lo cortamos
	if len(cleanTxt) > totalCeldas {
		fmt.Println("Desbordamiento! El texto se corta")
		cleanTxt = cleanTxt[:totalCeldas]
	}

	// Creamos la matriz bidimensional
	matriz := make([][]rune, filas)
	for i := range matriz {
		matriz[i] = make([]rune, columnas)
	}

	// Relleno la matriz por columnas!
	// [1,4,7]
	// [2,5,8]
	// [3,6,9]
	i := 0
	for col := 0; col < columnas; col++ {
		for fila := 0; fila < filas; fila++ {
			if i < len(cleanTxt) {
				matriz[fila][col] = cleanTxt[i]
				i++
			}
		}
	}

	// Leo la matriz por filas!
	// [1,2,3]
	// [4,5,6]
	// [7,8,9]
	for fila := 0; fila < filas; fila++ {
		for col := 0; col < columnas; col++ {
			fmt.Fprintf(fout, "%c", matriz[fila][col])
		}
	}

	// Si escribo en pantalla meto un salto de linea
	if fout == os.Stdout {
		fmt.Fprintln(fout)
	}
}
