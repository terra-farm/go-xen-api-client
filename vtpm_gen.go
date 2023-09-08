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

type VtpmOperations string

const (
	// Destroy a VTPM
	VtpmOperationsDestroy VtpmOperations = "destroy"
)

type PersistenceBackend string

const (
	// This VTPM is persisted in XAPI's DB
	PersistenceBackendXapi PersistenceBackend = "xapi"
)

type VTPMRecord struct {
	// Unique identifier/object reference
	UUID string
	// list of the operations allowed in this state. This list is advisory only and the server state may have changed by the time this field is read by a client.
	AllowedOperations []VtpmOperations
	// links each of the running tasks using this object (by reference) to a current_operation enum which describes the nature of the task.
	CurrentOperations map[string]VtpmOperations
	// The virtual machine the TPM is attached to
	VM VMRef
	// The domain where the backend is located (unused)
	Backend VMRef
	// The backend where the vTPM is persisted
	PersistenceBackend PersistenceBackend
	// Whether the contents are never copied, satisfying the TPM spec
	IsUnique bool
	// Whether the contents of the VTPM are secured according to the TPM spec
	IsProtected bool
}

type VTPMRef string

// A virtual TPM device
type VTPMClass struct {
	client *Client
}

// GetAllRecords Return a map of VTPM references to VTPM records for all VTPMs known to the system.
func (_class VTPMClass) GetAllRecords(sessionID SessionRef) (_retval map[VTPMRef]VTPMRecord, _err error) {
	_method := "VTPM.get_all_records"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVTPMRefToVTPMRecordMapToGo(_method + " -> ", _result.Value)
	return
}

// GetAll Return a list of all the VTPMs known to the system.
func (_class VTPMClass) GetAll(sessionID SessionRef) (_retval []VTPMRef, _err error) {
	_method := "VTPM.get_all"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVTPMRefSetToGo(_method + " -> ", _result.Value)
	return
}

// Destroy Destroy the specified VTPM instance, along with its state.
func (_class VTPMClass) Destroy(sessionID SessionRef, self VTPMRef) (_err error) {
	_method := "VTPM.destroy"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVTPMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg)
	return
}

// Create Create a new VTPM instance, and return its handle.
func (_class VTPMClass) Create(sessionID SessionRef, vm VMRef, isUnique bool) (_retval VTPMRef, _err error) {
	_method := "VTPM.create"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_vmArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "vM"), vm)
	if _err != nil {
		return
	}
	_isUniqueArg, _err := convertBoolToXen(fmt.Sprintf("%s(%s)", _method, "is_unique"), isUnique)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _vmArg, _isUniqueArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVTPMRefToGo(_method + " -> ", _result.Value)
	return
}

// GetIsProtected Get the is_protected field of the given VTPM.
func (_class VTPMClass) GetIsProtected(sessionID SessionRef, self VTPMRef) (_retval bool, _err error) {
	_method := "VTPM.get_is_protected"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVTPMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// GetIsUnique Get the is_unique field of the given VTPM.
func (_class VTPMClass) GetIsUnique(sessionID SessionRef, self VTPMRef) (_retval bool, _err error) {
	_method := "VTPM.get_is_unique"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVTPMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// GetPersistenceBackend Get the persistence_backend field of the given VTPM.
func (_class VTPMClass) GetPersistenceBackend(sessionID SessionRef, self VTPMRef) (_retval PersistenceBackend, _err error) {
	_method := "VTPM.get_persistence_backend"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVTPMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertEnumPersistenceBackendToGo(_method + " -> ", _result.Value)
	return
}

// GetBackend Get the backend field of the given VTPM.
func (_class VTPMClass) GetBackend(sessionID SessionRef, self VTPMRef) (_retval VMRef, _err error) {
	_method := "VTPM.get_backend"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVTPMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVMRefToGo(_method + " -> ", _result.Value)
	return
}

// GetVM Get the VM field of the given VTPM.
func (_class VTPMClass) GetVM(sessionID SessionRef, self VTPMRef) (_retval VMRef, _err error) {
	_method := "VTPM.get_VM"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVTPMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVMRefToGo(_method + " -> ", _result.Value)
	return
}

// GetCurrentOperations Get the current_operations field of the given VTPM.
func (_class VTPMClass) GetCurrentOperations(sessionID SessionRef, self VTPMRef) (_retval map[string]VtpmOperations, _err error) {
	_method := "VTPM.get_current_operations"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVTPMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertStringToEnumVtpmOperationsMapToGo(_method + " -> ", _result.Value)
	return
}

// GetAllowedOperations Get the allowed_operations field of the given VTPM.
func (_class VTPMClass) GetAllowedOperations(sessionID SessionRef, self VTPMRef) (_retval []VtpmOperations, _err error) {
	_method := "VTPM.get_allowed_operations"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVTPMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertEnumVtpmOperationsSetToGo(_method + " -> ", _result.Value)
	return
}

// GetUUID Get the uuid field of the given VTPM.
func (_class VTPMClass) GetUUID(sessionID SessionRef, self VTPMRef) (_retval string, _err error) {
	_method := "VTPM.get_uuid"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVTPMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// GetByUUID Get a reference to the VTPM instance with the specified UUID.
func (_class VTPMClass) GetByUUID(sessionID SessionRef, uuid string) (_retval VTPMRef, _err error) {
	_method := "VTPM.get_by_uuid"
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
	_retval, _err = convertVTPMRefToGo(_method + " -> ", _result.Value)
	return
}

// GetRecord Get a record containing the current state of the given VTPM.
func (_class VTPMClass) GetRecord(sessionID SessionRef, self VTPMRef) (_retval VTPMRecord, _err error) {
	_method := "VTPM.get_record"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVTPMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVTPMRecordToGo(_method + " -> ", _result.Value)
	return
}
