package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/lambda"
	"github.com/twizar/common/pkg/dto"
)

type Teams interface {
	AllTeams() ([]dto.Team, error)
	TeamsByID(ids []string) ([]dto.Team, error)
	SearchTeams(minRating float64, leagues []string, orderBy string, limit int) ([]dto.Team, error)
}

type lambdaRequest struct {
	Path       string `json:"path"`
	HTTPMethod string `json:"httpMethod"`
}

type lambdaResponse struct {
	StatusCode int               `json:"statusCode"`
	Headers    map[string]string `json:"headers"`
	Body       string            `json:"body"`
}

type AWSLambdaTeams struct {
	client          AWSLambdaClient
	lambdaTeamsName string
}

func NewAWSLambdaTeams(client AWSLambdaClient, lambdaTeamsName string) *AWSLambdaTeams {
	return &AWSLambdaTeams{client: client, lambdaTeamsName: lambdaTeamsName}
}

func (alt AWSLambdaTeams) AllTeams() ([]dto.Team, error) {
	return alt.call("/teams", http.MethodGet)
}

func (alt AWSLambdaTeams) TeamsByID(ids []string) ([]dto.Team, error) {
	idsQuery := strings.Join(ids, ",")
	path := fmt.Sprintf("/teams?ids=%s", idsQuery)

	return alt.call(path, http.MethodGet)
}

func (alt AWSLambdaTeams) SearchTeams(minRating float64, leagues []string, orderBy string, limit int) ([]dto.Team, error) {
	path := fmt.Sprintf("/teams/search?min-rating=%f&leagues=%s&order-by=%s&limit=%d", minRating, strings.Join(leagues, ","), orderBy, limit)

	return alt.call(path, http.MethodGet)
}

func (alt AWSLambdaTeams) call(path, method string) ([]dto.Team, error) {
	request := lambdaRequest{path, method}

	payload, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("request marshall error: %w", err)
	}

	result, err := alt.client.Invoke(&lambda.InvokeInput{
		FunctionName: aws.String(alt.lambdaTeamsName),
		Payload:      payload,
	})
	if err != nil {
		return nil, fmt.Errorf("lambda invoke error: %w", err)
	}

	resp := lambdaResponse{}

	if err = json.Unmarshal(result.Payload, &resp); err != nil {
		return nil, fmt.Errorf("response unmarshall error: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error getting items: %w", err)
	}

	var teams []dto.Team

	if err = json.Unmarshal([]byte(resp.Body), &teams); err != nil {
		return nil, fmt.Errorf("response unmarshall error: %w", err)
	}

	return teams, nil
}
