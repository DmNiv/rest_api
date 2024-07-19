package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"rest_api/db"
	"rest_api/models"
)

func main() {
	db.InitDB()
	server := gin.Default() // Cria o servidor default

	server.GET("/events", getEvents)    // Cria uma chamada GET e chama a função getEvents
	server.POST("/events", createEvent) // Cria uma chamada Post e chama a função createEvents

	server.Run(":8080") // Roda o server no endereço localhost:8080
}

func getEvents(context *gin.Context) {
	events := models.GetAllEvents()
	context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)
	fmt.Println(event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Mensagem": "Não foi possível criar o evento!"})
		return
	}

	event.ID = 1
	event.UserID = 1

	event.Save()
	context.JSON(http.StatusCreated, gin.H{"Mensagem": "Evento criado com sucesso!", "Evento": event})
}
