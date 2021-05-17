package api

import (
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"neophora/lib/type/uintval"
	"neophora/var/stderr"
)

func (me *T) GetBlockHeaderByBlockHeight(args struct {
	BlockHeight uintval.T
}, ret *json.RawMessage) error {
	if args.BlockHeight.Valid() == false {
		return stderr.ErrInvalidArgs
	}
	_, err := me.Data.Client.QueryOne(struct {
		Collection string
		Index      string
		Sort       bson.M
		Filter     bson.M
		Query      []string
	}{
		Collection: "Header",
		Index:      "someIndex",
		Sort:       bson.M{},
		Filter:     bson.M{"index": args.BlockHeight},
		Query:      []string{},
	}, ret)
	if err != nil {
		return err
	}
	return nil
}
