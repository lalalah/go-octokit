package octokit

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserSearchService_All(t *testing.T) {
	setup()
	defer tearDown()

	mux.HandleFunc("/search/users", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		respondWithJSON(w, loadFixture("user_search.json"))
	})

	url, err := SearchURL.Expand(map[string]interface{}{
		"type":  "users",
		"query": "dhruvsinghal"})
	assert.NoError(t, err)

	searchResults, result := client.Search(url).UserSearch()

	assert.False(t, result.HasError())
	assert.False(t, searchResults.IncompleteResults)
	assert.Equal(t, searchResults.TotalCount, 2)
	assert.Equal(t, len(searchResults.Items), 2)
	assert.Equal(t, searchResults.Items[0].ID, 3338555)
	assert.Equal(t, searchResults.Items[0].Login, "dhruvsinghal")
	assert.Equal(t, searchResults.Items[1].ID, 9294878)
	assert.Equal(t, searchResults.Items[1].Login, "dhruvsinghal5")
}

func TestIssueSearchService_All(t *testing.T) {
	setup()
	defer tearDown()

	mux.HandleFunc("/search/issues", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		respondWithJSON(w, loadFixture("issue_search.json"))
	})

	url, err := SearchURL.Expand(map[string]interface{}{
		"type":  "issues",
		"query": "color"})
	assert.NoError(t, err)

	searchResults, result := client.Search(url).IssueSearch()

	assert.False(t, result.HasError())
	assert.False(t, searchResults.IncompleteResults)
	assert.Equal(t, searchResults.TotalCount, 180172)
	assert.Equal(t, searchResults.Items[0].Number, 1551)
	assert.Equal(t, searchResults.Items[0].Title, "Colorizer")
	assert.Equal(t, searchResults.Items[1].Number, 3402)
	assert.Equal(t, searchResults.Items[1].Title, "Colorizer")
}