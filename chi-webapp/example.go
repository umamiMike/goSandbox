package main

import (
	"net/http"
	"path/filepath"

	"github.com/go-chi/chi"
)

func main() {
	r := NewRouter()

	http.ListenAndServe(":3333", r)
}

func NewRouter() http.Handler {
	router := chi.NewRouter()
	//router.Get("/{name}", HelloName)

	// Set up static file serving
	staticPath, _ := filepath.Abs("./client/build/")
	fs := http.FileServer(http.Dir(staticPath))
	router.Handle("/index.html", fs)

	return router
}

//func main() {

//r := chi.NewRouter()
//spew.Dump(r)
//r.Use(middleware.RequestID)
//r.Use(middleware.Logger)

//r.Get("/", func(w http.ResponseWriter, r *http.Request) {
//w.Write([]byte("sup"))
//})

//r.Get("/slow", func(w http.ResponseWriter, r *http.Request) {
//w.Write([]byte(fmt.Sprintf("all done.\n")))
//})

//r.Get("/site", func(router chi.Router, w http.ResponseWriter, r *http.Request) {
//staticPath, _ := filepath.Abs("./client/build")
//FileServer(router, "client/build", "/site", http.Dir("staticPath"))
////r.Handle("/*", fs)

//})
//fmt.Println("starting sterver on port http://localhost:3333")
//}

// FileServer conveniently sets up a http.FileServer handler to serve
// static files from a http.FileSystem.
//func FileServer(r chi.Router, basePath string, path string, root http.FileSystem) {
//if strings.ContainsAny(path, "{}*") {
//panic("FileServer does not permit URL parameters.")
//}

//fs := http.StripPrefix(basePath+path, http.FileServer(root))

//if path != "/" && path[len(path)-1] != '/' {
//r.Get(path, http.RedirectHandler(path+"/", 301).ServeHTTP)
//path += "/"
//fmt.Println(path)
//}
//path += "*"

//r.Get(path, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//fs.ServeHTTP(w, r)
//}))
//}
