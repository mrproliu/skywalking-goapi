// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package query

import (
	"fmt"
	"io"
	"strconv"
)

type AlarmMessage struct {
	StartTime int64       `json:"startTime"`
	Scope     *Scope      `json:"scope"`
	ID        string      `json:"id"`
	Message   string      `json:"message"`
	Events    []*Event    `json:"events"`
	Tags      []*KeyValue `json:"tags"`
}

type AlarmTag struct {
	Key   string  `json:"key"`
	Value *string `json:"value"`
}

type AlarmTrend struct {
	NumOfAlarm []*int `json:"numOfAlarm"`
}

type Alarms struct {
	Msgs  []*AlarmMessage `json:"msgs"`
	Total int             `json:"total"`
}

type Attribute struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type BasicTrace struct {
	SegmentID     string   `json:"segmentId"`
	EndpointNames []string `json:"endpointNames"`
	Duration      int      `json:"duration"`
	Start         string   `json:"start"`
	IsError       *bool    `json:"isError"`
	TraceIds      []string `json:"traceIds"`
}

type BatchMetricConditions struct {
	Name string   `json:"name"`
	Ids  []string `json:"ids"`
}

type BrowserErrorLog struct {
	Service            string        `json:"service"`
	ServiceVersion     string        `json:"serviceVersion"`
	Time               int64         `json:"time"`
	PagePath           string        `json:"pagePath"`
	Category           ErrorCategory `json:"category"`
	Grade              *string       `json:"grade"`
	Message            *string       `json:"message"`
	Line               *int          `json:"line"`
	Col                *int          `json:"col"`
	Stack              *string       `json:"stack"`
	ErrorURL           *string       `json:"errorUrl"`
	FirstReportedError bool          `json:"firstReportedError"`
}

type BrowserErrorLogQueryCondition struct {
	ServiceID        *string        `json:"serviceId"`
	ServiceVersionID *string        `json:"serviceVersionId"`
	PagePathID       *string        `json:"pagePathId"`
	Category         *ErrorCategory `json:"category"`
	QueryDuration    *Duration      `json:"queryDuration"`
	Paging           *Pagination    `json:"paging"`
}

type BrowserErrorLogs struct {
	Logs  []*BrowserErrorLog `json:"logs"`
	Total int                `json:"total"`
}

type Bucket struct {
	Min string `json:"min"`
	Max string `json:"max"`
}

type Call struct {
	Source           string        `json:"source"`
	SourceComponents []string      `json:"sourceComponents"`
	Target           string        `json:"target"`
	TargetComponents []string      `json:"targetComponents"`
	ID               string        `json:"id"`
	DetectPoints     []DetectPoint `json:"detectPoints"`
}

type ClusterBrief struct {
	NumOfService  int `json:"numOfService"`
	NumOfEndpoint int `json:"numOfEndpoint"`
	NumOfDatabase int `json:"numOfDatabase"`
	NumOfCache    int `json:"numOfCache"`
	NumOfMq       int `json:"numOfMQ"`
}

type DashboardConfiguration struct {
	Name          string       `json:"name"`
	Type          TemplateType `json:"type"`
	Configuration string       `json:"configuration"`
	Activated     bool         `json:"activated"`
	Disabled      bool         `json:"disabled"`
}

type DashboardSetting struct {
	Name          string       `json:"name"`
	Type          TemplateType `json:"type"`
	Configuration string       `json:"configuration"`
	Active        bool         `json:"active"`
}

type Database struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

type Duration struct {
	Start string `json:"start"`
	End   string `json:"end"`
	Step  Step   `json:"step"`
}

type Endpoint struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type EndpointInfo struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	ServiceID   string `json:"serviceId"`
	ServiceName string `json:"serviceName"`
}

type EndpointNode struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	ServiceID   string  `json:"serviceId"`
	ServiceName string  `json:"serviceName"`
	Type        *string `json:"type"`
	IsReal      bool    `json:"isReal"`
}

type EndpointTopology struct {
	Nodes []*EndpointNode `json:"nodes"`
	Calls []*Call         `json:"calls"`
}

type Entity struct {
	Scope                   Scope   `json:"scope"`
	ServiceName             *string `json:"serviceName"`
	Normal                  *bool   `json:"normal"`
	ServiceInstanceName     *string `json:"serviceInstanceName"`
	EndpointName            *string `json:"endpointName"`
	DestServiceName         *string `json:"destServiceName"`
	DestNormal              *bool   `json:"destNormal"`
	DestServiceInstanceName *string `json:"destServiceInstanceName"`
	DestEndpointName        *string `json:"destEndpointName"`
}

type Event struct {
	UUID       string      `json:"uuid"`
	Source     *Source     `json:"source"`
	Name       string      `json:"name"`
	Type       EventType   `json:"type"`
	Message    *string     `json:"message"`
	Parameters []*KeyValue `json:"parameters"`
	StartTime  int64       `json:"startTime"`
	EndTime    *int64      `json:"endTime"`
}

type EventQueryCondition struct {
	UUID   *string      `json:"uuid"`
	Source *SourceInput `json:"source"`
	Name   *string      `json:"name"`
	Type   *EventType   `json:"type"`
	Time   *Duration    `json:"time"`
	Order  *Order       `json:"order"`
	Paging *Pagination  `json:"paging"`
}

type Events struct {
	Events []*Event `json:"events"`
	Total  int      `json:"total"`
}

type HealthStatus struct {
	Score   int     `json:"score"`
	Details *string `json:"details"`
}

type HeatMap struct {
	Values  []*HeatMapColumn `json:"values"`
	Buckets []*Bucket        `json:"buckets"`
}

type HeatMapColumn struct {
	ID     string  `json:"id"`
	Values []int64 `json:"values"`
}

type IntValues struct {
	Values []*KVInt `json:"values"`
}

type KVInt struct {
	ID    string `json:"id"`
	Value int64  `json:"value"`
}

type KeyValue struct {
	Key   string  `json:"key"`
	Value *string `json:"value"`
}

type Log struct {
	ServiceName         *string     `json:"serviceName"`
	ServiceID           *string     `json:"serviceId"`
	ServiceInstanceName *string     `json:"serviceInstanceName"`
	ServiceInstanceID   *string     `json:"serviceInstanceId"`
	EndpointName        *string     `json:"endpointName"`
	EndpointID          *string     `json:"endpointId"`
	TraceID             *string     `json:"traceId"`
	Timestamp           int64       `json:"timestamp"`
	ContentType         ContentType `json:"contentType"`
	Content             *string     `json:"content"`
	Tags                []*KeyValue `json:"tags"`
}

type LogEntity struct {
	Time int64       `json:"time"`
	Data []*KeyValue `json:"data"`
}

type LogQueryCondition struct {
	ServiceID                  *string              `json:"serviceId"`
	ServiceInstanceID          *string              `json:"serviceInstanceId"`
	EndpointID                 *string              `json:"endpointId"`
	RelatedTrace               *TraceScopeCondition `json:"relatedTrace"`
	QueryDuration              *Duration            `json:"queryDuration"`
	Paging                     *Pagination          `json:"paging"`
	Tags                       []*LogTag            `json:"tags"`
	KeywordsOfContent          []string             `json:"keywordsOfContent"`
	ExcludingKeywordsOfContent []string             `json:"excludingKeywordsOfContent"`
	QueryOrder                 *Order               `json:"queryOrder"`
}

type LogTag struct {
	Key   string  `json:"key"`
	Value *string `json:"value"`
}

type LogTestMetrics struct {
	Name      string      `json:"name"`
	Tags      []*KeyValue `json:"tags"`
	Value     int64       `json:"value"`
	Timestamp int64       `json:"timestamp"`
}

type LogTestRequest struct {
	Log string `json:"log"`
	Dsl string `json:"dsl"`
}

type LogTestResponse struct {
	Log     *Log              `json:"log"`
	Metrics []*LogTestMetrics `json:"metrics"`
}

type Logs struct {
	Logs  []*Log `json:"logs"`
	Total int    `json:"total"`
}

type MetricCondition struct {
	Name string  `json:"name"`
	ID   *string `json:"id"`
}

type MetricDefinition struct {
	Name    string      `json:"name"`
	Type    MetricsType `json:"type"`
	Catalog *string     `json:"catalog"`
}

type MetricsCondition struct {
	Name   string  `json:"name"`
	Entity *Entity `json:"entity"`
}

type MetricsValues struct {
	Label  *string    `json:"label"`
	Values *IntValues `json:"values"`
}

type Node struct {
	ID     string  `json:"id"`
	Name   string  `json:"name"`
	Type   *string `json:"type"`
	IsReal bool    `json:"isReal"`
}

type Pagination struct {
	PageNum   *int  `json:"pageNum"`
	PageSize  int   `json:"pageSize"`
	NeedTotal *bool `json:"needTotal"`
}

type ProfileAnalyzation struct {
	Tip   *string             `json:"tip"`
	Trees []*ProfileStackTree `json:"trees"`
}

type ProfileAnalyzeTimeRange struct {
	Start int64 `json:"start"`
	End   int64 `json:"end"`
}

type ProfileStackElement struct {
	ID                    string `json:"id"`
	ParentID              string `json:"parentId"`
	CodeSignature         string `json:"codeSignature"`
	Duration              int    `json:"duration"`
	DurationChildExcluded int    `json:"durationChildExcluded"`
	Count                 int    `json:"count"`
}

type ProfileStackTree struct {
	Elements []*ProfileStackElement `json:"elements"`
}

type ProfileTask struct {
	ID                   string            `json:"id"`
	ServiceID            string            `json:"serviceId"`
	ServiceName          string            `json:"serviceName"`
	EndpointName         string            `json:"endpointName"`
	StartTime            int64             `json:"startTime"`
	Duration             int               `json:"duration"`
	MinDurationThreshold int               `json:"minDurationThreshold"`
	DumpPeriod           int               `json:"dumpPeriod"`
	MaxSamplingCount     int               `json:"maxSamplingCount"`
	Logs                 []*ProfileTaskLog `json:"logs"`
}

type ProfileTaskCreationRequest struct {
	ServiceID            string `json:"serviceId"`
	EndpointName         string `json:"endpointName"`
	StartTime            *int64 `json:"startTime"`
	Duration             int    `json:"duration"`
	MinDurationThreshold int    `json:"minDurationThreshold"`
	DumpPeriod           int    `json:"dumpPeriod"`
	MaxSamplingCount     int    `json:"maxSamplingCount"`
}

type ProfileTaskCreationResult struct {
	ErrorReason *string `json:"errorReason"`
	ID          *string `json:"id"`
}

type ProfileTaskLog struct {
	ID            string                      `json:"id"`
	InstanceID    string                      `json:"instanceId"`
	InstanceName  string                      `json:"instanceName"`
	OperationType ProfileTaskLogOperationType `json:"operationType"`
	OperationTime int64                       `json:"operationTime"`
}

type ProfiledSegment struct {
	Spans []*ProfiledSpan `json:"spans"`
}

type ProfiledSpan struct {
	SpanID              int          `json:"spanId"`
	ParentSpanID        int          `json:"parentSpanId"`
	ServiceCode         string       `json:"serviceCode"`
	ServiceInstanceName string       `json:"serviceInstanceName"`
	StartTime           int64        `json:"startTime"`
	EndTime             int64        `json:"endTime"`
	EndpointName        *string      `json:"endpointName"`
	Type                string       `json:"type"`
	Peer                *string      `json:"peer"`
	Component           *string      `json:"component"`
	IsError             *bool        `json:"isError"`
	Layer               *string      `json:"layer"`
	Tags                []*KeyValue  `json:"tags"`
	Logs                []*LogEntity `json:"logs"`
}

type Ref struct {
	TraceID         string  `json:"traceId"`
	ParentSegmentID string  `json:"parentSegmentId"`
	ParentSpanID    int     `json:"parentSpanId"`
	Type            RefType `json:"type"`
}

type SelectedRecord struct {
	Name  string  `json:"name"`
	ID    string  `json:"id"`
	Value *string `json:"value"`
	RefID *string `json:"refId"`
}

type Service struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Group string `json:"group"`
}

type ServiceInstance struct {
	ID           string       `json:"id"`
	Name         string       `json:"name"`
	Attributes   []*Attribute `json:"attributes"`
	Language     Language     `json:"language"`
	InstanceUUID string       `json:"instanceUUID"`
}

type ServiceInstanceNode struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	ServiceID   string  `json:"serviceId"`
	ServiceName string  `json:"serviceName"`
	Type        *string `json:"type"`
	IsReal      bool    `json:"isReal"`
}

type ServiceInstanceTopology struct {
	Nodes []*ServiceInstanceNode `json:"nodes"`
	Calls []*Call                `json:"calls"`
}

type Source struct {
	Service         *string `json:"service"`
	ServiceInstance *string `json:"serviceInstance"`
	Endpoint        *string `json:"endpoint"`
}

type SourceInput struct {
	Service         *string `json:"service"`
	ServiceInstance *string `json:"serviceInstance"`
	Endpoint        *string `json:"endpoint"`
}

type Span struct {
	TraceID             string       `json:"traceId"`
	SegmentID           string       `json:"segmentId"`
	SpanID              int          `json:"spanId"`
	ParentSpanID        int          `json:"parentSpanId"`
	Refs                []*Ref       `json:"refs"`
	ServiceCode         string       `json:"serviceCode"`
	ServiceInstanceName string       `json:"serviceInstanceName"`
	StartTime           int64        `json:"startTime"`
	EndTime             int64        `json:"endTime"`
	EndpointName        *string      `json:"endpointName"`
	Type                string       `json:"type"`
	Peer                *string      `json:"peer"`
	Component           *string      `json:"component"`
	IsError             *bool        `json:"isError"`
	Layer               *string      `json:"layer"`
	Tags                []*KeyValue  `json:"tags"`
	Logs                []*LogEntity `json:"logs"`
}

type SpanTag struct {
	Key   string  `json:"key"`
	Value *string `json:"value"`
}

type TemplateChangeStatus struct {
	Status  bool    `json:"status"`
	Message *string `json:"message"`
}

type Thermodynamic struct {
	Nodes     [][]*int `json:"nodes"`
	AxisYStep int      `json:"axisYStep"`
}

type TimeInfo struct {
	Timezone         *string `json:"timezone"`
	CurrentTimestamp *int64  `json:"currentTimestamp"`
}

type TopNCondition struct {
	Name          string  `json:"name"`
	ParentService *string `json:"parentService"`
	Normal        *bool   `json:"normal"`
	Scope         *Scope  `json:"scope"`
	TopN          int     `json:"topN"`
	Order         Order   `json:"order"`
}

type TopNEntity struct {
	Name  string `json:"name"`
	ID    string `json:"id"`
	Value int64  `json:"value"`
}

type TopNRecord struct {
	Statement *string `json:"statement"`
	Latency   int64   `json:"latency"`
	TraceID   *string `json:"traceId"`
}

type TopNRecordsCondition struct {
	ServiceID  string    `json:"serviceId"`
	MetricName string    `json:"metricName"`
	TopN       int       `json:"topN"`
	Order      Order     `json:"order"`
	Duration   *Duration `json:"duration"`
}

type Topology struct {
	Nodes []*Node `json:"nodes"`
	Calls []*Call `json:"calls"`
}

type Trace struct {
	Spans []*Span `json:"spans"`
}

type TraceBrief struct {
	Traces []*BasicTrace `json:"traces"`
	Total  int           `json:"total"`
}

type TraceQueryCondition struct {
	ServiceID         *string     `json:"serviceId"`
	ServiceInstanceID *string     `json:"serviceInstanceId"`
	TraceID           *string     `json:"traceId"`
	EndpointID        *string     `json:"endpointId"`
	QueryDuration     *Duration   `json:"queryDuration"`
	MinTraceDuration  *int        `json:"minTraceDuration"`
	MaxTraceDuration  *int        `json:"maxTraceDuration"`
	TraceState        TraceState  `json:"traceState"`
	QueryOrder        QueryOrder  `json:"queryOrder"`
	Tags              []*SpanTag  `json:"tags"`
	Paging            *Pagination `json:"paging"`
}

type TraceScopeCondition struct {
	TraceID   string  `json:"traceId"`
	SegmentID *string `json:"segmentId"`
	SpanID    *int    `json:"spanId"`
}

type ContentType string

const (
	ContentTypeText ContentType = "TEXT"
	ContentTypeJSON ContentType = "JSON"
	ContentTypeYaml ContentType = "YAML"
)

var AllContentType = []ContentType{
	ContentTypeText,
	ContentTypeJSON,
	ContentTypeYaml,
}

func (e ContentType) IsValid() bool {
	switch e {
	case ContentTypeText, ContentTypeJSON, ContentTypeYaml:
		return true
	}
	return false
}

func (e ContentType) String() string {
	return string(e)
}

func (e *ContentType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ContentType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ContentType", str)
	}
	return nil
}

func (e ContentType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type DetectPoint string

const (
	DetectPointClient DetectPoint = "CLIENT"
	DetectPointServer DetectPoint = "SERVER"
	DetectPointProxy  DetectPoint = "PROXY"
)

var AllDetectPoint = []DetectPoint{
	DetectPointClient,
	DetectPointServer,
	DetectPointProxy,
}

func (e DetectPoint) IsValid() bool {
	switch e {
	case DetectPointClient, DetectPointServer, DetectPointProxy:
		return true
	}
	return false
}

func (e DetectPoint) String() string {
	return string(e)
}

func (e *DetectPoint) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = DetectPoint(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid DetectPoint", str)
	}
	return nil
}

func (e DetectPoint) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type ErrorCategory string

const (
	ErrorCategoryAll      ErrorCategory = "ALL"
	ErrorCategoryAjax     ErrorCategory = "AJAX"
	ErrorCategoryResource ErrorCategory = "RESOURCE"
	ErrorCategoryVue      ErrorCategory = "VUE"
	ErrorCategoryPromise  ErrorCategory = "PROMISE"
	ErrorCategoryJs       ErrorCategory = "JS"
	ErrorCategoryUnknown  ErrorCategory = "UNKNOWN"
)

var AllErrorCategory = []ErrorCategory{
	ErrorCategoryAll,
	ErrorCategoryAjax,
	ErrorCategoryResource,
	ErrorCategoryVue,
	ErrorCategoryPromise,
	ErrorCategoryJs,
	ErrorCategoryUnknown,
}

func (e ErrorCategory) IsValid() bool {
	switch e {
	case ErrorCategoryAll, ErrorCategoryAjax, ErrorCategoryResource, ErrorCategoryVue, ErrorCategoryPromise, ErrorCategoryJs, ErrorCategoryUnknown:
		return true
	}
	return false
}

func (e ErrorCategory) String() string {
	return string(e)
}

func (e *ErrorCategory) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ErrorCategory(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ErrorCategory", str)
	}
	return nil
}

func (e ErrorCategory) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type EventType string

const (
	EventTypeNormal EventType = "Normal"
	EventTypeError  EventType = "Error"
)

var AllEventType = []EventType{
	EventTypeNormal,
	EventTypeError,
}

func (e EventType) IsValid() bool {
	switch e {
	case EventTypeNormal, EventTypeError:
		return true
	}
	return false
}

func (e EventType) String() string {
	return string(e)
}

func (e *EventType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = EventType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid EventType", str)
	}
	return nil
}

func (e EventType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type Language string

const (
	LanguageUnknown Language = "UNKNOWN"
	LanguageJava    Language = "JAVA"
	LanguageDotnet  Language = "DOTNET"
	LanguageNodejs  Language = "NODEJS"
	LanguagePython  Language = "PYTHON"
	LanguageRuby    Language = "RUBY"
	LanguageGo      Language = "GO"
	LanguageLua     Language = "LUA"
	LanguagePhp     Language = "PHP"
)

var AllLanguage = []Language{
	LanguageUnknown,
	LanguageJava,
	LanguageDotnet,
	LanguageNodejs,
	LanguagePython,
	LanguageRuby,
	LanguageGo,
	LanguageLua,
	LanguagePhp,
}

func (e Language) IsValid() bool {
	switch e {
	case LanguageUnknown, LanguageJava, LanguageDotnet, LanguageNodejs, LanguagePython, LanguageRuby, LanguageGo, LanguageLua, LanguagePhp:
		return true
	}
	return false
}

func (e Language) String() string {
	return string(e)
}

func (e *Language) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Language(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Language", str)
	}
	return nil
}

func (e Language) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type MetricsType string

const (
	MetricsTypeUnknown       MetricsType = "UNKNOWN"
	MetricsTypeRegularValue  MetricsType = "REGULAR_VALUE"
	MetricsTypeLabeledValue  MetricsType = "LABELED_VALUE"
	MetricsTypeHeatmap       MetricsType = "HEATMAP"
	MetricsTypeSampledRecord MetricsType = "SAMPLED_RECORD"
)

var AllMetricsType = []MetricsType{
	MetricsTypeUnknown,
	MetricsTypeRegularValue,
	MetricsTypeLabeledValue,
	MetricsTypeHeatmap,
	MetricsTypeSampledRecord,
}

func (e MetricsType) IsValid() bool {
	switch e {
	case MetricsTypeUnknown, MetricsTypeRegularValue, MetricsTypeLabeledValue, MetricsTypeHeatmap, MetricsTypeSampledRecord:
		return true
	}
	return false
}

func (e MetricsType) String() string {
	return string(e)
}

func (e *MetricsType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = MetricsType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid MetricsType", str)
	}
	return nil
}

func (e MetricsType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type NodeType string

const (
	NodeTypeService  NodeType = "SERVICE"
	NodeTypeEndpoint NodeType = "ENDPOINT"
	NodeTypeUser     NodeType = "USER"
)

var AllNodeType = []NodeType{
	NodeTypeService,
	NodeTypeEndpoint,
	NodeTypeUser,
}

func (e NodeType) IsValid() bool {
	switch e {
	case NodeTypeService, NodeTypeEndpoint, NodeTypeUser:
		return true
	}
	return false
}

func (e NodeType) String() string {
	return string(e)
}

func (e *NodeType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = NodeType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid NodeType", str)
	}
	return nil
}

func (e NodeType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type Order string

const (
	OrderAsc Order = "ASC"
	OrderDes Order = "DES"
)

var AllOrder = []Order{
	OrderAsc,
	OrderDes,
}

func (e Order) IsValid() bool {
	switch e {
	case OrderAsc, OrderDes:
		return true
	}
	return false
}

func (e Order) String() string {
	return string(e)
}

func (e *Order) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Order(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Order", str)
	}
	return nil
}

func (e Order) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type ProfileTaskLogOperationType string

const (
	ProfileTaskLogOperationTypeNotified          ProfileTaskLogOperationType = "NOTIFIED"
	ProfileTaskLogOperationTypeExecutionFinished ProfileTaskLogOperationType = "EXECUTION_FINISHED"
)

var AllProfileTaskLogOperationType = []ProfileTaskLogOperationType{
	ProfileTaskLogOperationTypeNotified,
	ProfileTaskLogOperationTypeExecutionFinished,
}

func (e ProfileTaskLogOperationType) IsValid() bool {
	switch e {
	case ProfileTaskLogOperationTypeNotified, ProfileTaskLogOperationTypeExecutionFinished:
		return true
	}
	return false
}

func (e ProfileTaskLogOperationType) String() string {
	return string(e)
}

func (e *ProfileTaskLogOperationType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ProfileTaskLogOperationType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ProfileTaskLogOperationType", str)
	}
	return nil
}

func (e ProfileTaskLogOperationType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type QueryOrder string

const (
	QueryOrderByStartTime QueryOrder = "BY_START_TIME"
	QueryOrderByDuration  QueryOrder = "BY_DURATION"
)

var AllQueryOrder = []QueryOrder{
	QueryOrderByStartTime,
	QueryOrderByDuration,
}

func (e QueryOrder) IsValid() bool {
	switch e {
	case QueryOrderByStartTime, QueryOrderByDuration:
		return true
	}
	return false
}

func (e QueryOrder) String() string {
	return string(e)
}

func (e *QueryOrder) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = QueryOrder(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid QueryOrder", str)
	}
	return nil
}

func (e QueryOrder) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type RefType string

const (
	RefTypeCrossProcess RefType = "CROSS_PROCESS"
	RefTypeCrossThread  RefType = "CROSS_THREAD"
)

var AllRefType = []RefType{
	RefTypeCrossProcess,
	RefTypeCrossThread,
}

func (e RefType) IsValid() bool {
	switch e {
	case RefTypeCrossProcess, RefTypeCrossThread:
		return true
	}
	return false
}

func (e RefType) String() string {
	return string(e)
}

func (e *RefType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = RefType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid RefType", str)
	}
	return nil
}

func (e RefType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type Scope string

const (
	ScopeAll                     Scope = "All"
	ScopeService                 Scope = "Service"
	ScopeServiceInstance         Scope = "ServiceInstance"
	ScopeEndpoint                Scope = "Endpoint"
	ScopeServiceRelation         Scope = "ServiceRelation"
	ScopeServiceInstanceRelation Scope = "ServiceInstanceRelation"
	ScopeEndpointRelation        Scope = "EndpointRelation"
)

var AllScope = []Scope{
	ScopeAll,
	ScopeService,
	ScopeServiceInstance,
	ScopeEndpoint,
	ScopeServiceRelation,
	ScopeServiceInstanceRelation,
	ScopeEndpointRelation,
}

func (e Scope) IsValid() bool {
	switch e {
	case ScopeAll, ScopeService, ScopeServiceInstance, ScopeEndpoint, ScopeServiceRelation, ScopeServiceInstanceRelation, ScopeEndpointRelation:
		return true
	}
	return false
}

func (e Scope) String() string {
	return string(e)
}

func (e *Scope) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Scope(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Scope", str)
	}
	return nil
}

func (e Scope) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type Step string

const (
	StepDay    Step = "DAY"
	StepHour   Step = "HOUR"
	StepMinute Step = "MINUTE"
	StepSecond Step = "SECOND"
)

var AllStep = []Step{
	StepDay,
	StepHour,
	StepMinute,
	StepSecond,
}

func (e Step) IsValid() bool {
	switch e {
	case StepDay, StepHour, StepMinute, StepSecond:
		return true
	}
	return false
}

func (e Step) String() string {
	return string(e)
}

func (e *Step) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Step(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Step", str)
	}
	return nil
}

func (e Step) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type TemplateType string

const (
	TemplateTypeDashboard                       TemplateType = "DASHBOARD"
	TemplateTypeTopologyService                 TemplateType = "TOPOLOGY_SERVICE"
	TemplateTypeTopologyInstance                TemplateType = "TOPOLOGY_INSTANCE"
	TemplateTypeTopologyEndpoint                TemplateType = "TOPOLOGY_ENDPOINT"
	TemplateTypeTopologyServiceRelation         TemplateType = "TOPOLOGY_SERVICE_RELATION"
	TemplateTypeTopologyServiceInstanceRelation TemplateType = "TOPOLOGY_SERVICE_INSTANCE_RELATION"
	TemplateTypeTopologyEndpointRelation        TemplateType = "TOPOLOGY_ENDPOINT_RELATION"
)

var AllTemplateType = []TemplateType{
	TemplateTypeDashboard,
	TemplateTypeTopologyService,
	TemplateTypeTopologyInstance,
	TemplateTypeTopologyEndpoint,
	TemplateTypeTopologyServiceRelation,
	TemplateTypeTopologyServiceInstanceRelation,
	TemplateTypeTopologyEndpointRelation,
}

func (e TemplateType) IsValid() bool {
	switch e {
	case TemplateTypeDashboard, TemplateTypeTopologyService, TemplateTypeTopologyInstance, TemplateTypeTopologyEndpoint, TemplateTypeTopologyServiceRelation, TemplateTypeTopologyServiceInstanceRelation, TemplateTypeTopologyEndpointRelation:
		return true
	}
	return false
}

func (e TemplateType) String() string {
	return string(e)
}

func (e *TemplateType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = TemplateType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid TemplateType", str)
	}
	return nil
}

func (e TemplateType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type TraceState string

const (
	TraceStateAll     TraceState = "ALL"
	TraceStateSuccess TraceState = "SUCCESS"
	TraceStateError   TraceState = "ERROR"
)

var AllTraceState = []TraceState{
	TraceStateAll,
	TraceStateSuccess,
	TraceStateError,
}

func (e TraceState) IsValid() bool {
	switch e {
	case TraceStateAll, TraceStateSuccess, TraceStateError:
		return true
	}
	return false
}

func (e TraceState) String() string {
	return string(e)
}

func (e *TraceState) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = TraceState(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid TraceState", str)
	}
	return nil
}

func (e TraceState) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
