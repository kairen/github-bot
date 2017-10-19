package hanlder

import (
	"bot/api"
	"fmt"
	"strconv"
	"strings"

	webhooks "gopkg.in/go-playground/webhooks.v3"
	"gopkg.in/go-playground/webhooks.v3/github"
	"gopkg.in/go-playground/webhooks.v3/gitlab"
)

// A Account is GitHub and GitLab client
type Account struct {
	githubToken    string
	gitlabToken    string
	gitlabEndpoint string
}

// A Server is an webhook server
type Server struct {
	port   int
	secret string
	path   string
}

// NewAccount returns a client using the given github token,.., etc
func NewAccount(ht string, lt string, ep string) *Account {
	return &Account{githubToken: ht, gitlabToken: lt, gitlabEndpoint: ep}
}

// NewServer returns a server using the given port,.., etc
func NewServer(port int, secret string, path string) *Server {
	return &Server{port: port, secret: secret, path: path}
}

// InitAccount is init github and gitlab account
func (a *Account) InitAccount() {
	ghc := api.NewGitHub(a.githubToken)
	ghc.InitGitHubClient()

	glc := api.NewGitLab(a.gitlabToken, a.gitlabEndpoint)
	glc.InitGitLabClient()
	api.GitLabPRComment("kairen", "gitlab-ci-test", 2, "test")
}

// RunServer is run webhook server
func (s *Server) RunServer() {
	gh := github.New(&github.Config{Secret: s.secret})
	gh.RegisterEvents(GitHubIssueCommentHandler, github.IssueCommentEvent)

	gl := gitlab.New(&gitlab.Config{Secret: s.secret})
	gl.RegisterEvents(GitLabPipelineHandler, gitlab.PipelineEvents)

	hooks := []webhooks.Webhook{gh, gl}
	err := webhooks.Run(hooks, ":"+strconv.Itoa(s.port), s.path)
	if err != nil {
		fmt.Println(err)
	}
}

// GitHubIssueCommentHandler handles GitHub pull request events
func GitHubIssueCommentHandler(payload interface{}, header webhooks.Header) {
	pl := payload.(github.IssueCommentPayload)
	command := pl.Comment.Body
	if pl.Action == "created" && isPullRequestComment(pl.Issue.HTMLURL) {
		switch command {
		case "/ok-to-test":
			user := pl.Sender.Login
			if user != "kairen" {
				msg := "@" + user + ": You can't run testing, because you are not a member or collaborator."
				repo := pl.Repository
				api.GitHubPRComment(repo.Owner.Login, repo.Name, int(pl.Issue.Number), msg)
			}
		default:
			fmt.Println("Event trigger")
		}
	}
}

// GitLabPipelineHandler handles GitLab pipeline events
func GitLabPipelineHandler(payload interface{}, header webhooks.Header) {
	pl := payload.(gitlab.PipelineEventPayload)
	fmt.Printf("%+v", pl.Project.Name)
}

func isPullRequestComment(url string) bool {
	i := strings.Index(url, "pull")
	return (i > 0)
}
