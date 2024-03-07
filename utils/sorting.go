package utils

import "scraper/ekstraklasa/models"

type GameSlice []models.Game

func (g GameSlice) Len() int {
	return len(g)
}

func (g GameSlice) Less(i, j int) bool {
	return g[i].DateTime.Before(g[j].DateTime)
}

func (g GameSlice) Swap(i, j int) {
	g[i], g[j] = g[j], g[i]
}
