package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	//string de conexao com o bnco MySql
	DBConnString = ""

	//Porta onde API está executando
	Port = 0
)

// Função que trabalha com variáveis de ambiente
func Load() {

	var err error

	if err = godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	Port, err = strconv.Atoi(os.Getenv("API_PORT"))
	if err != nil {
		Port = 8000
	}

	DBConnString = fmt.Sprintf("%s:%s/@%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_NAME"),
	)

}
