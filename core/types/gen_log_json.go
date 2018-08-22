// Code generated by github.com/fjl/gencoDEWH. DO NOT EDIT.

package types

import (
	"encoding/json"
	"errors"

	"github.com/DEWH/go-DEWH/common"
	"github.com/DEWH/go-DEWH/common/hexutil"
)

var _ = (*logMarshaling)(nil)

func (l Log) MarshalJSON() ([]byte, error) {
	type Log struct {
		Address     common.Address `json:"address" gencoDEWH:"required"`
		Topics      []common.Hash  `json:"topics" gencoDEWH:"required"`
		Data        hexutil.Bytes  `json:"data" gencoDEWH:"required"`
		BlockNumber hexutil.Uint64 `json:"blockNumber"`
		TxHash      common.Hash    `json:"transactionHash" gencoDEWH:"required"`
		TxIndex     hexutil.Uint   `json:"transactionIndex" gencoDEWH:"required"`
		BlockHash   common.Hash    `json:"blockHash"`
		Index       hexutil.Uint   `json:"logIndex" gencoDEWH:"required"`
		Removed     bool           `json:"removed"`
	}
	var enc Log
	enc.Address = l.Address
	enc.Topics = l.Topics
	enc.Data = l.Data
	enc.BlockNumber = hexutil.Uint64(l.BlockNumber)
	enc.TxHash = l.TxHash
	enc.TxIndex = hexutil.Uint(l.TxIndex)
	enc.BlockHash = l.BlockHash
	enc.Index = hexutil.Uint(l.Index)
	enc.Removed = l.Removed
	return json.Marshal(&enc)
}

func (l *Log) UnmarshalJSON(input []byte) error {
	type Log struct {
		Address     *common.Address `json:"address" gencoDEWH:"required"`
		Topics      []common.Hash   `json:"topics" gencoDEWH:"required"`
		Data        *hexutil.Bytes  `json:"data" gencoDEWH:"required"`
		BlockNumber *hexutil.Uint64 `json:"blockNumber"`
		TxHash      *common.Hash    `json:"transactionHash" gencoDEWH:"required"`
		TxIndex     *hexutil.Uint   `json:"transactionIndex" gencoDEWH:"required"`
		BlockHash   *common.Hash    `json:"blockHash"`
		Index       *hexutil.Uint   `json:"logIndex" gencoDEWH:"required"`
		Removed     *bool           `json:"removed"`
	}
	var DEWH Log
	if err := json.Unmarshal(input, &DEWH); err != nil {
		return err
	}
	if DEWH.Address == nil {
		return errors.New("missing required field 'address' for Log")
	}
	l.Address = *DEWH.Address
	if DEWH.Topics == nil {
		return errors.New("missing required field 'topics' for Log")
	}
	l.Topics = DEWH.Topics
	if DEWH.Data == nil {
		return errors.New("missing required field 'data' for Log")
	}
	l.Data = *DEWH.Data
	if DEWH.BlockNumber != nil {
		l.BlockNumber = uint64(*DEWH.BlockNumber)
	}
	if DEWH.TxHash == nil {
		return errors.New("missing required field 'transactionHash' for Log")
	}
	l.TxHash = *DEWH.TxHash
	if DEWH.TxIndex == nil {
		return errors.New("missing required field 'transactionIndex' for Log")
	}
	l.TxIndex = uint(*DEWH.TxIndex)
	if DEWH.BlockHash != nil {
		l.BlockHash = *DEWH.BlockHash
	}
	if DEWH.Index == nil {
		return errors.New("missing required field 'logIndex' for Log")
	}
	l.Index = uint(*DEWH.Index)
	if DEWH.Removed != nil {
		l.Removed = *DEWH.Removed
	}
	return nil
}
