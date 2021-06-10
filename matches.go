package faceit

type MatchTeam struct {
	FactionID string `json:"faction_id"`
	Leader    string `json:"leader"`
	Avatar    string `json:"avatar"`
	Roster    []struct {
		PlayerID          string `json:"player_id"`
		Nickname          string `json:"nickname"`
		Avatar            string `json:"avatar"`
		Membership        string `json:"membership"`
		GamePlayerID      string `json:"game_player_id"`
		GamePlayerName    string `json:"game_player_name"`
		GameSkillLevel    int    `json:"game_skill_level"`
		AnticheatRequired bool   `json:"anticheat_required"`
	} `json:"roster"`
	Substituted bool   `json:"substituted"`
	Name        string `json:"name"`
	Type        string `json:"type"`
}

type Match struct {
	MatchID         string               `json:"match_id"`
	Version         int                  `json:"version"`
	Game            string               `json:"game"`
	Region          string               `json:"region"`
	CompetitionID   string               `json:"competition_id"`
	CompetitionType string               `json:"competition_type"`
	CompetitionName string               `json:"competition_name"`
	OrganizerID     string               `json:"organizer_id"`
	Teams           map[string]MatchTeam `json:"teams"`
	Voting          struct {
		VotedEntityTypes []string `json:"voted_entity_types"`
		Location         struct {
			Entities []struct {
				ClassName      string `json:"className"`
				GameLocationID string `json:"game_location_id"`
				GUID           string `json:"guid"`
				ImageLg        string `json:"image_lg"`
				ImageSm        string `json:"image_sm"`
				Name           string `json:"name"`
			} `json:"entities"`
			Pick []string `json:"pick"`
		} `json:"location"`
		Map struct {
			Entities []struct {
				ClassName string `json:"class_name"`
				GameMapID string `json:"game_map_id"`
				GUID      string `json:"guid"`
				ImageLg   string `json:"image_lg"`
				ImageSm   string `json:"image_sm"`
				Name      string `json:"name"`
			} `json:"entities"`
			Pick []string `json:"pick"`
		} `json:"map"`
	} `json:"voting"`
	CalculateElo bool     `json:"calculate_elo"`
	ConfiguredAt int      `json:"configured_at"`
	StartedAt    int      `json:"started_at"`
	FinishedAt   int      `json:"finished_at"`
	DemoURL      []string `json:"demo_url"`
	ChatRoomID   string   `json:"chat_room_id"`
	BestOf       int      `json:"best_of"`
	Results      struct {
		Winner string         `json:"winner"`
		Score  map[string]int `json:"score"`
	} `json:"results"`
	Status    string `json:"status"`
	FaceitURL string `json:"faceit_url"`
}

type MatchStats struct {
	Rounds []struct {
		BestOf        string      `json:"best_of"`
		CompetitionID interface{} `json:"competition_id"`
		GameID        string      `json:"game_id"`
		GameMode      string      `json:"game_mode"`
		MatchID       string      `json:"match_id"`
		MatchRound    string      `json:"match_round"`
		Played        string      `json:"played"`
		RoundStats    struct {
			Region string `json:"Region"`
			Map    string `json:"Map"`
			Rounds string `json:"Rounds"`
			Winner string `json:"Winner"`
			Score  string `json:"Score"`
		} `json:"round_stats"`
		Teams []struct {
			TeamID    string `json:"team_id"`
			Premade   bool   `json:"premade"`
			TeamStats struct {
				TeamHeadshots   string `json:"Team Headshots"`
				FinalScore      string `json:"Final Score"`
				OvertimeScore   string `json:"Overtime score"`
				FirstHalfScore  string `json:"First Half Score"`
				SecondHalfScore string `json:"Second Half Score"`
				Team            string `json:"Team"`
				TeamWin         string `json:"Team Win"`
			} `json:"team_stats"`
			Players []struct {
				PlayerID    string `json:"player_id"`
				Nickname    string `json:"nickname"`
				PlayerStats struct {
					PentaKills  string `json:"Penta Kills"`
					KDRatio     string `json:"K/D Ratio"`
					HeadshotsPr string `json:"Headshots %"`
					MVPs        string `json:"MVPs"`
					Headshots   string `json:"Headshots"`
					TripleKills string `json:"Triple Kills"`
					KRRatio     string `json:"K/R Ratio"`
					Result      string `json:"Result"`
					QuadroKills string `json:"Quadro Kills"`
					Deaths      string `json:"Deaths"`
					Kills       string `json:"Kills"`
					Assists     string `json:"Assists"`
				} `json:"player_stats"`
			} `json:"players"`
		} `json:"teams"`
	} `json:"rounds"`
}

type MatchClient struct {
	*Client
}

func (c *Client) MatchClient() *MatchClient {
	return &MatchClient{c}
}

func (c *MatchClient) Get(id string) (Match, error) {
	match := Match{}
	err := c.sendRequest("/matches/"+id, &match)
	return match, err
}

func (c *MatchClient) Stats(id string) (MatchStats, error) {
	stats := MatchStats{}
	err := c.sendRequest("/matches/"+id+"/stats", &stats)
	return stats, err
}
