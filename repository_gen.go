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

type RepositoryRecord struct {
	// Unique identifier/object reference
	UUID string
	// a human-readable name
	NameLabel string
	// a notes field containing human-readable description
	NameDescription string
	// Base URL of binary packages in this repository
	BinaryURL string
	// Base URL of source packages in this repository
	SourceURL string
	// True if updateinfo.xml in this repository needs to be parsed
	Update bool
	// SHA256 checksum of latest updateinfo.xml.gz in this repository if its 'update' is true
	Hash string
	// True if all hosts in pool is up to date with this repository
	UpToDate bool
	// The file name of the GPG public key of this repository
	GpgkeyPath string
}

type RepositoryRef string

// Repository for updates
type RepositoryClass struct {
	client *Client
}

// GetAllRecords Return a map of Repository references to Repository records for all Repositorys known to the system.
func (_class RepositoryClass) GetAllRecords(sessionID SessionRef) (_retval map[RepositoryRef]RepositoryRecord, _err error) {
	_method := "Repository.get_all_records"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg)
	if _err != nil {
		return
	}
	_retval, _err = convertRepositoryRefToRepositoryRecordMapToGo(_method + " -> ", _result.Value)
	return
}

// GetAll Return a list of all the Repositorys known to the system.
func (_class RepositoryClass) GetAll(sessionID SessionRef) (_retval []RepositoryRef, _err error) {
	_method := "Repository.get_all"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg)
	if _err != nil {
		return
	}
	_retval, _err = convertRepositoryRefSetToGo(_method + " -> ", _result.Value)
	return
}

// SetGpgkeyPath Set the file name of the GPG public key of the repository
func (_class RepositoryClass) SetGpgkeyPath(sessionID SessionRef, self RepositoryRef, value string) (_err error) {
	_method := "Repository.set_gpgkey_path"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertRepositoryRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// Forget Remove the repository record from the database
func (_class RepositoryClass) Forget(sessionID SessionRef, self RepositoryRef) (_err error) {
	_method := "Repository.forget"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertRepositoryRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg)
	return
}

// Introduce Add the configuration for a new repository
func (_class RepositoryClass) Introduce(sessionID SessionRef, nameLabel string, nameDescription string, binaryURL string, sourceURL string, update bool, gpgkeyPath string) (_retval RepositoryRef, _err error) {
	_method := "Repository.introduce"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_nameLabelArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "name_label"), nameLabel)
	if _err != nil {
		return
	}
	_nameDescriptionArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "name_description"), nameDescription)
	if _err != nil {
		return
	}
	_binaryURLArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "binary_url"), binaryURL)
	if _err != nil {
		return
	}
	_sourceURLArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "source_url"), sourceURL)
	if _err != nil {
		return
	}
	_updateArg, _err := convertBoolToXen(fmt.Sprintf("%s(%s)", _method, "update"), update)
	if _err != nil {
		return
	}
	_gpgkeyPathArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "gpgkey_path"), gpgkeyPath)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _nameLabelArg, _nameDescriptionArg, _binaryURLArg, _sourceURLArg, _updateArg, _gpgkeyPathArg)
	if _err != nil {
		return
	}
	_retval, _err = convertRepositoryRefToGo(_method + " -> ", _result.Value)
	return
}

// SetNameDescription Set the name/description field of the given Repository.
func (_class RepositoryClass) SetNameDescription(sessionID SessionRef, self RepositoryRef, value string) (_err error) {
	_method := "Repository.set_name_description"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertRepositoryRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// SetNameLabel Set the name/label field of the given Repository.
func (_class RepositoryClass) SetNameLabel(sessionID SessionRef, self RepositoryRef, value string) (_err error) {
	_method := "Repository.set_name_label"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertRepositoryRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// GetGpgkeyPath Get the gpgkey_path field of the given Repository.
func (_class RepositoryClass) GetGpgkeyPath(sessionID SessionRef, self RepositoryRef) (_retval string, _err error) {
	_method := "Repository.get_gpgkey_path"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertRepositoryRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// GetUpToDate Get the up_to_date field of the given Repository.
func (_class RepositoryClass) GetUpToDate(sessionID SessionRef, self RepositoryRef) (_retval bool, _err error) {
	_method := "Repository.get_up_to_date"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertRepositoryRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// GetHash Get the hash field of the given Repository.
func (_class RepositoryClass) GetHash(sessionID SessionRef, self RepositoryRef) (_retval string, _err error) {
	_method := "Repository.get_hash"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertRepositoryRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// GetUpdate Get the update field of the given Repository.
func (_class RepositoryClass) GetUpdate(sessionID SessionRef, self RepositoryRef) (_retval bool, _err error) {
	_method := "Repository.get_update"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertRepositoryRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// GetSourceURL Get the source_url field of the given Repository.
func (_class RepositoryClass) GetSourceURL(sessionID SessionRef, self RepositoryRef) (_retval string, _err error) {
	_method := "Repository.get_source_url"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertRepositoryRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// GetBinaryURL Get the binary_url field of the given Repository.
func (_class RepositoryClass) GetBinaryURL(sessionID SessionRef, self RepositoryRef) (_retval string, _err error) {
	_method := "Repository.get_binary_url"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertRepositoryRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// GetNameDescription Get the name/description field of the given Repository.
func (_class RepositoryClass) GetNameDescription(sessionID SessionRef, self RepositoryRef) (_retval string, _err error) {
	_method := "Repository.get_name_description"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertRepositoryRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// GetNameLabel Get the name/label field of the given Repository.
func (_class RepositoryClass) GetNameLabel(sessionID SessionRef, self RepositoryRef) (_retval string, _err error) {
	_method := "Repository.get_name_label"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertRepositoryRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// GetUUID Get the uuid field of the given Repository.
func (_class RepositoryClass) GetUUID(sessionID SessionRef, self RepositoryRef) (_retval string, _err error) {
	_method := "Repository.get_uuid"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertRepositoryRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// GetByNameLabel Get all the Repository instances with the given label.
func (_class RepositoryClass) GetByNameLabel(sessionID SessionRef, label string) (_retval []RepositoryRef, _err error) {
	_method := "Repository.get_by_name_label"
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
	_retval, _err = convertRepositoryRefSetToGo(_method + " -> ", _result.Value)
	return
}

// GetByUUID Get a reference to the Repository instance with the specified UUID.
func (_class RepositoryClass) GetByUUID(sessionID SessionRef, uuid string) (_retval RepositoryRef, _err error) {
	_method := "Repository.get_by_uuid"
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
	_retval, _err = convertRepositoryRefToGo(_method + " -> ", _result.Value)
	return
}

// GetRecord Get a record containing the current state of the given Repository.
func (_class RepositoryClass) GetRecord(sessionID SessionRef, self RepositoryRef) (_retval RepositoryRecord, _err error) {
	_method := "Repository.get_record"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertRepositoryRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertRepositoryRecordToGo(_method + " -> ", _result.Value)
	return
}
