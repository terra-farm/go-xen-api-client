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

type PgpuDom0Access string

const (
	PgpuDom0AccessEnabled PgpuDom0Access = "enabled"
	PgpuDom0AccessDisableOnReboot PgpuDom0Access = "disable_on_reboot"
	PgpuDom0AccessDisabled PgpuDom0Access = "disabled"
	PgpuDom0AccessEnableOnReboot PgpuDom0Access = "enable_on_reboot"
)

type PGPURecord struct {
	UUID string
	PCI PCIRef
	GPUGroup GPUGroupRef
	Host HostRef
	OtherConfig map[string]string
	SupportedVGPUTypes []VGPUTypeRef
	EnabledVGPUTypes []VGPUTypeRef
	ResidentVGPUs []VGPURef
	SupportedVGPUMaxCapacities map[VGPUTypeRef]int
	Dom0Access PgpuDom0Access
	IsSystemDisplayDevice bool
}

type PGPURef string

type PGPUClass struct {
	client *Client
}

func (client *Client) PGPU() PGPUClass {
	return PGPUClass{client}
}

func (_class PGPUClass) GetAllRecords(sessionID SessionRef) (_retval map[PGPURef]PGPURecord, _err error) {
	_method := "PGPU.get_all_records"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg)
	if _err != nil {
		return
	}
	_retval, _err = convertPGPURefToPGPURecordMapToGo(_method + " -> ", _result.Value)
	return
}

func (_class PGPUClass) GetAll(sessionID SessionRef) (_retval []PGPURef, _err error) {
	_method := "PGPU.get_all"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg)
	if _err != nil {
		return
	}
	_retval, _err = convertPGPURefSetToGo(_method + " -> ", _result.Value)
	return
}

func (_class PGPUClass) DisableDom0Access(sessionID SessionRef, self PGPURef) (_retval PgpuDom0Access, _err error) {
	_method := "PGPU.disable_dom0_access"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPGPURefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertEnumPgpuDom0AccessToGo(_method + " -> ", _result.Value)
	return
}

func (_class PGPUClass) EnableDom0Access(sessionID SessionRef, self PGPURef) (_retval PgpuDom0Access, _err error) {
	_method := "PGPU.enable_dom0_access"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPGPURefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertEnumPgpuDom0AccessToGo(_method + " -> ", _result.Value)
	return
}

func (_class PGPUClass) GetRemainingCapacity(sessionID SessionRef, self PGPURef, vgpuType VGPUTypeRef) (_retval int, _err error) {
	_method := "PGPU.get_remaining_capacity"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPGPURefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_vgpuTypeArg, _err := convertVGPUTypeRefToXen(fmt.Sprintf("%s(%s)", _method, "vgpu_type"), vgpuType)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg, _vgpuTypeArg)
	if _err != nil {
		return
	}
	_retval, _err = convertIntToGo(_method + " -> ", _result.Value)
	return
}

func (_class PGPUClass) SetGPUGroup(sessionID SessionRef, self PGPURef, value GPUGroupRef) (_err error) {
	_method := "PGPU.set_GPU_group"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPGPURefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertGPUGroupRefToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}

func (_class PGPUClass) SetEnabledVGPUTypes(sessionID SessionRef, self PGPURef, value []VGPUTypeRef) (_err error) {
	_method := "PGPU.set_enabled_VGPU_types"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPGPURefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertVGPUTypeRefSetToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}

func (_class PGPUClass) RemoveEnabledVGPUTypes(sessionID SessionRef, self PGPURef, value VGPUTypeRef) (_err error) {
	_method := "PGPU.remove_enabled_VGPU_types"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPGPURefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertVGPUTypeRefToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}

func (_class PGPUClass) AddEnabledVGPUTypes(sessionID SessionRef, self PGPURef, value VGPUTypeRef) (_err error) {
	_method := "PGPU.add_enabled_VGPU_types"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPGPURefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertVGPUTypeRefToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}

func (_class PGPUClass) RemoveFromOtherConfig(sessionID SessionRef, self PGPURef, key string) (_err error) {
	_method := "PGPU.remove_from_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPGPURefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class PGPUClass) AddToOtherConfig(sessionID SessionRef, self PGPURef, key string, value string) (_err error) {
	_method := "PGPU.add_to_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPGPURefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class PGPUClass) SetOtherConfig(sessionID SessionRef, self PGPURef, value map[string]string) (_err error) {
	_method := "PGPU.set_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPGPURefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class PGPUClass) GetIsSystemDisplayDevice(sessionID SessionRef, self PGPURef) (_retval bool, _err error) {
	_method := "PGPU.get_is_system_display_device"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPGPURefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class PGPUClass) GetDom0Access(sessionID SessionRef, self PGPURef) (_retval PgpuDom0Access, _err error) {
	_method := "PGPU.get_dom0_access"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPGPURefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertEnumPgpuDom0AccessToGo(_method + " -> ", _result.Value)
	return
}

func (_class PGPUClass) GetSupportedVGPUMaxCapacities(sessionID SessionRef, self PGPURef) (_retval map[VGPUTypeRef]int, _err error) {
	_method := "PGPU.get_supported_VGPU_max_capacities"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPGPURefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVGPUTypeRefToIntMapToGo(_method + " -> ", _result.Value)
	return
}

func (_class PGPUClass) GetResidentVGPUs(sessionID SessionRef, self PGPURef) (_retval []VGPURef, _err error) {
	_method := "PGPU.get_resident_VGPUs"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPGPURefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVGPURefSetToGo(_method + " -> ", _result.Value)
	return
}

func (_class PGPUClass) GetEnabledVGPUTypes(sessionID SessionRef, self PGPURef) (_retval []VGPUTypeRef, _err error) {
	_method := "PGPU.get_enabled_VGPU_types"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPGPURefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVGPUTypeRefSetToGo(_method + " -> ", _result.Value)
	return
}

func (_class PGPUClass) GetSupportedVGPUTypes(sessionID SessionRef, self PGPURef) (_retval []VGPUTypeRef, _err error) {
	_method := "PGPU.get_supported_VGPU_types"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPGPURefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVGPUTypeRefSetToGo(_method + " -> ", _result.Value)
	return
}

func (_class PGPUClass) GetOtherConfig(sessionID SessionRef, self PGPURef) (_retval map[string]string, _err error) {
	_method := "PGPU.get_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPGPURefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class PGPUClass) GetHost(sessionID SessionRef, self PGPURef) (_retval HostRef, _err error) {
	_method := "PGPU.get_host"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPGPURefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class PGPUClass) GetGPUGroup(sessionID SessionRef, self PGPURef) (_retval GPUGroupRef, _err error) {
	_method := "PGPU.get_GPU_group"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPGPURefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertGPUGroupRefToGo(_method + " -> ", _result.Value)
	return
}

func (_class PGPUClass) GetPCI(sessionID SessionRef, self PGPURef) (_retval PCIRef, _err error) {
	_method := "PGPU.get_PCI"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPGPURefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertPCIRefToGo(_method + " -> ", _result.Value)
	return
}

func (_class PGPUClass) GetUUID(sessionID SessionRef, self PGPURef) (_retval string, _err error) {
	_method := "PGPU.get_uuid"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPGPURefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class PGPUClass) GetByUUID(sessionID SessionRef, uuid string) (_retval PGPURef, _err error) {
	_method := "PGPU.get_by_uuid"
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
	_retval, _err = convertPGPURefToGo(_method + " -> ", _result.Value)
	return
}

func (_class PGPUClass) GetRecord(sessionID SessionRef, self PGPURef) (_retval PGPURecord, _err error) {
	_method := "PGPU.get_record"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPGPURefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertPGPURecordToGo(_method + " -> ", _result.Value)
	return
}
