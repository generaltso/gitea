package models

// import "github.com/generaltso/linguist"
import api "code.gitea.io/sdk/gitea"

func getLanguageStatistics() []*api.LanguageStat {
	l1 := &api.LanguageStat{
		Language:   "PHP",
		Percentage: 0.47 * 100,
		Color:      "#ff0000",
	}
	l2 := &api.LanguageStat{
		Language:   "JavaScript",
		Percentage: 0.23 * 100,
		Color:      "#00ff00",
	}
	l3 := &api.LanguageStat{
		Language:   "CSS",
		Percentage: 0.3 * 100,
		Color:      "#0000ff",
	}
	return []*api.LanguageStat{l1, l2, l3}
}
