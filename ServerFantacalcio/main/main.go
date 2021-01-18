package main

import (
	"ServerFantacalcio/config"
	"ServerFantacalcio/entities"
	"ServerFantacalcio/team_player"
	"ServerFantacalcio/utils"
	"context"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"strconv"
)
var db,e = config.GetDB()
func main(){
	if e!=nil{
		log.Fatal(e)
	}
	defer db.Close()

	eb := db.Ping()
	if eb!=nil{
		panic(eb.Error())
	}
	/*
		1)Inserisci nuova squadra
		2)Restituisci una squadra
		3)Rimuovi una squadra
		4)Aggiungi un giocatore ad una squadra
		5)Modifica quotazione giocatore
		6)Svincola un giocatore
		7)Scambia giocatori
	*/
	http.HandleFunc("/addTeam",AddTeam)
	http.HandleFunc("/getTeam",GetTeam)
	http.HandleFunc("/deleteTeam",DeleteTeam)
	http.HandleFunc("/addPlayer",AddPlayer)
	http.HandleFunc("/putPlayer",UpdatePlayerPrice)
	http.HandleFunc("/removePlayer",RemovePlayer)
	http.HandleFunc("/changePlayer",ChangePlayer)

	//Metodi non implementati lato Client
	http.HandleFunc("/getTeams",GetTeams)
	http.HandleFunc("/getPlayer",GetPlayer)
	http.HandleFunc("/getPlayers",GetPlayers)


	err := http.ListenAndServe(":7000",nil)
	if err!=nil{
		log.Fatal(err)
	}
}
/*
	1)Inserisci nuova squadra
	2)Restituisci una squadra
	3)Rimuovi una squadra
	4)Aggiungi un giocatore ad una squadra
	5)Modifica quotazione giocatore
	6)Svincola un giocatore
	7)Scambia giocatori
*/
//Create Team
func AddTeam(w http.ResponseWriter, r *http.Request){
	if r.Method == "POST"{
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		//definisco una variabile di tipo Team
		var team entities.Team
		nameTeam := r.URL.Query().Get("nomeSquadra")
		if nameTeam == ""{
			utils.ResponseJSON(w, "campo nome squadra non deve essere vuoto",http.StatusBadRequest)
			return
		}
		team.TeamName = nameTeam
		namePresident := r.URL.Query().Get("nomePresidente")
		if namePresident == ""{
			utils.ResponseJSON(w, "campo nome presidente non deve essere vuoto",http.StatusBadRequest)
			return
		}
		team.TeamPresident = namePresident
		if err := team_player.InsertTeam(ctx,team); err != nil{
			utils.ResponseJSON(w, err, http.StatusBadRequest)
			return
		}
		res := map[string]string{
			"Status":"Succesfully",
		}
		utils.ResponseJSON(w,res,http.StatusCreated)
		return
	}
	http.Error(w, "comando non consentito", http.StatusMethodNotAllowed)
	return
}
//Get team
func GetTeam(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		var team entities.Team
		name := r.URL.Query().Get("nome")
		if name == ""{
			utils.ResponseJSON(w, "campo nome non deve essere vuoto",http.StatusBadRequest)
			return
		}
		team.TeamName = name
		team, err := team_player.GetTeam(ctx,team)
		if err != nil{
			utils.ResponseJSON(w, "Squadra non trovata",http.StatusNotFound)
			return
		}
		utils.ResponseJSON(w, team, http.StatusOK)
		return
	}
	http.Error(w, "Comando non consentito",http.StatusMethodNotAllowed)
	return

}
//Delete Team
func DeleteTeam(w http.ResponseWriter, r *http.Request){
	if r.Method == "DELETE" {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		var team entities.Team
		name := r.URL.Query().Get("nome")
		if name == ""{
			utils.ResponseJSON(w, "campo nome non deve essere vuoto",http.StatusBadRequest)
			return
		}
		team.TeamName = name
		//team.TeamId,_ = strconv.Atoi(id)
		if err := team_player.RemoveTeam(ctx,team); err != nil{
			errore := map[string]string{
				"error":fmt.Sprintf("%v",err),
			}

			utils.ResponseJSON(w,errore,http.StatusNotFound)
			return
		}
		res := map[string]string{
			"Status":"Squadra eliminata con successo",
		}
		utils.ResponseJSON(w,res,http.StatusOK)
		return
	}
	http.Error(w, "comando non consentito", http.StatusMethodNotAllowed)
	return
}
//Add player in a team
func AddPlayer(w http.ResponseWriter, r *http.Request){
	if r.Method == "PUT"{
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		var player entities.Player
		var team entities.Team
		nameTeam := r.URL.Query().Get("nomeSquadra")
		if nameTeam == ""{
			utils.ResponseJSON(w, "campo nome squadra non deve essere vuoto",http.StatusBadRequest)
			return
		}
		team.TeamName = nameTeam
		namePlayer := r.URL.Query().Get("nomePlayer")
		if namePlayer == ""{
			utils.ResponseJSON(w, "campo nome giocatore non deve essere vuoto",http.StatusBadRequest)
			return
		}
		player.PlayerName = namePlayer
		if err := team_player.AddPlayerToTeam(ctx,team,player); err != nil{
			errore := map[string]string{
				"error":fmt.Sprintf("%v",err),
			}
			utils.ResponseJSON(w,errore,http.StatusBadRequest)
			return
		}
		res := map[string]string{
			"Status":"Giocatore aggiunto con successo",
		}

		utils.ResponseJSON(w,res,http.StatusOK)
		return

	}
	http.Error(w, "comando non consentito", http.StatusMethodNotAllowed)
	return

}
//Modifica la quotazione del giocatore
func UpdatePlayerPrice(w http.ResponseWriter, r *http.Request){
	if r.Method == "PUT"{

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		var player entities.Player
		namePlayer := r.URL.Query().Get("nomePlayer")
		if namePlayer == ""{
			utils.ResponseJSON(w, "campo nome giocatore non deve essere vuoto",http.StatusBadRequest)
			return
		}
		player.PlayerName = namePlayer
		costPlayer := r.URL.Query().Get("quotazione")
		if namePlayer == ""{
			utils.ResponseJSON(w, "campo nuova quotazione non deve essere vuoto",http.StatusBadRequest)
			return
		}
		player.PlayerPrice,_ = strconv.Atoi(costPlayer)
		if err := team_player.UpdatePlayer(ctx,player);err!=nil{
			utils.ResponseJSON(w,err,http.StatusInternalServerError)
			return
		}
		res := map[string]string{
			"status":"Succesfully",
		}
		utils.ResponseJSON(w,res,http.StatusOK)
		return
	}
	http.Error(w,"Comando non trovato",http.StatusMethodNotAllowed)
	return
}
//Rimuove un giocatore dalla squadra
func RemovePlayer(w http.ResponseWriter, r *http.Request){
	if r.Method == "PUT"{

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		var player entities.Player
		var team entities.Team
		nameTeam := r.URL.Query().Get("nomeSquadra")
		if nameTeam == ""{
			utils.ResponseJSON(w, "campo nome della squadra non deve essere vuoto",http.StatusBadRequest)
			return
		}
		team.TeamName = nameTeam
		namePlayer := r.URL.Query().Get("nomePlayer")
		if namePlayer == ""{
			utils.ResponseJSON(w, "campo nome giocatore non deve essere vuoto",http.StatusBadRequest)
			return
		}
		player.PlayerName = namePlayer
		if err := team_player.RemovePlayerToTeam(ctx,team,player); err != nil{
			utils.ResponseJSON(w, "Operazione non consentita", http.StatusBadRequest)
			return
		}
		res := map[string]string{
			"Status":"Succesfully",
		}

		utils.ResponseJSON(w,res,http.StatusOK)
		return

	}
	http.Error(w, "comando non consentito", http.StatusMethodNotAllowed)
	return

}
//Scambia giocatori tra due squadre
func ChangePlayer(w http.ResponseWriter, r *http.Request){
	if r.Method == "PUT"{
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		var player1 entities.Player
		var player2 entities.Player
		var team1 entities.Team
		var team2 entities.Team

		name1 := r.URL.Query().Get("squadra1")
		if name1 == ""{
			utils.ResponseJSON(w, "campo nome1 non deve essere vuoto",http.StatusBadRequest)
			return
		}
		team1.TeamName = name1
		name2 := r.URL.Query().Get("squadra2")
		if name2 == ""{
			utils.ResponseJSON(w, "campo nome2 non deve essere vuoto",http.StatusBadRequest)
			return
		}
		team2.TeamName = name2
		nameP1 := r.URL.Query().Get("player1")
		if nameP1 == ""{
			utils.ResponseJSON(w, "campo nomeP1 non deve essere vuoto",http.StatusBadRequest)
			return
		}
		player1.PlayerName = nameP1
		nameP2 := r.URL.Query().Get("player2")
		if nameP2 == ""{
			utils.ResponseJSON(w, "campo nomeP2 non deve essere vuoto",http.StatusBadRequest)
			return
		}
		player2.PlayerName = nameP2
		if err := team_player.ScambiaPlayer(ctx,team1,team2,player1,player2); err != nil{
			utils.ResponseJSON(w, "Operazione non consentita", http.StatusBadRequest)
			return
		}
		res := map[string]string{
			"Status":"Succesfully",
		}

		utils.ResponseJSON(w,res,http.StatusOK)
		return

	}
	http.Error(w, "comando non consentito", http.StatusMethodNotAllowed)
	return
}
/*
Metodi non implementati lato client
*/
//Get all teams
func GetTeams(w http.ResponseWriter, r *http.Request){
	if r.Method == "GET" {
		ctx, cancel := context.WithCancel(context.Background())

		defer cancel()

		team, err := team_player.GetTeams(ctx)
		if err != nil{
			fmt.Println(err)
		}
		utils.ResponseJSON(w, team, http.StatusOK)
		return
	}
	http.Error(w, "Comando non consentito",http.StatusNotFound)
	return
}
//Get player
func GetPlayer(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		var player entities.Player
		name := r.URL.Query().Get("nome")
		if name == ""{
			utils.ResponseJSON(w, "campo nome non deve essere vuoto",http.StatusBadRequest)
			return
		}
		player.PlayerName = name
		player, err := team_player.GetPlayer(ctx,player)
		if err != nil{
			fmt.Println(err)
		}
		utils.ResponseJSON(w, player, http.StatusOK)
		return
	}
	http.Error(w, "Comando non consentito",http.StatusNotFound)
	return

}
//Get all free players
func GetPlayers(w http.ResponseWriter, r *http.Request){
	if r.Method == "GET" {
		ctx, cancel := context.WithCancel(context.Background())

		defer cancel()

		team, err := team_player.GetPlayers(ctx)
		if err != nil{
			fmt.Println(err)
		}
		utils.ResponseJSON(w, team, http.StatusOK)
		return
	}
	http.Error(w, "Comando non consentito",http.StatusNotFound)
	return
}



