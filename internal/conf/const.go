package conf

const (
	TypeString = "string"
	TypeSelect = "select"
	TypeBool   = "bool"
	TypeText   = "text"
	TypeNumber = "number"
)

const (
	// site
	VERSION      = "version"
	ApiUrl       = "api_url"
	BasePath     = "base_path"
	SiteTitle    = "site_title"
	Logo         = "logo"
	Favicon      = "favicon"
	Announcement = "announcement"
	IconColor    = "icon_color"

	// preview
	TextTypes     = "text_types"
	AudioTypes    = "audio_types"
	VideoTypes    = "video_types"
	ImageTypes    = "image_types"
	OfficeTypes   = "office_types"
	ProxyTypes    = "proxy_types"
	OfficeViewers = "office_viewers"
	PdfViewers    = "pdf_viewers"
	AudioAutoplay = "audio_autoplay"
	VideoAutoplay = "video_autoplay"

	// global
	HideFiles      = "hide_files"
	GlobalReadme   = "global_readme"
	CustomizeHead  = "customize_head"
	CustomizeBody  = "customize_body"
	LinkExpiration = "link_expiration"
	PrivacyRegs    = "privacy_regs"

	// aria2
	Aria2Uri    = "aria2_uri"
	Aria2Secret = "aria2_secret"

	// single
	Token = "token"
)

const (
	UNKNOWN = iota
	FOLDER
	OFFICE
	VIDEO
	AUDIO
	TEXT
	IMAGE
)
