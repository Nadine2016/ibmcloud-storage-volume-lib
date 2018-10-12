/*******************************************************************************
 * IBM Confidential
 * OCO Source Materials
 * IBM Cloud Container Service, 5737-D43
 * (C) Copyright IBM Corp. 2018 All Rights Reserved.
 * The source code for this program is not  published or otherwise divested of
 * its trade secrets, irrespective of what has been deposited with
 * the U.S. Copyright Office.
 ******************************************************************************/

package auth

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/IBM/ibmcloud-storage-volume-lib/config"

	"github.com/IBM/ibmcloud-storage-volume-lib/lib/provider"
	"go.uber.org/zap"
)

var (
	logger *zap.Logger
)

func init() {
	logger, _ = zap.NewDevelopment()
}

/*func Test_ForDefaultAccount(t *testing.T) {
	logger := zap.NewNop()
	username := "user1"
	apiKey := "abcdefg"
	endpointURL := "http://myEndpointUrl"

	ccf := &contextCredentialsFactory{
		softlayerConfig: &config.SoftlayerConfig{
			SoftlayerUsername:    username,
			SoftlayerAPIKey:      apiKey,
			SoftlayerEndpointURL: endpointURL,
		},
	}

	contextCredentials, err := ccf.ForDefaultAccount(logger)

	assert.NoError(t, err)

	assert.Equal(t, provider.ContextCredentials{
		AuthType:       provider.IaaSAPIKey,
		DefaultAccount: true,
		UserID:         username,
		Credential:     apiKey,
	}, contextCredentials)
}*/

func Test_ForIaaSAPIKey(t *testing.T) {
	account := "account1"
	username := "user1"
	apiKey := "abcdefg"
	endpointURL := "http://myEndpointUrl"

	ccf := &contextCredentialsFactory{
		softlayerConfig: &config.SoftlayerConfig{
			SoftlayerEndpointURL: endpointURL,
		},
	}

	contextCredentials, err := ccf.ForIaaSAPIKey(account, username, apiKey, *logger)

	assert.NoError(t, err)

	assert.Equal(t, provider.ContextCredentials{
		AuthType:     provider.IaaSAPIKey,
		IAMAccountID: account,
		UserID:       username,
		Credential:   apiKey,
	}, contextCredentials)

}
