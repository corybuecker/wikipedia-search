package main

import (
	"io/ioutil"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

var testPayload []byte
var exactMatchSearch = "Mount & Blade"
var nonExactMatchSearch = "Mount"

func TestRunner(t *testing.T) {
	testPayload, _ = ioutil.ReadFile("./test_json/mount_blade.json")

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", generateURL(exactMatchSearch), httpmock.NewBytesResponder(200, testPayload))
	httpmock.RegisterResponder("GET", generateURL(nonExactMatchSearch), httpmock.NewBytesResponder(200, testPayload))

	t.Run("results without exact match", resultsWithoutExactMatch)
	t.Run("results with exact match", resultsWithExactMatch)
}

func resultsWithoutExactMatch(t *testing.T) {
	results, _ := Search(exactMatchSearch, false)
	assert.Equal(t, results[0].ID, 2008127)
}

func resultsWithExactMatch(t *testing.T) {
	results, _ := Search(nonExactMatchSearch, true)
	assert.Empty(t, results)
}
