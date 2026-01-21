package types

type ChaosReport struct {
	Summary          Summary                   `json:"summary"`
	Hailer           Hailer                    `json:"hailer"`
	Hosts            map[string]HostData       `json:"hosts"`
	ResourceHistory  []HealthCheckEvent        `json:"resource_history"`
	OperationSummary map[string]OperationStats `json:"operation_summary"`
}

type Summary struct {
	TotalOps      int     `json:"total_operations"`
	ChangedOps    int     `json:"changed_operations"`
	SuccessOps    int     `json:"successful_operations"`
	FailedOps     int     `json:"failed_operations"`
	Status        string  `json:"status"`
	TotalDuration float64 `json:"total_duration"`
}

type Hailer struct {
	User      string `json:"user"`
	Boatswain string `json:"boatswain"`
	Hostname  string `json:"hostname"`
}

type HostData struct {
	TotalOps   int             `json:"total_operations"`
	ChangedOps int             `json:"changed_operations"`
	SuccessOps int             `json:"successful_operations"`
	FailedOps  int             `json:"failed_operations"`
	Duration   float64         `json:"duration"`
	History    []OperationData `json:"history"`
}

type OperationData struct {
	// Common / Setup Phase Fields
	Type      string  `json:"type,omitempty"`
	Stage 	  string  `json:"stage,omitempty"`
	Timestamp float64 `json:"timestamp,omitempty"`

	Name       string              `json:"operation,omitempty"`
	Args       *OperationArguments `json:"operation_arguments,omitempty"`
	Changed    bool                `json:"changed,omitempty"`
	Success    bool                `json:"success,omitempty"`
	Duration   float64             `json:"duration"`
	Stdout     string              `json:"stdout,omitempty"`
	Stderr     string              `json:"stderr,omitempty"`
	RetryStats *RetryStatistics    `json:"retry_statistics,omitempty"`
}

type OperationArguments struct {
	GlobalArguments map[string]string `json:"global_arguments"`
	OperationMeta   string            `json:"operation_meta"`
}

type RetryStatistics struct {
	Stdout     string      `json:"stdout"`
	Stderr     string      `json:"stderr"`
	Attempts   int         `json:"retry_attempts"`
	MaxRetries int         `json:"max_retries"`
	RetryInfo  interface{} `json:"retry_info"`
}

type HealthCheckEvent struct {
	Type      string      `json:"type"`
	Host      string      `json:"host"`
	Stage     string      `json:"stage"`
	Timestamp float64     `json:"timestamp"`
	Metrics   HostMetrics `json:"metrics"`
}

type HostMetrics struct {
	CPULoad1Min float64 `json:"cpu_load_1min"`
	CPULoad5Min float64 `json:"cpu_load_5min"`
	RAMPercent  float64 `json:"ram_percent"`
	RAMUsedGB   float64 `json:"ram_used_gb"`
	RAMTotalGB  float64 `json:"ram_total_gb"`
}

type OperationStats struct {
	Count           int     `json:"count"`
	TotalDuration   float64 `json:"total_duration"`
	AverageDuration float64 `json:"average_duration"`
	P50             float64 `json:"p50_duration"`
	P90             float64 `json:"p90_duration"`
	P95             float64 `json:"p95_duration"`
	P99             float64 `json:"p99_duration"`
}
