package handlers

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/cifra-city/comtools/cifractx"
	"github.com/cifra-city/comtools/httpkit"
	"github.com/cifra-city/comtools/httpkit/problems"
	"github.com/cifra-city/users-storage/internal/config"
	"github.com/cifra-city/users-storage/internal/data/db/sqlcore"
	"github.com/cifra-city/users-storage/resources"
	"github.com/sirupsen/logrus"
)

// SearchUsers выполняет поиск пользователей по параметру "q".
func SearchUsers(w http.ResponseWriter, r *http.Request) {
	server, err := cifractx.GetValue[*config.Service](r.Context(), config.SERVER)
	if err != nil {
		logrus.Errorf("Failed to retrieve service configuration: %v", err)
		httpkit.RenderErr(w, problems.InternalError("failed to retrieve service configuration"))
		return
	}

	log := server.Logger

	query := strings.TrimSpace(r.URL.Query().Get("q"))
	pageStr := r.URL.Query().Get("page")
	limitStr := r.URL.Query().Get("limit")

	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit < 1 || limit > 100 {
		limit = 10
	}

	offset := (page - 1) * limit

	if query == "" {
		log.Warn("empty search query")
		httpkit.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	users, err := server.Databaser.Users.Search(r.Context(), &query, limit, offset)
	if err != nil {
		log.Errorf("failed to search users: %v", err)
		httpkit.RenderErr(w, problems.InternalError())
		return
	}

	response := NewSearchResponse(users, resources.UserSearchType)
	httpkit.Render(w, response)
}

func NewSearchResponse(users []sqlcore.User, typeOperation string) resources.UserCollection {
	var result []resources.UserData
	for _, user := range users {
		result = append(result, resources.UserData{
			Type: typeOperation,
			Attributes: resources.UserDataAttributes{
				Id:       user.ID.String(),
				Username: user.Username,
				Title:    user.Title.String,
				Status:   user.Status.String,
				Avatar:   user.Avatar.String,
				Bio:      user.Bio.String,
				City:     user.City.UUID.String(),
			},
		})
	}

	return resources.UserCollection{Data: result}
}
