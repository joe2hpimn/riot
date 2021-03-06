// Copyright 2013 Hui Chen
// Copyright 2016 ego authors
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

package engine

// NumTokenIndexAdded added token index number
func (engine *Engine) NumTokenIndexAdded() uint64 {
	return engine.numTokenIndexAdded
}

// NumDocumentsIndexed documents indexed number
func (engine *Engine) NumDocumentsIndexed() uint64 {
	return engine.numDocumentsIndexed
}

// NumDocumentsRemoved documents removed number
func (engine *Engine) NumDocumentsRemoved() uint64 {
	return engine.numDocumentsRemoved
}
