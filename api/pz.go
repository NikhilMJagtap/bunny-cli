package api

import (
    "fmt"
    "github.com/NikhilMJagtap/bunny-cli/client"
)

type ListPullZoneOpts struct {
    Page               uint32 `json:"page"`
    PerPage            uint32 `json:"perPage"`
    Search             string `json:"search"`
    IncludeCertificate bool   `json:"includeCertificate"`
}

type AddPullZoneOpts struct {
    OriginURL                           string
    AllowedReferrers                    []string
    BlockedReferrers                    []string
    BlockedIPs                          []string
    EnableGeoZoneUS                     bool
    EnableGeoZoneEU                     bool
    EnableGeoZoneAsia                   bool
    EnableGeoZoneSA                     bool
    EnableGeoZoneAF                     bool
    BlockRootPathAccess                 bool
    EnableQueryStringOrdering           bool
    EnableWebpVary                      bool
    EnableAvifVary                      bool
    EnableMobileVary                    bool
    EnableCountryCodeVary               bool
    EnableHostnameVary                  bool
    EnableCacheSlice                    bool
    ZoneSecurityEnabled                 bool
    ZoneSecurityIncludeHashRemoteIP     bool
    IgnoreQueryStrings                  bool
    MonthlyBandwidthLimit               int64
    AccessControlOriginHeaderExtensions []string
    EnableAccessControlOriginHeader     bool
    DisableCookies                      bool
    BudgetRedirectedCountries           []string
    BlockedCountries                    []string
    CacheControlMaxAgeOverride          int64
    CacheControlPublicMaxAgeOverride    int64
    CacheControlBrowserMaxAgeOverride   int64
    AddHostHeader                       bool
    AddCanonicalHeader                  bool
    EnableLogging                       bool
    LoggingIPAnonymizationEnabled       bool
    PermaCacheStorageZoneId             int64
    AWSSigningEnabled                   bool
    AWSSigningKey                       string
    AWSSigningSecret                    string
    AWSSigningRegionName                string
    EnableOriginShield                  bool
    OriginShieldZoneCode                string
    EnableTLS1                          bool
    EnableTLS11                         bool
    CacheErrorResponses                 bool
    VerifyOriginSSL                     bool
    LogForwardingEnabled                bool
    LogForwardingHostname               string
    LogForwardingPort                   int32
    LogForwardingToken                  string
    LogForwardingProtocol               int
    LoggingSaveToStorage                bool
    LoggingStorageZoneId                int64
    FollowRedirects                     bool
    ConnectionLimitPerIPCount           int32
    RequestLimit                        int32
    LimitRateAfter                      float64
    LimitRatePerSecond                  int32
    BurstSize                           int32
    WAFEnabled                          bool
    WAFDisabledRuleGroups               []string
    WAFDisabledRules                    []string
    WAFEnableRequestHeaderLogging       bool
    WAFRequestHeaderIgnores             []string
    ErrorPageEnableCustomCode           bool
    ErrorPageCustomCode                 string
    ErrorPageEnableStatuspageWidget     bool
    ErrorPageStatuspageCode             string
    ErrorPageWhitelabel                 bool
    OptimizerEnabled                    bool
    OptimizerDesktopMaxWidth            int32
    OptimizerMobileMaxWidth             int32
    OptimizerImageQuality               int
    OptimizerMobileImageQuality         int
}

func ListPullZones(b *client.BunnyClient, options *ListPullZoneOpts) (interface{}, error) {
    return b.Get("/pullzone/", GetQueryParamsFromOptions(options))
}

func GetPullZone(b *client.BunnyClient, pullZoneId uint64) (interface{}, error) {
    return b.Get(fmt.Sprintf("/pullzone/%d/", pullZoneId), nil)
}
