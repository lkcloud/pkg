/*
 * Copyright © 2019 Lingfei Kong <466701708@qq.com>
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package core

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lkcloud/errors"
)

// Response defines project response format which in lkcloud organization.
type Response struct {
	Code      errors.Code `json:"code"`
	Message   string      `json:"message"`
	RequestID string      `json:"requestID"`
	Data      interface{} `json:"data,omitempty"`
}

// ListResponse defines list request response format which in lkcloud organization.
type ListResponse struct {
	TotalCount uint64      `json:"total_count,omitempty"`
	Items      interface{} `json:"items"`
}

// WriteResponse used to write an error and JSON data into response.
func WriteResponse(c *gin.Context, err error, data interface{}) {
	code, message := errors.DecodeErr(err)

	// always return http.StatusOK
	c.JSON(http.StatusOK, Response{
		Code:    code,
		Message: message,
		Data:    data,
	})
}
