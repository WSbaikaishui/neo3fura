package api

import (
	"encoding/json"
	"neophora/lib/type/addr"
	"neophora/lib/type/bins"
	"neophora/lib/type/h256"
	"neophora/lib/type/uintval"
	"neophora/var/stderr"
)

// GetUTXOByAddressTransactionHashOutputIndexInJSON ...
// as an example:
//
// ```
// $ curl https://example.neophora.io -d '{"jsonrpc": "2.0","id": 1,"method": "GetUTXOByAddressTransactionHashOutputIndexInJSON","params":{"Address":"AZ1QiX5nqgm8dsUY7iRyafKwmKnGP9bUhN" ,"TransactionHash":"4f6ab437d73a2b9f98f0571e3c55321f08175215642e30ef782d21374ac42323","OutputIndex":0}}'
// {"id":1,"result":{"address":"AZ1QiX5nqgm8dsUY7iRyafKwmKnGP9bUhN","asset":"0xc56f33fc6ecfcd0c225c4ab356fee59390af8560be0e930faebe74a6daff7c9b","value":"24"},"error":null}
// ```
func (me *T) GetUTXOByAddressTransactionHashOutputIndexInJSON(args struct {
	Address         addr.T
	TransactionHash h256.T
	OutputIndex     uintval.T
}, ret *json.RawMessage) error {
	if args.Address.Valid() == false {
		return stderr.ErrInvalidArgs
	}
	if args.TransactionHash.Valid() == false {
		return stderr.ErrInvalidArgs
	}
	if args.OutputIndex.Valid() == false {
		return stderr.ErrInvalidArgs
	}
	var result bins.T
	if err := me.Data.GetArgsInBins(struct {
		Target string
		Index  string
		Keys   []string
	}{
		Target: "bins.utx",
		Index:  "h160.act-h256.trx-uint.num",
		Keys:   []string{args.Address.H160(), args.TransactionHash.Val(), args.OutputIndex.Hex()},
	}, &result); err != nil {
		return err
	}
	if result.Valid() == false {
		return stderr.ErrNotFound
	}
	js, err := result.JSONViaUTXO()
	if err != nil {
		return stderr.ErrNotFound
	}
	*ret = js
	return nil
}