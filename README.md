# Ekstraklasa stats scraper :soccer:

## Quickstart

```bash
# launch docker container
./pull_and_compose.sh
```

## Dev quickstart

```bash
# launch live-reloaded dev version locally
air
```

## Checklist

- [x] scrap scoreboard
- [ ] scrap best ekstraklasa team stats (Puszcza Niepołomice :evergreen_tree:)
- [ ] scrap upcoming matches schedule

## Endpoints

```http
GET /table
```

```go
// returns scraped ekstraklasa table, in which each entry is of type Standing
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
```

## Motivation

<div style="display: flex;">
  <img src="./assets/58q6lo.png" alt="Ekstraklasa enjoyer"  style="padding: 5px; height: 300px;">
  <img src="./assets/8fpebz.jpg" alt="Lewandowski meme"  style="padding: 5px; height: 300px;">
<div>
