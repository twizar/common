package client_test

import (
	"net/http"
	"os"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/lambda"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/twizar/common/pkg/client"
	"github.com/twizar/common/test/mock"
)

const (
	liverpoolID = "d6548941-53f1-4d27-ad3d-0286cf512af1"
	milanID     = "5d912b4e-4932-496d-b706-c22b58f76a21"
	sevillaID   = "b0f6d915-da69-4681-bd7e-d933dd599ab2"
	bayernID    = "418ca28d-af10-4fbb-8b10-6afd74a001b7"
)

func TestAWSLambdaTeams_AllTeams(t *testing.T) {
	data, err := os.ReadFile("../../test/data/all_teams_payload.js")
	require.NoError(t, err)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	lambdaClient := mock.NewMockAWSLambdaClient(ctrl)
	lambdaClient.
		EXPECT().
		Invoke(gomock.Any()).
		Return(&lambda.InvokeOutput{Payload: data, StatusCode: aws.Int64(http.StatusOK)}, nil)

	c := client.NewAWSLambdaTeams(lambdaClient, "function")
	teams, err := c.AllTeams()
	require.NoError(t, err)
	assert.Equal(t, 703, len(teams))
}

func TestAWSLambdaTeams_TeamsByID(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	data, err := os.ReadFile("../../test/data/teams_by_ID_payload.js")
	require.NoError(t, err)
	lambdaClient := mock.NewMockAWSLambdaClient(ctrl)
	lambdaClient.
		EXPECT().
		Invoke(gomock.Any()).
		Return(&lambda.InvokeOutput{Payload: data, StatusCode: aws.Int64(http.StatusOK)}, nil)

	c := client.NewAWSLambdaTeams(lambdaClient, "function")
	stubIDs := []string{liverpoolID, milanID, sevillaID, bayernID}
	teams, err := c.TeamsByID(stubIDs)
	require.NoError(t, err)
	assert.Equal(t, 4, len(teams))
}

func TestAWSLambdaTeams_SearchTeams(t *testing.T) {
	data, err := os.ReadFile("../../test/data/search_teams_payload.js")
	require.NoError(t, err)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	lambdaClient := mock.NewMockAWSLambdaClient(ctrl)
	lambdaClient.
		EXPECT().
		Invoke(gomock.Any()).
		Return(&lambda.InvokeOutput{Payload: data, StatusCode: aws.Int64(http.StatusOK)}, nil)

	c := client.NewAWSLambdaTeams(lambdaClient, "function-name")
	teams, err := c.SearchTeams(4, []string{}, "rating", 0)
	require.NoError(t, err)
	assert.Equal(t, 82, len(teams))
}
