# 🚀 DevStack-Match Engine

**DevStack-Match** es un generador de sitios estáticos (SSG) de alto rendimiento construido en **Go**. Está diseñado específicamente para dominar las SERPs mediante **Programmatic SEO**, automatizando la creación de miles de páginas de aterrizaje optimizadas a partir de un simple archivo de datos.

## ✨ Características Técnicas (SEO Nivel Dios)

-   **⚡ Zero-Node Architecture:** Sin dependencias pesadas de JS. HTML puro y CSS (Bulma) para una carga instantánea.
-   **🖼️ Dynamic OG Image Generation:** Generación automática de imágenes Open Graph únicas por cada página usando la librería `gg`.
-   **🔍 Programmatic SEO:** Creación masiva de páginas (`ID.html`) con **Pretty URLs** configuradas.
-   **🗺️ Dynamic Sitemap & Schema:** Generación de `sitemap.xml` y datos estructurados JSON-LD (`SoftwareApplication`) en tiempo real.
-   **🛡️ Quality Gates:** Pre-commit hooks integrados para `gofmt` y `go vet` que aseguran código limpio.
-   **📦 Self-Contained:** Script de autodescarga de tipografías y servidor local integrado en Go.

## 🛠️ Stack Tecnológico

-   **Backend:** Go 1.2x (Módulos activos).
-   **Frontend:** Bulma CSS (Responsive Design).
-   **Gráficos:** `fogleman/gg` para renderizado de imágenes.
-   **Automatización:** Makefile.

## 🚀 Inicio Rápido

### 1. Preparar el entorno
```bash
make setup

2. Generar el sitio completo
Este comando compila el binario, descarga recursos, genera HTMLs, imágenes OG y el sitemap.

make all

3. Previsualizar en local
Lanza un servidor en Go que emula las rutas de producción (Pretty URLs).

make serve

Accede a: http://localhost:8080

📁 Estructura del Proyecto

/dist: El sitio estático final listo para desplegar (GitHub Pages/Vercel).

data.json: La fuente de verdad. Añade una línea aquí para crear una nueva página.

main.go: El motor de generación y lógica SEO.

server.go: Servidor de desarrollo optimizado.

Makefile: Orquestador de tareas.

🤝 Contribución
Agrega una nueva tecnología o comparativa en data.json.

Ejecuta make all.

Verifica el diseño y los metadatos.

Envía tu Pull Request.