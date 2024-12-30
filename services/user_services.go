package services

import (
    "go.mongodb.org/mongo-driver/bson/primitive"
    "quote-generator-backend/models"
    "quote-generator-backend/repositories"
)

type UserService struct {
    Repo *repositories.UserRepository
}

func (s *UserService) AddFavorite(userID primitive.ObjectID, quoteID primitive.ObjectID) error {
    return s.Repo.AddFavorite(userID, quoteID)
}

func (s *UserService) GetFavorites(userID primitive.ObjectID) ([]models.Quote, error) {
    return s.Repo.GetFavorites(userID)
}

func (s *UserService) LoginOrCreate(username string) (primitive.ObjectID, error) {
    return s.Repo.LoginOrCreate(username)
}