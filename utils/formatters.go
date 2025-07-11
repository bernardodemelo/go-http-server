package utils

import (
	"encoding/json"
	"net/http"
)

type ResponsePipeline struct {
	w http.ResponseWriter
}

func ResponsePipe(w http.ResponseWriter) *ResponsePipeline {
	return &ResponsePipeline{w: w}
}

func (r *ResponsePipeline) JSONHeaders(status int) *ResponsePipeline {
	r.w.WriteHeader(status)
	r.w.Header().Set("Content-Type", "application/json")
	return r
}

func (r *ResponsePipeline) JSON(data any) {
	json.NewEncoder(r.w).Encode(data)
}
