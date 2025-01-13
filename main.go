package main

import (
	"context"
	"github.com/google/go-github/v68/github"
	_ "github.com/joho/godotenv/autoload"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"
)

var templates = template.Must(template.ParseFiles(
	"templates/index.html",
	"templates/actions.html",
))

func indexHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string][]string{
		"Repos": repos,
	}

	err := templates.ExecuteTemplate(w, "index.html", data)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func actionsHandler(w http.ResponseWriter, r *http.Request) {
	repo_name := r.PathValue("repo_name")
	github_key := os.Getenv("GITHUB_KEY")

	client := github.NewClient(nil).WithAuthToken(github_key)

	actions, response, err := client.Actions.ListRepositoryWorkflowRuns(context.Background(), "edgargdev", repo_name, nil)

	repo := Repo{Name: repo_name, Actions: []Action{}}

	if response.StatusCode >= 200 && response.StatusCode <= 399 {
		for _, action := range actions.WorkflowRuns {
			action := Action{
				RepoName:   repo_name,
				ID:         strconv.FormatInt(*action.ID, 10),
				Status:     *action.Status,
				Conclusion: *action.Conclusion,
				HeadBranch: *action.HeadBranch,
				HeadSHA:    *action.HeadSHA,
				Event:      *action.Event,
				CreatedAt:  action.CreatedAt.Format("2006-01-02T15:04:05"),
				UpdatedAt:  action.UpdatedAt.Format("2006-01-02T15:04:05"),
				URL:        *action.HTMLURL,
			}
			repo.Actions = append(repo.Actions, action)
		}
	}
	data := map[string][]Action{
		"Actions": repo.Actions,
	}
	err = templates.ExecuteTemplate(w, "actions.html", data)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/actions/{repo_name}", actionsHandler)

	port := ":8080"
	log.Println("Server is running on http://localhost" + port)
	server := &http.Server{
		Addr:    port,
		Handler: mux,
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
