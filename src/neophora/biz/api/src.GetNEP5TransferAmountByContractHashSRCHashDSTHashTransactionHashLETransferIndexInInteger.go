package api

import (
	"encoding/json"
	"neophora/lib/type/bins"
	"neophora/lib/type/h160"
	"neophora/lib/type/h256"
	"neophora/lib/type/uintval"
	"neophora/var/stderr"
)

// GetNEP5TransferAmountByContractHashSRCHashDSTHashTransactionHashLETransferIndexInInteger ...
// as an example:
//
// ```
// $ curl https://example.neophora.io -d '{"jsonrpc": "2.0","id": 1,"method": "GetNEP5TransferAmountByContractHashSRCHashDSTHashTransactionHashLETransferIndexInInteger","params":{"ContractHash":"972e166ea1f8d3c3b14fd8766e7a0dad4084f9e8","SRCHash":"30c6c993a33e630056cfd05556361bf128b90fc8","DSTHash":"c7cb40c82c7920d2191f8924a6b73659dff3bfc3","TransactionHashLE":"9edfff14e7ac6d3c0bda2976de78fa774a119124cec7343dd66dad5cf4ffeddf","TransferIndex":0}}'
// {"id":1,"result":1650000000000,"error":null}
// ```
func (me *T) GetNEP5TransferAmountByContractHashSRCHashDSTHashTransactionHashLETransferIndexInInteger(args struct {
	ContractHash      h160.T
	SRCHash           h160.T
	DSTHash           h160.T
	TransactionHashLE h256.T
	TransferIndex     uintval.T
}, ret *json.RawMessage) error {
	if args.TransactionHashLE.Valid() == false {
		return stderr.ErrInvalidArgs
	}
	if args.TransferIndex.Valid() == false {
		return stderr.ErrInvalidArgs
	}
	if args.SRCHash.Valid() == false {
		return stderr.ErrInvalidArgs
	}
	if args.DSTHash.Valid() == false {
		return stderr.ErrInvalidArgs
	}
	if args.ContractHash.Valid() == false {
		return stderr.ErrInvalidArgs
	}
	var result bins.T
	if err := me.Data.GetArgsInBins(struct {
		Target string
		Index  string
		Keys   []string
	}{
		Target: "bigu.tsf",
		Index:  "h160.ctr-h160.src-h160.dst-h256.trx-uint.num",
		Keys:   []string{args.ContractHash.Val(), args.SRCHash.Val(), args.DSTHash.Val(), args.TransactionHashLE.RevVal(), args.TransferIndex.Hex()},
	}, &result); err != nil {
		return err
	}
	if result.Valid() == false {
		return stderr.ErrNotFound
	}
	*ret = json.RawMessage(result.BigString())
	return nil
}