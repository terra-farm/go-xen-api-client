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

type SriovConfigurationMode string

const (
  // Configure network sriov by sysfs, do not need reboot
	SriovConfigurationModeSysfs SriovConfigurationMode = "sysfs"
  // Configure network sriov by modprobe, need reboot
	SriovConfigurationModeModprobe SriovConfigurationMode = "modprobe"
  // Unknown mode
	SriovConfigurationModeUnknown SriovConfigurationMode = "unknown"
)

type NetworkSriovRecord struct {
  // Unique identifier/object reference
	UUID string
  // The PIF that has SR-IOV enabled
	PhysicalPIF PIFRef
  // The logical PIF to connect to the SR-IOV network after enable SR-IOV on the physical PIF
	LogicalPIF PIFRef
  // Indicates whether the host need to be rebooted before SR-IOV is enabled on the physical PIF
	RequiresReboot bool
  // The mode for configure network sriov
	ConfigurationMode SriovConfigurationMode
}

type NetworkSriovRef string

// network-sriov which connects logical pif and physical pif
type NetworkSriovClass struct {
	client *Client
}

// Return a map of network_sriov references to network_sriov records for all network_sriovs known to the system.
func (_class NetworkSriovClass) GetAllRecords(sessionID SessionRef) (_retval map[NetworkSriovRef]NetworkSriovRecord, _err error) {
	_method := "network_sriov.get_all_records"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg)
	if _err != nil {
		return
	}
	_retval, _err = convertNetworkSriovRefToNetworkSriovRecordMapToGo(_method + " -> ", _result.Value)
	return
}

// Return a list of all the network_sriovs known to the system.
func (_class NetworkSriovClass) GetAll(sessionID SessionRef) (_retval []NetworkSriovRef, _err error) {
	_method := "network_sriov.get_all"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg)
	if _err != nil {
		return
	}
	_retval, _err = convertNetworkSriovRefSetToGo(_method + " -> ", _result.Value)
	return
}

// Get the number of free SR-IOV VFs on the associated PIF
func (_class NetworkSriovClass) GetRemainingCapacity(sessionID SessionRef, self NetworkSriovRef) (_retval int, _err error) {
	_method := "network_sriov.get_remaining_capacity"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertNetworkSriovRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// Disable SR-IOV on the specific PIF. It will destroy the network-sriov and the logical PIF accordingly.
func (_class NetworkSriovClass) Destroy(sessionID SessionRef, self NetworkSriovRef) (_err error) {
	_method := "network_sriov.destroy"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertNetworkSriovRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg)
	return
}

// Enable SR-IOV on the specific PIF. It will create a network-sriov based on the specific PIF and automatically create a logical PIF to connect the specific network.
func (_class NetworkSriovClass) Create(sessionID SessionRef, pif PIFRef, network NetworkRef) (_retval NetworkSriovRef, _err error) {
	_method := "network_sriov.create"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_pifArg, _err := convertPIFRefToXen(fmt.Sprintf("%s(%s)", _method, "pif"), pif)
	if _err != nil {
		return
	}
	_networkArg, _err := convertNetworkRefToXen(fmt.Sprintf("%s(%s)", _method, "network"), network)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _pifArg, _networkArg)
	if _err != nil {
		return
	}
	_retval, _err = convertNetworkSriovRefToGo(_method + " -> ", _result.Value)
	return
}

// Get the configuration_mode field of the given network_sriov.
func (_class NetworkSriovClass) GetConfigurationMode(sessionID SessionRef, self NetworkSriovRef) (_retval SriovConfigurationMode, _err error) {
	_method := "network_sriov.get_configuration_mode"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertNetworkSriovRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertEnumSriovConfigurationModeToGo(_method + " -> ", _result.Value)
	return
}

// Get the requires_reboot field of the given network_sriov.
func (_class NetworkSriovClass) GetRequiresReboot(sessionID SessionRef, self NetworkSriovRef) (_retval bool, _err error) {
	_method := "network_sriov.get_requires_reboot"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertNetworkSriovRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// Get the logical_PIF field of the given network_sriov.
func (_class NetworkSriovClass) GetLogicalPIF(sessionID SessionRef, self NetworkSriovRef) (_retval PIFRef, _err error) {
	_method := "network_sriov.get_logical_PIF"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertNetworkSriovRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// Get the physical_PIF field of the given network_sriov.
func (_class NetworkSriovClass) GetPhysicalPIF(sessionID SessionRef, self NetworkSriovRef) (_retval PIFRef, _err error) {
	_method := "network_sriov.get_physical_PIF"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertNetworkSriovRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// Get the uuid field of the given network_sriov.
func (_class NetworkSriovClass) GetUUID(sessionID SessionRef, self NetworkSriovRef) (_retval string, _err error) {
	_method := "network_sriov.get_uuid"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertNetworkSriovRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// Get a reference to the network_sriov instance with the specified UUID.
func (_class NetworkSriovClass) GetByUUID(sessionID SessionRef, uuid string) (_retval NetworkSriovRef, _err error) {
	_method := "network_sriov.get_by_uuid"
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
	_retval, _err = convertNetworkSriovRefToGo(_method + " -> ", _result.Value)
	return
}

// Get a record containing the current state of the given network_sriov.
func (_class NetworkSriovClass) GetRecord(sessionID SessionRef, self NetworkSriovRef) (_retval NetworkSriovRecord, _err error) {
	_method := "network_sriov.get_record"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertNetworkSriovRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertNetworkSriovRecordToGo(_method + " -> ", _result.Value)
	return
}
