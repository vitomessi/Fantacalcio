package entities
//Definisco le mie strutture e l'equivalente digitura json
type Player struct {
	PlayerId   float64    `json:"id"`
	PlayerRole	string    `json:"ruolo"`
	PlayerName	string    `json:"nome"`
	PlayerTeam	string    `json:"squadra"`
	PlayerPrice   int    `json:"quotazione"`
	TeamId   int    `json:"TeamId"`
}
type Team struct {
	TeamId int    `json:"id"`
	TeamName string    `json:"nomeSquadra"`
	TeamCredit int    `json:"crediti"`
	TeamPresident string    `json:"nomePresidente"`
	TeamNumPlayer int    `json:"numPlayer"`
	TeamNumPor int    `json:"numPor"`
	TeamNumDif int    `json:"numDif"`
	TeamNumCen int    `json:"numCen"`
	TeamNumAtt int    `json:"numAtt"`
	Squadra []Player `json:"giocatori"`
}
