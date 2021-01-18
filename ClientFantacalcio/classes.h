#pragma once
#ifndef CLIENTFANTACALCIO_CLASSES_H
#define CLIENTFANTACALCIO_CLASSES_H

#include <iostream>

using namespace std;
class Team
{
private:
    string nome, presidente;
    int id;
public:
    Team();
    ~Team();
    void setNome(string& name);
    string getNome() { return nome; };
    void setPresidente(string& president);
    string getPresidente() { return presidente; };
    void setId(int& i);
    int getId() { return id; };
};
class Player
{
private:
    string nome;
    int quotazione;
public:
    Player();
    ~Player();
    void setNome(string& name);
    string getNome() { return nome; };
    void setQuotazione(int& newPrice);
    int getQuotazione() { return quotazione; };
};
#endif //CLIENTFANTACALCIO_CLASSES_H