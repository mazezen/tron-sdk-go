package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func TestGrpcClient_GetAssetIssueByAccount(t *testing.T) {
	client := NewGrpcClient("grpc.trongrid.io:50051")
	dialOptions := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	err := client.Start(dialOptions...)
	assert.Equal(t, err, nil)

	issueByAccount, err := client.GetAssetIssueByAccount("TK9bNoWKRexGucyQ7CFq8wuR2nw5NJkFdk")
	assert.NoError(t, err, "should not return error")

	assert.NotNil(t, issueByAccount)

	t.Logf("issueByAccount %+v", issueByAccount)
}

func TestGrpcClient_GetAssetIssueByName(t *testing.T) {
	client := NewGrpcClient("grpc.trongrid.io:50051")
	dialOptions := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	err := client.Start(dialOptions...)
	assert.Equal(t, err, nil)

	res, err := client.GetAssetIssueByAccount("TK9bNoWKRexGucyQ7CFq8wuR2nw5NJkFdk")
	assert.NoError(t, err, "should not return error")
	assert.NotNil(t, res)
	t.Logf("res %+v", res)
}

func TestGrpcClient_GetAssetIssueListByName(t *testing.T) {
	client := NewGrpcClient("grpc.trongrid.io:50051")
	dialOptions := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	err := client.Start(dialOptions...)
	assert.Equal(t, err, nil)

	res, err := client.GetAssetIssueListByName("TRC10org")
	assert.NoError(t, err, "should not return error")
	assert.NotNil(t, res)
	for _, asset := range res.GetAssetIssue() {
		t.Logf("asset %+v", asset)
	}
}

func TestGrpcClient_GetAssetIssueById(t *testing.T) {
	client := NewGrpcClient("grpc.trongrid.io:50051")
	dialOptions := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	err := client.Start(dialOptions...)
	assert.Equal(t, err, nil)
	res, err := client.GetAssetIssueById("1004114")
	assert.NoError(t, err, "should not return error")
	assert.NotNil(t, res)
	t.Logf("res %+v", res)
}
