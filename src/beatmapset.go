package gosu

type Beatmapset struct {
	BeatmapsetCompact
	downloadDisabled    bool
	moreInformation     string
	bpm                 float64
	canBeHyped          bool
	creator             string
	discussionEnabled   bool
	discussionLocked    bool
	currentHype         int
	requiredHype        int
	isScoreable         bool
	lastUpdated         Timestamp
	legacyThreadURL     string
	currentNominations  int
	requiredNominations int
	ranked              int
	rankedDate          Timestamp
	storyboard          bool
	submittedDate       Timestamp
	tags                string
}

type BeatmapsetCompact struct {
	artist        string
	artistUnicode string
	covers        Covers
	creator       string
	favoriteCount int
	id            int
	nsfw          bool
	playCount     int
	previewURL    string
	source        string
	status        RankStatus
	title         string
	titleUnicode  string
	userID        int
	video         string

	// Optional attributes
	// FIXME: Actual Types for most of these fields are unspecified
	beatmaps              []Beatmap
	converts              string
	currentUserAttributes string
	description           string
	discussions           string
	events                string
	genre                 string
	hasFavorited          bool // always included
	language              string
	nominations           string
	ratings               string
	recentFavorites       string
	relatedUsers          string
	user                  string
}