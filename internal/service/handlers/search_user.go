package handlers

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/cifra-city/cifractx"
	"github.com/cifra-city/httpkit"
	"github.com/cifra-city/httpkit/problems"
	"github.com/cifra-city/users-storage/internal/config"
	"github.com/cifra-city/users-storage/internal/data/db/dbcore"
	"github.com/cifra-city/users-storage/resources"
)

// SearchUsers выполняет поиск пользователей по параметру "q".
func SearchUsers(w http.ResponseWriter, r *http.Request) {
	service, err := cifractx.GetValue[*config.Service](r.Context(), config.SERVICE)
	if err != nil {
		httpkit.RenderErr(w, problems.InternalError("failed to retrieve service configuration"))
		return
	}

	log := service.Logger

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

	users, err := service.Databaser.Users.Search(r, &query, limit, offset)
	if err != nil {
		log.Errorf("failed to search users: %v", err)
		httpkit.RenderErr(w, problems.InternalError())
		return
	}

	response := NewSearchResponse(users)
	httpkit.Render(w, response)
}

func NewSearchResponse(users []dbcore.User) resources.UserCollection {
	var result []resources.UserData
	for _, user := range users {
		result = append(result, resources.UserData{
			Type: "user",
			Attributes: resources.UserDataAttributes{
				Id:       user.ID.String(),
				Username: user.Username,
				Title:    user.Title.String,
				Status:   user.Status.String,
				Avatar:   user.Avatar.String,
				Bio:      user.Bio.String,
			},
		})
	}

	return resources.UserCollection{Data: result}
}
