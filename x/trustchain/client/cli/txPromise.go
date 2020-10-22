package cli

import (
	"bufio"
    
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	"github.com/zeno-bg/trustchain/x/trustchain/types"
)

func GetCmdCreatePromise(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "create-promise [promiseDescription] [promiseKeeper] [reward]",
		Short: "Creates a new promise",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			argsPromiseDescription := string(args[0] )
			argsPromiseKeeper := string(args[1] )
			argsReward := string(args[2] )
			
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))
			msg := types.NewMsgCreatePromise(cliCtx.GetFromAddress(), string(argsPromiseDescription), string(argsPromiseKeeper), string(argsReward))
			err := msg.ValidateBasic()
			if err != nil {
				return err
			}
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}


func GetCmdSetPromise(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "set-promise [id]  [promiseDescription] [promiseKeeper] [reward]",
		Short: "Set a new promise",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) error {
			id := args[0]
			argsPromiseDescription := string(args[1])
			argsPromiseKeeper := string(args[2])
			argsReward := string(args[3])
			
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))
			msg := types.NewMsgSetPromise(cliCtx.GetFromAddress(), id, string(argsPromiseDescription), string(argsPromiseKeeper), string(argsReward))
			err := msg.ValidateBasic()
			if err != nil {
				return err
			}
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}

func GetCmdDeletePromise(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "delete-promise [id]",
		Short: "Delete a new promise by ID",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {

			cliCtx := context.NewCLIContext().WithCodec(cdc)
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))

			msg := types.NewMsgDeletePromise(args[0], cliCtx.GetFromAddress())
			err := msg.ValidateBasic()
			if err != nil {
				return err
			}
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}
