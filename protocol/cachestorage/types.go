package cachestorage

/*
Unique identifier of the Cache object.
*/
type CacheId string

/*
type of HTTP response cached
*/
type CachedResponseType string

/*
Data entry.
*/
type DataEntry struct {
	RequestURL         string             `json:"requestURL"`
	RequestMethod      string             `json:"requestMethod"`
	RequestHeaders     []*Header          `json:"requestHeaders"`
	ResponseTime       float64            `json:"responseTime"`
	ResponseStatus     int                `json:"responseStatus"`
	ResponseStatusText string             `json:"responseStatusText"`
	ResponseType       CachedResponseType `json:"responseType"`
	ResponseHeaders    []*Header          `json:"responseHeaders"`
}

/*
Cache identifier.
*/
type Cache struct {
	CacheId        CacheId `json:"cacheId"`
	SecurityOrigin string  `json:"securityOrigin"`
	StorageKey     string  `json:"storageKey"`
	CacheName      string  `json:"cacheName"`
}

/*
 */
type Header struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

/*
Cached response
*/
type CachedResponse struct {
	Body []byte `json:"body"`
}

type DeleteCacheArgs struct {
	CacheId CacheId `json:"cacheId"`
}

type DeleteEntryArgs struct {
	CacheId CacheId `json:"cacheId"`
	Request string  `json:"request"`
}

type RequestCacheNamesArgs struct {
	SecurityOrigin string `json:"securityOrigin,omitempty"`
	StorageKey     string `json:"storageKey,omitempty"`
}

type RequestCacheNamesVal struct {
	Caches []*Cache `json:"caches"`
}

type RequestCachedResponseArgs struct {
	CacheId        CacheId   `json:"cacheId"`
	RequestURL     string    `json:"requestURL"`
	RequestHeaders []*Header `json:"requestHeaders"`
}

type RequestCachedResponseVal struct {
	Response *CachedResponse `json:"response"`
}

type RequestEntriesArgs struct {
	CacheId    CacheId `json:"cacheId"`
	SkipCount  int     `json:"skipCount,omitempty"`
	PageSize   int     `json:"pageSize,omitempty"`
	PathFilter string  `json:"pathFilter,omitempty"`
}

type RequestEntriesVal struct {
	CacheDataEntries []*DataEntry `json:"cacheDataEntries"`
	ReturnCount      float64      `json:"returnCount"`
}
