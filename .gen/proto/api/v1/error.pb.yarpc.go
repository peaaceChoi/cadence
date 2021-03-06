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
// source: uber/cadence/api/v1/error.proto

package apiv1

var yarpcFileDescriptorClosurec8f91786c9aff272 = [][]byte{
	// uber/cadence/api/v1/error.proto
	[]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x41, 0x6f, 0xd3, 0x40,
		0x10, 0x85, 0x95, 0x22, 0x22, 0x18, 0x44, 0x5b, 0x4c, 0x80, 0x8a, 0x03, 0x05, 0x03, 0x6a, 0x2f,
		0xd8, 0x8a, 0x38, 0x72, 0x4a, 0x82, 0x91, 0x22, 0xa1, 0xaa, 0x34, 0x52, 0x91, 0xb8, 0x58, 0x9b,
		0xf5, 0x24, 0x8c, 0xb0, 0x77, 0xcd, 0xec, 0xae, 0x49, 0x2e, 0xfc, 0x0f, 0xfe, 0x2d, 0xb2, 0x77,
		0x5d, 0x82, 0x94, 0x43, 0x8f, 0xfb, 0xcd, 0xdb, 0x37, 0xb3, 0x3b, 0x0f, 0x4e, 0xdd, 0x12, 0x39,
		0x95, 0xa2, 0x40, 0x25, 0x31, 0x15, 0x35, 0xa5, 0xcd, 0x38, 0x45, 0x66, 0xcd, 0x49, 0xcd, 0xda,
		0xea, 0xe8, 0x71, 0x2b, 0x48, 0x82, 0x20, 0x11, 0x35, 0x25, 0xcd, 0x38, 0x5e, 0xc3, 0x9b, 0xaf,
		0x9a, 0x7f, 0xac, 0x4a, 0xfd, 0x2b, 0xdb, 0xa0, 0x74, 0x96, 0xb4, 0x9a, 0x94, 0x8c, 0xa2, 0xd8,
		0x2e, 0xac, 0x60, 0x8b, 0x45, 0xd6, 0x5a, 0x44, 0xe7, 0x70, 0x6c, 0xda, 0x73, 0xce, 0xf8, 0xd3,
		0xa1, 0xb1, 0x39, 0x15, 0x27, 0x83, 0x97, 0x83, 0xf3, 0xfb, 0x57, 0x87, 0x1d, 0xbf, 0xf2, 0x78,
		0x5e, 0x44, 0x4f, 0x60, 0xc8, 0x4e, 0xb5, 0xf5, 0x83, 0xae, 0x7e, 0x97, 0x9d, 0x9a, 0x17, 0xf1,
		0x0a, 0x46, 0x99, 0xb2, 0x64, 0xb7, 0x17, 0xda, 0x66, 0x1b, 0x32, 0xd6, 0x78, 0xe3, 0x33, 0x38,
		0x92, 0x8e, 0x19, 0x95, 0xcd, 0x65, 0xe9, 0x8c, 0x45, 0xee, 0x7d, 0x03, 0x9e, 0x79, 0x1a, 0xbd,
		0x85, 0x43, 0x21, 0x2d, 0x35, 0x78, 0xa3, 0xf3, 0xfe, 0x0f, 0x3d, 0x0d, 0xb2, 0xf8, 0x37, 0x8c,
		0x3e, 0xea, 0x4a, 0x90, 0xba, 0xd0, 0x76, 0xd2, 0x55, 0x7c, 0x9f, 0xa7, 0x30, 0x2c, 0x3a, 0x1e,
		0xec, 0xc3, 0x69, 0x5f, 0xff, 0x83, 0x5b, 0xf6, 0xbf, 0xb3, 0xaf, 0xff, 0x9f, 0x01, 0xbc, 0x98,
		0x95, 0x84, 0xca, 0x5e, 0x23, 0x1b, 0xd2, 0xed, 0x1c, 0x0b, 0x57, 0xd7, 0xfa, 0xdf, 0x5f, 0x9e,
		0xc1, 0xd1, 0x0a, 0x85, 0x75, 0x8c, 0x79, 0xe3, 0x35, 0xfd, 0x93, 0x03, 0x0e, 0x37, 0xa3, 0x53,
		0x78, 0x20, 0x3b, 0xab, 0x9c, 0xaa, 0xba, 0x0c, 0x73, 0x81, 0x47, 0xf3, 0xaa, 0x2e, 0xa3, 0x77,
		0x10, 0x99, 0xde, 0xbb, 0xf7, 0x32, 0x61, 0xae, 0x47, 0x37, 0x95, 0x60, 0x67, 0xe2, 0xd7, 0xf0,
		0x6a, 0x26, 0x94, 0xc4, 0xb2, 0x14, 0x3b, 0x7b, 0x0e, 0xab, 0x0b, 0xd3, 0xc5, 0xcf, 0xe1, 0xc4,
		0x7f, 0x60, 0x28, 0xef, 0x2c, 0x2b, 0x1e, 0x41, 0xf4, 0x99, 0x2a, 0xb2, 0xd9, 0x46, 0x22, 0x16,
		0xfd, 0x8d, 0x08, 0x8e, 0xbf, 0x38, 0xe4, 0xed, 0x27, 0x41, 0xe5, 0x0e, 0x5b, 0x20, 0x37, 0x24,
		0x71, 0xea, 0xcc, 0xb6, 0x63, 0xd3, 0x6b, 0x78, 0x26, 0x75, 0x95, 0xec, 0x89, 0xe1, 0xf4, 0xde,
		0xa4, 0xa6, 0xcb, 0x36, 0xa5, 0x97, 0x83, 0x6f, 0xe9, 0x9a, 0xec, 0x77, 0xb7, 0x4c, 0xa4, 0xae,
		0xd2, 0xff, 0x32, 0x9d, 0xac, 0x51, 0xa5, 0x5d, 0x96, 0x43, 0xbc, 0x3f, 0x88, 0x9a, 0x9a, 0xf1,
		0x72, 0xd8, 0xb1, 0xf7, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0xf2, 0x05, 0xa7, 0xdf, 0x02, 0x03,
		0x00, 0x00,
	},
}
