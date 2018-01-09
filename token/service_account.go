package token

import (
	"context"
	"github.com/fabric8-services/fabric8-notification/auth"
	"github.com/fabric8-services/fabric8-notification/auth/api"
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
	tokenString, err := auth.GetServiceAccountToken(ctx, c.client, c.accountID, c.accountSecret)
	if err != nil {
		return "", err
	}
	return *tokenString.AccessToken, nil
}
