package core

type CastlingRights int

const (
	WhiteCanKsCastle CastlingRights = 1 << iota
	WhiteCanQsCastle
	BlackCanKsCastle
	BlackCanQsCastle
)
