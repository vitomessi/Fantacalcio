#pragma once
#ifndef CLIENTFANTACALCIO_FANTACALCIO_H
#define CLIENTFANTACALCIO_FANTACALCIO_H

#include "classes.h"
class FantaCalcio {
public:
	FantaCalcio();
	void addTeam();
	void getTeam();
	void removeTeam();
	void addPlayerToTeam();
	void putPlayer();
	void toFreePlayer();
	void transferPlayers();
	int menu();
};
#endif //CLIENTFANTACALCIO_FANTACALCIO_H
