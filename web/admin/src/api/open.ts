import { request } from '@/utils/http'
const GITHUB_GRAPHQL_API = 'https://api.github.com/graphql'

export const queryGithubContributions = (username: string) => {
  const query = `
    query {
      user(login: "${username}") {
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
    }
  `
  return request({
    url: GITHUB_GRAPHQL_API, method: 'POST', data: { query }, headers: {
    }
  })
  //     Authorization: `bearer ${process.env.GITHUB_TOKEN}`,
}