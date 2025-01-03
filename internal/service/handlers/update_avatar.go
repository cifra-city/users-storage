package handlers

import (
	"net/http"

	"github.com/cifra-city/comtools/cifractx"
	"github.com/cifra-city/comtools/httpkit"
	"github.com/cifra-city/comtools/httpkit/problems"
	"github.com/cifra-city/tokens"
	"github.com/cifra-city/users-storage/internal/config"
	"github.com/cifra-city/users-storage/internal/service/requests"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

func UpdateAvatar(w http.ResponseWriter, r *http.Request) {
	server, err := cifractx.GetValue[*config.Service](r.Context(), config.SERVER)
	if err != nil {
		logrus.Errorf("Failed to retrieve service configuration: %v", err)
		httpkit.RenderErr(w, problems.InternalError("Failed to retrieve service configuration"))
		return
	}
	log := server.Logger

	req, err := requests.NewUpdateAvatarRequest(r)
	if err != nil {
		log.Info("Failed to parse request: ", err)
		httpkit.RenderErr(w, problems.BadRequest(err)...)
		return
	}
	defer req.File.Close()

	userID, ok := r.Context().Value(tokens.UserIDKey).(uuid.UUID)
	if !ok {
		server.Logger.Warn("UserID not found in context")
		httpkit.RenderErr(w, problems.Unauthorized("User not authenticated"))
		return
	}

	yes := true
	uploadParams := uploader.UploadParams{
		Folder:       "avatars",
		PublicID:     userID.String(),
		Overwrite:    &yes,
		ResourceType: "image",
	}
	uploadResult, err := server.Storage.Upload.Upload(r.Context(), req.File, uploadParams)
	if err != nil {
		log.Errorf("Failed to upload avatar to Cloudinary: %v", err)
		httpkit.RenderErr(w, problems.InternalError("Failed to upload avatar"))
		return
	}

	_, err = server.Databaser.Users.UpdateAvatar(r.Context(), userID, &uploadResult.SecureURL)
	if err != nil {
		log.Errorf("Failed to update avatar URL in database: %v", err)
		httpkit.RenderErr(w, problems.InternalError("Failed to save avatar"))
		return
	}

	response := map[string]string{"avatar_url": uploadResult.SecureURL}
	httpkit.Render(w, response)
}
