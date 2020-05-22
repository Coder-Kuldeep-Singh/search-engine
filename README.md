# search-engine

Google, DuckDuckGo search engine results collector

## Installation

```bash
git clone git@github.com:mavensingh/search-engine.git
```

```bash
go mod init github.com/googlesearch
```

## Running Software

There is two type of engine you can use with the flags

```bash
Flags
E  = Engine
q = query


go run main.go -E "google" -q "whatever you want to search"
```

## Search Engines We are Using

| No  | Engine     | Returns                    |
| --- | ---------- | -------------------------- |
| 1   | Google     | Titles, Links, Description |
| 2   | DuckDuckGo | Titles, Links, Description |
| 3   | Bing       | Titles, Links, Description |
| 4   | Yahoo      | Titles, Links, Description |
