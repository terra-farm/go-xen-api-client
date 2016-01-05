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

type HostPatchRecord struct {
	UUID string
	NameLabel string
	NameDescription string
	Version string
	Host HostRef
	Applied bool
	TimestampApplied time.Time
	Size int
	PoolPatch PoolPatchRef
	OtherConfig map[string]string
}

type HostPatchRef string

type HostPatchClass struct {
	client *Client
}

func (client *Client) HostPatch() HostPatchClass {
	return HostPatchClass{client}
}

func (_class HostPatchClass) GetAllRecords(sessionID SessionRef) (_retval map[HostPatchRef]HostPatchRecord, _err error) {
	_method := "host_patch.get_all_records"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg)
	if _err != nil {
		return
	}
	_retval, _err = convertHostPatchRefToHostPatchRecordMapToGo(_method + " -> ", _result.Value)
	return
}

func (_class HostPatchClass) GetAll(sessionID SessionRef) (_retval []HostPatchRef, _err error) {
	_method := "host_patch.get_all"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg)
	if _err != nil {
		return
	}
	_retval, _err = convertHostPatchRefSetToGo(_method + " -> ", _result.Value)
	return
}

func (_class HostPatchClass) Apply(sessionID SessionRef, self HostPatchRef) (_retval string, _err error) {
	_method := "host_patch.apply"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertHostPatchRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class HostPatchClass) Destroy(sessionID SessionRef, self HostPatchRef) (_err error) {
	_method := "host_patch.destroy"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertHostPatchRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg)
	return
}

func (_class HostPatchClass) RemoveFromOtherConfig(sessionID SessionRef, self HostPatchRef, key string) (_err error) {
	_method := "host_patch.remove_from_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertHostPatchRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class HostPatchClass) AddToOtherConfig(sessionID SessionRef, self HostPatchRef, key string, value string) (_err error) {
	_method := "host_patch.add_to_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertHostPatchRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class HostPatchClass) SetOtherConfig(sessionID SessionRef, self HostPatchRef, value map[string]string) (_err error) {
	_method := "host_patch.set_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertHostPatchRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class HostPatchClass) GetOtherConfig(sessionID SessionRef, self HostPatchRef) (_retval map[string]string, _err error) {
	_method := "host_patch.get_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertHostPatchRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class HostPatchClass) GetPoolPatch(sessionID SessionRef, self HostPatchRef) (_retval PoolPatchRef, _err error) {
	_method := "host_patch.get_pool_patch"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertHostPatchRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertPoolPatchRefToGo(_method + " -> ", _result.Value)
	return
}

func (_class HostPatchClass) GetSize(sessionID SessionRef, self HostPatchRef) (_retval int, _err error) {
	_method := "host_patch.get_size"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertHostPatchRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class HostPatchClass) GetTimestampApplied(sessionID SessionRef, self HostPatchRef) (_retval time.Time, _err error) {
	_method := "host_patch.get_timestamp_applied"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertHostPatchRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTimeToGo(_method + " -> ", _result.Value)
	return
}

func (_class HostPatchClass) GetApplied(sessionID SessionRef, self HostPatchRef) (_retval bool, _err error) {
	_method := "host_patch.get_applied"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertHostPatchRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class HostPatchClass) GetHost(sessionID SessionRef, self HostPatchRef) (_retval HostRef, _err error) {
	_method := "host_patch.get_host"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertHostPatchRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class HostPatchClass) GetVersion(sessionID SessionRef, self HostPatchRef) (_retval string, _err error) {
	_method := "host_patch.get_version"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertHostPatchRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class HostPatchClass) GetNameDescription(sessionID SessionRef, self HostPatchRef) (_retval string, _err error) {
	_method := "host_patch.get_name_description"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertHostPatchRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class HostPatchClass) GetNameLabel(sessionID SessionRef, self HostPatchRef) (_retval string, _err error) {
	_method := "host_patch.get_name_label"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertHostPatchRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class HostPatchClass) GetUUID(sessionID SessionRef, self HostPatchRef) (_retval string, _err error) {
	_method := "host_patch.get_uuid"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertHostPatchRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class HostPatchClass) GetByNameLabel(sessionID SessionRef, label string) (_retval []HostPatchRef, _err error) {
	_method := "host_patch.get_by_name_label"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_labelArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "label"), label)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _labelArg)
	if _err != nil {
		return
	}
	_retval, _err = convertHostPatchRefSetToGo(_method + " -> ", _result.Value)
	return
}

func (_class HostPatchClass) GetByUUID(sessionID SessionRef, uuid string) (_retval HostPatchRef, _err error) {
	_method := "host_patch.get_by_uuid"
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
	_retval, _err = convertHostPatchRefToGo(_method + " -> ", _result.Value)
	return
}

func (_class HostPatchClass) GetRecord(sessionID SessionRef, self HostPatchRef) (_retval HostPatchRecord, _err error) {
	_method := "host_patch.get_record"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertHostPatchRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertHostPatchRecordToGo(_method + " -> ", _result.Value)
	return
}
