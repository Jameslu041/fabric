/*
Copyright IBM Corp. 2016 All Rights Reserved.

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

package util

import (
	"testing"

	"github.com/golang/protobuf/proto"
	"github.com/stretchr/testify/assert"
)

func TestBuffer(t *testing.T) {
	pb := proto.NewBuffer(nil)
	pb.EncodeVarint(10)
	pos1 := len(pb.Bytes())
	pb.EncodeRawBytes([]byte("JunkText"))
	pos2 := len(pb.Bytes())
	pb.EncodeRawBytes([]byte("YetAnotherJunkText"))
	pos3 := len(pb.Bytes())
	pb.EncodeVarint(1000000)
	pos4 := len(pb.Bytes())

	b := NewBuffer(pb.Bytes())
	b.DecodeVarint()
	assert.Equal(t, pos1, b.GetBytesConsumed())
	b.DecodeRawBytes(false)
	assert.Equal(t, pos2, b.GetBytesConsumed())
	b.DecodeRawBytes(false)
	assert.Equal(t, pos3, b.GetBytesConsumed())
	b.DecodeVarint()
	assert.Equal(t, pos4, b.GetBytesConsumed())
}
