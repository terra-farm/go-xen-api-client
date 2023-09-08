//
// This file is generated. To change the content of this file, please do not
// apply the change to this file because it will get overwritten. Instead,
// change xenapi.go and execute 'go generate'.
//

package xenapi

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

type SrHealth string

const (
	// Storage is fully available
	SrHealthHealthy SrHealth = "healthy"
	// Storage is busy recovering, e.g. rebuilding mirrors.
	SrHealthRecovering SrHealth = "recovering"
)

type SrStatRecord struct {
	// Uuid that uniquely identifies this SR, if one is available.
	UUID string
	// Short, human-readable label for the SR.
	NameLabel string
	// Longer, human-readable description of the SR. Descriptions are generally only displayed by clients when the user is examining SRs in detail.
	NameDescription string
	// Number of bytes free on the backing storage (in bytes)
	FreeSpace int
	// Total physical size of the backing storage (in bytes)
	TotalSpace int
	// Indicates whether the SR uses clustered local storage.
	Clustered bool
	// The health status of the SR.
	Health SrHealth
}

type SrStatRef string

// A set of high-level properties associated with an SR.
type SrStatClass struct {
	client *Client
}
