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

type ObserverRecord struct {
	// Unique identifier/object reference
	UUID string
	// a human-readable name
	NameLabel string
	// a notes field containing human-readable description
	NameDescription string
	// The list of hosts the observer is active on. An empty list means all hosts
	Hosts []HostRef
	// Attributes that observer will add to the data they produce
	Attributes map[string]string
	// The list of endpoints where data is exported to. Each endpoint is a URL or the string 'bugtool' refering to the internal logs
	Endpoints []string
	// The list of xenserver components the observer will broadcast. An empty list means all components
	Components []string
	// This denotes if the observer is enabled. true if it is enabled and false if it is disabled
	Enabled bool
}

type ObserverRef string

// Describes a observer which will control observability activity in the Toolstack
type ObserverClass struct {
	client *Client
}

// GetAllRecords Return a map of Observer references to Observer records for all Observers known to the system.
func (_class ObserverClass) GetAllRecords(sessionID SessionRef) (_retval map[ObserverRef]ObserverRecord, _err error) {
	_method := "Observer.get_all_records"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg)
	if _err != nil {
		return
	}
	_retval, _err = convertObserverRefToObserverRecordMapToGo(_method + " -> ", _result.Value)
	return
}

// GetAll Return a list of all the Observers known to the system.
func (_class ObserverClass) GetAll(sessionID SessionRef) (_retval []ObserverRef, _err error) {
	_method := "Observer.get_all"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg)
	if _err != nil {
		return
	}
	_retval, _err = convertObserverRefSetToGo(_method + " -> ", _result.Value)
	return
}

// SetComponents Set the components on which the observer will broadcast to. i.e. xapi, xenopsd, networkd, etc
func (_class ObserverClass) SetComponents(sessionID SessionRef, self ObserverRef, value []string) (_err error) {
	_method := "Observer.set_components"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertObserverRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertStringSetToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}

// SetEndpoints Set the file/HTTP endpoints the observer sends data to
func (_class ObserverClass) SetEndpoints(sessionID SessionRef, self ObserverRef, value []string) (_err error) {
	_method := "Observer.set_endpoints"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertObserverRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertStringSetToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}

// SetAttributes Set the attributes of an observer. These are used to emit metadata by the observer
func (_class ObserverClass) SetAttributes(sessionID SessionRef, self ObserverRef, value map[string]string) (_err error) {
	_method := "Observer.set_attributes"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertObserverRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// SetEnabled Enable / disable this observer which will stop the observer from producing observability information
func (_class ObserverClass) SetEnabled(sessionID SessionRef, self ObserverRef, value bool) (_err error) {
	_method := "Observer.set_enabled"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertObserverRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertBoolToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}

// SetHosts Sets the hosts that the observer is to be registered on
func (_class ObserverClass) SetHosts(sessionID SessionRef, self ObserverRef, value []HostRef) (_err error) {
	_method := "Observer.set_hosts"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertObserverRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertHostRefSetToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}

// SetNameDescription Set the name/description field of the given Observer.
func (_class ObserverClass) SetNameDescription(sessionID SessionRef, self ObserverRef, value string) (_err error) {
	_method := "Observer.set_name_description"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertObserverRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}

// SetNameLabel Set the name/label field of the given Observer.
func (_class ObserverClass) SetNameLabel(sessionID SessionRef, self ObserverRef, value string) (_err error) {
	_method := "Observer.set_name_label"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertObserverRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}

// GetEnabled Get the enabled field of the given Observer.
func (_class ObserverClass) GetEnabled(sessionID SessionRef, self ObserverRef) (_retval bool, _err error) {
	_method := "Observer.get_enabled"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertObserverRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// GetComponents Get the components field of the given Observer.
func (_class ObserverClass) GetComponents(sessionID SessionRef, self ObserverRef) (_retval []string, _err error) {
	_method := "Observer.get_components"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertObserverRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertStringSetToGo(_method + " -> ", _result.Value)
	return
}

// GetEndpoints Get the endpoints field of the given Observer.
func (_class ObserverClass) GetEndpoints(sessionID SessionRef, self ObserverRef) (_retval []string, _err error) {
	_method := "Observer.get_endpoints"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertObserverRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertStringSetToGo(_method + " -> ", _result.Value)
	return
}

// GetAttributes Get the attributes field of the given Observer.
func (_class ObserverClass) GetAttributes(sessionID SessionRef, self ObserverRef) (_retval map[string]string, _err error) {
	_method := "Observer.get_attributes"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertObserverRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// GetHosts Get the hosts field of the given Observer.
func (_class ObserverClass) GetHosts(sessionID SessionRef, self ObserverRef) (_retval []HostRef, _err error) {
	_method := "Observer.get_hosts"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertObserverRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertHostRefSetToGo(_method + " -> ", _result.Value)
	return
}

// GetNameDescription Get the name/description field of the given Observer.
func (_class ObserverClass) GetNameDescription(sessionID SessionRef, self ObserverRef) (_retval string, _err error) {
	_method := "Observer.get_name_description"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertObserverRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// GetNameLabel Get the name/label field of the given Observer.
func (_class ObserverClass) GetNameLabel(sessionID SessionRef, self ObserverRef) (_retval string, _err error) {
	_method := "Observer.get_name_label"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertObserverRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// GetUUID Get the uuid field of the given Observer.
func (_class ObserverClass) GetUUID(sessionID SessionRef, self ObserverRef) (_retval string, _err error) {
	_method := "Observer.get_uuid"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertObserverRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// GetByNameLabel Get all the Observer instances with the given label.
func (_class ObserverClass) GetByNameLabel(sessionID SessionRef, label string) (_retval []ObserverRef, _err error) {
	_method := "Observer.get_by_name_label"
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
	_retval, _err = convertObserverRefSetToGo(_method + " -> ", _result.Value)
	return
}

// Destroy Destroy the specified Observer instance.
func (_class ObserverClass) Destroy(sessionID SessionRef, self ObserverRef) (_err error) {
	_method := "Observer.destroy"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertObserverRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg)
	return
}

// Create Create a new Observer instance, and return its handle. The constructor args are: name_label, name_description, hosts, attributes, endpoints, components, enabled (* = non-optional).
func (_class ObserverClass) Create(sessionID SessionRef, args ObserverRecord) (_retval ObserverRef, _err error) {
	_method := "Observer.create"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_argsArg, _err := convertObserverRecordToXen(fmt.Sprintf("%s(%s)", _method, "args"), args)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _argsArg)
	if _err != nil {
		return
	}
	_retval, _err = convertObserverRefToGo(_method + " -> ", _result.Value)
	return
}

// GetByUUID Get a reference to the Observer instance with the specified UUID.
func (_class ObserverClass) GetByUUID(sessionID SessionRef, uuid string) (_retval ObserverRef, _err error) {
	_method := "Observer.get_by_uuid"
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
	_retval, _err = convertObserverRefToGo(_method + " -> ", _result.Value)
	return
}

// GetRecord Get a record containing the current state of the given Observer.
func (_class ObserverClass) GetRecord(sessionID SessionRef, self ObserverRef) (_retval ObserverRecord, _err error) {
	_method := "Observer.get_record"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertObserverRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertObserverRecordToGo(_method + " -> ", _result.Value)
	return
}
