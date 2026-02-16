#!/bin/bash

# Abortar si hay errores
set -e

echo "🚀 Iniciando despliegue..."

# 1. Ejecutar el Makefile para generar la versión más reciente
make all

# 2. Entrar en la carpeta de salida
cd dist

# 3. Inicializar un repo temporal solo para la rama de hosting
git init
git add .
git commit -m "Deploy: $(date)"

# 4. Forzar la subida a la rama gh-pages (reemplaza con tu URL de repo)
# git push -f git@github.com:tu-usuario/devstack-match.git main:gh-pages

echo "✅ ¡Sitio en vivo en GitHub Pages!"