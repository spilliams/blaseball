package apiserver

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"strings"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/spilliams/blaseball/internal"
	"github.com/spilliams/blaseball/pkg"
)

// Server represents a web server that can handle requests about blaseball.
type Server struct {
	dataStore internal.LocalDataSession
	remoteAPI pkg.RemoteDataSession
}

// NewServer returns a new server with the given local data session (for
// storing) and remote data session (for fetching)
func NewServer(local internal.LocalDataSession, remote pkg.RemoteDataSession) *Server {
	return &Server{
		dataStore: local,
		remoteAPI: remote,
	}
}

// StartHTTPServer starts a TCP listener on the given port
func (s *Server) StartHTTPServer(port string) error {
	router := mux.NewRouter()

	router.Use(loggerMiddleware)

	router.Handle("/allDivisions", handler{s.GetDivisions})
	router.Handle("/division", handler{s.GetDivision})
	router.Handle("/allTeams", handler{s.GetTeams})
	router.Handle("/team", handler{s.GetTeam})
	router.Handle("/allPlayers", handler{s.GetAllPlayers})
	router.Handle("/players", handler{s.GetPlayers})

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

func loggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger := loggerFromRequest(r)
		r = requestWithLogger(r, logger)
		logger.Infof("new request from %s", r.UserAgent())
		next.ServeHTTP(w, r)
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

func getQueryStrings(r *http.Request, key string) []string {
	join := r.URL.Query().Get(key)
	if join == "" {
		return []string{}
	}
	return strings.Split(join, ",")
}
