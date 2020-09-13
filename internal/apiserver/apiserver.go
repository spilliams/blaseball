package apiserver

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/spilliams/blaseball/internal"
	"github.com/spilliams/blaseball/internal/memdata"
	"github.com/spilliams/blaseball/pkg"
	"github.com/spilliams/blaseball/pkg/api"
)

type Server struct {
	dataSession internal.LocalDataSession
	remoteAPI   internal.RemoteDataSession
}

func newServer() *Server {
	return &Server{
		dataSession: memdata.NewSession(),
		remoteAPI:   api.NewAPI("https://www.blaseball.com/database/", logrus.DebugLevel),
	}
}

func StartHTTPServer(port string) error {
	s := newServer()
	router := mux.NewRouter()
	// TODO logger middleware
	// TODO auth middleware
	router.Handle("/allDivisions", handler{s.GetDivisions})
	router.Handle("/division", handler{s.GetDivision})
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
	customError, ok := err.(pkg.Coded)
	if !ok {
		logrus.Warnf("Error came back that wasn't custom-typed: %v", err)
		http.Error(w, fmt.Sprintf(`{"error": "%s"}`, err), http.StatusInternalServerError)
		return
	}
	code := customError.StatusCode()
	if code == 0 {
		code = http.StatusInternalServerError
	}

	errMsg := fmt.Sprintf(`{"status": "%d", "error": "%s"}`, code, customError.Error())
	http.Error(w, errMsg, code)
}

func marshalAndWrite(obj interface{}, w http.ResponseWriter) error {
	bytes, err := json.Marshal(obj)
	if err != nil {
		return fmt.Errorf("could not marshal response: %v", err)
	}
	_, err = w.Write(bytes)
	return err
}

func getQueryString(r *http.Request, key string) string {
	return r.URL.Query().Get(key)
}
