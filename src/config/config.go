package config

import (
	"log"
	"os"
	"strconv"

	"github.com/diamondburned/arikawa/v3/discord"
	"github.com/joho/godotenv"
)

var (
	CanvasWidth = 300
	CanvasHeigh = 300
)

var (
	Token,
	ClientID,
	MongoDBURI,
	DatabaseName,
	ImageURL string
	ErrorLogChannel discord.ChannelID = 0
)

func init() {
	path, err := os.Getwd()
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	if err := godotenv.Load(path + "/.env"); err != nil {
		log.Fatalf("Error: %v", err)
	}
}

func init() {
	Token = os.Getenv("TOKEN")
	MongoDBURI = os.Getenv("MONGO_URI")
	DatabaseName = os.Getenv("DATABASE_NAME")
	ClientID = os.Getenv("CLIENT_ID")
	ImageURL = os.Getenv("IMAGE_URL")

	if i, err := strconv.ParseUint(os.Getenv("ERROR_LOG_CHANNEL"), 10, 64); err == nil {
		ErrorLogChannel = discord.ChannelID(i)
	}
}

func GetInviteURL() string {
	return "https://discord.com/api/oauth2/authorize?client_id=" + ClientID + "&permissions=414464724032&scope=bot%20applications.commands"
}
