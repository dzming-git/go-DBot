package dbot

import (
	"encoding/json"
	"net/http"
	"strconv"

	"dzmicro/app/message_handler"
)

type ServerCreator struct {
	MessageHandler func(string, []string)
}

func NewServerCreator() *ServerCreator {
	return &ServerCreator{}
}

type RequestData struct {
	PostType    string `json:"post_type"`
	MessageType string `json:"message_type"`
	Time        int64  `json:"time"`
	SelfID      int64  `json:"self_id"`
	SubType     string `json:"sub_type"`
	Message     string `json:"message"`
	RawMessage  string `json:"raw_message"`
	Font        int    `json:"font"`
	GroupID     *int64 `json:"group_id,omitempty"`
	Sender      struct {
		Age      int    `json:"age"`
		Nickname string `json:"nickname"`
		Sex      string `json:"sex"`
		UserID   int64  `json:"user_id"`
	} `json:"sender"`
	MessageID int64 `json:"message_id"`
	UserID    int64 `json:"user_id"`
	TargetID  int64 `json:"target_id"`
}

func InitServerCreator(serverCreatorPtr *ServerCreator) {
	serverCreatorPtr.MessageHandler = message_handler.MessageHandler
}

func CreatServer(serverCreatorPtr *ServerCreator) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/":
			if r.Method != http.MethodPost {
				http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
				return
			}

			// 从请求中解码JSON数据
			var requestData RequestData
			err := json.NewDecoder(r.Body).Decode(&requestData)
			if err != nil {
				http.Error(w, "Bad request", http.StatusBadRequest)
				return
			}
			var gID string
			gIDPtr := requestData.GroupID
			if gIDPtr != nil {
				gID = strconv.FormatInt(*gIDPtr, 10)
			} else {
				gID = "None"
			}
			qID := strconv.FormatInt(requestData.Sender.UserID, 10)
			sourceID := []string{gID, qID}
			serverCreatorPtr.MessageHandler(requestData.Message, sourceID)
		}
	}
}
