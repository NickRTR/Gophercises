package main

import "net/http"

func redirectToPortfolio(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "https://nickreutlinger.de", http.StatusSeeOther)
}

func redirectToBlog(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "https://blog.nickreutlinger.de", http.StatusSeeOther)
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/portfolio", redirectToPortfolio)
	mux.HandleFunc("/blog", redirectToBlog)

	http.ListenAndServe(":8080", mux)
}
