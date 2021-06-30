package api

import (
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"neo3fura/lib/type/h160"
	"neo3fura/var/stderr"
)

// this function may be not supported any more, we only support address in the formart of script hash
func (me *T) GetVotersByCandidateAddress(args struct {
	CandidateAddress h160.T
	Limit  int64
	Skip   int64
	Filter map[string]interface{}
}, ret *json.RawMessage) error {
	if args.CandidateAddress.Valid() == false {
		return stderr.ErrInvalidArgs
	}
	r1, count, err := me.Data.Client.QueryAll(struct {
		Collection string
		Index      string
		Sort       bson.M
		Filter     bson.M
		Query      []string
		Limit      int64
		Skip       int64
	}{
		Collection: "Vote",
		Index:      "someIndex",
		Sort:       bson.M{},
		Filter:     bson.M{"candidate": args.CandidateAddress.Val()},
		Query:      []string{},
		Limit:      args.Limit,
		Skip:       args.Skip,
	}, ret)
	if err != nil {
		return err
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