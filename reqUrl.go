package main

import "strings"

type ReqUrl struct {
	base  string
	param map[string]string
}

func NewReqUrl(base string) *ReqUrl {
	var url ReqUrl
	url.base = base
	url.param = make(map[string]string)

	return &url
}

func (req *ReqUrl) Get() string {
	url := req.base

	params := make([]string, len(req.param))
	for k, v := range req.param {
		params = append(params, k+"="+v)
	}

	url += strings.Join(params, "&")

	return url
}

func (req *ReqUrl) SetParam(key string, value string) *ReqUrl {
	req.param[key] = value
	return req
}
