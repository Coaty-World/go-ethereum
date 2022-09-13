// Copyright 2016 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

// Contains all the wrappers from the core/types package.

package geth

import (
	"errors"

	"github.com/Coaty-World/go-ethereum/core/types"
)

// Log represents a contract log event. These events are generated by the LOG
// opcode and stored/indexed by the node.
type Log struct {
	log *types.Log
}

func (l *Log) GetAddress() *Address  { return &Address{l.log.Address} }
func (l *Log) GetTopics() *Hashes    { return &Hashes{l.log.Topics} }
func (l *Log) GetData() []byte       { return l.log.Data }
func (l *Log) GetBlockNumber() int64 { return int64(l.log.BlockNumber) }
func (l *Log) GetTxHash() *Hash      { return &Hash{l.log.TxHash} }
func (l *Log) GetTxIndex() int       { return int(l.log.TxIndex) }
func (l *Log) GetBlockHash() *Hash   { return &Hash{l.log.BlockHash} }
func (l *Log) GetIndex() int         { return int(l.log.Index) }

// Logs represents a slice of VM logs.
type Logs struct{ logs []*types.Log }

// Size returns the number of logs in the slice.
func (l *Logs) Size() int {
	return len(l.logs)
}

// Get returns the log at the given index from the slice.
func (l *Logs) Get(index int) (log *Log, _ error) {
	if index < 0 || index >= len(l.logs) {
		return nil, errors.New("index out of bounds")
	}
	return &Log{l.logs[index]}, nil
}
