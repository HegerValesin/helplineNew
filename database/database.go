package database

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"helplineBKgo/models"
	"log"
	"os"
)

var DB *gorm.DB

func InitDB() {
	// Carrega as variáveis de ambiente do arquivo .env
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Erro ao carregar o arquivo .env: %v", err)
	}

	// Configura a string de conexão com o PostgreSQL
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s sslmode=%s password=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_SSLMODE"),
		os.Getenv("DB_PASSWORD"),
	)

	// Abre a conexão com o banco de dados
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Erro ao conectar ao banco de dados: %v", err)
	}

	// Migração automática das tabelas
	err = DB.AutoMigrate(
		&models.User{},
		&models.Service{},
		&models.Sector{},
		&models.Facility{},
		&models.Floor{},
		&models.Room{},
		&models.Occurrence{},
		&models.OccurrenceAudio{},
	)
	if err != nil {
		log.Fatalf("Erro ao migrar tabelas: %v", err)
	}

	log.Println("Conexão com o banco de dados estabelecida e migração concluída.")
}
