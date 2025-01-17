package handlers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/recovery-flow/comtools/cifractx"
	"github.com/recovery-flow/comtools/httpkit"
	"github.com/recovery-flow/comtools/httpkit/problems"
	"github.com/recovery-flow/tokens"
	"github.com/recovery-flow/users-storage/internal/config"
	"github.com/recovery-flow/users-storage/internal/service/roles"
	"github.com/sirupsen/logrus"
)

func MemberDelete(w http.ResponseWriter, r *http.Request) {
	server, err := cifractx.GetValue[*config.Service](r.Context(), config.SERVER)
	if err != nil {
		logrus.Errorf("Failed to retrieve service configuration: %v", err)
		httpkit.RenderErr(w, problems.InternalError("Failed to retrieve service configuration"))
		return
	}
	log := server.Logger

	userID, ok := r.Context().Value(tokens.UserIDKey).(uuid.UUID)
	if !ok {
		log.Warn("UserID not found in context")
		httpkit.RenderErr(w, problems.Unauthorized("User not authenticated"))
		return
	}

	teamId, err := uuid.Parse(chi.URLParam(r, "team_id"))
	if err != nil {
		httpkit.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	deletedUserId, err := uuid.Parse(chi.URLParam(r, "user_id"))
	if err != nil {
		httpkit.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	team, err := server.MongoDB.Teams.FilterById(teamId).Get(r.Context())
	if err != nil {
		log.Errorf("Failed to get team: %v", err)
		httpkit.RenderErr(w, problems.InternalError())
		return
	}

	cond := false
	for _, member := range team.Members {
		if member.ID == userID {
			if roles.CompareRolesTeam(member.Role, roles.RoleTeamAdmin) != -1 {
				cond = true
				break
			}
		}
	}

	if !cond {
		httpkit.RenderErr(w, problems.Forbidden("You don't have permissions to remove member"))
		return
	}

	teamMembers, err := server.MongoDB.Teams.FilterById(teamId).Members()
	if err != nil {
		log.Errorf("Failed to get team: %v", err)
		httpkit.RenderErr(w, problems.InternalError())
		return
	}

	_, err = teamMembers.FilterByUserId(deletedUserId).Delete(r.Context())
	if err != nil {
		log.Errorf("Failed to delete member: %v", err)
		httpkit.RenderErr(w, problems.InternalError())
		return
	}

	httpkit.Render(w, http.StatusOK)
}
