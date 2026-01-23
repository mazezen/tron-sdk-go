package client

import (
	"fmt"

	"github.com/golang/protobuf/proto"
	tronpb "github.com/mazezen/tron-sdk-go/pb/tron"
)

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
