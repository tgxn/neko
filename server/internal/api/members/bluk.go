package members

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/m1k1o/neko/server/pkg/types"
	"github.com/m1k1o/neko/server/pkg/utils"
)

type MemberBulkUpdatePayload struct {
	IDs     []string            `json:"ids"`
	Profile types.MemberProfile `json:"profile"`
}

func (h *MembersHandler) membersBulkUpdate(w http.ResponseWriter, r *http.Request) error {
	bytes, err := io.ReadAll(r.Body)
	if err != nil {
		return utils.HttpBadRequest("unable to read post body").WithInternalErr(err)
	}

	header := &MemberBulkUpdatePayload{}
	if err := json.Unmarshal(bytes, &header); err != nil {
		return utils.HttpBadRequest("unable to unmarshal payload").WithInternalErr(err)
	}

	for _, memberId := range header.IDs {
		// TODO: Bulk select?
		profile, err := h.members.Select(memberId)
		if err != nil {
			return utils.HttpInternalServerError().
				WithInternalErr(err).
				WithInternalMsg("unable to select member profile").
				Msgf("failed to update member %s", memberId)
		}

		body := &MemberBulkUpdatePayload{
			Profile: profile,
		}

		if err := json.Unmarshal(bytes, &body); err != nil {
			return utils.HttpBadRequest().
				WithInternalErr(err).
				Msgf("unable to unmarshal payload for member %s", memberId)
		}

		if err := h.members.UpdateProfile(memberId, body.Profile); err != nil {
			return utils.HttpInternalServerError().
				WithInternalErr(err).
				WithInternalMsg("unable to update member profile").
				Msgf("failed to update member %s", memberId)
		}
	}

	return utils.HttpSuccess(w)
}

type MemberBulkDeletePayload struct {
	IDs []string `json:"ids"`
}

func (h *MembersHandler) membersBulkDelete(w http.ResponseWriter, r *http.Request) error {
	bytes, err := io.ReadAll(r.Body)
	if err != nil {
		return utils.HttpBadRequest("unable to read post body").WithInternalErr(err)
	}

	data := &MemberBulkDeletePayload{}
	if err := json.Unmarshal(bytes, &data); err != nil {
		return utils.HttpBadRequest("unable to unmarshal payload").WithInternalErr(err)
	}

	for _, memberId := range data.IDs {
		if err := h.members.Delete(memberId); err != nil {
			return utils.HttpInternalServerError().
				WithInternalErr(err).
				WithInternalMsg("unable to delete member").
				Msgf("failed to delete member %s", memberId)
		}
	}

	return utils.HttpSuccess(w)
}
