package api

import (
	"encoding/json"
	"neophora/lib/type/bins"
	"neophora/lib/type/h160"
	"neophora/lib/type/h256"
	"neophora/lib/type/uintval"
	"neophora/var/stderr"
)

// GetNEP5TransferAmountByContractHashLEDSTHashLESRCHashLETransactionHashTransferIndexInInteger ...
// as an example:
//
// ```
// $ curl https://example.neophora.io -d '{"jsonrpc": "2.0","id": 1,"method": "GetNEP5TransferAmountByContractHashLEDSTHashLESRCHashLETransactionHashTransferIndexInInteger","params":{"ContractHashLE":"e8f98440ad0d7a6e76d84fb1c3d3f8a16e162e97","DSTHashLE":"c3bff3df5936b7a624891f19d220792cc840cbc7","SRCHashLE":"c80fb928f11b365655d0cf5600633ea393c9c630","TransactionHash":"dfedfff45cad6dd63d34c7ce2491114a77fa78de7629da0b3c6dace714ffdf9e","TransferIndex":0}}'
// {"id":1,"result":1650000000000,"error":null}
// ```
func (me *T) GetNEP5TransferAmountByContractHashLEDSTHashLESRCHashLETransactionHashTransferIndexInInteger(args struct {
	ContractHashLE  h160.T
	DSTHashLE       h160.T
	SRCHashLE       h160.T
	TransactionHash h256.T
	TransferIndex   uintval.T
}, ret *json.RawMessage) error {
	if args.TransactionHash.Valid() == false {
		return stderr.ErrInvalidArgs
	}
	if args.TransferIndex.Valid() == false {
		return stderr.ErrInvalidArgs
	}
	if args.SRCHashLE.Valid() == false {
		return stderr.ErrInvalidArgs
	}
	if args.DSTHashLE.Valid() == false {
		return stderr.ErrInvalidArgs
	}
	if args.ContractHashLE.Valid() == false {
		return stderr.ErrInvalidArgs
	}
	var result bins.T
	if err := me.Data.GetArgsInBins(struct {
		Target string
		Index  string
		Keys   []string
	}{
		Target: "bigu.tsf",
		Index:  "h160.ctr-h160.dst-h160.src-h256.trx-uint.num",
		Keys:   []string{args.ContractHashLE.RevVal(), args.DSTHashLE.RevVal(), args.SRCHashLE.RevVal(), args.TransactionHash.Val(), args.TransferIndex.Hex()},
	}, &result); err != nil {
		return err
	}
	if result.Valid() == false {
		return stderr.ErrNotFound
	}
	*ret = json.RawMessage(result.BigString())
	return nil
}