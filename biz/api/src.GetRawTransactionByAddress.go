package api

import (
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"neo3fura/lib/type/h160"
	"neo3fura/lib/type/h256"
	"neo3fura/var/stderr"
)

func (me *T) GetRawTransactionByAddress(args struct {
	Address h160.T
	Limit   int64
	Skip    int64
	Filter  map[string]interface{}
}, ret *json.RawMessage) error {
	if args.Address.Valid() == false {
		return stderr.ErrInvalidArgs
	}
	r1, count, err := me.Client.QueryAll(struct {
		Collection string
		Index      string
		Sort       bson.M
		Filter     bson.M
		Query      []string
		Limit      int64
		Skip       int64
	}{
		Collection: "Transaction",
		Index:      "GetRawTransactionByAddress",
		Sort:       bson.M{},
		Filter:     bson.M{"sender": args.Address.TransferAddress()},
		Query:      []string{},
		Limit:      args.Limit,
		Skip:       args.Skip,
	}, ret)
	if err != nil {
		return err
	}
	var raw1 map[string]interface{}
	for _, item := range r1 {
		err = me.GetVmStateByTransactionHash(struct {
			TransactionHash h256.T
			Filter          map[string]interface{}
			Raw             *map[string]interface{}
		}{
			TransactionHash: h256.T(fmt.Sprint(item["hash"])),
			Filter:          nil,
			Raw:             &raw1,
		}, ret)
		if err != nil {
			return err
		}
		item["vmstate"] = raw1["vmstate"].(string)
	}

	r2, err := me.FilterArrayAndAppendCount(r1, count, args.Filter)
	if err != nil {
		return err
	}
	r, err := json.Marshal(r2)
	if err != nil {
		return err
	}
	*ret = json.RawMessage(r)
	return nil
}
