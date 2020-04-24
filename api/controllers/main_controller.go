package controllers

import (
	"net/http"

	"github.com/mohitagarwal18/api-gateway/api/responses"
)

//sample response
type AsyncResponse struct {
	Status    string `json:"status"`
	RequestId string `json:"request_id"`
}
type AsyncFinalResponse struct {
	Status    string `json:"status"`
	RequestId string `json:"request_id"`
	ReqData   string `json: "data"`
}

//a healthcheck to API
func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "Welcome To This  API")

}

//function will checkif output is ready authenticate and provide it to user
func (server *Server) GetTask(w http.ResponseWriter, r *http.Request) {
	res := AsyncFinalResponse{Status: "OK", RequestId: "cc701a0e-0cfd-4160-b2e4-cadffd235b29", ReqData: "Here is your requested data"}
	responses.JSON(w, http.StatusOK, res)

}

//function will autheticate and queue task
func (server *Server) CreateTask(w http.ResponseWriter, r *http.Request) {
	var rmq = RMQ{}
	rmq.PublishData("cc701a0e-0cfd-4160-b2e4-cadffd235b29")
	res := AsyncResponse{Status: "OK", RequestId: "cc701a0e-0cfd-4160-b2e4-cadffd235b29"}

	responses.JSON(w, http.StatusOK, res)

}
