// Copyright 2016 The go-estchain Authors
// This file is part of the go-estchain library.
//
// The go-estchain library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-estchain library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-estchain library. If not, see <http://www.gnu.org/licenses/>.

package ethclient

import "github.com/estchain/go-estchain"

// Verify that Client implements the estchain interfaces.
var (
	_ = estchain.ChainReader(&Client{})
	_ = estchain.TransactionReader(&Client{})
	_ = estchain.ChainStateReader(&Client{})
	_ = estchain.ChainSyncReader(&Client{})
	_ = estchain.ContractCaller(&Client{})
	_ = estchain.GasEstimator(&Client{})
	_ = estchain.GasPricer(&Client{})
	_ = estchain.LogFilterer(&Client{})
	_ = estchain.PendingStateReader(&Client{})
	// _ = estchain.PendingStateEventer(&Client{})
	_ = estchain.PendingContractCaller(&Client{})
)