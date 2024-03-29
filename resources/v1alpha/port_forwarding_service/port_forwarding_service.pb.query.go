// Code generated by protoc-gen-goten-resource
// Resource: PortForwardingService
// DO NOT EDIT!!!

package port_forwarding_service

import (
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/cloudwan/goten-sdk/runtime/api/watch_type"
	"github.com/cloudwan/goten-sdk/runtime/goten"
	gotenobject "github.com/cloudwan/goten-sdk/runtime/object"
	gotenresource "github.com/cloudwan/goten-sdk/runtime/resource"
)

// proto imports
import (
	ntt_meta "github.com/cloudwan/edgelq-sdk/common/types/meta"
	iam_project "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha/project"
	access_point "github.com/cloudwan/ztna-sdk/resources/v1alpha/access_point"
)

// ensure the imports are used
var (
	_ = &timestamppb.Timestamp{}
	_ = watch_type.WatchType_STATELESS
	_ = new(goten.GotenMessage)
	_ = new(gotenobject.FieldMask)
	_ = new(gotenresource.ListQuery)
)

// make sure we're using proto imports
var (
	_ = &ntt_meta.Meta{}
	_ = &iam_project.Project{}
	_ = &access_point.AccessPoint{}
)

type GetQuery struct {
	Reference *Reference
	Mask      *PortForwardingService_FieldMask
}

func (q *GetQuery) GotenQuery() {}

func (q *GetQuery) String() string {
	return gotenresource.MakeSQLGetString(q)
}

func (q *GetQuery) GetResourceDescriptor() gotenresource.Descriptor {
	return descriptor
}

func (q *GetQuery) GetReference() gotenresource.Reference {
	return q.Reference
}

func (q *GetQuery) GetFieldMask() gotenobject.FieldMask {
	return q.Mask
}

func (q *GetQuery) SetReference(ref gotenresource.Reference) {
	if ref != nil {
		q.Reference = ref.(*Reference)
	} else {
		q.Reference = nil
	}
}

func (q *GetQuery) SetFieldMask(mask gotenobject.FieldMask) {
	if mask != nil {
		q.Mask = mask.(*PortForwardingService_FieldMask)
	} else {
		q.Mask = nil
	}
}

type ListQuery struct {
	Filter *Filter
	Pager  *PagerQuery
	Mask   *PortForwardingService_FieldMask
}

func (q *ListQuery) GotenQuery() {}

func (q *ListQuery) String() string {
	return gotenresource.MakeSQLListString(q)
}

func (q *ListQuery) GetResourceDescriptor() gotenresource.Descriptor {
	return descriptor
}

func (q *ListQuery) GetFilter() gotenresource.Filter {
	return q.Filter
}

func (q *ListQuery) GetPager() gotenresource.PagerQuery {
	return q.Pager
}

func (q *ListQuery) GetFieldMask() gotenobject.FieldMask {
	return q.Mask
}

func (q *ListQuery) SetFilter(filter gotenresource.Filter) {
	if filter != nil {
		q.Filter = filter.(*Filter)
	} else {
		q.Filter = nil
	}
}

func (q *ListQuery) SetPager(pager gotenresource.PagerQuery) {
	if pager != nil {
		q.Pager = pager.(*PagerQuery)
	} else {
		q.Pager = nil
	}
}

func (q *ListQuery) SetFieldMask(mask gotenobject.FieldMask) {
	if mask != nil {
		q.Mask = mask.(*PortForwardingService_FieldMask)
	} else {
		q.Mask = nil
	}
}

type WatchQuery struct {
	ListQuery
	WatchType    watch_type.WatchType
	ChunkSize    int
	ResumeToken  string
	StartingTime *timestamppb.Timestamp
}

func (q *WatchQuery) String() string {
	return gotenresource.MakeSQLWatchString(q)
}

func (q *WatchQuery) GetWatchType() watch_type.WatchType {
	return q.WatchType
}

func (q *WatchQuery) GetMaximumChunkSize() int {
	return q.ChunkSize
}

func (q *WatchQuery) GetResumeToken() string {
	return q.ResumeToken
}

func (q *WatchQuery) GetStartingTime() *timestamppb.Timestamp {
	return q.StartingTime
}

func (q *WatchQuery) SetWatchType(watchType watch_type.WatchType) {
	q.WatchType = watchType
}

func (q *WatchQuery) SetMaximumChunkSize(chunkSize int) {
	q.ChunkSize = chunkSize
}

func (q *WatchQuery) SetResumeToken(token string) {
	q.ResumeToken = token
}

func (q *WatchQuery) SetStartingTime(startingTime *timestamppb.Timestamp) {
	q.StartingTime = startingTime
}

type QueryResultSnapshot struct {
	PortForwardingServices []*PortForwardingService
	PrevPageCursor         *PagerCursor
	NextPageCursor         *PagerCursor
}

func (qr *QueryResultSnapshot) GetResults() gotenresource.ResourceList {
	return PortForwardingServiceList(qr.PortForwardingServices)
}

func (qr *QueryResultSnapshot) GetNextPageCursor() gotenresource.Cursor {
	return qr.NextPageCursor
}

func (qr *QueryResultSnapshot) GetPrevPageCursor() gotenresource.Cursor {
	return qr.PrevPageCursor
}

func (qr *QueryResultSnapshot) SetResults(results gotenresource.ResourceList) {
	if results != nil {
		qr.PortForwardingServices = results.(PortForwardingServiceList)
	} else {
		qr.PortForwardingServices = nil
	}
}

func (qr *QueryResultSnapshot) SetCursors(nextPageCursor, prevPageCursor gotenresource.Cursor) {
	if nextPageCursor != nil {
		qr.NextPageCursor = nextPageCursor.(*PagerCursor)
	} else {
		qr.NextPageCursor = nil
	}
	if prevPageCursor != nil {
		qr.PrevPageCursor = prevPageCursor.(*PagerCursor)
	} else {
		qr.PrevPageCursor = nil
	}
}

type QueryResultChange struct {
	Changes        []*PortForwardingServiceChange
	PrevPageCursor *PagerCursor
	NextPageCursor *PagerCursor
	ResumeToken    string
	IsCurrent      bool
	IsHardReset    bool
	IsSoftReset    bool
	SnapshotSize   int64
}

func (qr *QueryResultChange) GetResults() gotenresource.ResourceChangeList {
	return PortForwardingServiceChangeList(qr.Changes)
}

func (qr *QueryResultChange) GetNextPageCursor() gotenresource.Cursor {
	return qr.NextPageCursor
}

func (qr *QueryResultChange) GetPrevPageCursor() gotenresource.Cursor {
	return qr.PrevPageCursor
}

func (qr *QueryResultChange) GetIsCurrent() bool {
	return qr.IsCurrent
}

func (qr *QueryResultChange) GetIsHardReset() bool {
	return qr.IsHardReset
}

func (qr *QueryResultChange) GetIsSoftReset() bool {
	return qr.IsSoftReset
}

func (qr *QueryResultChange) GetSnapshotSize() int64 {
	return qr.SnapshotSize
}

func (qr *QueryResultChange) GetResumeToken() string {
	return qr.ResumeToken
}

func (qr *QueryResultChange) SetResults(results gotenresource.ResourceChangeList) {
	if results != nil {
		qr.Changes = results.(PortForwardingServiceChangeList)
	} else {
		qr.Changes = nil
	}
}

func (qr *QueryResultChange) SetCursors(nextPageCursor, prevPageCursor gotenresource.Cursor) {
	if nextPageCursor != nil {
		qr.NextPageCursor = nextPageCursor.(*PagerCursor)
	} else {
		qr.NextPageCursor = nil
	}
	if prevPageCursor != nil {
		qr.PrevPageCursor = prevPageCursor.(*PagerCursor)
	} else {
		qr.PrevPageCursor = nil
	}
}

func (qr *QueryResultChange) SetIsCurrent() {
	qr.IsCurrent = true
}

func (qr *QueryResultChange) SetIsSoftReset() {
	qr.IsSoftReset = true
}

func (qr *QueryResultChange) SetIsHardReset() {
	qr.IsHardReset = true
}

func (qr *QueryResultChange) SetSnapshotSize(size int64) {
	qr.SnapshotSize = size
}

func (qr *QueryResultChange) SetResumeToken(token string) {
	qr.ResumeToken = token
}
