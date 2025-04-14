package events

type Event struct {
	ClickAllInvalidReasons      uint32         `avro:"bl" json:"bl"`
	ImpressionAllInvalidReasons int64          `avro:"cw" json:"cw"`
	Advertiser                  int64          `avro:"af" json:"af"`
	Publisher                   int64          `avro:"ag" json:"ag"`
	Site                        int64          `avro:"ah" json:"ah"`
	Zone                        int64          `avro:"ai" json:"ai"`
	Campaign                    int64          `avro:"ak" json:"ak"`
	Variation                   int64          `avro:"al" json:"al"`
	Offer                       int64          `avro:"be" json:"be"`
	CdnUsage                    uint64         `avro:"el" json:"el"`
	Value                       float64        `avro:"as" json:"as"`
	PublisherValue              float64        `avro:"-" json:"-"`
	Exchange2Publisher          float64        `avro:"aw" json:"aw"`
	Exchange2Network            float64        `avro:"ax" json:"ax"`
	ExchangeEUR2USD             float64        `avro:"cz" json:"cz"`
	RtbContentSize              float64        `avro:"bo" json:"bo"`
	RtbValue                    float64        `avro:"bs" json:"bs"`
	AdvertiserValueEUR          float64        `avro:"-" json:"-"` // Calculated
	AdvertiserValueUSD          float64        `avro:"-" json:"-"` // Calculated
	PublisherValueEUR           float64        `avro:"-" json:"-"` // Calculated
	PublisherValueUSD           float64        `avro:"-" json:"-"` // Calculated
	MarginEUR                   float64        `avro:"-" json:"-"` // Calculated
	MarginUSD                   float64        `avro:"-" json:"-"` // Calculated
	RtbValueEUR                 float64        `avro:"-" json:"-"` // Calculated
	RtbValueUSD                 float64        `avro:"-" json:"-"` // Calculated
	ExecutionTime               float64        `avro:"cf" json:"cf"`
	ConversionValueEUR          float64        `avro:"dr" json:"dr"`
	ConversionValueUSD          float64        `avro:"ds" json:"ds"`
	BidShadingValue             float32        `avro:"em" json:"em"`
	SecondBestValue             float32        `avro:"en" json:"en"`
	AuctionFloor                float32        `avro:"eo" json:"eo"`
	AuctionCeil                 float32        `avro:"ep" json:"ep"`
	BidShadingValueEur          float32        `avro:"-" json:"-"`
	BidShadingValueUsd          float32        `avro:"-" json:"-"`
	SecondBestValueEur          float32        `avro:"-" json:"-"`
	SecondBestValueUsd          float32        `avro:"-" json:"-"`
	AuctionFloorEur             float32        `avro:"-" json:"-"`
	AuctionFloorUsd             float32        `avro:"-" json:"-"`
	AuctionCeilEur              float32        `avro:"-" json:"-"`
	AuctionCeilUsd              float32        `avro:"-" json:"-"`
	RtbTotalTime                uint32         `avro:"bn" json:"bn"`
	RegionGeonameID             uint32         `avro:"bh" json:"bh"`
	LandingPage                 uint32         `avro:"bf" json:"bf"`
	Sub                         uint32         `avro:"ar" json:"ar"`
	Sub2                        uint32         `avro:"dc" json:"dc"`
	Sub3                        uint32         `avro:"dd" json:"dd"`
	ResponsiveLayout            uint32         `avro:"dm" json:"dm"`
	CityGeonameID               uint32         `avro:"cp" json:"cp"`
	MultiZone                   uint32         `avro:"dn" json:"dn"`
	ISP                         uint32         `avro:"aq" json:"aq"`
	Type                        int32          `avro:"bu" json:"bu"`
	PublisherRate               int32          `avro:"at" json:"at"`
	ImpressionInvalidReason     int32          `avro:"bc" json:"bc"`
	ProductCategory             uint16         `avro:"by" json:"by"` //changed
	ErrorCode                   int32          `avro:"ch" json:"ch"`
	CarrierSource               int32          `avro:"cl" json:"cl"`
	ProxyStatus                 int32          `avro:"cq" json:"cq"`
	Datacenter                  int32          `avro:"cr" json:"cr"`
	ClickX                      int32          `avro:"ct" json:"ct"`
	ClickY                      int32          `avro:"cu" json:"cu"`
	ScriptVersion               int32          `avro:"dq" json:"dq"`
	KeyValue1                   uint16         `avro:"eb" json:"eb"`
	KeyValue2                   uint16         `avro:"ec" json:"ec"`
	KeyValue3                   uint16         `avro:"ed" json:"ed"`
	KeyValue4                   uint16         `avro:"ee" json:"ee"`
	KeyValue5                   uint16         `avro:"ef" json:"ef"`
	KeyValue6                   uint16         `avro:"eg" json:"eg"`
	KeyValue7                   uint16         `avro:"eh" json:"eh"`
	KeyValue8                   uint16         `avro:"ei" json:"ei"`
	KeyValue9                   uint16         `avro:"ej" json:"ej"`
	KeyValue10                  uint16         `avro:"ek" json:"ek"`
	Network                     uint16         `avro:"ba" json:"ba"`
	Category                    uint16         `avro:"aj" json:"aj"`
	Browser                     uint16         `avro:"an" json:"an"`
	Device                      uint16         `avro:"ao" json:"ao"`
	Carrier                     uint16         `avro:"ap" json:"ap"`
	Partner                     uint16         `avro:"bi" json:"bi"`
	HTTPCode                    uint16         `avro:"ck" json:"ck"`
	BrowserVer                  uint16         `avro:"dh" json:"dh"`
	EmailISP                    uint16         `avro:"do" json:"do"`
	IdAdvertiserAdType          uint16         `avro:"dt" json:"dt"`
	IdRtbErrorReason            uint8          `avro:"et" json:"et"`
	OS                          uint8          `avro:"am" json:"am"`
	AdType                      uint8          `avro:"ay" json:"ay"`
	AdBlock                     uint8          `avro:"bd" json:"bd"`
	Language                    uint8          `avro:"ad" json:"ad"`
	Position                    uint8          `avro:"bg" json:"bg"`
	RtbRequestType              uint8          `avro:"bt" json:"bt"`
	RtbQueries                  uint8          `avro:"-" json:"-"` // Calculated
	RtbTimeouts                 uint8          `avro:"-" json:"-"` // Calculated
	RtbErrors                   uint8          `avro:"-" json:"-"` // Calculated
	RtbWins                     uint8          `avro:"-" json:"-"` // Calculated
	RtbNoBid                    uint8          `avro:"-" json:"-"` // Calculated
	RtbSkipped                  uint8          `avro:"-" json:"-"` // Calculated
	RtbSuccess                  uint8          `avro:"-" json:"-"` // Calculated
	Impressions                 uint8          `avro:"-" json:"-"` // Calculated
	InvalidImpressions          uint8          `avro:"-" json:"-"` // Calculated
	Clicks                      uint8          `avro:"-" json:"-"` // Calculated
	InvalidClicks               uint8          `avro:"-" json:"-"` // Calculated
	Goals                       uint8          `avro:"-" json:"-"` // Calculated
	VideoViews                  uint8          `avro:"-" json:"-"` // Calculated
	VideoHits                   uint8          `avro:"-" json:"-"` // Calculated
	AdRequest                   uint8          `avro:"-" json:"-"` // Calculated
	VideoEvent                  uint8          `avro:"cg" json:"cg"`
	WpnEvent                    uint8          `avro:"da" json:"da"`
	WpnSubscriptions            uint8          `avro:"-" json:"-"`
	WpnUnsubscriptions          uint8          `avro:"-" json:"-"`
	WpnSent                     uint8          `avro:"-" json:"-"`
	PricingModel                uint8          `avro:"bz" json:"bz"`
	HTTPS                       uint8          `avro:"ca" json:"ca"`
	ImpressionsStatus           uint8          `avro:"bb" json:"bb"`
	ClickStatus                 uint8          `avro:"bj" json:"bj"`
	RtbQueryStatus              uint8          `avro:"br" json:"br"`
	NoScript                    uint8          `avro:"cm" json:"cm"`
	IncognitoMode               uint8          `avro:"cn" json:"cn"`
	TestMode                    uint8          `avro:"co" json:"co"`
	VR                          uint8          `avro:"di" json:"di"`
	Internal                    uint8          `avro:"dj" json:"dj"`
	Sampling                    uint8          `avro:"dk" json:"dk"`
	NoFunds                     uint8          `avro:"dl" json:"dl"`
	ScriptDetail                uint8          `avro:"dp" json:"dp"`
	Proxy                       uint8          `avro:"du" json:"du"`
	AdRequestResult             uint8          `avro:"dv" json:"dv"`
	TrafficType                 uint8          `avro:"dw" json:"dw"`
	IdCampaignType              uint8          `avro:"dz" json:"dz"`
	RefreshIteration            uint8          `avro:"ea" json:"ea"`
	TrafficSource               uint8          `avro:"eq" json:"eq"`
	FallbackPosition            uint8          `avro:"er" json:"er"`
	IdClientHints               uint8          `avro:"eu" json:"eu"`
	EmailStatus                 int8           `avro:"db" json:"db"`
	AdblockVersion              string         `avro:"cs" json:"cs"`
	Country                     string         `avro:"ac" json:"ac"`
	SiteHostname                string         `avro:"ae" json:"ae"`
	PublisherCurrency           string         `avro:"au" json:"au"`
	CampaignCurrency            string         `avro:"av" json:"av"`
	AdFormat                    string         `avro:"az" json:"az"`
	InvalidReason               string         `avro:"bk" json:"bk"`
	Goal                        string         `avro:"bm" json:"bm"`
	BidStatus                   string         `avro:"bp" json:"bp"`
	ErrorReason                 string         `avro:"bq" json:"bq"`
	UserIP                      string         `avro:"bv" json:"bv"`
	CountryOriginal             string         `avro:"bw" json:"bw"`
	HTTPXForwardedFor           string         `avro:"bx" json:"bx"`
	ScreenResolution            string         `avro:"cb" json:"cb"`
	UserID                      string         `avro:"cc" json:"cc"`
	URLHash                     string         `avro:"cd" json:"cd"`
	RefererSiteHostname         string         `avro:"ce" json:"ce"`
	ImpID                       string         `avro:"ci" json:"ci"`
	BidID                       string         `avro:"cj" json:"cj"`
	ContainerResolution         string         `avro:"cv" json:"cv"`
	NetacuityProxyType          string         `avro:"cy" json:"cy"`
	NetacuityProxyDescription   string         `avro:"cx" json:"cx"`
	ScriptType                  string         `avro:"de" json:"de"`
	ScriptName                  string         `avro:"df" json:"df"`
	ScriptVer                   string         `avro:"dg" json:"dg"`
	RtbInventoryId              string         `avro:"dx" json:"dx"`
	RtbInventoryName            string         `avro:"dy" json:"dy"`
	AdServingHost               string         `avro:"es" json:"es"`
	CustomValue1                string         `avro:"ev" json:"ev"`
	CustomValue2                string         `avro:"ew" json:"ew"`
	CustomValue3                string         `avro:"ex" json:"ex"`
	Time                        ClickhouseTime `avro:"ab" json:"ab"`
	Date                        ClickhouseDate `avro:"aa" json:"aa"`
}

/*
structure NewNameMap{
	[ad_type] => ay,
    [format] => az,
    [idnetwork] => ba,
    [country_original] => bw,
    [http_x_forwarded_for] => bx,
    [idproduct_category] => by,
    [pricing_model] => bz,
    [https] => ca,
    [screen_resolution] => cb,
    [user_id] => cc,
    [url_hash] => cd,
    [referer_site_hostname] => ce,
    [adblock] => bd,
    [idoffer] => be,
    [idlanding_page] => bf,
    [imp_id] => ci,
    [bid_id] => cj,
    [http_code] => ck,
    [rtb_total_time] => bn,
    [rtb_content_size] => bo,
    [bid_status] => bp,
    [error_reason] => bq,
    [rtb_query_status] => br,
    [rtb_value] => bs,
    [iddatacenter] => cr,
    [idcarrier_source] => cl,
    [rtb_request_type ] =>,
    [noscript] => cm,
    [incognito_mode] => cn,
    [test_mode] => co,
    [region_geoname_id] => bh,
    [proxy_status] =>,
    [idpartner] => bi,
    [city_geoname_id] => cp,
    [netacuity_proxy_type] => cy,
    [netacuity_proxy_description] => cx,
    [adblock_ver] => cs,
    [impression_all_invalid_reasons] => cw,
    [sub2] => dc,
    [sub3] => dd,
    [script_type] => de,
    [script_name] => df,
    [script_ver] => dg,
    [browser_ver] => dh,
    [vr] => di,
    [internal] => dj,
    [sampling] => dk,
    [no_funds] => dl,
    [idmultizone] => dn,
    [email_isp] => do,
    [script_detail] => dp,
    [script_version] => dq,
    [idadvertiser_ad_type] => dt,
    [idproxy] => du,
    [idnetwork_type] => dw,
    [rtb_inventory_id] => dx,
    [rtb_inventory_name] => dy,
    [idcampaign_type] => dz,
    [key_value_1] => eb,
    [key_value_2] => ec,
    [key_value_3] => ed,
    [key_value_4] => ee,
    [key_value_5] => ef,
    [key_value_6] => eg,
    [key_value_7] => eh,
    [key_value_8] => ei,
    [key_value_9] => ej,
    [key_value_10] => ek,
    [bid_shading_value] => em,
    [second_best_value] => en,
    [auction_floor] => eo,
    [auction_ceil] => ep,
    [ad_request_result] => dv,
    [idtraffic_source] => eq,
    [fallback_position] => er,
    [ad_serving_host] => es,
    [idrtb_error_reason] => et,
    [custom_value_1] => ev,
    [custom_value_2] => ew,
    [custom_value_3] => ex,
    [impression_status] => bb,
    [impression_invalid_reason] => bc,
    [execution_time] => cf,
    [position] => bg,
    [container_resolution] => cv,
    [idresponsive_layout] => dm,
    [refresh_iteration] => ea,
    [cdn_usage] => el,
    [client_hints_usage] => eu,
    [is_trusted_proxy] => cq,
    [is_internal] =>,
    [video_event] => cg,
    [errorcode] => ch,
    [click_status] => bj,
    [click_invalid_reason] => bk,
    [clickx] => ct,
    [clicky] => cu,
    [click_all_invalid_reasons] => bl,
    [email_status] => db,
    [id_traffic_source] =>,
    [idgoal] => bm,
    [date] =>,
    [user_ip] => bv,
    [country] => ac,
    [idlanguage] => ad,
    [site_hostname] => ae,
    [idadvertiser] => af,
    [idpublisher] => ag,
    [idsite] => ah,
    [idzone] => ai,
    [idcategory] => aj,
    [idcampaign] => ak,
    [idvariation] => al,
    [idos] => am,
    [idbrowser] => an,
    [iddevice] => ao,
    [idcarrier] => ap,
    [idisp] => aq,
    [sub] => ar,
    [value] => as,
    [publisher_rate] => at,
    [publisher_currency] => au,
    [campaign_currency] => av,
    [exchange_to_publisher] => aw,
    [exchange_to_network] => ax,
}
*/

/*
[rtb_request_type] - Has an empty value in NewNameMap
[proxy_status] - Has an empty value in NewNameMap
[is_internal] - Has an empty value in NewNameMap
[id_traffic_source] - Has an empty value in NewNameMap
[date] - Has an empty value in NewNameMap
*/
type Event2 struct {
	click_all_invalid_reasons      uint32  `avro:"bl" json:"bl"` // Original: ClickAllInvalidReasons
	impression_all_invalid_reasons int64   `avro:"cw" json:"cw"` // Original: ImpressionAllInvalidReasons
	id_advertiser                  int64   `avro:"af" json:"af"` // Original: Advertiser
	id_publisher                   int64   `avro:"ag" json:"ag"` // Original: Publisher
	id_site                        int64   `avro:"ah" json:"ah"` // Original: Site
	id_zone                        int64   `avro:"ai" json:"ai"` // Original: Zone
	id_campaign                    int64   `avro:"ak" json:"ak"` // Original: Campaign
	id_variation                   int64   `avro:"al" json:"al"` // Original: Variation
	id_offer                       int64   `avro:"be" json:"be"` // Original: Offer
	cdn_usage                      uint64  `avro:"el" json:"el"` // Original: CdnUsage
	value                          float64 `avro:"as" json:"as"` // Original: Value
	exchange_to_publisher          float64 `avro:"aw" json:"aw"` // Original: Exchange2Publisher
	exchange_to_network            float64 `avro:"ax" json:"ax"` // Original: Exchange2Network
	// ExchangeEUR2USD             float64               `avro:"cz" json:"cz"` // No match in NewNameMap
	rtb_content_size float64 `avro:"bo" json:"bo"` // Original: RtbContentSize
	rtb_value        float64 `avro:"bs" json:"bs"` // Original: RtbValue
	execution_time   float64 `avro:"cf" json:"cf"` // Original: ExecutionTime
	// ConversionValueEUR          float64               `avro:"dr" json:"dr"` // No match in NewNameMap
	// ConversionValueUSD          float64               `avro:"ds" json:"ds"` // No match in NewNameMap
	bid_shading_value    float32 `avro:"em" json:"em"` // Original: BidShadingValue
	second_best_value    float32 `avro:"en" json:"en"` // Original: SecondBestValue
	auction_floor        float32 `avro:"eo" json:"eo"` // Original: AuctionFloor
	auction_ceil         float32 `avro:"ep" json:"ep"` // Original: AuctionCeil
	rtb_total_time       uint32  `avro:"bn" json:"bn"` // Original: RtbTotalTime
	region_geoname_id    uint32  `avro:"bh" json:"bh"` // Original: RegionGeonameID
	id_landing_page      uint32  `avro:"bf" json:"bf"` // Original: LandingPage
	sub                  uint32  `avro:"ar" json:"ar"` // Original: Sub
	sub2                 uint32  `avro:"dc" json:"dc"` // Original: Sub2
	sub3                 uint32  `avro:"dd" json:"dd"` // Original: Sub3
	id_responsive_layout uint32  `avro:"dm" json:"dm"` // Original: ResponsiveLayout
	city_geoname_id      uint32  `avro:"cp" json:"cp"` // Original: CityGeonameID
	id_multizone         uint32  `avro:"dn" json:"dn"` // Original: MultiZone
	id_isp               uint32  `avro:"aq" json:"aq"` // Original: ISP
	// Type                        int32                 `avro:"bu" json:"bu"` // No match in NewNameMap
	publisher_rate            int32  `avro:"at" json:"at"` // Original: PublisherRate
	impression_invalid_reason int32  `avro:"bc" json:"bc"` // Original: ImpressionInvalidReason
	id_product_category       uint16 `avro:"by" json:"by"` // Original: ProductCategory
	errorcode                 int32  `avro:"ch" json:"ch"` // Original: ErrorCode
	id_carrier_source         int32  `avro:"cl" json:"cl"` // Original: CarrierSource
	is_trusted_proxy          int32  `avro:"cq" json:"cq"` // Original: ProxyStatus
	id_datacenter             int32  `avro:"cr" json:"cr"` // Original: Datacenter
	clickx                    int32  `avro:"ct" json:"ct"` // Original: ClickX
	clicky                    int32  `avro:"cu" json:"cu"` // Original: ClickY
	script_version            int32  `avro:"dq" json:"dq"` // Original: ScriptVersion
	key_value_1               uint16 `avro:"eb" json:"eb"` // Original: KeyValue1
	key_value_2               uint16 `avro:"ec" json:"ec"` // Original: KeyValue2
	key_value_3               uint16 `avro:"ed" json:"ed"` // Original: KeyValue3
	key_value_4               uint16 `avro:"ee" json:"ee"` // Original: KeyValue4
	key_value_5               uint16 `avro:"ef" json:"ef"` // Original: KeyValue5
	key_value_6               uint16 `avro:"eg" json:"eg"` // Original: KeyValue6
	key_value_7               uint16 `avro:"eh" json:"eh"` // Original: KeyValue7
	key_value_8               uint16 `avro:"ei" json:"ei"` // Original: KeyValue8
	key_value_9               uint16 `avro:"ej" json:"ej"` // Original: KeyValue9
	key_value_10              uint16 `avro:"ek" json:"ek"` // Original: KeyValue10
	id_network                uint16 `avro:"ba" json:"ba"` // Original: Network
	id_category               uint16 `avro:"aj" json:"aj"` // Original: Category
	id_browser                uint16 `avro:"an" json:"an"` // Original: Browser
	id_device                 uint16 `avro:"ao" json:"ao"` // Original: Device
	id_carrier                uint16 `avro:"ap" json:"ap"` // Original: Carrier
	id_partner                uint16 `avro:"bi" json:"bi"` // Original: Partner
	http_code                 uint16 `avro:"ck" json:"ck"` // Original: HTTPCode
	browser_ver               uint16 `avro:"dh" json:"dh"` // Original: BrowserVer
	email_isp                 uint16 `avro:"do" json:"do"` // Original: EmailISP
	id_advertiser_ad_type     uint16 `avro:"dt" json:"dt"` // Original: IdAdvertiserAdType
	id_rtb_error_reason       uint8  `avro:"et" json:"et"` // Original: IdRtbErrorReason
	id_os                     uint8  `avro:"am" json:"am"` // Original: OS
	ad_type                   uint8  `avro:"ay" json:"ay"` // Original: AdType
	adblock                   uint8  `avro:"bd" json:"bd"` // Original: AdBlock
	id_language               uint8  `avro:"ad" json:"ad"` // Original: Language
	position                  uint8  `avro:"bg" json:"bg"` // Original: Position
	// RtbRequestType              uint8                 `avro:"bt" json:"bt"` // No match in NewNameMap
	video_event uint8 `avro:"cg" json:"cg"` // Original: VideoEvent
	// WpnEvent                    uint8                 `avro:"da" json:"da"` // No match in NewNameMap
	pricing_model               uint8  `avro:"bz" json:"bz"` // Original: PricingModel
	https                       uint8  `avro:"ca" json:"ca"` // Original: HTTPS
	impression_status           uint8  `avro:"bb" json:"bb"` // Original: ImpressionsStatus
	click_status                uint8  `avro:"bj" json:"bj"` // Original: ClickStatus
	rtb_query_status            uint8  `avro:"br" json:"br"` // Original: RtbQueryStatus
	noscript                    uint8  `avro:"cm" json:"cm"` // Original: NoScript
	incognito_mode              uint8  `avro:"cn" json:"cn"` // Original: IncognitoMode
	test_mode                   uint8  `avro:"co" json:"co"` // Original: TestMode
	vr                          uint8  `avro:"di" json:"di"` // Original: VR
	internal                    uint8  `avro:"dj" json:"dj"` // Original: Internal
	sampling                    uint8  `avro:"dk" json:"dk"` // Original: Sampling
	no_funds                    uint8  `avro:"dl" json:"dl"` // Original: NoFunds
	script_detail               uint8  `avro:"dp" json:"dp"` // Original: ScriptDetail
	id_proxy                    uint8  `avro:"du" json:"du"` // Original: Proxy
	ad_request_result           uint8  `avro:"dv" json:"dv"` // Original: AdRequestResult
	id_network_type             uint8  `avro:"dw" json:"dw"` // Original: TrafficType
	id_campaign_type            uint8  `avro:"dz" json:"dz"` // Original: IdCampaignType
	refresh_iteration           uint8  `avro:"ea" json:"ea"` // Original: RefreshIteration
	id_traffic_source           uint8  `avro:"eq" json:"eq"` // Original: TrafficSource
	fallback_position           uint8  `avro:"er" json:"er"` // Original: FallbackPosition
	client_hints_usage          uint8  `avro:"eu" json:"eu"` // Original: IdClientHints
	email_status                int8   `avro:"db" json:"db"` // Original: EmailStatus
	adblock_ver                 string `avro:"cs" json:"cs"` // Original: AdblockVersion
	country                     string `avro:"ac" json:"ac"` // Original: Country
	site_hostname               string `avro:"ae" json:"ae"` // Original: SiteHostname
	publisher_currency          string `avro:"au" json:"au"` // Original: PublisherCurrency
	campaign_currency           string `avro:"av" json:"av"` // Original: CampaignCurrency
	format                      string `avro:"az" json:"az"` // Original: AdFormat
	click_invalid_reason        string `avro:"bk" json:"bk"` // Original: InvalidReason
	id_goal                     string `avro:"bm" json:"bm"` // Original: Goal
	bid_status                  string `avro:"bp" json:"bp"` // Original: BidStatus
	error_reason                string `avro:"bq" json:"bq"` // Original: ErrorReason
	user_ip                     string `avro:"bv" json:"bv"` // Original: UserIP
	country_original            string `avro:"bw" json:"bw"` // Original: CountryOriginal
	http_x_forwarded_for        string `avro:"bx" json:"bx"` // Original: HTTPXForwardedFor
	screen_resolution           string `avro:"cb" json:"cb"` // Original: ScreenResolution
	user_id                     string `avro:"cc" json:"cc"` // Original: UserID
	url_hash                    string `avro:"cd" json:"cd"` // Original: URLHash
	referer_site_hostname       string `avro:"ce" json:"ce"` // Original: RefererSiteHostname
	imp_id                      string `avro:"ci" json:"ci"` // Original: ImpID
	bid_id                      string `avro:"cj" json:"cj"` // Original: BidID
	container_resolution        string `avro:"cv" json:"cv"` // Original: ContainerResolution
	netacuity_proxy_type        string `avro:"cy" json:"cy"` // Original: NetacuityProxyType
	netacuity_proxy_description string `avro:"cx" json:"cx"` // Original: NetacuityProxyDescription
	script_type                 string `avro:"de" json:"de"` // Original: ScriptType
	script_name                 string `avro:"df" json:"df"` // Original: ScriptName
	script_ver                  string `avro:"dg" json:"dg"` // Original: ScriptVer
	rtb_inventory_id            string `avro:"dx" json:"dx"` // Original: RtbInventoryId
	rtb_inventory_name          string `avro:"dy" json:"dy"` // Original: RtbInventoryName
	ad_serving_host             string `avro:"es" json:"es"` // Original: AdServingHost
	custom_value_1              string `avro:"ev" json:"ev"` // Original: CustomValue1
	custom_value_2              string `avro:"ew" json:"ew"` // Original: CustomValue2
	custom_value_3              string `avro:"ex" json:"ex"` // Original: CustomValue3
	// Time                        ClickhouseTime        `avro:"ab" json:"ab"` // No match in NewNameMap
	// Date                        ClickhouseDate        `avro:"aa" json:"aa"` // No match in NewNameMap
}
