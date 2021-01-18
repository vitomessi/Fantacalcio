#include "classes.h"

Team::Team() {}
Team::~Team() {}
void Team::setNome(string& name) {
	Team::nome = name;
}
void Team::setPresidente(string& president) {
	Team::presidente = president;
}
void Team::setId(int& i) {
	Team::id = i;
}
Player::Player() {}
Player::~Player() {}
void Player::setNome(string& name) {
	Player::nome = name;
}
void Player::setQuotazione(int& newPrice) {
	Player::quotazione = newPrice;
}