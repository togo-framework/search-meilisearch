<!-- togo-header -->
<div align="center">
  <img src=".github/assets/togo-mark.svg" alt="togo" height="64" />
  <h1>togo-framework/search-meilisearch</h1>
  <p>
    <a href="https://to-go.dev/marketplace"><img src="https://img.shields.io/badge/marketplace-to--go.dev-1FC7DC" alt="marketplace" /></a>
    <a href="https://pkg.go.dev/github.com/togo-framework/search-meilisearch"><img src="https://pkg.go.dev/badge/github.com/togo-framework/search-meilisearch.svg" alt="pkg.go.dev" /></a>
    <img src="https://img.shields.io/badge/license-MIT-blue" alt="MIT" />
  </p>
  <p><strong>Part of the <a href="https://to-go.dev">togo</a> framework.</strong></p>
</div>

## Install

```bash
togo install togo-framework/search-meilisearch
```

<!-- /togo-header -->

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

<!-- togo-sponsors -->
---

<div align="center">
  <h3>Premium sponsors</h3>
  <p>
    <a href="https://id8media.com"><strong>ID8 Media</strong></a> &nbsp;·&nbsp;
    <a href="https://one-studio.co"><strong>One Studio</strong></a>
  </p>
  <p><sub>Support togo — <a href="https://github.com/sponsors/fadymondy">become a sponsor</a>.</sub></p>
</div>
<!-- /togo-sponsors -->
