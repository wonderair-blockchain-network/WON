// Copyright (c) 2018 The wonderair ecosystem Authors
// Distributed under the MIT software license, see the accompanying
// file COPYING or or or http://www.opensource.org/licenses/mit-license.php
package core

import (
	"github.com/wonderair ecosystem/go-wonderair ecosystem/common"
	"github.com/wonderair ecosystem/go-wonderair ecosystem/core/types"
	"github.com/wonderair ecosystem/go-wonderair ecosystem/wondb"
	"io/ioutil"
	"math/big"
	"testing"
)

type testChainReader struct {
	testHeader map[uint64]*types.Header
	curNumber  uint64
}

func newTestChainReader() *testChainReader {
	tcr := &testChainReader{
		testHeader: make(map[uint64]*types.Header),
		curNumber:  0,
	}

	header0 := &types.Header{
		ParentHash: common.Hash{},
		Number:     big.NewInt(0),
	}
	header0.Elect = make([]common.Elect, 0)
	header0.Elect = append(header0.Elect, common.Elect{
		Account: common.HexToAddress("100A"),
		Stock:   2,
		Type:    common.ElectRoleValidator,
	})
	header0.Elect = append(header0.Elect, common.Elect{
		Account: common.HexToAddress("100B"),
		Stock:   3,
		Type:    common.ElectRoleValidator,
	})
	header0.Elect = append(header0.Elect, common.Elect{
		Account: common.HexToAddress("100C"),
		Stock:   4,
		Type:    common.ElectRoleValidator,
	})
	header0.Elect = append(header0.Elect, common.Elect{
		Account: common.HexToAddress("200A"),
		Stock:   0,
		Type:    common.ElectRoleMiner,
	})

	header0.NetTopology = common.NetTopology{
		Type:            common.NetTopoTypeAll,
		NetTopologyData: make([]common.NetTopologyData, 0),
	}
	header0.NetTopology.NetTopologyData = append(header0.NetTopology.NetTopologyData, common.NetTopologyData{
		Account:  common.HexToAddress("100A"),
		Position: common.GeneratePosition(uint16(0), common.ElectRoleValidator),
	})
	header0.NetTopology.NetTopologyData = append(header0.NetTopology.NetTopologyData, common.NetTopologyData{
		Account:  common.HexToAddress("100B"),
		Position: common.GeneratePosition(uint16(1), common.ElectRoleValidator),
	})
	header0.NetTopology.NetTopologyData = append(header0.NetTopology.NetTopologyData, common.NetTopologyData{
		Account:  common.HexToAddress("100C"),
		Position: common.GeneratePosition(uint16(2), common.ElectRoleValidator),
	})
	header0.NetTopology.NetTopologyData = append(header0.NetTopology.NetTopologyData, common.NetTopologyData{
		Account:  common.HexToAddress("200A"),
		Position: common.GeneratePosition(uint16(0), common.ElectRoleMiner),
	})

	tcr.testHeader[0] = header0

	return tcr
}

func (tc *testChainReader) GetHeaderByNumber(number uint64) *types.Header {
	header, ok := tc.testHeader[number]
	if ok {
		return header
	} else {
		return nil
	}
}

func (tc *testChainReader) CurrentHeader() *types.Header {
	header, _ := tc.testHeader[tc.curNumber]
	return header
}

func TestTopologyStore_GetTopologyGraphByNumber(t *testing.T) {
	workspace, err := ioutil.TempDir("", "topology_store_test-")
	if err != nil {
		t.Fatalf("创建workspace失败, %v", err)
	}

	chainReader := newTestChainReader()
	db, err := wondb.NewLDBDatabase(workspace, 0, 0)
	if err != nil {
		t.Fatalf("创建db错误, %v", err)
	}

	store := NewTopologyStore(chainReader, db)
	store.WriteTopologyGraph(chainReader.GetHeaderByNumber(0))
}
