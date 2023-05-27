package urlsrv

import (
	"errors"

	"github.com/jannawro/hex-url-shortener/internal/core/domain"
	"github.com/jannawro/hex-url-shortener/internal/core/ports"
	hashids "github.com/speps/go-hashids"
)

var (
    ErrFailedRepoGet = errors.New("get url from repository had failed: ")
    ErrFailedEncodingUrl = errors.New("encoding the url had failed: ")
    ErrFailedRepoIncrement = errors.New("increment sequential ID in repository had failed: ")
    ErrFailedRepoSave = errors.New("saving url to repoistory had failed: ")
    ErrFailedRepoDump = errors.New("get all urls from repository had failed: ")
)

type service struct {
    urlRepository ports.UrlRepository
}

func New(urlRepository ports.UrlRepository) *service {
    return &service{
        urlRepository: urlRepository,
    }
}

func (srv *service) Get(shorthand string) (domain.Url, error) {
    url, err := srv.urlRepository.Get(shorthand)
    if err != nil {
        return domain.Url{}, errors.Join(ErrFailedRepoGet, err)
    }

    return url, nil
}

func (srv *service) GetAll() ([]domain.Url, error) {
    urls, err := srv.urlRepository.Dump()
    if err != nil {
        return []domain.Url{}, errors.Join(ErrFailedRepoDump, err)
    }

    return urls, nil
}

func (srv *service) Create(host, url string) (domain.Url, error) {
    seq, err := srv.urlRepository.IncrementSeq()
    if err != nil {
        return domain.Url{}, errors.Join(ErrFailedRepoIncrement, err)
    }

    uniqueID, err := generateHash(seq)
    if err != nil {
        return domain.Url{}, errors.Join(ErrFailedEncodingUrl, err)
    }

    saved, err := srv.urlRepository.Save(domain.Url{
        Original: url,
        Shorthand: host + "/" + uniqueID,
    })

    if err != nil {
        return domain.Url{}, errors.Join(ErrFailedRepoSave, err)
    }

    return saved, nil
}

func generateHash(uniqueInt int) (string, error) {
    hd := hashids.NewData()
    hd.Salt = "hex arch exercise"
    hd.MinLength = 6

    h, err := hashids.NewWithData(hd);
    if err != nil {
        return "", err
    }

    hash, err := h.Encode([]int{uniqueInt})
    if err != nil {
        return "", err
    }

    return hash, nil
}
