// Copyright 2016 Cyako Author

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//     http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package specvalue

import (
	cyako "github.com/Cyako/Cyako.go"
	"github.com/Cyako/Cyako.go/kvstore"
)

// Specific values include: string/struct sets or lists.

// type Interface interface{
// 	Get(string)
// 	Set(string)
// }

type SpecValue struct {
	kvstore *kvstore.KVStore
}

func (s *SpecValue) SetInt(key string) {

}

func (s *SpecValue) GetInt(key string) {

}

func init() {
	specValue := &SpecValue{
		kvstore: cyako.Ins().Middleware.Map["KVStore"].(*kvstore.KVStore),
	}
	cyako.LoadMiddleware(specValue)
	// type SpecValue2 SpecValue
	// specValue2 := &SpecValue2{
	// 	kvstore: cyako.Ins().Middleware.Map["KVStore"].(*kvstore.KVStore),
	// }
	// cyako.LoadMiddleware(specValue2)
}
