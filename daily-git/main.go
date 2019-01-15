package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/shurcooL/githubv4"
	"golang.org/x/oauth2"
)

func main() {
	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")},
	)
	httpClient := oauth2.NewClient(context.Background(), src)

	client := githubv4.NewClient(httpClient)

	var query struct {
		Viewer struct {
			ContributionsCollection struct {
				TotalCommitContributions githubv4.Int
			}
			Login     githubv4.String
			CreatedAt githubv4.DateTime
		}
	}
	err := client.Query(context.Background(), &query, nil)
	if err != nil {
		panic(err)
	}

	fmt.Println(query.Viewer.Login)
	fmt.Println("Joined", query.Viewer.CreatedAt.Format(time.RFC850))
	fmt.Println(query.Viewer.ContributionsCollection.TotalCommitContributions, "Contributions in the last year")

}
