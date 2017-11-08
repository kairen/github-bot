package webhook

import (
	"github-bot/pkg/api"
	"strconv"

	webhooks "gopkg.in/go-playground/webhooks.v3"
	"gopkg.in/go-playground/webhooks.v3/gitlab"
)

// GitLabPipelineHandler handles gitlab pipeline events
func GitLabPipelineHandler(payload interface{}, header webhooks.Header) {
	pl := payload.(gitlab.PipelineEventPayload)
	id := strconv.FormatInt(pl.ObjectAttributes.ID, 10)
	for _, build := range pl.Builds {
		stats := &api.GitHubStatus{
			State:       build.Status,
			Context:     build.Name,
			Description: "Job " + build.Status,
			TargetURL:   pl.Project.WebURL + "/pipelines/" + id,
		}
		stats.CheckStatus()
		go api.CreateGitHubStatus(
			pl.Project.Namespace,
			pl.Project.Name,
			pl.ObjectAttributes.SHA,
			stats,
		)
	}
}
