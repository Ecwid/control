package tracing

/*
Configuration for memory dump. Used only when "memory-infra" category is enabled.
*/
type MemoryDumpConfig interface{}

/*
 */
type TraceConfig struct {
	RecordMode           string            `json:"recordMode,omitempty"`
	TraceBufferSizeInKb  float64           `json:"traceBufferSizeInKb,omitempty"`
	EnableSampling       bool              `json:"enableSampling,omitempty"`
	EnableSystrace       bool              `json:"enableSystrace,omitempty"`
	EnableArgumentFilter bool              `json:"enableArgumentFilter,omitempty"`
	IncludedCategories   []string          `json:"includedCategories,omitempty"`
	ExcludedCategories   []string          `json:"excludedCategories,omitempty"`
	SyntheticDelays      []string          `json:"syntheticDelays,omitempty"`
	MemoryDumpConfig     *MemoryDumpConfig `json:"memoryDumpConfig,omitempty"`
}

/*
	Data format of a trace. Can be either the legacy JSON format or the

protocol buffer format. Note that the JSON format will be deprecated soon.
*/
type StreamFormat string

/*
Compression type to use for traces returned via streams.
*/
type StreamCompression string

/*
	Details exposed when memory request explicitly declared.

Keep consistent with memory_dump_request_args.h and
memory_instrumentation.mojom
*/
type MemoryDumpLevelOfDetail string

/*
	Backend type to use for tracing. `chrome` uses the Chrome-integrated

tracing service and is supported on all platforms. `system` is only
supported on Chrome OS and uses the Perfetto system tracing service.
`auto` chooses `system` when the perfettoConfig provided to Tracing.start
specifies at least one non-Chrome data source; otherwise uses `chrome`.
*/
type TracingBackend string

type GetCategoriesVal struct {
	Categories []string `json:"categories"`
}

type RecordClockSyncMarkerArgs struct {
	SyncId string `json:"syncId"`
}

type RequestMemoryDumpArgs struct {
	Deterministic bool                    `json:"deterministic,omitempty"`
	LevelOfDetail MemoryDumpLevelOfDetail `json:"levelOfDetail,omitempty"`
}

type RequestMemoryDumpVal struct {
	DumpGuid string `json:"dumpGuid"`
	Success  bool   `json:"success"`
}

type StartArgs struct {
	BufferUsageReportingInterval float64           `json:"bufferUsageReportingInterval,omitempty"`
	TransferMode                 string            `json:"transferMode,omitempty"`
	StreamFormat                 StreamFormat      `json:"streamFormat,omitempty"`
	StreamCompression            StreamCompression `json:"streamCompression,omitempty"`
	TraceConfig                  *TraceConfig      `json:"traceConfig,omitempty"`
	PerfettoConfig               []byte            `json:"perfettoConfig,omitempty"`
	TracingBackend               TracingBackend    `json:"tracingBackend,omitempty"`
}
