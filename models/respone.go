package models

type ResponeMessage struct {
	Message string      ` json : "message"`
	Status  int         ` json : "status"`
	Data    interface{} ` json : "data"`
}
