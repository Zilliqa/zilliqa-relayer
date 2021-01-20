package service

import (
	"github.com/Zilliqa/gozilliqa-sdk/account"
	"github.com/Zilliqa/gozilliqa-sdk/provider"
	poly "github.com/polynetwork/poly-go-sdk"
	"github.com/polynetwork/zilliqa-relayer/config"
	"github.com/polynetwork/zilliqa-relayer/db"
)

type ZilliqaSyncManager struct {
	polySigner               *poly.Account
	polySdk                  *poly.PolySdk
	relaySyncHeight          uint32
	zilAccount               *account.Account
	currentHeight            uint64
	zilSdk                   *provider.Provider
	crossChainManagerAddress string
	cfg                      *config.Config
	db                       *db.BoltDB
	exitChan                 chan int
}

func NewZilliqaSyncManager(cfg *config.Config, zilSdk *provider.Provider, boltDB *db.BoltDB) *ZilliqaSyncManager {
	return &ZilliqaSyncManager{
		db:                       boltDB,
		cfg:                      cfg,
		zilSdk:                   zilSdk,
		currentHeight:            uint64(cfg.ZilConfig.ZilStartHeight),
		crossChainManagerAddress: cfg.ZilConfig.CrossChainManagerContract,
	}
}

func (s *ZilliqaSyncManager) Run() {
	go s.MonitorChain()
	go s.MonitorDeposit()
}
