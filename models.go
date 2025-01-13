package main

type Repo struct {
	Name    string
	Actions []Action
}

type Action struct {
	RepoName   string
	ID         string
	Status     string
	Conclusion string
	HeadBranch string
	HeadSHA    string
	Event      string
	CreatedAt  string
	UpdatedAt  string
	URL        string
}
