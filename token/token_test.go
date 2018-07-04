package token_test

import (
	//"context"
	"os"
	"testing"

	"github.com/fabric8-services/fabric8-notification/configuration"
	"github.com/fabric8-services/fabric8-notification/token"
	//"github.com/fabric8-services/fabric8-notification/collector"
	//"github.com/goadesign/goa/uuid"
	"github.com/stretchr/testify/assert"
)

const (
	OpenShiftIOAuthAPI = "https://auth.openshift.io"
)

func TestManager(t *testing.T) {

	old := os.Getenv("F8_AUTH_URL")
	defer os.Setenv("F8_AUTH_URL", old)

	os.Setenv("F8_AUTH_URL", "https://auth.openshift.io")
	config, err := configuration.GetData()
	assert.Nil(t, err)

	manager, err := token.NewManager(config)
	assert.Nil(t, err)

	keys := manager.PublicKeys()
	assert.NotEmpty(t, keys)
	for _, k := range keys {
		assert.NotNil(t, k.N)
	}
}
