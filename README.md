# Práctica 0 - Cifrados Clásicos en Go

Este proyecto implementa diferentes métodos de cifrado clásico en Go, incluyendo el cifrado César y el cifrado Escítala.

## 📁 Estructura del Proyecto

```
P0/
├── Cesar/              # Implementación del cifrado César
│   ├── cesar.go
│   └── cesar.exe
├── Escitala/           # Implementación del cifrado Escítala
│   └── escitala.go
├── Tests/              # Archivos de prueba
│   ├── entrada.txt
│   ├── salida.txt
│   └── salidaa.txt
├── Documentacion/      # Documentación del proyecto
│   ├── INSTRUCCIONES_COLABORACION.md
│   └── SC - Práctica 0.pdf
├── .vscode/            # Configuración de VS Code y debugger
│   └── launch.json
├── go.mod              # Módulo de Go
├── .gitignore
└── README.md
```

## 🔐 Cifrado César

El programa implementa el cifrado César para el alfabeto español (incluyendo la Ñ).

### Características

- ✅ Alfabeto español completo (A-Z + Ñ, 27 letras)
- ✅ Conversión automática a mayúsculas
- ✅ Ignora espacios y caracteres no alfabéticos
- ✅ Soporta claves positivas y negativas
- ✅ Lectura desde stdin o archivo
- ✅ Escritura a stdout o archivo

### Uso

#### 1. Sin parámetros (clave por defecto = 3)
```bash
cd Cesar
go run cesar.go
```

#### 2. Con clave personalizada
```bash
cd Cesar
go run cesar.go 5
```

#### 3. Con archivos de entrada/salida
```bash
cd Cesar
go run cesar.go ../Tests/entrada.txt ../Tests/salida.txt
```

#### 4. Con clave y archivos
```bash
cd Cesar
go run cesar.go 7 ../Tests/entrada.txt ../Tests/salida.txt
```

### Ejemplos

**Cifrar con clave 3:**
- Entrada: `HOLAMUNDO`
- Salida: `KRODPXQGR`

**Descifrar (clave negativa):**
```bash
echo "KRODPXQGR" | go run cesar.go -3
```
- Salida: `HOLAMUNDO`

## 🔧 Compilación

### Compilar Cesar
```bash
cd Cesar
go build -o cesar.exe cesar.go
```

### Compilar Escitala
```bash
cd Escitala
go build -o escitala.exe escitala.go
```

## 🐛 Debugging en VS Code

El proyecto incluye configuraciones de debug en `.vscode/launch.json`:

1. **Launch Current File** - Ejecuta el archivo Go actual
2. **Debug Cesar** - Debug específico para el programa Cesar
3. **Debug Escitala** - Debug específico para el programa Escitala

Para usar el debugger:
1. Abre el archivo que quieres debuggear
2. Presiona `F5` o ve a Run > Start Debugging
3. Selecciona la configuración apropiada

## 📋 Requisitos

- Go 1.21.6 o superior

## 👥 Autor

Proyecto de práctica de cifrados clásicos
