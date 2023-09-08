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

type CertificateType string

const (
	// Certificate that is trusted by the whole pool
	CertificateTypeCa CertificateType = "ca"
	// Certificate that identifies a single host to entities outside the pool
	CertificateTypeHost CertificateType = "host"
	// Certificate that identifies a single host to other pool members
	CertificateTypeHostInternal CertificateType = "host_internal"
)

type CertificateRecord struct {
	// Unique identifier/object reference
	UUID string
	// The name of the certificate, only present on certificates of type 'ca'
	Name string
	// The type of the certificate, either 'ca', 'host' or 'host_internal'
	Type CertificateType
	// The host where the certificate is installed
	Host HostRef
	// Date after which the certificate is valid
	NotBefore time.Time
	// Date before which the certificate is valid
	NotAfter time.Time
	// The certificate's fingerprint / hash
	Fingerprint string
}

type CertificateRef string

// Description
type CertificateClass struct {
	client *Client
}

// GetAllRecords Return a map of Certificate references to Certificate records for all Certificates known to the system.
func (_class CertificateClass) GetAllRecords(sessionID SessionRef) (_retval map[CertificateRef]CertificateRecord, _err error) {
	_method := "Certificate.get_all_records"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg)
	if _err != nil {
		return
	}
	_retval, _err = convertCertificateRefToCertificateRecordMapToGo(_method + " -> ", _result.Value)
	return
}

// GetAll Return a list of all the Certificates known to the system.
func (_class CertificateClass) GetAll(sessionID SessionRef) (_retval []CertificateRef, _err error) {
	_method := "Certificate.get_all"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg)
	if _err != nil {
		return
	}
	_retval, _err = convertCertificateRefSetToGo(_method + " -> ", _result.Value)
	return
}

// GetFingerprint Get the fingerprint field of the given Certificate.
func (_class CertificateClass) GetFingerprint(sessionID SessionRef, self CertificateRef) (_retval string, _err error) {
	_method := "Certificate.get_fingerprint"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertCertificateRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// GetNotAfter Get the not_after field of the given Certificate.
func (_class CertificateClass) GetNotAfter(sessionID SessionRef, self CertificateRef) (_retval time.Time, _err error) {
	_method := "Certificate.get_not_after"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertCertificateRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// GetNotBefore Get the not_before field of the given Certificate.
func (_class CertificateClass) GetNotBefore(sessionID SessionRef, self CertificateRef) (_retval time.Time, _err error) {
	_method := "Certificate.get_not_before"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertCertificateRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// GetHost Get the host field of the given Certificate.
func (_class CertificateClass) GetHost(sessionID SessionRef, self CertificateRef) (_retval HostRef, _err error) {
	_method := "Certificate.get_host"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertCertificateRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// GetType Get the type field of the given Certificate.
func (_class CertificateClass) GetType(sessionID SessionRef, self CertificateRef) (_retval CertificateType, _err error) {
	_method := "Certificate.get_type"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertCertificateRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertEnumCertificateTypeToGo(_method + " -> ", _result.Value)
	return
}

// GetName Get the name field of the given Certificate.
func (_class CertificateClass) GetName(sessionID SessionRef, self CertificateRef) (_retval string, _err error) {
	_method := "Certificate.get_name"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertCertificateRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// GetUUID Get the uuid field of the given Certificate.
func (_class CertificateClass) GetUUID(sessionID SessionRef, self CertificateRef) (_retval string, _err error) {
	_method := "Certificate.get_uuid"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertCertificateRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// GetByUUID Get a reference to the Certificate instance with the specified UUID.
func (_class CertificateClass) GetByUUID(sessionID SessionRef, uuid string) (_retval CertificateRef, _err error) {
	_method := "Certificate.get_by_uuid"
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
	_retval, _err = convertCertificateRefToGo(_method + " -> ", _result.Value)
	return
}

// GetRecord Get a record containing the current state of the given Certificate.
func (_class CertificateClass) GetRecord(sessionID SessionRef, self CertificateRef) (_retval CertificateRecord, _err error) {
	_method := "Certificate.get_record"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertCertificateRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertCertificateRecordToGo(_method + " -> ", _result.Value)
	return
}
