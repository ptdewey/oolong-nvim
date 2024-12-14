package internal

import (
	"fmt"
	"sort"
)

type KeywordResp struct {
	Keyword   string  `json:"keyword"`
	Weight    float64 `json:"weight"`
	Count     int     `json:"count"`
	Documents map[string]struct {
		DocumentWeight float64 `json:"DocumentWeight"`
		DocumentCount  int     `json:"DocumentCount"`
	} `json:"documents"`
}

type NoteResp struct {
	Path    string             `json:"path"`
	Weights map[string]float64 `json:"weights"`
}

type kv struct {
	Key    string
	Weight float64
}

func formatKeywordSearchResp(r KeywordResp) [][]byte {
	lines := []string{
		"# Keyword: " + r.Keyword,
		fmt.Sprintf("- Global Weight: %f", r.Weight),
		fmt.Sprintf("- Count: %d", r.Count),
		"",
		"## Document Occurences",
	}

	kvList := []kv{}
	for k, v := range r.Documents {
		kvList = append(kvList, kv{k, v.DocumentWeight})
	}

	sort.Slice(kvList, func(i, j int) bool {
		return kvList[i].Weight > kvList[j].Weight
	})

	for _, kv := range kvList {
		v := r.Documents[kv.Key]
		lines = append(lines, []string{
			kv.Key,
			fmt.Sprintf("- Weight: %f, Count: %d", v.DocumentWeight, v.DocumentCount),
			"",
		}...)
	}

	out := [][]byte{}
	for _, line := range lines {
		out = append(out, []byte(line))
	}

	return out
}

// TODO: Update with count for each keyword (requires update to API)
func formatNoteSearchResp(r NoteResp) [][]byte {
	lines := []string{
		"# Note: " + r.Path,
		fmt.Sprintf("- Keyword Count: %d", len(r.Weights)),
		"",
		"## Document Keywords",
	}

	var kvList []kv
	for k, v := range r.Weights {
		kvList = append(kvList, kv{k, v})
	}

	sort.Slice(kvList, func(i, j int) bool {
		return kvList[i].Weight > kvList[j].Weight
	})

	for _, kv := range kvList {
		lines = append(lines, fmt.Sprintf("- %s: %f", kv.Key, kv.Weight))
	}

	out := [][]byte{}
	for _, line := range lines {
		out = append(out, []byte(line))
	}

	return out
}
