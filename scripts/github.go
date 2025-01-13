package main

import (
	"context"
	"github.com/google/go-github/v68/github"
	_ "github.com/joho/godotenv/autoload"
	"log"
	"os"
	// "strconv"
)

func main() {
	github_key := os.Getenv("GITHUB_KEY")
	client := github.NewClient(nil).WithAuthToken(github_key)

	opt := &github.RepositoryListByUserOptions{Type: "all"}

	repos, _, err := client.Repositories.ListByUser(context.Background(), "edgargdev", opt)

	if err != nil {
		log.Fatal(err)
	}

	workflow_runs := github.WorkflowRuns{}

	for _, repo := range repos {
		log.Println(*repo.Name)
		actions, response, err := client.Actions.ListRepositoryWorkflowRuns(context.Background(), "edgargdev", *repo.Name, nil)
		if response.StatusCode == 404 {
			log.Println("No actions found for repo: ", *repo.Name)
			continue
		}
		if err != nil {
			log.Fatal(err)
		}
		log.Println("Repo: ", *repo.Name)
		if *actions.TotalCount > 0 {
			for _, action := range actions.WorkflowRuns {
				workflow_runs.WorkflowRuns = append(workflow_runs.WorkflowRuns, action)
			}
		}
	}

	for _, workflow_run := range workflow_runs.WorkflowRuns {
		log.Println("WorkflowRun: ", *workflow_run.ID)
		log.Println("Status: ", *workflow_run.Status)
		log.Println("Conclusion: ", *workflow_run.Conclusion)
		log.Println("HeadBranch: ", *workflow_run.HeadBranch)
		log.Println("HeadSha: ", *workflow_run.HeadSHA)
		log.Println("Event: ", *workflow_run.Event)
		log.Println("CreatedAt: ", *workflow_run.CreatedAt)
		log.Println("UpdatedAt: ", *workflow_run.UpdatedAt)
		log.Println("URL: ", *workflow_run.HTMLURL)
	}

}
