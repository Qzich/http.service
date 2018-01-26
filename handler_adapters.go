package http_service

import "net/http"

type HandlerAdapter func(handler http.Handler) http.Handler

type HandlerAdaptersList []HandlerAdapter

func (handlerAdaptersList *HandlerAdaptersList) Add(handlerAdapters ...HandlerAdapter) {
	for _, adapter := range handlerAdapters {
		*handlerAdaptersList = append(*handlerAdaptersList, adapter)
	}
}
