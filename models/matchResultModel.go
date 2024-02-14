package models

type MatchResult int

const (
	Lost MatchResult = iota
	Draw
	Won
)
