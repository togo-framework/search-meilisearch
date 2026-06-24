// Package meilisearch is a Meilisearch driver for togo full-text search.
// Blank-import it and set SEARCH_DRIVER=meilisearch, MEILI_HOST, MEILI_API_KEY.
package meilisearch

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/togo-framework/search"
	"github.com/togo-framework/togo"
)

func init() {
	search.RegisterDriver("meilisearch", func(k *togo.Kernel) (search.Searcher, error) {
		host := strings.TrimRight(os.Getenv("MEILI_HOST"), "/")
		if host == "" {
			return nil, errors.New("search-meilisearch: MEILI_HOST not set")
		}
		return &searcher{host: host, key: os.Getenv("MEILI_API_KEY"), client: &http.Client{Timeout: 15 * time.Second}}, nil
	})
}

type searcher struct {
	host, key string
	client    *http.Client
}

func (s *searcher) req(ctx context.Context, method, path string, body any) (*http.Response, error) {
	var r io.Reader
	if body != nil {
		b, _ := json.Marshal(body)
		r = bytes.NewReader(b)
	}
	req, err := http.NewRequestWithContext(ctx, method, s.host+path, r)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	if s.key != "" {
		req.Header.Set("Authorization", "Bearer "+s.key)
	}
	return s.client.Do(req)
}

func (s *searcher) Index(ctx context.Context, index, id string, doc map[string]any) error {
	d := map[string]any{"id": id}
	for k, v := range doc {
		d[k] = v
	}
	resp, err := s.req(ctx, http.MethodPost, "/indexes/"+url.PathEscape(index)+"/documents", []map[string]any{d})
	if err != nil {
		return err
	}
	return drain(resp)
}

func (s *searcher) Search(ctx context.Context, index, query string, limit int) ([]search.Hit, error) {
	if limit <= 0 {
		limit = 20
	}
	resp, err := s.req(ctx, http.MethodPost, "/indexes/"+url.PathEscape(index)+"/search", map[string]any{"q": query, "limit": limit})
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode >= 300 {
		b, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("search-meilisearch: %s: %s", resp.Status, b)
	}
	var out struct {
		Hits []map[string]any `json:"hits"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return nil, err
	}
	hits := make([]search.Hit, 0, len(out.Hits))
	for _, h := range out.Hits {
		id, _ := h["id"].(string)
		hits = append(hits, search.Hit{ID: id, Doc: h})
	}
	return hits, nil
}

func (s *searcher) Delete(ctx context.Context, index, id string) error {
	resp, err := s.req(ctx, http.MethodDelete, "/indexes/"+url.PathEscape(index)+"/documents/"+url.PathEscape(id), nil)
	if err != nil {
		return err
	}
	return drain(resp)
}

func drain(resp *http.Response) error {
	defer resp.Body.Close()
	if resp.StatusCode >= 300 {
		b, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("search-meilisearch: %s: %s", resp.Status, b)
	}
	io.Copy(io.Discard, resp.Body)
	return nil
}
