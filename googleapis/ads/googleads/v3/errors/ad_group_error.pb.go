// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v3/errors/ad_group_error.proto

package errors

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Enum describing possible ad group errors.
type AdGroupErrorEnum_AdGroupError int32

const (
	// Enum unspecified.
	AdGroupErrorEnum_UNSPECIFIED AdGroupErrorEnum_AdGroupError = 0
	// The received error code is not known in this version.
	AdGroupErrorEnum_UNKNOWN AdGroupErrorEnum_AdGroupError = 1
	// AdGroup with the same name already exists for the campaign.
	AdGroupErrorEnum_DUPLICATE_ADGROUP_NAME AdGroupErrorEnum_AdGroupError = 2
	// AdGroup name is not valid.
	AdGroupErrorEnum_INVALID_ADGROUP_NAME AdGroupErrorEnum_AdGroupError = 3
	// Advertiser is not allowed to target sites or set site bids that are not
	// on the Google Search Network.
	AdGroupErrorEnum_ADVERTISER_NOT_ON_CONTENT_NETWORK AdGroupErrorEnum_AdGroupError = 5
	// Bid amount is too big.
	AdGroupErrorEnum_BID_TOO_BIG AdGroupErrorEnum_AdGroupError = 6
	// AdGroup bid does not match the campaign's bidding strategy.
	AdGroupErrorEnum_BID_TYPE_AND_BIDDING_STRATEGY_MISMATCH AdGroupErrorEnum_AdGroupError = 7
	// AdGroup name is required for Add.
	AdGroupErrorEnum_MISSING_ADGROUP_NAME AdGroupErrorEnum_AdGroupError = 8
	// No link found between the ad group and the label.
	AdGroupErrorEnum_ADGROUP_LABEL_DOES_NOT_EXIST AdGroupErrorEnum_AdGroupError = 9
	// The label has already been attached to the ad group.
	AdGroupErrorEnum_ADGROUP_LABEL_ALREADY_EXISTS AdGroupErrorEnum_AdGroupError = 10
	// The CriterionTypeGroup is not supported for the content bid dimension.
	AdGroupErrorEnum_INVALID_CONTENT_BID_CRITERION_TYPE_GROUP AdGroupErrorEnum_AdGroupError = 11
	// The ad group type is not compatible with the campaign channel type.
	AdGroupErrorEnum_AD_GROUP_TYPE_NOT_VALID_FOR_ADVERTISING_CHANNEL_TYPE AdGroupErrorEnum_AdGroupError = 12
	// The ad group type is not supported in the country of sale of the
	// campaign.
	AdGroupErrorEnum_ADGROUP_TYPE_NOT_SUPPORTED_FOR_CAMPAIGN_SALES_COUNTRY AdGroupErrorEnum_AdGroupError = 13
	// Ad groups of AdGroupType.SEARCH_DYNAMIC_ADS can only be added to
	// campaigns that have DynamicSearchAdsSetting attached.
	AdGroupErrorEnum_CANNOT_ADD_ADGROUP_OF_TYPE_DSA_TO_CAMPAIGN_WITHOUT_DSA_SETTING AdGroupErrorEnum_AdGroupError = 14
)

var AdGroupErrorEnum_AdGroupError_name = map[int32]string{
	0:  "UNSPECIFIED",
	1:  "UNKNOWN",
	2:  "DUPLICATE_ADGROUP_NAME",
	3:  "INVALID_ADGROUP_NAME",
	5:  "ADVERTISER_NOT_ON_CONTENT_NETWORK",
	6:  "BID_TOO_BIG",
	7:  "BID_TYPE_AND_BIDDING_STRATEGY_MISMATCH",
	8:  "MISSING_ADGROUP_NAME",
	9:  "ADGROUP_LABEL_DOES_NOT_EXIST",
	10: "ADGROUP_LABEL_ALREADY_EXISTS",
	11: "INVALID_CONTENT_BID_CRITERION_TYPE_GROUP",
	12: "AD_GROUP_TYPE_NOT_VALID_FOR_ADVERTISING_CHANNEL_TYPE",
	13: "ADGROUP_TYPE_NOT_SUPPORTED_FOR_CAMPAIGN_SALES_COUNTRY",
	14: "CANNOT_ADD_ADGROUP_OF_TYPE_DSA_TO_CAMPAIGN_WITHOUT_DSA_SETTING",
}

var AdGroupErrorEnum_AdGroupError_value = map[string]int32{
	"UNSPECIFIED":                                                    0,
	"UNKNOWN":                                                        1,
	"DUPLICATE_ADGROUP_NAME":                                         2,
	"INVALID_ADGROUP_NAME":                                           3,
	"ADVERTISER_NOT_ON_CONTENT_NETWORK":                              5,
	"BID_TOO_BIG":                                                    6,
	"BID_TYPE_AND_BIDDING_STRATEGY_MISMATCH":                         7,
	"MISSING_ADGROUP_NAME":                                           8,
	"ADGROUP_LABEL_DOES_NOT_EXIST":                                   9,
	"ADGROUP_LABEL_ALREADY_EXISTS":                                   10,
	"INVALID_CONTENT_BID_CRITERION_TYPE_GROUP":                       11,
	"AD_GROUP_TYPE_NOT_VALID_FOR_ADVERTISING_CHANNEL_TYPE":           12,
	"ADGROUP_TYPE_NOT_SUPPORTED_FOR_CAMPAIGN_SALES_COUNTRY":          13,
	"CANNOT_ADD_ADGROUP_OF_TYPE_DSA_TO_CAMPAIGN_WITHOUT_DSA_SETTING": 14,
}

func (x AdGroupErrorEnum_AdGroupError) String() string {
	return proto.EnumName(AdGroupErrorEnum_AdGroupError_name, int32(x))
}

func (AdGroupErrorEnum_AdGroupError) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e609153f44ae1ab8, []int{0, 0}
}

// Container for enum describing possible ad group errors.
type AdGroupErrorEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AdGroupErrorEnum) Reset()         { *m = AdGroupErrorEnum{} }
func (m *AdGroupErrorEnum) String() string { return proto.CompactTextString(m) }
func (*AdGroupErrorEnum) ProtoMessage()    {}
func (*AdGroupErrorEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_e609153f44ae1ab8, []int{0}
}

func (m *AdGroupErrorEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AdGroupErrorEnum.Unmarshal(m, b)
}
func (m *AdGroupErrorEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AdGroupErrorEnum.Marshal(b, m, deterministic)
}
func (m *AdGroupErrorEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AdGroupErrorEnum.Merge(m, src)
}
func (m *AdGroupErrorEnum) XXX_Size() int {
	return xxx_messageInfo_AdGroupErrorEnum.Size(m)
}
func (m *AdGroupErrorEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_AdGroupErrorEnum.DiscardUnknown(m)
}

var xxx_messageInfo_AdGroupErrorEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v3.errors.AdGroupErrorEnum_AdGroupError", AdGroupErrorEnum_AdGroupError_name, AdGroupErrorEnum_AdGroupError_value)
	proto.RegisterType((*AdGroupErrorEnum)(nil), "google.ads.googleads.v3.errors.AdGroupErrorEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v3/errors/ad_group_error.proto", fileDescriptor_e609153f44ae1ab8)
}

var fileDescriptor_e609153f44ae1ab8 = []byte{
	// 553 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0xd1, 0x6a, 0xd4, 0x40,
	0x14, 0x86, 0xed, 0xb6, 0xb6, 0x3a, 0xad, 0x3a, 0x0e, 0x22, 0x52, 0x4a, 0xd1, 0x05, 0x45, 0x44,
	0xb2, 0x17, 0x51, 0xd0, 0x08, 0xc2, 0x49, 0x66, 0x9a, 0x0e, 0xcd, 0xce, 0x84, 0xcc, 0x64, 0xeb,
	0xca, 0xc2, 0xb0, 0x9a, 0x12, 0x0a, 0x6d, 0xb2, 0x24, 0x6d, 0x5f, 0xc2, 0x97, 0x10, 0x2f, 0x7d,
	0x14, 0x1f, 0x45, 0x7c, 0x08, 0x99, 0x8c, 0xd9, 0xb6, 0x88, 0x5e, 0xe5, 0xcc, 0x99, 0xef, 0x3f,
	0xe7, 0x4f, 0xf2, 0x23, 0xbf, 0xac, 0xeb, 0xf2, 0xe4, 0x68, 0x34, 0x2f, 0xda, 0x91, 0x2b, 0x6d,
	0x75, 0xe1, 0x8f, 0x8e, 0x9a, 0xa6, 0x6e, 0xda, 0xd1, 0xbc, 0x30, 0x65, 0x53, 0x9f, 0x2f, 0x4c,
	0x77, 0xf6, 0x16, 0x4d, 0x7d, 0x56, 0x93, 0x5d, 0x47, 0x7a, 0xf3, 0xa2, 0xf5, 0x96, 0x22, 0xef,
	0xc2, 0xf7, 0x9c, 0x68, 0x7b, 0xa7, 0x1f, 0xba, 0x38, 0x1e, 0xcd, 0xab, 0xaa, 0x3e, 0x9b, 0x9f,
	0x1d, 0xd7, 0x55, 0xeb, 0xd4, 0xc3, 0xaf, 0x6b, 0x08, 0x43, 0x11, 0xdb, 0xa9, 0xcc, 0xf2, 0xac,
	0x3a, 0x3f, 0x1d, 0x7e, 0x59, 0x43, 0x5b, 0x57, 0x9b, 0xe4, 0x1e, 0xda, 0xcc, 0x85, 0x4a, 0x59,
	0xc4, 0xf7, 0x38, 0xa3, 0xf8, 0x06, 0xd9, 0x44, 0x1b, 0xb9, 0x38, 0x10, 0xf2, 0x50, 0xe0, 0x15,
	0xb2, 0x8d, 0x1e, 0xd2, 0x3c, 0x4d, 0x78, 0x04, 0x9a, 0x19, 0xa0, 0x71, 0x26, 0xf3, 0xd4, 0x08,
	0x18, 0x33, 0x3c, 0x20, 0x8f, 0xd0, 0x03, 0x2e, 0x26, 0x90, 0x70, 0x7a, 0xfd, 0x66, 0x95, 0x3c,
	0x45, 0x4f, 0x80, 0x4e, 0x58, 0xa6, 0xb9, 0x62, 0x99, 0x11, 0x52, 0x1b, 0x29, 0x4c, 0x24, 0x85,
	0x66, 0x42, 0x1b, 0xc1, 0xf4, 0xa1, 0xcc, 0x0e, 0xf0, 0x4d, 0xbb, 0x3a, 0xe4, 0xd4, 0x68, 0x29,
	0x4d, 0xc8, 0x63, 0xbc, 0x4e, 0x5e, 0xa0, 0x67, 0x5d, 0x63, 0x9a, 0x32, 0x03, 0x82, 0x9a, 0x90,
	0x53, 0xca, 0x45, 0x6c, 0x94, 0xce, 0x40, 0xb3, 0x78, 0x6a, 0xc6, 0x5c, 0x8d, 0x41, 0x47, 0xfb,
	0x78, 0xc3, 0x6e, 0x1f, 0x73, 0xa5, 0xec, 0xf5, 0xb5, 0xed, 0xb7, 0xc8, 0x63, 0xb4, 0xd3, 0x77,
	0x12, 0x08, 0x59, 0x62, 0xa8, 0x64, 0xaa, 0x73, 0xc1, 0x3e, 0x70, 0xa5, 0xf1, 0xed, 0xbf, 0x09,
	0x48, 0x32, 0x06, 0x74, 0xea, 0x00, 0x85, 0x11, 0x79, 0x89, 0x9e, 0xf7, 0xef, 0xd6, 0xfb, 0xb6,
	0xce, 0xa2, 0x8c, 0x6b, 0x96, 0x71, 0x29, 0x9c, 0xc7, 0x6e, 0x04, 0xde, 0x24, 0x6f, 0xd0, 0x2b,
	0xa0, 0xee, 0xe4, 0x2e, 0xec, 0x32, 0xa7, 0xde, 0x93, 0x99, 0xe9, 0xbf, 0x85, 0xf5, 0x1a, 0xed,
	0x83, 0x10, 0x2c, 0xe9, 0x30, 0xbc, 0x45, 0xde, 0xa2, 0xd7, 0xbd, 0x93, 0xa5, 0x50, 0xe5, 0x69,
	0x2a, 0x33, 0xcd, 0x9c, 0x38, 0x82, 0x71, 0x0a, 0x3c, 0x16, 0x46, 0x41, 0xc2, 0x94, 0x89, 0x64,
	0x2e, 0x74, 0x36, 0xc5, 0x77, 0x48, 0x88, 0xde, 0x47, 0x20, 0xac, 0x00, 0xe8, 0xe5, 0x1f, 0x90,
	0x7b, 0x6e, 0x10, 0x55, 0x60, 0xb4, 0xbc, 0x54, 0x1f, 0x72, 0xbd, 0x2f, 0x73, 0xdd, 0xf5, 0x15,
	0xd3, 0x9a, 0x8b, 0x18, 0xdf, 0x0d, 0x7f, 0xad, 0xa0, 0xe1, 0xe7, 0xfa, 0xd4, 0xfb, 0x7f, 0xce,
	0xc2, 0xfb, 0x57, 0x13, 0x93, 0xda, 0x70, 0xa5, 0x2b, 0x1f, 0xe9, 0x1f, 0x51, 0x59, 0x9f, 0xcc,
	0xab, 0xd2, 0xab, 0x9b, 0x72, 0x54, 0x1e, 0x55, 0x5d, 0xf4, 0xfa, 0x84, 0x2f, 0x8e, 0xdb, 0x7f,
	0x05, 0xfe, 0x9d, 0x7b, 0x7c, 0x1b, 0xac, 0xc6, 0x00, 0xdf, 0x07, 0xbb, 0xb1, 0x1b, 0x06, 0x45,
	0xeb, 0xb9, 0xd2, 0x56, 0x13, 0xdf, 0xeb, 0x56, 0xb6, 0x3f, 0x7a, 0x60, 0x06, 0x45, 0x3b, 0x5b,
	0x02, 0xb3, 0x89, 0x3f, 0x73, 0xc0, 0xcf, 0xc1, 0xd0, 0x75, 0x83, 0x00, 0x8a, 0x36, 0x08, 0x96,
	0x48, 0x10, 0x4c, 0xfc, 0x20, 0x70, 0xd0, 0xa7, 0xf5, 0xce, 0x9d, 0xff, 0x3b, 0x00, 0x00, 0xff,
	0xff, 0xf3, 0x1a, 0x86, 0xaa, 0x8d, 0x03, 0x00, 0x00,
}