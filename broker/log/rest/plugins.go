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

// Package rest exposes a simple REST API for communicating with the GRPC package.
package rest

import (
	"github.com/pydio/cells/common"
	"github.com/pydio/cells/common/service"
)

func init() {
	service.NewService(
		service.Name(common.SERVICE_REST_NAMESPACE_+common.SERVICE_LOG),
		service.Tag(common.SERVICE_TAG_BROKER),
		service.Description("RESTful Gateway to search in the log repositories"),
		service.WithWeb(func() service.WebHandler {
			return new(Handler)
		}),
	)
}
