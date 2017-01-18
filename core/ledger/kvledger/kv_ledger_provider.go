/*
Copyright IBM Corp. 2016 All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

		 http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package kvledger

import (
	"errors"

	"github.com/hyperledger/fabric/core/ledger"
	"github.com/hyperledger/fabric/core/ledger/kvledger/txmgmt/statedb"
	"github.com/hyperledger/fabric/core/ledger/kvledger/txmgmt/statedb/statecouchdb"
	"github.com/hyperledger/fabric/core/ledger/kvledger/txmgmt/statedb/stateleveldb"
	"github.com/hyperledger/fabric/core/ledger/ledgerconfig"
	"github.com/hyperledger/fabric/core/ledger/util/db"
)

var (
	// ErrLedgerIDExists is thrown by a CreateLedger call if a ledger with the given id already exists
	ErrLedgerIDExists = errors.New("LedgerID already exists")
	// ErrNonExistingLedgerID is thrown by a OpenLedger call if a ledger with the given id does not exist
	ErrNonExistingLedgerID = errors.New("LedgerID does not exist")
	// ErrLedgerNotOpened is thrown by a CloseLedger call if a ledger with the given id has not been opened
	ErrLedgerNotOpened = errors.New("Ledger is not opened yet")
)

// Provider implements interface ledger.PeerLedgerProvider
type Provider struct {
	idStore     *idStore
	vdbProvider statedb.VersionedDBProvider
}

// NewProvider instantiates a new Provider.
// This is not thread-safe and assumed to be synchronized be the caller
func NewProvider() (ledger.PeerLedgerProvider, error) {
	logger.Info("Initializing ledger provider")
	var vdbProvider statedb.VersionedDBProvider
	if !ledgerconfig.IsCouchDBEnabled() {
		logger.Debugf("Constructing leveldb VersionedDBProvider")
		vdbProvider = stateleveldb.NewVersionedDBProvider()
	} else {
		logger.Debugf("Constructing CouchDB VersionedDBProvider")
		var err error
		vdbProvider, err = statecouchdb.NewVersionedDBProvider()
		if err != nil {
			return nil, err
		}
	}
	ledgerMgmtPath := ledgerconfig.GetLedgerProviderPath()
	idStore := openIDStore(ledgerMgmtPath)
	logger.Info("ledger provider Initialized")
	return &Provider{idStore, vdbProvider}, nil
}

// Create implements the corresponding method from interface ledger.PeerLedgerProvider
func (provider *Provider) Create(ledgerID string) (ledger.PeerLedger, error) {
	exists, err := provider.idStore.ledgerIDExists(ledgerID)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, ErrLedgerIDExists
	}
	provider.idStore.createLedgerID(ledgerID)
	l, err := NewKVLedger(provider.vdbProvider, ledgerID)
	if err != nil {
		return nil, err
	}
	return l, nil
}

// Open implements the corresponding method from interface ledger.PeerLedgerProvider
func (provider *Provider) Open(ledgerID string) (ledger.PeerLedger, error) {
	exists, err := provider.idStore.ledgerIDExists(ledgerID)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, ErrNonExistingLedgerID
	}
	l, err := NewKVLedger(provider.vdbProvider, ledgerID)
	if err != nil {
		return nil, err
	}
	return l, nil
}

// Exists implements the corresponding method from interface ledger.PeerLedgerProvider
func (provider *Provider) Exists(ledgerID string) (bool, error) {
	return provider.idStore.ledgerIDExists(ledgerID)
}

// List implements the corresponding method from interface ledger.PeerLedgerProvider
func (provider *Provider) List() ([]string, error) {
	return provider.idStore.getAllLedgerIds()
}

// Close implements the corresponding method from interface ledger.PeerLedgerProvider
func (provider *Provider) Close() {
	provider.vdbProvider.Close()
	provider.idStore.close()
}

type idStore struct {
	db *db.DB
}

func openIDStore(path string) *idStore {
	db := db.CreateDB(&db.Conf{DBPath: path})
	db.Open()
	return &idStore{db}
}

func (s *idStore) createLedgerID(ledgerID string) error {
	key := []byte(ledgerID)
	val := []byte{}
	err := error(nil)
	if val, err = s.db.Get(key); err != nil {
		return err
	}
	if val != nil {
		return ErrLedgerIDExists
	}
	return s.db.Put(key, val, true)
}

func (s *idStore) ledgerIDExists(ledgerID string) (bool, error) {
	key := []byte(ledgerID)
	val := []byte{}
	err := error(nil)
	if val, err = s.db.Get(key); err != nil {
		return false, err
	}
	return val != nil, nil
}

func (s *idStore) getAllLedgerIds() ([]string, error) {
	var ids []string
	itr := s.db.GetIterator(nil, nil)
	itr.First()
	for itr.Valid() {
		key := string(itr.Key())
		ids = append(ids, key)
		itr.Next()
	}
	return ids, nil
}

func (s *idStore) close() {
	s.db.Close()
}
