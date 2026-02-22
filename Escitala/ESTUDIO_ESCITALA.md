# 🔐 Cifrado Escítala en Go — Guía de Estudio Completa

> **Objetivo:** Entender completamente el programa `escitala.go` desde cero para el examen.

---

## 📚 Índice

1. [¿Qué es la Escítala?](#1--qué-es-la-escítala)
2. [Conceptos Básicos de Go](#2--conceptos-básicos-de-go)
3. [Estructura del Programa](#3--estructura-del-programa)
4. [Explicación Línea a Línea](#4--explicación-línea-a-línea)
5. [El Algoritmo Paso a Paso](#5--el-algoritmo-paso-a-paso)
6. [Tipos de Datos Clave](#6--tipos-de-datos-clave)
7. [Flujo Completo con Ejemplo](#7--flujo-completo-con-ejemplo)
8. [Comandos de Ejecución](#8--comandos-de-ejecución)
9. [Posibles Preguntas de Examen](#9--posibles-preguntas-de-examen)
10. [Resumen Rápido](#10--resumen-rápido)

---

## 1. 🏺 ¿Qué es la Escítala?

La **escítala** es uno de los cifrados más antiguos de la historia, usado por los espartanos en la Antigua Grecia.

### ¿Cómo funciona físicamente?

```
┌─────────────────────────────────────────┐
│  Se enrolla una tira de cuero en un     │
│  bastón de grosor específico            │
│                                         │
│  El mensaje se escribe a lo largo       │
│  del bastón → al desenrollar la tira,   │
│  el texto queda desordenado = CIFRADO   │
└─────────────────────────────────────────┘
```

### ¿Cómo lo simulamos con una matriz?

| Concepto físico | Concepto en el programa |
|----------------|------------------------|
| Grosor del bastón | Número de **filas** |
| Longitud del bastón | Número de **columnas** |
| Enrollar la tira | Llenar la matriz por **columnas** |
| Leer el bastón | Leer la matriz por **filas** |

---

## 2. 🧱 Conceptos Básicos de Go

Antes de entrar al código, estos son los conceptos de Go que necesitas saber:

### `package main`
```go
package main
```
- Todo programa en Go pertenece a un **paquete**
- El paquete `main` indica que este archivo es un **programa ejecutable**
- Sin `package main`, Go no sabe que tiene que generar un ejecutable

### `import`
```go
import (
    "fmt"
    "os"
    "strconv"
    "unicode"
)
```
- Carga **librerías estándar** de Go
- `fmt` → imprimir y leer datos (`fmt.Printf`, `fmt.Fscanf`...)
- `os` → interactuar con el sistema operativo (archivos, argumentos...)
- `strconv` → convertir tipos (`string` ↔ `int`)
- `unicode` → manipular caracteres (`ToUpper`, `IsLetter`...)

### `func main()`
```go
func main() {
    // El programa empieza aquí
}
```
- `func` = palabra clave para definir una función
- `main()` = función especial que Go ejecuta **automáticamente al iniciar**
- Es el **punto de entrada** del programa

### Variables en Go

```go
// Forma 1: declaración explícita
var nombre string
var número int

// Forma 2: declaración + asignación (más común)
nombre := "Hola"   // Go infiere el tipo automáticamente
número := 42
```

### El operador `:=`
- Declara la variable **Y** le asigna un valor en una sola línea
- Solo se puede usar **dentro de funciones**
- Go detecta automáticamente el tipo

### Errores en Go
```go
var err error  // tipo especial para errores

resultado, err = algunaFuncion()
if err != nil {  // nil = "nada/vacío" - si hay error, err no es nil
    // manejar el error
}
```
- En Go, los errores **se devuelven como valores**, no se lanzan como excepciones
- Si `err == nil` → todo fue bien ✅
- Si `err != nil` → algo falló ❌

---

## 3. 🗺️ Estructura del Programa

```
escitala.go
│
├── 1. Declarar variables (archivos, dimensiones)
│
├── 2. Definir el alfabeto válido (mapa)
│
├── 3. Procesar argumentos de línea de comandos
│   ├── 3 args → stdin/stdout
│   ├── 4 args → archivo entrada / stdout
│   └── 5 args → archivo entrada / archivo salida
│
├── 4. Leer el texto carácter a carácter
│   └── Filtrar: solo letras válidas en MAYÚSCULAS
│
├── 5. Ajustar el texto a la matriz
│   ├── Si falta texto → rellenar con 'X'
│   └── Si sobra texto → truncar
│
├── 6. Crear la matriz [filas][columnas]
│
├── 7. Llenar la matriz POR COLUMNAS
│
└── 8. Leer la matriz POR FILAS → imprimir cifrado
```

---

## 4. 📝 Explicación Línea a Línea

### 4.1 Variables de archivos

```go
var fin *os.File  // fichero de entrada (por defecto teclado/stdin)
var fout *os.File // fichero de salida (por defecto pantalla/stdout)
var err error     // receptor de errores
```

| Variable | Tipo | ¿Qué es? |
|----------|------|----------|
| `fin` | `*os.File` | Puntero al archivo de entrada. El `*` indica que es un **puntero** (dirección de memoria del archivo) |
| `fout` | `*os.File` | Puntero al archivo de salida |
| `err` | `error` | Almacena cualquier error que ocurra |

> 💡 **¿Qué es un puntero?** En lugar de guardar el archivo completo, guardamos su **dirección de memoria** (como la dirección de una casa en lugar de la casa entera). Es más eficiente.

### 4.2 El mapa del alfabeto

```go
alfabeto := map[rune]bool{
    'A': true, 'B': true, 'C': true, 'D': true, 'E': true,
    // ... todas las letras castellanas + Ñ
    'Ñ': true,
}
```

- `map[rune]bool` → diccionario donde la **clave** es un carácter (`rune`) y el **valor** es verdadero/falso (`bool`)
- `rune` → tipo de Go para representar **cualquier carácter Unicode** (incluye Ñ, acentos, etc.)
- ¿Para qué sirve? → Para filtrar el texto y quedarnos **solo con letras válidas**

```go
// Ejemplo de uso:
alfabeto['A']  // → true  ✅
alfabeto['1']  // → false ❌
alfabeto[' ']  // → false ❌ (los espacios se ignoran)
```

### 4.3 Valores por defecto

```go
fin = os.Stdin   // leer del teclado
fout = os.Stdout // escribir en pantalla
```

- `os.Stdin` = **entrada estándar** = el teclado
- `os.Stdout` = **salida estándar** = la pantalla
- Son las entradas/salidas por defecto si el usuario no proporciona archivos

### 4.4 Procesamiento de argumentos — `os.Args`

```go
if len(os.Args) == 3 {
    // go run escitala.go filas columnas
}
```

- `os.Args` es un **slice (lista)** con todos los argumentos del programa
- `os.Args[0]` = nombre del programa (siempre)
- `os.Args[1]`, `os.Args[2]`... = argumentos del usuario
- `len()` = función que devuelve el número de elementos

```
Comando:  go run escitala.go entrada.txt 4 3
           ─────────────────  ──────────  ─ ─
                Args[0]         Args[1]  [2][3]
           len(os.Args) == 4
```

#### Los 3 casos posibles:

```
┌──────────────────────────────────────────────────────────┐
│ len == 3 │ escitala filas columnas                       │
│          │ Lee del TECLADO, escribe en PANTALLA          │
├──────────┼───────────────────────────────────────────────┤
│ len == 4 │ escitala entrada.txt filas columnas           │
│          │ Lee de ARCHIVO, escribe en PANTALLA           │
├──────────┼───────────────────────────────────────────────┤
│ len == 5 │ escitala entrada.txt salida.txt filas columnas│
│          │ Lee de ARCHIVO, escribe en ARCHIVO            │
└──────────┴───────────────────────────────────────────────┘
```

### 4.5 Abrir/Crear archivos

```go
// Abrir para LEER (el archivo debe existir)
fin, err = os.Open(os.Args[1])
if err != nil {
    fmt.Fprintf(os.Stderr, "Error: %v\n", err)
    os.Exit(1)
}
defer fin.Close()
```

```go
// Crear para ESCRIBIR (lo crea si no existe, lo sobreescribe si existe)
fout, err = os.Create(os.Args[2])
```

| Función | Uso |
|---------|-----|
| `os.Open(ruta)` | Abre un archivo **solo lectura** |
| `os.Create(ruta)` | Crea o sobreescribe un archivo para **escritura** |
| `defer x.Close()` | Cierra el archivo cuando la función termina |

> 💡 **`defer`**: Pospone la ejecución hasta el final de la función. Se usa para asegurarse de que el archivo siempre se cierre, aunque haya un error. Es como decir: *"acuérdate de cerrar esto cuando termines"*.

### 4.6 `strconv.Atoi` — Convertir texto a número

```go
filas, err = strconv.Atoi(os.Args[1])
if err != nil {
    fmt.Fprintln(os.Stderr, "Error: filas debe ser un entero")
    os.Exit(1)
}
```

- `os.Args` siempre devuelve **strings** (texto)
- Necesitamos **convertir** `"4"` → `4` para hacer cálculos
- `Atoi` = **A**SCII **to** **i**nteger
- Si el usuario escribe `"abc"` en lugar de un número, `err != nil` y mostramos error

### 4.7 Validación de parámetros

```go
if filas <= 0 || columnas <= 0 {
    fmt.Fprintln(os.Stderr, "Error: filas y columnas deben ser positivos")
    os.Exit(1)
}
```

- `||` = operador **OR** lógico
- `<=` = menor o igual
- No tiene sentido una matriz con 0 o menos filas/columnas

### 4.8 Lectura del texto — Bucle infinito con `break`

```go
var textoLimpio []rune  // slice vacío de runas

for {                              // bucle infinito
    var c rune
    _, err = fmt.Fscanf(fin, "%c", &c)  // leer un carácter

    if err != nil {   // si hay error (fin de archivo)
        break         // salir del bucle
    }

    C := unicode.ToUpper(c)   // convertir a mayúscula

    if alfabeto[C] {           // si es una letra válida
        textoLimpio = append(textoLimpio, C)  // añadirla
    }
}
```

Desglose detallado:

| Parte | Explicación |
|-------|-------------|
| `[]rune` | Slice (lista dinámica) de caracteres |
| `for { }` | Bucle infinito — se repite hasta que algo lo rompa |
| `fmt.Fscanf(fin, "%c", &c)` | Lee **1 carácter** de `fin` y lo guarda en `c` |
| `&c` | `&` = "dame la dirección de `c`" (puntero). Fscanf necesita saber dónde guardar el dato |
| `_` | Descarta el primer valor de retorno (bytes leídos) que no nos interesa |
| `break` | **Sale del bucle** inmediatamente |
| `unicode.ToUpper(c)` | Convierte el carácter a mayúscula |
| `append(slice, elem)` | Añade un elemento al final del slice |

### 4.9 Rellenar y truncar el texto

```go
// Si el texto es MÁS CORTO que la matriz → rellenar con X
for len(textoLimpio) < totalCeldas {
    textoLimpio = append(textoLimpio, 'X')
}

// Si el texto es MÁS LARGO que la matriz → truncar
if len(textoLimpio) > totalCeldas {
    textoLimpio = textoLimpio[:totalCeldas]
}
```

- `totalCeldas = filas * columnas` → número total de caracteres que caben en la matriz
- `slice[:n]` → devuelve los primeros `n` elementos del slice (como cortar con tijeras)
- La `X` es el **carácter de relleno** estándar en criptografía

### 4.10 Crear la matriz bidimensional

```go
// Crear un slice de "filas" elementos, cada uno vacío
matriz := make([][]rune, filas)

// Para cada fila, crear un slice de "columnas" runas
for i := range matriz {
    matriz[i] = make([]rune, columnas)
}
```

- `make(tipo, tamaño)` → crea un slice con tamaño específico (inicializado con valores cero)
- `[][]rune` → "slice de slices de rune" = **matriz 2D**
- `range matriz` → itera sobre los índices del slice (equivale a `i = 0; i < len(matriz); i++`)

```
make([][]rune, 4) crea:
[ nil, nil, nil, nil ]
  ↓
Después del bucle:
[ [_,_,_], [_,_,_], [_,_,_], [_,_,_] ]
  fila 0    fila 1    fila 2    fila 3
```

### 4.11 Llenar la matriz POR COLUMNAS (el cifrado)

```go
indice := 0
for col := 0; col < columnas; col++ {      // columna exterior
    for fila := 0; fila < filas; fila++ {  // fila interior
        if indice < len(textoLimpio) {
            matriz[fila][col] = textoLimpio[indice]
            indice++
        }
    }
}
```

> ⚠️ **CRUCIAL para el examen**: El bucle exterior es por **columnas** y el interior por **filas**. Esto es lo que hace el cifrado: el texto se escribe de arriba a abajo en cada columna.

```
Texto: "HOLAM"  →  filas=2, columnas=3

col=0:  matriz[0][0]='H',  matriz[1][0]='O'
col=1:  matriz[0][1]='L',  matriz[1][1]='A'
col=2:  matriz[0][2]='M',  matriz[1][2]='X'  (relleno)

Resultado:
       col0  col1  col2
fila0: [ H  ,  L  ,  M ]
fila1: [ O  ,  A  ,  X ]
```

### 4.12 Leer la matriz POR FILAS (obtener el cifrado)

```go
for fila := 0; fila < filas; fila++ {
    for col := 0; col < columnas; col++ {
        fmt.Fprintf(fout, "%c", matriz[fila][col])
    }
}
```

- Ahora leemos **fila por fila** (de izquierda a derecha)
- `fmt.Fprintf(fout, "%c", c)` → escribe el carácter `c` en `fout` (archivo o pantalla)
- Con el ejemplo anterior: lee `H,L,M` luego `O,A,X` → **"HLMOAX"**

### 4.13 Salto de línea final

```go
if fout == os.Stdout {
    fmt.Fprintln(fout)
}
```

- Solo añade `\n` si escribimos en la pantalla (para que el terminal quede bien)
- Si escribimos en un archivo, no necesitamos el salto extra

---

## 5. 🔢 El Algoritmo Paso a Paso

### Ejemplo Completo

**Entrada:** `"HOLA MUNDO ES PAÑA"` | **Filas:** 4 | **Columnas:** 4

#### Paso 1: Limpiar el texto
```
"HOLA MUNDO ES PAÑA"
 ↓ eliminar espacios, pasar a mayúsculas
"HOLAMUNDOSPAÑA"
 ↓ 14 letras, necesitamos 4×4=16 → rellenar con XX
"HOLAMUNDOSPAÑAXX"
```

#### Paso 2: Llenar la matriz por columnas
```
Texto: H O L A M U N D O S P A Ñ A X X
       0 1 2 3 4 5 6 7 8 9 ...

Escribir columna 0 (índices 0,1,2,3): H,O,L,A
Escribir columna 1 (índices 4,5,6,7): M,U,N,D
Escribir columna 2 (índices 8,9,10,11): O,S,P,A
Escribir columna 3 (índices 12,13,14,15): Ñ,A,X,X

       col0  col1  col2  col3
fila0: [ H  ,  M  ,  O  ,  Ñ ]
fila1: [ O  ,  U  ,  S  ,  A ]
fila2: [ L  ,  N  ,  P  ,  X ]
fila3: [ A  ,  D  ,  A  ,  X ]
```

#### Paso 3: Leer por filas = texto cifrado
```
fila0: H M O Ñ
fila1: O U S A
fila2: L N P X
fila3: A D A X

Resultado: "HMOÑOUSALN PXADAX"
```

---

## 6. 📊 Tipos de Datos Clave

| Tipo | ¿Qué es? | Ejemplo |
|------|----------|---------|
| `int` | Número entero | `4`, `-1`, `100` |
| `string` | Cadena de texto | `"hola"`, `"escitala"` |
| `rune` | Carácter Unicode (= `int32`) | `'A'`, `'Ñ'`, `'€'` |
| `bool` | Verdadero o falso | `true`, `false` |
| `error` | Tipo para errores | `nil` (sin error) o mensaje de error |
| `[]rune` | Slice (lista) de caracteres | `['H','O','L','A']` |
| `[][]rune` | Matriz 2D de caracteres | `[['H','M'],['O','U']]` |
| `map[rune]bool` | Diccionario carácter→booleano | `{'A':true, 'B':true}` |
| `*os.File` | Puntero a un archivo | Manejador de fichero |

---

## 7. 🔄 Flujo Completo con Ejemplo

```
                    ┌─────────────────────┐
                    │   Inicio del main   │
                    └──────────┬──────────┘
                               │
                    ┌──────────▼──────────┐
                    │  Declarar variables │
                    │  fin, fout, err,    │
                    │  filas, columnas    │
                    └──────────┬──────────┘
                               │
                    ┌──────────▼──────────┐
                    │  Definir alfabeto   │
                    │  (mapa A-Z + Ñ)     │
                    └──────────┬──────────┘
                               │
                    ┌──────────▼──────────┐
                    │  Asignar stdin/     │
                    │  stdout por defecto │
                    └──────────┬──────────┘
                               │
                    ┌──────────▼──────────┐
                    │  ¿Cuántos args?     │
                    └──┬────────┬─────┬───┘
                       │        │     │
                  3 args     4 args  5 args
                  stdin/   archivo/  archivo/
                  stdout    stdout   archivo
                       │        │     │
                    ┌──▼────────▼─────▼───┐
                    │  Leer texto letra   │
                    │  a letra → filtrar  │
                    │  solo letras válidas│
                    └──────────┬──────────┘
                               │
                    ┌──────────▼──────────┐
                    │  ¿Texto suficiente? │
                    │  No → rellenar 'X'  │
                    │  Sí → truncar       │
                    └──────────┬──────────┘
                               │
                    ┌──────────▼──────────┐
                    │  Crear matriz       │
                    │  make([][]rune,     │
                    │       filas)        │
                    └──────────┬──────────┘
                               │
                    ┌──────────▼──────────┐
                    │  Llenar por COLUMNAS│
                    │  (bucle col fuera,  │
                    │   fila dentro)      │
                    └──────────┬──────────┘
                               │
                    ┌──────────▼──────────┐
                    │  Leer por FILAS     │
                    │  → imprimir cifrado │
                    └──────────┬──────────┘
                               │
                    ┌──────────▼──────────┐
                    │       FIN           │
                    └─────────────────────┘
```

---

## 8. ⌨️ Comandos de Ejecución

```bash
# Caso 1: leer del teclado, imprimir en pantalla
go run escitala.go 4 4

# Caso 2: leer de un archivo, imprimir en pantalla
go run escitala.go entrada.txt 4 4

# Caso 3: leer de un archivo, guardar en otro archivo
go run escitala.go entrada.txt salida.txt 4 4

# Construir el ejecutable
go build escitala.go

# Ejecutar el binario compilado (Windows)
.\escitala.exe 4 4
```

---

## 9. ❓ Posibles Preguntas de Examen

### Pregunta 1: ¿Qué hace `defer`?
> `defer` pospone la ejecución de una sentencia hasta que la función que la contiene termina. Se usa principalmente para cerrar recursos (archivos, conexiones) y garantizar que siempre se cierren, aunque ocurra un error.

### Pregunta 2: ¿Por qué se usa `rune` en lugar de `byte`?
> `byte` solo puede representar caracteres ASCII (0-127). `rune` (que es un alias de `int32`) puede representar cualquier carácter Unicode, incluyendo la **Ñ** y letras con acento, que son necesarias en el alfabeto castellano.

### Pregunta 3: ¿Por qué el bucle exterior es por columnas y el interior por filas?
> Porque el algoritmo de la escítala escribe el texto **columna por columna** (de arriba a abajo en cada columna). Si invirtiéramos los bucles, estaríamos llenando por filas y el cifrado sería incorrecto.

### Pregunta 4: ¿Qué diferencia hay entre `os.Open` y `os.Create`?
> - `os.Open` abre un archivo **existente** solo para **lectura**. Falla si el archivo no existe.
> - `os.Create` crea un archivo nuevo para **escritura**. Si ya existe, lo sobreescribe desde cero.

### Pregunta 5: ¿Para qué sirve `os.Exit(1)`?
> Termina el programa inmediatamente con un **código de salida 1** (error). Por convención, el código 0 significa éxito y cualquier otro número indica un error. Los scripts y sistemas operativos usan este código para saber si el programa terminó correctamente.

### Pregunta 6: ¿Qué es `os.Args`?
> Es un **slice de strings** que contiene los argumentos pasados al programa en la línea de comandos. `os.Args[0]` siempre es el nombre del programa, y los demás son los parámetros del usuario.

### Pregunta 7: ¿Por qué rellenamos con `'X'` y no con espacios?
> Los espacios son ignorados durante la lectura del texto (el programa solo acepta letras del alfabeto). Si usáramos un espacio como relleno, sería ignorado en una segunda pasada. La `'X'` es un carácter válido del alfabeto, así que se procesa correctamente.

### Pregunta 8: ¿Qué hace `unicode.ToUpper(c)`?
> Convierte un carácter a su equivalente en **mayúscula**. Funciona con cualquier carácter Unicode, no solo con el ASCII básico, por lo que maneja correctamente la ñ, vocales con tilde, etc.

### Pregunta 9: ¿Qué significa `nil` en Go?
> `nil` representa la **ausencia de valor** para tipos como punteros, slices, maps, interfaces y errores. Si `err == nil`, significa que no hubo ningún error. Si un puntero es `nil`, no apunta a ningún lugar en memoria.

### Pregunta 10: ¿Qué diferencia hay entre `fmt.Fprintf`, `fmt.Fprintln` y `fmt.Printf`?
> | Función | Escribe en... | Formato |
> |---------|--------------|---------|
> | `fmt.Printf(formato, ...)` | `os.Stdout` (pantalla) | Con formato |
> | `fmt.Fprintf(w, formato, ...)` | Cualquier `io.Writer` (archivo, pantalla...) | Con formato |
> | `fmt.Fprintln(w, ...)` | Cualquier `io.Writer` | Sin formato, añade `\n` al final |

---

## 10. ⚡ Resumen Rápido

```
╔══════════════════════════════════════════════════════════════╗
║                    CHULETA ESCÍTALA GO                       ║
╠══════════════════════════════════════════════════════════════╣
║                                                              ║
║  ESTRUCTURA:      package main → import → func main()        ║
║                                                              ║
║  VARIABLES:       var x tipo  ó  x := valor                  ║
║                                                              ║
║  TIPOS CLAVE:     rune (carácter), []rune (texto),           ║
║                   map[rune]bool (alfabeto)                   ║
║                                                              ║
║  ARCHIVOS:        os.Open (leer) | os.Create (escribir)      ║
║                   defer f.Close() → siempre cerrar           ║
║                                                              ║
║  ARGUMENTOS:      os.Args[0]=programa, [1],[2]...=usuario    ║
║                   strconv.Atoi("4") → 4 (texto→número)       ║
║                                                              ║
║  LECTURA:         fmt.Fscanf(fin, "%c", &c) → un carácter   ║
║                   unicode.ToUpper(c) → mayúscula             ║
║                                                              ║
║  ERRORES:         _, err = funcion()                         ║
║                   if err != nil { os.Exit(1) }               ║
║                                                              ║
║  MATRIZ:          make([][]rune, filas)                       ║
║                   Llenar por COLUMNAS → leer por FILAS        ║
║                                                              ║
║  ALGORITMO:       1. Leer texto → filtrar letras válidas      ║
║                   2. Rellenar con X si falta                 ║
║                   3. Poner en matriz por COLUMNAS            ║
║                   4. Leer de la matriz por FILAS             ║
║                   5. Resultado = texto CIFRADO               ║
╚══════════════════════════════════════════════════════════════╝
```

---

> 📌 **Recuerda:** La clave del algoritmo está en que **se escribe por columnas pero se lee por filas**. Eso es lo que desordena el texto y produce el cifrado.

> 🍀 **¡Mucha suerte en el examen!**
