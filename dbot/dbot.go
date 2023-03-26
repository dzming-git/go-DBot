package dbot

import (
	"net/http"
)

type DBot struct {
	serverCreatorPtr *ServerCreator
}

func NewDBot() *DBot {
	return &DBot{
		serverCreatorPtr: NewServerCreator(),
	}
}

func Start(dBotPtr *DBot) {
	InitServerCreator(dBotPtr.serverCreatorPtr)
	http.HandleFunc("/", CreatServer(dBotPtr.serverCreatorPtr))
	http.ListenAndServe(":5701", nil)
}
