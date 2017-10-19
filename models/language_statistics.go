package models

import (
	"sort"

	api "code.gitea.io/sdk/gitea"

	"github.com/generaltso/linguist"
)

type LanguageStatistics []*api.LanguageStat

func (l LanguageStatistics) Len() int {
	return len(l)
}

func (l LanguageStatistics) Less(i, j int) bool {
	return l[i].Percentage < l[j].Percentage
}

func (l LanguageStatistics) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

func getLanguageStatistics() LanguageStatistics {
	l1 := &api.LanguageStat{
		Language:   "PHP",
		Percentage: 0.47,
	}
	l2 := &api.LanguageStat{
		Language:   "JavaScript",
		Percentage: 0.23,
	}
	l3 := &api.LanguageStat{
		Language:   "CSS",
		Percentage: 0.3,
	}
	stats := LanguageStatistics([]*api.LanguageStat{l1, l2, l3})
	for i, l := range stats {
		stats[i].Percentage *= 100
		stats[i].Color = linguist.LanguageColor(l.Language)
	}
	sort.Sort(sort.Reverse(stats))
	return stats
}
