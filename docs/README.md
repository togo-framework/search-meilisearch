# search-meilisearch — documentation

Meilisearch driver for togo full-text search

## Overview

Package meilisearch is a Meilisearch driver for togo full-text search.
Blank-import it and set SEARCH_DRIVER=meilisearch, MEILI_HOST, MEILI_API_KEY.

## Install

```bash
togo install togo-framework/search-meilisearch
```

Set `SEARCH_DRIVER=meilisearch`.

## Configuration

Environment variables read by this plugin (extracted from the source — see the gateway/provider docs for each value):

| Env var |
|---|
| `MEILI_API_KEY` |
| `MEILI_HOST` |

## Usage

```go
s := k.Search
s.Index(ctx, "posts", doc)
hits, _ := s.Search(ctx, "posts", "query")
```

## Links

- Marketplace: https://to-go.dev/marketplace
- Source: https://github.com/togo-framework/search-meilisearch
- Full README: ../README.md
