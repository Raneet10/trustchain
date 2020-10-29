package rest

import (
	"github.com/gorilla/mux"

	"github.com/cosmos/cosmos-sdk/client/context"
)

// RegisterRoutes registers trustchain-related REST handlers to a router
func RegisterRoutes(cliCtx context.CLIContext, r *mux.Router) {
	// this line is used by starport scaffolding # 1
	r.HandleFunc("/trustchain/promise", createPromiseHandler(cliCtx)).Methods("POST")
	r.HandleFunc("/trustchain/promise", listPromiseHandler(cliCtx, "trustchain")).Methods("GET")
	r.HandleFunc("/trustchain/promise/keeper/{keeper}", listPromiseForKeeperHandler(cliCtx, "trustchain")).Methods("GET")
	r.HandleFunc("/trustchain/promise/{key}", getPromiseHandler(cliCtx, "trustchain")).Methods("GET")
	r.HandleFunc("/trustchain/promise/confirm/{key}", confirmPromiseHandler(cliCtx)).Methods("POST")
	// r.HandleFunc("/trustchain/promise", setPromiseHandler(cliCtx)).Methods("PUT")
	// r.HandleFunc("/trustchain/promise", deletePromiseHandler(cliCtx)).Methods("DELETE")

}
