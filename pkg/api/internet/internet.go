package internet

type Result struct {
	Result  string `json:"result"`
	Code    string `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

type ValidateAddressInternet struct {
	Result  bool   `json:"result,omitempty"`
	Message string `json:"message,omitempty"`
}

type CreateAccountInternet struct {
	Error string `json:"error,omitempty"`
}

type DelegatedResourceInternet struct {
	From                      string   `json:"from,omitempty"`
	To                        string   `json:"to,omitempty"`
	FrozenBalanceForBandwidth int64    `json:"frozen_balance_for_bandwidth,omitempty"`
	FrozenBalanceForEnergy    int64    `json:"frozen_balance_for_energy,omitempty"`
	ExpireTimeForBandwidth    int64    `json:"expire_time_for_bandwidth,omitempty"`
	ExpireTimeForEnergy       int64    `json:"expire_time_for_energy,omitempty"`
	Account                   string   `json:"account,omitempty"`
	FromAccounts              []string `json:"fromAccounts,omitempty"`
	ToAccounts                []string `json:"toAccounts,omitempty"`
}

type AccountResourceInternet struct {
	FreeNetUsed          int64 `json:"freeNetUsed,omitempty"`
	FreeNetLimit         int64 `json:"freeNetLimit,omitempty"`
	NetUsed              int64 `json:"NetUsed,omitempty"`
	NetLimit             int64 `json:"NetLimit,omitempty"`
	TotalNetLimit        int64 `json:"TotalNetLimit,omitempty"`
	TotalNetWeight       int64 `json:"TotalNetWeight,omitempty"`
	TotalTronPowerWeight int64 `json:"totalTronPowerWeight,omitempty"`
	TronPowerLimit       int64 `json:"tronPowerLimit,omitempty"`
	TronPowerUsed        int64 `json:"tronPowerUsed,omitempty"`
	EnergyUsed           int64 `json:"EnergyUsed,omitempty"`
	EnergyLimit          int64 `json:"EnergyLimit,omitempty"`
	TotalEnergyLimit     int64 `json:"TotalEnergyLimit,omitempty"`
	TotalEnergyWeight    int64 `json:"TotalEnergyWeight,omitempty"`
	AssetNetUsed         []struct {
		Key   string `json:"key,omitempty"`
		Value int64  `json:"value"`
	} `json:"assetNetUsed,omitempty"`
	AssetNetLimit []struct {
		Key   string `json:"key,omitempty"`
		Value int64  `json:"value"`
	} `json:"assetNetLimit,omitempty"`
}

type AccountBalanceInternet struct {
	Balance         int64 `json:"balance,omitempty"`
	BlockIdentifier struct {
		Hash   string `json:"hash,omitempty"`
		Number int64  `json:"number,omitempty"`
	} `json:"block_identifier,omitempty"`
}

type BlockInternet struct {
	BlockID     string `json:"blockID,omitempty"`
	BlockHeader struct {
		RawData struct {
			Number           int64  `json:"number,omitempty"`
			TxTrieRoot       string `json:"txTrieRoot,omitempty"`
			WitnessAddress   string `json:"witness_address,omitempty"`
			ParentHash       string `json:"parentHash,omitempty"`
			Version          int32  `json:"version,omitempty"`
			Timestamp        int64  `json:"timestamp,omitempty"`
			WitnessId        int64  `json:"witness_id,omitempty"`
			AccountStateRoot string `json:"accountStateRoot,omitempty"`
		} `json:"raw_data,omitempty"`
		WitnessSignature string `json:"witness_signature,omitempty"`
	} `json:"block_header,omitempty"`
	Transactions []Transaction `json:"transactions,omitempty"`
}

type BlockListInternet struct {
	Block []BlockInternet `json:"block,omitempty"`
}

type Transaction struct {
	Result struct {
		Result bool `json:"result,omitempty"`
	} `json:"result,omitempty"`
	Visible bool `json:"visible,omitempty"`
	Ret     []struct {
		ContractRet string `json:"contractRet,omitempty"`
	} `json:"ret,omitempty"`
	Signature []string `json:"signature,omitempty"`
	TxID      string   `json:"txID,omitempty"`
	RawData   struct {
		Contract []struct {
			Parameter struct {
				Value struct {
					Data            string `json:"data,omitempty"`
					Balance         int64  `json:"balance,omitempty"`
					Resource        string `json:"resource,omitempty"`
					ReceiverAddress string `json:"receiver_address,omitempty"`
					Amount          int64  `json:"amount,omitempty"`
					AccountName     string `json:"account_name,omitempty"`
					FrozenBalance   int64  `json:"frozen_balance,omitempty"`
					UnfreezeBalance int64  `json:"unfreeze_balance,omitempty"`
					OwnerAddress    string `json:"owner_address,omitempty"`
					ToAddress       string `json:"to_address,omitempty"`
					ContractAddress string `json:"contract_address,omitempty"`
					Lock            bool   `json:"lock,omitempty"`
					LockPeriod      int64  `json:"lock_period,omitempty"`
					AssetName       string `json:"asset_name,omitempty"`

					CallValue      int64 `json:"call_value,omitempty"`
					CallTokenValue int64 `json:"call_token_value,omitempty"`
					TokenId        int64 `json:"token_id,omitempty"`

					// --------------------------------- TRC -10 -------------------------------------------
					Name         string `json:",omitempty"`             // TRC-10 token name
					Abbr         string `json:"abbr,omitempty"`         // TRC-10 token symbol
					TotalSupply  int64  `json:"total_supply,omitempty"` // Total supply.
					FrozenSupply []struct {
						FrozenAmount int64 `json:"frozen_amount,omitempty"`
						FrozenDays   int64 `json:"frozen_days,omitempty"`
					} `json:"frozen_supply,omitempty"`
					TrxNum                  int32  `json:"trx_num,omitempty"`
					Precision               int32  `json:"precision,omitempty"`
					Num                     int32  `json:"num,omitempty"`
					StartTime               int64  `json:"start_time,omitempty"`
					EndTime                 int64  `json:"end_time,omitempty"`
					Description             string `json:"description,omitempty"`
					Url                     string `json:"url,omitempty"`
					FreeAssetNetLimit       int64  `json:"free_asset_net_limit,omitempty"`
					PublicFreeAssetNetLimit int64  `json:"public_free_asset_net_limit,omitempty"`
					PublicLatestFreeNetTime int64  `json:"public_latest_free_net_time,omitempty"`
					NewLimit                int64  `json:"new_limit,omitempty"`
					NewPublicLimit          int64  `json:"new_public_limit,omitempty"`
					// --------------------------------- TRC -10 -------------------------------------------

					// --------------------------------- Proposals -------------------------------------------
					Parameters    map[int64]int64 `json:"parameters,omitempty"`
					ProposalId    int64           `json:"proposal_id,omitempty"`
					IsAddApproval bool            `json:"is_add_approval,omitempty"`
					// --------------------------------- Proposals -------------------------------------------

					// --------------------------------- witness -------------------------------------------
					UpdateUrl string `json:"update_url,omitempty"`
					Brokerage int32  `json:"brokerage,omitempty"`
					Votes     []struct {
						VoteAddress string `json:"vote_address,omitempty"`
						VoteCount   int64  `json:"vote_count,omitempty"`
					} `json:"votes,omitempty"`
					// --------------------------------- witness -------------------------------------------

					NewContract                string `json:"new_contract,omitempty"`
					ConsumeUserResourcePercent int64  `json:"consume_user_resource_percent,omitempty"`
					OriginEnergyLimit          int64  `json:"origin_energy_limit,omitempty"`
				} `json:"value,omitempty"`
				TypeURL string `json:"type_url,omitempty"`
			} `json:"parameter,omitempty"`
			Type string `json:"type,omitempty"`
		} `json:"contract,omitempty"`
		RefBlockBytes string `json:"ref_block_bytes,omitempty"`
		RefBlockHash  string `json:"ref_block_hash,omitempty"`
		Expiration    int64  `json:"expiration,omitempty"`
		FeeLimit      int64  `json:"fee_limit,omitempty"`
		Timestamp     int64  `json:"timestamp,omitempty"`
	} `json:"raw_data,omitempty"`
	RawDataHex string `json:"raw_data_hex,omitempty"`
	Error      string `json:"Error,omitempty,omitempty"`
}

type TransactionInfoByIdInternet struct {
	ID              string   `json:"id,omitempty"`
	Fee             int64    `json:"fee,omitempty"`
	BlockNumber     int64    `json:"blockNumber,omitempty"`
	BlockTimeStamp  int64    `json:"blockTimeStamp,omitempty"`
	ContractResult  []string `json:"contractResult,omitempty"`
	ContractAddress string   `json:"contract_address,omitempty"`
	Receipt         struct {
		EnergyUsage        int64  `json:"energy_usage,omitempty"`
		EnergyFee          int64  `json:"energy_fee,omitempty"`
		OriginEnergyUsage  int64  `json:"origin_energy_usage,omitempty"`
		EnergyUsaotal      int64  `json:"energy_usage_total,omitempty"`
		NetUsage           int64  `json:"net_usage,omitempty"`
		NetFee             int64  `json:"net_fee,omitempty"`
		Result             string `json:"result,omitempty"`
		EnergyPenaltyTotal int64  `json:"energy_penalty_total,omitempty"`
	} `json:"receipt,omitempty"`
	Log []struct {
		Address string   `json:"address,omitempty"`
		Topics  []string `json:"topics,omitempty"`
		Data    string   `json:"data,omitempty"`
	} `json:"log,omitempty"`
	Result               string `json:"result,omitempty"`
	ResMessage           string `json:"resMessage,omitempty"`
	AssetIssueID         string `json:"assetIssueID,omitempty"`
	WithdrawAmount       int64  `json:"withdraw_amount,omitempty"`
	UnfreezeAmount       int64  `json:"unfreeze_amount,omitempty"`
	InternalTransactions []struct {
		Hash              string `json:"hash,omitempty"`
		CallerAddress     string `json:"caller_address,omitempty"`
		TransferToAddress string `json:"transferTo_address,omitempty"`
		CallValueInfo     []struct {
			CallValue int64 `json:"callValue,omitempty"`
			TokenId   int64 `json:"tokenId,omitempty"`
		}
	} `json:"internal_transactions,omitempty"`
	WithdrawExpireAmount          int64            `json:"withdraw_expire_amount,omitempty"`
	CancelUnfreezeV2Amount        map[string]int64 `json:"cancel_unfreeze_v_2_amount,omitempty"`
	ExchangeReceivedAmount        int64            `json:"exchange_received_amount,omitempty"`
	ExchangeInjectAnotherAmount   int64            `json:"exchange_inject_another_amount,omitempty"`
	ExchangeWithdrawAnotherAmount int64            `json:"exchange_withdraw_another_amount,omitempty"`
	ExchangeId                    int64            `json:"exchange_id,omitempty"`
	ShieldedTransactionFee        int64            `json:"shielded_transaction_fee,omitempty"`
}

type ContractInternet struct {
	OriginAddress              string                 `json:"origin_address,omitempty"`
	ContractAddress            string                 `json:"contract_address,omitempty"`
	ABI                        map[string]interface{} `json:"abi,omitempty"`
	Bytecode                   string                 `json:"bytecode,omitempty"`
	CallValue                  int64                  `json:"call_value,omitempty"`
	ConsumeUserResourcePercent int64                  `json:"consume_user_resource_percent,omitempty"`
	Name                       string                 `json:"name,omitempty"`
	OriginEnergyLimit          int64                  `json:"origin_energy_limit,omitempty"`
	CodeHash                   string                 `json:"code_hash,omitempty"`
}

type ContractInfoInternet struct {
	RuntimeCode   string           `json:"runtimecode,omitempty"`
	SmartContract ContractInternet `json:"smart_contract,omitempty"`
	ContractState struct {
		EnergyUsage  int64 `json:"energy_usage,omitempty"`
		EnergyFactor int64 `json:"energy_factor,omitempty"`
		UpdateCycle  int64 `json:"update_cycle,omitempty"`
	} `json:"contract_state,omitempty"`
	Error string `json:"Error,omitempty"`
}

type TriggerConstantContractInternet struct {
	Result         Result      `json:"result,omitempty"`
	EnergyUsed     int64       `json:"energy_used,omitempty"`
	EnergyPenalty  int64       `json:"energy_penalty,omitempty"`
	ConstantResult []string    `json:"constant_result,omitempty"`
	Transaction    Transaction `json:"transaction,omitempty"`
}

type ListNodesInternet struct {
	Nodes []*struct {
		Address struct {
			Host string `json:"host,omitempty"`
			Port int    `json:"port,omitempty"`
		} `json:"address"`
	} `json:"nodes"`
}

type NodeInfoInternet struct {
	ActiveConnectCount  int64  `json:"activeConnectCount,omitempty"`
	BeginSyncNum        int64  `json:"beginSyncNum,omitempty"`
	Block               string `json:"block,omitempty"`
	CheatWitnessInfoMap struct {
	} `json:"cheatWitnessInfoMap,omitempty"`
	ConfigNodeInfo struct {
		ActiveNodeSize           int32   `json:"activeNodeSize,omitempty"`
		AllowAdaptiveEnergy      int32   `json:"allowAdaptiveEnergy,omitempty"`
		AllowCreationOfContracts int32   `json:"allowCreationOfContracts,omitempty"`
		BackupListenPort         int32   `json:"backupListenPort,omitempty"`
		BackupMemberSize         int32   `json:"backupMemberSize,omitempty"`
		BackupPriority           int32   `json:"backupPriority,omitempty"`
		CodeVersion              string  `json:"codeVersion,omitempty"`
		DbVersion                int32   `json:"dbVersion,omitempty"`
		DiscoverEnable           bool    `json:"discoverEnable,omitempty"`
		ListenPort               int32   `json:"listenPort,omitempty"`
		MaxConnectCount          int32   `json:"maxConnectCount,omitempty"`
		MaxTimeRatio             float64 `json:"maxTimeRatio,omitempty"`
		MinParticipationRate     int32   `json:"minParticipationRate,omitempty"`
		MinTimeRatio             float64 `json:"minTimeRatio,omitempty"`
		P2PVersion               string  `json:"p2pVersion,omitempty"`
		PassiveNodeSize          int32   `json:"passiveNodeSize,omitempty"`
		SameIPMaxConnectCount    int32   `json:"sameIpMaxConnectCount,omitempty"`
		SendNodeSize             int32   `json:"sendNodeSize,omitempty"`
		SupportConstant          bool    `json:"supportConstant,omitempty"`
		VersionNum               string  `json:"versionNum,omitempty"`
	} `json:"configNodeInfo,omitempty"`
	CurrentConnectCount int64 `json:"currentConnectCount,omitempty"`
	MachineInfo         struct {
		CPUCount               int64         `json:"cpuCount,omitempty"`
		CPURate                float64       `json:"cpuRate,omitempty"`
		DeadLockThreadCount    int64         `json:"deadLockThreadCount,omitempty"`
		DeadLockThreadInfoList []interface{} `json:"deadLockThreadInfoList,omitempty"`
		FreeMemory             int64         `json:"freeMemory,omitempty"`
		JavaVersion            string        `json:"javaVersion,omitempty"`
		JvmFreeMemory          int64         `json:"jvmFreeMemory,omitempty"`
		JvmTotalMemory         int64         `json:"jvmTotalMemory,omitempty"`
		MemoryDescInfoList     []struct {
			InitSize int64   `json:"initSize,omitempty"`
			MaxSize  int64   `json:"maxSize,omitempty"`
			Name     string  `json:"name,omitempty"`
			UseRate  float64 `json:"useRate,omitempty"`
			UseSize  int64   `json:"useSize,omitempty"`
		} `json:"memoryDescInfoList,omitempty"`
		OsName         string  `json:"osName,omitempty"`
		ProcessCPURate float64 `json:"processCpuRate,omitempty"`
		ThreadCount    int64   `json:"threadCount,omitempty"`
		TotalMemory    int64   `json:"totalMemory,omitempty"`
	} `json:"machineInfo,omitempty"`
	PassiveConnectCount int `json:"passiveConnectCount,omitempty"`
	PeerList            []struct {
		Active                  bool    `json:"active,omitempty"`
		AvgLatency              float64 `json:"avgLatency,omitempty"`
		BlockInPorcSize         int64   `json:"blockInPorcSize,omitempty"`
		ConnectTime             int64   `json:"connectTime,omitempty"`
		DisconnectTimes         int64   `json:"disconnectTimes,omitempty"`
		HeadBlockTimeWeBothHave int64   `json:"headBlockTimeWeBothHave,omitempty"`
		HeadBlockWeBothHave     string  `json:"headBlockWeBothHave,omitempty"`
		Host                    string  `json:"host,omitempty,omitempty"`
		InFlow                  int64   `json:"inFlow,omitempty,omitempty"`
		LastBlockUpdateTime     int64   `json:"lastBlockUpdateTime,omitempty"`
		LastSyncBlock           string  `json:"lastSyncBlock,omitempty"`
		LocalDisconnectReason   string  `json:"localDisconnectReason,omitempty"`
		NeedSyncFromPeer        bool    `json:"needSyncFromPeer,omitempty"`
		NeedSyncFromUs          bool    `json:"needSyncFromUs,omitempty"`
		NodeCount               int64   `json:"nodeCount,omitempty"`
		NodeID                  string  `json:"nodeId,omitempty"`
		Port                    int64   `json:"port,omitempty"`
		RemainNum               int64   `json:"remainNum,omitempty"`
		RemoteDisconnectReason  string  `json:"remoteDisconnectReason,omitempty"`
		Score                   int64   `json:"score,omitempty"`
		SyncBlockRequestedSize  int64   `json:"syncBlockRequestedSize,omitempty"`
		SyncFlag                bool    `json:"syncFlag,omitempty"`
		SyncToFetchSize         int64   `json:"syncToFetchSize,omitempty"`
		SyncToFetchSizePeekNum  int64   `json:"syncToFetchSizePeekNum,omitempty"`
		UnFetchSynNum           int64   `json:"unFetchSynNum,omitempty"`
	} `json:"peerList,omitempty"`
	SolidityBlock string `json:"solidityBlock,omitempty"`
	TotalFlow     int64  `json:"totalFlow,omitempty"`
}

type ChainParametersInternet struct {
	ChainParameter []struct {
		Key   string `json:"key"`
		Value int    `json:"value"`
	} `json:"chainParameter"`
}

type ApproveListInternet struct {
	Result struct {
		Code    int    `json:"code,omitempty"`
		Message string `json:"message,omitempty"`
	} `json:"result,omitempty"`
	ApprovedList []string    `json:"approved_list,omitempty"`
	Transaction  Transaction `json:"transaction,omitempty"`
	Error        string      `json:"Error,omitempty"`
}

type AccountInternet struct {
	Address            string `json:"address,omitempty"`
	Balance            int    `json:"balance,omitempty"`
	CreateTime         int64  `json:"create_time,omitempty"`
	LatestConsumeTime  int64  `json:"latest_consume_time,omitempty"`
	NetWindowSize      int    `json:"net_window_size,omitempty"`
	NetWindowOptimized bool   `json:"net_window_optimized,omitempty"`
	AccountResource    struct {
		LatestConsumeTimeForEnergy                int64 `json:"latest_consume_time_for_energy,omitempty"`
		EnergyWindowSize                          int   `json:"energy_window_size,omitempty"`
		AcquiredDelegatedFrozenV2BalanceForEnergy int   `json:"acquired_delegated_frozenV2_balance_for_energy,omitempty"`
		EnergyWindowOptimized                     bool  `json:"energy_window_optimized,omitempty"`
	} `json:"account_resource,omitempty"`
	OwnerPermission struct {
		PermissionName string `json:"permission_name,omitempty"`
		Threshold      int    `json:"threshold,omitempty"`
		Keys           []struct {
			Address string `json:"address,omitempty"`
			Weight  int    `json:"weight,omitempty"`
		} `json:"keys,omitempty"`
	} `json:"owner_permission,omitempty"`
	ActivePermission []struct {
		Type           string `json:"type,omitempty"`
		ID             int    `json:"id,omitempty"`
		PermissionName string `json:"permission_name,omitempty"`
		Threshold      int    `json:"threshold,omitempty"`
		Operations     string `json:"operations,omitempty"`
		Keys           []struct {
			Address string `json:"address,omitempty"`
			Weight  int    `json:"weight,omitempty"`
		} `json:"keys,omitempty"`
	} `json:"active_permission,omitempty"`
	FrozenV2 []struct {
		Type string `json:"type,omitempty,omitempty"`
	} `json:"frozenV2,omitempty"`
	AssetV2 []struct {
		Key   string `json:"key,omitempty"`
		Value int64  `json:"value,omitempty"`
	} `json:"assetV2,omitempty"`
	FreeAssetNetUsageV2 []struct {
		Key   string `json:"key,omitempty"`
		Value int    `json:"value,omitempty"`
	} `json:"free_asset_net_usageV2,omitempty"`
	AssetOptimized bool `json:"asset_optimized,omitempty"`
}

type AssetIssueInternet struct {
	Id           string `json:"id,omitempty"`
	OwnerAddress string `json:"owner_address,omitempty"`
	Name         string `json:"name,omitempty"`
	Abbr         string `json:"abbr,omitempty"`
	TotalSupply  int64  `json:"total_supply,omitempty"`
	FrozenSupply []struct {
		FrozenAmount int64 `json:"frozen_amount,omitempty"`
		FrozenDays   int64 `json:"frozen_days,omitempty"`
	} `json:"frozen_supply,omitempty"`
	TrxNum                  int32  `json:"trx_num,omitempty"`
	Precision               int32  `json:"precision,omitempty"`
	Num                     int32  `json:"num,omitempty"`
	StartTime               int64  `json:"start_time,omitempty"`
	EndTime                 int64  `json:"end_time,omitempty"`
	VoteScore               int32  `json:"vote_score,omitempty"`
	Description             string `json:"description,omitempty"`
	Url                     string `json:"url,omitempty"`
	FreeAssetNetLimit       int64  `json:"free_asset_net_limit,omitempty"`
	PublicFreeAssetNetLimit int64  `json:"public_free_asset_net_limit,omitempty"`
	PublicFreeAssetNetUsage int64  `json:"public_free_asset_net_usage,omitempty"`
	PublicLatestFreeNetTime int64  `json:"public_latest_free_net_time,omitempty"`
	Error                   string `json:"Error,omitempty"`
}

type AssetIssueListInternet struct {
	AssetIssue []AssetIssueInternet `json:"assetIssue,omitempty"`
	Error      string               `json:"Error,omitempty"`
}

type ProposalListInternet struct {
	Proposals []ProposalInternet `json:"proposals,omitempty"`
	Error     string             `json:"Error,omitempty"`
}
type ProposalInternet struct {
	ProposalID      int    `json:"proposal_id,omitempty"`
	ProposerAddress string `json:"proposer_address,omitempty"`
	Parameters      []struct {
		Key   int `json:"key,omitempty"`
		Value int `json:"value,omitempty"`
	} `json:"parameters,omitempty"`
	ExpirationTime int64    `json:"expiration_time,omitempty"`
	CreateTime     int64    `json:"create_time,omitempty"`
	Approvals      []string `json:"approvals,omitempty"`
	State          string   `json:"state,omitempty"` // state of the proposal. (Enum: PENDING, DISAPPROVED, APPROVED, CANCELED)state
	Error          string   `json:"Error,omitempty"`
}

type WitnessListInternet struct {
	Witnesses []WitnessInternet `json:"witnesses,omitempty"`
	Error     string            `json:"Error,omitempty"`
}

type WitnessInternet struct {
	Address        string `json:"address,omitempty"`
	VoteCount      int64  `json:"voteCount,omitempty"`
	Url            string `json:"url,omitempty"`
	TotalProduced  int64  `json:"totalProduced,omitempty"`
	TotalMissed    int64  `json:"totalMissed,omitempty"`
	LatestBlockNum int64  `json:"latestBlockNum,omitempty"`
	LatestSlotNum  int64  `json:"latestSlotNum,omitempty"`
	IsJobs         bool   `json:"isJobs,omitempty"`
	Error          string `json:"Error,omitempty"`
}
