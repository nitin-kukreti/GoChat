package server

import (
	"fmt"
	"log"
	"net/http"
)

type App interface {
	GET(path string, handler http.HandlerFunc)
	POST(path string, handler http.HandlerFunc)
	DELETE(path string, handler http.HandlerFunc)
	PUT(path string, handler http.HandlerFunc)
	PATCH(path string, handler http.HandlerFunc)
	Listen(port string)
	Group(prefix string) *Group
}

type Server struct {
	mux *http.ServeMux
}

func NewServer() *Server {
	return &Server{
		mux: http.NewServeMux(),
	}
}

func (s *Server) register(method, path string, handler http.HandlerFunc) {
	pattern := fmt.Sprintf("%s %s", method, path)
	s.mux.HandleFunc(pattern, handler)
}

func (s *Server) GET(path string, handler http.HandlerFunc) {
	s.register(http.MethodGet, path, handler)
}

func (s *Server) POST(path string, handler http.HandlerFunc) {
	s.register(http.MethodPost, path, handler)
}

func (s *Server) DELETE(path string, handler http.HandlerFunc) {
	s.register(http.MethodDelete, path, handler)
}

func (s *Server) PUT(path string, handler http.HandlerFunc) {
	s.register(http.MethodPut, path, handler)
}

func (s *Server) PATCH(path string, handler http.HandlerFunc) {
	s.register(http.MethodPatch, path, handler)
}

func (s *Server) Listen(port string) {
	log.Printf("🚀 Server listening on http://localhost%s\n", port)
	err := http.ListenAndServe(port, s.mux)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func (s *Server) Group(prefix string) *Group {
	return &Group{prefix: prefix, app: s}
}





type Group struct {
	prefix string
	app    App
}

func (g *Group) GET(path string, h http.HandlerFunc)    { g.app.GET(g.prefix+path, h) }
func (g *Group) POST(path string, h http.HandlerFunc)   { g.app.POST(g.prefix+path, h) }
func (g *Group) PUT(path string, h http.HandlerFunc)    { g.app.PUT(g.prefix+path, h) }
func (g *Group) PATCH(path string, h http.HandlerFunc)  { g.app.PATCH(g.prefix+path, h) }
func (g *Group) DELETE(path string, h http.HandlerFunc) { g.app.DELETE(g.prefix+path, h) }



