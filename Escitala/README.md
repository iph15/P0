# Cifrado Escítala (Scytale Cipher)

## Descripción

La **escítala** es un método de cifrado por transposición usado en la antigua Grecia y Esparta. Consiste en enrollar una tira de cuero o pergamino alrededor de un bastón (vara) de un grosor específico. El mensaje se escribía a lo largo del bastón, y al desenrollarlo, las letras quedaban desordenadas.

## Algoritmo

1. Se crea una **matriz** de `FILAS x COLUMNAS` (donde FILAS = capacidad de la vara)
2. El texto se **escribe** en la matriz **COLUMNA POR COLUMNA** (de arriba a abajo, de izquierda a derecha)
3. El texto cifrado se **lee FILA POR FILA** (de izquierda a derecha, de arriba a abajo)

## Ejemplo Visual

**Texto original:** `HOLA MUNDO ES PAÑA`  
**Filas:** 4  
**Columnas:** 4

### Paso 1: Limpiar el texto (eliminar espacios)
```
HOLAMUNDOESPAÑA → 15 caracteres
```

### Paso 2: Rellenar hasta completar la matriz (4×4 = 16 caracteres)
```
HOLAMUNDOESPAÑA + X → HOLAMUNDOESPAÑAX
```

### Paso 3: Escribir por COLUMNAS en la matriz

```
Columna:  1  2  3  4
        ┌──┬──┬──┬──┐
Fila 1  │ H│ M│ O│ A│
        ├──┼──┼──┼──┤
Fila 2  │ O│ U│ E│ Ñ│
        ├──┼──┼──┼──┤
Fila 3  │ L│ N│ S│ A│
        ├──┼──┼──┼──┤
Fila 4  │ A│ D│ P│ X│
        └──┴──┴──┴──┘
```

### Paso 4: Leer por FILAS (texto cifrado)

```
Fila 1: HMOA
Fila 2: OUEÑ
Fila 3: LNSA
Fila 4: ADPX

Resultado: HMOAOUEÑLNSAADPX
```

## Uso del Programa

### Sintaxis

```bash
# Opción 1: Entrada desde teclado (stdin), salida por pantalla (stdout)
go run escitala.go <filas> <columnas>

# Opción 2: Entrada desde archivo, salida por pantalla
go run escitala.go <archivo_entrada> <filas> <columnas>

# Opción 3: Entrada desde archivo, salida a archivo
go run escitala.go <archivo_entrada> <archivo_salida> <filas> <columnas>
```

### Ejemplos de Uso

#### Ejemplo 1: Entrada desde teclado
```bash
go run escitala.go 4 4
# Escribe: HOLA MUNDO ES PAÑA
# Presiona Ctrl+Z (Windows) o Ctrl+D (Linux/Mac)
# Salida: HMOAOUEÑLNSAADPX
```

#### Ejemplo 2: Desde archivo de entrada
```bash
go run escitala.go entrada.txt 4 4
# Lee de entrada.txt y muestra el resultado en pantalla
```

#### Ejemplo 3: Desde archivo a archivo
```bash
go run escitala.go entrada.txt salida.txt 4 4
# Lee de entrada.txt y guarda el resultado en salida.txt
```

#### Ejemplo 4: Diferentes dimensiones de matriz
```bash
# Matriz de 3 filas x 5 columnas
go run escitala.go entrada.txt 3 5

# Matriz de 5 filas x 3 columnas (resultado diferente)
go run escitala.go entrada.txt 5 3
```

## Características

✅ **Alfabeto castellano completo:** A-Z incluyendo la Ñ  
✅ **Conversión automática:** Entrada en minúsculas/mayúsculas → Salida en MAYÚSCULAS  
✅ **Ignora espacios:** Los espacios en blanco son eliminados automáticamente  
✅ **Relleno automático:** Si el texto es más corto que la matriz, se rellena con 'X'  
✅ **Truncado inteligente:** Si el texto es más largo, se trunca al tamaño de la matriz  
✅ **Manejo de errores:** Validación de parámetros y archivos  

## Requisitos

- **Go 1.16 o superior**
- Sistema operativo: Windows, Linux, macOS

## Compilación

Para compilar el programa:

```bash
go build escitala.go
```

Esto generará un ejecutable `escitala.exe` (Windows) o `escitala` (Linux/Mac).

## Pruebas

### Prueba 1: Texto del ejemplo
```bash
echo "HOLA MUNDO ES PAÑA" | go run escitala.go 4 4
# Resultado esperado: HMOAOUEÑLNSAADPX
```

### Prueba 2: Texto corto con relleno
```bash
echo "HOLA" | go run escitala.go 2 3
# Texto limpio: HOLA (4 caracteres)
# Matriz 2x3 necesita 6 caracteres → se añaden 2 'X'
# Resultado: HOLXAX
```

### Prueba 3: Verificar archivo de entrada
```bash
# Crear archivo de prueba
echo "HOLAMUNDOESPAÑA" > entrada.txt

# Cifrar
go run escitala.go entrada.txt salida.txt 4 4

# Ver resultado
type salida.txt  # Windows
cat salida.txt   # Linux/Mac
```

## Notas Técnicas

- El programa utiliza `rune` para manejar correctamente caracteres Unicode (como la Ñ)
- La matriz se implementa como `[][]rune` (slice bidimensional)
- Se usa `unicode.ToUpper()` para la conversión a mayúsculas
- Los archivos se manejan con `defer Close()` para garantizar su cierre

## Descifrado

Para descifrar un mensaje cifrado con escítala, se debe:
1. Usar las **mismas dimensiones** de matriz (filas y columnas)
2. **Invertir el proceso:** escribir por FILAS y leer por COLUMNAS

*Nota: Este programa solo implementa el cifrado. El descifrado requeriría un programa complementario.*

## Autor

Práctica de Criptografía - Cifrado Escítala
