#include <iostream>
#include <stdlib.h>
#include <cpr/cpr.h>
#include "fantacalcio.h"

using namespace std;
string url = "http://localhost:7000";

FantaCalcio::FantaCalcio() {}
void FantaCalcio::addTeam() {
    string nome, nomePresidente;
    Team* squadra = new Team();
    cout << "Inserisci nome squadra" << endl;
    getline(cin, nome);
    squadra->setNome(nome);
    cout << "Inserisci nome presidente" << endl;
    getline(cin, nomePresidente);
    squadra->setPresidente(nomePresidente);
    auto r = cpr::Post(cpr::Url{ url + "/addTeam?nomeSquadra=" + squadra->getNome() + "&nomePresidente=" + squadra->getPresidente() });
    cout << "Returned Status:" << r.status_code << std::endl;
    cout << r.text << endl;
    delete squadra;

}
void FantaCalcio::getTeam() {
    string nome;
    Team* squadra = new Team();
    cout << "Inserisci il nome della squadra:" << endl;
    getline(cin, nome);
    squadra->setNome(nome);
    auto r = cpr::Get(cpr::Url{ url + "/getTeam?nome=" + squadra->getNome() });
    cout << r.text << endl;
    cout << "Returned Status:" << r.status_code << std::endl;
    delete squadra;
}
void FantaCalcio::removeTeam() {
    string nome;
    Team* squadra = new Team();
    cout << "Inserisci nome della squadra da eliminare" << endl;
    getline(cin, nome);
    squadra->setNome(nome);
    auto r = cpr::Delete(cpr::Url{ url + "/deleteTeam?nome=" + squadra->getNome() });
    cout << "\nReturned Status:" << r.status_code << std::endl;
    cout << r.text << endl;
    delete squadra;
}
void FantaCalcio::addPlayerToTeam() {
    string nomeSquadra, nomePlayer;
    Team* squadra = new Team();
    Player* player = new Player();
    cout << "Inserisci nome squadra" << endl;
    getline(cin, nomeSquadra);
    squadra->setNome(nomeSquadra);
    cout << "Inserisci nome giocatore" << endl;
    getline(cin, nomePlayer);
    player->setNome(nomePlayer);
    cout << "Stai inserendo il giocatore: " << player->getNome() << " nella squadra: " << squadra->getNome() << endl;
    auto r = cpr::Put(cpr::Url{ url + "/addPlayer?nomeSquadra=" + squadra->getNome() + "&nomePlayer=" + player->getNome() });
    cout << "Returned Status:" << r.status_code << std::endl;
    cout << r.text << endl;
    delete squadra;
    delete player;
}
void FantaCalcio::putPlayer() {
    string nome;
    int price;
    Player* player = new Player();
    cout << "Inserisci il nome del giocatore di cui aggiornare la quotazione" << endl;
    getline(cin, nome);
    player->setNome(nome);
    cout << "Inserisci la nuova quotazione del giocatore" << endl;
    cin >> price;
    player->setQuotazione(price);
    auto r = cpr::Put(cpr::Url{ url + "/putPlayer?nomePlayer=" + player->getNome() + "&quotazione=" + to_string(player->getQuotazione()) });
    cout << "Returned Status:" << r.status_code << std::endl;
    cout << r.text << endl;
    delete player;
}
void FantaCalcio::toFreePlayer() {
    Team* squadra = new Team();
    Player* giocatore = new Player();
    string nomeSquadra, nomeGiocatore;
    cout << "Inserisci il nome della squadra in cui si trova il giocatore " << endl;
    getline(cin, nomeSquadra);
    squadra->setNome(nomeSquadra);
    cout << "Inserisci il nome del giocatore da rimuovere dalla squadra" << endl;
    getline(cin, nomeGiocatore);
    giocatore->setNome(nomeGiocatore);
    auto r = cpr::Put(cpr::Url{ url + "/removePlayer?nomeSquadra=" + squadra->getNome() + "&nomePlayer=" + giocatore->getNome() });
    cout << "Returned Status:" << r.status_code << std::endl;
    cout << r.text << endl;
    delete squadra;
    delete giocatore;
}
void FantaCalcio::transferPlayers() {
    string squadra1, squadra2, player1, player2;
    Team* t1 = new Team();
    Team* t2 = new Team();
    Player* p1 = new Player();
    Player* p2 = new Player();
    cout << "Inserisci nome squadra 1:" << endl;
    getline(cin, squadra1);
    t1->setNome(squadra1);
    cout << "Inserisci nome squadra 2:" << endl;
    getline(cin, squadra2);
    t2->setNome(squadra2);
    cout << "Inserisci giocatore della squadra 1:" << endl;
    getline(cin, player1);
    p1->setNome(player1);
    cout << "Inserisci giocatore della squadra 2:" << endl;
    getline(cin, player2);
    p2->setNome(player2);
    auto res = cpr::Put(cpr::Url{ url + "/changePlayer?squadra1=" + t1->getNome() + "&squadra2=" + t2->getNome() + "&player1=" + p1->getNome() + "&player2=" + p2->getNome() });
    cout << "Returned Status:" << res.status_code << std::endl;
    cout << res.text << endl;
    delete t1;
    delete t2;
    delete p1;
    delete p2;
}
int FantaCalcio::menu() {
    string msg = string("\n\nScegli una delle seguenti opzioni:\n") +
        "1) Inserisci una nuova squadra\n" +
        "2) Visualizza una squadra\n" +
        "3) Rimuovi una squadra\n" +
        "4) Aggiungi un giocatore in squadra\n" +
        "5) Modifica la quotazione di un giocatore\n" +
        "6) Svincola un giocatore\n" +
        "7) Effettua uno scambio tra due squadre\n" +
        "0) Uscita\n" +
        "La tua scelta: ";
    string scelta = "";

    while (true) {
        cout << msg << endl;
        getline(cin, scelta);
        switch (scelta[0]) {
        case '1': addTeam(); break;
        case '2': getTeam(); break;
        case '3': removeTeam(); break;
        case '4': addPlayerToTeam(); break;
        case '5': putPlayer(); break;
        case '6': toFreePlayer(); break;
        case '7': transferPlayers(); break;
        case '0': return (0);

        }
    }
}