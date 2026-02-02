package client

import (
	"encoding/binary"
	"fmt"

	tronpb "github.com/mazezen/tron-sdk-go/pb/tron"
	"google.golang.org/protobuf/proto"
)

// ExchangeCreate create a trading pair
// https://developers.tron.network/reference/exchangecreate
func (c *GrpcClient) ExchangeCreate(
	from, firstTokenId string,
	firstTokenBalance int64,
	secondTokenId string,
	secondTokenBalance int64,
) (*tronpb.TransactionExtention, error) {
	var err error
	var in = new(tronpb.ExchangeCreateContract)
	if in.OwnerAddress, err = c.convert(from); err != nil {
		return nil, err
	}
	in.FirstTokenId = []byte(firstTokenId)
	in.SecondTokenId = []byte(secondTokenId)
	in.FirstTokenBalance = firstTokenBalance
	in.SecondTokenBalance = secondTokenBalance

	ctx, cancelFunc := c.getContext()
	defer cancelFunc()

	tx, err := c.WalletClient.ExchangeCreate(ctx, in)
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

// ExchangeInject
// https://developers.tron.network/reference/exchangeinject
func (c *GrpcClient) ExchangeInject(
	from string,
	exchangeId int64,
	tokenId string,
	quant int64,
) (*tronpb.TransactionExtention, error) {
	var err error
	var in = &tronpb.ExchangeInjectContract{
		ExchangeId: exchangeId,
		TokenId:    []byte(tokenId),
		Quant:      quant,
	}
	if in.OwnerAddress, err = c.convert(from); err != nil {
		return nil, err
	}

	ctx, cancelFunc := c.getContext()
	defer cancelFunc()

	tx, err := c.WalletClient.ExchangeInject(ctx, in)
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

// ExchangeWithdraw withdraw the transaction pair.
// https://developers.tron.network/reference/exchangewithdraw
func (c *GrpcClient) ExchangeWithdraw(
	from string,
	exchangeId int64,
	tokenId string,
	quant int64,
) (*tronpb.TransactionExtention, error) {
	var err error
	var in = &tronpb.ExchangeWithdrawContract{
		ExchangeId: exchangeId,
		TokenId:    []byte(tokenId),
		Quant:      quant,
	}
	if in.OwnerAddress, err = c.convert(from); err != nil {
		return nil, err
	}

	ctx, cancelFunc := c.getContext()
	defer cancelFunc()

	tx, err := c.WalletClient.ExchangeWithdraw(ctx, in)
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

// ExchangeTransaction exchange transaction
func (c *GrpcClient) ExchangeTransaction(
	from string,
	exchangeID int64,
	tokenID string,
	amountToken int64,
	amountExpected int64,
) (*tronpb.TransactionExtention, error) {
	var err error
	var in = &tronpb.ExchangeTransactionContract{
		ExchangeId: exchangeID,
		TokenId:    []byte(tokenID),
		Quant:      amountToken,
		Expected:   amountExpected,
	}
	if in.OwnerAddress, err = c.convert(from); err != nil {
		return nil, err
	}

	ctx, cancelFunc := c.getContext()
	defer cancelFunc()

	tx, err := c.WalletClient.ExchangeTransaction(ctx, in)
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

// ListExchanges list all exchange pairs (Confirmed state)
// https://developers.tron.network/reference/listexchanges-1
func (c *GrpcClient) ListExchanges() (*tronpb.ExchangeList, error) {
	ctx, cancelFunc := c.getContext()
	defer cancelFunc()

	return c.WalletClient.ListExchanges(ctx, &tronpb.EmptyMessage{})
}

// GetPaginatedExchangeList return pagination data
func (c *GrpcClient) GetPaginatedExchangeList(offset, limit int64) (*tronpb.ExchangeList, error) {
	ctx, cancelFunc := c.getContext()
	defer cancelFunc()

	return c.WalletClient.GetPaginatedExchangeList(ctx, &tronpb.PaginatedMessage{
		Offset: offset,
		Limit:  limit,
	})
}

// GetExchangeById query exchange pair based on id (Confirmed state)
// https://developers.tron.network/reference/getexchangebyid-1
func (c *GrpcClient) GetExchangeById(id int64) (*tronpb.Exchange, error) {
	ctx, cancelFunc := c.getContext()
	defer cancelFunc()
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(id))
	return c.WalletClient.GetExchangeById(ctx, &tronpb.BytesMessage{Value: b})
}
