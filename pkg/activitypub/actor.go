package activitypub

import (
	"fmt"
	pb "github.com/disism/hvxahv/api/accounts/v1alpha1"
	"github.com/spf13/viper"
	"strings"
	"time"
)

// GetActorName Get the username in the request url such,
// as "/.well-known/webFinger?resource=acct:hvturingga@0efb43b41a8a.ngrok.io" Will get hvturingga,
// If the match fails, it will return a custom username not found error.
func GetActorName(resource string) string {
	if strings.HasPrefix(resource, "acct:") {
		resource = resource[5:]
		if ali := strings.IndexByte(resource, '@'); ali != -1 {
			resource = resource[:ali]
		}
	}

	return resource
}

// IsRemote Get host to determine whether it is a remote instance user (not a user of this instance).
func IsRemote(resource string) bool {
	host := GetHost(resource)
	if !strings.Contains(resource, "@") {
		return false
	}
	if host != viper.GetString("localhost") {
		return true
	}
	return false

}

// GetHost Get the host of the received resource.
func GetHost(resource string) string {
	if strings.HasPrefix(resource, "acct:") {
		ali := strings.IndexByte(resource, '@')
		if ali != -1 {
			return resource[ali+1:]
		}
	}
	return resource[5:]
}

// EXAMPLE 9
//{
//  "@context": ["https://www.w3.org/ns/activitystreams",
//               {"@language": "ja"}],
//  "type": "Person",
//  "id": "https://kenzoishii.example.com/",
//  "following": "https://kenzoishii.example.com/following.json",
//  "followers": "https://kenzoishii.example.com/followers.json",
//  "liked": "https://kenzoishii.example.com/liked.json",
//  "inbox": "https://kenzoishii.example.com/inbox.json",
//  "outbox": "https://kenzoishii.example.com/feed.json",
//  "preferredUsername": "kenzoishii",
//  "name": "石井健蔵",
//  "summary": "この方はただの例です",
//  "icon": [
//    "https://kenzoishii.example.com/image/165987aklre4"
//  ]
//}

type icon struct {
	Type      string `json:"type"`
	MediaType string `json:"mediaType"`
	Url       string `json:"url"`
}

func NewContext() []string {
	ctx := []string{"https://www.w3.org/ns/activitystreams", "https://w3id.org/security/v1alpha1"}

	return ctx

}

func NewIcon(url string) *icon {
	t := "Image"
	mt := "image/jpg"
	return &icon{Type: t, MediaType: mt, Url: url}
}

type actor struct {
	Context           []string    `json:"@context"`
	Type              string      `json:"type"`
	ID                string      `json:"id"`
	Following         string      `json:"following"`
	Followers         string      `json:"followers"`
	Liked             string      `json:"liked"`
	Inbox             string      `json:"inbox"`
	Outbox            string      `json:"outbox"`
	PreferredUsername string      `json:"preferredUsername"`
	Name              string      `json:"name"`
	Summary           string      `json:"summary"`
	PublicKey         interface{} `json:"public_key"`
	Icon              *icon       `json:"icon"`
}

// NewActor Return standard ActivityPub protocol user data.
func NewActor(a *pb.AccountData) *actor {
	addr := viper.GetString("localhost")

	act := fmt.Sprintf("https://%s/u/%s", addr, a.Username)
	box := fmt.Sprintf("https://%s/u/%s/", addr, a.Username)

	publicKey := map[string]string{
		"id":           fmt.Sprintf("https://%s/u/%s#main-key", addr, a.Username),
		"owner":        act,
		"publicKeyPem": a.PublicKey,
	}

	actor := &actor{
		Context:           NewContext(),
		Type:              "Person",
		ID:                act,
		Inbox:             box + "inbox",
		Outbox:            box + "outbox",
		PreferredUsername: a.Username,
		Name:              a.Name,
		Summary:           a.Bio,
		PublicKey:         publicKey,
		Icon:              NewIcon(a.Avatar),
	}
	return actor
}

type Actor struct {
	Context                   []interface{} `json:"@context"`
	Id                        string        `json:"id"`
	Type                      string        `json:"type"`
	Following                 string        `json:"following"`
	Followers                 string        `json:"followers"`
	Inbox                     string        `json:"inbox"`
	Outbox                    string        `json:"outbox"`
	Featured                  string        `json:"featured"`
	FeaturedTags              string        `json:"featuredTags"`
	PreferredUsername         string        `json:"preferredUsername"`
	Name                      string        `json:"name"`
	Summary                   string        `json:"summary"`
	Url                       string        `json:"url"`
	ManuallyApprovesFollowers bool          `json:"manuallyApprovesFollowers"`
	Discoverable              bool          `json:"discoverable"`
	Published                 time.Time     `json:"published"`
	Devices                   string        `json:"devices"`
	PublicKey                 struct {
		Id           string `json:"id"`
		Owner        string `json:"owner"`
		PublicKeyPem string `json:"publicKeyPem"`
	} `json:"publicKey"`
	Tag        []interface{} `json:"tag"`
	Attachment []interface{} `json:"attachment"`
	Endpoints  struct {
		SharedInbox string `json:"sharedInbox"`
	} `json:"endpoints"`
	Icon struct {
		Type      string `json:"type"`
		MediaType string `json:"mediaType"`
		Url       string `json:"url"`
	} `json:"icon"`
}