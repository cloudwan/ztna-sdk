// Code generated by protoc-gen-goten-object
// File: ztna/proto/v1alpha/access_point.proto
// DO NOT EDIT!!!

package access_point

import (
	"fmt"
	"sort"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	gotenobject "github.com/cloudwan/goten-sdk/runtime/object"
)

// proto imports
import (
	iam_project "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha/project"
)

// ensure the imports are used
var (
	_ = new(fmt.Stringer)
	_ = new(sort.Interface)

	_ = new(proto.Message)
	_ = fieldmaskpb.FieldMask{}

	_ = new(gotenobject.FieldPath)
)

// make sure we're using proto imports
var (
	_ = &iam_project.Project{}
)

func (o *AccessPoint) GotenObjectExt() {}

func (o *AccessPoint) MakeFullFieldMask() *AccessPoint_FieldMask {
	return FullAccessPoint_FieldMask()
}

func (o *AccessPoint) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullAccessPoint_FieldMask()
}

func (o *AccessPoint) MakeDiffFieldMask(other *AccessPoint) *AccessPoint_FieldMask {
	if o == nil && other == nil {
		return &AccessPoint_FieldMask{}
	}
	if o == nil || other == nil {
		return FullAccessPoint_FieldMask()
	}

	res := &AccessPoint_FieldMask{}
	if o.GetName().String() != other.GetName().String() {
		res.Paths = append(res.Paths, &AccessPoint_FieldTerminalPath{selector: AccessPoint_FieldPathSelectorName})
	}
	return res
}

func (o *AccessPoint) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*AccessPoint))
}

func (o *AccessPoint) Clone() *AccessPoint {
	if o == nil {
		return nil
	}
	result := &AccessPoint{}
	if o.Name == nil {
		result.Name = nil
	} else if data, err := o.Name.ProtoString(); err != nil {
		panic(err)
	} else {
		result.Name = &Name{}
		if err := result.Name.ParseProtoString(data); err != nil {
			panic(err)
		}
	}
	return result
}

func (o *AccessPoint) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *AccessPoint) Merge(source *AccessPoint) {
	if source.GetName() != nil {
		if data, err := source.GetName().ProtoString(); err != nil {
			panic(err)
		} else {
			o.Name = &Name{}
			if err := o.Name.ParseProtoString(data); err != nil {
				panic(err)
			}
		}
	} else {
		o.Name = nil
	}
}

func (o *AccessPoint) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*AccessPoint))
}
