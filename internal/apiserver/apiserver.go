package apiserver

import (
	"fmt"
	"net"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/spilliams/blaseball/internal"
	"github.com/spilliams/blaseball/internal/memdata"
	"github.com/spilliams/blaseball/pkg/api"
)

type Server struct {
	dataSession internal.LocalDataSession
	remoteAPI   internal.RemoteDataSession
}

func newServer() *Server {
	return &Server{
		dataSession: memdata.NewSession(),
		remoteAPI:   api.NewAPI("https://www.blaseball.com/database/"),
	}
}

func StartHTTPServer(port string) error {
	s := newServer()
	router := mux.NewRouter()
	// TODO logger middleware
	// TODO auth middleware
	router.Handle("/allDivisions", handler{s.GetDivisions})
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return err
	}
	handler := handlers.CompressHandler(router)
	handler = setContentTypeJSON(handler)
	go http.Serve(listener, handler)

	fmt.Printf("server started at port :%v\n", port)
	return nil
}

func setContentTypeJSON(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		h.ServeHTTP(w, r)
	})
}

type handler struct {
	httpFunc func(http.ResponseWriter, *http.Request) error
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := h.httpFunc(w, r)
	if err == nil {
		return
	}

	// log the error
	code := 0 // TODO get the code out of the error
	if code == 0 {
		code = http.StatusInternalServerError
	}

	errMsg := fmt.Sprintf(`{"error": "%s"}`, err)
	http.Error(w, errMsg, code)
}
