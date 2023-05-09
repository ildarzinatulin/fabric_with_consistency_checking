/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package privdata

import (
	"context"

	cb "github.com/hyperledger/fabric-protos-go/common"
	"github.com/hyperledger/fabric-protos-go/orderer"
	"github.com/hyperledger/fabric/core/deliverservice"
	"github.com/hyperledger/fabric/gossip/util"
	"github.com/hyperledger/fabric/internal/peer/common"
	"github.com/hyperledger/fabric/internal/pkg/comm"
	"github.com/hyperledger/fabric/internal/pkg/peer/orderers"
	"github.com/hyperledger/fabric/protoutil"
)

//go:generate mockery --dir ./ --name AttestationMessageSender --case underscore --output mocks/
type AttestationMessageSender interface {
	Send(blockNUmber uint64, mptHead []byte) error
}

func NewAttestationMessageSender(chainID string, signer common.Signer, ordererSource *orderers.ConnectionSource) AttestationMessageSender {
	return &attestationMessageSender{
		chainID:       chainID,
		logger:        logger.With("channel", chainID),
		signer:        signer,
		ordererSource: ordererSource,
	}
}

type attestationMessageSender struct {
	logger          util.Logger
	signer          common.Signer
	chainID         string
	broadcastClient orderer.AtomicBroadcast_BroadcastClient
	ordererSource   *orderers.ConnectionSource
}

func (s *attestationMessageSender) Send(blockNUmber uint64, mptHead []byte) error {
	s.logger.Error("Sending attestation message")

	if !s.tryInitBroadcastClient() {
		s.logger.Error("Broadcast client not yet created for attestation messages")
		return nil
	}
	s.logger.Debug("Broadcast client was created for attestation messages")

	env, err := protoutil.CreateSignedEnvelope(cb.HeaderType_ATTESTATION, s.chainID, s.signer, &cb.AttestationEnvelope{TrieHead: mptHead, BlockNumber: blockNUmber}, 0, 0)
	if err != nil {
		s.logger.Errorf("Error while creating envelope for attestation: %s", err)
		return err
	}

	s.logger.Debug("Sending attestation message is finished")
	return s.broadcastClient.Send(env)
}

func (s *attestationMessageSender) tryInitBroadcastClient() bool {
	if s.broadcastClient != nil {
		return true
	}

	endpoint, err := s.ordererSource.RandomEndpoint()
	if err != nil {
		logger.Errorf("Error while getting orderer endpoint for channel %s: %s", s.chainID, err)
		return false
	}

	config := deliverservice.GlobalConfig()
	cc := comm.ClientConfig{
		DialTimeout: config.ConnectionTimeout,
		KaOpts:      config.KeepaliveOptions,
		SecOpts:     config.SecOpts,
	}
	cc.SecOpts.ServerRootCAs = endpoint.RootCerts
	conn, err := cc.Dial(endpoint.Address)
	if err != nil {
		logger.Errorf("Error while dial for channel %s: %s", s.chainID, err)
		return false
	}

	broadcastClient, err := orderer.NewAtomicBroadcastClient(conn).Broadcast(context.TODO())
	if err != nil {
		logger.Errorf("Error while creating AtomicBroadcastClient for channel %s: %s", s.chainID, err)
		return false
	}

	s.broadcastClient = broadcastClient

	return true
}
