package api

import (
	"neophora/lib/type/addr"
	"neophora/lib/type/bins"
	"neophora/lib/type/h160"
	"neophora/lib/type/h256"
	"neophora/lib/type/uintval"
	"neophora/var/stderr"
)

// GetNEP5TransferAmountByContractHashLESRCAddressDSTAddressTransactionHashTransferIndexInUint64 ...
// as an example:
//
// ```
// TODO
// ```
func (me *T) GetNEP5TransferAmountByContractHashLESRCAddressDSTAddressTransactionHashTransferIndexInUint64(args struct {
	ContractHashLE  h160.T
	SRCAddress      addr.T
	DSTAddress      addr.T
	TransactionHash h256.T
	TransferIndex   uintval.T
}, ret *uint64) error {
	if args.TransactionHash.Valid() == false {
		return stderr.ErrInvalidArgs
	}
	if args.TransferIndex.Valid() == false {
		return stderr.ErrInvalidArgs
	}
	if args.SRCAddress.Valid() == false {
		return stderr.ErrInvalidArgs
	}
	if args.DSTAddress.Valid() == false {
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
		Index:  "h160.ctr-h160.src-h160.dst-h256.trx-uint.num",
		Keys:   []string{args.ContractHashLE.RevVal(), args.SRCAddress.H160(), args.DSTAddress.H160(), args.TransactionHash.Val(), args.TransferIndex.Hex()},
	}, &result); err != nil {
		return err
	}
	if result.Valid() == false {
		return stderr.ErrNotFound
	}
	*ret = result.Uint64()
	return nil
}
