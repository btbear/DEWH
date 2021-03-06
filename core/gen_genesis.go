// Code generated by github.com/fjl/gencoDEWH. DO NOT EDIT.

package core

import (
	"encoding/json"
	"errors"
	"math/big"

	"github.com/DEWH/go-DEWH/common"
	"github.com/DEWH/go-DEWH/common/hexutil"
	"github.com/DEWH/go-DEWH/common/math"
	"github.com/DEWH/go-DEWH/params"
)

var _ = (*genesisSpecMarshaling)(nil)

func (g Genesis) MarshalJSON() ([]byte, error) {
	type Genesis struct {
		Config     *params.ChainConfig                         `json:"config"`
		Nonce      math.HexOrDEWHimal64                         `json:"nonce"`
		Timestamp  math.HexOrDEWHimal64                         `json:"timestamp"`
		ExtraData  hexutil.Bytes                               `json:"extraData"`
		GasLimit   math.HexOrDEWHimal64                         `json:"gasLimit"   gencoDEWH:"required"`
		Difficulty *math.HexOrDEWHimal256                       `json:"difficulty" gencoDEWH:"required"`
		Mixhash    common.Hash                                 `json:"mixHash"`
		Coinbase   common.Address                              `json:"coinbase"`
		Alloc      map[common.UnprefixedAddress]GenesisAccount `json:"alloc"      gencoDEWH:"required"`
		Number     math.HexOrDEWHimal64                         `json:"number"`
		GasUsed    math.HexOrDEWHimal64                         `json:"gasUsed"`
		ParentHash common.Hash                                 `json:"parentHash"`
	}
	var enc Genesis
	enc.Config = g.Config
	enc.Nonce = math.HexOrDEWHimal64(g.Nonce)
	enc.Timestamp = math.HexOrDEWHimal64(g.Timestamp)
	enc.ExtraData = g.ExtraData
	enc.GasLimit = math.HexOrDEWHimal64(g.GasLimit)
	enc.Difficulty = (*math.HexOrDEWHimal256)(g.Difficulty)
	enc.Mixhash = g.Mixhash
	enc.Coinbase = g.Coinbase
	if g.Alloc != nil {
		enc.Alloc = make(map[common.UnprefixedAddress]GenesisAccount, len(g.Alloc))
		for k, v := range g.Alloc {
			enc.Alloc[common.UnprefixedAddress(k)] = v
		}
	}
	enc.Number = math.HexOrDEWHimal64(g.Number)
	enc.GasUsed = math.HexOrDEWHimal64(g.GasUsed)
	enc.ParentHash = g.ParentHash
	return json.Marshal(&enc)
}

func (g *Genesis) UnmarshalJSON(input []byte) error {
	type Genesis struct {
		Config     *params.ChainConfig                         `json:"config"`
		Nonce      *math.HexOrDEWHimal64                        `json:"nonce"`
		Timestamp  *math.HexOrDEWHimal64                        `json:"timestamp"`
		ExtraData  *hexutil.Bytes                              `json:"extraData"`
		GasLimit   *math.HexOrDEWHimal64                        `json:"gasLimit"   gencoDEWH:"required"`
		Difficulty *math.HexOrDEWHimal256                       `json:"difficulty" gencoDEWH:"required"`
		Mixhash    *common.Hash                                `json:"mixHash"`
		Coinbase   *common.Address                             `json:"coinbase"`
		Alloc      map[common.UnprefixedAddress]GenesisAccount `json:"alloc"      gencoDEWH:"required"`
		Number     *math.HexOrDEWHimal64                        `json:"number"`
		GasUsed    *math.HexOrDEWHimal64                        `json:"gasUsed"`
		ParentHash *common.Hash                                `json:"parentHash"`
	}
	var DEWH Genesis
	if err := json.Unmarshal(input, &DEWH); err != nil {
		return err
	}
	if DEWH.Config != nil {
		g.Config = DEWH.Config
	}
	if DEWH.Nonce != nil {
		g.Nonce = uint64(*DEWH.Nonce)
	}
	if DEWH.Timestamp != nil {
		g.Timestamp = uint64(*DEWH.Timestamp)
	}
	if DEWH.ExtraData != nil {
		g.ExtraData = *DEWH.ExtraData
	}
	if DEWH.GasLimit == nil {
		return errors.New("missing required field 'gasLimit' for Genesis")
	}
	g.GasLimit = uint64(*DEWH.GasLimit)
	if DEWH.Difficulty == nil {
		return errors.New("missing required field 'difficulty' for Genesis")
	}
	g.Difficulty = (*big.Int)(DEWH.Difficulty)
	if DEWH.Mixhash != nil {
		g.Mixhash = *DEWH.Mixhash
	}
	if DEWH.Coinbase != nil {
		g.Coinbase = *DEWH.Coinbase
	}
	if DEWH.Alloc == nil {
		return errors.New("missing required field 'alloc' for Genesis")
	}
	g.Alloc = make(GenesisAlloc, len(DEWH.Alloc))
	for k, v := range DEWH.Alloc {
		g.Alloc[common.Address(k)] = v
	}
	if DEWH.Number != nil {
		g.Number = uint64(*DEWH.Number)
	}
	if DEWH.GasUsed != nil {
		g.GasUsed = uint64(*DEWH.GasUsed)
	}
	if DEWH.ParentHash != nil {
		g.ParentHash = *DEWH.ParentHash
	}
	return nil
}
