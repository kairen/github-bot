package webhook

import (
	"github-bot/pkg/api"

	webhooks "gopkg.in/go-playground/webhooks.v3"
	"gopkg.in/go-playground/webhooks.v3/gitlab"
)

// GitLabPipelineHandler handles gitlab pipeline events
func GitLabPipelineHandler(payload interface{}, header webhooks.Header) {
	pl := payload.(gitlab.PipelineEventPayload)
	pepelineStatus := pl.ObjectAttributes.Status

	status := &api.GitHubStatus{
		State:       pepelineStatus,
		Context:     "gitlab-ci/pipeline",
		Description: "Pipeline " + pepelineStatus,
		TargetURL:   pl.Project.WebURL + "/" + string(pl.ObjectAttributes.ID),
	}
	status.CheckStatus()

	api.CreateGitHubStatus(
		pl.Project.Namespace,
		pl.Project.Name,
		pl.ObjectAttributes.SHA,
		status,
	)
}
