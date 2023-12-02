package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/google/uuid"
)

type heartbeat struct {
	Name    string
	Version string
	Message string
}

func heartbeatRoutes() *chi.Mux {
	r := chi.NewRouter()
	r.Get("/", getHeartbeat())
	return r
}

func getHeartbeat() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		log.Println("starting heartbeat request...")
		resp := heartbeat{
			Name:    "todo backend",
			Version: "1.0",
			Message: "heartbeat",
		}
		log.Println("heartbeat request done.")
		render.JSON(rw, r, resp)
	}
}

type apiResp struct {
	ReqId string `json:"request_id"`
	Data  []todo `json:"data"`
}

type todo struct {
	Id      int
	Title   string
	Message string
}

func todoRoutes() *chi.Mux {
	r := chi.NewRouter()
	r.Get("/", getAllTodos())
	return r
}

func getAllTodos() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		log.Println("starting getAllTodos request...")
		// TODO: add code to retrieve todos from DB

		todos := []todo{
			{Id: 1, Title: "Eat", Message: "Make sure to eat"},
			{Id: 2, Title: "Sleep", Message: "Get good sleep"},
			{Id: 3, Title: "Code", Message: "Always be coding"},
		}

		resp := apiResp{
			ReqId: uuid.New().String(),
			Data:  todos,
		}
		log.Println("getAllTodos requeset done.")
		render.JSON(rw, r, resp)
	}
}
