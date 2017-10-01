package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"context"
)

type HTTP_ROUTER_PARAMS_CTX_KEY string

// helpers
func routerFuncWrapper(hf http.HandlerFunc) httprouter.Handle {
	return routerWrapper(http.Handler(hf))
}

func routerWrapper(h http.Handler) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		h.ServeHTTP(w, r.WithContext(withRouterParams(r.Context(), params)))
	}
}

func withRouterParams(ctx context.Context, params httprouter.Params) context.Context {
	return context.WithValue(ctx, HTTP_ROUTER_PARAMS_CTX_KEY("httprouter_params"), params)
}

func getRouterParams(ctx context.Context) httprouter.Params {
	params, ok := ctx.Value(HTTP_ROUTER_PARAMS_CTX_KEY("httprouter_params")).(httprouter.Params)

	if (!ok) {
		return nil
	}

	return params
}
