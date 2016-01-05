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

type BondMode string

const (
	BondModeBalanceSlb BondMode = "balance-slb"
	BondModeActiveBackup BondMode = "active-backup"
	BondModeLacp BondMode = "lacp"
)

type BondRecord struct {
	UUID string
	Master PIFRef
	Slaves []PIFRef
	OtherConfig map[string]string
	PrimarySlave PIFRef
	Mode BondMode
	Properties map[string]string
	LinksUp int
}

type BondRef string

type BondClass struct {
	client *Client
}

func (client *Client) Bond() BondClass {
	return BondClass{client}
}

func (_class BondClass) GetAllRecords(sessionID SessionRef) (_retval map[BondRef]BondRecord, _err error) {
	_method := "Bond.get_all_records"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg)
	if _err != nil {
		return
	}
	_retval, _err = convertBondRefToBondRecordMapToGo(_method + " -> ", _result.Value)
	return
}

func (_class BondClass) GetAll(sessionID SessionRef) (_retval []BondRef, _err error) {
	_method := "Bond.get_all"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg)
	if _err != nil {
		return
	}
	_retval, _err = convertBondRefSetToGo(_method + " -> ", _result.Value)
	return
}

func (_class BondClass) SetProperty(sessionID SessionRef, self BondRef, name string, value string) (_err error) {
	_method := "Bond.set_property"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertBondRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_nameArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "name"), name)
	if _err != nil {
		return
	}
	_valueArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _nameArg, _valueArg)
	return
}

func (_class BondClass) SetMode(sessionID SessionRef, self BondRef, value BondMode) (_err error) {
	_method := "Bond.set_mode"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertBondRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertEnumBondModeToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}

func (_class BondClass) Destroy(sessionID SessionRef, self BondRef) (_err error) {
	_method := "Bond.destroy"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertBondRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg)
	return
}

func (_class BondClass) Create(sessionID SessionRef, network NetworkRef, members []PIFRef, mac string, mode BondMode, properties map[string]string) (_retval BondRef, _err error) {
	_method := "Bond.create"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_networkArg, _err := convertNetworkRefToXen(fmt.Sprintf("%s(%s)", _method, "network"), network)
	if _err != nil {
		return
	}
	_membersArg, _err := convertPIFRefSetToXen(fmt.Sprintf("%s(%s)", _method, "members"), members)
	if _err != nil {
		return
	}
	_macArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "MAC"), mac)
	if _err != nil {
		return
	}
	_modeArg, _err := convertEnumBondModeToXen(fmt.Sprintf("%s(%s)", _method, "mode"), mode)
	if _err != nil {
		return
	}
	_propertiesArg, _err := convertStringToStringMapToXen(fmt.Sprintf("%s(%s)", _method, "properties"), properties)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _networkArg, _membersArg, _macArg, _modeArg, _propertiesArg)
	if _err != nil {
		return
	}
	_retval, _err = convertBondRefToGo(_method + " -> ", _result.Value)
	return
}

func (_class BondClass) RemoveFromOtherConfig(sessionID SessionRef, self BondRef, key string) (_err error) {
	_method := "Bond.remove_from_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertBondRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class BondClass) AddToOtherConfig(sessionID SessionRef, self BondRef, key string, value string) (_err error) {
	_method := "Bond.add_to_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertBondRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class BondClass) SetOtherConfig(sessionID SessionRef, self BondRef, value map[string]string) (_err error) {
	_method := "Bond.set_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertBondRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class BondClass) GetLinksUp(sessionID SessionRef, self BondRef) (_retval int, _err error) {
	_method := "Bond.get_links_up"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertBondRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class BondClass) GetProperties(sessionID SessionRef, self BondRef) (_retval map[string]string, _err error) {
	_method := "Bond.get_properties"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertBondRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class BondClass) GetMode(sessionID SessionRef, self BondRef) (_retval BondMode, _err error) {
	_method := "Bond.get_mode"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertBondRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertEnumBondModeToGo(_method + " -> ", _result.Value)
	return
}

func (_class BondClass) GetPrimarySlave(sessionID SessionRef, self BondRef) (_retval PIFRef, _err error) {
	_method := "Bond.get_primary_slave"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertBondRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class BondClass) GetOtherConfig(sessionID SessionRef, self BondRef) (_retval map[string]string, _err error) {
	_method := "Bond.get_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertBondRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class BondClass) GetSlaves(sessionID SessionRef, self BondRef) (_retval []PIFRef, _err error) {
	_method := "Bond.get_slaves"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertBondRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertPIFRefSetToGo(_method + " -> ", _result.Value)
	return
}

func (_class BondClass) GetMaster(sessionID SessionRef, self BondRef) (_retval PIFRef, _err error) {
	_method := "Bond.get_master"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertBondRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class BondClass) GetUUID(sessionID SessionRef, self BondRef) (_retval string, _err error) {
	_method := "Bond.get_uuid"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertBondRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class BondClass) GetByUUID(sessionID SessionRef, uuid string) (_retval BondRef, _err error) {
	_method := "Bond.get_by_uuid"
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
	_retval, _err = convertBondRefToGo(_method + " -> ", _result.Value)
	return
}

func (_class BondClass) GetRecord(sessionID SessionRef, self BondRef) (_retval BondRecord, _err error) {
	_method := "Bond.get_record"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertBondRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertBondRecordToGo(_method + " -> ", _result.Value)
	return
}
