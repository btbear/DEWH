// Code generated by github.com/fjl/gencoDEWH. DO NOT EDIT.

package types

import (
	"encoding/json"
	"errors"
	"math/big"

	"github.com/DEWH/go-DEWH/common"
	"github.com/DEWH/go-DEWH/common/hexutil"
)

var _ = (*headerMarshaling)(nil)

func (h Header) MarshalJSON() ([]byte, error) {
	type Header struct {
		ParentHash  common.Hash    `json:"parentHash"       gencoDEWH:"required"`
		UncleHash   common.Hash    `json:"sha3Uncles"       gencoDEWH:"required"`
		Coinbase    common.Address `json:"miner"            gencoDEWH:"required"`
		Root        common.Hash    `json:"stateRoot"        gencoDEWH:"required"`
		TxHash      common.Hash    `json:"transactionsRoot" gencoDEWH:"required"`
		ReceiptHash common.Hash    `json:"receiptsRoot"     gencoDEWH:"required"`
		Bloom       Bloom          `json:"logsBloom"        gencoDEWH:"required"`
		Difficulty  *hexutil.Big   `json:"difficulty"       gencoDEWH:"required"`
		Number      *hexutil.Big   `json:"number"           gencoDEWH:"required"`
		GasLimit    hexutil.Uint64 `json:"gasLimit"         gencoDEWH:"required"`
		GasUsed     hexutil.Uint64 `json:"gasUsed"          gencoDEWH:"required"`
		Time        *hexutil.Big   `json:"timestamp"        gencoDEWH:"required"`
		Extra       hexutil.Bytes  `json:"extraData"        gencoDEWH:"required"`
		MixDigest   common.Hash    `json:"mixHash"          gencoDEWH:"required"`
		Nonce       BlockNonce     `json:"nonce"            gencoDEWH:"required"`
		Hash        common.Hash    `json:"hash"`
	}
	var enc Header
	enc.ParentHash = h.ParentHash
	enc.UncleHash = h.UncleHash
	enc.Coinbase = h.Coinbase
	enc.Root = h.Root
	enc.TxHash = h.TxHash
	enc.ReceiptHash = h.ReceiptHash
	enc.Bloom = h.Bloom
	enc.Difficulty = (*hexutil.Big)(h.Difficulty)
	enc.Number = (*hexutil.Big)(h.Number)
	enc.GasLimit = hexutil.Uint64(h.GasLimit)
	enc.GasUsed = hexutil.Uint64(h.GasUsed)
	enc.Time = (*hexutil.Big)(h.Time)
	enc.Extra = h.Extra
	enc.MixDigest = h.MixDigest
	enc.Nonce = h.Nonce
	enc.Hash = h.Hash()
	return json.Marshal(&enc)
}

func (h *Header) UnmarshalJSON(input []byte) error {
	type Header struct {
		ParentHash  *common.Hash    `json:"parentHash"       gencoDEWH:"required"`
		UncleHash   *common.Hash    `json:"sha3Uncles"       gencoDEWH:"required"`
		Coinbase    *common.Address `json:"miner"            gencoDEWH:"required"`
		Root        *common.Hash    `json:"stateRoot"        gencoDEWH:"required"`
		TxHash      *common.Hash    `json:"transactionsRoot" gencoDEWH:"required"`
		ReceiptHash *common.Hash    `json:"receiptsRoot"     gencoDEWH:"required"`
		Bloom       *Bloom          `json:"logsBloom"        gencoDEWH:"required"`
		Difficulty  *hexutil.Big    `json:"difficulty"       gencoDEWH:"required"`
		Number      *hexutil.Big    `json:"number"           gencoDEWH:"required"`
		GasLimit    *hexutil.Uint64 `json:"gasLimit"         gencoDEWH:"required"`
		GasUsed     *hexutil.Uint64 `json:"gasUsed"          gencoDEWH:"required"`
		Time        *hexutil.Big    `json:"timestamp"        gencoDEWH:"required"`
		Extra       *hexutil.Bytes  `json:"extraData"        gencoDEWH:"required"`
		MixDigest   *common.Hash    `json:"mixHash"          gencoDEWH:"required"`
		Nonce       *BlockNonce     `json:"nonce"            gencoDEWH:"required"`
	}
	var DEWH Header
	if err := json.Unmarshal(input, &DEWH); err != nil {
		return err
	}
	if DEWH.ParentHash == nil {
		return errors.New("missing required field 'parentHash' for Header")
	}
	h.ParentHash = *DEWH.ParentHash
	if DEWH.UncleHash == nil {
		return errors.New("missing required field 'sha3Uncles' for Header")
	}
	h.UncleHash = *DEWH.UncleHash
	if DEWH.Coinbase == nil {
		return errors.New("missing required field 'miner' for Header")
	}
	h.Coinbase = *DEWH.Coinbase
	if DEWH.Root == nil {
		return errors.New("missing required field 'stateRoot' for Header")
	}
	h.Root = *DEWH.Root
	if DEWH.TxHash == nil {
		return errors.New("missing required field 'transactionsRoot' for Header")
	}
	h.TxHash = *DEWH.TxHash
	if DEWH.ReceiptHash == nil {
		return errors.New("missing required field 'receiptsRoot' for Header")
	}
	h.ReceiptHash = *DEWH.ReceiptHash
	if DEWH.Bloom == nil {
		return errors.New("missing required field 'logsBloom' for Header")
	}
	h.Bloom = *DEWH.Bloom
	if DEWH.Difficulty == nil {
		return errors.New("missing required field 'difficulty' for Header")
	}
	h.Difficulty = (*big.Int)(DEWH.Difficulty)
	if DEWH.Number == nil {
		return errors.New("missing required field 'number' for Header")
	}
	h.Number = (*big.Int)(DEWH.Number)
	if DEWH.GasLimit == nil {
		return errors.New("missing required field 'gasLimit' for Header")
	}
	h.GasLimit = uint64(*DEWH.GasLimit)
	if DEWH.GasUsed == nil {
		return errors.New("missing required field 'gasUsed' for Header")
	}
	h.GasUsed = uint64(*DEWH.GasUsed)
	if DEWH.Time == nil {
		return errors.New("missing required field 'timestamp' for Header")
	}
	h.Time = (*big.Int)(DEWH.Time)
	if DEWH.Extra == nil {
		return errors.New("missing required field 'extraData' for Header")
	}
	h.Extra = *DEWH.Extra
	if DEWH.MixDigest == nil {
		return errors.New("missing required field 'mixHash' for Header")
	}
	h.MixDigest = *DEWH.MixDigest
	if DEWH.Nonce == nil {
		return errors.New("missing required field 'nonce' for Header")
	}
	h.Nonce = *DEWH.Nonce
	return nil
}
