#setting della directory in cui si trova il file excel
setwd("D:/ProgettoAPL")
#caricamento del listone excel
library(readxl)
giocatori <- read_excel("Quotazioni_Fantacalcio.xlsx")
#Importo libreria RMySql
library(RMySQL)
#Connessione
my = dbConnect(MySQL(), username = 'root', password = 'a1b2c3d4e5')
#Creazione database
dbSendQuery(my, "CREATE DATABASE IF NOT EXISTS fantacalcio")
dbSendQuery(my, "USE fantacalcio")
dbSendQuery(my, "SET GLOBAL local_infile = true;")
dbWriteTable(my, "giocatori", giocatori, overwrite = TRUE, temporary = FALSE)
dbSendQuery(my, "ALTER TABLE giocatori DROP COLUMN row_names")
dbSendQuery(my, "ALTER TABLE giocatori 
            CHANGE COLUMN `id` `id` DOUBLE NOT NULL ,
            CHANGE COLUMN `ruolo` `ruolo` TEXT NOT NULL ,
            CHANGE COLUMN `nome` `nome` TEXT NOT NULL ,
            CHANGE COLUMN `squadra` `squadra` TEXT NOT NULL ,
            CHANGE COLUMN `quotazione` `quotazione` DOUBLE NOT NULL ,
            ADD PRIMARY KEY (`id`);")
dbSendQuery(my,"UPDATE giocatori set team_Id = 0;")
dbSendQuery(my, "CREATE TABLE IF NOT EXISTS squadre(
            team_Id int NOT NULL AUTO_INCREMENT UNIQUE,
            nome_squadra varchar(50) NOT NULL UNIQUE,
            crediti int default 300,
            nome_presidente varchar(50) NOT NULL UNIQUE,
            num_giocatori int,
            num_portieri int,
            num_difensori int,
            num_centrocampisti int,
            num_attaccanti int,
            PRIMARY KEY (team_Id));")
dbListTables(my)