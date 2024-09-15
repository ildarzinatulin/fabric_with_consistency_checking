/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package privdata

import (
	"github.com/hyperledger/fabric/common/channelconfig"
	"github.com/hyperledger/fabric/internal/peer/common"
)

// appCapabilities local interface used to generate mock for foreign interface.
//
//go:generate mockery -dir ./ -name AppCapabilities -case underscore -output mocks/
type AppCapabilities interface {
	channelconfig.ApplicationCapabilities
}

//go:generate mockery --dir ./ --name AttestationCheckingParameters --case underscore --output mocks/
type AttestationCheckingParameters interface {
	channelconfig.AttestationCheckingParameters
}

//go:generate mockery --dir ./ --name BroadcastClient --case underscore --output mocks/
type BroadcastClient interface {
	common.BroadcastClient
}
