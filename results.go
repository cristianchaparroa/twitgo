package twitgo

// SearchResult
type SearchResult struct {
	SearchMetadata SearchMetaData `json:"search_metadata"`
	Statuses       []Status       `json:"statuses"`
}

// SearchMetaData
type SearchMetaData struct {
	CompletedIn float64 `json:"completed_in"`
	Count       int     `json:"count"`
	MaxID       int     `json:"max_id"`
	MaxIDStr    string  `json:"max_id_str"`
	NextResults string  `json:"next_results"`
	Query       string  `json:"query"`
	RefreshURL  string  `json:"refresh_url"`
	SinceID     int     `json:"since_id"`
	SinceIDStr  string  `json:"since_id_str"`
}

// Status
type Status struct {
	Contributors         interface{}      `json:"contributors"`
	Coordinates          interface{}      `json:"coordinates"`
	CreatedAt            string           `json:"created_at"`
	Entities             StatusesEntities `json:"entities"`
	FavoriteCount        int              `json:"favorite_count"`
	Favorited            bool             `json:"favorited"`
	Geo                  interface{}      `json:"geo"`
	ID                   int              `json:"id"`
	IDStr                string           `json:"id_str"`
	InReplyToScreenName  string           `json:"in_reply_to_screen_name"`
	InReplyToStatusID    interface{}      `json:"in_reply_to_status_id"`
	InReplyToStatusIDStr interface{}      `json:"in_reply_to_status_id_str"`
	InReplyToUserID      int              `json:"in_reply_to_user_id"`
	InReplyToUserIDStr   string           `json:"in_reply_to_user_id_str"`
	IsQuoteStatus        bool             `json:"is_quote_status"`
	Lang                 string           `json:"lang"`
	Metadata             StatusMetaData   `json:"metadata"`
	Place                interface{}      `json:"place"`
	// Number of times retweet this tweet
	RetweetCount         int              `json:"retweet_count"`
	// Indicate if this tweet was retweet
	Retweeted            bool             `json:"retweeted"`
	Source               string           `json:"source"`
	// This is the text of the tweet
	Text                 string           `json:"text"`
	Truncated            bool             `json:"truncated"`
	// User that make the tweet
	User                 User             `json:"user"`
}

// User
type User struct {
	ContributorsEnabled            bool         `json:"contributors_enabled"`
	CreatedAt                      string       `json:"created_at"`
	DefaultProfile                 bool         `json:"default_profile"`
	DefaultProfileImage            bool         `json:"default_profile_image"`
	Description                    string       `json:"description"`
	Entities                       UserEntities `json:"entities"`
	FavouritesCount                int          `json:"favourites_count"`
	FollowRequestSent              interface{}  `json:"follow_request_sent"`
	FollowersCount                 int          `json:"followers_count"`
	Following                      interface{}  `json:"following"`
	FriendsCount                   int          `json:"friends_count"`
	GeoEnabled                     bool         `json:"geo_enabled"`
	HasExtendedProfile             bool         `json:"has_extended_profile"`
	ID                             int          `json:"id"`
	IDStr                          string       `json:"id_str"`
	IsTranslationEnabled           bool         `json:"is_translation_enabled"`
	IsTranslator                   bool         `json:"is_translator"`
	Lang                           string       `json:"lang"`
	ListedCount                    int          `json:"listed_count"`
	Location                       string       `json:"location"`
	Name                           string       `json:"name"`
	Notifications                  interface{}  `json:"notifications"`
	ProfileBackgroundColor         string       `json:"profile_background_color"`
	ProfileBackgroundImageURL      string       `json:"profile_background_image_url"`
	ProfileBackgroundImageURLHTTPS string       `json:"profile_background_image_url_https"`
	ProfileBackgroundTile          bool         `json:"profile_background_tile"`
	ProfileBannerURL               string       `json:"profile_banner_url"`
	ProfileImageURL                string       `json:"profile_image_url"`
	ProfileImageURLHTTPS           string       `json:"profile_image_url_https"`
	ProfileLinkColor               string       `json:"profile_link_color"`
	ProfileSidebarBorderColor      string       `json:"profile_sidebar_border_color"`
	ProfileSidebarFillColor        string       `json:"profile_sidebar_fill_color"`
	ProfileTextColor               string       `json:"profile_text_color"`
	ProfileUseBackgroundImage      bool         `json:"profile_use_background_image"`
	Protected                      bool         `json:"protected"`
	ScreenName                     string       `json:"screen_name"`
	StatusesCount                  int          `json:"statuses_count"`
	TimeZone                       string       `json:"time_zone"`
	URL                            interface{}  `json:"url"`
	UtcOffset                      int          `json:"utc_offset"`
	Verified                       bool         `json:"verified"`
}

// UserMention
type UserMention struct {
	ID         int    `json:"id"`
	IDStr      string `json:"id_str"`
	Indices    []int  `json:"indices"`
	Name       string `json:"name"`
	ScreenName string `json:"screen_name"`
}

// StatusMetaData
type StatusMetaData struct {
	IsoLanguageCode string `json:"iso_language_code"`
	ResultType      string `json:"result_type"`
}

// UserEntities
type UserEntities struct {
	Description struct {
			    Urls []interface{} `json:"urls"`
		    }
}

// StatusesEntities
type StatusesEntities struct {
	Hashtags     []interface{} `json:"hashtags"`
	Symbols      []interface{} `json:"symbols"`
	Urls         []interface{} `json:"urls"`
	UserMentions []UserMention `json:"user_mentions"`
}
