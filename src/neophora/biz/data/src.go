package data

import (
	"encoding/hex"
	"neophora/lib/cli"
	"neophora/var/stderr"
	"net/url"
	"path"
)

// T ...
type T struct {
	Client *cli.T
}

// Get ...
func (me *T) Get(args struct {
	Key []byte
}, ret *[]byte) error {
	if err := me.Client.Call("T.Get", args, ret); err != nil {
		return stderr.ErrUnknown
	}
	return nil
}

// GetStr ...
func (me *T) GetStr(args struct {
	Key string
}, ret *[]byte) error {
	return me.Get(struct{ Key []byte }{[]byte(args.Key)}, ret)
}

// GetArgs ...
func (me *T) GetArgs(args struct {
	Target string
	Index  string
	Keys   []string
}, ret *[]byte) error {
	uri := &url.URL{
		Scheme: args.Target,
		Host:   args.Index,
		Path:   path.Join(args.Keys...),
	}
	return me.GetStr(struct{ Key string }{uri.String()}, ret)
}

// GetLastKey ...
func (me *T) GetLastKey(args struct {
	Target string
	Index  string
	Keys   []string
}, ret *[]byte) error {
	if len(args.Keys) == 0 {
		return stderr.ErrInvalidArgs
	}
	uri := &url.URL{
		Scheme: args.Target,
		Host:   args.Index,
		Path:   path.Join(args.Keys...),
	}
	key := []byte(uri.String())
	parameter := struct {
		Key    []byte
		Prefix uint
	}{
		Key:    key,
		Prefix: uint(len(key) - len(args.Keys[len(args.Keys)-1])),
	}
	if err := me.Client.Call("T.GetLastKey", parameter, ret); err != nil {
		return stderr.ErrUnknown
	}
	return nil
}

// GetLastVal ...
func (me *T) GetLastVal(args struct {
	Target string
	Index  string
	Keys   []string
}, ret *[]byte) error {
	if len(args.Keys) == 0 {
		return stderr.ErrInvalidArgs
	}
	uri := &url.URL{
		Scheme: args.Target,
		Host:   args.Index,
		Path:   path.Join(args.Keys...),
	}
	key := []byte(uri.String())
	parameter := struct {
		Key    []byte
		Prefix uint
	}{
		Key:    key,
		Prefix: uint(len(key) - len(args.Keys[len(args.Keys)-1])),
	}
	if err := me.Client.Call("T.GetLastVal", parameter, ret); err != nil {
		return stderr.ErrUnknown
	}
	return nil
}

// GetLastestUint64Key ...
func (me *T) GetLastestUint64Key(args struct {
	Target string
	Index  string
	Keys   []string
}, ret *[]byte) error {
	args.Keys = append(args.Keys, "ffffffffffffffff")
	return me.GetLastKey(args, ret)
}

// GetLastestUint64Val ...
func (me *T) GetLastestUint64Val(args struct {
	Target string
	Index  string
	Keys   []string
}, ret *[]byte) error {
	args.Keys = append(args.Keys, "ffffffffffffffff")
	return me.GetLastVal(args, ret)
}

// GetInHex ...
func (me *T) GetInHex(args struct {
	Key []byte
}, ret *string) error {
	var result []byte
	if err := me.Get(args, &result); err != nil {
		return err
	}
	*ret = hex.EncodeToString(result)
	return nil
}

// GetStrHex ...
func (me *T) GetStrHex(args struct {
	Key string
}, ret *string) error {
	var result []byte
	if err := me.GetStr(args, &result); err != nil {
		return err
	}
	*ret = hex.EncodeToString(result)
	return nil
}

// GetArgsInHex ...
func (me *T) GetArgsInHex(args struct {
	Target string
	Index  string
	Keys   []string
}, ret *string) error {
	var result []byte
	if err := me.GetArgs(args, &result); err != nil {
		return err
	}
	*ret = hex.EncodeToString(result)
	return nil
}

// GetLastKeyInStr ...
func (me *T) GetLastKeyInStr(args struct {
	Target string
	Index  string
	Keys   []string
}, ret *string) error {
	var result []byte
	if err := me.GetLastKey(args, &result); err != nil {
		return err
	}
	*ret = string(result)
	return nil
}

// GetLastKeyInURL ...
func (me *T) GetLastKeyInURL(args struct {
	Target string
	Index  string
	Keys   []string
}, ret *url.URL) error {
	var result string
	if err := me.GetLastKeyInStr(args, &result); err != nil {
		return err
	}
	uri, err := url.Parse(result)
	if err != nil {
		return stderr.ErrNotFound
	}
	*ret = *uri
	return nil
}

// GetLastValInHex ...
func (me *T) GetLastValInHex(args struct {
	Target string
	Index  string
	Keys   []string
}, ret *string) error {
	var result []byte
	if err := me.GetLastVal(args, &result); err != nil {
		return err
	}
	*ret = hex.EncodeToString(result)
	return nil
}

// GetLastestUint64KeyInStr ...
func (me *T) GetLastestUint64KeyInStr(args struct {
	Target string
	Index  string
	Keys   []string
}, ret *string) error {
	var result []byte
	if err := me.GetLastestUint64Key(args, &result); err != nil {
		return err
	}
	*ret = string(result)
	return nil
}

// GetLastestUint64KeyInURL ...
func (me *T) GetLastestUint64KeyInURL(args struct {
	Target string
	Index  string
	Keys   []string
}, ret *url.URL) error {
	var result string
	if err := me.GetLastestUint64KeyInStr(args, &result); err != nil {
		return err
	}
	uri, err := url.Parse(result)
	if err != nil {
		return stderr.ErrNotFound
	}
	*ret = *uri
	return nil
}

// GetLastestUint64ValInHex ...
func (me *T) GetLastestUint64ValInHex(args struct {
	Target string
	Index  string
	Keys   []string
}, ret *string) error {
	var result []byte
	if err := me.GetLastestUint64Val(args, &result); err != nil {
		return err
	}
	*ret = hex.EncodeToString(result)
	return nil
}
