package service

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/amyunfei/glassy-sky/internal/admin/dto"
)

type GithubService interface {
	GetGithubContributions(ctx context.Context)
}

type DefaultGithubService struct{}

func (s *DefaultGithubService) GetGithubContributions(ctx context.Context) {
	client := &http.Client{}
	query := dto.GithubGraphQLQuery{
		Query: `query {
		  user(login: "amyunfei") {
		    contributionsCollection {
		      contributionCalendar {
		        totalContributions
		        weeks {
		          contributionDays {
		            contributionCount
		            date
		            weekday
		          }
		        }
		      }
		    }
		  }
		}`,
	}
	params, err := json.Marshal(query)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(params))
	request, err := http.NewRequest("POST", "https://api.github.com/graphql", strings.NewReader(string(params)))
	if err != nil {
		fmt.Println(err)
	}
	request.Header.Add("Authorization", "Bearer")
	request.Header.Add("Content-Type", "application/json")
	res, err := client.Do(request)
	fmt.Println(res, err)
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(body))
	defer res.Body.Close()
}
