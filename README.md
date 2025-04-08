## Blog aggregator in Go

- Add RSS feeds from across the internet to be collected
- Store the collected posts in a PostgreSQL database
- Follow and unfollow RSS feeds that other users have added
- View summaries of the aggregated posts in the terminal, with a link to the full post

## Setup 
for running you need postgre (15 and later) and go 

1. create ~/.gatorconfig.json, with the following content:

```json
{
  "db_url": "protocol://username:password@host:port/database?sslmode=disable"
}
```
> dont forget ssl mode

2. create database, run migrations with goose

- install goose
```bash
go install github.com/pressly/goose/v3/cmd/goose@latest

```
- check if installed
```bash
goose version

```
- run migration
```bash
cd sql/schema
goose postgres <connection_string> up
```


### Avialable commands

#### register my_name

#### login my_name

#### feeds
> list of all feeds (creator, name of feed, feed url)
 
#### addfeed "Lanes Blog" "https://www.wagslane.dev/index.xml"
> loged user adding feed

#### follow <feed url>
> follow feed of another user

#### unfollow <feed url>
> unfollow feed with url

#### following
> list of my feeds (titles)

#### agg <time interval>
> time interval is string: 1s, 3s, 3h ...
> fetch the actual posts from the feed URLs and store them in our database

#### browse <number of posts to display>
> view all the posts from the feeds the user follows
