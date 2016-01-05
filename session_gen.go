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

type SessionRecord struct {
	UUID string
	ThisHost HostRef
	ThisUser UserRef
	LastActive time.Time
	Pool bool
	OtherConfig map[string]string
	IsLocalSuperuser bool
	Subject SubjectRef
	ValidationTime time.Time
	AuthUserSid string
	AuthUserName string
	RbacPermissions []string
	Tasks []TaskRef
	Parent SessionRef
	Originator string
}

type SessionRef string

type SessionClass struct {
	client *Client
}

func (client *Client) Session() SessionClass {
	return SessionClass{client}
}

func (_class SessionClass) LogoutSubjectIdentifier(sessionID SessionRef, subjectIdentifier string) (_err error) {
	_method := "session.logout_subject_identifier"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_subjectIdentifierArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "subject_identifier"), subjectIdentifier)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _subjectIdentifierArg)
	return
}

func (_class SessionClass) GetAllSubjectIdentifiers(sessionID SessionRef) (_retval []string, _err error) {
	_method := "session.get_all_subject_identifiers"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg)
	if _err != nil {
		return
	}
	_retval, _err = convertStringSetToGo(_method + " -> ", _result.Value)
	return
}

func (_class SessionClass) LocalLogout(sessionID SessionRef) (_err error) {
	_method := "session.local_logout"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg)
	return
}

func (_class SessionClass) CreateFromDbFile(sessionID SessionRef, filename string) (_retval SessionRef, _err error) {
	_method := "session.create_from_db_file"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_filenameArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "filename"), filename)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _filenameArg)
	if _err != nil {
		return
	}
	_retval, _err = convertSessionRefToGo(_method + " -> ", _result.Value)
	return
}

func (_class SessionClass) SlaveLocalLoginWithPassword(uname string, pwd string) (_retval SessionRef, _err error) {
	_method := "session.slave_local_login_with_password"
	_unameArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "uname"), uname)
	if _err != nil {
		return
	}
	_pwdArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "pwd"), pwd)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _unameArg, _pwdArg)
	if _err != nil {
		return
	}
	_retval, _err = convertSessionRefToGo(_method + " -> ", _result.Value)
	return
}

func (_class SessionClass) ChangePassword(sessionID SessionRef, oldPwd string, newPwd string) (_err error) {
	_method := "session.change_password"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_oldPwdArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "old_pwd"), oldPwd)
	if _err != nil {
		return
	}
	_newPwdArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "new_pwd"), newPwd)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _oldPwdArg, _newPwdArg)
	return
}

func (_class SessionClass) Logout(sessionID SessionRef) (_err error) {
	_method := "session.logout"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg)
	return
}

func (_class SessionClass) LoginWithPassword(uname string, pwd string, version string, originator string) (_retval SessionRef, _err error) {
	_method := "session.login_with_password"
	_unameArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "uname"), uname)
	if _err != nil {
		return
	}
	_pwdArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "pwd"), pwd)
	if _err != nil {
		return
	}
	_versionArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "version"), version)
	if _err != nil {
		return
	}
	_originatorArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "originator"), originator)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _unameArg, _pwdArg, _versionArg, _originatorArg)
	if _err != nil {
		return
	}
	_retval, _err = convertSessionRefToGo(_method + " -> ", _result.Value)
	return
}

func (_class SessionClass) RemoveFromOtherConfig(sessionID SessionRef, self SessionRef, key string) (_err error) {
	_method := "session.remove_from_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class SessionClass) AddToOtherConfig(sessionID SessionRef, self SessionRef, key string, value string) (_err error) {
	_method := "session.add_to_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class SessionClass) SetOtherConfig(sessionID SessionRef, self SessionRef, value map[string]string) (_err error) {
	_method := "session.set_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class SessionClass) GetOriginator(sessionID SessionRef, self SessionRef) (_retval string, _err error) {
	_method := "session.get_originator"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class SessionClass) GetParent(sessionID SessionRef, self SessionRef) (_retval SessionRef, _err error) {
	_method := "session.get_parent"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertSessionRefToGo(_method + " -> ", _result.Value)
	return
}

func (_class SessionClass) GetTasks(sessionID SessionRef, self SessionRef) (_retval []TaskRef, _err error) {
	_method := "session.get_tasks"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefSetToGo(_method + " -> ", _result.Value)
	return
}

func (_class SessionClass) GetRbacPermissions(sessionID SessionRef, self SessionRef) (_retval []string, _err error) {
	_method := "session.get_rbac_permissions"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class SessionClass) GetAuthUserName(sessionID SessionRef, self SessionRef) (_retval string, _err error) {
	_method := "session.get_auth_user_name"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class SessionClass) GetAuthUserSid(sessionID SessionRef, self SessionRef) (_retval string, _err error) {
	_method := "session.get_auth_user_sid"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class SessionClass) GetValidationTime(sessionID SessionRef, self SessionRef) (_retval time.Time, _err error) {
	_method := "session.get_validation_time"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class SessionClass) GetSubject(sessionID SessionRef, self SessionRef) (_retval SubjectRef, _err error) {
	_method := "session.get_subject"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertSubjectRefToGo(_method + " -> ", _result.Value)
	return
}

func (_class SessionClass) GetIsLocalSuperuser(sessionID SessionRef, self SessionRef) (_retval bool, _err error) {
	_method := "session.get_is_local_superuser"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class SessionClass) GetOtherConfig(sessionID SessionRef, self SessionRef) (_retval map[string]string, _err error) {
	_method := "session.get_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class SessionClass) GetPool(sessionID SessionRef, self SessionRef) (_retval bool, _err error) {
	_method := "session.get_pool"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class SessionClass) GetLastActive(sessionID SessionRef, self SessionRef) (_retval time.Time, _err error) {
	_method := "session.get_last_active"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class SessionClass) GetThisUser(sessionID SessionRef, self SessionRef) (_retval UserRef, _err error) {
	_method := "session.get_this_user"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertUserRefToGo(_method + " -> ", _result.Value)
	return
}

func (_class SessionClass) GetThisHost(sessionID SessionRef, self SessionRef) (_retval HostRef, _err error) {
	_method := "session.get_this_host"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class SessionClass) GetUUID(sessionID SessionRef, self SessionRef) (_retval string, _err error) {
	_method := "session.get_uuid"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class SessionClass) GetByUUID(sessionID SessionRef, uuid string) (_retval SessionRef, _err error) {
	_method := "session.get_by_uuid"
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
	_retval, _err = convertSessionRefToGo(_method + " -> ", _result.Value)
	return
}

func (_class SessionClass) GetRecord(sessionID SessionRef, self SessionRef) (_retval SessionRecord, _err error) {
	_method := "session.get_record"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertSessionRecordToGo(_method + " -> ", _result.Value)
	return
}
