package sdk_rpc

type EthGetBlockByHashResult struct {
	Number           string               `json:"number,omitempty"`
	Hash             string               `json:"hash,omitempty"`
	ParentHash       string               `json:"parentHash,omitempty"`
	Nonce            string               `json:"nonce,omitempty"`
	Sha3Uncle        string               `json:"sha3Uncle,omitempty"`
	LogsBloom        string               `json:"logsBloom,omitempty"`
	TransactionsRoot string               `json:"transactionsRoot,omitempty"`
	StateRoot        string               `json:"stateRoot,omitempty"`
	ReceiptsRoot     string               `json:"receiptsRoot,omitempty"`
	Miner            string               `json:"miner,omitempty"`
	Difficulty       string               `json:"difficulty,omitempty"`
	TotalDifficulty  string               `json:"totalDifficulty,omitempty"`
	ExtraData        string               `json:"extraData,omitempty"`
	Size             string               `json:"size,omitempty"`
	GasLimit         string               `json:"gasLimit,omitempty"`
	GasUsed          string               `json:"gasUsed,omitempty"`
	Timestamp        string               `json:"timestamp,omitempty"`
	Transactions     []*TransactionResult `json:"transactions,omitempty"`
	Uncles           []string             `json:"uncles,omitempty"`
}

type TransactionResult struct {
	BlockHash        string `json:"blockHash,omitempty"`
	BlockNumber      string `json:"blockNumber,omitempty"`
	From             string `json:"from,omitempty"`
	Gas              string `json:"gas,omitempty"`
	GasPrice         string `json:"gasPrice,omitempty"`
	Hash             string `json:"hash,omitempty"`
	Input            string `json:"input,omitempty"`
	Nonce            string `json:"nonce,omitempty"`
	To               string `json:"to,omitempty"`
	TransactionIndex string `json:"transactionIndex,omitempty"`
	Value            string `json:"value,omitempty"`
	V                string `json:"v,omitempty"`
	R                string `json:"r,omitempty"`
	S                string `json:"s,omitempty"`
}

type TransactionReceiptResult struct {
	TransactionHash   string `json:"transactionHash,omitempty"`
	TransactionIndex  string `json:"transactionIndex,omitempty"`
	BlockHash         string `json:"blockHash,omitempty"`
	BlockNumber       string `json:"blockNumber,omitempty"`
	From              string `json:"from,omitempty"`
	To                string `json:"to,omitempty"`
	CumulativeGasUsed string `json:"cumulativeGasUsed,omitempty"`
	GasUsed           string `json:"gasUsed,omitempty"`
	ContractAddress   string `json:"contractAddress,omitempty"`
	Logs              []*Log `json:"logs,omitempty"`
	LogsBloom         string `json:"logsBloom,omitempty"`
	Root              string `json:"root,omitempty"`
	Status            string `json:"status,omitempty"`
}

type SyncingResult struct {
	StartingBlock string `json:"startingBlock,omitempty"`
	CurrentBlock  string `json:"currentBlock,omitempty"`
	HighestBlock  string `json:"highestBlock,omitempty"`
}

type Log struct {
	Removed          bool     `json:"removed"`
	LogIndex         string   `json:"logIndex"`
	TransactionIndex string   `json:"transactionIndex"`
	TransactionHash  string   `json:"transactionHash"`
	BlockHash        string   `json:"blockHash"`
	BlockNumber      string   `json:"blockNumber"`
	Address          string   `json:"address"`
	Data             string   `json:"data"`
	Topics           []string `json:"topics"`
}

type TransferResult struct {
	Transaction struct {
		Visible bool   `json:"visible"`
		TxID    string `json:"txID"`
		RawData struct {
			Contract []struct {
				Parameter struct {
					Value struct {
						Amount       int    `json:"amount"`
						OwnerAddress string `json:"owner_address"`
						ToAddress    string `json:"to_address"`
					} `json:"value"`
					TypeURL string `json:"type_url"`
				} `json:"parameter"`
				Type string `json:"type"`
			} `json:"contract"`
			RefBlockBytes string `json:"ref_block_bytes"`
			RefBlockHash  string `json:"ref_block_hash"`
			Expiration    int64  `json:"expiration"`
			Timestamp     int64  `json:"timestamp"`
		} `json:"raw_data"`
		RawDataHex string `json:"raw_data_hex"`
	} `json:"transaction"`
}

type CreateSmartContractResult struct {
	Transaction struct {
		Visible         bool   `json:"visible"`
		TxID            string `json:"txID"`
		ContractAddress string `json:"contract_address"`
		RawData         struct {
			Contract []struct {
				Parameter struct {
					Value struct {
						TokenID        int    `json:"token_id"`
						OwnerAddress   string `json:"owner_address"`
						CallTokenValue int    `json:"call_token_value"`
						NewContract    struct {
							Bytecode                   string `json:"bytecode"`
							ConsumeUserResourcePercent int    `json:"consume_user_resource_percent"`
							Name                       string `json:"name"`
							OriginAddress              string `json:"origin_address"`
							Abi                        struct {
								Entrys []struct {
									Outputs []struct {
										Type string `json:"type"`
									} `json:"outputs,omitempty"`
									Payable         bool   `json:"payable"`
									Name            string `json:"name,omitempty"`
									StateMutability string `json:"stateMutability"`
									Type            string `json:"type"`
									Inputs          []struct {
										Name string `json:"name"`
										Type string `json:"type"`
									} `json:"inputs,omitempty"`
								} `json:"entrys"`
							} `json:"abi"`
							OriginEnergyLimit int64 `json:"origin_energy_limit"`
							CallValue         int   `json:"call_value"`
						} `json:"new_contract"`
					} `json:"value"`
					TypeURL string `json:"type_url"`
				} `json:"parameter"`
				Type string `json:"type"`
			} `json:"contract"`
			RefBlockBytes string `json:"ref_block_bytes"`
			RefBlockHash  string `json:"ref_block_hash"`
			Expiration    int64  `json:"expiration"`
			FeeLimit      int    `json:"fee_limit"`
			Timestamp     int64  `json:"timestamp"`
		} `json:"raw_data"`
		RawDataHex string `json:"raw_data_hex"`
	} `json:"transaction"`
}

type TriggerSmartContractResult struct {
	Transaction struct {
		Visible bool   `json:"visible"`
		TxID    string `json:"txID"`
		RawData struct {
			Contract []struct {
				Parameter struct {
					Value struct {
						Amount       int    `json:"amount"`
						OwnerAddress string `json:"owner_address"`
						ToAddress    string `json:"to_address"`
					} `json:"value"`
					TypeURL string `json:"type_url"`
				} `json:"parameter"`
				Type string `json:"type"`
			} `json:"contract"`
			RefBlockBytes string `json:"ref_block_bytes"`
			RefBlockHash  string `json:"ref_block_hash"`
			Expiration    int64  `json:"expiration"`
			Timestamp     int64  `json:"timestamp"`
		} `json:"raw_data"`
		RawDataHex string `json:"raw_data_hex"`
	} `json:"transaction"`
}
