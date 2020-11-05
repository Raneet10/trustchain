package rest

import (
	"net/http"
	"strconv"
	"time"

	"github.com/cosmos/cosmos-sdk/client/context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	"github.com/zeno-bg/trustchain/x/trustchain/types"
)

// Used to not have an error if strconv is unused
var _ = strconv.Itoa(42)

type createPromiseRequest struct {
	BaseReq            rest.BaseReq `json:"base_req"`
	Creator            string       `json:"creator"`
	PromiseDescription string       `json:"promiseDescription"`
	PromiseKeeper      string       `json:"promiseKeeper"`
	Reward             string       `json:"reward"`
	Deadline           string       `json:"deadline"`
}

func createPromiseHandler(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req createPromiseRequest
		if !rest.ReadRESTReq(w, r, cliCtx.Codec, &req) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "failed to parse request")
			return
		}
		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}
		creator, err := sdk.AccAddressFromBech32(req.Creator)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		promiseKeeper, err := sdk.AccAddressFromBech32(req.PromiseKeeper)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		reward, err := sdk.ParseCoins(req.Reward)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		parsedPromiseDescription := req.PromiseDescription

		deadlineParsed, err := strconv.ParseInt(req.Deadline, 10, 64)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		deadline := time.Unix(deadlineParsed/1000, 0)

		msg := types.NewMsgCreatePromise(
			creator,
			parsedPromiseDescription,
			promiseKeeper,
			reward,
			deadline,
		)

		err = msg.ValidateBasic()
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		utils.WriteGenerateStdTxResponse(w, cliCtx, baseReq, []sdk.Msg{msg})
	}
}

type confirmPromiseRequest struct {
	BaseReq   rest.BaseReq `json:"base_req"`
	Creator   string       `json:"creator"`
	Fulfiller string       `json:"fulfiller"`
	Reward    string       `json:"reward"`
	ID        string       `json:"id"`
}

func confirmPromiseHandler(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req confirmPromiseRequest
		if !rest.ReadRESTReq(w, r, cliCtx.Codec, &req) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "failed to parse request")
			return
		}

		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}
		creator, err := sdk.AccAddressFromBech32(req.Creator)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		fulfiller, err := sdk.AccAddressFromBech32(req.Fulfiller)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		reward, err := sdk.ParseCoins(req.Reward)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}
		msg := types.NewMsgConfirmPromise(creator, fulfiller, reward, req.ID)

		err = msg.ValidateBasic()
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		utils.WriteGenerateStdTxResponse(w, cliCtx, baseReq, []sdk.Msg{msg})
	}
}

// type setPromiseRequest struct {
//     BaseReq            rest.BaseReq `json:"base_req"`
//     ID                 string       `json:"id"`
//     Creator            string       `json:"creator"`
//     PromiseDescription string       `json:"promiseDescription"`
//     PromiseKeeper      string       `json:"promiseKeeper"`
//     Reward             string       `json:"reward"`
// }

// func setPromiseHandler(cliCtx context.CLIContext) http.HandlerFunc {
//     return func(w http.ResponseWriter, r *http.Request) {
//         var req setPromiseRequest
//         if !rest.ReadRESTReq(w, r, cliCtx.Codec, &req) {
//             rest.WriteErrorResponse(w, http.StatusBadRequest, "failed to parse request")
//             return
//         }
//         baseReq := req.BaseReq.Sanitize()
//         if !baseReq.ValidateBasic(w) {
//             return
//         }
//         creator, err := sdk.AccAddressFromBech32(req.Creator)
//         if err != nil {
//             rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
//             return
//         }

//         parsedPromiseDescription := req.PromiseDescription

//         parsedPromiseKeeper := req.PromiseKeeper

//         parsedReward := req.Reward

//         msg := types.NewMsgSetPromise(
//             creator,
//             req.ID,
//             parsedPromiseDescription,
//             parsedPromiseKeeper,
//             parsedReward,
//         )

//         err = msg.ValidateBasic()
//         if err != nil {
//             rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
//             return
//         }

//         utils.WriteGenerateStdTxResponse(w, cliCtx, baseReq, []sdk.Msg{msg})
//     }
// }

// type deletePromiseRequest struct {
//     BaseReq rest.BaseReq `json:"base_req"`
//     Creator string `json:"creator"`
//     ID 		string `json:"id"`
// }

// func deletePromiseHandler(cliCtx context.CLIContext) http.HandlerFunc {
//     return func(w http.ResponseWriter, r *http.Request) {
//         var req deletePromiseRequest
//         if !rest.ReadRESTReq(w, r, cliCtx.Codec, &req) {
//             rest.WriteErrorResponse(w, http.StatusBadRequest, "failed to parse request")
//             return
//         }
//         baseReq := req.BaseReq.Sanitize()
//         if !baseReq.ValidateBasic(w) {
//             return
//         }
//         creator, err := sdk.AccAddressFromBech32(req.Creator)
//         if err != nil {
//             rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
//             return
//         }
//         msg := types.NewMsgDeletePromise(req.ID, creator)

//         err = msg.ValidateBasic()
//         if err != nil {
//             rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
//             return
//         }

//         utils.WriteGenerateStdTxResponse(w, cliCtx, baseReq, []sdk.Msg{msg})
//     }
// }
