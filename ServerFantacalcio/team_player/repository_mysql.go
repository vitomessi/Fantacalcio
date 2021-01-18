package team_player

import (
	"ServerFantacalcio/config"
	"ServerFantacalcio/entities"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
)
const (
	tableTeam = "squadre"
	tablePlayer = "giocatori"
)
//METODI IMPLEMENTATI ANCHE LATO CLIENT:
/*
	1)Inserisci nuova squadra
	2)Restituisci una squadra
	3)Rimuovi una squadra
	4)Aggiungi un giocatore ad una squadra
	5)Modifica quotazione giocatore
	6)Svincola un giocatore
	7)Scambia giocatori
*/
//Inserimento di una nuova squadra
func InsertTeam(ctx context.Context, team entities.Team)error{
	db, err := config.GetDB()
	if err != nil{
		log.Fatal("Errore di connessione al database",err)
	}
	queryText := fmt.Sprintf("INSERT INTO %v (nome_squadra, nome_presidente,num_giocatori,num_portieri,num_difensori,num_centrocampisti,num_attaccanti) VALUES ('%v','%v',0,0,0,0,0)", tableTeam,
		team.TeamName,
		team.TeamPresident)
	//eseguiamo la query
	_, err = db.ExecContext(ctx, queryText)
	if err != nil{
		return err
	}
	return nil
}
//restituisce una squadra
func GetTeam(ctx context.Context,team entities.Team)(entities.Team, error){
	var squadra entities.Team

	db,err := config.GetDB()
	if err != nil{
		log.Fatal("Errore di connessione al database",err)
	}
	queryText := fmt.Sprintf("SELECT * FROM %v WHERE nome_squadra='%v'",tableTeam,team.TeamName)
	rowQuery,err := db.QueryContext(ctx,queryText)
	if err!=nil{
		log.Fatal(err)
	}
	for rowQuery.Next(){
		var team entities.Team
		err = rowQuery.Scan( &team.TeamId,
			&team.TeamName,
			&team.TeamCredit,
			&team.TeamPresident,
			&team.TeamNumPlayer,
			&team.TeamNumPor,
			&team.TeamNumDif,
			&team.TeamNumCen,
			&team.TeamNumAtt)
		squadra.TeamId = team.TeamId
		squadra.TeamName = team.TeamName
		squadra.TeamCredit = team.TeamCredit
		squadra.TeamPresident = team.TeamPresident
		squadra.TeamNumPlayer = team.TeamNumPlayer
		squadra.TeamNumPor = team.TeamNumPor
		squadra.TeamNumDif = team.TeamNumDif
		squadra.TeamNumCen = team.TeamNumCen
		squadra.TeamNumAtt = team.TeamNumAtt
		squadra.Squadra = team.Squadra
	}
	if squadra.TeamName == ""{
		err = sql.ErrNoRows
	}
	giocatore, err := db.Query("Select * FROM giocatori where team_Id=(?)",squadra.TeamId)
	if err!=nil{
		log.Fatal(err)
	}
	for giocatore.Next(){
		var player entities.Player
		err = giocatore.Scan(&player.PlayerId,
			&player.PlayerRole,
			&player.PlayerName,
			&player.PlayerTeam,
			&player.PlayerPrice,
			&player.TeamId)
		squadra.Squadra = append(squadra.Squadra,player)
	}
	return squadra,err
}
//Rimozione di una squadra
func RemoveTeam(ctx context.Context, team entities.Team)error{
	db, err := config.GetDB()
	if err != nil{
		log.Fatal("Errore di connessione al database",err)
	}
	var squadra entities.Team
	fsquadra,err := db.Query("SELECT * from squadre WHERE nome_squadra=(?)",team.TeamName)
	if err!=nil{
		log.Fatal(err)
	}
	for fsquadra.Next(){
		var team entities.Team
		err = fsquadra.Scan(&team.TeamId,
			&team.TeamName,
			&team.TeamCredit,
			&team.TeamPresident,
			&team.TeamNumPlayer,
			&team.TeamNumPor,
			&team.TeamNumDif,
			&team.TeamNumCen,
			&team.TeamNumAtt)
		squadra.TeamId = team.TeamId
		squadra.TeamName = team.TeamName
		squadra.TeamCredit = team.TeamCredit
		squadra.TeamPresident = team.TeamPresident
		squadra.TeamNumPlayer = team.TeamNumPlayer
		squadra.TeamNumPor = team.TeamNumPor
		squadra.TeamNumDif = team.TeamNumDif
		squadra.TeamNumCen = team.TeamNumCen
		squadra.TeamNumAtt = team.TeamNumAtt
	}
	if squadra.TeamName == ""{
		err = sql.ErrNoRows
	}
	queryText := fmt.Sprintf("DELETE FROM %v where team_Id='%d'", tableTeam, squadra.TeamId)
	s, err := db.ExecContext(ctx, queryText)
	if err != nil && err != sql.ErrNoRows{
		return err
	}
	queryText1 := fmt.Sprintf("UPDATE %v SET team_Id = 0 where team_Id='%d'",tablePlayer,squadra.TeamId)
	_, err = db.ExecContext(ctx, queryText1)
	if err!=nil{
		log.Fatal(err)
	}
	check, err := s.RowsAffected()
	if check == 0 {
		return errors.New("Squadra non trovata")
	}
	return nil
}
//Aggiunta di un giocatore in squadra
func AddPlayerToTeam(ctx context.Context, team entities.Team, player entities.Player)error{
	db, err := config.GetDB()
	if err != nil{
		log.Fatal("Errore di connessione al database",err)
	}
	//Mi faccio restituire la squadra a cui voglio aggiungere il giocatore e il giocatore da aggiungere
	var squadra entities.Team
	var giocatore entities.Player
	fsquadra,err := db.Query("SELECT * from squadre WHERE nome_squadra=(?)",team.TeamName)
	if err!=nil{
		log.Fatal(err)
	}
	for fsquadra.Next(){
		var team entities.Team
		err = fsquadra.Scan(&team.TeamId,
			&team.TeamName,
			&team.TeamCredit,
			&team.TeamPresident,
			&team.TeamNumPlayer,
			&team.TeamNumPor,
			&team.TeamNumDif,
			&team.TeamNumCen,
			&team.TeamNumAtt)
		squadra.TeamId = team.TeamId
		squadra.TeamName = team.TeamName
		squadra.TeamCredit = team.TeamCredit
		squadra.TeamPresident = team.TeamPresident
		squadra.TeamNumPlayer = team.TeamNumPlayer
		squadra.TeamNumPor = team.TeamNumPor
		squadra.TeamNumDif = team.TeamNumDif
		squadra.TeamNumCen = team.TeamNumCen
		squadra.TeamNumAtt = team.TeamNumAtt
	}
	if squadra.TeamName == ""{
		err = errors.New("Squadra non presente")
		log.Println(err)
	}
	fplayer,err := db.Query("SELECT * FROM giocatori WHERE nome=(?)",player.PlayerName)
	if err!=nil{
		log.Fatal(err)
	}
	for fplayer.Next(){
		var player entities.Player
		err = fplayer.Scan(&player.PlayerId,
			&player.PlayerRole,
			&player.PlayerName,
			&player.PlayerTeam,
			&player.PlayerPrice,
			&player.TeamId)
		giocatore.PlayerId = player.PlayerId
		giocatore.PlayerRole = player.PlayerRole
		giocatore.PlayerName = player.PlayerName
		giocatore.PlayerTeam = player.PlayerTeam
		giocatore.PlayerPrice = player.PlayerPrice
		giocatore.TeamId = player.TeamId
	}
	if giocatore.PlayerName == ""{
		err = errors.New("Giocatore non presente")
		log.Println(err)
	}
	if squadra.TeamNumPlayer<25 && squadra.TeamCredit>giocatore.PlayerPrice && giocatore.TeamId == 0{
		switch giocatore.PlayerRole {
		case "A":
			if squadra.TeamNumAtt<6{
				queryText1 := fmt.Sprintf("UPDATE %v set team_Id = %d where nome = '%v'",tablePlayer,squadra.TeamId,giocatore.PlayerName)
				_, err = db.ExecContext(ctx, queryText1)
				if err!=nil{
					log.Fatal(err)
				}
				queryText2 := fmt.Sprintf("UPDATE %v set crediti = %v, num_giocatori=%v, num_attaccanti=%v where team_Id=%v",tableTeam, squadra.TeamCredit-giocatore.PlayerPrice,squadra.TeamNumPlayer+1,squadra.TeamNumAtt+1,squadra.TeamId)
				_, err = db.ExecContext(ctx, queryText2)
				if err!=nil{
					log.Fatal(err)
				}

			}else{
				err = errors.New("Troppi Attaccanti")
				log.Println(err)
			}
		case "C":
			if squadra.TeamNumCen<8{
				queryText1 := fmt.Sprintf("UPDATE %v set team_Id = %d where nome = '%v'",tablePlayer,squadra.TeamId,giocatore.PlayerName)
				_, err = db.ExecContext(ctx, queryText1)
				if err!=nil{
					log.Fatal(err)
				}
				queryText2 := fmt.Sprintf("UPDATE %v set crediti = %v, num_giocatori=%v, num_centrocampisti=%v where team_Id=%v",tableTeam, squadra.TeamCredit-giocatore.PlayerPrice,squadra.TeamNumPlayer+1,squadra.TeamNumCen+1,squadra.TeamId)
				_, err = db.ExecContext(ctx, queryText2)
				if err!=nil{
					log.Fatal(err)
				}

			}else{
				err = errors.New("Troppi Centrocampisti")
				log.Println(err)
			}
		case "D":
			if squadra.TeamNumDif<8{
				queryText1 := fmt.Sprintf("UPDATE %v set team_Id = %d where nome = '%v'",tablePlayer,squadra.TeamId,giocatore.PlayerName)
				_, err = db.ExecContext(ctx, queryText1)
				if err!=nil{
					log.Fatal(err)
				}
				queryText2 := fmt.Sprintf("UPDATE %v set crediti = %v, num_giocatori=%v, num_difensori=%v where team_Id=%v",tableTeam, squadra.TeamCredit-giocatore.PlayerPrice,squadra.TeamNumPlayer+1,squadra.TeamNumDif+1,squadra.TeamId)
				_, err = db.ExecContext(ctx, queryText2)
				if err!=nil{
					log.Fatal(err)
				}

			}else{
				err = errors.New("Troppi Difensori")
				log.Println(err)
			}

		case "P":
			if squadra.TeamNumPor<3{
				queryText1 := fmt.Sprintf("UPDATE %v set team_Id = %d where nome = '%v'",tablePlayer,squadra.TeamId,giocatore.PlayerName)
				_, err = db.ExecContext(ctx, queryText1)
				if err!=nil{
					log.Fatal(err)
				}
				queryText2 := fmt.Sprintf("UPDATE %v set crediti = %v, num_giocatori=%v, num_portieri=%v where team_Id=%v",tableTeam, (squadra.TeamCredit-giocatore.PlayerPrice),squadra.TeamNumPlayer+1,squadra.TeamNumPor+1,squadra.TeamId)
				_, err = db.ExecContext(ctx, queryText2)
				if err!=nil{
					log.Fatal(err)
				}

			}else{
				err = errors.New("Troppi Portieri")
				log.Println(err)
			}

		}
	}else{
		err =  errors.New("Operazione non consentita")
		log.Println(err)
	}
	return err
}
//Aggiornamento valore giocatore
func UpdatePlayer(ctx context.Context, player entities.Player)error{
	db, err := config.GetDB()
	if err != nil{
		log.Fatal("Errore di connessione al database",err)
	}
	queryText := fmt.Sprintf("UPDATE %v set quotazione=%d where nome = '%v'",tablePlayer,player.PlayerPrice,player.PlayerName)
	_,err = db.ExecContext(ctx,queryText)
	if err!=nil{
		return err
	}
	return nil
}
//Rimozione di un giocatore in squadra
func RemovePlayerToTeam(ctx context.Context, team entities.Team, player entities.Player)error{
	db, err := config.GetDB()
	if err != nil{
		log.Fatal("Errore di connessione al database",err)
	}
	//Mi faccio restituire la squadra da cui voglio rimuovere il giocatore e il giocatore da rimuovere
	var squadra entities.Team
	var giocatore entities.Player
	fsquadra,err := db.Query("SELECT * from squadre WHERE nome_squadra=(?)",team.TeamName)
	if err!=nil{
		log.Fatal(err)
	}
	for fsquadra.Next(){
		var team entities.Team
		err = fsquadra.Scan(&team.TeamId,
			&team.TeamName,
			&team.TeamCredit,
			&team.TeamPresident,
			&team.TeamNumPlayer,
			&team.TeamNumPor,
			&team.TeamNumDif,
			&team.TeamNumCen,
			&team.TeamNumAtt)
		squadra.TeamId = team.TeamId
		squadra.TeamName = team.TeamName
		squadra.TeamCredit = team.TeamCredit
		squadra.TeamPresident = team.TeamPresident
		squadra.TeamNumPlayer = team.TeamNumPlayer
		squadra.TeamNumPor = team.TeamNumPor
		squadra.TeamNumDif = team.TeamNumDif
		squadra.TeamNumCen = team.TeamNumCen
		squadra.TeamNumAtt = team.TeamNumAtt
	}
	if squadra.TeamName == ""{
		err = sql.ErrNoRows
	}
	fplayer,err := db.Query("SELECT * FROM giocatori WHERE nome=(?)",player.PlayerName)
	if err!=nil{
		log.Fatal(err)
	}
	for fplayer.Next(){
		var player entities.Player
		err = fplayer.Scan(&player.PlayerId,
			&player.PlayerRole,
			&player.PlayerName,
			&player.PlayerTeam,
			&player.PlayerPrice,
			&player.TeamId)
		giocatore.PlayerId = player.PlayerId
		giocatore.PlayerRole = player.PlayerRole
		giocatore.PlayerName = player.PlayerName
		giocatore.PlayerTeam = player.PlayerTeam
		giocatore.PlayerPrice = player.PlayerPrice
		giocatore.TeamId = player.TeamId
	}
	if giocatore.PlayerName == ""{
		err = sql.ErrNoRows
	}
	if giocatore.TeamId != 0{
		switch giocatore.PlayerRole {
		case "A":
			queryText1 := fmt.Sprintf("UPDATE %v set team_Id = 0 where nome = '%v'",tablePlayer,giocatore.PlayerName)
			_, err = db.ExecContext(ctx, queryText1)
			if err!=nil{
				log.Fatal(err)
			}
			queryText2 := fmt.Sprintf("UPDATE %v set crediti = %v, num_giocatori=%v, num_attaccanti=%v where team_Id=%v",tableTeam, squadra.TeamCredit+(giocatore.PlayerPrice/2),squadra.TeamNumPlayer-1,squadra.TeamNumAtt-1,squadra.TeamId)
			_, err = db.ExecContext(ctx, queryText2)
			if err!=nil{
				log.Fatal(err)
			}


		case "C":
			queryText1 := fmt.Sprintf("UPDATE %v set team_Id = 0 where nome = '%v'",tablePlayer,giocatore.PlayerName)
			_, err = db.ExecContext(ctx, queryText1)
			if err!=nil{
				log.Fatal(err)
			}
			queryText2 := fmt.Sprintf("UPDATE %v set crediti = %v, num_giocatori=%v, num_centrocampisti=%v where team_Id=%v",tableTeam, squadra.TeamCredit+(giocatore.PlayerPrice/2),squadra.TeamNumPlayer-1,squadra.TeamNumCen-1,squadra.TeamId)
			_, err = db.ExecContext(ctx, queryText2)
			if err!=nil{
				log.Fatal(err)
			}
		case "D":
			queryText1 := fmt.Sprintf("UPDATE %v set team_Id = 0 where nome = '%v'",tablePlayer,giocatore.PlayerName)
			_, err = db.ExecContext(ctx, queryText1)
			if err!=nil{
				log.Fatal(err)
			}
			queryText2 := fmt.Sprintf("UPDATE %v set crediti = %v, num_giocatori=%v, num_difensori=%v where team_Id=%v",tableTeam, squadra.TeamCredit+(giocatore.PlayerPrice/2),squadra.TeamNumPlayer-1,squadra.TeamNumDif-1,squadra.TeamId)
			_, err = db.ExecContext(ctx, queryText2)
			if err!=nil {
				log.Fatal(err)
			}

		case "P":
			queryText1 := fmt.Sprintf("UPDATE %v set team_Id = 0 where nome = '%v'",tablePlayer,giocatore.PlayerName)
			_, err = db.ExecContext(ctx, queryText1)
			if err!=nil{
				log.Fatal(err)
			}
			queryText2 := fmt.Sprintf("UPDATE %v set crediti = %v, num_giocatori=%v, num_portieri=%v where team_Id=%v",tableTeam, squadra.TeamCredit+(giocatore.PlayerPrice/2),squadra.TeamNumPlayer-1,squadra.TeamNumPor-1,squadra.TeamId)
			_, err = db.ExecContext(ctx, queryText2)
			if err!=nil{
				log.Fatal(err)
			}
		}
	}else{
		err =  errors.New("Operazione non consentita")
		log.Println(err)
	}

	if err != nil && err!=sql.ErrNoRows{
		return err
	}
	return nil
}
//Scambio tra due giocatori
func ScambiaPlayer(ctx context.Context, team1 entities.Team, team2 entities.Team,player1 entities.Player,player2 entities.Player)error{
	var squadre [2]entities.Team
	var giocatori [2]entities.Player
	db, err := config.GetDB()
	if err != nil{
		log.Fatal("Errore di connessione al database",err)
	}
	fsquadra,err := db.Query("SELECT * from squadre WHERE nome_squadra=(?) ",team1.TeamName)
	if err!=nil{
		log.Fatal(err)
	}
	for fsquadra.Next(){
		var team entities.Team
		err = fsquadra.Scan(&team.TeamId,
			&team.TeamName,
			&team.TeamCredit,
			&team.TeamPresident,
			&team.TeamNumPlayer,
			&team.TeamNumPor,
			&team.TeamNumDif,
			&team.TeamNumCen,
			&team.TeamNumAtt)
		squadre[0] = team
	}

	fsquadra1,err := db.Query("SELECT * from squadre WHERE nome_squadra=(?) ",team2.TeamName)
	if err!=nil{
		log.Fatal(err)
	}
	for fsquadra1.Next(){
		var team entities.Team
		err = fsquadra1.Scan(&team.TeamId,
			&team.TeamName,
			&team.TeamCredit,
			&team.TeamPresident,
			&team.TeamNumPlayer,
			&team.TeamNumPor,
			&team.TeamNumDif,
			&team.TeamNumCen,
			&team.TeamNumAtt)
		squadre[1] = team
	}
	fgiocatore,err := db.Query("SELECT * from giocatori WHERE nome=(?) ",player1.PlayerName)
	if err!=nil{
		log.Fatal(err)
	}
	for fgiocatore.Next(){
		var player entities.Player
		err = fgiocatore.Scan(&player.PlayerId,
			&player.PlayerRole,
			&player.PlayerName,
			&player.PlayerTeam,
			&player.PlayerPrice,
			&player.TeamId)
		giocatori[0] = player
	}
	fgiocatore1,err := db.Query("SELECT * from giocatori WHERE nome=(?) ",player2.PlayerName)
	if err!=nil{
		log.Fatal(err)
	}
	for fgiocatore1.Next(){
		var player entities.Player
		err = fgiocatore1.Scan(&player.PlayerId,
			&player.PlayerRole,
			&player.PlayerName,
			&player.PlayerTeam,
			&player.PlayerPrice,
			&player.TeamId)
		giocatori[1] = player
	}

	if squadre[0].TeamId == giocatori[0].TeamId && squadre[1].TeamId == giocatori[1].TeamId && giocatori[0].PlayerRole==giocatori[1].PlayerRole{
		queryText1 := fmt.Sprintf("UPDATE %v set team_Id = %d where nome = '%v'",tablePlayer,giocatori[1].TeamId,giocatori[0].PlayerName)
		_, err = db.ExecContext(ctx, queryText1)
		if err!=nil{
			log.Fatal(err)
		}
		queryText2 := fmt.Sprintf("UPDATE %v set team_Id = %d where nome = '%v'",tablePlayer,giocatori[0].TeamId,giocatori[1].PlayerName)
		_, err = db.ExecContext(ctx, queryText2)
		if err!=nil{
			log.Fatal(err)
		}
	}else{
		err =  errors.New("Operazione non consentita")
		log.Println(err)
		return err
	}
	return nil
}
/*METODI NON IMPLEMENTATI LATO CLIENT
*/
//restituisce tutte le squadre
func GetTeams(ctx context.Context)([]entities.Team, error){
	var squadre []entities.Team

	db,err := config.GetDB()
	if err != nil{
		log.Fatal("Errore di connessione al database",err)
	}
	queryText := fmt.Sprintf("SELECT * FROM %v ORDER by team_Id ASC" +
		" ", tableTeam)
	rowQuery,err := db.QueryContext(ctx,queryText)
	if err!=nil{
		log.Fatal(err)
	}
	for rowQuery.Next(){
		var team entities.Team
		err = rowQuery.Scan(&team.TeamId,
			&team.TeamName,
			&team.TeamCredit,
			&team.TeamPresident,
			&team.TeamNumPlayer,
			&team.TeamNumPor,
			&team.TeamNumDif,
			&team.TeamNumCen,
			&team.TeamNumAtt)
		giocatore, err := db.Query("Select * FROM giocatori where team_Id=(?)",team.TeamId)
		if err!=nil{
			log.Fatal(err)
		}
		for giocatore.Next(){
			var player entities.Player
			err = giocatore.Scan(&player.PlayerId,
				&player.PlayerRole,
				&player.PlayerName,
				&player.PlayerTeam,
				&player.PlayerPrice,
				&player.TeamId)
			team.Squadra = append(team.Squadra,player)
		}
		squadre = append(squadre,team)
	}
	return squadre,nil
}
//restituisce un giocatore
func GetPlayer(ctx context.Context,player entities.Player)(entities.Player, error){
	var giocatore entities.Player
	db,err := config.GetDB()
	if err != nil{
		log.Fatal("Errore di connessione al database",err)
	}
	queryText := fmt.Sprintf("SELECT * FROM %v WHERE nome = '%v'",tablePlayer,player.PlayerName)
	rowQuery,err := db.QueryContext(ctx,queryText)
	if err!=nil{
		log.Fatal(err)
	}
	for rowQuery.Next(){
		var player entities.Player
		err = rowQuery.Scan(&player.PlayerId,
			&player.PlayerRole,
			&player.PlayerName,
			&player.PlayerTeam,
			&player.PlayerPrice,
			&player.TeamId)
		giocatore.PlayerId = player.PlayerId
		giocatore.PlayerRole = player.PlayerRole
		giocatore.PlayerName = player.PlayerName
		giocatore.PlayerTeam = player.PlayerTeam
		giocatore.PlayerPrice = player.PlayerPrice
		giocatore.TeamId = player.TeamId

	}
	return giocatore,nil
}
//restituisce tutti i giocatori svincolati
func GetPlayers(ctx context.Context)([]entities.Player, error){
	var giocatori []entities.Player

	db,err := config.GetDB()
	if err != nil{
		log.Fatal("Errore di connessione al database",err)
	}
	queryText := fmt.Sprintf("SELECT * FROM %v WHERE team_Id = 0 ORDER by nome ASC", tablePlayer)
	rowQuery,err := db.QueryContext(ctx,queryText)
	if err!=nil{
		log.Fatal(err)
	}
	for rowQuery.Next(){
		var player entities.Player
		err = rowQuery.Scan(&player.PlayerId,
			&player.PlayerRole,
			&player.PlayerName,
			&player.PlayerTeam,
			&player.PlayerPrice,
			&player.TeamId)
		giocatori = append(giocatori,player)
	}
	return giocatori,nil
}




