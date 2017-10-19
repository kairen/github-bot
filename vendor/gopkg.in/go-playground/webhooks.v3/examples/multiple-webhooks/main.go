package main

import (
	"fmt"
	"strconv"
        
	"gopkg.in/go-playground/webhooks.v3"
	"gopkg.in/go-playground/webhooks.v3/github"
	"gopkg.in/go-playground/webhooks.v3/gitlab"
)

const (
	path = "/webhooks"
	port = 8080
)

func main() {

	githubHook := github.New(&github.Config{Secret: "r00tme"})
	githubHook.RegisterEvents(GitHubHandlePullRequest, github.PullRequestEvent)

	gitlabHook := gitlab.New(&gitlab.Config{Secret: "r00tme"})
	gitlabHook.RegisterEvents(GitLabHandlePipeline, gitlab.PipelineEvents)

	hooks := []webhooks.Webhook{gitlabHook, githubHook}
	err := webhooks.Run(hooks, ":"+strconv.Itoa(port), path)
	if err != nil {
		fmt.Println(err)
	}
}

// GitHubHandlePullRequest handles GitHub pull_request events
func GitHubHandlePullRequest(payload interface{}, header webhooks.Header) {

	fmt.Println("Handling GitHub Pull Request")

	pl := payload.(github.PullRequestPayload)

	// Do whatever you want from here...
	fmt.Printf("%+v", pl)
}

// GitLabHandlePipeline handles GitLab pipeline events
func GitLabHandlePipeline(payload interface{}, header webhooks.Header) {

	fmt.Println("Handling GitLab Pipeline")

	pl := payload.(gitlab.PipelineEventPayload)

	// Do whatever you want from here...
	fmt.Printf("%+v", pl)
}
