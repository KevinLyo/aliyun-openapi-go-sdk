// Copyright 2015 Chen Xianren. All rights reserved.
// Code generated by openapi-meta; DO NOT EDIT

package batchcompute // import "github.com/cxr29/aliyun-openapi-go-sdk/apis/batchcompute/v20150630"

import (
	"errors"
	"fmt"

	"github.com/cxr29/aliyun-openapi-go-sdk"
)

var (
	_ = errors.New("")
	_ = fmt.Sprint("")
)

type API struct {
	openapi.Service
}

func New(accessKeyId, accessKeySecret string) API {
	return API{openapi.NewService(accessKeyId, accessKeySecret)}
}

func NewParams() openapi.Params {
	args := openapi.NewParams()
	args.Product = Product
	args.Style = Style
	args.Version = Version
	return args
}

const (
	Product = "BatchCompute"
	Style   = "ROA"
	Version = "2015-06-30"
)

// DeleteJob version 2015-06-30
//
// required parameters:
//  name: ResourceName, type: string
//
// optional parameters:
//  name: _region, type: string
//  name: _scheme, type: string, optional values: http|https
func (api API) DeleteJob(ResourceName string, optional openapi.M) (*openapi.Response, error) {
	args := NewParams()

	args.Query.Set("Action", "DeleteJob")
	args.Method = "DELETE"
	args.Path["ResourceName"] = ResourceName
	args.Pattern = `/jobs/[ResourceName]`
	if v, ok := optional["_region"]; ok {
		if s, ok := v.(string); ok {
			args.Region = s
		} else {
			return nil, errors.New("_region must be string")
		}
	}
	if v, ok := optional["_scheme"]; ok {
		if s, ok := v.(string); ok {
			if !openapi.IsIn(s, "http|https") {
				return nil, errors.New("_scheme must be http|https")
			}
			args.Scheme = s
		} else {
			return nil, errors.New("_scheme must be string")
		}
	}

	result := new(openapi.Response)
	if err := api.Service.Do(result, args); err != nil {
		return nil, err
	}
	return result, nil
}

// GetJob version 2015-06-30
//
// required parameters:
//  name: ResourceName, type: string
//
// optional parameters:
//  name: _region, type: string
//  name: _scheme, type: string, optional values: http|https
func (api API) GetJob(ResourceName string, optional openapi.M) (*openapi.Response, error) {
	args := NewParams()

	args.Query.Set("Action", "GetJob")
	args.Method = "GET"
	args.Path["ResourceName"] = ResourceName
	args.Pattern = `/jobs/[ResourceName]`
	if v, ok := optional["_region"]; ok {
		if s, ok := v.(string); ok {
			args.Region = s
		} else {
			return nil, errors.New("_region must be string")
		}
	}
	if v, ok := optional["_scheme"]; ok {
		if s, ok := v.(string); ok {
			if !openapi.IsIn(s, "http|https") {
				return nil, errors.New("_scheme must be http|https")
			}
			args.Scheme = s
		} else {
			return nil, errors.New("_scheme must be string")
		}
	}

	result := new(openapi.Response)
	if err := api.Service.Do(result, args); err != nil {
		return nil, err
	}
	return result, nil
}

// GetJobDescription version 2015-06-30
//
// required parameters:
//  name: ResourceName, type: string
//
// optional parameters:
//  name: _region, type: string
//  name: _scheme, type: string, optional values: http|https
func (api API) GetJobDescription(ResourceName string, optional openapi.M) (*openapi.Response, error) {
	args := NewParams()

	args.Query.Set("Action", "GetJobDescription")
	args.Method = "GET"
	args.Path["ResourceName"] = ResourceName
	args.Pattern = `/jobs/[ResourceName]/description`
	if v, ok := optional["_region"]; ok {
		if s, ok := v.(string); ok {
			args.Region = s
		} else {
			return nil, errors.New("_region must be string")
		}
	}
	if v, ok := optional["_scheme"]; ok {
		if s, ok := v.(string); ok {
			if !openapi.IsIn(s, "http|https") {
				return nil, errors.New("_scheme must be http|https")
			}
			args.Scheme = s
		} else {
			return nil, errors.New("_scheme must be string")
		}
	}

	result := new(openapi.Response)
	if err := api.Service.Do(result, args); err != nil {
		return nil, err
	}
	return result, nil
}

// GetTasks version 2015-06-30
//
// required parameters:
//  name: ResourceName, type: string
//
// optional parameters:
//  name: _region, type: string
//  name: _scheme, type: string, optional values: http|https
func (api API) GetTasks(ResourceName string, optional openapi.M) (*openapi.Response, error) {
	args := NewParams()

	args.Query.Set("Action", "GetTasks")
	args.Method = "GET"
	args.Path["ResourceName"] = ResourceName
	args.Pattern = `/jobs/[ResourceName]/tasks`
	if v, ok := optional["_region"]; ok {
		if s, ok := v.(string); ok {
			args.Region = s
		} else {
			return nil, errors.New("_region must be string")
		}
	}
	if v, ok := optional["_scheme"]; ok {
		if s, ok := v.(string); ok {
			if !openapi.IsIn(s, "http|https") {
				return nil, errors.New("_scheme must be http|https")
			}
			args.Scheme = s
		} else {
			return nil, errors.New("_scheme must be string")
		}
	}

	result := new(openapi.Response)
	if err := api.Service.Do(result, args); err != nil {
		return nil, err
	}
	return result, nil
}

// ListImages version 2015-06-30
//
// optional parameters:
//  name: _region, type: string
//  name: _scheme, type: string, optional values: http|https
func (api API) ListImages(optional openapi.M) (*openapi.Response, error) {
	args := NewParams()

	args.Query.Set("Action", "ListImages")
	args.Method = "GET"
	args.Pattern = `/images`
	if v, ok := optional["_region"]; ok {
		if s, ok := v.(string); ok {
			args.Region = s
		} else {
			return nil, errors.New("_region must be string")
		}
	}
	if v, ok := optional["_scheme"]; ok {
		if s, ok := v.(string); ok {
			if !openapi.IsIn(s, "http|https") {
				return nil, errors.New("_scheme must be http|https")
			}
			args.Scheme = s
		} else {
			return nil, errors.New("_scheme must be string")
		}
	}

	result := new(openapi.Response)
	if err := api.Service.Do(result, args); err != nil {
		return nil, err
	}
	return result, nil
}

// ListJobs version 2015-06-30
//
// optional parameters:
//  name: _region, type: string
//  name: _scheme, type: string, optional values: http|https
func (api API) ListJobs(optional openapi.M) (*openapi.Response, error) {
	args := NewParams()

	args.Query.Set("Action", "ListJobs")
	args.Method = "GET"
	args.Pattern = `/jobs`
	if v, ok := optional["_region"]; ok {
		if s, ok := v.(string); ok {
			args.Region = s
		} else {
			return nil, errors.New("_region must be string")
		}
	}
	if v, ok := optional["_scheme"]; ok {
		if s, ok := v.(string); ok {
			if !openapi.IsIn(s, "http|https") {
				return nil, errors.New("_scheme must be http|https")
			}
			args.Scheme = s
		} else {
			return nil, errors.New("_scheme must be string")
		}
	}

	result := new(openapi.Response)
	if err := api.Service.Do(result, args); err != nil {
		return nil, err
	}
	return result, nil
}

// ListSnapshots version 2015-06-30
//
// optional parameters:
//  name: _region, type: string
//  name: _scheme, type: string, optional values: http|https
func (api API) ListSnapshots(optional openapi.M) (*openapi.Response, error) {
	args := NewParams()

	args.Query.Set("Action", "ListSnapshots")
	args.Method = "GET"
	args.Pattern = `/snapshots`
	if v, ok := optional["_region"]; ok {
		if s, ok := v.(string); ok {
			args.Region = s
		} else {
			return nil, errors.New("_region must be string")
		}
	}
	if v, ok := optional["_scheme"]; ok {
		if s, ok := v.(string); ok {
			if !openapi.IsIn(s, "http|https") {
				return nil, errors.New("_scheme must be http|https")
			}
			args.Scheme = s
		} else {
			return nil, errors.New("_scheme must be string")
		}
	}

	result := new(openapi.Response)
	if err := api.Service.Do(result, args); err != nil {
		return nil, err
	}
	return result, nil
}

// PostJob version 2015-06-30
//
// optional parameters:
//  name: _region, type: string
//  name: _scheme, type: string, optional values: http|https
func (api API) PostJob(optional openapi.M) (*openapi.Response, error) {
	args := NewParams()

	args.Query.Set("Action", "PostJob")
	args.Method = "POST"
	args.Pattern = `/jobs`
	if v, ok := optional["_region"]; ok {
		if s, ok := v.(string); ok {
			args.Region = s
		} else {
			return nil, errors.New("_region must be string")
		}
	}
	if v, ok := optional["_scheme"]; ok {
		if s, ok := v.(string); ok {
			if !openapi.IsIn(s, "http|https") {
				return nil, errors.New("_scheme must be http|https")
			}
			args.Scheme = s
		} else {
			return nil, errors.New("_scheme must be string")
		}
	}

	result := new(openapi.Response)
	if err := api.Service.Do(result, args); err != nil {
		return nil, err
	}
	return result, nil
}

// PutJob version 2015-06-30
//
// required parameters:
//  name: ResourceName, type: string
//
// optional parameters:
//  name: _region, type: string
//  name: _scheme, type: string, optional values: http|https
func (api API) PutJob(ResourceName string, optional openapi.M) (*openapi.Response, error) {
	args := NewParams()

	args.Query.Set("Action", "PutJob")
	args.Method = "PUT"
	args.Path["ResourceName"] = ResourceName
	args.Pattern = `/jobs/[ResourceName]/Priority`
	if v, ok := optional["_region"]; ok {
		if s, ok := v.(string); ok {
			args.Region = s
		} else {
			return nil, errors.New("_region must be string")
		}
	}
	if v, ok := optional["_scheme"]; ok {
		if s, ok := v.(string); ok {
			if !openapi.IsIn(s, "http|https") {
				return nil, errors.New("_scheme must be http|https")
			}
			args.Scheme = s
		} else {
			return nil, errors.New("_scheme must be string")
		}
	}

	result := new(openapi.Response)
	if err := api.Service.Do(result, args); err != nil {
		return nil, err
	}
	return result, nil
}

// ReleaseJob version 2015-06-30
//
// required parameters:
//  name: ResourceName, type: string
//
// optional parameters:
//  name: _region, type: string
//  name: _scheme, type: string, optional values: http|https
func (api API) ReleaseJob(ResourceName string, optional openapi.M) (*openapi.Response, error) {
	args := NewParams()

	args.Query.Set("Action", "ReleaseJob")
	args.Method = "DELETE"
	args.Path["ResourceName"] = ResourceName
	args.Pattern = `/2013-01-11/jobs/[ResourceName]`
	if v, ok := optional["_region"]; ok {
		if s, ok := v.(string); ok {
			args.Region = s
		} else {
			return nil, errors.New("_region must be string")
		}
	}
	if v, ok := optional["_scheme"]; ok {
		if s, ok := v.(string); ok {
			if !openapi.IsIn(s, "http|https") {
				return nil, errors.New("_scheme must be http|https")
			}
			args.Scheme = s
		} else {
			return nil, errors.New("_scheme must be string")
		}
	}

	result := new(openapi.Response)
	if err := api.Service.Do(result, args); err != nil {
		return nil, err
	}
	return result, nil
}

// StartJob version 2015-06-30
//
// required parameters:
//  name: ResourceName, type: string
//
// optional parameters:
//  name: _region, type: string
//  name: _scheme, type: string, optional values: http|https
func (api API) StartJob(ResourceName string, optional openapi.M) (*openapi.Response, error) {
	args := NewParams()

	args.Query.Set("Action", "Start")
	args.Method = "PUT"
	args.Path["ResourceName"] = ResourceName
	args.Pattern = `/jobs/[ResourceName]`
	if v, ok := optional["_region"]; ok {
		if s, ok := v.(string); ok {
			args.Region = s
		} else {
			return nil, errors.New("_region must be string")
		}
	}
	if v, ok := optional["_scheme"]; ok {
		if s, ok := v.(string); ok {
			if !openapi.IsIn(s, "http|https") {
				return nil, errors.New("_scheme must be http|https")
			}
			args.Scheme = s
		} else {
			return nil, errors.New("_scheme must be string")
		}
	}

	result := new(openapi.Response)
	if err := api.Service.Do(result, args); err != nil {
		return nil, err
	}
	return result, nil
}

// StopJob version 2015-06-30
//
// required parameters:
//  name: ResourceName, type: string
//
// optional parameters:
//  name: _region, type: string
//  name: _scheme, type: string, optional values: http|https
func (api API) StopJob(ResourceName string, optional openapi.M) (*openapi.Response, error) {
	args := NewParams()

	args.Query.Set("Action", "Stop")
	args.Method = "PUT"
	args.Path["ResourceName"] = ResourceName
	args.Pattern = `/jobs/[ResourceName]`
	if v, ok := optional["_region"]; ok {
		if s, ok := v.(string); ok {
			args.Region = s
		} else {
			return nil, errors.New("_region must be string")
		}
	}
	if v, ok := optional["_scheme"]; ok {
		if s, ok := v.(string); ok {
			if !openapi.IsIn(s, "http|https") {
				return nil, errors.New("_scheme must be http|https")
			}
			args.Scheme = s
		} else {
			return nil, errors.New("_scheme must be string")
		}
	}

	result := new(openapi.Response)
	if err := api.Service.Do(result, args); err != nil {
		return nil, err
	}
	return result, nil
}
