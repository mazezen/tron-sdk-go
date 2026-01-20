package client

import (
	"fmt"

	tronpb "github.com/mazezen/tron-sdk-go/pb/tron"
	"github.com/mazezen/tron-sdk-go/pkg/common"
)

// GetNowBlock return TIP block
// Please use GetNowBlock2 instead of this function.
func (c *GrpcClient) GetNowBlock() (*tronpb.Block, error) {
	ctx, cancelFunc := c.getContext()
	defer cancelFunc()

	block, err := c.WalletClient.GetNowBlock(ctx, new(tronpb.EmptyMessage))
	if err != nil {
		return nil, fmt.Errorf("GetNowBlock: %w", err)
	}
	return block, nil
}

// GetNowBlock2 return TIP block
func (c *GrpcClient) GetNowBlock2() (*tronpb.BlockExtention, error) {
	ctx, cancelFunc := c.getContext()
	defer cancelFunc()

	block, err := c.WalletClient.GetNowBlock2(ctx, new(tronpb.EmptyMessage))
	if err != nil {
		return nil, fmt.Errorf("GetNowBlock2: %w", err)
	}
	return block, err
}

// GetBlockByNum
// Queries whether a specified block is confirmed
// If the block can be obtained, the block is confirmed by the network;
// if the block cannot be obtained, the block is not confirmed by the network
// https://developers.tron.network/reference/getblockbynum
// Please use GetBlockByNum2 instead of this function.
func (c *GrpcClient) GetBlockByNum(num int64) (*tronpb.Block, error) {
	ctx, cancelFunc := c.getContext()
	defer cancelFunc()

	block, err := c.WalletClient.GetBlockByNum(ctx, &tronpb.NumberMessage{Num: num})
	if err != nil {
		return nil, fmt.Errorf("GetBlockByNum: %w", err)
	}
	return block, nil
}

// GetBlockByNum2
// Queries whether a specified block is confirmed
// If the block can be obtained, the block is confirmed by the network;
// if the block cannot be obtained, the block is not confirmed by the network
// https://developers.tron.network/reference/getblockbynum
// Use this function instead of GetBlockByNum.
func (c *GrpcClient) GetBlockByNum2(num int64) (*tronpb.BlockExtention, error) {
	ctx, cancelFunc := c.getContext()
	defer cancelFunc()

	block, err := c.WalletClient.GetBlockByNum2(ctx, &tronpb.NumberMessage{Num: num})
	if err != nil {
		return nil, fmt.Errorf("GetBlockByNum2: %w", err)
	}
	return block, nil
}

// GetTransactionCountByBlockNum get transaction count through given block number
func (c *GrpcClient) GetTransactionCountByBlockNum(num int64) (int64, error) {
	ctx, cancelFunc := c.getContext()
	defer cancelFunc()

	numberMessage, err := c.WalletClient.GetTransactionCountByBlockNum(ctx, &tronpb.NumberMessage{Num: num})
	if err != nil {
		return 0, fmt.Errorf("GetTransactionCountByBlockNum: %w", err)
	}
	return numberMessage.Num, nil
}

// GetBlockById Query block by ID(block hash). (Confirmed state)
// https://developers.tron.network/reference/getblockbyid-1
func (c *GrpcClient) GetBlockById(id string) (*tronpb.Block, error) {
	ctx, cancelFunc := c.getContext()
	defer cancelFunc()

	hexId, err := common.FromHex(id)
	if err != nil {
		return nil, fmt.Errorf("FromHex: %w", err)
	}

	block, err := c.WalletClient.GetBlockById(ctx, &tronpb.BytesMessage{Value: hexId})
	if err != nil {
		return nil, fmt.Errorf("GetBlockById: %w", err)
	}
	return block, nil
}

// GetBlockByLimitNext
// Returns the list of Block Objects included in the 'Block Height' range specified. (Confirmed state)
// https://developers.tron.network/reference/getblockbylimitnext-1
// Please use GetBlockByLimitNext2 instead of this function.
func (c *GrpcClient) GetBlockByLimitNext(startNum, endNum int64) (*tronpb.BlockList, error) {
	ctx, cancelFunc := c.getContext()
	defer cancelFunc()

	block, err := c.WalletClient.GetBlockByLimitNext(ctx, &tronpb.BlockLimit{StartNum: startNum, EndNum: endNum})
	if err != nil {
		return nil, fmt.Errorf("GetBlockByLimitNext: %w", err)
	}
	return block, nil
}

// GetBlockByLimitNext2
// Returns the list of Block Objects included in the 'Block Height' range specified. (Confirmed state)
// https://developers.tron.network/reference/getblockbylimitnext-1
// Use this function instead of GetBlockByLimitNext.
func (c *GrpcClient) GetBlockByLimitNext2(startNum, endNum int64) (*tronpb.BlockListExtention, error) {
	ctx, cancelFunc := c.getContext()
	defer cancelFunc()

	block, err := c.WalletClient.GetBlockByLimitNext2(ctx, &tronpb.BlockLimit{StartNum: startNum, EndNum: endNum})
	if err != nil {
		return nil, fmt.Errorf("GetBlockByLimitNext: %w", err)
	}
	return block, nil
}

// GetBlockByLatestNum block list
// Return the most recent number of blocks
// https://developers.tron.network/reference/getblockbylatestnum
// Please use GetBlockByLatestNum2 instead of this function.
func (c *GrpcClient) GetBlockByLatestNum(num int64) (*tronpb.BlockList, error) {
	ctx, cancelFunc := c.getContext()
	defer cancelFunc()

	blockList, err := c.WalletClient.GetBlockByLatestNum(ctx, &tronpb.NumberMessage{Num: num})
	if err != nil {
		return nil, fmt.Errorf("GetBlockByLatestNum: %w", err)
	}

	return blockList, nil
}

// GetBlockByLatestNum2
// Return the most recent number of blocks
// https://developers.tron.network/reference/getblockbylatestnum
// Use this function instead of GetBlockByLatestNum.
func (c *GrpcClient) GetBlockByLatestNum2(num int64) (*tronpb.BlockListExtention, error) {
	ctx, cancelFunc := c.getContext()
	defer cancelFunc()

	blockList, err := c.WalletClient.GetBlockByLatestNum2(ctx, &tronpb.NumberMessage{Num: num})
	if err != nil {
		return nil, fmt.Errorf("GetBlockByLatestNum2: %w", err)
	}
	return blockList, nil
}
