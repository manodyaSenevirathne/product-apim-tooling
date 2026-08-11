/*
*  Copyright (c) WSO2 Inc. (http://www.wso2.org) All Rights Reserved.
*
*  WSO2 Inc. licenses this file to you under the Apache License,
*  Version 2.0 (the "License"); you may not use this file except
*  in compliance with the License.
*  You may obtain a copy of the License at
*
*    http://www.apache.org/licenses/LICENSE-2.0
*
* Unless required by applicable law or agreed to in writing,
* software distributed under the License is distributed on an
* "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
* KIND, either express or implied.  See the License for the
* specific language governing permissions and limitations
* under the License.
 */

package impl

import (
	"archive/zip"
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"

	v2 "github.com/wso2/product-apim-tooling/import-export-cli/specs/v2"
	"github.com/wso2/product-apim-tooling/import-export-cli/utils"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestExtractAPIInfoWithCorrectJSON(t *testing.T) {
	// Correct json
	content := `{
		"type": "api",
		"version": "v4.6.0",
		"data": {
		  "id": "e4d0c1be-44e9-43ad-b434-f8e2f02dad11",
		  "name": "APIName",
		  "provider": "devops",
		  "version": "1.0.0"
		}
	  }`

	api, err := extractAPIDefinition([]byte(content))
	assert.Equal(t, err, nil, "Should return nil error for correct json")
	assert.Equal(t, api.Data.Name, "APIName", "Should parse correct json")
	assert.Equal(t, api.Data.Provider, "devops", "Should parse correct json")
	assert.Equal(t, api.Data.Version, "1.0.0", "Should parse correct json")
}

func TestExtractAPIInfoWhenDataTagMissing(t *testing.T) {
	// When ID tag missing
	content := `{
		"type": "api",
		"version": "v4.6.0"
	  }`

	api, err := extractAPIDefinition([]byte(content))
	assert.Nil(t, err, "Should return nil error")
	assert.Equal(t, v2.APIDTODefinition{}, api.Data, "Should return empty Data when ID tag missing")
}

func TestExtractAPIInfoWithMalformedJSON(t *testing.T) {
	// Malformed json
	content := `{
	  "uuid": "e4d0c1be-44e9-43ad-b434-f8e2f02dad11",
	  "description": "Some API Description",
	  "type": "HTTP",
	  "context": "/api/1.0.0",
	  "contextTemplate": "/api/{version}",
	  "tags": [
		"api"
	  
	}`

	api, err := extractAPIDefinition([]byte(content))
	assert.Nil(t, api, "Should return nil API struct")
	assert.Error(t, err, "Should return an error regarding malformed json")
}

func TestGetAPIInfoCorrectDirectoryStructure(t *testing.T) {
	api, _, err := GetAPIDefinition(utils.GetRelativeTestDataPathFromImpl() + "PizzaShackAPI-1.0.0")
	assert.Nil(t, err, "Should return nil error on reading correct directories")
	assert.Equal(t, api.Data.Name, "PizzaShackAPI", "Should return correct values for API name")
	assert.Equal(t, api.Data.Provider, "admin", "Should return correct values for API provider")
	assert.Equal(t, api.Data.Version, "1.0.0", "Should return correct values for API version")
}

func TestGetAPIInfoMalformedDirectory(t *testing.T) {
	api, _, err := GetAPIDefinition("testdata/PizzaShackAPI_1.0.0-malformed")
	assert.Error(t, err, "Should return error on reading malformed directories")
	assert.True(t, os.IsNotExist(err), "File not found error must be thrown")
	assert.Nil(t, api,
		"Should return nil for malformed directories")
}

// TestImportAPI_DryRunWithParams_DoesNotRestructureProject verifies that ImportAPI
// skips params-based project restructuring when --dry-run is set. The uploaded zip
// is inspected via a mock server in both cases.
func TestImportAPI_DryRunWithParams_DoesNotRestructureProject(t *testing.T) {
	srcDir := utils.GetRelativeTestDataPathFromImpl() + "PizzaShackAPI-1.0.0"

	paramsFile := filepath.Join(t.TempDir(), "api_params.yaml")
	require.NoError(t, os.WriteFile(paramsFile, []byte(`environments:
  - name: testenv
    configs:
      endpoints:
        production:
          url: 'https://example.com'`), 0644))

	// getUploadedZipEntries calls ImportAPI against a mock server and returns
	// the names of entries in the zip that was uploaded. Entries carry a directory
	// prefix (e.g. "PizzaShackAPI-1.0.0/api.yaml"), so callers should match by suffix.
	getUploadedZipEntries := func(dryRun bool) []string {
		var uploadedZip []byte

		srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			_ = r.ParseMultipartForm(32 << 20)
			if f, _, err := r.FormFile("file"); err == nil {
				defer f.Close()
				uploadedZip, _ = io.ReadAll(f)
			}
			if dryRun {
				w.WriteHeader(http.StatusOK)
				_, _ = w.Write([]byte(`{"compliance-check":{"result":"pass","violations":[]}}`))
			} else {
				w.WriteHeader(http.StatusCreated)
			}
		}))
		defer srv.Close()

		_ = ImportAPI("token", srv.URL, "testenv", srcDir, paramsFile,
			false, true, false, false, false, dryRun, "table")

		zr, err := zip.NewReader(bytes.NewReader(uploadedZip), int64(len(uploadedZip)))
		if err != nil {
			return nil
		}
		entries := make([]string, 0, len(zr.File))
		for _, f := range zr.File {
			entries = append(entries, f.Name)
		}
		return entries
	}

	hasEntry := func(entries []string, suffix string) bool {
		for _, e := range entries {
			if strings.HasSuffix(e, suffix) {
				return true
			}
		}
		return false
	}

	t.Run("without dry-run params restructures the uploaded artifact", func(t *testing.T) {
		entries := getUploadedZipEntries(false)
		assert.True(t, hasEntry(entries, "SourceArchive.zip"),
			"uploaded zip must contain SourceArchive.zip, got: %v", entries)
		assert.False(t, hasEntry(entries, "api.yaml"),
			"api.yaml must not be at root of the uploaded zip, got: %v", entries)
	})

	t.Run("with dry-run params restructuring is skipped", func(t *testing.T) {
		entries := getUploadedZipEntries(true)
		assert.True(t, hasEntry(entries, "api.yaml"),
			"api.yaml must be at root of the uploaded zip, got: %v", entries)
		assert.False(t, hasEntry(entries, "SourceArchive.zip"),
			"uploaded zip must not contain SourceArchive.zip, got: %v", entries)
	})
}
