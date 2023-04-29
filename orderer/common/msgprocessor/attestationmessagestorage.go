/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package msgprocessor

import (
	"encoding/json"
	"strconv"

	cb "github.com/hyperledger/fabric-protos-go/common"
	"github.com/syndtr/goleveldb/leveldb"
)

type AttestationMessageStorage interface {
	Add(blockNumber uint64, newMessage *cb.Envelope) error
	Get(blockNumber uint64) ([]cb.Envelope, error)
}

func NewAttestationMessageStorage(channelID string) AttestationMessageStorage {
	storage, err := leveldb.OpenFile("/mptDB/orderer/attestationMessagesStorages/"+channelID, nil)
	if err != nil {
		logger.Errorf("Error while init level db in /mptDB/orderer/attestationMessagesStorage/%s: %s", channelID, err)
	}
	return &attestationMessageStorage{storage: storage}
}

type attestationMessageStorage struct {
	storage *leveldb.DB
}

func (s *attestationMessageStorage) Add(blockNumber uint64, newMessage *cb.Envelope) error {
	messages, err := s.Get(blockNumber)
	if err != nil {
		return err
	}

	messages = append(messages, *newMessage)
	marshaledMessages, err := json.Marshal(messages)
	if err != nil {
		logger.Warningf("error while marshaling messages: %s", err)
		return err
	}
	err = s.storage.Put(keyByBlockNumber(blockNumber), marshaledMessages, nil)
	if err != nil {
		logger.Warningf("error while saving message: %s", err)
		return err
	}

	return nil
}

func (s *attestationMessageStorage) Get(blockNumber uint64) ([]cb.Envelope, error) {
	key := keyByBlockNumber(blockNumber)
	envelopesBytes, err := s.storage.Get(key, nil)
	if err == leveldb.ErrNotFound {
		return nil, nil
	} else if err != nil {
		logger.Warningf("error while getting exist messages: %s", err)
		return nil, err
	}

	var envelopes []cb.Envelope
	err = json.Unmarshal(envelopesBytes, &envelopes)
	if err != nil {
		logger.Warningf("error while unmarshaling exist messages: %s", err)
		return nil, err
	}
	return envelopes, nil
}

func keyByBlockNumber(blockNumber uint64) []byte {
	return []byte(strconv.FormatUint(blockNumber, 10))
}

type attestationMessageInMemoryStorage struct {
	store map[string][]cb.Envelope
}

func (s attestationMessageInMemoryStorage) Add(blockNumber uint64, newMessage *cb.Envelope) error {
	messages, err := s.Get(blockNumber)
	if err != nil {
		return err
	}

	messages = append(messages, *newMessage)
	s.store[string(keyByBlockNumber(blockNumber))] = messages

	return nil
}

func (s attestationMessageInMemoryStorage) Get(blockNumber uint64) ([]cb.Envelope, error) {
	return s.store[string(keyByBlockNumber(blockNumber))], nil
}

func NewAttestationMessageInMemoryStorage() AttestationMessageStorage {
	return &attestationMessageInMemoryStorage{store: make(map[string][]cb.Envelope)}
}
