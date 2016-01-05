//
// This file is generated. To change the content of this file, please do not
// apply the change to this file because it will get overwritten. Instead,
// change xenapi.go and execute 'go generate'.
//

package xenAPI

import (
	"fmt"
	"github.com/amfranz/go-xmlrpc-client"
	"reflect"
	"strconv"
	"time"
)

var _ = fmt.Errorf
var _ = xmlrpc.NewClient
var _ = reflect.TypeOf
var _ = strconv.Atoi
var _ = time.UTC

type DataSourceRecord struct {
	NameLabel string
	NameDescription string
	Enabled bool
	Standard bool
	Units string
	Min float64
	Max float64
	Value float64
}

type DataSourceRef string

type DataSourceClass struct {
	client *Client
}

func (client *Client) DataSource() DataSourceClass {
	return DataSourceClass{client}
}
