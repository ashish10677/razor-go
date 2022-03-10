package cmd

import (
	"errors"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"math/big"
	"razor/core"
	"razor/core/types"
	"razor/logger"
	"razor/pkg/bindings"
	"razor/utils"
)

var initiateWithdrawCmd = &cobra.Command{
	Use:   "initiateWithdraw",
	Short: "initiateWithdraw for your razors once you've unstaked",
	Long: `initiateWithdraw command can be used once the user has unstaked their token and the withdraw period is upon them.

Example:
  ./razor initiateWithdraw --address 0x5a0b54d5dc17e0aadc383d2db43b0a0d3e029c4c --stakerId 1
`,
	Run: func(cmd *cobra.Command, args []string) {
		cmdUtils.ExecuteInitiateWithdraw(cmd.Flags())
	},
}

func (*UtilsStruct) ExecuteInitiateWithdraw(flagSet *pflag.FlagSet) {
	address, err := flagSetUtils.GetStringAddress(flagSet)
	utils.CheckError("Error in getting address: ", err)

	logger.Address = address

	config, err := cmdUtils.GetConfigData()
	utils.CheckError("Error in getting config: ", err)

	password := razorUtils.AssignPassword(flagSet)

	client := razorUtils.ConnectToClient(config.Provider)

	razorUtils.CheckEthBalanceIsZero(client, address)

	stakerId, err := razorUtils.AssignStakerId(flagSet, client, address)
	utils.CheckError("Error in fetching stakerId:  ", err)

	txn, err := cmdUtils.HandleUnstakeLock(client, types.Account{
		Address:  address,
		Password: password,
	}, config, stakerId)

	utils.CheckError("InitiateWithdraw error: ", err)
	if txn != core.NilHash {
		razorUtils.WaitForBlockCompletion(client, txn.String())
	}
}

func (*UtilsStruct) HandleUnstakeLock(client *ethclient.Client, account types.Account, configurations types.Configurations, stakerId uint32) (common.Hash, error) {

	unstakeLock, err := razorUtils.GetLock(client, account.Address, stakerId, 0)
	if err != nil {
		log.Error("Error in fetching unstakeLock")
		return core.NilHash, err
	}

	if unstakeLock.UnlockAfter.Cmp(big.NewInt(0)) == 0 {
		log.Error("Unstake command not called before initiating withdrawal!")
		return core.NilHash, errors.New("unstake Razors before withdrawing")
	}

	withdrawInitiationPeriod, err := razorUtils.GetWithdrawInitiationPeriod(client)
	if err != nil {
		log.Error("Error in fetching withdraw release period")
		return core.NilHash, err
	}

	withdrawBefore := big.NewInt(0).Add(unstakeLock.UnlockAfter, big.NewInt(int64(withdrawInitiationPeriod)))
	epoch, err := razorUtils.GetEpoch(client)
	if err != nil {
		log.Error("Error in fetching epoch")
		return core.NilHash, err
	}

	if big.NewInt(int64(epoch)).Cmp(withdrawBefore) > 0 {
		log.Info("Withdraw initiation period has passed. Cannot withdraw now, please reset the unstakeLock!")
		return core.NilHash, nil
	}

	txnArgs := types.TransactionOptions{
		Client:          client,
		Password:        account.Password,
		AccountAddress:  account.Address,
		ChainId:         core.ChainId,
		Config:          configurations,
		ContractAddress: core.StakeManagerAddress,
		MethodName:      "initiateWithdrawCmd",
		ABI:             bindings.StakeManagerABI,
		Parameters:      []interface{}{stakerId},
	}
	txnOpts := razorUtils.GetTxnOpts(txnArgs)

	if big.NewInt(int64(epoch)).Cmp(unstakeLock.UnlockAfter) >= 0 && big.NewInt(int64(epoch)).Cmp(withdrawBefore) <= 0 {
		return cmdUtils.InitiateWithdraw(client, txnOpts, stakerId)
	}
	return core.NilHash, errors.New("unstakeLock period not over yet! Please try after some time")
}

func (*UtilsStruct) InitiateWithdraw(client *ethclient.Client, txnOpts *bind.TransactOpts, stakerId uint32) (common.Hash, error) {
	log.Info("Initiating withdrawal of funds...")

	txn, err := stakeManagerUtils.InitiateWithdraw(client, txnOpts, stakerId)
	if err != nil {
		log.Error("Error in initiating withdrawal of funds")
		return core.NilHash, err
	}

	log.Info("Txn Hash: ", transactionUtils.Hash(txn))

	return transactionUtils.Hash(txn), nil
}

func init() {
	rootCmd.AddCommand(initiateWithdrawCmd)

	var (
		Address  string
		Password string
		StakerId uint32
	)

	initiateWithdrawCmd.Flags().StringVarP(&Address, "address", "a", "", "address of the user")
	initiateWithdrawCmd.Flags().StringVarP(&Password, "password", "", "", "password path of user to protect the keystore")
	initiateWithdrawCmd.Flags().Uint32VarP(&StakerId, "stakerId", "", 0, "password path of user to protect the keystore")

	addrErr := initiateWithdrawCmd.MarkFlagRequired("address")
	utils.CheckError("Address error: ", addrErr)
}