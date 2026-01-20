package client

import (
	"fmt"

	tronpb "github.com/mazezen/tron-sdk-go/pb/tron"
)

// GetNowBlock
// 请使用 GetNowBlock2 替代.
func (c *GrpcClient) GetNowBlock() (*tronpb.Block, error) {
	ctx, cancelFunc := c.getContext()
	defer cancelFunc()

	block, err := c.WalletClient.GetNowBlock(ctx, new(tronpb.EmptyMessage))
	if err != nil {
		return nil, fmt.Errorf("GetNowBlock: %w", err)
	}
	return block, nil
}

// GetNowBlock2 获取 TIP 区块
func (c *GrpcClient) GetNowBlock2() (*tronpb.BlockExtention, error) {
	ctx, cancelFunc := c.getContext()
	defer cancelFunc()

	block, err := c.WalletClient.GetNowBlock2(ctx, new(tronpb.EmptyMessage))
	if err != nil {
		return nil, fmt.Errorf("get now block: %w", err)
	}
	return block, err
}
