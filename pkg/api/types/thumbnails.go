package types

// ThumbnailType represents the type of thumbnail.
type ThumbnailType string

const (
	AvatarHeadShot   ThumbnailType = "AvatarHeadShot"
	AvatarBust       ThumbnailType = "AvatarBust"
	AvatarFullBody   ThumbnailType = "AvatarFullBody"
	GameIcon         ThumbnailType = "GameIcon"
	BadgeIcon        ThumbnailType = "BadgeIcon"
	GamePass         ThumbnailType = "GamePass"
	DeveloperProduct ThumbnailType = "DeveloperProduct"
	GroupIcon        ThumbnailType = "GroupIcon"
)

// ThumbnailFormat represents the format of the thumbnail.
type ThumbnailFormat string

const (
	PNG  ThumbnailFormat = "png"
	JPEG ThumbnailFormat = "jpeg"
)

// ThumbnailSize represents the size of the thumbnail.
type ThumbnailSize string

const (
	Size30x30   ThumbnailSize = "30x30"
	Size48x48   ThumbnailSize = "48x48"
	Size60x60   ThumbnailSize = "60x60"
	Size75x75   ThumbnailSize = "75x75"
	Size100x100 ThumbnailSize = "100x100"
	Size110x110 ThumbnailSize = "110x110"
	Size140x140 ThumbnailSize = "140x140"
	Size150x150 ThumbnailSize = "150x150"
	Size180x180 ThumbnailSize = "180x180"
	Size250x250 ThumbnailSize = "250x250"
	Size352x352 ThumbnailSize = "352x352"
	Size420x420 ThumbnailSize = "420x420"
	Size720x720 ThumbnailSize = "720x720"
)

// ThumbnailRequest represents a single thumbnail request.
type ThumbnailRequest struct {
	Type       ThumbnailType   `json:"type"`
	Token      string          `json:"token,omitempty"`
	RequestID  string          `json:"requestId"`
	TargetID   uint64          `json:"targetId"`
	IsCircular bool            `json:"isCircular,omitempty"`
	Format     ThumbnailFormat `json:"format,omitempty"`
	Size       ThumbnailSize   `json:"size"`
}

// ThumbnailDataResponse represents the data for a single thumbnail in the response returned by the Roblox API.
type ThumbnailDataResponse struct {
	RequestID    string `json:"requestId"`
	ErrorCode    int    `json:"errorCode"`
	ErrorMessage string `json:"errorMessage"`
	TargetID     uint64 `json:"targetId"`
	State        string `json:"state"`
	ImageURL     string `json:"imageUrl"`
	Version      string `json:"version"`
}
