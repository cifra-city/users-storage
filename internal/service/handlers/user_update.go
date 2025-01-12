package handlers

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/recovery-flow/comtools/cifractx"
	"github.com/recovery-flow/comtools/httpkit"
	"github.com/recovery-flow/comtools/httpkit/problems"
	"github.com/recovery-flow/tokens"
	"github.com/recovery-flow/users-storage/internal/config"
	"github.com/recovery-flow/users-storage/internal/service/requests"
	"github.com/recovery-flow/users-storage/resources"
	"github.com/sirupsen/logrus"
)

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	server, err := cifractx.GetValue[*config.Service](r.Context(), config.SERVER)
	if err != nil {
		logrus.Errorf("Failed to retrieve service configuration: %v", err)
		httpkit.RenderErr(w, problems.InternalError("Failed to retrieve service configuration"))
		return
	}
	log := server.Logger

	req, err := requests.NewUpdateUser(r)
	if err != nil {
		httpkit.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	username := req.Data.Attributes.Username
	avatar := req.Data.Attributes.Avatar

	userID, ok := r.Context().Value(tokens.UserIDKey).(uuid.UUID)
	if !ok {
		log.Warn("UserID not found in context")
		httpkit.RenderErr(w, problems.Unauthorized("User not authenticated"))
		return
	}

	num, err := server.MongoDB.Users.FilterById(userID).UpdateAvatar(r.Context(), avatar)
	if err != nil {
		log.Errorf("Failed to update username: %v", err)
		httpkit.RenderErr(w, problems.InternalError())
		return
	}
	if num == 0 {
		httpkit.RenderErr(w, problems.InternalError())
		return
	}

	num, err = server.MongoDB.Users.FilterById(userID).UpdateUsername(r.Context(), username)
	if err != nil {
		log.Errorf("Failed to update username: %v", err)
		httpkit.RenderErr(w, problems.InternalError())
		return
	}
	if num == 0 {
		httpkit.RenderErr(w, problems.InternalError())
		return
	}

	users, err := server.MongoDB.Users.FilterById(userID).Get(r.Context())
	if err != nil {
		log.Errorf("Failed to update city: %v", err)
		httpkit.RenderErr(w, problems.InternalError())
		return
	}

	httpkit.Render(w, NewUserResponse(users, resources.UserUpdateType))
}