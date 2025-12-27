package main

import "net/http"
import "log"

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")

		w.Write([]byte(`
			<!DOCTYPE html>
			<html lang="es">
			<head>
				<meta charset="UTF-8">
				<title>Go Server</title>
			</head>
			<body>
				<h1>Hola desde Go 👋</h1>
				<p>Servidor HTTP funcionando.</p>
			</body>
			</html>
		`))
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}

