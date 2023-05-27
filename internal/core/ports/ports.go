package ports

import "github.com/jannawro/hex-url-shortener/internal/core/domain"

type UrlRepository interface {
    Get(shorthand string) (domain.Url, error)
    Save(url domain.Url) (domain.Url, error)
    IncrementSeq() (int, error) 
    Dump() ([]domain.Url, error)
} 

type UrlService interface {
    Get(shorthand string) (domain.Url, error)
    Create(domain, url string) (domain.Url, error)
    GetAll() ([]domain.Url, error)
}
