package api

import (
	"encoding/json"
	"neophora/lib/transex"
)

// GetHeaderByHashInJSON ...
func (me *T) GetHeaderByHashInJSON(args struct {
	Hash string
}, ret *json.RawMessage) error {
	var result []byte
	var tr transex.T
	if err := me.Data.GetArgs(struct {
		Target string
		Index  string
		Keys   []string
	}{
		Target: "header",
		Index:  "hash",
		Keys:   []string{args.Hash},
	}, &result); err != nil {
		return err
	}
	tr.V = result
	if err := tr.BytesToJSONViaHeader(); err != nil {
		return err
	}
	*ret = tr.V.(json.RawMessage)
	return nil
}