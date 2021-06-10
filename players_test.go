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
	apiKey       = "API_KEY"
	nickname     = "testnickname"
	playerID     = "bb75c4d7-8167-4494-9a9d-efdf9bda3862"
	game         = "csgo"
	errorMessage = `{
		"errors": [
		  {
			"message": "Error message",
			"code": "err_nf0",
			"http_status": 400,
			"parameters": []
		  }
		]
	  }`
)

func findDummy(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("nickname")
	switch name {
	case nickname:
		w.WriteHeader(http.StatusOK)
		file, _ := os.Open("./testdata/player.json")
		b, _ := ioutil.ReadAll(file)
		w.Write(b)
	default:
		w.WriteHeader(http.StatusNotFound)
		io.WriteString(w, errorMessage)
	}
}

func TestFind(t *testing.T) {
	c := NewClient(apiKey)
	ts := httptest.NewServer(http.HandlerFunc(findDummy))
	c.host = ts.URL
	defer ts.Close()
	req := FindPlayerRequest{
		Nickname: nickname,
	}
	player, err := c.PlayerClient().Find(req)
	assert.NoError(t, err)
	assert.Equal(t, nickname, player.Nickname)
	req.Nickname = "404"
	player, err = c.PlayerClient().Find(req)
	assert.Error(t, err)
}

func playerDummy(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/players/")
	switch id {
	case playerID:
		w.WriteHeader(http.StatusOK)
		file, _ := os.Open("./testdata/player.json")
		b, _ := ioutil.ReadAll(file)
		w.Write(b)
	default:
		w.WriteHeader(http.StatusNotFound)
		io.WriteString(w, errorMessage)
	}
}

func TestPlayer(t *testing.T) {
	c := NewClient(apiKey)
	ts := httptest.NewServer(http.HandlerFunc(playerDummy))
	c.host = ts.URL
	defer ts.Close()
	player, err := c.PlayerClient().Get(playerID)
	assert.NoError(t, err)
	assert.Equal(t, playerID, player.PlayerID)
	player, err = c.PlayerClient().Get("WRONG")
	assert.Error(t, err)
}

func historyDummy(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/players/")
	id = strings.TrimSuffix(id, "/history")
	switch id {
	case playerID:
		w.WriteHeader(http.StatusOK)
		file, _ := os.Open("./testdata/history.json")
		b, _ := ioutil.ReadAll(file)
		w.Write(b)
	default:
		w.WriteHeader(http.StatusOK)
		file, _ := os.Open("./testdata/empty_history.json")
		b, _ := ioutil.ReadAll(file)
		w.Write(b)
	}
}

func TestHistory(t *testing.T) {
	c := NewClient(apiKey)
	ts := httptest.NewServer(http.HandlerFunc(historyDummy))
	c.host = ts.URL
	defer ts.Close()
	req := HistoryRequest{
		PlayerID: playerID,
	}
	history, err := c.PlayerClient().History(req)
	assert.NoError(t, err)
	var found bool
	assert.Equal(t, 1, len(history.Items))
	for _, item := range history.Items {
		found = false
		for _, player := range item.PlayingPlayers {
			if player == playerID {
				found = true
				break
			}
		}
		assert.True(t, found)
	}
	req.PlayerID = "WRONG"
	history, err = c.PlayerClient().History(req)
	assert.NoError(t, err)
	assert.Empty(t, history.Items)
}

func hubsDummy(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/players/")
	id = strings.TrimSuffix(id, "/hubs")
	switch id {
	case playerID:
		w.WriteHeader(http.StatusOK)
		file, _ := os.Open("./testdata/hubs.json")
		b, _ := ioutil.ReadAll(file)
		w.Write(b)
	default:
		w.WriteHeader(http.StatusOK)
		file, _ := os.Open("./testdata/empty_hubs.json")
		b, _ := ioutil.ReadAll(file)
		w.Write(b)
	}
}

func TestHubs(t *testing.T) {
	c := NewClient(apiKey)
	ts := httptest.NewServer(http.HandlerFunc(hubsDummy))
	c.host = ts.URL
	defer ts.Close()
	req := HubsRequest{
		PlayerID: playerID,
	}
	hubs, err := c.PlayerClient().Hubs(req)
	assert.NoError(t, err)
	assert.NotEmpty(t, hubs.Items)
	req.PlayerID = "WRONG"
	hubs, err = c.PlayerClient().Hubs(req)
	assert.NoError(t, err)
	assert.Empty(t, hubs.Items)
}

func statsDummy(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/players/")
	id = strings.TrimSuffix(id, "/stats/"+game)
	switch id {
	case playerID:
		w.WriteHeader(http.StatusOK)
		file, _ := os.Open("./testdata/stats.json")
		b, _ := ioutil.ReadAll(file)
		w.Write(b)
	default:
		w.WriteHeader(http.StatusNotFound)
		io.WriteString(w, errorMessage)
	}
}

func TestStats(t *testing.T) {
	c := NewClient(apiKey)
	ts := httptest.NewServer(http.HandlerFunc(statsDummy))
	c.host = ts.URL
	defer ts.Close()
	req := StatsRequest{
		PlayerID: playerID,
		GameID:   game,
	}
	stats, err := c.PlayerClient().Stats(req)
	assert.NoError(t, err)
	assert.Equal(t, game, stats.GameID)
	assert.Equal(t, playerID, stats.PlayerID)
	assert.NotEmpty(t, stats.Segments)
	req.PlayerID = "WRONG"
	stats, err = c.PlayerClient().Stats(req)
	assert.Error(t, err)
	assert.Empty(t, stats.Segments)
}

func tDummy(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/players/")
	id = strings.TrimSuffix(id, "/tournaments")
	switch id {
	case playerID:
		w.WriteHeader(http.StatusOK)
		file, _ := os.Open("./testdata/tournaments.json")
		b, _ := ioutil.ReadAll(file)
		w.Write(b)
	default:
		w.WriteHeader(http.StatusNotFound)
		io.WriteString(w, errorMessage)
	}
}

func TestTournaments(t *testing.T) {
	c := NewClient(apiKey)
	ts := httptest.NewServer(http.HandlerFunc(tDummy))
	c.host = ts.URL
	defer ts.Close()
	req := TournamentsRequest{
		PlayerID: playerID,
	}
	tournaments, err := c.PlayerClient().Tournaments(req)
	assert.NoError(t, err)
	assert.NotEmpty(t, tournaments.Items)
	req.PlayerID = "WRONG"
	tournaments, err = c.PlayerClient().Tournaments(req)
	assert.Error(t, err)
	assert.Empty(t, tournaments.Items)
}
