package client

import (
	"fmt"

	tronpb "github.com/mazezen/tron-sdk-go/pb/tron"
	"google.golang.org/protobuf/proto"
)

// CreateWitness apply to become a super representative candidate
// https://developers.tron.network/reference/createwitness
// Please use CreateWitness2 instead of this function.
func (c *GrpcClient) CreateWitness(from, url string) (*tronpb.Transaction, error) {
	var err error
	var in = new(tronpb.WitnessCreateContract)
	if in.OwnerAddress, err = c.convert(from); err != nil {
		return nil, err
	}
	in.Url = []byte(url)

	ctx, cancelFunc := c.getContext()
	defer cancelFunc()

	return c.WalletClient.CreateWitness(ctx, in)
}

// CreateWitness2 apply to become a super representative candidate
// https://developers.tron.network/reference/createwitness
// Use this function instead of CreateWitness.
func (c *GrpcClient) CreateWitness2(from, url string) (*tronpb.TransactionExtention, error) {
	var err error
	var in = new(tronpb.WitnessCreateContract)
	if in.OwnerAddress, err = c.convert(from); err != nil {
		return nil, err
	}
	in.Url = []byte(url)

	ctx, cancelFunc := c.getContext()
	defer cancelFunc()

	tx, err := c.WalletClient.CreateWitness2(ctx, in)
	if err != nil {
		return nil, err
	}

	if proto.Size(tx) == 0 {
		return nil, fmt.Errorf("empty transaction")
	}
	if tx.GetResult().GetCode() != 0 {
		return nil, fmt.Errorf("%v", tx.GetResult().GetMessage())
	}
	return tx, nil
}

// ListWitnesses List all Super Representatives
// https://developers.tron.network/reference/listwitnesses
func (c *GrpcClient) ListWitnesses() (*tronpb.WitnessList, error) {
	ctx, cancelFunc := c.getContext()
	defer cancelFunc()

	return c.WalletClient.ListWitnesses(ctx, &tronpb.EmptyMessage{})
}

// VoteWitnessAccount Vote for super representatives
// from: base58 (T...) | hex (tron hex 41...) | eth hex (0x...)
// votes: List of Vote objects, each containing vote_address (SR address) and vote_count (number of votes)
// SR address: base58 (T...) | hex (tron hex 41...) | eth hex (0x...)
// https://developers.tron.network/reference/votewitnessaccount
// Please use VoteWitnessAccount2 instead of this function.
func (c *GrpcClient) VoteWitnessAccount(from string, votes map[string]int64) (*tronpb.Transaction, error) {
	var err error
	var req = new(tronpb.VoteWitnessContract)
	if req.OwnerAddress, err = c.convert(from); err != nil {
		return nil, fmt.Errorf("invalid address: %w", err)
	}
	for k, v := range votes {
		voteAddress, err := c.convert(k)
		if err != nil {
			return nil, fmt.Errorf("invalid vote address: %w", err)
		}
		req.Votes = append(req.Votes, &tronpb.VoteWitnessContract_Vote{
			VoteAddress: voteAddress,
			VoteCount:   v,
		})
	}

	ctx, cancelFunc := c.getContext()
	defer cancelFunc()

	return c.WalletClient.VoteWitnessAccount(ctx, req)
}

// VoteWitnessAccount2 Vote for super representatives
// from: base58 (T...) | hex (tron hex 41...) | eth hex (0x...)
// votes: List of Vote objects, each containing vote_address (SR address) and vote_count (number of votes)
// SR address: base58 (T...) | hex (tron hex 41...) | eth hex (0x...)
// https://developers.tron.network/reference/votewitnessaccount
// Use this function instead of VoteWitnessAccount.
func (c *GrpcClient) VoteWitnessAccount2(from string, votes map[string]int64) (*tronpb.TransactionExtention, error) {
	var err error
	var req = new(tronpb.VoteWitnessContract)
	if req.OwnerAddress, err = c.convert(from); err != nil {
		return nil, fmt.Errorf("invalid address: %w", err)
	}
	for k, v := range votes {
		voteAddress, err := c.convert(k)
		if err != nil {
			return nil, fmt.Errorf("invalid vote address: %w", err)
		}
		req.Votes = append(req.Votes, &tronpb.VoteWitnessContract_Vote{
			VoteAddress: voteAddress,
			VoteCount:   v,
		})
	}

	ctx, cancelFunc := c.getContext()
	defer cancelFunc()

	tx, err := c.WalletClient.VoteWitnessAccount2(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("VoteWitnessAccount2: %w", err)
	}

	if proto.Size(tx) == 0 {
		return nil, fmt.Errorf("empty transaction")
	}

	if tx.GetResult().GetCode() != tronpb.Return_SUCCESS {
		return nil, fmt.Errorf("%s", tx.GetResult().GetMessage())
	}

	return tx, nil
}

// UpdateWitness edit the URL of the SR's official website
// https://developers.tron.network/reference/updatewitness
// Please use UpdateWitness2 instead of this function.
func (c *GrpcClient) UpdateWitness(from, url string) (*tronpb.Transaction, error) {
	var err error
	var in = new(tronpb.WitnessUpdateContract)

	if in.OwnerAddress, err = c.convert(from); err != nil {
		return nil, err
	}
	in.UpdateUrl = []byte(url)
	ctx, cancelFunc := c.getContext()
	defer cancelFunc()

	return c.WalletClient.UpdateWitness(ctx, in)
}

// UpdateWitness2 edit the URL of the SR's official website
// https://developers.tron.network/reference/updatewitness
// Use this function instead of UpdateWitness.
func (c *GrpcClient) UpdateWitness2(from, url string) (*tronpb.TransactionExtention, error) {
	var err error
	var in = new(tronpb.WitnessUpdateContract)

	if in.OwnerAddress, err = c.convert(from); err != nil {
		return nil, err
	}
	in.UpdateUrl = []byte(url)

	ctx, cancelFunc := c.getContext()
	defer cancelFunc()

	tx, err := c.WalletClient.UpdateWitness2(ctx, in)
	if err != nil {
		return nil, err
	}

	if proto.Size(tx) == 0 {
		return nil, fmt.Errorf("empty transaction")
	}

	if tx.GetResult().GetCode() != tronpb.Return_SUCCESS {
		return nil, fmt.Errorf("%s", tx.GetResult().GetMessage())
	}
	return tx, nil
}

// GetWitnessBrokerageInfo get witness brokerage info
func (c *GrpcClient) GetWitnessBrokerageInfo(witness string) (*tronpb.NumberMessage, error) {
	var err error
	var in = new(tronpb.BytesMessage)

	if in.Value, err = c.convert(witness); err != nil {
		return nil, err
	}

	ctx, cancelFunc := c.getContext()
	defer cancelFunc()

	return c.WalletClient.GetBrokerageInfo(ctx, in)
}

// UpdateWitnessBrokerage change SR comission fees
func (c *GrpcClient) UpdateWitnessBrokerage(from string, brokerage int32) (*tronpb.TransactionExtention, error) {
	var err error
	var in = new(tronpb.UpdateBrokerageContract)
	if in.OwnerAddress, err = c.convert(from); err != nil {
		return nil, err
	}
	in.Brokerage = brokerage

	ctx, cancelFunc := c.getContext()
	defer cancelFunc()

	tx, err := c.WalletClient.UpdateBrokerage(ctx, in)
	if err != nil {
		return nil, err
	}
	if proto.Size(tx) == 0 {
		return nil, fmt.Errorf("empty transaction")
	}
	if tx.GetResult().GetCode() != tronpb.Return_SUCCESS {
		return nil, fmt.Errorf("%s", tx.GetResult().GetMessage())
	}
	return tx, nil
}
