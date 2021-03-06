package handlers

import (
	"context"
	"github.com/llygcd/block-compensation/pkg/pool"
	"github.com/llygcd/block-compensation/utils"
	"github.com/sirupsen/logrus"
	ctypes "github.com/tendermint/tendermint/rpc/core/types"
	"github.com/tendermint/tendermint/types"
	"time"
)

type chanTxResult struct {
	TxHash   string
	TxResult *ctypes.ResultTx
	Err      error
}

// parse tx with more goroutine concurrency
func handleTxResult(client *pool.Client, block *types.Block) map[string]chanTxResult {
	if _conf == nil {
		logrus.Error("InitRouter don't work")
	}
	if _conf.Server.ThreadNumParseTx <= 0 {
		_conf.Server.ThreadNumParseTx = 1
	}

	chanParseTxLimit := make(chan bool, _conf.Server.ThreadNumParseTx)
	chanRes := make(chan chanTxResult, len(block.Txs))
	for _, v := range block.Txs {
		chanParseTxLimit <- true
		// parse txReult with more goroutine concurrency
		go getTxResult(client, v, block.Height, chanParseTxLimit, chanRes)
	}
	txRetMap := make(map[string]chanTxResult, len(block.Txs))
	for i := 0; i < len(block.Txs); i++ {
		chanValue := <-chanRes
		txRetMap[chanValue.TxHash] = chanValue

	}
	return txRetMap
}

func getTxResult(c *pool.Client, txBytes types.Tx, height int64, chanLimit chan bool, chanRes chan chanTxResult) {
	defer func() {
		if r := recover(); r != nil {
			logrus.Error("execute getTxResult fail ")
		}
		<-chanLimit
	}()
	txHash := utils.BuildHex(txBytes.Hash())
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	txResult, err := c.Tx(ctx, txBytes.Hash(), false)
	if err != nil {
		time.Sleep(1 * time.Second)
		if v, err := c.Tx(ctx, txBytes.Hash(), false); err != nil {
			logrus.Error(utils.ConvertErr(height, txHash, "TxResult", err).Error())
		} else {
			txResult = v
		}
	}

	if txResult == nil {
		chanRes <- chanTxResult{Err: err}
		return
	}
	ret := chanTxResult{
		TxHash:   txResult.Hash.String(),
		TxResult: txResult,
		Err:      err,
	}
	chanRes <- ret

	return
}
