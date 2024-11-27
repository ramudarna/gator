# Gator CLI
Gator is a cli tool to fetch and manage RSS feeds. This guide will help you set the environment and run the program.

## Prerequisites
To run Gator, you will need the following:
- [Postgres](https://www.postgresql.org/download/) 
- [Go](https://golang.org/doc/install)

## Installation
### Step 1: Clone the repository

### Step 2: Install Gator cli
bash:
`go install`

### Step 3: Set up the config file\
Create a Json config file in the root directory(according to your OS) with the following code
Json:
```
{
  "current_user_name": "your_username",
  "db": {
    "host": "localhost",
    "port": 5432,
    "user": "your_db_user",
    "password": "your_db_password",
    "dbname": "your_db_name"
  }
}
```
Replace all the placeholders with your credentials

### Step 4: Run the Program
Here are some commands that u can run in the cli
-Add a Feed
bash:
`gator addfeed <name> https://example.com/feed`

-Follow a Feeed by giving the feed name and the feed url
bash:
`gator follow https://example.com/feed`

-Unfollow a Feed
bash:
`gator unfollow https://example.com/feed`

-List Feeds You Are Following
bash:
`gator following`

-Aggregate Feeds
bash:
`gator agg 1m`
The agg command will fetch and print the posts from the feeds every specified duration (e.g., 1m for 1 minute).

-Browse posts
bash:
`gator browse 2`
The browse command will browse all the posts for the current user and the limit of posts is set to 2(default is 2, can be left empty)

Example RSS Feeds
Here are a few RSS feeds to get you started:

TechCrunch: https://techcrunch.com/feed/

Hacker News: https://news.ycombinator.com/rss

Boot.devBlog: https://blog.boot.dev/index.xml
