package handler

import "github.com/99designs/gqlgen/graphql/handler"

type HttpStr struct {
	srv *handler.Server
}

func NewHttp(sv *handler.Server) *HttpStr {
	return &HttpStr{srv: sv}
}
