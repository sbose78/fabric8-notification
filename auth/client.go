package auth

import (
	"context"
	"net/http"
	"net/url"

	"fmt"

	"github.com/fabric8-services/fabric8-notification/auth/api"
	"github.com/fabric8-services/fabric8-wit/goasupport"
	"github.com/fabric8-services/fabric8-wit/log"
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
