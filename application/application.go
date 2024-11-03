package application

import (
	"context"

	"github.com/SoftwareArch-BackstreetBoys/club-service/model"
	"github.com/SoftwareArch-BackstreetBoys/club-service/repository"
	"github.com/google/uuid"
)

type Application interface {
	CreateClub(ctx context.Context, club model.Club) (*model.Club, error)
	JoinClub(ctx context.Context, clubID string, userID string) error
	LeaveClub(ctx context.Context, clubID string, userID string) error
	GetClubInfo(ctx context.Context, clubID string) (*model.Club, error)
	GetJoinedClub(ctx context.Context, userID string) ([]model.Club, error)
	IsBelongToClub(ctx context.Context, clubID string, userID string) (bool, error)
	SearchClubs(ctx context.Context, keyword string) ([]model.Club, error)
	GetAllClubs(ctx context.Context) ([]model.Club, error)
	CheckDatabaseConnection(ctx context.Context) error
}

type application struct {
	repository repository.Repository
}

func NewApplication(repository repository.Repository) Application {
	return &application{
		repository: repository,
	}
}

func (self *application) CreateClub(ctx context.Context, club model.Club) (*model.Club, error) {
	club.ID = uuid.NewString()

	return self.repository.CreateClub(ctx, club)
}

func (self *application) JoinClub(ctx context.Context, clubID string, userID string) error {
	return self.repository.CreateClubMemberShip(ctx, clubID, userID)
}

func (self *application) LeaveClub(ctx context.Context, clubID string, userID string) error {
	return self.repository.DeleteClubMemberShip(ctx, clubID, userID)
}

func (self *application) GetClubInfo(ctx context.Context, clubID string) (*model.Club, error) {
	return self.repository.GetClub(ctx, clubID)
}

func (self *application) GetJoinedClub(ctx context.Context, userID string) ([]model.Club, error) {
	joinedClubIDS, err := self.repository.GetJoinedClubIDS(ctx, userID)

	if err != nil {
		return nil, err
	}

	return self.repository.GetClubs(ctx, joinedClubIDS)
}

func (self *application) IsBelongToClub(ctx context.Context, clubID string, userID string) (bool, error) {
	membership, err := self.repository.GetClubMemberShip(ctx, clubID, userID)

	if err != nil {
		return false, err
	}

	if membership == nil {
		return false, nil
	}

	return true, nil
}

func (self *application) SearchClubs(ctx context.Context, keyword string) ([]model.Club, error) {
	return self.repository.SearchClubs(ctx, keyword)
}

func (self *application) GetAllClubs(ctx context.Context) ([]model.Club, error) {
	return self.repository.GetAllClubs(ctx)
}

func (self *application) CheckDatabaseConnection(ctx context.Context) error {
	return self.repository.PingDatabase(ctx)
}