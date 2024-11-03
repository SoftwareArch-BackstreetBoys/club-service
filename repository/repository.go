package repository

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/SoftwareArch-BackstreetBoys/club-service/model"
	"github.com/joho/godotenv"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MONGO_URI string
var MONGO_DATABASE string
var CLUB_COLLECTION_NAME string
var CLUB_MEMBERSHIP_COLLECTION_NAME string

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	MONGO_URI = os.Getenv("MONGO_URI")
	MONGO_DATABASE = os.Getenv("MONGO_DATABASE")
	CLUB_COLLECTION_NAME = os.Getenv("CLUB_COLLECTION_NAME")
	CLUB_MEMBERSHIP_COLLECTION_NAME = os.Getenv("CLUB_MEMBERSHIP_COLLECTION_NAME")
}

type Repository interface {
	// club table
	CreateClub(ctx context.Context, club model.Club) (*model.Club, error)
	GetAllClubs(ctx context.Context) ([]model.Club, error)
	GetClub(ctx context.Context, clubID string) (*model.Club, error)
	GetClubs(ctx context.Context, clubIDs []string) ([]model.Club, error)
	SearchClubs(ctx context.Context, keyword string) ([]model.Club, error)

	// club_membership table
	CreateClubMemberShip(ctx context.Context, clubID string, userID string) error
	GetClubMemberShip(ctx context.Context, clubID string, userID string) (*model.ClubMembership, error)
	DeleteClubMemberShip(ctx context.Context, clubID string, userID string) error
	GetJoinedClubIDS(ctx context.Context, userID string) ([]string, error)

	// health check
	PingDatabase(ctx context.Context) error
}

type repository struct {
	mongoDB *mongo.Database
}

func NewRepository(ctx context.Context) Repository {
	mongoDB := connectMongoDB(MONGO_URI)

	repository := &repository{
		mongoDB: mongoDB,
	}

	if err := repository.createClubIndexes(ctx); err != nil {
		panic(err.Error())
	}

	if err := repository.createClubMembershipIndexes(ctx); err != nil {
		panic(err.Error())
	}

	return repository
}

func connectMongoDB(mongoURI string) *mongo.Database {
	// Create a context with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel() // Ensure cancel is called to avoid context leak

	// Connect to MongoDB using the mongo.Connect method
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatal(err)
	}

	// Ping the database to verify connection
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB")
	return client.Database(MONGO_DATABASE)
}

func (r *repository) PingDatabase(ctx context.Context) error {
	return r.mongoDB.Client().Ping(ctx, nil)
}