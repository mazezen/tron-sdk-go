package client

import (
	"fmt"

	"github.com/golang/protobuf/proto"
	tronpb "github.com/mazezen/tron-sdk-go/pb/tron"
)

// GetDelegatedResourceAccountIndex Query the resource delegation by an account during stake1.0 phase.
// i.e. list all addresses that have delegated resources to an account
// value: base58 (T...) | hex (tron hex 41...) | eth hex (0x...)
// https://developers.tron.network/reference/getdelegatedresourceaccountindex
func (c *GrpcClient) GetDelegatedResourceAccountIndex(value string) (*tronpb.DelegatedResourceAccountIndex, error) {
	var err error
	req := new(tronpb.BytesMessage)
	req.Value, err = c.convert(value)
	if err != nil {
		return nil, fmt.Errorf("invalid value address: %w", err)
	}

	ctx, cancelFunc := c.getContext()
	defer cancelFunc()

	return c.WalletClient.GetDelegatedResourceAccountIndex(ctx, req)
}

// GetDelegatedResourceAccountIndexV2 In Stake2.0, query the resource delegation index of an account.
// Two lists will return, one is the list of addresses the account has delegated its resources(toAddress),
// and the other is the list of addresses that have delegated resources to the account(fromAddress).
// value: base58 (T...) | hex (tron hex 41...) | eth hex (0x...)
// https://developers.tron.network/reference/getdelegatedresourceaccountindexv2-1
func (c *GrpcClient) GetDelegatedResourceAccountIndexV2(value string) (*tronpb.DelegatedResourceAccountIndex, error) {
	var err error
	req := new(tronpb.BytesMessage)
	req.Value, err = c.convert(value)
	if err != nil {
		return nil, fmt.Errorf("invalid value address: %w", err)
	}

	ctx, cancelFunc := c.getContext()
	defer cancelFunc()

	return c.WalletClient.GetDelegatedResourceAccountIndexV2(ctx, req)
}

// GetDelegatedResource
// Returns all resources delegations during stake1.0 phase from an account to another account.
// The to address can be retrieved from the GetDelegatedResourceAccountIndex API. (Confirmed state)
// from: base58 (T...) | hex (tron hex 41...) | eth hex (0x...)
// https://developers.tron.network/reference/getdelegatedresource-1
func (c *GrpcClient) GetDelegatedResource(from string) ([]*tronpb.DelegatedResourceList, error) {
	var err error
	req := new(tronpb.DelegatedResourceMessage)
	req.FromAddress, err = c.convert(from)
	if err != nil {
		return nil, fmt.Errorf("invalid from address: %w", err)
	}
	ai, err := c.GetDelegatedResourceAccountIndex(from)
	if err != nil {
		return nil, err
	}

	ctx, cancelFunc := c.getContext()
	defer cancelFunc()

	result := make([]*tronpb.DelegatedResourceList, len(ai.GetToAccounts()))
	for _, addrTo := range ai.GetToAccounts() {
		req.ToAddress = addrTo
		resource, err := c.WalletClient.GetDelegatedResource(ctx, req)
		if err != nil {
			return nil, err
		}
		result = append(result, resource)
	}

	return result, nil
}

// GetDelegatedResourceV2 In Stake2.0, query the detail of resource share delegated from fromAddress to toAddress.
// The to address can be retrieved from the GetDelegatedResourceAccountIndexV2 API. (Confirmed state)
// from: base58 (T...) | hex (tron hex 41...) | eth hex (0x...)
// https://developers.tron.network/reference/getdelegatedresourcev2
func (c *GrpcClient) GetDelegatedResourceV2(from string) ([]*tronpb.DelegatedResourceList, error) {
	var err error
	req := new(tronpb.DelegatedResourceMessage)
	req.FromAddress, err = c.convert(from)
	if err != nil {
		return nil, fmt.Errorf("invalid from address: %w", err)
	}
	ai, err := c.GetDelegatedResourceAccountIndexV2(from)
	if err != nil {
		return nil, err
	}

	ctx, cancelFunc := c.getContext()
	defer cancelFunc()

	result := make([]*tronpb.DelegatedResourceList, len(ai.GetToAccounts()))
	for _, addrTo := range ai.GetToAccounts() {
		req.ToAddress = addrTo
		resource, err := c.WalletClient.GetDelegatedResourceV2(ctx, req)
		if err != nil {
			return nil, err
		}
		result = append(result, resource)
	}

	return result, nil
}

// GetDelegatedResourceByGivenTo
// In Stake1.0, Specify the receiving resource address To and query all resources delegated to the To address
// from: base58 (T...) | hex (tron hex 41...) | eth hex (0x...)
// to: base58 (T...) | hex (tron hex 41...) | eth hex (0x...)
func (c *GrpcClient) GetDelegatedResourceByGivenTo(from, to string) (*tronpb.DelegatedResourceList, error) {
	var err error
	req := new(tronpb.DelegatedResourceMessage)

	if req.FromAddress, err = c.convert(from); err != nil {
		return nil, fmt.Errorf("invalid from address: %w", err)
	}
	if req.ToAddress, err = c.convert(to); err != nil {
		return nil, fmt.Errorf("invalid to address: %w", err)
	}

	ctx, cancelFunc := c.getContext()
	defer cancelFunc()

	return c.WalletClient.GetDelegatedResource(ctx, req)
}

// GetDelegatedResourceV2ByGivenTo
// In Stake2.0, Specify the receiving resource address To and query all resources delegated to the To address
// from: base58 (T...) | hex (tron hex 41...) | eth hex (0x...)
// to: base58 (T...) | hex (tron hex 41...) | eth hex (0x...)
func (c *GrpcClient) GetDelegatedResourceV2ByGivenTo(from, to string) (*tronpb.DelegatedResourceList, error) {
	var err error
	req := new(tronpb.DelegatedResourceMessage)
	if req.FromAddress, err = c.convert(from); err != nil {
		return nil, fmt.Errorf("invalid from address: %w", err)
	}
	if req.ToAddress, err = c.convert(to); err != nil {
		return nil, fmt.Errorf("invalid to address: %w", err)
	}

	ctx, cancelFunc := c.getContext()
	defer cancelFunc()

	return c.WalletClient.GetDelegatedResourceV2(ctx, req)
}

// GetCanDelegatedMaxSize In Stake2.0, query the amount of delegatable resources share of the specified resource type for an address, unit is sun.
// addr: base58 (T...) | hex (tron hex 41...) | eth hex (0x...)
// resourceType: Resource type, 0 is bandwidth, 1 is energy.
// https://developers.tron.network/reference/getcandelegatedmaxsize
func (c *GrpcClient) GetCanDelegatedMaxSize(addr string, resourceType int32) (*tronpb.CanDelegatedMaxSizeResponseMessage, error) {
	var err error
	var req = new(tronpb.CanDelegatedMaxSizeRequestMessage)
	req.OwnerAddress, err = c.convert(addr)
	if err != nil {
		return nil, fmt.Errorf("invalid owner address: %w", err)
	}
	req.Type = resourceType

	ctx, cancelFunc := c.getContext()
	defer cancelFunc()

	return c.WalletClient.GetCanDelegatedMaxSize(ctx, req)
}

// DelegateResource Delegate bandwidth or energy resources to other accounts in Stake2.0.
// https://developers.tron.network/reference/delegateresource-1
func (c *GrpcClient) DelegateResource(
	from, to string, // from: resource sender address  to: resource receiver address
	core tronpb.ResourceCode, // Type of resource. (Enum: "BANDWIDTH" or "ENERGY")
	balance int64, // Amount of TRX to delegate or undelegate for resources. (Unit: sun)
	lock bool, // Optional. Whether to lock the resource delegation. If true, the delegation cannot be canceled during the lock_period.
	lockPeriod int64, // Lock duration in blocks (1 block ≈ 3 seconds). Only valid if lock is true. (e.g., 1 day = 28800)
) (*tronpb.TransactionExtention, error) {
	var err error
	var req = new(tronpb.DelegateResourceContract)

	req.OwnerAddress, err = c.convert(from)
	if err != nil {
		return nil, fmt.Errorf("invalid owner address: %w", err)
	}
	req.ReceiverAddress, err = c.convert(to)
	if err != nil {
		return nil, fmt.Errorf("invalid receiver address: %w", err)
	}

	req.Resource = core
	req.Balance = balance
	req.Lock = lock
	req.LockPeriod = lockPeriod

	ctx, cancelFunc := c.getContext()
	defer cancelFunc()

	tx, err := c.WalletClient.DelegateResource(ctx, req)
	if err != nil {
		return nil, err
	}

	if proto.Size(tx) == 0 {
		return nil, fmt.Errorf("transaction empty")
	}

	if tx.GetResult().GetCode() != tronpb.Return_SUCCESS {
		return nil, fmt.Errorf("%s", tx.GetResult().GetMessage())
	}

	return tx, nil
}

// UnDelegateResource
// Cancel the delegation of bandwidth or energy resources to other accounts in Stake2.0.
// https://developers.tron.network/reference/undelegateresource-1
func (c *GrpcClient) UnDelegateResource(
	from, to string, // from: resource sender address  to: resource receiver address
	code tronpb.ResourceCode, // Type of resource. (Enum: "BANDWIDTH" or "ENERGY")
	balance int64, // Amount of TRX to delegate or undelegate for resources. (Unit: sun)
) (*tronpb.TransactionExtention, error) {
	var err error
	var req = new(tronpb.UnDelegateResourceContract)

	req.OwnerAddress, err = c.convert(from)
	if err != nil {
		return nil, fmt.Errorf("invalid owner address: %w", err)
	}
	req.ReceiverAddress, err = c.convert(to)
	if err != nil {
		return nil, fmt.Errorf("invalid receiver address: %w", err)
	}
	req.Resource = code
	req.Balance = balance

	ctx, cancelFunc := c.getContext()
	defer cancelFunc()

	tx, err := c.WalletClient.UnDelegateResource(ctx, req)
	if err != nil {
		return nil, err
	}
	if proto.Size(tx) == 0 {
		return nil, fmt.Errorf("transaction empty")
	}
	if tx.GetResult().GetCode() != tronpb.Return_SUCCESS {
		return nil, fmt.Errorf("%s", tx.GetResult().GetMessage())
	}

	return tx, nil
}
