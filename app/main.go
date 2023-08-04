package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})
	http.HandleFunc("/admin", adminHandler)
	http.HandleFunc("/get-url", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Método inválido", http.StatusMethodNotAllowed)
			return
		}

		url := r.FormValue("url")

		// Fazendo a requisição GET
		resp, err := http.Get(url)
		if err != nil {
			http.Error(w, "Erro ao fazer a requisição: "+err.Error(), http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()

		// Lendo o corpo da resposta
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			http.Error(w, "Erro ao ler o corpo da resposta: "+err.Error(), http.StatusInternalServerError)
			return
		}

		// Escrevendo a resposta no cliente
		w.Header().Set("Access-Control-Allow-Origin", "*") // Permitir acesso de qualquer origem para /get-url
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.Write(body)
	})

	fmt.Println("Servidor rodando na porta 8081...")
	http.ListenAndServe(":8081", nil) // Alterar a porta para 8081
}

func adminHandler(w http.ResponseWriter, r *http.Request) {
	// Verificar o endereço IP de origem da solicitação
	remoteAddr := strings.Split(r.RemoteAddr, ":")[0]

	// Permitir o acesso apenas se o endereço IP de origem for localhost (127.0.0.1)
	if remoteAddr != "127.0.0.1" {
		http.Error(w, "Acesso negado.", http.StatusForbidden)
		return
	}

	fmt.Fprintf(w, "Parabéns! Você acessou o endpoint /admin.")
}
