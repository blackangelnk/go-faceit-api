package faceit

import (
	"net/url"
)

type Region struct {
	SelectedLadderID string `json:"selected_ladder_id"`
}

type Game struct {
	GameProfileID   string            `json:"game_profile_id"`
	Region          string            `json:"region"`
	Regions         map[string]Region `json:"regions"`
	SkillLevelLabel string            `json:"skill_level_label"`
	GamePlayerID    string            `json:"game_player_id"`
	SkillLevel      int               `json:"skill_level"`
	FaceitElo       int               `json:"faceit_elo"`
	GamePlayerName  string            `json:"game_player_name"`
}

type Player struct {
	PlayerID           string `json:"player_id"`
	Nickname           string `json:"nickname"`
	Avatar             string `json:"avatar"`
	Country            string `json:"country"`
	CoverImage         string `json:"cover_image"`
	CoverFeaturedImage string `json:"cover_featured_image"`
	Infractions        struct {
		LastInfractionDate string `json:"last_infraction_date"`
		Afk                int    `json:"afk"`
		Leaver             int    `json:"leaver"`
		QmNotCheckedin     int    `json:"qm_not_checkedin"`
		QmNotVoted         int    `json:"qm_not_voted"`
	} `json:"infractions"`
	Platforms map[string]string `json:"platforms"`
	Games     map[string]Game   `json:"games"`
	Settings  struct {
		Language string `json:"language"`
	} `json:"settings"`
	FriendsIds     []string `json:"friends_ids"`
	NewSteamID     string   `json:"new_steam_id"`
	SteamID64      string   `json:"steam_id_64"`
	SteamNickname  string   `json:"steam_nickname"`
	MembershipType string   `json:"membership_type"`
	Memberships    []string `json:"memberships"`
	FaceitURL      string   `json:"faceit_url"`
}

type Team struct {
	TeamID   string `json:"team_id"`
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar"`
	Type     string `json:"type"`
	Players  []struct {
		PlayerID       string `json:"player_id"`
		Nickname       string `json:"nickname"`
		Avatar         string `json:"avatar"`
		SkillLevel     int    `json:"skill_level"`
		GamePlayerID   string `json:"game_player_id"`
		GamePlayerName string `json:"game_player_name"`
		FaceitURL      string `json:"faceit_url"`
	} `json:"players"`
}

type HistoryItem struct {
	MatchID         string          `json:"match_id"`
	GameID          string          `json:"game_id"`
	Region          string          `json:"region"`
	MatchType       string          `json:"match_type"`
	GameMode        string          `json:"game_mode"`
	MaxPlayers      int             `json:"max_players"`
	TeamsSize       int             `json:"teams_size"`
	Teams           map[string]Team `json:"teams"`
	PlayingPlayers  []string        `json:"playing_players"`
	CompetitionID   string          `json:"competition_id"`
	CompetitionName string          `json:"competition_name"`
	CompetitionType string          `json:"competition_type"`
	OrganizerID     string          `json:"organizer_id"`
	Status          string          `json:"status"`
	StartedAt       int             `json:"started_at"`
	FinishedAt      int             `json:"finished_at"`
	Results         struct {
		Winner string `json:"winner"`
		Score  struct {
			Faction1 int `json:"faction1"`
			Faction2 int `json:"faction2"`
		} `json:"score"`
	} `json:"results"`
	FaceitURL string `json:"faceit_url"`
}

type History struct {
	Items []HistoryItem `json:"items"`
	rPagination
	rTime
}

type Hub struct {
	HubID       string `json:"hub_id"`
	Name        string `json:"name"`
	Avatar      string `json:"avatar"`
	GameID      string `json:"game_id"`
	OrganizerID string `json:"organizer_id"`
	FaceitURL   string `json:"faceit_url"`
}

type Hubs struct {
	Items []Hub `json:"items"`
	rPagination
}

type Tournament struct {
	AnticheatRequired           bool   `json:"anticheat_required"`
	Custom                      bool   `json:"custom"`
	FaceitURL                   string `json:"faceit_url"`
	FeaturedImage               string `json:"featured_image"`
	GameID                      string `json:"game_id"`
	InviteType                  string `json:"invite_type"`
	MatchType                   string `json:"match_type"`
	MaxSkill                    int    `json:"max_skill"`
	MembershipType              string `json:"membership_type"`
	MinSkill                    int    `json:"min_skill"`
	Name                        string `json:"name"`
	NumberOfPlayers             int    `json:"number_of_players"`
	NumberOfPlayersCheckedin    int    `json:"number_of_players_checkedin"`
	NumberOfPlayersJoined       int    `json:"number_of_players_joined"`
	NumberOfPlayersParticipants int    `json:"number_of_players_participants"`
	OrganizerID                 string `json:"organizer_id"`
	PrizeType                   string `json:"prize_type"`
	Region                      string `json:"region"`
	StartedAt                   int    `json:"started_at"`
	Status                      string `json:"status"`
	SubscriptionsCount          int    `json:"subscriptions_count"`
	TeamSize                    int    `json:"team_size"`
	TotalPrize                  struct {
	} `json:"total_prize"`
	TournamentID       string   `json:"tournament_id"`
	WhitelistCountries []string `json:"whitelist_countries"`
}

type Tournaments struct {
	Items []Tournament `json:"items"`
	rPagination
}

type Stats struct {
	PlayerID string `json:"player_id"`
	GameID   string `json:"game_id"`
	Lifetime struct {
		WinRate          string   `json:"Win Rate %"`
		AverageHeadshots string   `json:"Average Headshots %"`
		RecentResults    []string `json:"Recent Results"`
		Wins             string   `json:"Wins"`
		KDRatio          string   `json:"K/D Ratio"`
		CurrentWinStreak string   `json:"Current Win Streak"`
		LongestWinStreak string   `json:"Longest Win Streak"`
		TotalHeadshots   string   `json:"Total Headshots %"`
		AverageKDRatio   string   `json:"Average K/D Ratio"`
		Matches          string   `json:"Matches"`
	} `json:"lifetime"`
	Segments []struct {
		Label      string `json:"label"`
		ImgSmall   string `json:"img_small"`
		ImgRegular string `json:"img_regular"`
		Stats      struct {
			AverageMVPs        string `json:"Average MVPs"`
			Headshots          string `json:"Headshots"`
			HeadshotsPerMatch  string `json:"Headshots per Match"`
			KRRatio            string `json:"K/R Ratio"`
			TotalHeadshots     string `json:"Total Headshots %"`
			AverageKills       string `json:"Average Kills"`
			Matches            string `json:"Matches"`
			Deaths             string `json:"Deaths"`
			AveragePentaKills  string `json:"Average Penta Kills"`
			AverageDeaths      string `json:"Average Deaths"`
			Assists            string `json:"Assists"`
			AverageTripleKills string `json:"Average Triple Kills"`
			AverageQuadroKills string `json:"Average Quadro Kills"`
			AverageAssists     string `json:"Average Assists"`
			AverageKRRatio     string `json:"Average K/R Ratio"`
			WinRate            string `json:"Win Rate %"`
			Rounds             string `json:"Rounds"`
			PentaKills         string `json:"Penta Kills"`
			Kills              string `json:"Kills"`
			KDRatio            string `json:"K/D Ratio"`
			TripleKills        string `json:"Triple Kills"`
			MVPs               string `json:"MVPs"`
			AverageHeadshots   string `json:"Average Headshots %"`
			Wins               string `json:"Wins"`
			QuadroKills        string `json:"Quadro Kills"`
			AverageKDRatio     string `json:"Average K/D Ratio"`
		} `json:"stats"`
		Type string `json:"type"`
		Mode string `json:"mode"`
	} `json:"segments"`
}

type FindPlayerRequest struct {
	Nickname     string
	Game         string
	GamePlayerID string
}

type HistoryRequest struct {
	PlayerID string
	Game     string
	pagination
	timestamps
}

type HubsRequest struct {
	PlayerID string
	pagination
}

type TournamentsRequest struct {
	PlayerID string
	pagination
}

type StatsRequest struct {
	PlayerID string
	GameID   string
}

type PlayerClient struct {
	*Client
}

func (c *Client) PlayerClient() *PlayerClient {
	return &PlayerClient{c}
}

func (c *PlayerClient) Find(req FindPlayerRequest) (Player, error) {
	player := Player{}
	err := c.sendRequest("/players?"+url.Values{
		"nickname":       {req.Nickname},
		"game":           {req.Game},
		"game_player_id": {req.GamePlayerID},
	}.Encode(), &player)
	return player, err
}

func (c *PlayerClient) Get(id string) (Player, error) {
	player := Player{}
	err := c.sendRequest("/players/"+id, &player)
	return player, err
}

func (c *PlayerClient) History(req HistoryRequest) (History, error) {
	v := url.Values{
		"game": {req.Game},
	}
	req.pagination.toValues(&v)
	req.timestamps.toValues(&v)
	history := History{}
	err := c.sendRequest("/players/"+req.PlayerID+"/history?"+v.Encode(), &history)
	return history, err
}

func (c *PlayerClient) Hubs(req HubsRequest) (Hubs, error) {
	v := url.Values{}
	req.pagination.toValues(&v)
	hubs := Hubs{}
	err := c.sendRequest("/players/"+req.PlayerID+"/hubs?"+v.Encode(), &hubs)
	return hubs, err
}

func (c *PlayerClient) Stats(req StatsRequest) (Stats, error) {
	stats := Stats{}
	err := c.sendRequest("/players/"+req.PlayerID+"/stats/"+req.GameID, &stats)
	return stats, err
}

func (c *PlayerClient) Tournaments(req TournamentsRequest) (Tournaments, error) {
	v := url.Values{}
	req.pagination.toValues(&v)
	t := Tournaments{}
	err := c.sendRequest("/players/"+req.PlayerID+"/tournaments?"+v.Encode(), &t)
	return t, err
}
