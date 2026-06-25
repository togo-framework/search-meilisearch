# search-meilisearch — documentation

  <img src=".github/assets/togo-mark.svg" alt="togo" height="64" />

## Overview

Package meilisearch is a Meilisearch driver for togo full-text search.
Blank-import it and set SEARCH_DRIVER=meilisearch, MEILI_HOST, MEILI_API_KEY.

## Install

```bash
togo install togo-framework/search-meilisearch
```

Set `SEARCH_DRIVER=meilisearch`.

## Configuration

Environment variables read by this plugin (extracted from the source):

| Env var | Notes |
|---|---|
| `G` | _see provider docs_ |
| `MEILI_API_KEY` | _see provider docs_ |
| `MEILI_HOST` | _see provider docs_ |

## Usage

```go
s := k.Search
s.Index(ctx, "posts", doc)
hits, _ := s.Search(ctx, "posts", "query")
```

## Links

- Marketplace: https://to-go.dev/marketplace
- Source: https://github.com/togo-framework/search-meilisearch
- README: ../README.md
