package memorykv

import (
	"errors"
	"sync"

	"github.com/jannawro/hex-url-shortener/internal/core/domain"
)

var (
    ErrDoesNotExist = errors.New("No url for this key was found.")
)

type store struct {
    increment int
    kvStore map[string]string
    mu sync.Mutex
}

func New() *store {
    return &store{
        increment: 0,
        kvStore: make(map[string]string),
    }
}

func(s *store) Get(shorthand string) (domain.Url, error) {
    value := s.kvStore[shorthand]
    if value == "" {
        return domain.Url{}, ErrDoesNotExist
    }

    return domain.Url{
        Shorthand: shorthand,
        Original: value,
    }, nil
}

func(s *store) Save(url domain.Url) (domain.Url, error) {
    s.kvStore[url.Shorthand] = url.Original
    return url, nil
}

func(s *store) IncrementSeq() (int, error) {
    s.mu.Lock()
    defer s.mu.Unlock()
    s.increment++
    return s.increment, nil
}

func(s *store) Dump() ([]domain.Url, error) {
    urls := make([]domain.Url, 0, len(s.kvStore))
    for key, value := range s.kvStore {
        urls = append(urls, domain.Url{
            Shorthand: key,
            Original: value,
        })
    }

    return urls, nil
}
