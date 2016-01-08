//go:generate go run xenapi.go

package xenAPI

import (
	"crypto/tls"
	"fmt"
	"github.com/amfranz/go-xmlrpc-client"
	"net/http"
)

type APIResult struct {
	Value interface{}
}

func (client *Client) APICall(method string, params ...interface{}) (result APIResult, err error) {
	rpcParams := xmlrpc.Params{
		Params: params,
	}

	rpcResult := xmlrpc.Struct{}

	err = client.rpc.Call(method, rpcParams, &rpcResult)
	if err != nil {
		return
	}

	status, ok := rpcResult["Status"].(string)
	if !ok {
		err = fmt.Errorf("Expected a field named %q with a string value in the response", "Status")
		return
	}

	if status != "Success" {
		details := rpcResult["ErrorDescription"].([]interface{})
		err = &Error{
			code:    details[0].(string),
			objtype: details[1].(string), // might be nil
			uuid:    details[2].(string), // optional
		}
		return
	}

	result.Value = rpcResult["Value"]
	return
}

func NewClient(url string, transport *http.Transport) (*Client, error) {
	if transport == nil {
		transport = &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
	}

	rpc, err := xmlrpc.NewClient(url, transport)
	if err != nil {
		return nil, err
	}

	return prepClient(rpc), nil
}
