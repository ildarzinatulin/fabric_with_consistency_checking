package msgprocessor

import (
	"encoding/json"
	cb "github.com/hyperledger/fabric-protos-go/common"
	"github.com/syndtr/goleveldb/leveldb"
	"strconv"
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
	err = s.storage.Put(s.keyByBlockNumber(blockNumber), marshaledMessages, nil)
	if err != nil {
		logger.Warningf("error while saving message: %s", err)
		return err
	}

	return nil
}

func (s *attestationMessageStorage) Get(blockNumber uint64) ([]cb.Envelope, error) {
	key := s.keyByBlockNumber(blockNumber)
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

func (s *attestationMessageStorage) keyByBlockNumber(blockNumber uint64) []byte {
	return []byte(strconv.FormatUint(blockNumber, 10))
}
