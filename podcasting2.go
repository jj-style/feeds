package feeds

import "encoding/xml"

// https://podcastindex.org/namespace/1.0
// https://podcasting2.org/podcast-namespace
// https://podcasting2.org/podcast-namespace/tags/liveItem

const (
	Podcasting2Namespace = "https://podcastindex.org/namespace/1.0"
)

type Podcasting2Feed struct {
	Locked  Podcasting2Locked  `xml:"podcast:locked,omitempty"`
	Podping Podcasting2Podping `xml:"podcast:podping,omitempty"`
}

type Podcasting2Podping struct {
	UsesPodping bool `xml:"usesPodping,attr,omitempty"`
}

type Podcasting2LockedState string

var (
	Locked   Podcasting2LockedState = "yes"
	Unlocked Podcasting2LockedState = "no"
)

type Podcasting2Locked struct {
	Owner string                 `xml:"owner,attr,omitempty"`
	Value Podcasting2LockedState `xml:",chardata"`
}

type Podcasting2Item struct {
	ContentLink        *Podcasting2ContentLink        `xml:"podcast:contentLink,omitempty"`
	AlternateEnclosure *Podcasting2AlternateEnclosure `xml:"podcast:alternateEnclosure,omitempty"`
}

type Podcasting2ContentLink struct {
	Href string `xml:"href,attr,omitempty"`
	Text string `xml:",chardata"`
}

type Podcasting2AlternateEnclosure struct {
	Type    string            `xml:"type,attr"`
	Length  string            `xml:"length,attr"`
	Default bool              `xml:"default,attr,omitempty"`
	Source  Podcasting2Source `xml:"podcast:source"`
}

type Podcasting2Source struct {
	Url string `xml:"url,attr"`
}

type Podcasting2LiveItem struct {
	XMLName xml.Name `xml:"podcast:liveItem"`
	*RssItem
	Status string `xml:"status,attr,omitempty"` // A string that must be one of pending, live or ended.
	Start  string `xml:"start,attr,omitempty"`  // A string representing an ISO8601 timestamp that denotes the time when the stream is intended to start.
	End    string `xml:"end,attr,omitempty"`    // A string representing an ISO8601 timestamp that denotes the time when the stream is intended to end.
}
