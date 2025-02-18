/*
Genome Nexus API

Testing PdbControllerApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package genome_nexus_internal_api

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func Test_genome_nexus_internal_api_PdbControllerApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test PdbControllerApiService FetchPdbHeaderGET", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pdbId string

		resp, httpRes, err := apiClient.PdbControllerApi.FetchPdbHeaderGET(context.Background(), pdbId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PdbControllerApiService FetchPdbHeaderPOST", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.PdbControllerApi.FetchPdbHeaderPOST(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
