# 📚 Guía de Colaboración

## 🚀 PARTE 1: Subir el Proyecto a GitHub (TÚ)

### Paso 1: Crear el repositorio en GitHub

1. Ve a [github.com](https://github.com) e inicia sesión
2. Haz clic en el botón **"+"** (arriba a la derecha) → **"New repository"**
3. Configura el repositorio:
   - **Repository name**: `cifrado-cesar-go` (o el nombre que prefieras)
   - **Description**: "Implementación del cifrado César en Go con alfabeto español"
   - **Visibilidad**: 
     - ✅ **Public** - Cualquiera puede verlo (recomendado para proyectos de aprendizaje)
     - 🔒 **Private** - Solo tú y colaboradores invitados pueden verlo
   - ⚠️ **NO marques** "Add a README file" (ya lo tenemos)
   - ⚠️ **NO marques** "Add .gitignore" (ya lo tenemos)
4. Haz clic en **"Create repository"**

### Paso 2: Conectar tu repositorio local con GitHub

Después de crear el repositorio, GitHub te mostrará instrucciones. Ejecuta estos comandos **desde la carpeta P0**:

```bash
# Navega a tu carpeta del proyecto (si no estás ahí)
cd C:\Users\IkerP\Desktop\P0

# Conecta tu repositorio local con GitHub
git remote add origin https://github.com/TU_USUARIO/cifrado-cesar-go.git

# Renombra la rama principal a 'main' (estándar de GitHub)
git branch -M main

# Sube tu código a GitHub
git push -u origin main
```

⚠️ **Importante**: Reemplaza `TU_USUARIO` con tu nombre de usuario de GitHub.

### Paso 3: Invitar a tu compañero (si el repo es privado)

Si elegiste **repositorio privado**:

1. Ve a tu repositorio en GitHub
2. Haz clic en **"Settings"** (Configuración)
3. En el menú lateral, haz clic en **"Collaborators"**
4. Haz clic en **"Add people"**
5. Busca a tu compañero por su usuario de GitHub
6. Envía la invitación

---

## 👥 PARTE 2: Cómo Trabaja tu Compañero

### Opción A: Clonar el Repositorio (Primera vez)

Tu compañero debe ejecutar estos comandos **desde donde quiera guardar el proyecto**:

```bash
# Ejemplo: Si quiere el proyecto en su Desktop
cd C:\Users\SU_NOMBRE\Desktop

# Clonar el repositorio
git clone https://github.com/TU_USUARIO/cifrado-cesar-go.git

# Entrar a la carpeta del proyecto
cd cifrado-cesar-go
```

Ahora tu compañero tiene una copia completa del proyecto en su computadora.

---

## 🔄 Flujo de Trabajo Colaborativo

### Para TU COMPAÑERO: Hacer cambios y subirlos

```bash
# 1. Asegurarse de estar en la carpeta del proyecto
cd C:\Users\SU_NOMBRE\Desktop\cifrado-cesar-go

# 2. Descargar los últimos cambios (por si tú hiciste algo nuevo)
git pull origin main

# 3. Hacer modificaciones en los archivos (editar cesar.go, etc.)

# 4. Ver qué archivos cambiaron
git status

# 5. Agregar los archivos modificados
git add .
# O agregar archivos específicos:
# git add cesar.go

# 6. Hacer commit con un mensaje descriptivo
git commit -m "Descripción de los cambios que hice"

# 7. Subir los cambios a GitHub
git push origin main
```

### Para TI: Descargar los cambios de tu compañero

```bash
# 1. Ir a tu carpeta del proyecto
cd C:\Users\IkerP\Desktop\P0

# 2. Descargar los cambios que hizo tu compañero
git pull origin main
```

---

## 📋 Comandos Esenciales de Git

| Comando | Descripción |
|---------|-------------|
| `git status` | Ver qué archivos han cambiado |
| `git add .` | Preparar TODOS los archivos modificados |
| `git add archivo.go` | Preparar un archivo específico |
| `git commit -m "mensaje"` | Guardar los cambios con un mensaje |
| `git push origin main` | Subir cambios a GitHub |
| `git pull origin main` | Descargar cambios de GitHub |
| `git log --oneline` | Ver historial de commits |
| `git diff` | Ver diferencias en archivos modificados |

---

## 🎯 Ejemplo Práctico Completo

### Escenario: Tu compañero quiere agregar una función de descifrado

**Tu compañero hace:**

```bash
# 1. Clonar el proyecto (solo la primera vez)
cd C:\Users\Compañero\Desktop
git clone https://github.com/IkerP/cifrado-cesar-go.git
cd cifrado-cesar-go

# 2. Editar el archivo cesar.go (agregar nueva función)

# 3. Probar que funciona
go run cesar.go

# 4. Guardar y subir cambios
git add cesar.go
git commit -m "Agregada función de descifrado automático"
git push origin main
```

**Tú descargas sus cambios:**

```bash
cd C:\Users\IkerP\Desktop\P0
git pull origin main
# Ahora tienes la nueva función en tu computadora
```

---

## ⚠️ Problemas Comunes y Soluciones

### Problema 1: "Permission denied" al hacer push

**Solución**: Necesitas autenticarte en GitHub. Usa un Personal Access Token:

1. Ve a GitHub → Settings → Developer settings → Personal access tokens → Tokens (classic)
2. Genera un nuevo token con permisos de "repo"
3. Copia el token
4. Cuando Git te pida contraseña, pega el token (no tu contraseña de GitHub)

### Problema 2: Conflictos al hacer pull

**Solución**: Si ambos editaron el mismo archivo:

```bash
# Git te mostrará qué archivos tienen conflictos
git status

# Abre el archivo y busca las marcas de conflicto:
# <<<<<<< HEAD
# Tu código
# =======
# Código de tu compañero
# >>>>>>> 

# Edita el archivo para resolver el conflicto
# Luego:
git add archivo_resuelto.go
git commit -m "Resuelto conflicto en archivo.go"
git push origin main
```

### Problema 3: Olvidé hacer pull antes de hacer cambios

**Solución**:

```bash
# Guarda tus cambios temporalmente
git stash

# Descarga los cambios de GitHub
git pull origin main

# Recupera tus cambios
git stash pop

# Si hay conflictos, resuélvelos como en el Problema 2
```

---

## 🎓 Buenas Prácticas

1. ✅ **Siempre hacer `git pull` antes de empezar a trabajar**
2. ✅ **Hacer commits frecuentes con mensajes descriptivos**
3. ✅ **Probar el código antes de hacer push**
4. ✅ **Comunicarse con tu compañero sobre qué archivos están editando**
5. ✅ **Usar mensajes de commit claros**: 
   - ✅ "Agregada validación de entrada"
   - ❌ "cambios"

---

## 📞 Resumen Rápido

**Para TI (primera vez):**
```bash
cd C:\Users\IkerP\Desktop\P0
git remote add origin https://github.com/TU_USUARIO/cifrado-cesar-go.git
git branch -M main
git push -u origin main
```

**Para TU COMPAÑERO (primera vez):**
```bash
cd C:\Users\Compañero\Desktop
git clone https://github.com/TU_USUARIO/cifrado-cesar-go.git
cd cifrado-cesar-go
```

**Flujo diario (ambos):**
```bash
git pull origin main          # Descargar cambios
# ... hacer modificaciones ...
git add .                     # Preparar cambios
git commit -m "descripción"   # Guardar cambios
git push origin main          # Subir cambios
```

---

## 🔗 Recursos Adicionales

- [Documentación oficial de Git](https://git-scm.com/doc)
- [GitHub Guides](https://guides.github.com/)
- [Visualizador de Git](https://git-school.github.io/visualizing-git/)

---

**¡Listo para colaborar! 🚀**
