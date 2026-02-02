package client

import (
	"fmt"

	tronpb "github.com/mazezen/tron-sdk-go/pb/tron"
	"google.golang.org/protobuf/proto"
)

// BuyStorage  trx quantity for buy storage (sun)
func (c *GrpcClient) BuyStorage(from string, quant int64) (*tronpb.TransactionExtention, error) {
	var err error
	var in = new(tronpb.BuyStorageContract)

	if in.OwnerAddress, err = c.convert(from); err != nil {
		return nil, err
	}
	in.Quant = quant

	ctx, cancelFunc := c.getContext()
	defer cancelFunc()

	tx, err := c.WalletClient.BuyStorage(ctx, in)
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

// BuyStorageBytes storage bytes for buy

func (c *GrpcClient) BuyStorageBytes(from string, bytes int64) (*tronpb.TransactionExtention, error) {
	var err error
	var in = new(tronpb.BuyStorageBytesContract)

	if in.OwnerAddress, err = c.convert(from); err != nil {
		return nil, err
	}
	in.Bytes = bytes

	ctx, cancelFunc := c.getContext()
	defer cancelFunc()

	tx, err := c.WalletClient.BuyStorageBytes(ctx, in)
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

// SellStorage sell storage
func (c *GrpcClient) SellStorage(from string, storageBytes int64) (*tronpb.TransactionExtention, error) {
	var err error
	var in = new(tronpb.SellStorageContract)

	if in.OwnerAddress, err = c.convert(from); err != nil {
		return nil, err
	}
	in.StorageBytes = storageBytes

	ctx, cancelFunc := c.getContext()
	defer cancelFunc()

	tx, err := c.WalletClient.SellStorage(ctx, in)
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
