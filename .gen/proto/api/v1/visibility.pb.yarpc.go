// The MIT License (MIT)

// Copyright (c) 2017-2020 Uber Technologies Inc.

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

// Code generated by protoc-gen-yarpc-go. DO NOT EDIT.
// source: uber/cadence/api/v1/visibility.proto

package apiv1

var yarpcFileDescriptorClosurec1b2132ef24217c8 = [][]byte{
	// uber/cadence/api/v1/visibility.proto
	[]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0xcd, 0x6e, 0xd3, 0x40,
		0x14, 0x85, 0x71, 0x68, 0x23, 0xb8, 0x29, 0x60, 0x0d, 0x82, 0x80, 0x2b, 0x08, 0xb2, 0x58, 0x54,
		0x2c, 0xc6, 0x4a, 0x59, 0x76, 0x81, 0x12, 0x3c, 0xa0, 0x11, 0x21, 0x09, 0x8e, 0x9b, 0x12, 0x36,
		0xd6, 0xd8, 0x9e, 0x86, 0x11, 0xb6, 0xc7, 0xb2, 0xc7, 0x6e, 0xfb, 0x14, 0xbc, 0x27, 0x4f, 0x81,
		0xfc, 0x87, 0x84, 0x70, 0xc5, 0xce, 0x3e, 0xf7, 0x9c, 0x4f, 0x73, 0x7f, 0xe0, 0x75, 0xe1, 0xf3,
		0xcc, 0x0a, 0x58, 0xc8, 0x93, 0x80, 0x5b, 0x2c, 0x15, 0x56, 0x39, 0xb5, 0x4a, 0x91, 0x0b, 0x5f,
		0x44, 0x42, 0xdd, 0xe0, 0x34, 0x93, 0x4a, 0xa2, 0xc7, 0x95, 0x0b, 0xb7, 0x2e, 0xcc, 0x52, 0x81,
		0xcb, 0xa9, 0x31, 0xd9, 0x4b, 0xb9, 0x8f, 0xb8, 0x55, 0x5b, 0xfc, 0xe2, 0xd2, 0x52, 0x22, 0xe6,
		0xb9, 0x62, 0x71, 0xda, 0xa4, 0x0c, 0xb3, 0x8f, 0x7d, 0x25, 0xb3, 0x1f, 0x97, 0x91, 0xbc, 0x6a,
		0x3c, 0xe6, 0x17, 0x18, 0x5f, 0xb4, 0x0a, 0xb9, 0xe6, 0x41, 0xa1, 0x84, 0x4c, 0x3e, 0x88, 0x48,
		0xf1, 0x0c, 0x4d, 0x60, 0xd4, 0x99, 0x3d, 0x11, 0x3e, 0xd3, 0x5e, 0x69, 0x27, 0xf7, 0x1d, 0xe8,
		0x24, 0x1a, 0xa2, 0x27, 0x30, 0xcc, 0x8a, 0xa4, 0xaa, 0x0d, 0xea, 0xda, 0x61, 0x56, 0x24, 0x34,
		0x34, 0x4f, 0x00, 0x75, 0x48, 0xf7, 0x26, 0xe5, 0x2d, 0x0d, 0xc1, 0x41, 0xc2, 0x62, 0xde, 0x62,
		0xea, 0x6f, 0xf3, 0xa7, 0x06, 0x8f, 0x36, 0x8a, 0x65, 0xca, 0x15, 0x71, 0xe7, 0x7b, 0x07, 0x0f,
		0x38, 0xcb, 0x22, 0xc1, 0x73, 0xe5, 0x55, 0x0d, 0xd5, 0x81, 0xd1, 0xa9, 0x81, 0x9b, 0x6e, 0x71,
		0xd7, 0x2d, 0x76, 0xbb, 0x6e, 0x9d, 0xa3, 0x2e, 0x50, 0x49, 0xe8, 0x0c, 0x46, 0x11, 0x53, 0x7f,
		0xe2, 0x83, 0xff, 0xc6, 0xa1, 0xb1, 0x57, 0x82, 0xb9, 0x83, 0xa3, 0x8d, 0x62, 0xaa, 0xc8, 0xdb,
		0xd7, 0x50, 0x18, 0xe6, 0xf5, 0x7f, 0xfd, 0x8c, 0x87, 0xa7, 0x53, 0xdc, 0xb3, 0x09, 0xfc, 0xcf,
		0x04, 0xdf, 0x47, 0x32, 0xe7, 0x0d, 0xc8, 0x69, 0x01, 0x6f, 0x7e, 0x69, 0xa0, 0xd3, 0x24, 0xe4,
		0xd7, 0x3c, 0xdc, 0xb2, 0xa8, 0xe0, 0xd5, 0x6c, 0xd0, 0x4b, 0x30, 0xe8, 0xd2, 0x26, 0x5f, 0x89,
		0xed, 0x6d, 0x67, 0x8b, 0x73, 0xe2, 0xb9, 0xbb, 0x35, 0xf1, 0xe8, 0x72, 0x3b, 0x5b, 0x50, 0x5b,
		0xbf, 0x83, 0x5e, 0xc0, 0xf3, 0x9e, 0xfa, 0xc6, 0x75, 0xe8, 0xf2, 0xa3, 0xae, 0xdd, 0x12, 0xff,
		0x44, 0x76, 0x17, 0x2b, 0xc7, 0xd6, 0x07, 0xc8, 0x80, 0xa7, 0xbd, 0x78, 0x57, 0xbf, 0x7b, 0x0b,
		0xda, 0x5e, 0x9d, 0xcf, 0x17, 0x44, 0x3f, 0x40, 0xc7, 0x30, 0xee, 0x29, 0xcf, 0x57, 0xab, 0x85,
		0x7e, 0x88, 0x26, 0x70, 0xdc, 0x97, 0x9d, 0xb9, 0xc4, 0xa5, 0x9f, 0x89, 0x3e, 0x9c, 0x6f, 0x61,
		0x1c, 0xc8, 0xb8, 0x6f, 0x58, 0xf3, 0x7b, 0xb3, 0x54, 0xac, 0xab, 0x2d, 0xac, 0xb5, 0x6f, 0xd6,
		0x5e, 0xa8, 0xef, 0x85, 0x8f, 0x03, 0x19, 0x5b, 0x7f, 0x1d, 0x2b, 0xde, 0xf3, 0xa4, 0x39, 0xec,
		0xf6, 0x6e, 0xcf, 0x58, 0x2a, 0xca, 0xa9, 0x3f, 0xac, 0xb5, 0xb7, 0xbf, 0x03, 0x00, 0x00, 0xff,
		0xff, 0xd7, 0x22, 0x00, 0x73, 0x37, 0x03, 0x00, 0x00,
	},
	// google/protobuf/timestamp.proto
	[]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4f, 0xcf, 0xcf, 0x4f,
		0xcf, 0x49, 0xd5, 0x2f, 0x28, 0xca, 0x2f, 0xc9, 0x4f, 0x2a, 0x4d, 0xd3, 0x2f, 0xc9, 0xcc, 0x4d,
		0x2d, 0x2e, 0x49, 0xcc, 0x2d, 0xd0, 0x03, 0x0b, 0x09, 0xf1, 0x43, 0x14, 0xe8, 0xc1, 0x14, 0x28,
		0x59, 0x73, 0x71, 0x86, 0xc0, 0xd4, 0x08, 0x49, 0x70, 0xb1, 0x17, 0xa7, 0x26, 0xe7, 0xe7, 0xa5,
		0x14, 0x4b, 0x30, 0x2a, 0x30, 0x6a, 0x30, 0x07, 0xc1, 0xb8, 0x42, 0x22, 0x5c, 0xac, 0x79, 0x89,
		0x79, 0xf9, 0xc5, 0x12, 0x4c, 0x0a, 0x8c, 0x1a, 0xac, 0x41, 0x10, 0x8e, 0x53, 0x2b, 0x23, 0x97,
		0x70, 0x72, 0x7e, 0xae, 0x1e, 0x9a, 0xa1, 0x4e, 0x7c, 0x70, 0x23, 0x03, 0x40, 0x42, 0x01, 0x8c,
		0x51, 0x46, 0x50, 0x25, 0xe9, 0xf9, 0x39, 0x89, 0x79, 0xe9, 0x7a, 0xf9, 0x45, 0xe9, 0x48, 0x6e,
		0xac, 0x2c, 0x48, 0x2d, 0xd6, 0xcf, 0xce, 0xcb, 0x2f, 0xcf, 0x43, 0xb8, 0xb7, 0x20, 0xe9, 0x07,
		0x23, 0xe3, 0x22, 0x26, 0x66, 0xf7, 0x00, 0xa7, 0x55, 0x4c, 0x72, 0xee, 0x10, 0xdd, 0x01, 0x50,
		0x2d, 0x7a, 0xe1, 0xa9, 0x39, 0x39, 0xde, 0x20, 0x0d, 0x21, 0x20, 0xbd, 0x49, 0x6c, 0x60, 0xb3,
		0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0xae, 0x65, 0xce, 0x7d, 0xff, 0x00, 0x00, 0x00,
	},
	// uber/cadence/api/v1/workflow.proto
	[]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x59, 0x4f, 0x6f, 0xdb, 0xca,
		0xf5, 0xfd, 0x49, 0xb2, 0x1d, 0xfb, 0xca, 0xb1, 0xe9, 0x71, 0x9c, 0x28, 0xff, 0x1d, 0xbd, 0x5f,
		0x12, 0x47, 0x7d, 0x91, 0x9f, 0x93, 0x97, 0x97, 0xe6, 0xa5, 0x69, 0x4a, 0x93, 0x74, 0xcc, 0x44,
		0xa6, 0xd4, 0x21, 0x15, 0xc7, 0x0f, 0x45, 0x09, 0x5a, 0x1a, 0xdb, 0x44, 0x24, 0x52, 0x10, 0x47,
		0x49, 0xbc, 0x2f, 0xd0, 0x75, 0x77, 0x45, 0xbb, 0xe9, 0xae, 0x9b, 0x02, 0x45, 0x3f, 0x40, 0x51,
		0xa0, 0x8b, 0xee, 0xba, 0xed, 0xb2, 0xfb, 0x7e, 0x8b, 0x62, 0x86, 0x43, 0x89, 0x92, 0x28, 0x51,
		0x69, 0x81, 0xd7, 0x9d, 0x79, 0x79, 0xce, 0xe1, 0x9d, 0x3b, 0xf7, 0x1e, 0x0e, 0x2d, 0x28, 0xf6,
		0x8e, 0x49, 0x77, 0xbb, 0xe1, 0x34, 0x89, 0xd7, 0x20, 0xdb, 0x4e, 0xc7, 0xdd, 0xfe, 0xb0, 0xb3,
		0xfd, 0xd1, 0xef, 0xbe, 0x3f, 0x69, 0xf9, 0x1f, 0xcb, 0x9d, 0xae, 0x4f, 0x7d, 0xb4, 0xce, 0x30,
		0x65, 0x81, 0x29, 0x3b, 0x1d, 0xb7, 0xfc, 0x61, 0xe7, 0xda, 0xad, 0x53, 0xdf, 0x3f, 0x6d, 0x91,
		0x6d, 0x0e, 0x39, 0xee, 0x9d, 0x6c, 0x37, 0x7b, 0x5d, 0x87, 0xba, 0xbe, 0x17, 0x92, 0xae, 0xdd,
		0x1e, 0xbd, 0x4f, 0xdd, 0x36, 0x09, 0xa8, 0xd3, 0xee, 0x08, 0xc0, 0x66, 0xd2, 0x93, 0x1b, 0x7e,
		0xbb, 0xdd, 0x97, 0x48, 0xcc, 0x8d, 0x3a, 0xc1, 0xfb, 0x96, 0x1b, 0xd0, 0x10, 0x53, 0xfc, 0xfd,
		0x02, 0x6c, 0x1c, 0x8a, 0x74, 0xb5, 0x4f, 0xa4, 0xd1, 0x63, 0x29, 0xe8, 0xde, 0x89, 0x8f, 0xea,
		0x80, 0xa2, 0x75, 0xd8, 0x24, 0xba, 0x53, 0xc8, 0x6c, 0x66, 0xb6, 0xf2, 0x8f, 0xee, 0x95, 0x13,
		0x96, 0x54, 0x1e, 0xd3, 0xc1, 0x6b, 0x1f, 0x47, 0x43, 0xe8, 0x09, 0xcc, 0xd1, 0xf3, 0x0e, 0x29,
		0x64, 0xb9, 0xd0, 0x9d, 0xa9, 0x42, 0xd6, 0x79, 0x87, 0x60, 0x0e, 0x47, 0xcf, 0x00, 0x02, 0xea,
		0x74, 0xa9, 0xcd, 0xca, 0x50, 0xc8, 0x71, 0xf2, 0xb5, 0x72, 0x58, 0xa3, 0x72, 0x54, 0xa3, 0xb2,
		0x15, 0xd5, 0x08, 0x2f, 0x71, 0x34, 0xbb, 0x66, 0xd4, 0x46, 0xcb, 0x0f, 0x48, 0x48, 0x9d, 0x4b,
		0xa7, 0x72, 0x34, 0xa7, 0x5a, 0xb0, 0x1c, 0x52, 0x03, 0xea, 0xd0, 0x5e, 0x50, 0x98, 0xdf, 0xcc,
		0x6c, 0xad, 0x3c, 0xda, 0x99, 0x6d, 0xf5, 0x0a, 0x63, 0x9a, 0x9c, 0x88, 0xf3, 0x8d, 0xc1, 0x05,
		0xba, 0x0b, 0x2b, 0x67, 0x6e, 0x40, 0xfd, 0xee, 0xb9, 0xdd, 0x22, 0xde, 0x29, 0x3d, 0x2b, 0x2c,
		0x6c, 0x66, 0xb6, 0x72, 0xf8, 0xa2, 0x88, 0x56, 0x78, 0x10, 0xfd, 0x0c, 0x36, 0x3a, 0x4e, 0x97,
		0x78, 0x74, 0x50, 0x7e, 0xdb, 0xf5, 0x4e, 0xfc, 0xc2, 0x05, 0xbe, 0x84, 0xad, 0xc4, 0x2c, 0x6a,
		0x9c, 0x31, 0xb4, 0x93, 0x78, 0xbd, 0x33, 0x1e, 0x44, 0x32, 0xac, 0x0c, 0x64, 0x79, 0x65, 0x16,
		0x53, 0x2b, 0x73, 0xb1, 0xcf, 0xe0, 0xd5, 0x79, 0x08, 0x73, 0x6d, 0xd2, 0xf6, 0x0b, 0x4b, 0x9c,
		0x78, 0x35, 0x31, 0x9f, 0x03, 0xd2, 0xf6, 0x31, 0x87, 0x21, 0x0c, 0x6b, 0x01, 0x71, 0xba, 0x8d,
		0x33, 0xdb, 0xa1, 0xb4, 0xeb, 0x1e, 0xf7, 0x28, 0x09, 0x0a, 0xc0, 0xb9, 0x77, 0x13, 0xb9, 0x26,
		0x47, 0xcb, 0x7d, 0x30, 0x96, 0x82, 0x91, 0x08, 0xaa, 0xc0, 0x9a, 0xd3, 0xa3, 0xbe, 0xdd, 0x25,
		0x01, 0xa1, 0x76, 0xc7, 0x77, 0x3d, 0x1a, 0x14, 0xf2, 0x5c, 0x73, 0x33, 0x51, 0x13, 0x33, 0x60,
		0x8d, 0xe3, 0xf0, 0x2a, 0xa3, 0xc6, 0x02, 0xe8, 0x3a, 0x2c, 0xb1, 0xf1, 0xb0, 0xd9, 0x7c, 0x14,
		0x96, 0x37, 0x33, 0x5b, 0x4b, 0x78, 0x91, 0x05, 0x2a, 0x6e, 0x40, 0x8b, 0xbf, 0xce, 0xc2, 0xad,
		0xf1, 0x3d, 0xf6, 0xbd, 0x13, 0xf7, 0x54, 0x4c, 0x2e, 0xfa, 0x36, 0xce, 0x0f, 0x27, 0xe5, 0x66,
		0x62, 0x16, 0x96, 0x10, 0x1d, 0xc8, 0x23, 0x07, 0x36, 0x07, 0xfb, 0x21, 0x5a, 0xdd, 0xb7, 0x07,
		0x8d, 0xeb, 0xf7, 0xa8, 0x98, 0x99, 0xab, 0x63, 0x3b, 0xa4, 0x8a, 0x04, 0xf0, 0x8d, 0xbe, 0x84,
		0xc9, 0xdb, 0xdf, 0x57, 0xa2, 0x56, 0xf6, 0x7b, 0x14, 0x1d, 0xc2, 0x75, 0x9e, 0xde, 0x04, 0xf5,
		0x5c, 0x9a, 0xfa, 0x15, 0xc6, 0x4e, 0x10, 0x2e, 0xfe, 0x3d, 0x03, 0xeb, 0x09, 0x8d, 0xc7, 0xea,
		0xd9, 0xf4, 0xdb, 0x8e, 0xeb, 0xd9, 0x6e, 0x93, 0xd7, 0x63, 0x09, 0x2f, 0x86, 0x01, 0xbd, 0x89,
		0x6e, 0x43, 0x5e, 0xdc, 0xf4, 0x9c, 0x76, 0xe8, 0x07, 0x4b, 0x18, 0xc2, 0x90, 0xe1, 0xb4, 0xc9,
		0x04, 0x03, 0xca, 0xfd, 0xb7, 0x06, 0x74, 0x07, 0x96, 0x5d, 0xcf, 0xa5, 0xae, 0x43, 0x49, 0x93,
		0xe5, 0x35, 0xc7, 0x67, 0x2f, 0xdf, 0x8f, 0xe9, 0xcd, 0xe2, 0xaf, 0x32, 0xb0, 0xa1, 0x7d, 0xa2,
		0xa4, 0xeb, 0x39, 0xad, 0xef, 0xc5, 0x14, 0x47, 0x73, 0xca, 0x8e, 0xe7, 0xf4, 0xcf, 0x79, 0x58,
		0xaf, 0x11, 0xaf, 0xe9, 0x7a, 0xa7, 0x72, 0x83, 0xba, 0x1f, 0x5c, 0x7a, 0xce, 0x33, 0xba, 0x0d,
		0x79, 0x47, 0x5c, 0x0f, 0xaa, 0x0c, 0x51, 0x48, 0x6f, 0xa2, 0x3d, 0xb8, 0xd8, 0x07, 0xa4, 0x3a,
		0x6f, 0x24, 0xcd, 0x9d, 0x77, 0xd9, 0x89, 0x5d, 0xa1, 0x97, 0x30, 0xcf, 0x5c, 0x30, 0x34, 0xdf,
		0x95, 0x47, 0x0f, 0x92, 0xed, 0x67, 0x38, 0x43, 0x66, 0x78, 0x04, 0x87, 0x3c, 0xa4, 0xc3, 0xda,
		0x19, 0x71, 0xba, 0xf4, 0x98, 0x38, 0xd4, 0x6e, 0x12, 0xea, 0xb8, 0xad, 0x40, 0xd8, 0xf1, 0x8d,
		0x09, 0x5e, 0x76, 0xde, 0xf2, 0x9d, 0x26, 0x96, 0xfa, 0x34, 0x35, 0x64, 0xa1, 0xd7, 0xb0, 0xde,
		0x72, 0x02, 0x6a, 0x0f, 0xf4, 0xb8, 0x83, 0xcd, 0xa7, 0x3a, 0xd8, 0x1a, 0xa3, 0xed, 0x47, 0x2c,
		0xee, 0x62, 0x7b, 0xc0, 0x83, 0xe1, 0x54, 0x90, 0x66, 0xa8, 0xb4, 0x90, 0xaa, 0xb4, 0xca, 0x48,
		0x66, 0xc8, 0xe1, 0x3a, 0x05, 0xb8, 0xe0, 0x50, 0x4a, 0xda, 0x1d, 0xca, 0x0d, 0x7a, 0x1e, 0x47,
		0x97, 0xe8, 0x01, 0x48, 0x6d, 0xe7, 0x93, 0xdb, 0xee, 0xb5, 0x6d, 0x11, 0x0a, 0xb8, 0xd9, 0xce,
		0xe3, 0x55, 0x11, 0x97, 0x45, 0x98, 0xb9, 0x72, 0xd0, 0x38, 0x23, 0xcd, 0x5e, 0x2b, 0xca, 0x64,
		0x29, 0xdd, 0x95, 0xfb, 0x0c, 0x9e, 0x87, 0x02, 0xab, 0xe4, 0x53, 0xc7, 0x0d, 0x67, 0x36, 0xd4,
		0x80, 0x54, 0x8d, 0x95, 0x01, 0x85, 0x8b, 0xbc, 0x84, 0x65, 0x5e, 0x94, 0x13, 0xc7, 0x6d, 0xf5,
		0xba, 0x44, 0x58, 0x6a, 0xf2, 0x36, 0xed, 0x85, 0x18, 0x9c, 0x67, 0x0c, 0x71, 0x81, 0xbe, 0x82,
		0x4b, 0x5c, 0x80, 0xf5, 0x3a, 0xe9, 0xda, 0x6e, 0x93, 0x78, 0xd4, 0xa5, 0xe7, 0xc2, 0x55, 0x11,
		0xbb, 0x77, 0xc8, 0x6f, 0xe9, 0xe2, 0x4e, 0xf1, 0xb7, 0x59, 0xb8, 0x2a, 0xda, 0x47, 0x39, 0x73,
		0x5b, 0xcd, 0xef, 0x65, 0xf0, 0xbe, 0x8c, 0xc9, 0xb2, 0xe1, 0x88, 0x7b, 0x91, 0xf4, 0x31, 0x76,
		0x0c, 0xe1, 0x8e, 0x34, 0x3a, 0xa6, 0xb9, 0xb1, 0x31, 0x45, 0x6f, 0x41, 0xbc, 0x6d, 0x85, 0xb9,
		0x76, 0xfc, 0x96, 0xdb, 0x38, 0xe7, 0x6d, 0xbe, 0x32, 0x21, 0xd1, 0xd0, 0x39, 0xb9, 0xa1, 0xd6,
		0x38, 0x1a, 0xaf, 0x75, 0x46, 0x43, 0xc5, 0xbf, 0x65, 0xfb, 0xe3, 0xaf, 0x92, 0x86, 0x1b, 0x44,
		0x75, 0xe9, 0x4f, 0x65, 0x26, 0x7d, 0x2a, 0x23, 0xe2, 0xd0, 0x54, 0x8e, 0x77, 0x5c, 0xf6, 0x73,
		0x3b, 0xee, 0x05, 0x2c, 0x0f, 0x0d, 0x4f, 0xfa, 0xe9, 0x2c, 0x1f, 0x24, 0x0f, 0xce, 0xdc, 0xf0,
		0xe0, 0x60, 0xb8, 0xe2, 0x77, 0xdd, 0x53, 0xd7, 0x73, 0x5a, 0xf6, 0x48, 0x92, 0xe9, 0xa3, 0xbe,
		0x11, 0x51, 0xcd, 0x78, 0xb2, 0xc5, 0x3f, 0x67, 0xe1, 0x6a, 0x64, 0x4f, 0x15, 0xbf, 0xe1, 0xb4,
		0x54, 0x37, 0xe8, 0x38, 0xb4, 0x71, 0x36, 0x9b, 0x9b, 0xfe, 0xef, 0xcb, 0xf5, 0x73, 0xb8, 0x35,
		0x9c, 0x81, 0xed, 0x9f, 0xd8, 0xf4, 0xcc, 0x0d, 0xec, 0x78, 0x15, 0xa7, 0x0b, 0x5e, 0x1b, 0xca,
		0xa8, 0x7a, 0x62, 0x9d, 0xb9, 0x81, 0xf0, 0x20, 0x74, 0x13, 0x80, 0x9f, 0x12, 0xa8, 0xff, 0x9e,
		0x78, 0xbc, 0xce, 0xcb, 0x98, 0x1f, 0x6b, 0x2c, 0x16, 0x28, 0xbe, 0x86, 0x7c, 0xfc, 0xc8, 0xf4,
		0x1c, 0x16, 0xc4, 0xa9, 0x2b, 0xb3, 0x99, 0xdb, 0xca, 0x3f, 0xfa, 0x22, 0xe5, 0xd4, 0xc5, 0x0f,
		0xa4, 0x82, 0x52, 0xfc, 0x63, 0x16, 0x56, 0x86, 0x6f, 0xa1, 0xfb, 0xb0, 0x7a, 0xec, 0x7a, 0x4e,
		0xf7, 0xdc, 0x6e, 0x9c, 0x91, 0xc6, 0xfb, 0xa0, 0xd7, 0x16, 0x9b, 0xb0, 0x12, 0x86, 0x15, 0x11,
		0x45, 0x1b, 0xb0, 0xd0, 0xed, 0x79, 0xd1, 0xcb, 0x72, 0x09, 0xcf, 0x77, 0x7b, 0xec, 0x54, 0xf1,
		0x02, 0xae, 0x9f, 0xb8, 0xdd, 0x80, 0xbd, 0x60, 0xc2, 0x66, 0xb7, 0x1b, 0x7e, 0xbb, 0xd3, 0x22,
		0x43, 0x13, 0x5b, 0xe0, 0x90, 0x68, 0x1c, 0x94, 0x08, 0xc0, 0xe9, 0xcb, 0x8d, 0x2e, 0x71, 0xfa,
		0x7b, 0x93, 0x5e, 0xca, 0xbc, 0xc0, 0x0b, 0xdb, 0xbc, 0xc8, 0x8d, 0xd4, 0xf5, 0x4e, 0x67, 0x6d,
		0xd3, 0xe5, 0x88, 0xc0, 0x05, 0x6e, 0x01, 0xf0, 0xa3, 0x2c, 0x75, 0x8e, 0x5b, 0xe1, 0x5b, 0x68,
		0x11, 0xc7, 0x22, 0xa5, 0x3f, 0x65, 0xe0, 0x52, 0xd2, 0x3b, 0x16, 0x15, 0xe1, 0x56, 0x4d, 0x33,
		0x54, 0xdd, 0x78, 0x65, 0xcb, 0x8a, 0xa5, 0xbf, 0xd5, 0xad, 0x23, 0xdb, 0xb4, 0x64, 0x4b, 0xb3,
		0x75, 0xe3, 0xad, 0x5c, 0xd1, 0x55, 0xe9, 0xff, 0xd0, 0xff, 0xc3, 0xe6, 0x04, 0x8c, 0xa9, 0xec,
		0x6b, 0x6a, 0xbd, 0xa2, 0xa9, 0x52, 0x66, 0x8a, 0x92, 0x69, 0xc9, 0xd8, 0xd2, 0x54, 0x29, 0x8b,
		0x7e, 0x00, 0xf7, 0x27, 0x60, 0x14, 0xd9, 0x50, 0xb4, 0x8a, 0x8d, 0xb5, 0x9f, 0xd6, 0x35, 0x93,
		0x81, 0x73, 0xa5, 0x5f, 0x0c, 0x72, 0x1e, 0x72, 0xa0, 0xf8, 0x93, 0x54, 0x4d, 0xd1, 0x4d, 0xbd,
		0x6a, 0x4c, 0xcb, 0x79, 0x04, 0x33, 0x21, 0xe7, 0x51, 0x54, 0x94, 0x73, 0xe9, 0x97, 0xd9, 0xc1,
		0x97, 0xae, 0xde, 0xc4, 0xa4, 0x17, 0x79, 0x2b, 0x7b, 0xc6, 0x61, 0x15, 0xbf, 0xd9, 0xab, 0x54,
		0x0f, 0x6d, 0x5d, 0xb5, 0xb1, 0x56, 0x37, 0x35, 0xbb, 0x56, 0xad, 0xe8, 0xca, 0x51, 0x2c, 0x93,
		0x1f, 0xc2, 0xd7, 0x13, 0x51, 0x72, 0x85, 0x45, 0xd5, 0x7a, 0xad, 0xa2, 0x2b, 0xec, 0xa9, 0x7b,
		0xb2, 0x5e, 0xd1, 0x54, 0xbb, 0x6a, 0x54, 0x8e, 0xa4, 0x0c, 0xfa, 0x12, 0xb6, 0x66, 0x65, 0x4a,
		0x59, 0xf4, 0x10, 0x1e, 0x4c, 0x44, 0x63, 0xed, 0xb5, 0xa6, 0x58, 0x31, 0x78, 0x0e, 0xed, 0xc0,
		0xc3, 0x89, 0x70, 0x4b, 0xc3, 0x07, 0xba, 0xc1, 0x0b, 0xba, 0x67, 0xe3, 0xba, 0x61, 0xe8, 0xc6,
		0x2b, 0x69, 0xae, 0xf4, 0xbb, 0x0c, 0xac, 0x8d, 0xbd, 0x74, 0xd0, 0x6d, 0xb8, 0x5e, 0x93, 0xb1,
		0x66, 0x58, 0xb6, 0x52, 0xa9, 0x26, 0x15, 0x60, 0x02, 0x40, 0xde, 0x95, 0x0d, 0xb5, 0x6a, 0x48,
		0x19, 0x74, 0x0f, 0x8a, 0x49, 0x00, 0xd1, 0x0b, 0xa2, 0x35, 0xa4, 0x2c, 0xba, 0x03, 0x37, 0x93,
		0x70, 0xfd, 0x6c, 0xa5, 0x5c, 0xe9, 0x5f, 0x59, 0xb8, 0x31, 0xed, 0x83, 0x9a, 0x75, 0x60, 0x7f,
		0xd9, 0xda, 0x3b, 0x4d, 0xa9, 0x5b, 0x6c, 0xcf, 0x43, 0x3d, 0xb6, 0xf3, 0x75, 0x33, 0x96, 0x79,
		0xbc, 0xa4, 0x13, 0xc0, 0x4a, 0xf5, 0xa0, 0x56, 0xd1, 0x2c, 0xde, 0x4d, 0x25, 0xb8, 0x97, 0x06,
		0x0f, 0x37, 0x58, 0xca, 0x0e, 0xed, 0xed, 0x24, 0x69, 0xbe, 0x6e, 0x36, 0x0a, 0xa8, 0x0c, 0xa5,
		0x34, 0x74, 0xbf, 0x0a, 0xaa, 0x34, 0x87, 0xbe, 0x86, 0xaf, 0xd2, 0x13, 0x37, 0x2c, 0xdd, 0xa8,
		0x6b, 0xaa, 0x2d, 0x9b, 0xb6, 0xa1, 0x1d, 0x4a, 0xf3, 0xb3, 0x2c, 0xd7, 0xd2, 0x0f, 0x58, 0x7f,
		0xd6, 0x2d, 0x69, 0xa1, 0xf4, 0x97, 0x0c, 0x5c, 0x56, 0x7c, 0x8f, 0xba, 0x5e, 0x8f, 0xc8, 0x81,
		0x41, 0x3e, 0xea, 0xe1, 0x79, 0xc6, 0xef, 0xa2, 0xbb, 0x70, 0x27, 0xd2, 0x17, 0xf2, 0xb6, 0x6e,
		0xe8, 0x96, 0x2e, 0x5b, 0x55, 0x1c, 0xab, 0xef, 0x54, 0x18, 0x1b, 0x48, 0x55, 0xc3, 0x61, 0x5d,
		0x27, 0xc3, 0xb0, 0x66, 0xe1, 0x23, 0xd1, 0x0a, 0xa1, 0xc3, 0x4c, 0xc6, 0x2a, 0x98, 0xcd, 0xb7,
		0x98, 0x7f, 0x29, 0x57, 0xfa, 0x43, 0x06, 0xf2, 0xe2, 0x5b, 0x94, 0x7f, 0xaa, 0x14, 0xe0, 0x12,
		0x5b, 0x60, 0xb5, 0x6e, 0xd9, 0xd6, 0x51, 0x4d, 0x1b, 0xee, 0xe1, 0xa1, 0x3b, 0xdc, 0x1e, 0x6c,
		0xab, 0x1a, 0x56, 0x27, 0x74, 0x92, 0x61, 0x80, 0x78, 0x0a, 0xc3, 0x70, 0xb0, 0x94, 0x9d, 0x8a,
		0x09, 0x75, 0x72, 0xe8, 0x1a, 0x5c, 0x1e, 0xc2, 0xec, 0x6b, 0x32, 0xb6, 0x76, 0x35, 0xd9, 0x92,
		0xe6, 0x4a, 0xbf, 0xc9, 0xc0, 0xd5, 0xc8, 0x09, 0x2d, 0xf6, 0x62, 0x75, 0xdb, 0xa4, 0x59, 0xed,
		0x51, 0xc5, 0xe9, 0x05, 0x04, 0x3d, 0x80, 0xbb, 0x7d, 0x0f, 0xb3, 0x64, 0xf3, 0xcd, 0x60, 0xaf,
		0x6c, 0x45, 0x66, 0xc3, 0x3d, 0x58, 0x4d, 0x2a, 0x54, 0xa4, 0x20, 0x65, 0xd0, 0x7d, 0xf8, 0x62,
		0x3a, 0x14, 0x6b, 0xa6, 0x66, 0x49, 0xd9, 0xd2, 0x3f, 0xf2, 0x70, 0x25, 0x9e, 0x1c, 0x3b, 0xd0,
		0x93, 0x66, 0x98, 0xda, 0x3d, 0x28, 0x0e, 0x8b, 0x08, 0x9f, 0x1b, 0xcd, 0x6b, 0x07, 0x1e, 0x4e,
		0xc1, 0xd5, 0x8d, 0x7d, 0xd9, 0x50, 0xd9, 0x75, 0x04, 0x92, 0x32, 0xe8, 0x25, 0x3c, 0x9f, 0x42,
		0xd9, 0x95, 0xd5, 0x41, 0x95, 0xfb, 0x6f, 0x1c, 0xd9, 0xb2, 0xb0, 0xbe, 0x5b, 0xb7, 0x34, 0x53,
		0xca, 0x22, 0x0d, 0xe4, 0x14, 0x81, 0x61, 0x1f, 0x4a, 0x94, 0xc9, 0xa1, 0x67, 0xf0, 0x24, 0x2d,
		0x8f, 0xb0, 0x65, 0xf4, 0x03, 0x0d, 0xc7, 0xa9, 0x73, 0xe8, 0x5b, 0xf8, 0x26, 0x85, 0x2a, 0x9e,
		0x3c, 0xc6, 0x9d, 0x47, 0xcf, 0xe1, 0x69, 0x6a, 0xf6, 0x4a, 0x15, 0xab, 0xf6, 0x81, 0x8c, 0xdf,
		0x0c, 0x93, 0x17, 0x90, 0x0e, 0x5a, 0xda, 0x83, 0x85, 0xbb, 0xd9, 0x09, 0xbe, 0x10, 0x93, 0xba,
		0x30, 0x43, 0x15, 0x59, 0x20, 0x45, 0x66, 0x11, 0xbd, 0x02, 0x65, 0xb6, 0x52, 0x4c, 0x17, 0x5a,
		0x42, 0xef, 0xc0, 0xfa, 0xbc, 0x5d, 0xd5, 0xde, 0x59, 0x1a, 0x36, 0xe4, 0x34, 0x65, 0x40, 0x2f,
		0xe0, 0x59, 0x6a, 0xd1, 0x86, 0xfd, 0x27, 0x46, 0xcf, 0xa3, 0xa7, 0xf0, 0x78, 0x0a, 0x3d, 0xde,
		0x23, 0x83, 0x53, 0x81, 0xae, 0x4a, 0xcb, 0xe8, 0x09, 0xec, 0x4c, 0x21, 0xf2, 0x29, 0xb4, 0x4d,
		0x4b, 0x57, 0xde, 0x1c, 0x85, 0xb7, 0x2b, 0xba, 0x69, 0x49, 0x17, 0xd1, 0x4f, 0xe0, 0x47, 0x53,
		0x68, 0xfd, 0xc5, 0xb2, 0x3f, 0x34, 0x1c, 0x1b, 0x31, 0x06, 0xab, 0x63, 0x4d, 0x5a, 0x99, 0x61,
		0x4f, 0x4c, 0xfd, 0x55, 0x7a, 0xe5, 0x56, 0x91, 0x02, 0x2f, 0x67, 0x1a, 0x11, 0x65, 0x5f, 0xaf,
		0xa8, 0xc9, 0x22, 0x12, 0x7a, 0x0c, 0xdb, 0x53, 0x44, 0xf6, 0xaa, 0x58, 0xd1, 0xc4, 0x1b, 0xab,
		0x6f, 0x12, 0x6b, 0xe8, 0x1b, 0x78, 0x34, 0x8d, 0x24, 0xeb, 0x95, 0xea, 0x5b, 0x0d, 0x8f, 0xf2,
		0x10, 0x7b, 0x8d, 0xce, 0xb6, 0x74, 0xdd, 0xa8, 0xd5, 0x2d, 0xdb, 0xd4, 0xbf, 0xd3, 0xa4, 0x75,
		0xf6, 0x1a, 0x4d, 0xdd, 0xa9, 0xa8, 0x56, 0xd2, 0xa5, 0x71, 0x33, 0x1e, 0x7b, 0xc8, 0xae, 0x6e,
		0xc8, 0xf8, 0x48, 0xda, 0x48, 0xe9, 0xbd, 0x71, 0xa3, 0x1b, 0x6a, 0xa1, 0xcb, 0xb3, 0x2c, 0x47,
		0x93, 0xb1, 0xb2, 0x1f, 0xaf, 0xf8, 0x15, 0xf6, 0xd6, 0xb9, 0xc3, 0xff, 0xb1, 0x32, 0x76, 0xae,
		0x8a, 0x5b, 0xfc, 0x0e, 0x3c, 0x0c, 0xf7, 0x2d, 0xa1, 0x0b, 0x26, 0xb8, 0xfd, 0x2e, 0xfc, 0x78,
		0x36, 0x4a, 0xff, 0xbe, 0x5c, 0xc1, 0x9a, 0xac, 0x1e, 0xf5, 0x8f, 0xa4, 0x99, 0xd2, 0x5f, 0x33,
		0x50, 0x52, 0x1c, 0xaf, 0x41, 0x5a, 0xd1, 0xff, 0x5d, 0xa7, 0x66, 0xf9, 0x1c, 0x9e, 0xce, 0x30,
		0xef, 0x13, 0xf2, 0x3d, 0x04, 0xf3, 0x73, 0xc9, 0x75, 0xe3, 0x8d, 0x51, 0x3d, 0x34, 0xa6, 0x11,
		0xc4, 0x22, 0x4c, 0xf7, 0x94, 0xff, 0xd3, 0x78, 0xb6, 0x45, 0x88, 0xb6, 0xfb, 0xcf, 0x16, 0xf1,
		0xb9, 0xe4, 0x99, 0x16, 0xb1, 0xfb, 0x16, 0xae, 0x34, 0xfc, 0x76, 0xd2, 0x57, 0xfc, 0xee, 0xa2,
		0xdc, 0x71, 0x6b, 0xec, 0x0b, 0xb6, 0x96, 0xf9, 0x6e, 0xfb, 0xd4, 0xa5, 0x67, 0xbd, 0xe3, 0x72,
		0xc3, 0x6f, 0x6f, 0x0f, 0xfd, 0xcc, 0x58, 0x3e, 0x25, 0x5e, 0xf8, 0xa3, 0xa5, 0xf8, 0xc5, 0xf1,
		0xb9, 0xd3, 0x71, 0x3f, 0xec, 0x1c, 0x2f, 0xf0, 0xd8, 0xe3, 0x7f, 0x07, 0x00, 0x00, 0xff, 0xff,
		0x05, 0xd4, 0x20, 0x14, 0x31, 0x1d, 0x00, 0x00,
	},
	// google/protobuf/duration.proto
	[]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4b, 0xcf, 0xcf, 0x4f,
		0xcf, 0x49, 0xd5, 0x2f, 0x28, 0xca, 0x2f, 0xc9, 0x4f, 0x2a, 0x4d, 0xd3, 0x4f, 0x29, 0x2d, 0x4a,
		0x2c, 0xc9, 0xcc, 0xcf, 0xd3, 0x03, 0x8b, 0x08, 0xf1, 0x43, 0xe4, 0xf5, 0x60, 0xf2, 0x4a, 0x56,
		0x5c, 0x1c, 0x2e, 0x50, 0x25, 0x42, 0x12, 0x5c, 0xec, 0xc5, 0xa9, 0xc9, 0xf9, 0x79, 0x29, 0xc5,
		0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0xcc, 0x41, 0x30, 0xae, 0x90, 0x08, 0x17, 0x6b, 0x5e, 0x62, 0x5e,
		0x7e, 0xb1, 0x04, 0x93, 0x02, 0xa3, 0x06, 0x6b, 0x10, 0x84, 0xe3, 0xd4, 0xcc, 0xc8, 0x25, 0x9c,
		0x9c, 0x9f, 0xab, 0x87, 0x66, 0xa6, 0x13, 0x2f, 0xcc, 0xc4, 0x00, 0x90, 0x48, 0x00, 0x63, 0x94,
		0x21, 0x54, 0x45, 0x7a, 0x7e, 0x4e, 0x62, 0x5e, 0xba, 0x5e, 0x7e, 0x51, 0x3a, 0xc2, 0x81, 0x25,
		0x95, 0x05, 0xa9, 0xc5, 0xfa, 0xd9, 0x79, 0xf9, 0xe5, 0x79, 0x70, 0xc7, 0x16, 0x24, 0xfd, 0x60,
		0x64, 0x5c, 0xc4, 0xc4, 0xec, 0x1e, 0xe0, 0xb4, 0x8a, 0x49, 0xce, 0x1d, 0xa2, 0x39, 0x00, 0xaa,
		0x43, 0x2f, 0x3c, 0x35, 0x27, 0xc7, 0x1b, 0xa4, 0x3e, 0x04, 0xa4, 0x35, 0x89, 0x0d, 0x6c, 0x94,
		0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0xef, 0x8a, 0xb4, 0xc3, 0xfb, 0x00, 0x00, 0x00,
	},
	// uber/cadence/api/v1/common.proto
	[]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x55, 0x4f, 0x6f, 0xdb, 0x36,
		0x14, 0x9f, 0xe2, 0xda, 0x49, 0x9f, 0xdd, 0xd4, 0x63, 0xd6, 0xd4, 0xc9, 0xfe, 0x79, 0x06, 0x86,
		0x66, 0x3b, 0x48, 0x88, 0x7b, 0x29, 0x56, 0x14, 0x83, 0x13, 0x3b, 0xab, 0xda, 0x2d, 0x31, 0x64,
		0x23, 0xc1, 0x76, 0x98, 0x40, 0x4b, 0x4f, 0x2e, 0x67, 0x89, 0x14, 0x28, 0xca, 0x89, 0x6f, 0xfb,
		0x24, 0x3b, 0xec, 0x2b, 0xed, 0x0b, 0x0d, 0x94, 0xe8, 0xd8, 0xee, 0x3c, 0xf4, 0x32, 0xec, 0x46,
		0xbe, 0xdf, 0x9f, 0xf7, 0xa3, 0xf0, 0x48, 0x41, 0x3b, 0x9f, 0xa0, 0x74, 0x02, 0x1a, 0x22, 0x0f,
		0xd0, 0xa1, 0x29, 0x73, 0xe6, 0xa7, 0x4e, 0x20, 0x92, 0x44, 0x70, 0x3b, 0x95, 0x42, 0x09, 0x72,
		0xa0, 0x19, 0xb6, 0x61, 0xd8, 0x34, 0x65, 0xf6, 0xfc, 0xf4, 0xf8, 0x8b, 0xa9, 0x10, 0xd3, 0x18,
		0x9d, 0x82, 0x32, 0xc9, 0x23, 0x27, 0xcc, 0x25, 0x55, 0x6c, 0x29, 0xea, 0xbc, 0x85, 0x8f, 0x6f,
		0x84, 0x9c, 0x45, 0xb1, 0xb8, 0x1d, 0xdc, 0x61, 0x90, 0x6b, 0x88, 0x7c, 0x09, 0xf5, 0x5b, 0x53,
		0xf4, 0x59, 0xd8, 0xb2, 0xda, 0xd6, 0xc9, 0x43, 0x0f, 0x96, 0x25, 0x37, 0x24, 0x4f, 0xa0, 0x26,
		0x73, 0xae, 0xb1, 0x9d, 0x02, 0xab, 0xca, 0x9c, 0xbb, 0x61, 0xa7, 0x03, 0x8d, 0xa5, 0xd9, 0x78,
		0x91, 0x22, 0x21, 0xf0, 0x80, 0xd3, 0x04, 0x8d, 0x41, 0xb1, 0xd6, 0x9c, 0x5e, 0xa0, 0xd8, 0x9c,
		0xa9, 0xc5, 0xbf, 0x72, 0x3e, 0x87, 0xdd, 0x21, 0x5d, 0xc4, 0x82, 0x86, 0x1a, 0x0e, 0xa9, 0xa2,
		0x05, 0xdc, 0xf0, 0x8a, 0x75, 0xe7, 0x25, 0xec, 0x5e, 0x50, 0x16, 0xe7, 0x12, 0xc9, 0x21, 0xd4,
		0x24, 0xd2, 0x4c, 0x70, 0xa3, 0x37, 0x3b, 0xd2, 0x82, 0xdd, 0x10, 0x15, 0x65, 0x71, 0x56, 0x24,
		0x6c, 0x78, 0xcb, 0x6d, 0xe7, 0x0f, 0x0b, 0x1e, 0xfc, 0x84, 0x89, 0x20, 0xaf, 0xa0, 0x16, 0x31,
		0x8c, 0xc3, 0xac, 0x65, 0xb5, 0x2b, 0x27, 0xf5, 0xee, 0xd7, 0xf6, 0x96, 0xef, 0x67, 0x6b, 0xaa,
		0x7d, 0x51, 0xf0, 0x06, 0x5c, 0xc9, 0x85, 0x67, 0x44, 0xc7, 0x37, 0x50, 0x5f, 0x2b, 0x93, 0x26,
		0x54, 0x66, 0xb8, 0x30, 0x29, 0xf4, 0x92, 0x74, 0xa1, 0x3a, 0xa7, 0x71, 0x8e, 0x45, 0x80, 0x7a,
		0xf7, 0xb3, 0xad, 0xf6, 0xe6, 0x98, 0x5e, 0x49, 0xfd, 0x6e, 0xe7, 0x85, 0xd5, 0xf9, 0xd3, 0x82,
		0xda, 0x6b, 0xa4, 0x21, 0x4a, 0xf2, 0xfd, 0x7b, 0x11, 0x9f, 0x6d, 0xf5, 0x28, 0xc9, 0xff, 0x6f,
		0xc8, 0xbf, 0x2c, 0x68, 0x8e, 0x90, 0xca, 0xe0, 0x5d, 0x4f, 0x29, 0xc9, 0x26, 0xb9, 0xc2, 0x8c,
		0xf8, 0xb0, 0xcf, 0x78, 0x88, 0x77, 0x18, 0xfa, 0x1b, 0xb1, 0x5f, 0x6c, 0x75, 0x7d, 0x5f, 0x6e,
		0xbb, 0xa5, 0x76, 0xfd, 0x1c, 0x8f, 0xd8, 0x7a, 0xed, 0xf8, 0x57, 0x20, 0xff, 0x24, 0xfd, 0x87,
		0xa7, 0x8a, 0x60, 0xaf, 0x4f, 0x15, 0x3d, 0x8b, 0xc5, 0x84, 0x5c, 0xc0, 0x23, 0xe4, 0x81, 0x08,
		0x19, 0x9f, 0xfa, 0x6a, 0x91, 0x96, 0x03, 0xba, 0xdf, 0xfd, 0x6a, 0xab, 0xd7, 0xc0, 0x30, 0xf5,
		0x44, 0x7b, 0x0d, 0x5c, 0xdb, 0xdd, 0x0f, 0xf0, 0xce, 0xda, 0x00, 0x0f, 0xcb, 0x4b, 0x87, 0xf2,
		0x1a, 0x65, 0xc6, 0x04, 0x77, 0x79, 0x24, 0x34, 0x91, 0x25, 0x69, 0xbc, 0xbc, 0x08, 0x7a, 0x4d,
		0x9e, 0xc1, 0xe3, 0x08, 0xa9, 0xca, 0x25, 0xfa, 0xf3, 0x92, 0x6a, 0x2e, 0xdc, 0xbe, 0x29, 0x1b,
		0x83, 0xce, 0x5b, 0x78, 0x3a, 0xca, 0xd3, 0x54, 0x48, 0x85, 0xe1, 0x79, 0xcc, 0x90, 0x2b, 0x83,
		0x64, 0xfa, 0xae, 0x4e, 0x85, 0x9f, 0x85, 0x33, 0xe3, 0x5c, 0x9d, 0x8a, 0x51, 0x38, 0x23, 0x47,
		0xb0, 0xf7, 0x1b, 0x9d, 0xd3, 0x02, 0x28, 0x3d, 0x77, 0xf5, 0x7e, 0x14, 0xce, 0x3a, 0xbf, 0x57,
		0xa0, 0xee, 0xa1, 0x92, 0x8b, 0xa1, 0x88, 0x59, 0xb0, 0x20, 0x7d, 0x68, 0x32, 0xce, 0x14, 0xa3,
		0xb1, 0xcf, 0xb8, 0x42, 0x39, 0xa7, 0x65, 0xca, 0x7a, 0xf7, 0xc8, 0x2e, 0x9f, 0x17, 0x7b, 0xf9,
		0xbc, 0xd8, 0x7d, 0xf3, 0xbc, 0x78, 0x8f, 0x8d, 0xc4, 0x35, 0x0a, 0xe2, 0xc0, 0xc1, 0x84, 0x06,
		0x33, 0x11, 0x45, 0x7e, 0x20, 0x30, 0x8a, 0x58, 0xa0, 0x63, 0x16, 0xbd, 0x2d, 0x8f, 0x18, 0xe8,
		0x7c, 0x85, 0xe8, 0xb6, 0x09, 0xbd, 0x63, 0x49, 0x9e, 0xac, 0xda, 0x56, 0x3e, 0xd8, 0xd6, 0x48,
		0xee, 0xdb, 0x7e, 0xb3, 0x72, 0xa1, 0x4a, 0x61, 0x92, 0xaa, 0xac, 0xf5, 0xa0, 0x6d, 0x9d, 0x54,
		0xef, 0xa9, 0x3d, 0x53, 0x26, 0xaf, 0xe0, 0x53, 0x2e, 0xb8, 0x2f, 0xf5, 0xd1, 0xe9, 0x24, 0x46,
		0x1f, 0xa5, 0x14, 0xd2, 0x2f, 0x9f, 0x94, 0xac, 0x55, 0x6d, 0x57, 0x4e, 0x1e, 0x7a, 0x2d, 0x2e,
		0xb8, 0xb7, 0x64, 0x0c, 0x34, 0xc1, 0x2b, 0x71, 0xf2, 0x06, 0x0e, 0xf0, 0x2e, 0x65, 0x65, 0x90,
		0x55, 0xe4, 0xda, 0x87, 0x22, 0x93, 0x95, 0x6a, 0x99, 0xfa, 0xdb, 0x5b, 0x68, 0xac, 0xcf, 0x14,
		0x39, 0x82, 0x27, 0x83, 0xcb, 0xf3, 0xab, 0xbe, 0x7b, 0xf9, 0x83, 0x3f, 0xfe, 0x79, 0x38, 0xf0,
		0xdd, 0xcb, 0xeb, 0xde, 0x8f, 0x6e, 0xbf, 0xf9, 0x11, 0x39, 0x86, 0xc3, 0x4d, 0x68, 0xfc, 0xda,
		0x73, 0x2f, 0xc6, 0xde, 0x4d, 0xd3, 0x22, 0x87, 0x40, 0x36, 0xb1, 0x37, 0xa3, 0xab, 0xcb, 0xe6,
		0x0e, 0x69, 0xc1, 0x27, 0x9b, 0xf5, 0xa1, 0x77, 0x35, 0xbe, 0x7a, 0xde, 0xac, 0x9c, 0x5d, 0xc3,
		0xd3, 0x40, 0x24, 0xdb, 0x86, 0xfc, 0x6c, 0xaf, 0x97, 0xb2, 0xa1, 0x4e, 0x3f, 0xb4, 0x7e, 0x71,
		0xa6, 0x4c, 0xbd, 0xcb, 0x27, 0x76, 0x20, 0x12, 0x67, 0xe3, 0xc7, 0x64, 0x4f, 0x91, 0x97, 0x3f,
		0x1b, 0xf3, 0x8f, 0x7a, 0x49, 0x53, 0x36, 0x3f, 0x9d, 0xd4, 0x8a, 0xda, 0xf3, 0xbf, 0x03, 0x00,
		0x00, 0xff, 0xff, 0xdc, 0x8c, 0x77, 0x9a, 0xc7, 0x06, 0x00, 0x00,
	},
	// uber/cadence/api/v1/tasklist.proto
	[]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x54, 0xdd, 0x6e, 0xdb, 0x36,
		0x14, 0x9e, 0xe2, 0xb4, 0x73, 0x98, 0x25, 0xd5, 0xb8, 0xb5, 0x8d, 0xdd, 0x75, 0xcb, 0x74, 0x51,
		0x04, 0xc5, 0x20, 0x21, 0x19, 0x76, 0xb5, 0x8b, 0xc1, 0x89, 0x83, 0x55, 0x88, 0xe3, 0x1a, 0x92,
		0x1a, 0x20, 0xbb, 0xe1, 0x28, 0xf1, 0xd4, 0x21, 0xf4, 0x43, 0x81, 0xa4, 0x92, 0xf8, 0x45, 0xf6,
		0x30, 0x7b, 0xa2, 0x3d, 0xc6, 0x40, 0x4a, 0xf6, 0xdc, 0xc4, 0xeb, 0x1d, 0x79, 0xbe, 0xf3, 0x9d,
		0x9f, 0x8f, 0xe7, 0x10, 0x79, 0x4d, 0x0a, 0x32, 0xc8, 0x28, 0x83, 0x2a, 0x83, 0x80, 0xd6, 0x3c,
		0xb8, 0x3d, 0x0e, 0x34, 0x55, 0x79, 0xc1, 0x95, 0xf6, 0x6b, 0x29, 0xb4, 0xc0, 0xdf, 0x18, 0x1f,
		0xbf, 0xf3, 0xf1, 0x69, 0xcd, 0xfd, 0xdb, 0xe3, 0xe1, 0xf7, 0x73, 0x21, 0xe6, 0x05, 0x04, 0xd6,
		0x25, 0x6d, 0x3e, 0x06, 0xac, 0x91, 0x54, 0x73, 0x51, 0xb5, 0xa4, 0xe1, 0x0f, 0x0f, 0x71, 0xcd,
		0x4b, 0x50, 0x9a, 0x96, 0x75, 0xe7, 0xf0, 0x28, 0xc0, 0x9d, 0xa4, 0x75, 0x0d, 0x52, 0xb5, 0xb8,
		0xf7, 0x01, 0xf5, 0x13, 0xaa, 0xf2, 0x09, 0x57, 0x1a, 0x63, 0xb4, 0x5d, 0xd1, 0x12, 0x0e, 0x9c,
		0x43, 0xe7, 0x68, 0x27, 0xb2, 0x67, 0xfc, 0x0b, 0xda, 0xce, 0x79, 0xc5, 0x0e, 0xb6, 0x0e, 0x9d,
		0xa3, 0xfd, 0x93, 0x1f, 0xfd, 0x0d, 0x45, 0xfa, 0xcb, 0x00, 0x17, 0xbc, 0x62, 0x91, 0x75, 0xf7,
		0x28, 0x72, 0x97, 0xd6, 0x4b, 0xd0, 0x94, 0x51, 0x4d, 0xf1, 0x25, 0xfa, 0xb6, 0xa4, 0xf7, 0xc4,
		0xb4, 0xad, 0x48, 0x0d, 0x92, 0x28, 0xc8, 0x44, 0xc5, 0x6c, 0xba, 0xdd, 0x93, 0xef, 0xfc, 0xb6,
		0x52, 0x7f, 0x59, 0xa9, 0x3f, 0x16, 0x4d, 0x5a, 0xc0, 0x15, 0x2d, 0x1a, 0x88, 0xbe, 0x2e, 0xe9,
		0xbd, 0x09, 0xa8, 0x66, 0x20, 0x63, 0x4b, 0xf3, 0x3e, 0xa0, 0xc1, 0x32, 0xc5, 0x8c, 0x4a, 0xcd,
		0x8d, 0x2a, 0xab, 0x5c, 0x2e, 0xea, 0xe5, 0xb0, 0xe8, 0x3a, 0x31, 0x47, 0xfc, 0x06, 0x3d, 0x13,
		0x77, 0x15, 0x48, 0x72, 0x23, 0x94, 0x26, 0xb6, 0xcf, 0x2d, 0x8b, 0xee, 0x59, 0xf3, 0x3b, 0xa1,
		0xf4, 0x94, 0x96, 0xe0, 0xfd, 0xe3, 0xa0, 0xfd, 0x65, 0xdc, 0x58, 0x53, 0xdd, 0x28, 0xfc, 0x13,
		0xc2, 0x29, 0xcd, 0xf2, 0x42, 0xcc, 0x49, 0x26, 0x9a, 0x4a, 0x93, 0x1b, 0x5e, 0x69, 0x1b, 0xbb,
		0x17, 0xb9, 0x1d, 0x72, 0x66, 0x80, 0x77, 0xbc, 0xd2, 0xf8, 0x35, 0x42, 0x12, 0x28, 0x23, 0x05,
		0xdc, 0x42, 0x61, 0x73, 0xf4, 0xa2, 0x1d, 0x63, 0x99, 0x18, 0x03, 0x7e, 0x85, 0x76, 0x68, 0x96,
		0x77, 0x68, 0xcf, 0xa2, 0x7d, 0x9a, 0xe5, 0x2d, 0xf8, 0x06, 0x3d, 0x93, 0x54, 0xc3, 0xba, 0x3a,
		0xdb, 0x87, 0xce, 0x91, 0x13, 0xed, 0x19, 0xf3, 0xaa, 0x77, 0x3c, 0x46, 0x7b, 0x46, 0x46, 0xc2,
		0x19, 0x49, 0x0b, 0x91, 0xe5, 0x07, 0x4f, 0xac, 0x86, 0x87, 0xff, 0xfb, 0x3c, 0xe1, 0xf8, 0xd4,
		0xf8, 0x45, 0xbb, 0x86, 0x16, 0x32, 0x7b, 0xf1, 0x7e, 0x43, 0xbb, 0x6b, 0x18, 0x1e, 0xa0, 0xbe,
		0xd2, 0x54, 0x6a, 0xc2, 0x59, 0xd7, 0xdc, 0x97, 0xf6, 0x1e, 0x32, 0xfc, 0x1c, 0x3d, 0x85, 0x8a,
		0x19, 0xa0, 0xed, 0xe7, 0x09, 0x54, 0x2c, 0x64, 0xde, 0x5f, 0x0e, 0x42, 0x33, 0x51, 0x14, 0x20,
		0xc3, 0xea, 0xa3, 0xc0, 0x63, 0xe4, 0x16, 0x54, 0x69, 0x42, 0xb3, 0x0c, 0x94, 0x22, 0x66, 0x14,
		0xbb, 0xc7, 0x1d, 0x3e, 0x7a, 0xdc, 0x64, 0x39, 0xa7, 0xd1, 0xbe, 0xe1, 0x8c, 0x2c, 0xc5, 0x18,
		0xf1, 0x10, 0xf5, 0x39, 0x83, 0x4a, 0x73, 0xbd, 0xe8, 0x5e, 0x68, 0x75, 0xdf, 0xa4, 0x4f, 0x6f,
		0x83, 0x3e, 0xde, 0xdf, 0x0e, 0x1a, 0xc4, 0x9a, 0x67, 0xf9, 0xe2, 0xfc, 0x1e, 0xb2, 0xc6, 0x8c,
		0xc6, 0x48, 0x6b, 0xc9, 0xd3, 0x46, 0x83, 0xc2, 0xbf, 0x23, 0xf7, 0x4e, 0xc8, 0x1c, 0xa4, 0x9d,
		0x45, 0x62, 0x76, 0xb0, 0xab, 0xf3, 0xf5, 0x67, 0xe7, 0x3b, 0xda, 0x6f, 0x69, 0xab, 0x85, 0x49,
		0xd0, 0x40, 0x65, 0x37, 0xc0, 0x9a, 0x02, 0x88, 0x16, 0xa4, 0x55, 0xcf, 0xb4, 0x2d, 0x1a, 0x6d,
		0x6b, 0xdf, 0x3d, 0x19, 0x3c, 0x1e, 0xeb, 0x6e, 0x83, 0xa3, 0x17, 0x4b, 0x6e, 0x22, 0x62, 0xc3,
		0x4c, 0x5a, 0xe2, 0xdb, 0x3f, 0xd1, 0x57, 0xeb, 0x1b, 0x85, 0x87, 0xe8, 0x45, 0x32, 0x8a, 0x2f,
		0xc8, 0x24, 0x8c, 0x13, 0x72, 0x11, 0x4e, 0xc7, 0x24, 0x9c, 0x5e, 0x8d, 0x26, 0xe1, 0xd8, 0xfd,
		0x02, 0x0f, 0xd0, 0xf3, 0x07, 0xd8, 0xf4, 0x7d, 0x74, 0x39, 0x9a, 0xb8, 0xce, 0x06, 0x28, 0x4e,
		0xc2, 0xb3, 0x8b, 0x6b, 0x77, 0xeb, 0x2d, 0xfb, 0x2f, 0x43, 0xb2, 0xa8, 0xe1, 0xd3, 0x0c, 0xc9,
		0xf5, 0xec, 0x7c, 0x2d, 0xc3, 0x2b, 0xf4, 0xf2, 0x01, 0x36, 0x3e, 0x3f, 0x0b, 0xe3, 0xf0, 0xfd,
		0xd4, 0x75, 0x36, 0x80, 0xa3, 0xb3, 0x24, 0xbc, 0x0a, 0x93, 0x6b, 0x77, 0xeb, 0xf4, 0x0a, 0xbd,
		0xcc, 0x44, 0xb9, 0x49, 0xd1, 0xd3, 0xfe, 0xa8, 0xe6, 0x33, 0x23, 0xc8, 0xcc, 0xf9, 0x23, 0x98,
		0x73, 0x7d, 0xd3, 0xa4, 0x7e, 0x26, 0xca, 0xe0, 0x93, 0x6f, 0xd2, 0x9f, 0x43, 0xd5, 0xfe, 0x5b,
		0xdd, 0x8f, 0xf9, 0x2b, 0xad, 0xf9, 0xed, 0x71, 0xfa, 0xd4, 0xda, 0x7e, 0xfe, 0x37, 0x00, 0x00,
		0xff, 0xff, 0x58, 0x62, 0x2b, 0xbf, 0x55, 0x05, 0x00, 0x00,
	},
	// google/protobuf/wrappers.proto
	[]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4b, 0xcf, 0xcf, 0x4f,
		0xcf, 0x49, 0xd5, 0x2f, 0x28, 0xca, 0x2f, 0xc9, 0x4f, 0x2a, 0x4d, 0xd3, 0x2f, 0x2f, 0x4a, 0x2c,
		0x28, 0x48, 0x2d, 0x2a, 0xd6, 0x03, 0x8b, 0x08, 0xf1, 0x43, 0xe4, 0xf5, 0x60, 0xf2, 0x4a, 0xca,
		0x5c, 0xdc, 0x2e, 0xf9, 0xa5, 0x49, 0x39, 0xa9, 0x61, 0x89, 0x39, 0xa5, 0xa9, 0x42, 0x22, 0x5c,
		0xac, 0x65, 0x20, 0x86, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x63, 0x10, 0x84, 0xa3, 0xa4, 0xc4, 0xc5,
		0xe5, 0x96, 0x93, 0x9f, 0x58, 0x82, 0x45, 0x0d, 0x13, 0x92, 0x1a, 0xcf, 0xbc, 0x12, 0x33, 0x13,
		0x2c, 0x6a, 0x98, 0x61, 0x6a, 0x94, 0xb9, 0xb8, 0x43, 0x71, 0x29, 0x62, 0x41, 0x35, 0xc8, 0xd8,
		0x08, 0x8b, 0x1a, 0x56, 0x34, 0x83, 0xb0, 0x2a, 0xe2, 0x85, 0x29, 0x52, 0xe4, 0xe2, 0x74, 0xca,
		0xcf, 0xcf, 0xc1, 0xa2, 0x84, 0x03, 0xc9, 0x9c, 0xe0, 0x92, 0xa2, 0xcc, 0xbc, 0x74, 0x2c, 0x8a,
		0x38, 0x91, 0x1c, 0xe4, 0x54, 0x59, 0x92, 0x5a, 0x8c, 0x45, 0x0d, 0x0f, 0x54, 0x8d, 0x53, 0x33,
		0x23, 0x97, 0x70, 0x72, 0x7e, 0xae, 0x1e, 0x5a, 0xf0, 0x3a, 0xf1, 0x86, 0x43, 0xc3, 0x3f, 0x00,
		0x24, 0x12, 0xc0, 0x18, 0x65, 0x08, 0x55, 0x91, 0x9e, 0x9f, 0x93, 0x98, 0x97, 0xae, 0x97, 0x5f,
		0x94, 0x8e, 0x88, 0xab, 0x92, 0xca, 0x82, 0xd4, 0x62, 0xfd, 0xec, 0xbc, 0xfc, 0xf2, 0x3c, 0x78,
		0xbc, 0x15, 0x24, 0xfd, 0x60, 0x64, 0x5c, 0xc4, 0xc4, 0xec, 0x1e, 0xe0, 0xb4, 0x8a, 0x49, 0xce,
		0x1d, 0xa2, 0x39, 0x00, 0xaa, 0x43, 0x2f, 0x3c, 0x35, 0x27, 0xc7, 0x1b, 0xa4, 0x3e, 0x04, 0xa4,
		0x35, 0x89, 0x0d, 0x6c, 0x94, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x3c, 0x92, 0x48, 0x30, 0x06,
		0x02, 0x00, 0x00,
	},
}
