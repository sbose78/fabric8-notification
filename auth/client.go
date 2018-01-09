package auth

import (
	"context"
	"net/http"
	"net/url"

	"fmt"

	"github.com/fabric8-services/fabric8-notification/auth/api"
	"github.com/fabric8-services/fabric8-wit/goasupport"
	goaclient "github.com/goadesign/goa/client"
	"github.com/goadesign/goa/uuid"
	"github.com/gregjones/httpcache"
)

func NewCachedClient(hostURL string) (*api.Client, error) {

	u, err := url.Parse(hostURL)
	if err != nil {
		return nil, err
	}

	tp := httpcache.NewMemoryCacheTransport()
	client := http.Client{Transport: tp}

	c := api.New(goaclient.HTTPClientDoer(&client))
	c.Host = u.Host
	c.Scheme = u.Scheme
	return c, nil
}

func GetUser(ctx context.Context, client *api.Client, uID uuid.UUID) (*api.User, error) {
	resp, err := client.ShowUsers(goasupport.ForwardContextRequestID(ctx), api.ShowUsersPath(uID.String()), nil, nil)
	if resp != nil {
		defer resp.Body.Close()
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("non %v status code for %v, returned %v", http.StatusOK, "GET user", resp.StatusCode)
	}

	if err != nil {
		return nil, err
	}
	return client.DecodeUser(resp)
}

func GetSpaceCollaborators(ctx context.Context, client *api.Client, spaceID uuid.UUID) (*api.UserList, error) {
	pageLimit := 100
	pageOffset := "0"
	resp, err := client.ListCollaborators(goasupport.ForwardContextRequestID(ctx), api.ListCollaboratorsPath(spaceID), &pageLimit, &pageOffset, nil, nil)
	if resp != nil {
		defer resp.Body.Close()
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("non %v status code for %v, returned %v", http.StatusOK, "GET collaborators", resp.StatusCode)
	}

	if err != nil {
		return nil, err
	}
	return client.DecodeUserList(resp)
}

func GetServiceAccountToken(ctx context.Context, client *api.Client, serviceAccountID string, serviceAccountSecret string) (*api.OauthToken, error) {
	payload := api.TokenExchange{
		ClientID:     serviceAccountID,
		ClientSecret: &serviceAccountSecret,
	}
	resp, err := client.ExchangeToken(goasupport.ForwardContextRequestID(ctx), api.ExchangeTokenPath(), &payload, "application/x-www-form-urlencoded")

	if err != nil {
		return nil, err
	}

	return client.DecodeOauthToken(resp)
}
