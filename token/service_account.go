package token

import (
	"context"
	"github.com/fabric8-services/fabric8-notification/auth/api"
	"github.com/fabric8-services/fabric8-wit/goasupport"
)

type Fabric8ServiceAccountTokenClient struct {
	client        *api.Client
	accountID     string
	accountSecret string
}

func NewFabric8ServiceAccountTokenClient(client *api.Client, accountID string, accountSecret string) *Fabric8ServiceAccountTokenClient {
	return &Fabric8ServiceAccountTokenClient{
		client:        client,
		accountID:     accountID,
		accountSecret: accountSecret,
	}
}

type Fabric8ServiceAccountTokenService interface {
	Get(ctx context.Context) (string, error)
}

func (c *Fabric8ServiceAccountTokenClient) Get(ctx context.Context) (string, error) {
	tokenString, err := getServiceAccountToken(ctx, c.client, c.accountID, c.accountSecret)
	if err != nil {
		return "", err
	}
	return *tokenString.AccessToken, nil
}

func getServiceAccountToken(ctx context.Context, client *api.Client, serviceAccountID string, serviceAccountSecret string) (*api.OauthToken, error) {
	payload := api.TokenExchange{
		ClientID:     serviceAccountID,
		ClientSecret: &serviceAccountSecret,
		GrantType:    "client_credentials",
	}
	resp, err := client.ExchangeToken(goasupport.ForwardContextRequestID(ctx), api.ExchangeTokenPath(), &payload, "application/x-www-form-urlencoded")

	if err != nil {
		return nil, err
	}

	return client.DecodeOauthToken(resp)
}
