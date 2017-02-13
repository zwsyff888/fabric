/*
Copyright IBM Corp. 2017 All Rights Reserved.

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

package application

import (
	"fmt"

	cb "github.com/hyperledger/fabric/protos/common"
	pb "github.com/hyperledger/fabric/protos/peer"

	"github.com/golang/protobuf/proto"
	"github.com/op/go-logging"
)

var orgSchema = &cb.ConfigGroupSchema{
	Groups: map[string]*cb.ConfigGroupSchema{},
	Values: map[string]*cb.ConfigValueSchema{
		"MSP": nil, // TODO, consolidate into a constant once common org code exists
	},
	Policies: map[string]*cb.ConfigPolicySchema{
	// TODO, set appropriately once hierarchical policies are implemented
	},
}

var Schema = &cb.ConfigGroupSchema{
	Groups: map[string]*cb.ConfigGroupSchema{
		"": orgSchema,
	},
	Values: map[string]*cb.ConfigValueSchema{
		AnchorPeersKey: nil,
	},
	Policies: map[string]*cb.ConfigPolicySchema{
	// TODO, set appropriately once hierarchical policies are implemented
	},
}

// Peer config keys
const (
	// AnchorPeersKey is the cb.ConfigItem type key name for the AnchorPeers message
	AnchorPeersKey = "AnchorPeers"
)

var logger = logging.MustGetLogger("peer/sharedconfig")

type sharedConfig struct {
	anchorPeers []*pb.AnchorPeer
}

// SharedConfigImpl is an implementation of Manager and configtx.ConfigHandler
// In general, it should only be referenced as an Impl for the configtx.Manager
type SharedConfigImpl struct {
	pendingConfig *sharedConfig
	config        *sharedConfig
}

// NewSharedConfigImpl creates a new SharedConfigImpl with the given CryptoHelper
func NewSharedConfigImpl() *SharedConfigImpl {
	return &SharedConfigImpl{
		config: &sharedConfig{},
	}
}

// AnchorPeers returns the list of valid orderer addresses to connect to to invoke Broadcast/Deliver
func (di *SharedConfigImpl) AnchorPeers() []*pb.AnchorPeer {
	return di.config.anchorPeers
}

// BeginConfig is used to start a new config proposal
func (di *SharedConfigImpl) BeginConfig() {
	logger.Debugf("Beginning a possible new peer shared config")
	if di.pendingConfig != nil {
		logger.Panicf("Programming error, cannot call begin in the middle of a proposal")
	}
	di.pendingConfig = &sharedConfig{}
}

// RollbackConfig is used to abandon a new config proposal
func (di *SharedConfigImpl) RollbackConfig() {
	logger.Debugf("Rolling back proposed peer shared config")
	di.pendingConfig = nil
}

// CommitConfig is used to commit a new config proposal
func (di *SharedConfigImpl) CommitConfig() {
	logger.Debugf("Committing new peer shared config")
	if di.pendingConfig == nil {
		logger.Panicf("Programming error, cannot call commit without an existing proposal")
	}
	di.config = di.pendingConfig
	di.pendingConfig = nil
}

// ProposeConfig is used to add new config to the config proposal
func (di *SharedConfigImpl) ProposeConfig(configItem *cb.ConfigItem) error {
	if configItem.Type != cb.ConfigItem_PEER {
		return fmt.Errorf("Expected type of ConfigItem_Peer, got %v", configItem.Type)
	}

	switch configItem.Key {
	case AnchorPeersKey:
		anchorPeers := &pb.AnchorPeers{}
		if err := proto.Unmarshal(configItem.Value, anchorPeers); err != nil {
			return fmt.Errorf("Unmarshaling error for %s: %s", configItem.Key, err)
		}
		if logger.IsEnabledFor(logging.DEBUG) {
			logger.Debugf("Setting %s to %v", configItem.Key, anchorPeers.AnchorPeers)
		}
		di.pendingConfig.anchorPeers = anchorPeers.AnchorPeers
	default:
		logger.Warningf("Uknown Peer config item with key %s", configItem.Key)
	}
	return nil
}
