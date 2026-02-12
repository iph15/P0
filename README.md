# Cifrado César en Go

Este programa implementa el cifrado César para el alfabeto español (incluyendo la Ñ).

## Descripción

El programa lee texto de la entrada estándar o de un archivo, aplica el cifrado César con una clave de desplazamiento, y escribe el resultado en la salida estándar o en un archivo.

### Características

- ✅ Alfabeto español completo (A-Z + Ñ, 27 letras)
- ✅ Conversión automática a mayúsculas
- ✅ Ignora espacios y caracteres no alfabéticos
- ✅ Soporta claves positivas y negativas
- ✅ Lectura desde stdin o archivo
- ✅ Escritura a stdout o archivo

## Uso

El programa acepta diferentes combinaciones de parámetros:

### 1. Sin parámetros (clave por defecto = 3)
```bash
go run cesar.go
```
Lee de la entrada estándar y escribe en la salida estándar con clave 3.

### 2. Con clave personalizada
```bash
go run cesar.go 5
```
Lee de stdin, escribe a stdout, usa la clave especificada.

### 3. Con archivos de entrada/salida (clave por defecto = 3)
```bash
go run cesar.go entrada.txt salida.txt
```
Lee del archivo de entrada, escribe al archivo de salida, usa clave 3.

### 4. Con clave y archivos
```bash
go run cesar.go 7 entrada.txt salida.txt
```
Lee del archivo de entrada, escribe al archivo de salida, usa la clave especificada.

## Ejemplos

### Cifrar con clave 3
**Entrada:** `HOLAMUNDO`  
**Salida:** `KRODPXQGR`

### Descifrar (clave negativa)
```bash
echo "KRODPXQGR" | go run cesar.go -3
```
**Salida:** `HOLAMUNDO`

## Compilación

Para compilar el programa:
```bash
go build cesar.go
```

Esto generará un ejecutable `cesar.exe` (en Windows) o `cesar` (en Linux/Mac).

## Requisitos

- Go 1.16 o superior

## Autor

Proyecto de práctica del cifrado César
