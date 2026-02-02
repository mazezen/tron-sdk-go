package client

import (
	"encoding/binary"
	"fmt"

	tronpb "github.com/mazezen/tron-sdk-go/pb/tron"
	"google.golang.org/protobuf/proto"
)

// ProposalCreate creates a proposal transaction
// https://developers.tron.network/reference/proposalcreate
func (c *GrpcClient) ProposalCreate(from string, parameters map[int64]int64) (*tronpb.TransactionExtention, error) {
	var err error
	var in = new(tronpb.ProposalCreateContract)
	if in.OwnerAddress, err = c.convert(from); err != nil {
		return nil, err
	}

	in.Parameters = parameters

	ctx, cancelFunc := c.getContext()
	defer cancelFunc()

	tx, err := c.WalletClient.ProposalCreate(ctx, in)
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

// ProposalApprove approves proposed transaction
// https://developers.tron.network/reference/proposalapprove
func (c *GrpcClient) ProposalApprove(from string, proposalId int64, isAddApproval bool) (*tronpb.TransactionExtention, error) {
	var err error
	var in = new(tronpb.ProposalApproveContract)
	if in.OwnerAddress, err = c.convert(from); err != nil {
		return nil, err
	}
	in.ProposalId = proposalId
	in.IsAddApproval = isAddApproval

	ctx, cancelFunc := c.getContext()
	defer cancelFunc()
	tx, err := c.WalletClient.ProposalApprove(ctx, in)
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

// ProposalDelete delete proposal transaction
// https://developers.tron.network/reference/proposaldelete
func (c *GrpcClient) ProposalDelete(from string, proposalId int64) (*tronpb.TransactionExtention, error) {
	var err error
	var in = new(tronpb.ProposalDeleteContract)
	if in.OwnerAddress, err = c.convert(from); err != nil {
		return nil, err
	}
	in.ProposalId = proposalId

	ctx, cancelFunc := c.getContext()
	defer cancelFunc()

	tx, err := c.WalletClient.ProposalDelete(ctx, in)
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

// ListProposals list app proposals
func (c *GrpcClient) ListProposals() (*tronpb.ProposalList, error) {
	var ctx, cancelFunc = c.getContext()
	defer cancelFunc()

	return c.WalletClient.ListProposals(ctx, &tronpb.EmptyMessage{})
}

func (c *GrpcClient) GetPaginatedProposalList(offset, limit int64) (*tronpb.ProposalList, error) {
	var ctx, cancelFunc = c.getContext()
	defer cancelFunc()
	return c.WalletClient.GetPaginatedProposalList(ctx, &tronpb.PaginatedMessage{
		Offset: offset,
		Limit:  limit,
	})
}

// GetProposalById queries proposal based on ID and returns proposal details
// https://developers.tron.network/reference/getproposalbyid
func (c *GrpcClient) GetProposalById(id int64) (*tronpb.Proposal, error) {
	var ctx, cancelFunc = c.getContext()
	defer cancelFunc()
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(id))
	return c.WalletClient.GetProposalById(ctx, &tronpb.BytesMessage{Value: b})
}
