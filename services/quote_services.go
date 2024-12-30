package services

import (
    "quote-generator-backend/models"
    "quote-generator-backend/repositories"
)

type QuoteService struct {
    Repo *repositories.QuoteRepository
}
func (s *QuoteService) AddQuote(quote models.Quote) error {
    return s.Repo.AddQuote(quote)
}

func (s *QuoteService) GetQuotesByCategory(category string) ([]models.Quote, error) {
    return s.Repo.GetQuotesByCategory(category)
}

func (s *QuoteService) GetRandomQuotes(limit int, categories ...string) ([]models.Quote, error) {
    return s.Repo.GetRandomQuotes(limit, categories...)
}
