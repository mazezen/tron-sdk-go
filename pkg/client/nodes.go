package client

import tronpb "github.com/mazezen/tron-sdk-go/pb/tron"

// ListNodes query the list of peers connected to the current node.
// https://developers.tron.network/reference/wallet-listnodes
func (c *GrpcClient) ListNodes() (*tronpb.NodeList, error) {
	ctx, cancelFunc := c.getContext()
	defer cancelFunc()
	return c.WalletClient.ListNodes(ctx, &tronpb.EmptyMessage{})
}

// GetNodeInfo query Node Information. Returns information about current state of node. (Confirmed state)
// https://developers.tron.network/reference/getnodeinfo-1
func (c *GrpcClient) GetNodeInfo() (*tronpb.NodeInfo, error) {
	ctx, cancelFunc := c.getContext()
	defer cancelFunc()

	return c.WalletClient.GetNodeInfo(ctx, &tronpb.EmptyMessage{})
}
