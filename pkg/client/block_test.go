package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func TestGrpcClient_GetNowBlock(t *testing.T) {
	client := NewGrpcClient("grpc.trongrid.io:50051")
	dialOptions := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	err := client.Start(dialOptions...)
	assert.NoError(t, err)
	defer client.Stop()

	block, err := client.GetNowBlock()
	assert.Equal(t, err, nil)
	assert.NotEqual(t, block, nil)
	t.Logf("block: %v\n", block)
}

func TestGrpcClient_GetNowBlock2(t *testing.T) {
	client := NewGrpcClient("grpc.trongrid.io:50051")
	dialOptions := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	err := client.Start(dialOptions...)
	assert.NoError(t, err)
	defer client.Stop()

	block, err := client.GetNowBlock2()
	assert.Equal(t, err, nil)
	assert.NotEqual(t, block, nil)
	t.Logf("block: %v\n", block)
}

func TestGrpcClient_GetBlockByNum(t *testing.T) {
	client := NewGrpcClient("grpc.trongrid.io:50051")
	dialOptions := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	err := client.Start(dialOptions...)
	assert.NoError(t, err)
	defer client.Stop()

	block, err := client.GetBlockByNum(79408785)
	assert.Equal(t, err, nil)
	assert.NotEqual(t, block, nil)
	t.Logf("79408785 block: %v\n", block)
}

func TestGrpcClient_GetBlockByNum2(t *testing.T) {
	client := NewGrpcClient("grpc.trongrid.io:50051")
	dialOptions := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	err := client.Start(dialOptions...)
	assert.NoError(t, err)
	defer client.Stop()

	block, err := client.GetBlockByNum2(79408785)
	assert.Equal(t, err, nil)
	assert.NotEqual(t, block, nil)
	t.Logf("block: %v\n", block)
}

func TestGrpcClient_GetTransactionCountByBlockNum(t *testing.T) {
	client := NewGrpcClient("grpc.trongrid.io:50051")
	dialOptions := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	err := client.Start(dialOptions...)
	assert.NoError(t, err)
	defer client.Stop()

	count, err := client.GetTransactionCountByBlockNum(79408785)
	assert.Equal(t, err, nil)
	assert.NotEqual(t, count, 0)
	t.Logf("79408785 block transaction count: %d\n", count)
}

func TestGrpcClient_GetBlockById(t *testing.T) {
	client := NewGrpcClient("grpc.trongrid.io:50051")
	dialOptions := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	err := client.Start(dialOptions...)
	assert.NoError(t, err)
	defer client.Stop()

	block, err := client.GetBlockById("0000000004bbae918ed02f63313398281299978d707ed2e0880942322ed06a13")
	assert.NoError(t, err)
	assert.NotEqual(t, block, nil)
	t.Logf("block: %v\n", block)

	block2, err2 := client.GetBlockById("0x0000000004bbae918ed02f63313398281299978d707ed2e0880942322ed06a13")
	assert.NoError(t, err2)
	assert.NotEqual(t, block2, nil)
	t.Logf("block2: %v\n", block2)

	block3, err3 := client.GetBlockById("0X0000000004bbae918ed02f63313398281299978d707ed2e0880942322ed06a13")
	assert.NoError(t, err3)
	assert.NotEqual(t, block3, nil)
	t.Logf("block3: %v\n", block3)
}

func TestGrpcClient_GetBlockByLimitNext(t *testing.T) {
	client := NewGrpcClient("grpc.trongrid.io:50051")
	dialOptions := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	err := client.Start(dialOptions...)
	assert.NoError(t, err)
	defer client.Stop()

	blockList, err := client.GetBlockByLimitNext(79408784, 79408785)
	assert.NoError(t, err)
	assert.NotEqual(t, blockList, nil)

	t.Logf("blockList: %v\n", blockList)
}

func TestGrpcClient_GetBlockByLimitNext2(t *testing.T) {
	client := NewGrpcClient("grpc.trongrid.io:50051")
	dialOptions := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	err := client.Start(dialOptions...)
	assert.NoError(t, err)
	defer client.Stop()

	blockList, err := client.GetBlockByLimitNext2(79408784, 79408785)
	assert.NoError(t, err)
	assert.NotEqual(t, blockList, nil)

	t.Logf("blockList: %v\n", blockList)
}

func TestGrpcClient_GetBlockByLatestNum(t *testing.T) {
	client := NewGrpcClient("grpc.trongrid.io:50051")
	dialOptions := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	err := client.Start(dialOptions...)
	assert.NoError(t, err)
	defer client.Stop()

	blockList, err := client.GetBlockByLatestNum(2)
	assert.NoError(t, err)
	assert.NotEqual(t, blockList, nil)

	t.Logf("blockList: %v\n", blockList)
}

func TestGrpcClient_GetBlockByLatestNum2(t *testing.T) {
	client := NewGrpcClient("grpc.trongrid.io:50051")
	dialOptions := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	err := client.Start(dialOptions...)
	assert.NoError(t, err)
	defer client.Stop()

	blockList, err := client.GetBlockByLatestNum2(2)
	assert.NoError(t, err)
	assert.NotEqual(t, blockList, nil)

	t.Logf("blockList: %v\n", blockList)
}
