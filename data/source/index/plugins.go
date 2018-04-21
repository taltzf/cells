/*
 * Copyright (c) 2018. Abstrium SAS <team (at) pydio.com>
 * This file is part of Pydio Cells.
 *
 * Pydio Cells is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * Pydio Cells is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with Pydio Cells.  If not, see <http://www.gnu.org/licenses/>.
 *
 * The latest code can be found at <https://pydio.com>.
 */

package index

import (
	"github.com/pydio/cells/common"
	"github.com/pydio/cells/common/service"
)

var (
	// Name of the current plugin
	Name        = common.SERVICE_GRPC_NAMESPACE_ + common.SERVICE_DATA_INDEX
	ChildPrefix = common.SERVICE_GRPC_NAMESPACE_ + common.SERVICE_DATA_INDEX_
)

func init() {
	service.NewService(
		service.Name(Name),
		service.Tag(common.SERVICE_TAG_DATASOURCE),
		service.Description("Starter for data sources indexes"),
		service.Dependency(common.SERVICE_GRPC_NAMESPACE_+common.SERVICE_TREE, []string{}),
		service.WithMicroChildrenRunner(Name, ChildPrefix, true, nil),
	)
}
