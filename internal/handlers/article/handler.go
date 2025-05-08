package article

import (
	"errors"
	"fmt"
	"go-articles-manager-bot/internal/keyboards"
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

type ArticleId struct {
	DbId   uint32
	ListId uint16
}

func (ah *ArticleHandler) getArticlesList(tgId int64, page uint16, read bool) ([]*keyboards.ArticleListEl, error) {
	limit := ah.articlesPerPage

	offset := page*limit - limit

	articles, err := ah.articleRepo.GetArticlesByTgId(tgId, read, offset, limit)
	if err != nil {
		return nil, errors.New("Internal error")
	}

	if len(articles) == 0 {
		return nil, errors.New("No articles found")
	}

	var out []*keyboards.ArticleListEl

	for pos, a := range articles {
		out = append(
			out,
			&keyboards.ArticleListEl{
				DbId:   a.Id,
				ListId: uint16(pos) + 1 + uint16(page)*limit - limit,
				Read:   a.ReadAt != nil,
				Text:   a.GetTitleLink(),
			},
		)
	}
	return out, nil
}

func (ah *ArticleHandler) generateAtriclesMessage(articles []*keyboards.ArticleListEl) string {
	var out []string
	for _, a := range articles {
		var rText string
		if a.Read {
			rText = "(read)"
		}
		out = append(out, fmt.Sprintf("%d. %s %s", a.ListId, a.Text, rText))
	}
	return strings.Join(out, "\n")
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
