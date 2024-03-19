package mirip

import (
	"errors"
	"sort"
	"strings"

	"github.com/samber/lo"
)

type kvScore struct {
	Key   string
	Score float64
}

type KvScores []kvScore

type Adapter interface {
	Compare(string, string) float64
}

func Compare(adapter Adapter, txt string, minimal float64, ignoreCase bool, others ...string) (string, float64, error) {
	originalTxt := txt
	_, ok := lo.Find(others, func(other string) bool {
		if ignoreCase {
			other = strings.ToLower(other)
			txt = strings.ToLower(txt)
		}
		return txt == other
	})
	if ok {
		return originalTxt, 0, nil
	}

	mapScores := make(map[string]float64)
	for _, other := range others {
		originalOther := other
		if ignoreCase {
			other = strings.ToLower(other)
		}

		_, ok := mapScores[other]
		if ok {
			continue
		}
		other = sortNamePhrase(other)
		txt = sortNamePhrase(txt)
		cmp := adapter.Compare(txt, other)

		if cmp < minimal {
			continue
		}
		mapScores[originalOther] = cmp
	}

	if len(mapScores) == 0 {
		return "", 0, errors.New("not found")
	}

	scores := KvScores(lo.MapToSlice(mapScores, func(k string, v float64) kvScore {
		return kvScore{k, v}
	}))

	sort.Sort(scores)
	return scores[0].Key, scores[0].Score, nil
}

func (r KvScores) Len() int {
	return len(r)
}

func (r KvScores) Less(i, j int) bool {
	return r[i].Score > r[j].Score
}

func (r KvScores) Swap(i, j int) {
	tmp := r[i]
	r[i] = r[j]
	r[j] = tmp
}

func sortNamePhrase(name string) string {
	words := strings.Split(name, " ")
	if len(words) == 1 {
		return name
	}

	sort.Strings(words)
	return strings.Join(words, " ")
}
