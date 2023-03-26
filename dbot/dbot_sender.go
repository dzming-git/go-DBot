package dbot

import (
	"fmt"
	"net/http"
)

func SendMessageToCqhttp(message string, sourceID []string) {
	gid := sourceID[0]
	qid := sourceID[1]
	urlencoded_message := urlencoding_message(message)

	url := ""

	if gid == "None" {
		url = "http://127.0.0.1:5700/send_private_msg?user_id=%s&message=%s"
		url = fmt.Sprintf(url, qid, urlencoded_message)
	} else {
		url = "http://127.0.0.1:5700/send_group_msg?group_id=%s&message=%s"
		url = fmt.Sprintf(url, gid, urlencoded_message)
	}

	resp, err := http.Get(url)
	if err != nil {
		// 处理错误
	}
	defer resp.Body.Close()

	// 处理响应
}
