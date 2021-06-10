package faceit

import (
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	matchID = "1-61e5c416-cb78-4951-bb1c-5a04ceea9a3e"
)

func getDummy(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/matches/")
	switch id {
	case matchID:
		w.WriteHeader(http.StatusOK)
		file, _ := os.Open("./testdata/match.json")
		b, _ := ioutil.ReadAll(file)
		w.Write(b)
	default:
		w.WriteHeader(http.StatusNotFound)
		io.WriteString(w, errorMessage)
	}
}

func TestGet(t *testing.T) {
	c := NewClient(apiKey)
	ts := httptest.NewServer(http.HandlerFunc(getDummy))
	c.host = ts.URL
	defer ts.Close()
	match, err := c.MatchClient().Get(matchID)
	assert.NoError(t, err)
	assert.Equal(t, matchID, match.MatchID)
	match, err = c.MatchClient().Get("WRONG")
	assert.Error(t, err)
}

func matchStatsDummy(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/matches/")
	id = strings.TrimSuffix(id, "/stats")
	switch id {
	case matchID:
		w.WriteHeader(http.StatusOK)
		file, _ := os.Open("./testdata/match_stats.json")
		b, _ := ioutil.ReadAll(file)
		w.Write(b)
	default:
		w.WriteHeader(http.StatusNotFound)
		io.WriteString(w, errorMessage)
	}
}

func TestMatchStats(t *testing.T) {
	c := NewClient(apiKey)
	ts := httptest.NewServer(http.HandlerFunc(matchStatsDummy))
	c.host = ts.URL
	defer ts.Close()
	stats, err := c.MatchClient().Stats(matchID)
	assert.NoError(t, err)
	for _, r := range stats.Rounds {
		assert.Equal(t, matchID, r.MatchID)
	}
	stats, err = c.MatchClient().Stats("WRONG")
	assert.Error(t, err)
}
