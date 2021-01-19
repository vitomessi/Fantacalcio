package team_player

import (
	"ServerFantacalcio/config"
	"ServerFantacalcio/entities"
	"database/sql"
	"errors"
	"log"
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
func InsertTeam(team entities.Team)error{
	db, err := config.GetDB()
	if err != nil{
		log.Fatal("Errore di connessione al database",err)
	}
	//Eseguo la query
	_,err = db.Exec("INSERT INTO squadre (nome_squadra, nome_presidente,num_giocatori,num_portieri,num_difensori,num_centrocampisti,num_attaccanti) VALUES ((?),(?),0,0,0,0,0)",team.TeamName,team.TeamPresident)
	if err != nil{
		return err
	}
	return nil
}
//restituisce una squadra
func GetTeam(team entities.Team)(entities.Team, error){
	var squadra entities.Team
	db,err := config.GetDB()
	if err != nil{
		log.Fatal("Errore di connessione al database",err)
	}
	rowQuery,err := db.Query("SELECT * FROM squadre WHERE nome_squadra=(?)",team.TeamName)
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
	if squadra.TeamId == 0{
		return squadra,err
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
func RemoveTeam(team entities.Team)error{
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
	//elimino la squadra
	s,err := db.Exec("DELETE FROM squadre where team_Id=(?)",squadra.TeamId)
	if err != nil && err != sql.ErrNoRows{
		return err
	}
	//svincolo i giocatori che appartenevano alla squadra eliminata
	_,err = db.Exec("UPDATE giocatori SET team_Id = 0 where team_Id=(?)",squadra.TeamId)
	if err!=nil{
		log.Fatal(err)
	}
	/*RowsAffected() restituisce il numero di righe del database che sono state modificate
	o eliminate, in questo caso, se check == 0, vuol dire che la squadra non è presente nel database
	*/
	check, err := s.RowsAffected()
	if check == 0 {
		return errors.New("Squadra non trovata")
	}
	return nil
}
//Aggiunta di un giocatore in squadra
func AddPlayerToTeam(team entities.Team, player entities.Player)error{
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
	/*affichè possa aggiungere il giocatore alla squadra, il numero di giocatori totali in squadra <25, il costo del giocatore < crediti squadra
	e il giocatore dovrà avere un team_id = 0, ovvero svincolato*/
	if squadra.TeamNumPlayer<25 && squadra.TeamCredit>giocatore.PlayerPrice && giocatore.TeamId == 0{
		switch giocatore.PlayerRole {
		case "A":
			if squadra.TeamNumAtt<6{
				//aggiorno il team_id del giocatore
				_,err = db.Exec("UPDATE giocatori set team_Id = (?) where nome = (?)",squadra.TeamId,giocatore.PlayerName)
				if err!=nil{
					log.Fatal(err)
				}
				//ai crediti della squadra tolgo i credidi del giocatore, e incremento numero di giocatori e attaccanti presenti in squadra
				_, err = db.Exec("UPDATE squadre set crediti = (?), num_giocatori=(?), num_attaccanti=(?) where team_Id=(?)", squadra.TeamCredit-giocatore.PlayerPrice,squadra.TeamNumPlayer+1,squadra.TeamNumAtt+1,squadra.TeamId)
				if err!=nil{
					log.Fatal(err)
				}

			}else{
				err = errors.New("Troppi Attaccanti")
				log.Println(err)
			}
		case "C":
			if squadra.TeamNumCen<8{
				_,err = db.Exec("UPDATE giocatori set team_Id = (?) where nome = (?)",squadra.TeamId,giocatore.PlayerName)
				if err!=nil{
					log.Fatal(err)
				}
				_, err = db.Exec("UPDATE squadre set crediti = (?), num_giocatori=(?), num_centrocampisti=(?) where team_Id=(?)",squadra.TeamCredit-giocatore.PlayerPrice,squadra.TeamNumPlayer+1,squadra.TeamNumCen+1,squadra.TeamId)
				if err!=nil{
					log.Fatal(err)
				}

			}else{
				err = errors.New("Troppi Centrocampisti")
				log.Println(err)
			}
		case "D":
			if squadra.TeamNumDif<8{
				_,err = db.Exec("UPDATE giocatori set team_Id = (?) where nome = (?)",squadra.TeamId,giocatore.PlayerName)
				if err!=nil{
					log.Fatal(err)
				}
				_, err = db.Exec("UPDATE squadre set crediti = (?), num_giocatori=(?), num_difensori=(?) where team_Id=(?)",squadra.TeamCredit-giocatore.PlayerPrice,squadra.TeamNumPlayer+1,squadra.TeamNumDif+1,squadra.TeamId)
				if err!=nil{
					log.Fatal(err)
				}

			}else{
				err = errors.New("Troppi Difensori")
				log.Println(err)
			}

		case "P":
			if squadra.TeamNumPor<3{
				_,err = db.Exec("UPDATE giocatori set team_Id = (?) where nome = (?)",squadra.TeamId,giocatore.PlayerName)
				if err!=nil{
					log.Fatal(err)
				}
				_, err = db.Exec("UPDATE squadre set crediti = (?), num_giocatori=(?), num_portieri=(?) where team_Id=(?)",(squadra.TeamCredit-giocatore.PlayerPrice),squadra.TeamNumPlayer+1,squadra.TeamNumPor+1,squadra.TeamId)
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
func UpdatePlayer( player entities.Player)error{
	db, err := config.GetDB()
	if err != nil{
		log.Fatal("Errore di connessione al database",err)
	}
	//modifico la quotazione del giocatore
	_,err = db.Exec("UPDATE giocatori set quotazione=(?) where nome = (?)",player.PlayerPrice,player.PlayerName)
	if err!=nil{
		return err
	}
	return nil
}
//Rimozione di un giocatore in squadra
func RemovePlayerToTeam( team entities.Team, player entities.Player)error{
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
	/*l'aggiornamento consisterà nel settare a 0 il team_id del giocatore da rimuovere e aggiornare la squadra, incrementando il numero di crediti della metà del valore del giocatore
	e nel decrementare il numero di giocatori e il numero di giocatori nel ruolo del giocatore*/
	if giocatore.TeamId != 0{
		switch giocatore.PlayerRole {
		case "A":
			_,err = db.Exec("UPDATE giocatori set team_Id = 0 where nome = (?)",giocatore.PlayerName)
			if err!=nil{
				log.Fatal(err)
			}
			_,err = db.Exec("UPDATE squadre set crediti = (?), num_giocatori=(?), num_attaccanti=(?) where team_Id=(?)", squadra.TeamCredit+(giocatore.PlayerPrice/2),squadra.TeamNumPlayer-1,squadra.TeamNumAtt-1,squadra.TeamId)
			if err!=nil{
				log.Fatal(err)
			}
		case "C":
			_,err = db.Exec("UPDATE giocatori set team_Id = 0 where nome = (?)",giocatore.PlayerName)
			if err!=nil{
				log.Fatal(err)
			}
			_,err = db.Exec("UPDATE squadre set crediti = (?), num_giocatori=(?), num_centrocampisti=(?) where team_Id=(?)",squadra.TeamCredit+(giocatore.PlayerPrice/2),squadra.TeamNumPlayer-1,squadra.TeamNumCen-1,squadra.TeamId)
			if err!=nil{
				log.Fatal(err)
			}
		case "D":
			_,err = db.Exec("UPDATE giocatori set team_Id = 0 where nome = (?)",giocatore.PlayerName)
			if err!=nil{
				log.Fatal(err)
			}
			_,err = db.Exec("UPDATE squadre set crediti = (?), num_giocatori=(?), num_difensori=(?) where team_Id=(?)", squadra.TeamCredit+(giocatore.PlayerPrice/2),squadra.TeamNumPlayer-1,squadra.TeamNumDif-1,squadra.TeamId)
			if err!=nil {
				log.Fatal(err)
			}

		case "P":
			_,err = db.Exec("UPDATE giocatori set team_Id = 0 where nome = (?)",giocatore.PlayerName)
			if err!=nil{
				log.Fatal(err)
			}
			_,err = db.Exec("UPDATE squadre set crediti = (?), num_giocatori=(?), num_portieri=(?) where team_Id=(?)",squadra.TeamCredit+(giocatore.PlayerPrice/2),squadra.TeamNumPlayer-1,squadra.TeamNumPor-1,squadra.TeamId)
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
func ScambiaPlayer( team1 entities.Team, team2 entities.Team,player1 entities.Player,player2 entities.Player)error{
	//creo due vettori in cui inserire le due squadre e i due giocatori coinvolti nello scambio
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
	//affinchè possa essere effettuato lo scambio, deve esserci corrispondenza tra i team_id delle squadre e dei giocatori, e anche tra il ruolo dei giocatori
	if squadre[0].TeamId == giocatori[0].TeamId && squadre[1].TeamId == giocatori[1].TeamId && giocatori[0].PlayerRole==giocatori[1].PlayerRole{
		_,err = db.Exec("UPDATE giocatori set team_Id = (?) where nome = (?)",giocatori[1].TeamId,giocatori[0].PlayerName)
		if err!=nil{
			log.Fatal(err)
		}
		_,err = db.Exec("UPDATE giocatori set team_Id = (?) where nome = (?)",giocatori[0].TeamId,giocatori[1].PlayerName)
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







