package kernel

import "github.com/aws/jsii-runtime-go/internal/api"

type GetProps struct {
	Property   string        `json:"property"`
	ObjRef     api.ObjectRef `json:"objref"`
	StackTrace []string			 `json:"stacktrace"`
}

type StaticGetProps struct {
	FQN        api.FQN  `json:"fqn"`
	Property   string   `json:"property"`
	StackTrace []string `json:"stacktrace"`
}

type GetResponse struct {
	kernelResponse
	Value interface{} `json:"value"`
}

func (c *Client) Get(props GetProps) (response GetResponse, err error) {
	type request struct {
		kernelRequest
		GetProps
	}
	props.StackTrace = captureStack()
	err = c.request(request{kernelRequest{"get"}, props}, &response)
	return
}

func (c *Client) SGet(props StaticGetProps) (response GetResponse, err error) {
	type request struct {
		kernelRequest
		StaticGetProps
	}
	props.StackTrace = captureStack()
	err = c.request(request{kernelRequest{"sget"}, props}, &response)
	return
}

// UnmarshalJSON provides custom unmarshalling implementation for response
// structs. Creating new types is required in order to avoid infinite recursion.
func (r *GetResponse) UnmarshalJSON(data []byte) error {
	type response GetResponse
	return unmarshalKernelResponse(data, (*response)(r), r)
}
