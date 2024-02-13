package models

type Standing struct {
	Position        int
	TeamName        string
	GamesPlayed     int
	Wins            int
	Draws           int
	Losses          int
	GoalsFor        int
	GoalsAgainst    int
	GoalsDifference int
	TeamPoints      int
}
