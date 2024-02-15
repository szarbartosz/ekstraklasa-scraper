package models

type Standing struct {
	Position        int           `json:"position"`
	TeamName        string        `json:"teamName"`
	LogoUrl         string        `json:"logoUrl"`
	LastResults     []MatchResult `json:"lastResults"`
	GamesPlayed     int           `json:"gamesPlayed"`
	Wins            int           `json:"wins"`
	Draws           int           `json:"draws"`
	Losses          int           `json:"losses"`
	GoalsFor        int           `json:"goalsFor"`
	GoalsAgainst    int           `json:"goalsAgainst"`
	GoalsDifference int           `json:"goalsDifference"`
	TeamPoints      int           `json:"teamPoints"`
}
