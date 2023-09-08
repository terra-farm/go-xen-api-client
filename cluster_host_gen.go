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

type ClusterHostOperation string

const (
	// enabling cluster membership on a particular host
	ClusterHostOperationEnable ClusterHostOperation = "enable"
	// disabling cluster membership on a particular host
	ClusterHostOperationDisable ClusterHostOperation = "disable"
	// completely destroying a cluster host
	ClusterHostOperationDestroy ClusterHostOperation = "destroy"
)

type ClusterHostRecord struct {
	// Unique identifier/object reference
	UUID string
	// Reference to the Cluster object
	Cluster ClusterRef
	// Reference to the Host object
	Host HostRef
	// Whether the cluster host believes that clustering should be enabled on this host
	Enabled bool
	// Reference to the PIF object
	PIF PIFRef
	// Whether the cluster host has joined the cluster
	Joined bool
	// list of the operations allowed in this state. This list is advisory only and the server state may have changed by the time this field is read by a client.
	AllowedOperations []ClusterHostOperation
	// links each of the running tasks using this object (by reference) to a current_operation enum which describes the nature of the task.
	CurrentOperations map[string]ClusterHostOperation
	// Additional configuration
	OtherConfig map[string]string
}

type ClusterHostRef string

// Cluster member metadata
type ClusterHostClass struct {
	client *Client
}

// GetAllRecords Return a map of Cluster_host references to Cluster_host records for all Cluster_hosts known to the system.
func (_class ClusterHostClass) GetAllRecords(sessionID SessionRef) (_retval map[ClusterHostRef]ClusterHostRecord, _err error) {
	_method := "Cluster_host.get_all_records"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg)
	if _err != nil {
		return
	}
	_retval, _err = convertClusterHostRefToClusterHostRecordMapToGo(_method + " -> ", _result.Value)
	return
}

// GetAll Return a list of all the Cluster_hosts known to the system.
func (_class ClusterHostClass) GetAll(sessionID SessionRef) (_retval []ClusterHostRef, _err error) {
	_method := "Cluster_host.get_all"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg)
	if _err != nil {
		return
	}
	_retval, _err = convertClusterHostRefSetToGo(_method + " -> ", _result.Value)
	return
}

// Disable Disable cluster membership for an enabled cluster host.
//
// Errors:
//  CLUSTER_STACK_IN_USE - The cluster stack is still in use by at least one plugged PBD.
func (_class ClusterHostClass) Disable(sessionID SessionRef, self ClusterHostRef) (_err error) {
	_method := "Cluster_host.disable"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertClusterHostRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg)
	return
}

// ForceDestroy Remove a host from an existing cluster forcefully.
//
// Errors:
//  CLUSTER_STACK_IN_USE - The cluster stack is still in use by at least one plugged PBD.
func (_class ClusterHostClass) ForceDestroy(sessionID SessionRef, self ClusterHostRef) (_err error) {
	_method := "Cluster_host.force_destroy"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertClusterHostRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg)
	return
}

// Enable Enable cluster membership for a disabled cluster host.
//
// Errors:
//  PIF_ALLOWS_UNPLUG - The operation you requested cannot be performed because the specified PIF allows unplug.
//  REQUIRED_PIF_IS_UNPLUGGED - The operation you requested cannot be performed because the specified PIF is currently unplugged.
func (_class ClusterHostClass) Enable(sessionID SessionRef, self ClusterHostRef) (_err error) {
	_method := "Cluster_host.enable"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertClusterHostRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg)
	return
}

// Destroy Remove a host from an existing cluster.
//
// Errors:
//  CLUSTER_STACK_IN_USE - The cluster stack is still in use by at least one plugged PBD.
//  CLUSTERING_DISABLED - An operation was attempted while clustering was disabled on the cluster_host.
//  CLUSTER_HOST_IS_LAST - The last cluster host cannot be destroyed. Destroy the cluster instead
func (_class ClusterHostClass) Destroy(sessionID SessionRef, self ClusterHostRef) (_err error) {
	_method := "Cluster_host.destroy"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertClusterHostRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg)
	return
}

// Create Add a new host to an existing cluster.
//
// Errors:
//  PIF_NOT_ATTACHED_TO_HOST - Cluster_host creation failed as the PIF provided is not attached to the host.
//  NO_CLUSTER_HOSTS_REACHABLE - No other cluster host was reachable when joining
func (_class ClusterHostClass) Create(sessionID SessionRef, cluster ClusterRef, host HostRef, pif PIFRef) (_retval ClusterHostRef, _err error) {
	_method := "Cluster_host.create"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_clusterArg, _err := convertClusterRefToXen(fmt.Sprintf("%s(%s)", _method, "cluster"), cluster)
	if _err != nil {
		return
	}
	_hostArg, _err := convertHostRefToXen(fmt.Sprintf("%s(%s)", _method, "host"), host)
	if _err != nil {
		return
	}
	_pifArg, _err := convertPIFRefToXen(fmt.Sprintf("%s(%s)", _method, "pif"), pif)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _clusterArg, _hostArg, _pifArg)
	if _err != nil {
		return
	}
	_retval, _err = convertClusterHostRefToGo(_method + " -> ", _result.Value)
	return
}

// GetOtherConfig Get the other_config field of the given Cluster_host.
func (_class ClusterHostClass) GetOtherConfig(sessionID SessionRef, self ClusterHostRef) (_retval map[string]string, _err error) {
	_method := "Cluster_host.get_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertClusterHostRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertStringToStringMapToGo(_method + " -> ", _result.Value)
	return
}

// GetCurrentOperations Get the current_operations field of the given Cluster_host.
func (_class ClusterHostClass) GetCurrentOperations(sessionID SessionRef, self ClusterHostRef) (_retval map[string]ClusterHostOperation, _err error) {
	_method := "Cluster_host.get_current_operations"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertClusterHostRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertStringToEnumClusterHostOperationMapToGo(_method + " -> ", _result.Value)
	return
}

// GetAllowedOperations Get the allowed_operations field of the given Cluster_host.
func (_class ClusterHostClass) GetAllowedOperations(sessionID SessionRef, self ClusterHostRef) (_retval []ClusterHostOperation, _err error) {
	_method := "Cluster_host.get_allowed_operations"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertClusterHostRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertEnumClusterHostOperationSetToGo(_method + " -> ", _result.Value)
	return
}

// GetJoined Get the joined field of the given Cluster_host.
func (_class ClusterHostClass) GetJoined(sessionID SessionRef, self ClusterHostRef) (_retval bool, _err error) {
	_method := "Cluster_host.get_joined"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertClusterHostRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertBoolToGo(_method + " -> ", _result.Value)
	return
}

// GetPIF Get the PIF field of the given Cluster_host.
func (_class ClusterHostClass) GetPIF(sessionID SessionRef, self ClusterHostRef) (_retval PIFRef, _err error) {
	_method := "Cluster_host.get_PIF"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertClusterHostRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertPIFRefToGo(_method + " -> ", _result.Value)
	return
}

// GetEnabled Get the enabled field of the given Cluster_host.
func (_class ClusterHostClass) GetEnabled(sessionID SessionRef, self ClusterHostRef) (_retval bool, _err error) {
	_method := "Cluster_host.get_enabled"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertClusterHostRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertBoolToGo(_method + " -> ", _result.Value)
	return
}

// GetHost Get the host field of the given Cluster_host.
func (_class ClusterHostClass) GetHost(sessionID SessionRef, self ClusterHostRef) (_retval HostRef, _err error) {
	_method := "Cluster_host.get_host"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertClusterHostRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertHostRefToGo(_method + " -> ", _result.Value)
	return
}

// GetCluster Get the cluster field of the given Cluster_host.
func (_class ClusterHostClass) GetCluster(sessionID SessionRef, self ClusterHostRef) (_retval ClusterRef, _err error) {
	_method := "Cluster_host.get_cluster"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertClusterHostRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertClusterRefToGo(_method + " -> ", _result.Value)
	return
}

// GetUUID Get the uuid field of the given Cluster_host.
func (_class ClusterHostClass) GetUUID(sessionID SessionRef, self ClusterHostRef) (_retval string, _err error) {
	_method := "Cluster_host.get_uuid"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertClusterHostRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertStringToGo(_method + " -> ", _result.Value)
	return
}

// GetByUUID Get a reference to the Cluster_host instance with the specified UUID.
func (_class ClusterHostClass) GetByUUID(sessionID SessionRef, uuid string) (_retval ClusterHostRef, _err error) {
	_method := "Cluster_host.get_by_uuid"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_uuidArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "uuid"), uuid)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _uuidArg)
	if _err != nil {
		return
	}
	_retval, _err = convertClusterHostRefToGo(_method + " -> ", _result.Value)
	return
}

// GetRecord Get a record containing the current state of the given Cluster_host.
func (_class ClusterHostClass) GetRecord(sessionID SessionRef, self ClusterHostRef) (_retval ClusterHostRecord, _err error) {
	_method := "Cluster_host.get_record"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertClusterHostRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertClusterHostRecordToGo(_method + " -> ", _result.Value)
	return
}
