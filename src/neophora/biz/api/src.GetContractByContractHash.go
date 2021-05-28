package api

import (
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"neophora/lib/type/h160"
	"neophora/var/stderr"
)

func (me *T) GetContractByContractHash(args struct{
	ContractHash h160.T
}, ret *json.RawMessage) error {
	if args.ContractHash.Valid() == false {
		return stderr.ErrInvalidArgs
	}
	_, err := me.Data.Client.QueryOne(struct {
		Collection string
		Index      string
		Sort       bson.M
		Filter     bson.M
		Query      []string
	}{
		Collection: "Contract",
		Index:      "someIndex",
		Sort:       bson.M{"_id": -1},
		Filter:     bson.M{"hash":args.ContractHash},
		Query:      []string{},
	}, ret)
	if err != nil {
		return err
	}
	return nil
}
