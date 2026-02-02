package client

import (
	"fmt"

	tronpb "github.com/mazezen/tron-sdk-go/pb/tron"
	"google.golang.org/protobuf/proto"
)

func (c *GrpcClient) MarketSellAsset(
	ownerAddress string,
	sellTokenId string,
	sellTokenQuantity int64,
	buyTokenId string,
	buyTokenQuantity int64, // min to receive
) (*tronpb.TransactionExtention, error) {
	var err error
	var in = &tronpb.MarketSellAssetContract{
		SellTokenId:       []byte(sellTokenId),
		SellTokenQuantity: sellTokenQuantity,
		BuyTokenId:        []byte(buyTokenId),
		BuyTokenQuantity:  buyTokenQuantity,
	}
	if in.OwnerAddress, err = c.convert(ownerAddress); err != nil {
		return nil, err
	}

	ctx, cancelFunc := c.getContext()
	defer cancelFunc()

	tx, err := c.WalletClient.MarketSellAsset(ctx, in)
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

func (c *GrpcClient) MarketCancelOrder(
	ownerAddress string,
	orderId string,
) (*tronpb.TransactionExtention, error) {
	var err error
	var in = &tronpb.MarketCancelOrderContract{
		OrderId: []byte(orderId),
	}
	if in.OwnerAddress, err = c.convert(ownerAddress); err != nil {
		return nil, err
	}
	ctx, cancelFunc := c.getContext()
	defer cancelFunc()
	tx, err := c.WalletClient.MarketCancelOrder(ctx, in)
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

func (c *GrpcClient) GetMarketOrderById(value string) (*tronpb.MarketOrder, error) {
	var in = &tronpb.BytesMessage{
		Value: []byte(value),
	}
	ctx, cancelFunc := c.getContext()
	defer cancelFunc()
	return c.WalletClient.GetMarketOrderById(ctx, in)
}

func (c *GrpcClient) GetMarketOrderByAccount(value string) (*tronpb.MarketOrderList, error) {
	var in = &tronpb.BytesMessage{
		Value: []byte(value),
	}
	ctx, cancelFunc := c.getContext()
	defer cancelFunc()
	return c.WalletClient.GetMarketOrderByAccount(ctx, in)
}

func (c *GrpcClient) GetMarketPriceByPair(
	sellTokenId string,
	buyTokenId string,
) (*tronpb.MarketPriceList, error) {
	var in = &tronpb.MarketOrderPair{
		SellTokenId: []byte(sellTokenId),
		BuyTokenId:  []byte(buyTokenId),
	}
	ctx, cancelFunc := c.getContext()
	defer cancelFunc()
	return c.WalletClient.GetMarketPriceByPair(ctx, in)
}

func (c *GrpcClient) GetMarketOrderListByPair(
	sellTokenId string,
	buyTokenId string,
) (*tronpb.MarketOrderList, error) {
	var in = &tronpb.MarketOrderPair{
		SellTokenId: []byte(sellTokenId),
		BuyTokenId:  []byte(buyTokenId),
	}
	ctx, cancelFunc := c.getContext()
	defer cancelFunc()
	return c.WalletClient.GetMarketOrderListByPair(ctx, in)
}

func (c *GrpcClient) GetMarketPairList() (*tronpb.MarketOrderPairList, error) {
	ctx, cancelFunc := c.getContext()
	defer cancelFunc()

	return c.WalletClient.GetMarketPairList(ctx, &tronpb.EmptyMessage{})
}
