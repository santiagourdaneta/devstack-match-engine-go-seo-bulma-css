package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"os"
	"time"
)

type TechPage struct {
	ID          string `json:"ID"`
	Title       string `json:"Title"`
	Description string `json:"Description"`
	Category    string `json:"Category"`
	Rating      string `json:"Rating"`
	Votes       string `json:"Votes"`
}

func main() {
	// 1. Cargar Datos
	file, _ := os.ReadFile("data.json")
	var pages []TechPage
	json.Unmarshal(file, &pages)

	distDir := "dist"
	os.MkdirAll(distDir, 0755)
	currentDate := time.Now().Format("2006-01-02")

	// --- TEMPLATE PARA LAS PÁGINAS INTERNAS ---
	itemTmpl := template.Must(template.New("item").Parse(`
<!DOCTYPE html>
<html lang="es">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>{{.Current.Title}} | DevStack-Match</title>
    <link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/bulma@0.9.4/css/bulma.min.css">
    <style>body { background: #fafafa; } .hero { background: #00d1b2; color: white; }</style>
</head>
<body>
    <section class="hero is-small">
        <div class="hero-body"><div class="container"><h1 class="title">{{.Current.Title}}</h1></div></div>
    </section>
    <section class="section">
        <div class="container">
            <div class="box">
                <span class="tag is-info">{{.Current.Category}}</span>
                <p class="is-size-4 mt-4">{{.Current.Description}}</p>
                <div class="notification is-light mt-5">⭐ Puntaje: {{.Current.Rating}} ({{.Current.Votes}} votos)</div>
                <hr>
                <h3 class="subtitle">Herramientas relacionadas:</h3>
                <div class="buttons">
                    {{range .Related}}<a href="/{{.ID}}" class="button is-light">{{.Title}}</a>{{end}}
                </div>
                <a href="/" class="button is-primary mt-6">← Volver al buscador</a>
            </div>
        </div>
    </section>
</body>
</html>`))

	// --- TEMPLATE PARA EL INDEX (EL HOME) ---
	indexTmpl := template.Must(template.New("index").Parse(`
<!DOCTYPE html>
<html lang="es">
<head>
    <meta charset="UTF-8">
    <title>DevStack-Match | Comparador de Tecnología</title>
    <link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/bulma@0.9.4/css/bulma.min.css">
</head>
<body>
    <section class="hero is-medium is-dark">
        <div class="hero-body has-text-centered">
            <h1 class="title is-1">🚀 DevStack-Match</h1>
            <p class="subtitle">Encuentra tu próximo stack con datos reales</p>
        </div>
    </section>
    <section class="section">
        <div class="container">
            <div class="columns is-multiline">
                {{range .}}
                <div class="column is-4">
                    <div class="card">
                        <div class="card-content">
                            <p class="title is-4">{{.Title}}</p>
                            <p class="subtitle is-6">{{.Category}}</p>
                            <a href="/{{.ID}}" class="button is-primary is-fullwidth">Ver comparativa</a>
                        </div>
                    </div>
                </div>
                {{end}}
            </div>
        </div>
    </section>
</body>
</html>`))

	// 2. Generar el Home (Index)
	fIndex, _ := os.Create(distDir + "/index.html")
	indexTmpl.Execute(fIndex, pages)
	fIndex.Close()

	// 3. Generar Páginas Internas y Sitemap
	sitemapBody := ""
	for _, p := range pages {
		var related []TechPage
		for _, r := range pages {
			if r.Category == p.Category && r.ID != p.ID {
				related = append(related, r)
			}
		}

		data := struct {
			Current TechPage
			Related []TechPage
		}{p, related}

		f, _ := os.Create(fmt.Sprintf("%s/%s.html", distDir, p.ID))
		itemTmpl.Execute(f, data)
		f.Close()

		sitemapBody += fmt.Sprintf("\n  <url><loc>https://devstack-match.com/%s</loc><lastmod>%s</lastmod></url>", p.ID, currentDate)
	}

	// 4. Guardar Sitemap y Favicon
	os.WriteFile(distDir+"/sitemap.xml", []byte(`<?xml version="1.0" encoding="UTF-8"?><urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">`+sitemapBody+`</urlset>`), 0644)
	os.WriteFile(distDir+"/favicon.ico", []byte(""), 0644)

	fmt.Println("🚀 Sitio completo generado con éxito.")
}
