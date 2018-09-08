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

type ClusterOperation string

const (
  // adding a new member to the cluster
	ClusterOperationAdd ClusterOperation = "add"
  // removing a member from the cluster
	ClusterOperationRemove ClusterOperation = "remove"
  // enabling any cluster member
	ClusterOperationEnable ClusterOperation = "enable"
  // disabling any cluster member
	ClusterOperationDisable ClusterOperation = "disable"
  // completely destroying a cluster
	ClusterOperationDestroy ClusterOperation = "destroy"
)

type ClusterRecord struct {
  // Unique identifier/object reference
	UUID string
  // A list of the cluster_host objects associated with the Cluster
	ClusterHosts []ClusterHostRef
  // Reference to the single network on which corosync carries out its inter-host communications
	Network NetworkRef
  // The secret key used by xapi-clusterd when it talks to itself on other hosts
	ClusterToken string
  // Simply the string 'corosync'. No other cluster stacks are currently supported
	ClusterStack string
  // list of the operations allowed in this state. This list is advisory only and the server state may have changed by the time this field is read by a client.
	AllowedOperations []ClusterOperation
  // links each of the running tasks using this object (by reference) to a current_operation enum which describes the nature of the task.
	CurrentOperations map[string]ClusterOperation
  // True if xapi is automatically joining new pool members to the cluster. This will be `true` in the first release
	PoolAutoJoin bool
  // The corosync token timeout in ms
	TokenTimeout int
  // The corosync token timeout coefficient in ms
	TokenTimeoutCoefficient int
  // Contains read-only settings for the cluster, such as timeouts and other options. It can only be set at cluster create time
	ClusterConfig map[string]string
  // Additional configuration
	OtherConfig map[string]string
}

type ClusterRef string

// Cluster-wide Cluster metadata
type ClusterClass struct {
	client *Client
}

// Return a map of Cluster references to Cluster records for all Clusters known to the system.
func (_class ClusterClass) GetAllRecords(sessionID SessionRef) (_retval map[ClusterRef]ClusterRecord, _err error) {
	_method := "Cluster.get_all_records"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg)
	if _err != nil {
		return
	}
	_retval, _err = convertClusterRefToClusterRecordMapToGo(_method + " -> ", _result.Value)
	return
}

// Return a list of all the Clusters known to the system.
func (_class ClusterClass) GetAll(sessionID SessionRef) (_retval []ClusterRef, _err error) {
	_method := "Cluster.get_all"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg)
	if _err != nil {
		return
	}
	_retval, _err = convertClusterRefSetToGo(_method + " -> ", _result.Value)
	return
}

// Resynchronise the cluster_host objects across the pool. Creates them where they need creating and then plugs them
func (_class ClusterClass) PoolResync(sessionID SessionRef, self ClusterRef) (_err error) {
	_method := "Cluster.pool_resync"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertClusterRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg)
	return
}

// Attempt to destroy the Cluster_host objects for all hosts in the pool and then destroy the Cluster.
func (_class ClusterClass) PoolDestroy(sessionID SessionRef, self ClusterRef) (_err error) {
	_method := "Cluster.pool_destroy"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertClusterRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg)
	return
}

// Attempt to force destroy the Cluster_host objects, and then destroy the Cluster.
func (_class ClusterClass) PoolForceDestroy(sessionID SessionRef, self ClusterRef) (_err error) {
	_method := "Cluster.pool_force_destroy"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertClusterRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg)
	return
}

// Attempt to create a Cluster from the entire pool
func (_class ClusterClass) PoolCreate(sessionID SessionRef, network NetworkRef, clusterStack string, tokenTimeout float64, tokenTimeoutCoefficient float64) (_retval ClusterRef, _err error) {
	_method := "Cluster.pool_create"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_networkArg, _err := convertNetworkRefToXen(fmt.Sprintf("%s(%s)", _method, "network"), network)
	if _err != nil {
		return
	}
	_clusterStackArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "cluster_stack"), clusterStack)
	if _err != nil {
		return
	}
	_tokenTimeoutArg, _err := convertFloatToXen(fmt.Sprintf("%s(%s)", _method, "token_timeout"), tokenTimeout)
	if _err != nil {
		return
	}
	_tokenTimeoutCoefficientArg, _err := convertFloatToXen(fmt.Sprintf("%s(%s)", _method, "token_timeout_coefficient"), tokenTimeoutCoefficient)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _networkArg, _clusterStackArg, _tokenTimeoutArg, _tokenTimeoutCoefficientArg)
	if _err != nil {
		return
	}
	_retval, _err = convertClusterRefToGo(_method + " -> ", _result.Value)
	return
}

// Destroys a Cluster object and the one remaining Cluster_host member
func (_class ClusterClass) Destroy(sessionID SessionRef, self ClusterRef) (_err error) {
	_method := "Cluster.destroy"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertClusterRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg)
	return
}

// Creates a Cluster object and one Cluster_host object as its first member
func (_class ClusterClass) Create(sessionID SessionRef, network NetworkRef, clusterStack string, poolAutoJoin bool, tokenTimeout float64, tokenTimeoutCoefficient float64) (_retval ClusterRef, _err error) {
	_method := "Cluster.create"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_networkArg, _err := convertNetworkRefToXen(fmt.Sprintf("%s(%s)", _method, "network"), network)
	if _err != nil {
		return
	}
	_clusterStackArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "cluster_stack"), clusterStack)
	if _err != nil {
		return
	}
	_poolAutoJoinArg, _err := convertBoolToXen(fmt.Sprintf("%s(%s)", _method, "pool_auto_join"), poolAutoJoin)
	if _err != nil {
		return
	}
	_tokenTimeoutArg, _err := convertFloatToXen(fmt.Sprintf("%s(%s)", _method, "token_timeout"), tokenTimeout)
	if _err != nil {
		return
	}
	_tokenTimeoutCoefficientArg, _err := convertFloatToXen(fmt.Sprintf("%s(%s)", _method, "token_timeout_coefficient"), tokenTimeoutCoefficient)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _networkArg, _clusterStackArg, _poolAutoJoinArg, _tokenTimeoutArg, _tokenTimeoutCoefficientArg)
	if _err != nil {
		return
	}
	_retval, _err = convertClusterRefToGo(_method + " -> ", _result.Value)
	return
}

// Remove the given key and its corresponding value from the other_config field of the given Cluster.  If the key is not in that Map, then do nothing.
func (_class ClusterClass) RemoveFromOtherConfig(sessionID SessionRef, self ClusterRef, key string) (_err error) {
	_method := "Cluster.remove_from_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertClusterRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_keyArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "key"), key)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _keyArg)
	return
}

// Add the given key-value pair to the other_config field of the given Cluster.
func (_class ClusterClass) AddToOtherConfig(sessionID SessionRef, self ClusterRef, key string, value string) (_err error) {
	_method := "Cluster.add_to_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertClusterRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_keyArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "key"), key)
	if _err != nil {
		return
	}
	_valueArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _keyArg, _valueArg)
	return
}

// Set the other_config field of the given Cluster.
func (_class ClusterClass) SetOtherConfig(sessionID SessionRef, self ClusterRef, value map[string]string) (_err error) {
	_method := "Cluster.set_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertClusterRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertStringToStringMapToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}

// Get the other_config field of the given Cluster.
func (_class ClusterClass) GetOtherConfig(sessionID SessionRef, self ClusterRef) (_retval map[string]string, _err error) {
	_method := "Cluster.get_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertClusterRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// Get the cluster_config field of the given Cluster.
func (_class ClusterClass) GetClusterConfig(sessionID SessionRef, self ClusterRef) (_retval map[string]string, _err error) {
	_method := "Cluster.get_cluster_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertClusterRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// Get the token_timeout_coefficient field of the given Cluster.
func (_class ClusterClass) GetTokenTimeoutCoefficient(sessionID SessionRef, self ClusterRef) (_retval int, _err error) {
	_method := "Cluster.get_token_timeout_coefficient"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertClusterRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertIntToGo(_method + " -> ", _result.Value)
	return
}

// Get the token_timeout field of the given Cluster.
func (_class ClusterClass) GetTokenTimeout(sessionID SessionRef, self ClusterRef) (_retval int, _err error) {
	_method := "Cluster.get_token_timeout"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertClusterRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertIntToGo(_method + " -> ", _result.Value)
	return
}

// Get the pool_auto_join field of the given Cluster.
func (_class ClusterClass) GetPoolAutoJoin(sessionID SessionRef, self ClusterRef) (_retval bool, _err error) {
	_method := "Cluster.get_pool_auto_join"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertClusterRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// Get the current_operations field of the given Cluster.
func (_class ClusterClass) GetCurrentOperations(sessionID SessionRef, self ClusterRef) (_retval map[string]ClusterOperation, _err error) {
	_method := "Cluster.get_current_operations"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertClusterRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertStringToEnumClusterOperationMapToGo(_method + " -> ", _result.Value)
	return
}

// Get the allowed_operations field of the given Cluster.
func (_class ClusterClass) GetAllowedOperations(sessionID SessionRef, self ClusterRef) (_retval []ClusterOperation, _err error) {
	_method := "Cluster.get_allowed_operations"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertClusterRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertEnumClusterOperationSetToGo(_method + " -> ", _result.Value)
	return
}

// Get the cluster_stack field of the given Cluster.
func (_class ClusterClass) GetClusterStack(sessionID SessionRef, self ClusterRef) (_retval string, _err error) {
	_method := "Cluster.get_cluster_stack"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertClusterRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// Get the cluster_token field of the given Cluster.
func (_class ClusterClass) GetClusterToken(sessionID SessionRef, self ClusterRef) (_retval string, _err error) {
	_method := "Cluster.get_cluster_token"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertClusterRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// Get the network field of the given Cluster.
func (_class ClusterClass) GetNetwork(sessionID SessionRef, self ClusterRef) (_retval NetworkRef, _err error) {
	_method := "Cluster.get_network"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertClusterRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertNetworkRefToGo(_method + " -> ", _result.Value)
	return
}

// Get the cluster_hosts field of the given Cluster.
func (_class ClusterClass) GetClusterHosts(sessionID SessionRef, self ClusterRef) (_retval []ClusterHostRef, _err error) {
	_method := "Cluster.get_cluster_hosts"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertClusterRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertClusterHostRefSetToGo(_method + " -> ", _result.Value)
	return
}

// Get the uuid field of the given Cluster.
func (_class ClusterClass) GetUUID(sessionID SessionRef, self ClusterRef) (_retval string, _err error) {
	_method := "Cluster.get_uuid"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertClusterRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// Get a reference to the Cluster instance with the specified UUID.
func (_class ClusterClass) GetByUUID(sessionID SessionRef, uuid string) (_retval ClusterRef, _err error) {
	_method := "Cluster.get_by_uuid"
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
	_retval, _err = convertClusterRefToGo(_method + " -> ", _result.Value)
	return
}

// Get a record containing the current state of the given Cluster.
func (_class ClusterClass) GetRecord(sessionID SessionRef, self ClusterRef) (_retval ClusterRecord, _err error) {
	_method := "Cluster.get_record"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertClusterRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertClusterRecordToGo(_method + " -> ", _result.Value)
	return
}
