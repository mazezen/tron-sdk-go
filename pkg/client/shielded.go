package client

import tronpb "github.com/mazezen/tron-sdk-go/pb/tron"

// =============================================
// for shieldedTransaction (隐私交易相关)
// =============================================

// CreateShieldedTransaction 创建隐私交易（带 spend auth sig）
func (c *GrpcClient) CreateShieldedTransaction(in *tronpb.PrivateParameters) (*tronpb.TransactionExtention, error) {
	ctx, cancel := c.getContext()
	defer cancel()
	return c.WalletClient.CreateShieldedTransaction(ctx, in)
}

// GetMerkleTreeVoucherInfo 获取 Merkle 树凭证信息
func (c *GrpcClient) GetMerkleTreeVoucherInfo(in *tronpb.OutputPointInfo) (*tronpb.IncrementalMerkleVoucherInfo, error) {
	ctx, cancel := c.getContext()
	defer cancel()
	return c.WalletClient.GetMerkleTreeVoucherInfo(ctx, in)
}

// ScanNoteByIvk 通过 IVK 扫描笔记（未标记）
func (c *GrpcClient) ScanNoteByIvk(in *tronpb.IvkDecryptParameters) (*tronpb.DecryptNotes, error) {
	ctx, cancel := c.getContext()
	defer cancel()
	return c.WalletClient.ScanNoteByIvk(ctx, in)
}

// ScanAndMarkNoteByIvk 通过 IVK 扫描并标记笔记
func (c *GrpcClient) ScanAndMarkNoteByIvk(in *tronpb.IvkDecryptAndMarkParameters) (*tronpb.DecryptNotesMarked, error) {
	ctx, cancel := c.getContext()
	defer cancel()
	return c.WalletClient.ScanAndMarkNoteByIvk(ctx, in)
}

// ScanNoteByOvk 通过 OVK 扫描笔记
func (c *GrpcClient) ScanNoteByOvk(in *tronpb.OvkDecryptParameters) (*tronpb.DecryptNotes, error) {
	ctx, cancel := c.getContext()
	defer cancel()
	return c.WalletClient.ScanNoteByOvk(ctx, in)
}

// GetSpendingKey 获取随机 spending key
func (c *GrpcClient) GetSpendingKey() (*tronpb.BytesMessage, error) {
	ctx, cancel := c.getContext()
	defer cancel()
	return c.WalletClient.GetSpendingKey(ctx, new(tronpb.EmptyMessage))
}

// GetExpandedSpendingKey 从 spending key 扩展出 ask/nsk/ovk
func (c *GrpcClient) GetExpandedSpendingKey(in *tronpb.BytesMessage) (*tronpb.ExpandedSpendingKeyMessage, error) {
	ctx, cancel := c.getContext()
	defer cancel()
	return c.WalletClient.GetExpandedSpendingKey(ctx, in)
}

// GetAkFromAsk 从 ask 获取 ak
func (c *GrpcClient) GetAkFromAsk(in *tronpb.BytesMessage) (*tronpb.BytesMessage, error) {
	ctx, cancel := c.getContext()
	defer cancel()
	return c.WalletClient.GetAkFromAsk(ctx, in)
}

// GetNkFromNsk 从 nsk 获取 nk
func (c *GrpcClient) GetNkFromNsk(in *tronpb.BytesMessage) (*tronpb.BytesMessage, error) {
	ctx, cancel := c.getContext()
	defer cancel()
	return c.WalletClient.GetNkFromNsk(ctx, in)
}

// GetIncomingViewingKey 从 ask + nsk 获取 ivk
func (c *GrpcClient) GetIncomingViewingKey(in *tronpb.ViewingKeyMessage) (*tronpb.IncomingViewingKeyMessage, error) {
	ctx, cancel := c.getContext()
	defer cancel()
	return c.WalletClient.GetIncomingViewingKey(ctx, in)
}

// GetDiversifier 获取随机 diversifier d
func (c *GrpcClient) GetDiversifier() (*tronpb.DiversifierMessage, error) {
	ctx, cancel := c.getContext()
	defer cancel()
	return c.WalletClient.GetDiversifier(ctx, new(tronpb.EmptyMessage))
}

// GetNewShieldedAddress 生成新的屏蔽地址（完整 shielded address）
func (c *GrpcClient) GetNewShieldedAddress() (*tronpb.ShieldedAddressInfo, error) {
	ctx, cancel := c.getContext()
	defer cancel()
	return c.WalletClient.GetNewShieldedAddress(ctx, new(tronpb.EmptyMessage))
}

// GetZenPaymentAddress 从 ivk + d 生成 payment address
func (c *GrpcClient) GetZenPaymentAddress(in *tronpb.IncomingViewingKeyDiversifierMessage) (*tronpb.PaymentAddressMessage, error) {
	ctx, cancel := c.getContext()
	defer cancel()
	return c.WalletClient.GetZenPaymentAddress(ctx, in)
}

// GetRcm 获取随机 rcm（用于生成 note）
func (c *GrpcClient) GetRcm() (*tronpb.BytesMessage, error) {
	ctx, cancel := c.getContext()
	defer cancel()
	return c.WalletClient.GetRcm(ctx, new(tronpb.EmptyMessage))
}

// IsSpend 检查某个 note 是否已被花费
func (c *GrpcClient) IsSpend(in *tronpb.NoteParameters) (*tronpb.SpendResult, error) {
	ctx, cancel := c.getContext()
	defer cancel()
	return c.WalletClient.IsSpend(ctx, in)
}

// CreateShieldedTransactionWithoutSpendAuthSig 创建隐私交易（不带 spend auth sig，需要后续签名）
func (c *GrpcClient) CreateShieldedTransactionWithoutSpendAuthSig(in *tronpb.PrivateParametersWithoutAsk) (*tronpb.TransactionExtention, error) {
	ctx, cancel := c.getContext()
	defer cancel()
	return c.WalletClient.CreateShieldedTransactionWithoutSpendAuthSig(ctx, in)
}

// GetShieldTransactionHash 获取隐私交易的 hash（用于签名）
func (c *GrpcClient) GetShieldTransactionHash(tx *tronpb.Transaction) (*tronpb.BytesMessage, error) {
	ctx, cancel := c.getContext()
	defer cancel()
	return c.WalletClient.GetShieldTransactionHash(ctx, tx)
}

// CreateSpendAuthSig 创建 spend authorization signature
func (c *GrpcClient) CreateSpendAuthSig(in *tronpb.SpendAuthSigParameters) (*tronpb.BytesMessage, error) {
	ctx, cancel := c.getContext()
	defer cancel()
	return c.WalletClient.CreateSpendAuthSig(ctx, in)
}

// CreateShieldNullifier 生成 nullifier（用于防止双花）
func (c *GrpcClient) CreateShieldNullifier(in *tronpb.NfParameters) (*tronpb.BytesMessage, error) {
	ctx, cancel := c.getContext()
	defer cancel()
	return c.WalletClient.CreateShieldNullifier(ctx, in)
}

// =============================================
// for shielded contract (屏蔽 TRC-20 相关)
// =============================================

// CreateShieldedContractParameters 创建屏蔽 TRC-20 合约参数（带 ask）
func (c *GrpcClient) CreateShieldedContractParameters(in *tronpb.PrivateShieldedTRC20Parameters) (*tronpb.ShieldedTRC20Parameters, error) {
	ctx, cancel := c.getContext()
	defer cancel()
	return c.WalletClient.CreateShieldedContractParameters(ctx, in)
}

// CreateShieldedContractParametersWithoutAsk 创建屏蔽 TRC-20 合约参数（不带 ask）
func (c *GrpcClient) CreateShieldedContractParametersWithoutAsk(in *tronpb.PrivateShieldedTRC20ParametersWithoutAsk) (*tronpb.ShieldedTRC20Parameters, error) {
	ctx, cancel := c.getContext()
	defer cancel()
	return c.WalletClient.CreateShieldedContractParametersWithoutAsk(ctx, in)
}

// ScanShieldedTRC20NotesByIvk 通过 ivk 扫描屏蔽 TRC-20 笔记
func (c *GrpcClient) ScanShieldedTRC20NotesByIvk(in *tronpb.IvkDecryptTRC20Parameters) (*tronpb.DecryptNotesTRC20, error) {
	ctx, cancel := c.getContext()
	defer cancel()
	return c.WalletClient.ScanShieldedTRC20NotesByIvk(ctx, in)
}

// ScanShieldedTRC20NotesByOvk 通过 ovk 扫描屏蔽 TRC-20 笔记
func (c *GrpcClient) ScanShieldedTRC20NotesByOvk(in *tronpb.OvkDecryptTRC20Parameters) (*tronpb.DecryptNotesTRC20, error) {
	ctx, cancel := c.getContext()
	defer cancel()
	return c.WalletClient.ScanShieldedTRC20NotesByOvk(ctx, in)
}

// IsShieldedTRC20ContractNoteSpent 检查屏蔽 TRC-20 note 是否已被花费
func (c *GrpcClient) IsShieldedTRC20ContractNoteSpent(in *tronpb.NfTRC20Parameters) (*tronpb.NullifierResult, error) {
	ctx, cancel := c.getContext()
	defer cancel()
	return c.WalletClient.IsShieldedTRC20ContractNoteSpent(ctx, in)
}

// GetTriggerInputForShieldedTRC20Contract 获取触发屏蔽 TRC-20 合约的 input data
func (c *GrpcClient) GetTriggerInputForShieldedTRC20Contract(in *tronpb.ShieldedTRC20TriggerContractParameters) (*tronpb.BytesMessage, error) {
	ctx, cancel := c.getContext()
	defer cancel()
	return c.WalletClient.GetTriggerInputForShieldedTRC20Contract(ctx, in)
}
