package main

import (
	"allora_offchain_node/lib"
	"os"
)

//// If in the event inferences, forecasts, truth, or loss functions are
//// not derived from the Python Flask server, new adapters can be made
//// and added to the UserConfig struct below for the same or different topics.

var UserConfig = lib.UserConfig{
	Wallet: lib.WalletConfig{
		AddressKeyName:           os.Getenv("ALLORA_ACCOUNT_NAME"),       // load a address by key from the keystore testing = allo1wmfp4xuurjsvh3qzjhkdxqgepmshpy7ny88pc7
		AddressRestoreMnemonic:   os.Getenv("ALLORA_ACCOUNT_MNEMONIC"),   // mnemonic for the allora account
		AddressAccountPassphrase: os.Getenv("ALLORA_ACCOUNT_PASSPHRASE"), // passphrase for the allora account
		AlloraHomeDir:            "",                                     // home directory for the allora keystore, if "", it will automatically create in "$HOME/.allorad"
		Gas:                      "1000000",                              // gas to use for the allora client in uallo
		GasAdjustment:            1.0,                                    // gas adjustment to use for the allora client
		SubmitTx:                 false,                                  // set to false to run in dry-run processes without committing to the chain. useful for dev/testing
		NodeRpc:                  os.Getenv("ALLORA_NODE_RPC"),
		MaxRetries:               3,
		MinDelay:                 1,
		MaxDelay:                 6,
	},
	Worker: []lib.WorkerConfig{
		{
			TopicId:                 1,
			InferenceEntrypointName: "api-worker-reputer",
			LoopSeconds:             5,
			Parameters: map[string]string{
				//// These communicate with local Python Flask server
				"InferenceEndpoint": os.Getenv("INFERENCE_URL"),
				"Token":             "ETH",
				"ForecastEndpoint":  os.Getenv("FORECAST_URL"),
			},
		},
		{
			TopicId:                 1,
			InferenceEntrypointName: "api-worker-reputer",
			LoopSeconds:             5,
			Parameters: map[string]string{
				//// These communicate with local Python Flask server
				"inferenceEndpoint": os.Getenv("INFERENCE_URL"),
				"token":             "ETH",
				"forecastEndpoint":  os.Getenv("FORECAST_URL"),
			},
		},
	},
	Reputer: []lib.ReputerConfig{
		{
			TopicId:               1,
			ReputerEntrypointName: "api-worker-reputer",
			LoopSeconds:           30,
			MinStake:              100000,
			Parameters: map[string]string{
				"SourceOfTruthEndpoint": os.Getenv("TRUTH_URL"),
				"Token":                 "ethereum",
				//// Could put this in Python Flask server as well
				// "cgSimpleEndpoint": "https://api.coingecko.com/api/v3/simple/price?vs_currencies=usd&ids=",
				// "apiKey":           os.Getenv("CG_API_KEY"),
			},
		},
		{
			TopicId:               1,
			ReputerEntrypointName: "api-worker-reputer",
			LoopSeconds:           30,
			MinStake:              100000,
			Parameters: map[string]string{
				"truthEndpoint": os.Getenv("TRUTH_URL"),
				"token":         "ethereum",
				//// Could put this in Python Flask server as well
				// "cgSimpleEndpoint": "https://api.coingecko.com/api/v3/simple/price?vs_currencies=usd&ids=",
				// "apiKey":           os.Getenv("CG_API_KEY"),
			},
		},
	},
}

//// The config above implies that I have a local server running on port 8000
//// that can handle the following endpoints: /inference, /forecast, /truth
//// The server should be able to handle GET requests to these endpoints
//// and return the appropriate data for the worker and reputer processes.
//// The server's endpoints are all assigned for topic 1

//// It is up to the user to add more topics and endpoints as necessary,
//// and to ensure that models and sources of truth relayed or generated by
//// adapters are compatible with their assigned topic(s).
