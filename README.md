# search-meilisearch

> A **Meilisearch** driver for [togo](https://to-go.dev) full-text search.

Registers the `meilisearch` driver on the [`search`](https://github.com/togo-framework/search) plugin.

## Install

```sh
togo install togo-framework/search-meilisearch
```

Then in `.env`:

```sh
SEARCH_DRIVER=meilisearch
MEILI_HOST=https://your-instance.meilisearch.io
MEILI_API_KEY=your-key   # optional for unsecured dev instances
```

## Usage

Auto-registers on `togo serve`. Use the kernel search service:

```go
s, _ := search.FromKernel(k)
_ = s.Index(ctx, "posts", id, map[string]any{"title": "Hello"})
hits, _ := s.Search(ctx, "posts", "hello", 10)
```

## License

MIT
