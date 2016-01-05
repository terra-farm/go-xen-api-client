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

type PoolAllowedOperations string

const (
	PoolAllowedOperationsHaEnable PoolAllowedOperations = "ha_enable"
	PoolAllowedOperationsHaDisable PoolAllowedOperations = "ha_disable"
)

type PoolRecord struct {
	UUID string
	NameLabel string
	NameDescription string
	Master HostRef
	DefaultSR SRRef
	SuspendImageSR SRRef
	CrashDumpSR SRRef
	OtherConfig map[string]string
	HaEnabled bool
	HaConfiguration map[string]string
	HaStatefiles []string
	HaHostFailuresToTolerate int
	HaPlanExistsFor int
	HaAllowOvercommit bool
	HaOvercommitted bool
	Blobs map[string]BlobRef
	Tags []string
	GuiConfig map[string]string
	HealthCheckConfig map[string]string
	WlbURL string
	WlbUsername string
	WlbEnabled bool
	WlbVerifyCert bool
	RedoLogEnabled bool
	RedoLogVdi VDIRef
	VswitchController string
	Restrictions map[string]string
	MetadataVDIs []VDIRef
	HaClusterStack string
	AllowedOperations []PoolAllowedOperations
	CurrentOperations map[string]PoolAllowedOperations
}

type PoolRef string

type PoolClass struct {
	client *Client
}

func (client *Client) Pool() PoolClass {
	return PoolClass{client}
}

func (_class PoolClass) GetAllRecords(sessionID SessionRef) (_retval map[PoolRef]PoolRecord, _err error) {
	_method := "pool.get_all_records"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg)
	if _err != nil {
		return
	}
	_retval, _err = convertPoolRefToPoolRecordMapToGo(_method + " -> ", _result.Value)
	return
}

func (_class PoolClass) GetAll(sessionID SessionRef) (_retval []PoolRef, _err error) {
	_method := "pool.get_all"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg)
	if _err != nil {
		return
	}
	_retval, _err = convertPoolRefSetToGo(_method + " -> ", _result.Value)
	return
}

func (_class PoolClass) DisableSslLegacy(sessionID SessionRef, self PoolRef) (_err error) {
	_method := "pool.disable_ssl_legacy"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg)
	return
}

func (_class PoolClass) EnableSslLegacy(sessionID SessionRef, self PoolRef) (_err error) {
	_method := "pool.enable_ssl_legacy"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg)
	return
}

func (_class PoolClass) ApplyEdition(sessionID SessionRef, self PoolRef, edition string) (_err error) {
	_method := "pool.apply_edition"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_editionArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "edition"), edition)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _editionArg)
	return
}

func (_class PoolClass) GetLicenseState(sessionID SessionRef, self PoolRef) (_retval map[string]string, _err error) {
	_method := "pool.get_license_state"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class PoolClass) DisableLocalStorageCaching(sessionID SessionRef, self PoolRef) (_err error) {
	_method := "pool.disable_local_storage_caching"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg)
	return
}

func (_class PoolClass) EnableLocalStorageCaching(sessionID SessionRef, self PoolRef) (_err error) {
	_method := "pool.enable_local_storage_caching"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg)
	return
}

func (_class PoolClass) TestArchiveTarget(sessionID SessionRef, self PoolRef, config map[string]string) (_retval string, _err error) {
	_method := "pool.test_archive_target"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_configArg, _err := convertStringToStringMapToXen(fmt.Sprintf("%s(%s)", _method, "config"), config)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg, _configArg)
	if _err != nil {
		return
	}
	_retval, _err = convertStringToGo(_method + " -> ", _result.Value)
	return
}

func (_class PoolClass) SetVswitchController(sessionID SessionRef, address string) (_err error) {
	_method := "pool.set_vswitch_controller"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_addressArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "address"), address)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _addressArg)
	return
}

func (_class PoolClass) DisableRedoLog(sessionID SessionRef) (_err error) {
	_method := "pool.disable_redo_log"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg)
	return
}

func (_class PoolClass) EnableRedoLog(sessionID SessionRef, sr SRRef) (_err error) {
	_method := "pool.enable_redo_log"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_srArg, _err := convertSRRefToXen(fmt.Sprintf("%s(%s)", _method, "sr"), sr)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _srArg)
	return
}

func (_class PoolClass) CertificateSync(sessionID SessionRef) (_err error) {
	_method := "pool.certificate_sync"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg)
	return
}

func (_class PoolClass) CrlList(sessionID SessionRef) (_retval []string, _err error) {
	_method := "pool.crl_list"
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

func (_class PoolClass) CrlUninstall(sessionID SessionRef, name string) (_err error) {
	_method := "pool.crl_uninstall"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_nameArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "name"), name)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _nameArg)
	return
}

func (_class PoolClass) CrlInstall(sessionID SessionRef, name string, cert string) (_err error) {
	_method := "pool.crl_install"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_nameArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "name"), name)
	if _err != nil {
		return
	}
	_certArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "cert"), cert)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _nameArg, _certArg)
	return
}

func (_class PoolClass) CertificateList(sessionID SessionRef) (_retval []string, _err error) {
	_method := "pool.certificate_list"
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

func (_class PoolClass) CertificateUninstall(sessionID SessionRef, name string) (_err error) {
	_method := "pool.certificate_uninstall"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_nameArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "name"), name)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _nameArg)
	return
}

func (_class PoolClass) CertificateInstall(sessionID SessionRef, name string, cert string) (_err error) {
	_method := "pool.certificate_install"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_nameArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "name"), name)
	if _err != nil {
		return
	}
	_certArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "cert"), cert)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _nameArg, _certArg)
	return
}

func (_class PoolClass) SendTestPost(sessionID SessionRef, host string, port int, body string) (_retval string, _err error) {
	_method := "pool.send_test_post"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_hostArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "host"), host)
	if _err != nil {
		return
	}
	_portArg, _err := convertIntToXen(fmt.Sprintf("%s(%s)", _method, "port"), port)
	if _err != nil {
		return
	}
	_bodyArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "body"), body)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _hostArg, _portArg, _bodyArg)
	if _err != nil {
		return
	}
	_retval, _err = convertStringToGo(_method + " -> ", _result.Value)
	return
}

func (_class PoolClass) RetrieveWlbRecommendations(sessionID SessionRef) (_retval map[VMRef][]string, _err error) {
	_method := "pool.retrieve_wlb_recommendations"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVMRefToStringSetMapToGo(_method + " -> ", _result.Value)
	return
}

func (_class PoolClass) RetrieveWlbConfiguration(sessionID SessionRef) (_retval map[string]string, _err error) {
	_method := "pool.retrieve_wlb_configuration"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg)
	if _err != nil {
		return
	}
	_retval, _err = convertStringToStringMapToGo(_method + " -> ", _result.Value)
	return
}

func (_class PoolClass) SendWlbConfiguration(sessionID SessionRef, config map[string]string) (_err error) {
	_method := "pool.send_wlb_configuration"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_configArg, _err := convertStringToStringMapToXen(fmt.Sprintf("%s(%s)", _method, "config"), config)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _configArg)
	return
}

func (_class PoolClass) DeconfigureWlb(sessionID SessionRef) (_err error) {
	_method := "pool.deconfigure_wlb"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg)
	return
}

func (_class PoolClass) InitializeWlb(sessionID SessionRef, wlbURL string, wlbUsername string, wlbPassword string, xenserverUsername string, xenserverPassword string) (_err error) {
	_method := "pool.initialize_wlb"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_wlbURLArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "wlb_url"), wlbURL)
	if _err != nil {
		return
	}
	_wlbUsernameArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "wlb_username"), wlbUsername)
	if _err != nil {
		return
	}
	_wlbPasswordArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "wlb_password"), wlbPassword)
	if _err != nil {
		return
	}
	_xenserverUsernameArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "xenserver_username"), xenserverUsername)
	if _err != nil {
		return
	}
	_xenserverPasswordArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "xenserver_password"), xenserverPassword)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _wlbURLArg, _wlbUsernameArg, _wlbPasswordArg, _xenserverUsernameArg, _xenserverPasswordArg)
	return
}

func (_class PoolClass) DetectNonhomogeneousExternalAuth(sessionID SessionRef, pool PoolRef) (_err error) {
	_method := "pool.detect_nonhomogeneous_external_auth"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_poolArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "pool"), pool)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _poolArg)
	return
}

func (_class PoolClass) DisableExternalAuth(sessionID SessionRef, pool PoolRef, config map[string]string) (_err error) {
	_method := "pool.disable_external_auth"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_poolArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "pool"), pool)
	if _err != nil {
		return
	}
	_configArg, _err := convertStringToStringMapToXen(fmt.Sprintf("%s(%s)", _method, "config"), config)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _poolArg, _configArg)
	return
}

func (_class PoolClass) EnableExternalAuth(sessionID SessionRef, pool PoolRef, config map[string]string, serviceName string, authType string) (_err error) {
	_method := "pool.enable_external_auth"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_poolArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "pool"), pool)
	if _err != nil {
		return
	}
	_configArg, _err := convertStringToStringMapToXen(fmt.Sprintf("%s(%s)", _method, "config"), config)
	if _err != nil {
		return
	}
	_serviceNameArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "service_name"), serviceName)
	if _err != nil {
		return
	}
	_authTypeArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "auth_type"), authType)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _poolArg, _configArg, _serviceNameArg, _authTypeArg)
	return
}

func (_class PoolClass) CreateNewBlob(sessionID SessionRef, pool PoolRef, name string, mimeType string, public bool) (_retval BlobRef, _err error) {
	_method := "pool.create_new_blob"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_poolArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "pool"), pool)
	if _err != nil {
		return
	}
	_nameArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "name"), name)
	if _err != nil {
		return
	}
	_mimeTypeArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "mime_type"), mimeType)
	if _err != nil {
		return
	}
	_publicArg, _err := convertBoolToXen(fmt.Sprintf("%s(%s)", _method, "public"), public)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _poolArg, _nameArg, _mimeTypeArg, _publicArg)
	if _err != nil {
		return
	}
	_retval, _err = convertBlobRefToGo(_method + " -> ", _result.Value)
	return
}

func (_class PoolClass) SetHaHostFailuresToTolerate(sessionID SessionRef, self PoolRef, value int) (_err error) {
	_method := "pool.set_ha_host_failures_to_tolerate"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertIntToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}

func (_class PoolClass) HaComputeVMFailoverPlan(sessionID SessionRef, failedHosts []HostRef, failedVms []VMRef) (_retval map[VMRef]map[string]string, _err error) {
	_method := "pool.ha_compute_vm_failover_plan"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_failedHostsArg, _err := convertHostRefSetToXen(fmt.Sprintf("%s(%s)", _method, "failed_hosts"), failedHosts)
	if _err != nil {
		return
	}
	_failedVmsArg, _err := convertVMRefSetToXen(fmt.Sprintf("%s(%s)", _method, "failed_vms"), failedVms)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _failedHostsArg, _failedVmsArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVMRefToStringToStringMapMapToGo(_method + " -> ", _result.Value)
	return
}

func (_class PoolClass) HaComputeHypotheticalMaxHostFailuresToTolerate(sessionID SessionRef, configuration map[VMRef]string) (_retval int, _err error) {
	_method := "pool.ha_compute_hypothetical_max_host_failures_to_tolerate"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_configurationArg, _err := convertVMRefToStringMapToXen(fmt.Sprintf("%s(%s)", _method, "configuration"), configuration)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _configurationArg)
	if _err != nil {
		return
	}
	_retval, _err = convertIntToGo(_method + " -> ", _result.Value)
	return
}

func (_class PoolClass) HaComputeMaxHostFailuresToTolerate(sessionID SessionRef) (_retval int, _err error) {
	_method := "pool.ha_compute_max_host_failures_to_tolerate"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg)
	if _err != nil {
		return
	}
	_retval, _err = convertIntToGo(_method + " -> ", _result.Value)
	return
}

func (_class PoolClass) HaFailoverPlanExists(sessionID SessionRef, n int) (_retval bool, _err error) {
	_method := "pool.ha_failover_plan_exists"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_nArg, _err := convertIntToXen(fmt.Sprintf("%s(%s)", _method, "n"), n)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _nArg)
	if _err != nil {
		return
	}
	_retval, _err = convertBoolToGo(_method + " -> ", _result.Value)
	return
}

func (_class PoolClass) HaPreventRestartsFor(sessionID SessionRef, seconds int) (_err error) {
	_method := "pool.ha_prevent_restarts_for"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_secondsArg, _err := convertIntToXen(fmt.Sprintf("%s(%s)", _method, "seconds"), seconds)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _secondsArg)
	return
}

func (_class PoolClass) DesignateNewMaster(sessionID SessionRef, host HostRef) (_err error) {
	_method := "pool.designate_new_master"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_hostArg, _err := convertHostRefToXen(fmt.Sprintf("%s(%s)", _method, "host"), host)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _hostArg)
	return
}

func (_class PoolClass) SyncDatabase(sessionID SessionRef) (_err error) {
	_method := "pool.sync_database"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg)
	return
}

func (_class PoolClass) DisableHa(sessionID SessionRef) (_err error) {
	_method := "pool.disable_ha"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg)
	return
}

func (_class PoolClass) EnableHa(sessionID SessionRef, heartbeatSrs []SRRef, configuration map[string]string) (_err error) {
	_method := "pool.enable_ha"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_heartbeatSrsArg, _err := convertSRRefSetToXen(fmt.Sprintf("%s(%s)", _method, "heartbeat_srs"), heartbeatSrs)
	if _err != nil {
		return
	}
	_configurationArg, _err := convertStringToStringMapToXen(fmt.Sprintf("%s(%s)", _method, "configuration"), configuration)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _heartbeatSrsArg, _configurationArg)
	return
}

func (_class PoolClass) CreateVLANFromPIF(sessionID SessionRef, pif PIFRef, network NetworkRef, vlan int) (_retval []PIFRef, _err error) {
	_method := "pool.create_VLAN_from_PIF"
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
	_vlanArg, _err := convertIntToXen(fmt.Sprintf("%s(%s)", _method, "VLAN"), vlan)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _pifArg, _networkArg, _vlanArg)
	if _err != nil {
		return
	}
	_retval, _err = convertPIFRefSetToGo(_method + " -> ", _result.Value)
	return
}

func (_class PoolClass) CreateVLAN(sessionID SessionRef, device string, network NetworkRef, vlan int) (_retval []PIFRef, _err error) {
	_method := "pool.create_VLAN"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_deviceArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "device"), device)
	if _err != nil {
		return
	}
	_networkArg, _err := convertNetworkRefToXen(fmt.Sprintf("%s(%s)", _method, "network"), network)
	if _err != nil {
		return
	}
	_vlanArg, _err := convertIntToXen(fmt.Sprintf("%s(%s)", _method, "VLAN"), vlan)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _deviceArg, _networkArg, _vlanArg)
	if _err != nil {
		return
	}
	_retval, _err = convertPIFRefSetToGo(_method + " -> ", _result.Value)
	return
}

func (_class PoolClass) RecoverSlaves(sessionID SessionRef) (_retval []HostRef, _err error) {
	_method := "pool.recover_slaves"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg)
	if _err != nil {
		return
	}
	_retval, _err = convertHostRefSetToGo(_method + " -> ", _result.Value)
	return
}

func (_class PoolClass) EmergencyResetMaster(sessionID SessionRef, masterAddress string) (_err error) {
	_method := "pool.emergency_reset_master"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_masterAddressArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "master_address"), masterAddress)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _masterAddressArg)
	return
}

func (_class PoolClass) EmergencyTransitionToMaster(sessionID SessionRef) (_err error) {
	_method := "pool.emergency_transition_to_master"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg)
	return
}

func (_class PoolClass) Eject(sessionID SessionRef, host HostRef) (_err error) {
	_method := "pool.eject"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_hostArg, _err := convertHostRefToXen(fmt.Sprintf("%s(%s)", _method, "host"), host)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _hostArg)
	return
}

func (_class PoolClass) JoinForce(sessionID SessionRef, masterAddress string, masterUsername string, masterPassword string) (_err error) {
	_method := "pool.join_force"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_masterAddressArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "master_address"), masterAddress)
	if _err != nil {
		return
	}
	_masterUsernameArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "master_username"), masterUsername)
	if _err != nil {
		return
	}
	_masterPasswordArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "master_password"), masterPassword)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _masterAddressArg, _masterUsernameArg, _masterPasswordArg)
	return
}

func (_class PoolClass) Join(sessionID SessionRef, masterAddress string, masterUsername string, masterPassword string) (_err error) {
	_method := "pool.join"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_masterAddressArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "master_address"), masterAddress)
	if _err != nil {
		return
	}
	_masterUsernameArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "master_username"), masterUsername)
	if _err != nil {
		return
	}
	_masterPasswordArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "master_password"), masterPassword)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _masterAddressArg, _masterUsernameArg, _masterPasswordArg)
	return
}

func (_class PoolClass) SetWlbVerifyCert(sessionID SessionRef, self PoolRef, value bool) (_err error) {
	_method := "pool.set_wlb_verify_cert"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class PoolClass) SetWlbEnabled(sessionID SessionRef, self PoolRef, value bool) (_err error) {
	_method := "pool.set_wlb_enabled"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class PoolClass) RemoveFromHealthCheckConfig(sessionID SessionRef, self PoolRef, key string) (_err error) {
	_method := "pool.remove_from_health_check_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class PoolClass) AddToHealthCheckConfig(sessionID SessionRef, self PoolRef, key string, value string) (_err error) {
	_method := "pool.add_to_health_check_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class PoolClass) SetHealthCheckConfig(sessionID SessionRef, self PoolRef, value map[string]string) (_err error) {
	_method := "pool.set_health_check_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class PoolClass) RemoveFromGuiConfig(sessionID SessionRef, self PoolRef, key string) (_err error) {
	_method := "pool.remove_from_gui_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class PoolClass) AddToGuiConfig(sessionID SessionRef, self PoolRef, key string, value string) (_err error) {
	_method := "pool.add_to_gui_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class PoolClass) SetGuiConfig(sessionID SessionRef, self PoolRef, value map[string]string) (_err error) {
	_method := "pool.set_gui_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class PoolClass) RemoveTags(sessionID SessionRef, self PoolRef, value string) (_err error) {
	_method := "pool.remove_tags"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class PoolClass) AddTags(sessionID SessionRef, self PoolRef, value string) (_err error) {
	_method := "pool.add_tags"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class PoolClass) SetTags(sessionID SessionRef, self PoolRef, value []string) (_err error) {
	_method := "pool.set_tags"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class PoolClass) SetHaAllowOvercommit(sessionID SessionRef, self PoolRef, value bool) (_err error) {
	_method := "pool.set_ha_allow_overcommit"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class PoolClass) RemoveFromOtherConfig(sessionID SessionRef, self PoolRef, key string) (_err error) {
	_method := "pool.remove_from_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class PoolClass) AddToOtherConfig(sessionID SessionRef, self PoolRef, key string, value string) (_err error) {
	_method := "pool.add_to_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class PoolClass) SetOtherConfig(sessionID SessionRef, self PoolRef, value map[string]string) (_err error) {
	_method := "pool.set_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class PoolClass) SetCrashDumpSR(sessionID SessionRef, self PoolRef, value SRRef) (_err error) {
	_method := "pool.set_crash_dump_SR"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertSRRefToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}

func (_class PoolClass) SetSuspendImageSR(sessionID SessionRef, self PoolRef, value SRRef) (_err error) {
	_method := "pool.set_suspend_image_SR"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertSRRefToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}

func (_class PoolClass) SetDefaultSR(sessionID SessionRef, self PoolRef, value SRRef) (_err error) {
	_method := "pool.set_default_SR"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertSRRefToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}

func (_class PoolClass) SetNameDescription(sessionID SessionRef, self PoolRef, value string) (_err error) {
	_method := "pool.set_name_description"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class PoolClass) SetNameLabel(sessionID SessionRef, self PoolRef, value string) (_err error) {
	_method := "pool.set_name_label"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class PoolClass) GetCurrentOperations(sessionID SessionRef, self PoolRef) (_retval map[string]PoolAllowedOperations, _err error) {
	_method := "pool.get_current_operations"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertStringToEnumPoolAllowedOperationsMapToGo(_method + " -> ", _result.Value)
	return
}

func (_class PoolClass) GetAllowedOperations(sessionID SessionRef, self PoolRef) (_retval []PoolAllowedOperations, _err error) {
	_method := "pool.get_allowed_operations"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertEnumPoolAllowedOperationsSetToGo(_method + " -> ", _result.Value)
	return
}

func (_class PoolClass) GetHaClusterStack(sessionID SessionRef, self PoolRef) (_retval string, _err error) {
	_method := "pool.get_ha_cluster_stack"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class PoolClass) GetMetadataVDIs(sessionID SessionRef, self PoolRef) (_retval []VDIRef, _err error) {
	_method := "pool.get_metadata_VDIs"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVDIRefSetToGo(_method + " -> ", _result.Value)
	return
}

func (_class PoolClass) GetRestrictions(sessionID SessionRef, self PoolRef) (_retval map[string]string, _err error) {
	_method := "pool.get_restrictions"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class PoolClass) GetVswitchController(sessionID SessionRef, self PoolRef) (_retval string, _err error) {
	_method := "pool.get_vswitch_controller"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class PoolClass) GetRedoLogVdi(sessionID SessionRef, self PoolRef) (_retval VDIRef, _err error) {
	_method := "pool.get_redo_log_vdi"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVDIRefToGo(_method + " -> ", _result.Value)
	return
}

func (_class PoolClass) GetRedoLogEnabled(sessionID SessionRef, self PoolRef) (_retval bool, _err error) {
	_method := "pool.get_redo_log_enabled"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class PoolClass) GetWlbVerifyCert(sessionID SessionRef, self PoolRef) (_retval bool, _err error) {
	_method := "pool.get_wlb_verify_cert"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class PoolClass) GetWlbEnabled(sessionID SessionRef, self PoolRef) (_retval bool, _err error) {
	_method := "pool.get_wlb_enabled"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class PoolClass) GetWlbUsername(sessionID SessionRef, self PoolRef) (_retval string, _err error) {
	_method := "pool.get_wlb_username"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class PoolClass) GetWlbURL(sessionID SessionRef, self PoolRef) (_retval string, _err error) {
	_method := "pool.get_wlb_url"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class PoolClass) GetHealthCheckConfig(sessionID SessionRef, self PoolRef) (_retval map[string]string, _err error) {
	_method := "pool.get_health_check_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class PoolClass) GetGuiConfig(sessionID SessionRef, self PoolRef) (_retval map[string]string, _err error) {
	_method := "pool.get_gui_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class PoolClass) GetTags(sessionID SessionRef, self PoolRef) (_retval []string, _err error) {
	_method := "pool.get_tags"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class PoolClass) GetBlobs(sessionID SessionRef, self PoolRef) (_retval map[string]BlobRef, _err error) {
	_method := "pool.get_blobs"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertStringToBlobRefMapToGo(_method + " -> ", _result.Value)
	return
}

func (_class PoolClass) GetHaOvercommitted(sessionID SessionRef, self PoolRef) (_retval bool, _err error) {
	_method := "pool.get_ha_overcommitted"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class PoolClass) GetHaAllowOvercommit(sessionID SessionRef, self PoolRef) (_retval bool, _err error) {
	_method := "pool.get_ha_allow_overcommit"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class PoolClass) GetHaPlanExistsFor(sessionID SessionRef, self PoolRef) (_retval int, _err error) {
	_method := "pool.get_ha_plan_exists_for"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class PoolClass) GetHaHostFailuresToTolerate(sessionID SessionRef, self PoolRef) (_retval int, _err error) {
	_method := "pool.get_ha_host_failures_to_tolerate"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class PoolClass) GetHaStatefiles(sessionID SessionRef, self PoolRef) (_retval []string, _err error) {
	_method := "pool.get_ha_statefiles"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class PoolClass) GetHaConfiguration(sessionID SessionRef, self PoolRef) (_retval map[string]string, _err error) {
	_method := "pool.get_ha_configuration"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class PoolClass) GetHaEnabled(sessionID SessionRef, self PoolRef) (_retval bool, _err error) {
	_method := "pool.get_ha_enabled"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class PoolClass) GetOtherConfig(sessionID SessionRef, self PoolRef) (_retval map[string]string, _err error) {
	_method := "pool.get_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class PoolClass) GetCrashDumpSR(sessionID SessionRef, self PoolRef) (_retval SRRef, _err error) {
	_method := "pool.get_crash_dump_SR"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertSRRefToGo(_method + " -> ", _result.Value)
	return
}

func (_class PoolClass) GetSuspendImageSR(sessionID SessionRef, self PoolRef) (_retval SRRef, _err error) {
	_method := "pool.get_suspend_image_SR"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertSRRefToGo(_method + " -> ", _result.Value)
	return
}

func (_class PoolClass) GetDefaultSR(sessionID SessionRef, self PoolRef) (_retval SRRef, _err error) {
	_method := "pool.get_default_SR"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertSRRefToGo(_method + " -> ", _result.Value)
	return
}

func (_class PoolClass) GetMaster(sessionID SessionRef, self PoolRef) (_retval HostRef, _err error) {
	_method := "pool.get_master"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class PoolClass) GetNameDescription(sessionID SessionRef, self PoolRef) (_retval string, _err error) {
	_method := "pool.get_name_description"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class PoolClass) GetNameLabel(sessionID SessionRef, self PoolRef) (_retval string, _err error) {
	_method := "pool.get_name_label"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class PoolClass) GetUUID(sessionID SessionRef, self PoolRef) (_retval string, _err error) {
	_method := "pool.get_uuid"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class PoolClass) GetByUUID(sessionID SessionRef, uuid string) (_retval PoolRef, _err error) {
	_method := "pool.get_by_uuid"
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
	_retval, _err = convertPoolRefToGo(_method + " -> ", _result.Value)
	return
}

func (_class PoolClass) GetRecord(sessionID SessionRef, self PoolRef) (_retval PoolRecord, _err error) {
	_method := "pool.get_record"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertPoolRecordToGo(_method + " -> ", _result.Value)
	return
}
