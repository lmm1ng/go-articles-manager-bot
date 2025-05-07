package article

import (
	"errors"
	"fmt"
	"strings"
)

type ArticleHandler struct {
	articleRepo     articleRepository
	userRepo        userRepository
	articlesPerPage uint16
}

func New(articleRepo articleRepository, userRepo userRepository) *ArticleHandler {
	return &ArticleHandler{
		articleRepo:     articleRepo,
		userRepo:        userRepo,
		articlesPerPage: 5,
	}
}

func (ah *ArticleHandler) generateArticlesList(tgId int64, page uint16, read bool) ([]string, error) {
	limit := ah.articlesPerPage

	offset := page*limit - limit

	articles, err := ah.articleRepo.GetArticlesByTgId(tgId, read, offset, limit)
	if err != nil {
		return nil, errors.New("Internal error")
	}

	if len(articles) == 0 {
		return nil, errors.New("No articles found")
	}

	var aSlice []string

	for pos, a := range articles {
		text := a.GetTitleLink()

		if a.ReadAt != nil {
			text += " (read)"
		}

		aSlice = append(aSlice, fmt.Sprintf("%d. %s", pos+1+int(page*limit-limit), text))
	}
	return aSlice, nil
}

func (ah *ArticleHandler) getCallbackArgs(cbStr string, buttons ...string) []string {
	argsStr := strings.Clone(cbStr)

	for _, b := range buttons {
		s, found := strings.CutPrefix(argsStr, b+" ")
		if !found {
			continue
		}
		argsStr = s
	}

	return strings.Split(argsStr, " ")
}
