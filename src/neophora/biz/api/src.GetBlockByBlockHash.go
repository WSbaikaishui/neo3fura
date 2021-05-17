package api

import (
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"neophora/lib/type/h256"
	"neophora/var/stderr"
)

func (me *T) GetBlockByBlockHash(args struct {
	BlockHash h256.T
}, ret *json.RawMessage) error {
	if args.BlockHash.Valid() == false {
		return stderr.ErrInvalidArgs
	}
	_, err := me.Data.Client.QueryOne(struct {
		Collection string
		Index      string
		Sort       bson.M
		Filter     bson.M
		Query      []string
	}{
		Collection: "Block",
		Index:      "someIndex",
		Sort:       bson.M{},
		Filter:     bson.M{"hash": args.BlockHash},
		Query:      []string{},
	}, ret)
	if err != nil {
		return err
	}
	return nil
}
