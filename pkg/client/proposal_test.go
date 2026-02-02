package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func TestGrpcClient_GetPaginatedProposalList(t *testing.T) {
	client := NewGrpcClient("grpc.trongrid.io:50051")
	dialOptions := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	err := client.Start(dialOptions...)
	assert.NoError(t, err, "failed to start grpc client")
	defer client.Stop()

	proposalList, err := client.GetPaginatedProposalList(1, 5)
	assert.NoError(t, err, "failed to get paginated proposal")
	for _, proposal := range proposalList.GetProposals() {
		t.Logf("proposal: %v", proposal)
	}
}

func TestGrpcClient_GetProposalById(t *testing.T) {
	client := NewGrpcClient("grpc.trongrid.io:50051")
	dialOptions := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	err := client.Start(dialOptions...)
	assert.NoError(t, err, "failed to start grpc client")
	defer client.Stop()

	proposal, err := client.GetProposalById(2)
	assert.NoError(t, err, "failed to get proposal")
	assert.NotNil(t, proposal, "proposal should not be nil")

	t.Logf("proposal: %+v", proposal)
}
