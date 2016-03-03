package fast_http

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/fasthttpadaptor"
)

type Message struct {
	Id  int    `json:"id"`
	msg string `json:"string"`
}

var jsonHandle = func(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	messageList := make([]Message, 0)
	for i := 0; i < 100; i++ {
		messageList = append(messageList, Message{
			i,
			fmt.Sprintf("ezioruan%d", i),
		})
	}
	json.NewEncoder(w).Encode(&messageList)
}

var fastJsonHandle1 = fasthttpadaptor.NewFastHTTPHandlerFunc(jsonHandle)

var fastJsonHandle2 = func(ctx *fasthttp.RequestCtx) {
	ctx.SetContentType("application/json")
	messageList := make([]Message, 0)
	for i := 0; i < 100; i++ {
		messageList = append(messageList, Message{
			i,
			fmt.Sprintf("ezioruan%d", i),
		})
	}
	json.NewEncoder(ctx).Encode(&messageList)
}
