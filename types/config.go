package types

import (
	"log"

	emissions "github.com/allora-network/allora-chain/x/emissions/types"
	bank "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/ignite/cli/v28/ignite/pkg/cosmosaccount"
	"github.com/ignite/cli/v28/ignite/pkg/cosmosclient"
)

type WalletConfig struct {
	Address                  string
	AddressPrefix            string // prefix for the allora addresses
	AddressKeyName           string // load a address by key from the keystore
	AddressRestoreMnemonic   string
	AddressAccountPassphrase string
	AlloraHomeDir            string  // home directory for the allora keystore
	Gas                      string  // gas to use for the allora client
	GasAdjustment            float64 // gas adjustment to use for the allora client
	SubmitTx                 bool    // do we need to commit these to the chain, might be a reason not to
	LoopSeconds              int     // how often to run the main loops per worker and per reputer
	// Minimum stake to repute. will try to add stake from wallet if current stake is less than this.
	// Will not repute if current stake is less than this, after trying to add any necessary stake.
	// This is idempotent in that it will not add more stake than specified here.
	// Set to 0 to effectively disable this feature and use whatever stake has already been added.
	InitialStake             int64   // uallo to initially stake upon registration on a new topi
	MinStakeToRepute 		 string
	NodeRpc          		 string // rpc node for allora chain
	RequestRetries   		 int    // retry to get data from chain this many times per query or tx
}

type ChainConfig struct {
	Address              string
	Account              cosmosaccount.Account
	Client               *cosmosclient.Client
	EmissionsQueryClient emissions.QueryClient
	BankQueryClient      bank.QueryClient
}

type WorkerConfig struct {
	TopicId             emissions.TopicId
	InferenceEntrypoint AlloraEntrypoint
	ForecastEntrypoint  AlloraEntrypoint
}

type ReputerConfig struct {
	TopicId           emissions.TopicId
	ReputerEntrypoint AlloraEntrypoint
}

type UserConfig struct {
	Wallet  WalletConfig
	Worker  []WorkerConfig
	Reputer []ReputerConfig
}

type Config struct {
	Chain    ChainConfig
	Wallet  WalletConfig
	Worker  []WorkerConfig
	Reputer []ReputerConfig
}

func (c *UserConfig) MapUserConfigToFullConfig() Config {
	return Config{
		Chain:    ChainConfig{},
		Wallet:  c.Wallet,
		Worker:  c.Worker,
		Reputer: c.Reputer,
	}
}

// Check that each assigned entrypoint in `TheConfig` actually can be used
// for the intended purpose, else throw error
func (c *UserConfig) ValidateConfigEntrypoints() {
	for _, workerConfig := range c.Worker {
		if workerConfig.InferenceEntrypoint != nil && !workerConfig.InferenceEntrypoint.CanInfer() {
			log.Fatal("Invalid inference entrypoint: ", workerConfig.InferenceEntrypoint)
		}
		if workerConfig.ForecastEntrypoint != nil && !workerConfig.ForecastEntrypoint.CanForecast() {
			log.Fatal("Invalid forecast entrypoint: ", workerConfig.ForecastEntrypoint)
		}
	}

	for _, reputerConfig := range c.Reputer {
		if reputerConfig.ReputerEntrypoint != nil && !reputerConfig.ReputerEntrypoint.CanSourceTruth() {
			log.Fatal("Invalid loss entrypoint: ", reputerConfig.ReputerEntrypoint)
		}
	}
}
