//  Copyright (c) 2017-2018 Uber Technologies, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package common

import (
	memCom "github.com/uber/aresdb/memstore/common"
	queryCom "github.com/uber/aresdb/query/common"
	"io"
	"net/http"
)

// SchemaManager keeps table schema from all namespaces in current ENV up to date
type SchemaManager interface {
	// Run initializes all schema and starts jobs to sync from controller
	Run()
	// GetTable gets schema by namespace and table name
	GetTableSchemaReader(namespace string) (memCom.TableSchemaReader, error)
}

// QueryExecutor defines query executor
type QueryExecutor interface {
	// Execute executes query and flush result to connection
	Execute(namespace, sqlQuery string, w http.ResponseWriter) (err error)
}

// BlockingPlanNode defines query plan nodes that waits for children to finish
type BlockingPlanNode interface {
	Run() (queryCom.AQLQueryResult, error)
	Children() []BlockingPlanNode
	Add(...BlockingPlanNode)
}

// StreamingPlanNode defines query plan nodes that eager flushes to output
type StreamingPlanNode interface {
	Run(writer io.Writer) error
}