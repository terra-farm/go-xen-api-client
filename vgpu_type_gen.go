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

type VgpuTypeImplementation string

const (
	VgpuTypeImplementationPassthrough VgpuTypeImplementation = "passthrough"
	VgpuTypeImplementationNvidia VgpuTypeImplementation = "nvidia"
	VgpuTypeImplementationGvtG VgpuTypeImplementation = "gvt_g"
)

type VGPUTypeRecord struct {
	UUID string
	VendorName string
	ModelName string
	FramebufferSize int
	MaxHeads int
	MaxResolutionX int
	MaxResolutionY int
	SupportedOnPGPUs []PGPURef
	EnabledOnPGPUs []PGPURef
	VGPUs []VGPURef
	SupportedOnGPUGroups []GPUGroupRef
	EnabledOnGPUGroups []GPUGroupRef
	Implementation VgpuTypeImplementation
	Identifier string
	Experimental bool
}

type VGPUTypeRef string

type VGPUTypeClass struct {
	client *Client
}

func (client *Client) VGPUType() VGPUTypeClass {
	return VGPUTypeClass{client}
}

func (_class VGPUTypeClass) GetAllRecords(sessionID SessionRef) (_retval map[VGPUTypeRef]VGPUTypeRecord, _err error) {
	_method := "VGPU_type.get_all_records"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVGPUTypeRefToVGPUTypeRecordMapToGo(_method + " -> ", _result.Value)
	return
}

func (_class VGPUTypeClass) GetAll(sessionID SessionRef) (_retval []VGPUTypeRef, _err error) {
	_method := "VGPU_type.get_all"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVGPUTypeRefSetToGo(_method + " -> ", _result.Value)
	return
}

func (_class VGPUTypeClass) GetExperimental(sessionID SessionRef, self VGPUTypeRef) (_retval bool, _err error) {
	_method := "VGPU_type.get_experimental"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVGPUTypeRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class VGPUTypeClass) GetIdentifier(sessionID SessionRef, self VGPUTypeRef) (_retval string, _err error) {
	_method := "VGPU_type.get_identifier"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVGPUTypeRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class VGPUTypeClass) GetImplementation(sessionID SessionRef, self VGPUTypeRef) (_retval VgpuTypeImplementation, _err error) {
	_method := "VGPU_type.get_implementation"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVGPUTypeRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertEnumVgpuTypeImplementationToGo(_method + " -> ", _result.Value)
	return
}

func (_class VGPUTypeClass) GetEnabledOnGPUGroups(sessionID SessionRef, self VGPUTypeRef) (_retval []GPUGroupRef, _err error) {
	_method := "VGPU_type.get_enabled_on_GPU_groups"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVGPUTypeRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertGPUGroupRefSetToGo(_method + " -> ", _result.Value)
	return
}

func (_class VGPUTypeClass) GetSupportedOnGPUGroups(sessionID SessionRef, self VGPUTypeRef) (_retval []GPUGroupRef, _err error) {
	_method := "VGPU_type.get_supported_on_GPU_groups"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVGPUTypeRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertGPUGroupRefSetToGo(_method + " -> ", _result.Value)
	return
}

func (_class VGPUTypeClass) GetVGPUs(sessionID SessionRef, self VGPUTypeRef) (_retval []VGPURef, _err error) {
	_method := "VGPU_type.get_VGPUs"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVGPUTypeRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class VGPUTypeClass) GetEnabledOnPGPUs(sessionID SessionRef, self VGPUTypeRef) (_retval []PGPURef, _err error) {
	_method := "VGPU_type.get_enabled_on_PGPUs"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVGPUTypeRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertPGPURefSetToGo(_method + " -> ", _result.Value)
	return
}

func (_class VGPUTypeClass) GetSupportedOnPGPUs(sessionID SessionRef, self VGPUTypeRef) (_retval []PGPURef, _err error) {
	_method := "VGPU_type.get_supported_on_PGPUs"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVGPUTypeRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertPGPURefSetToGo(_method + " -> ", _result.Value)
	return
}

func (_class VGPUTypeClass) GetMaxResolutionY(sessionID SessionRef, self VGPUTypeRef) (_retval int, _err error) {
	_method := "VGPU_type.get_max_resolution_y"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVGPUTypeRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class VGPUTypeClass) GetMaxResolutionX(sessionID SessionRef, self VGPUTypeRef) (_retval int, _err error) {
	_method := "VGPU_type.get_max_resolution_x"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVGPUTypeRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class VGPUTypeClass) GetMaxHeads(sessionID SessionRef, self VGPUTypeRef) (_retval int, _err error) {
	_method := "VGPU_type.get_max_heads"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVGPUTypeRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class VGPUTypeClass) GetFramebufferSize(sessionID SessionRef, self VGPUTypeRef) (_retval int, _err error) {
	_method := "VGPU_type.get_framebuffer_size"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVGPUTypeRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class VGPUTypeClass) GetModelName(sessionID SessionRef, self VGPUTypeRef) (_retval string, _err error) {
	_method := "VGPU_type.get_model_name"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVGPUTypeRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class VGPUTypeClass) GetVendorName(sessionID SessionRef, self VGPUTypeRef) (_retval string, _err error) {
	_method := "VGPU_type.get_vendor_name"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVGPUTypeRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class VGPUTypeClass) GetUUID(sessionID SessionRef, self VGPUTypeRef) (_retval string, _err error) {
	_method := "VGPU_type.get_uuid"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVGPUTypeRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class VGPUTypeClass) GetByUUID(sessionID SessionRef, uuid string) (_retval VGPUTypeRef, _err error) {
	_method := "VGPU_type.get_by_uuid"
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
	_retval, _err = convertVGPUTypeRefToGo(_method + " -> ", _result.Value)
	return
}

func (_class VGPUTypeClass) GetRecord(sessionID SessionRef, self VGPUTypeRef) (_retval VGPUTypeRecord, _err error) {
	_method := "VGPU_type.get_record"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVGPUTypeRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVGPUTypeRecordToGo(_method + " -> ", _result.Value)
	return
}
