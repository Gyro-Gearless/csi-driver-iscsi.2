/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: k8s.io/kubernetes/vendor/k8s.io/apimachinery/pkg/apis/meta/v1beta1/generated.proto

package v1beta1

import (
	fmt "fmt"

	io "io"

	proto "github.com/gogo/protobuf/proto"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	math "math"
	math_bits "math/bits"
	reflect "reflect"
	strings "strings"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal

var (
	_ = fmt.Errorf
	_ = math.Inf
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

func (m *PartialObjectMetadataList) Reset()      { *m = PartialObjectMetadataList{} }
func (*PartialObjectMetadataList) ProtoMessage() {}
func (*PartialObjectMetadataList) Descriptor() ([]byte, []int) {
	return fileDescriptor_90ec10f86b91f9a8, []int{0}
}

func (m *PartialObjectMetadataList) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *PartialObjectMetadataList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}

func (m *PartialObjectMetadataList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PartialObjectMetadataList.Merge(m, src)
}

func (m *PartialObjectMetadataList) XXX_Size() int {
	return m.Size()
}

func (m *PartialObjectMetadataList) XXX_DiscardUnknown() {
	xxx_messageInfo_PartialObjectMetadataList.DiscardUnknown(m)
}

var xxx_messageInfo_PartialObjectMetadataList proto.InternalMessageInfo

func init() {
	proto.RegisterType((*PartialObjectMetadataList)(nil), "k8s.io.apimachinery.pkg.apis.meta.v1beta1.PartialObjectMetadataList")
}

func init() {
	proto.RegisterFile("k8s.io/kubernetes/vendor/k8s.io/apimachinery/pkg/apis/meta/v1beta1/generated.proto", fileDescriptor_90ec10f86b91f9a8)
}

var fileDescriptor_90ec10f86b91f9a8 = []byte{
	// 317 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x41, 0x4b, 0xf3, 0x30,
	0x1c, 0xc6, 0x9b, 0xf7, 0x65, 0x30, 0x3a, 0x04, 0xd9, 0x69, 0xee, 0x90, 0x0d, 0x4f, 0xdb, 0xc1,
	0x84, 0x0d, 0x11, 0xc1, 0xdb, 0x6e, 0x82, 0x32, 0xd9, 0x51, 0x3c, 0x98, 0x76, 0x7f, 0xbb, 0x58,
	0xd3, 0x94, 0xe4, 0xdf, 0x81, 0x37, 0x3f, 0x82, 0x1f, 0x6b, 0xc7, 0x1d, 0x07, 0xc2, 0x70, 0xf5,
	0x8b, 0x48, 0xda, 0x2a, 0x32, 0x14, 0x7a, 0xeb, 0xf3, 0x94, 0xdf, 0x2f, 0x4f, 0x20, 0xfe, 0x2c,
	0x3e, 0xb7, 0x4c, 0x6a, 0x1e, 0x67, 0x01, 0x98, 0x04, 0x10, 0x2c, 0x5f, 0x42, 0x32, 0xd7, 0x86,
	0x57, 0x3f, 0x44, 0x2a, 0x95, 0x08, 0x17, 0x32, 0x01, 0xf3, 0xcc, 0xd3, 0x38, 0x72, 0x85, 0xe5,
	0x0a, 0x50, 0xf0, 0xe5, 0x28, 0x00, 0x14, 0x23, 0x1e, 0x41, 0x02, 0x46, 0x20, 0xcc, 0x59, 0x6a,
	0x34, 0xea, 0xf6, 0xb0, 0x44, 0xd9, 0x4f, 0x94, 0xa5, 0x71, 0xe4, 0x0a, 0xcb, 0x1c, 0xca, 0x2a,
	0xb4, 0x7b, 0x12, 0x49, 0x5c, 0x64, 0x01, 0x0b, 0xb5, 0xe2, 0x91, 0x8e, 0x34, 0x2f, 0x0c, 0x41,
	0xf6, 0x50, 0xa4, 0x22, 0x14, 0x5f, 0xa5, 0xb9, 0x7b, 0x5a, 0x67, 0xd4, 0xfe, 0x9e, 0xee, 0xd9,
	0x5f, 0x94, 0xc9, 0x12, 0x94, 0x0a, 0xb8, 0x0d, 0x17, 0xa0, 0xc4, 0x3e, 0x77, 0xfc, 0x46, 0xfc,
	0xa3, 0x1b, 0x61, 0x50, 0x8a, 0xa7, 0x69, 0xf0, 0x08, 0x21, 0x5e, 0x03, 0x8a, 0xb9, 0x40, 0x71,
	0x25, 0x2d, 0xb6, 0xef, 0xfc, 0xa6, 0xaa, 0x72, 0xe7, 0x5f, 0x9f, 0x0c, 0x5a, 0x63, 0xc6, 0xea,
	0x5c, 0x9c, 0x39, 0xda, 0x99, 0x26, 0x87, 0xab, 0x6d, 0xcf, 0xcb, 0xb7, 0xbd, 0xe6, 0x57, 0x33,
	0xfb, 0x36, 0xb6, 0xef, 0xfd, 0x86, 0x44, 0x50, 0xb6, 0x43, 0xfa, 0xff, 0x07, 0xad, 0xf1, 0x45,
	0x3d, 0xf5, 0xaf, 0x6b, 0x27, 0x07, 0xd5, 0x39, 0x8d, 0x4b, 0x67, 0x9c, 0x95, 0xe2, 0xc9, 0x74,
	0xb5, 0xa3, 0xde, 0x7a, 0x47, 0xbd, 0xcd, 0x8e, 0x7a, 0x2f, 0x39, 0x25, 0xab, 0x9c, 0x92, 0x75,
	0x4e, 0xc9, 0x26, 0xa7, 0xe4, 0x3d, 0xa7, 0xe4, 0xf5, 0x83, 0x7a, 0xb7, 0xc3, 0xda, 0xcf, 0xe0,
	0x33, 0x00, 0x00, 0xff, 0xff, 0x30, 0x97, 0x8b, 0x11, 0x4b, 0x02, 0x00, 0x00,
}

func (m *PartialObjectMetadataList) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PartialObjectMetadataList) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PartialObjectMetadataList) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.ListMeta.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenerated(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Items) > 0 {
		for iNdEx := len(m.Items) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Items[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenerated(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintGenerated(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenerated(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}

func (m *PartialObjectMetadataList) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Items) > 0 {
		for _, e := range m.Items {
			l = e.Size()
			n += 1 + l + sovGenerated(uint64(l))
		}
	}
	l = m.ListMeta.Size()
	n += 1 + l + sovGenerated(uint64(l))
	return n
}

func sovGenerated(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}

func sozGenerated(x uint64) (n int) {
	return sovGenerated(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}

func (this *PartialObjectMetadataList) String() string {
	if this == nil {
		return "nil"
	}
	repeatedStringForItems := "[]PartialObjectMetadata{"
	for _, f := range this.Items {
		repeatedStringForItems += fmt.Sprintf("%v", f) + ","
	}
	repeatedStringForItems += "}"
	s := strings.Join([]string{
		`&PartialObjectMetadataList{`,
		`Items:` + repeatedStringForItems + `,`,
		`ListMeta:` + strings.Replace(strings.Replace(fmt.Sprintf("%v", this.ListMeta), "ListMeta", "v1.ListMeta", 1), `&`, ``, 1) + `,`,
		`}`,
	}, "")
	return s
}

func valueToStringGenerated(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}

func (m *PartialObjectMetadataList) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: PartialObjectMetadataList: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PartialObjectMetadataList: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Items", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Items = append(m.Items, v1.PartialObjectMetadata{})
			if err := m.Items[len(m.Items)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ListMeta", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ListMeta.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenerated
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}

func skipGenerated(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthGenerated
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenerated
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenerated
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenerated        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenerated          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenerated = fmt.Errorf("proto: unexpected end of group")
)
