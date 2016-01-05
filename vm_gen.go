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

type VMPowerState string

const (
	VMPowerStateHalted VMPowerState = "Halted"
	VMPowerStatePaused VMPowerState = "Paused"
	VMPowerStateRunning VMPowerState = "Running"
	VMPowerStateSuspended VMPowerState = "Suspended"
)

type OnNormalExit string

const (
	OnNormalExitDestroy OnNormalExit = "destroy"
	OnNormalExitRestart OnNormalExit = "restart"
)

type OnCrashBehaviour string

const (
	OnCrashBehaviourDestroy OnCrashBehaviour = "destroy"
	OnCrashBehaviourCoredumpAndDestroy OnCrashBehaviour = "coredump_and_destroy"
	OnCrashBehaviourRestart OnCrashBehaviour = "restart"
	OnCrashBehaviourCoredumpAndRestart OnCrashBehaviour = "coredump_and_restart"
	OnCrashBehaviourPreserve OnCrashBehaviour = "preserve"
	OnCrashBehaviourRenameRestart OnCrashBehaviour = "rename_restart"
)

type VMOperations string

const (
	VMOperationsSnapshot VMOperations = "snapshot"
	VMOperationsClone VMOperations = "clone"
	VMOperationsCopy VMOperations = "copy"
	VMOperationsCreateTemplate VMOperations = "create_template"
	VMOperationsRevert VMOperations = "revert"
	VMOperationsCheckpoint VMOperations = "checkpoint"
	VMOperationsSnapshotWithQuiesce VMOperations = "snapshot_with_quiesce"
	VMOperationsProvision VMOperations = "provision"
	VMOperationsStart VMOperations = "start"
	VMOperationsStartOn VMOperations = "start_on"
	VMOperationsPause VMOperations = "pause"
	VMOperationsUnpause VMOperations = "unpause"
	VMOperationsCleanShutdown VMOperations = "clean_shutdown"
	VMOperationsCleanReboot VMOperations = "clean_reboot"
	VMOperationsHardShutdown VMOperations = "hard_shutdown"
	VMOperationsPowerStateReset VMOperations = "power_state_reset"
	VMOperationsHardReboot VMOperations = "hard_reboot"
	VMOperationsSuspend VMOperations = "suspend"
	VMOperationsCsvm VMOperations = "csvm"
	VMOperationsResume VMOperations = "resume"
	VMOperationsResumeOn VMOperations = "resume_on"
	VMOperationsPoolMigrate VMOperations = "pool_migrate"
	VMOperationsMigrateSend VMOperations = "migrate_send"
	VMOperationsGetBootRecord VMOperations = "get_boot_record"
	VMOperationsSendSysrq VMOperations = "send_sysrq"
	VMOperationsSendTrigger VMOperations = "send_trigger"
	VMOperationsQueryServices VMOperations = "query_services"
	VMOperationsShutdown VMOperations = "shutdown"
	VMOperationsCallPlugin VMOperations = "call_plugin"
	VMOperationsChangingMemoryLive VMOperations = "changing_memory_live"
	VMOperationsAwaitingMemoryLive VMOperations = "awaiting_memory_live"
	VMOperationsChangingDynamicRange VMOperations = "changing_dynamic_range"
	VMOperationsChangingStaticRange VMOperations = "changing_static_range"
	VMOperationsChangingMemoryLimits VMOperations = "changing_memory_limits"
	VMOperationsChangingShadowMemory VMOperations = "changing_shadow_memory"
	VMOperationsChangingShadowMemoryLive VMOperations = "changing_shadow_memory_live"
	VMOperationsChangingVCPUs VMOperations = "changing_VCPUs"
	VMOperationsChangingVCPUsLive VMOperations = "changing_VCPUs_live"
	VMOperationsAssertOperationValid VMOperations = "assert_operation_valid"
	VMOperationsDataSourceOp VMOperations = "data_source_op"
	VMOperationsUpdateAllowedOperations VMOperations = "update_allowed_operations"
	VMOperationsMakeIntoTemplate VMOperations = "make_into_template"
	VMOperationsImport VMOperations = "import"
	VMOperationsExport VMOperations = "export"
	VMOperationsMetadataExport VMOperations = "metadata_export"
	VMOperationsReverting VMOperations = "reverting"
	VMOperationsDestroy VMOperations = "destroy"
)

type VMRecord struct {
	UUID string
	AllowedOperations []VMOperations
	CurrentOperations map[string]VMOperations
	PowerState VMPowerState
	NameLabel string
	NameDescription string
	UserVersion int
	IsATemplate bool
	SuspendVDI VDIRef
	ResidentOn HostRef
	Affinity HostRef
	MemoryOverhead int
	MemoryTarget int
	MemoryStaticMax int
	MemoryDynamicMax int
	MemoryDynamicMin int
	MemoryStaticMin int
	VCPUsParams map[string]string
	VCPUsMax int
	VCPUsAtStartup int
	ActionsAfterShutdown OnNormalExit
	ActionsAfterReboot OnNormalExit
	ActionsAfterCrash OnCrashBehaviour
	Consoles []ConsoleRef
	VIFs []VIFRef
	VBDs []VBDRef
	CrashDumps []CrashdumpRef
	VTPMs []VTPMRef
	PVBootloader string
	PVKernel string
	PVRamdisk string
	PVArgs string
	PVBootloaderArgs string
	PVLegacyArgs string
	HVMBootPolicy string
	HVMBootParams map[string]string
	HVMShadowMultiplier float64
	Platform map[string]string
	PCIBus string
	OtherConfig map[string]string
	Domid int
	Domarch string
	LastBootCPUFlags map[string]string
	IsControlDomain bool
	Metrics VMMetricsRef
	GuestMetrics VMGuestMetricsRef
	LastBootedRecord string
	Recommendations string
	XenstoreData map[string]string
	HaAlwaysRun bool
	HaRestartPriority string
	IsASnapshot bool
	SnapshotOf VMRef
	Snapshots []VMRef
	SnapshotTime time.Time
	TransportableSnapshotID string
	Blobs map[string]BlobRef
	Tags []string
	BlockedOperations map[VMOperations]string
	SnapshotInfo map[string]string
	SnapshotMetadata string
	Parent VMRef
	Children []VMRef
	BiosStrings map[string]string
	ProtectionPolicy VMPPRef
	IsSnapshotFromVmpp bool
	Appliance VMApplianceRef
	StartDelay int
	ShutdownDelay int
	Order int
	VGPUs []VGPURef
	AttachedPCIs []PCIRef
	SuspendSR SRRef
	Version int
	GenerationID string
	HardwarePlatformVersion int
	AutoUpdateDrivers bool
}

type VMRef string

type VMClass struct {
	client *Client
}

func (client *Client) VM() VMClass {
	return VMClass{client}
}

func (_class VMClass) GetAllRecords(sessionID SessionRef) (_retval map[VMRef]VMRecord, _err error) {
	_method := "VM.get_all_records"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVMRefToVMRecordMapToGo(_method + " -> ", _result.Value)
	return
}

func (_class VMClass) GetAll(sessionID SessionRef) (_retval []VMRef, _err error) {
	_method := "VM.get_all"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVMRefSetToGo(_method + " -> ", _result.Value)
	return
}

func (_class VMClass) XenprepStart(sessionID SessionRef, self VMRef) (_err error) {
	_method := "VM.xenprep_start"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg)
	return
}

func (_class VMClass) Import(sessionID SessionRef, url string, sr SRRef, fullRestore bool, force bool) (_retval []VMRef, _err error) {
	_method := "VM.import"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_urlArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "url"), url)
	if _err != nil {
		return
	}
	_srArg, _err := convertSRRefToXen(fmt.Sprintf("%s(%s)", _method, "sr"), sr)
	if _err != nil {
		return
	}
	_fullRestoreArg, _err := convertBoolToXen(fmt.Sprintf("%s(%s)", _method, "full_restore"), fullRestore)
	if _err != nil {
		return
	}
	_forceArg, _err := convertBoolToXen(fmt.Sprintf("%s(%s)", _method, "force"), force)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _urlArg, _srArg, _fullRestoreArg, _forceArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVMRefSetToGo(_method + " -> ", _result.Value)
	return
}

func (_class VMClass) AssertCanSetAutoUpdateDrivers(sessionID SessionRef, self VMRef, value bool) (_err error) {
	_method := "VM.assert_can_set_auto_update_drivers"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class VMClass) SetAutoUpdateDrivers(sessionID SessionRef, self VMRef, value bool) (_err error) {
	_method := "VM.set_auto_update_drivers"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class VMClass) CallPlugin(sessionID SessionRef, vm VMRef, plugin string, fn string, args map[string]string) (_retval string, _err error) {
	_method := "VM.call_plugin"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_vmArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "vm"), vm)
	if _err != nil {
		return
	}
	_pluginArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "plugin"), plugin)
	if _err != nil {
		return
	}
	_fnArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "fn"), fn)
	if _err != nil {
		return
	}
	_argsArg, _err := convertStringToStringMapToXen(fmt.Sprintf("%s(%s)", _method, "args"), args)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _vmArg, _pluginArg, _fnArg, _argsArg)
	if _err != nil {
		return
	}
	_retval, _err = convertStringToGo(_method + " -> ", _result.Value)
	return
}

func (_class VMClass) QueryServices(sessionID SessionRef, self VMRef) (_retval map[string]string, _err error) {
	_method := "VM.query_services"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class VMClass) SetAppliance(sessionID SessionRef, self VMRef, value VMApplianceRef) (_err error) {
	_method := "VM.set_appliance"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertVMApplianceRefToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}

func (_class VMClass) ImportConvert(sessionID SessionRef, atype string, username string, password string, sr SRRef, remoteConfig map[string]string) (_err error) {
	_method := "VM.import_convert"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_atypeArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "type"), atype)
	if _err != nil {
		return
	}
	_usernameArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "username"), username)
	if _err != nil {
		return
	}
	_passwordArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "password"), password)
	if _err != nil {
		return
	}
	_srArg, _err := convertSRRefToXen(fmt.Sprintf("%s(%s)", _method, "sr"), sr)
	if _err != nil {
		return
	}
	_remoteConfigArg, _err := convertStringToStringMapToXen(fmt.Sprintf("%s(%s)", _method, "remote_config"), remoteConfig)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _atypeArg, _usernameArg, _passwordArg, _srArg, _remoteConfigArg)
	return
}

func (_class VMClass) Recover(sessionID SessionRef, self VMRef, sessionTo SessionRef, force bool) (_err error) {
	_method := "VM.recover"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_sessionToArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_to"), sessionTo)
	if _err != nil {
		return
	}
	_forceArg, _err := convertBoolToXen(fmt.Sprintf("%s(%s)", _method, "force"), force)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _sessionToArg, _forceArg)
	return
}

func (_class VMClass) GetSRsRequiredForRecovery(sessionID SessionRef, self VMRef, sessionTo SessionRef) (_retval []SRRef, _err error) {
	_method := "VM.get_SRs_required_for_recovery"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_sessionToArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_to"), sessionTo)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg, _sessionToArg)
	if _err != nil {
		return
	}
	_retval, _err = convertSRRefSetToGo(_method + " -> ", _result.Value)
	return
}

func (_class VMClass) AssertCanBeRecovered(sessionID SessionRef, self VMRef, sessionTo SessionRef) (_err error) {
	_method := "VM.assert_can_be_recovered"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_sessionToArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_to"), sessionTo)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _sessionToArg)
	return
}

func (_class VMClass) SetSuspendVDI(sessionID SessionRef, self VMRef, value VDIRef) (_err error) {
	_method := "VM.set_suspend_VDI"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertVDIRefToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}

func (_class VMClass) SetOrder(sessionID SessionRef, self VMRef, value int) (_err error) {
	_method := "VM.set_order"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class VMClass) SetShutdownDelay(sessionID SessionRef, self VMRef, value int) (_err error) {
	_method := "VM.set_shutdown_delay"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class VMClass) SetStartDelay(sessionID SessionRef, self VMRef, value int) (_err error) {
	_method := "VM.set_start_delay"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class VMClass) SetProtectionPolicy(sessionID SessionRef, self VMRef, value VMPPRef) (_err error) {
	_method := "VM.set_protection_policy"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertVMPPRefToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}

func (_class VMClass) CopyBiosStrings(sessionID SessionRef, vm VMRef, host HostRef) (_err error) {
	_method := "VM.copy_bios_strings"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_vmArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "vm"), vm)
	if _err != nil {
		return
	}
	_hostArg, _err := convertHostRefToXen(fmt.Sprintf("%s(%s)", _method, "host"), host)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _vmArg, _hostArg)
	return
}

func (_class VMClass) RetrieveWlbRecommendations(sessionID SessionRef, vm VMRef) (_retval map[HostRef][]string, _err error) {
	_method := "VM.retrieve_wlb_recommendations"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_vmArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "vm"), vm)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _vmArg)
	if _err != nil {
		return
	}
	_retval, _err = convertHostRefToStringSetMapToGo(_method + " -> ", _result.Value)
	return
}

func (_class VMClass) AssertAgile(sessionID SessionRef, self VMRef) (_err error) {
	_method := "VM.assert_agile"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg)
	return
}

func (_class VMClass) CreateNewBlob(sessionID SessionRef, vm VMRef, name string, mimeType string, public bool) (_retval BlobRef, _err error) {
	_method := "VM.create_new_blob"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_vmArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "vm"), vm)
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
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _vmArg, _nameArg, _mimeTypeArg, _publicArg)
	if _err != nil {
		return
	}
	_retval, _err = convertBlobRefToGo(_method + " -> ", _result.Value)
	return
}

func (_class VMClass) AssertCanBootHere(sessionID SessionRef, self VMRef, host HostRef) (_err error) {
	_method := "VM.assert_can_boot_here"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_hostArg, _err := convertHostRefToXen(fmt.Sprintf("%s(%s)", _method, "host"), host)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _hostArg)
	return
}

func (_class VMClass) GetPossibleHosts(sessionID SessionRef, vm VMRef) (_retval []HostRef, _err error) {
	_method := "VM.get_possible_hosts"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_vmArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "vm"), vm)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _vmArg)
	if _err != nil {
		return
	}
	_retval, _err = convertHostRefSetToGo(_method + " -> ", _result.Value)
	return
}

func (_class VMClass) GetAllowedVIFDevices(sessionID SessionRef, vm VMRef) (_retval []string, _err error) {
	_method := "VM.get_allowed_VIF_devices"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_vmArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "vm"), vm)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _vmArg)
	if _err != nil {
		return
	}
	_retval, _err = convertStringSetToGo(_method + " -> ", _result.Value)
	return
}

func (_class VMClass) GetAllowedVBDDevices(sessionID SessionRef, vm VMRef) (_retval []string, _err error) {
	_method := "VM.get_allowed_VBD_devices"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_vmArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "vm"), vm)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _vmArg)
	if _err != nil {
		return
	}
	_retval, _err = convertStringSetToGo(_method + " -> ", _result.Value)
	return
}

func (_class VMClass) UpdateAllowedOperations(sessionID SessionRef, self VMRef) (_err error) {
	_method := "VM.update_allowed_operations"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg)
	return
}

func (_class VMClass) AssertOperationValid(sessionID SessionRef, self VMRef, op VMOperations) (_err error) {
	_method := "VM.assert_operation_valid"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_opArg, _err := convertEnumVMOperationsToXen(fmt.Sprintf("%s(%s)", _method, "op"), op)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _opArg)
	return
}

func (_class VMClass) ForgetDataSourceArchives(sessionID SessionRef, self VMRef, dataSource string) (_err error) {
	_method := "VM.forget_data_source_archives"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_dataSourceArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "data_source"), dataSource)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _dataSourceArg)
	return
}

func (_class VMClass) QueryDataSource(sessionID SessionRef, self VMRef, dataSource string) (_retval float64, _err error) {
	_method := "VM.query_data_source"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_dataSourceArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "data_source"), dataSource)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg, _dataSourceArg)
	if _err != nil {
		return
	}
	_retval, _err = convertFloatToGo(_method + " -> ", _result.Value)
	return
}

func (_class VMClass) RecordDataSource(sessionID SessionRef, self VMRef, dataSource string) (_err error) {
	_method := "VM.record_data_source"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_dataSourceArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "data_source"), dataSource)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _dataSourceArg)
	return
}

func (_class VMClass) GetDataSources(sessionID SessionRef, self VMRef) (_retval []DataSourceRecord, _err error) {
	_method := "VM.get_data_sources"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertDataSourceRecordSetToGo(_method + " -> ", _result.Value)
	return
}

func (_class VMClass) GetBootRecord(sessionID SessionRef, self VMRef) (_retval VMRecord, _err error) {
	_method := "VM.get_boot_record"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVMRecordToGo(_method + " -> ", _result.Value)
	return
}

func (_class VMClass) AssertCanMigrate(sessionID SessionRef, vm VMRef, dest map[string]string, live bool, vdiMap map[VDIRef]SRRef, vifMap map[VIFRef]NetworkRef, options map[string]string) (_err error) {
	_method := "VM.assert_can_migrate"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_vmArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "vm"), vm)
	if _err != nil {
		return
	}
	_destArg, _err := convertStringToStringMapToXen(fmt.Sprintf("%s(%s)", _method, "dest"), dest)
	if _err != nil {
		return
	}
	_liveArg, _err := convertBoolToXen(fmt.Sprintf("%s(%s)", _method, "live"), live)
	if _err != nil {
		return
	}
	_vdiMapArg, _err := convertVDIRefToSRRefMapToXen(fmt.Sprintf("%s(%s)", _method, "vdi_map"), vdiMap)
	if _err != nil {
		return
	}
	_vifMapArg, _err := convertVIFRefToNetworkRefMapToXen(fmt.Sprintf("%s(%s)", _method, "vif_map"), vifMap)
	if _err != nil {
		return
	}
	_optionsArg, _err := convertStringToStringMapToXen(fmt.Sprintf("%s(%s)", _method, "options"), options)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _vmArg, _destArg, _liveArg, _vdiMapArg, _vifMapArg, _optionsArg)
	return
}

func (_class VMClass) MigrateSend(sessionID SessionRef, vm VMRef, dest map[string]string, live bool, vdiMap map[VDIRef]SRRef, vifMap map[VIFRef]NetworkRef, options map[string]string) (_retval VMRef, _err error) {
	_method := "VM.migrate_send"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_vmArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "vm"), vm)
	if _err != nil {
		return
	}
	_destArg, _err := convertStringToStringMapToXen(fmt.Sprintf("%s(%s)", _method, "dest"), dest)
	if _err != nil {
		return
	}
	_liveArg, _err := convertBoolToXen(fmt.Sprintf("%s(%s)", _method, "live"), live)
	if _err != nil {
		return
	}
	_vdiMapArg, _err := convertVDIRefToSRRefMapToXen(fmt.Sprintf("%s(%s)", _method, "vdi_map"), vdiMap)
	if _err != nil {
		return
	}
	_vifMapArg, _err := convertVIFRefToNetworkRefMapToXen(fmt.Sprintf("%s(%s)", _method, "vif_map"), vifMap)
	if _err != nil {
		return
	}
	_optionsArg, _err := convertStringToStringMapToXen(fmt.Sprintf("%s(%s)", _method, "options"), options)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _vmArg, _destArg, _liveArg, _vdiMapArg, _vifMapArg, _optionsArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVMRefToGo(_method + " -> ", _result.Value)
	return
}

func (_class VMClass) MaximiseMemory(sessionID SessionRef, self VMRef, total int, approximate bool) (_retval int, _err error) {
	_method := "VM.maximise_memory"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_totalArg, _err := convertIntToXen(fmt.Sprintf("%s(%s)", _method, "total"), total)
	if _err != nil {
		return
	}
	_approximateArg, _err := convertBoolToXen(fmt.Sprintf("%s(%s)", _method, "approximate"), approximate)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg, _totalArg, _approximateArg)
	if _err != nil {
		return
	}
	_retval, _err = convertIntToGo(_method + " -> ", _result.Value)
	return
}

func (_class VMClass) SendTrigger(sessionID SessionRef, vm VMRef, trigger string) (_err error) {
	_method := "VM.send_trigger"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_vmArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "vm"), vm)
	if _err != nil {
		return
	}
	_triggerArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "trigger"), trigger)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _vmArg, _triggerArg)
	return
}

func (_class VMClass) SendSysrq(sessionID SessionRef, vm VMRef, key string) (_err error) {
	_method := "VM.send_sysrq"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_vmArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "vm"), vm)
	if _err != nil {
		return
	}
	_keyArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "key"), key)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _vmArg, _keyArg)
	return
}

func (_class VMClass) SetVCPUsAtStartup(sessionID SessionRef, self VMRef, value int) (_err error) {
	_method := "VM.set_VCPUs_at_startup"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class VMClass) SetVCPUsMax(sessionID SessionRef, self VMRef, value int) (_err error) {
	_method := "VM.set_VCPUs_max"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class VMClass) SetShadowMultiplierLive(sessionID SessionRef, self VMRef, multiplier float64) (_err error) {
	_method := "VM.set_shadow_multiplier_live"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_multiplierArg, _err := convertFloatToXen(fmt.Sprintf("%s(%s)", _method, "multiplier"), multiplier)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _multiplierArg)
	return
}

func (_class VMClass) SetHVMShadowMultiplier(sessionID SessionRef, self VMRef, value float64) (_err error) {
	_method := "VM.set_HVM_shadow_multiplier"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertFloatToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}

func (_class VMClass) GetCooperative(sessionID SessionRef, self VMRef) (_retval bool, _err error) {
	_method := "VM.get_cooperative"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class VMClass) WaitMemoryTargetLive(sessionID SessionRef, self VMRef) (_err error) {
	_method := "VM.wait_memory_target_live"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg)
	return
}

func (_class VMClass) SetMemoryTargetLive(sessionID SessionRef, self VMRef, target int) (_err error) {
	_method := "VM.set_memory_target_live"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_targetArg, _err := convertIntToXen(fmt.Sprintf("%s(%s)", _method, "target"), target)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _targetArg)
	return
}

func (_class VMClass) SetMemoryLimits(sessionID SessionRef, self VMRef, staticMin int, staticMax int, dynamicMin int, dynamicMax int) (_err error) {
	_method := "VM.set_memory_limits"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_staticMinArg, _err := convertIntToXen(fmt.Sprintf("%s(%s)", _method, "static_min"), staticMin)
	if _err != nil {
		return
	}
	_staticMaxArg, _err := convertIntToXen(fmt.Sprintf("%s(%s)", _method, "static_max"), staticMax)
	if _err != nil {
		return
	}
	_dynamicMinArg, _err := convertIntToXen(fmt.Sprintf("%s(%s)", _method, "dynamic_min"), dynamicMin)
	if _err != nil {
		return
	}
	_dynamicMaxArg, _err := convertIntToXen(fmt.Sprintf("%s(%s)", _method, "dynamic_max"), dynamicMax)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _staticMinArg, _staticMaxArg, _dynamicMinArg, _dynamicMaxArg)
	return
}

func (_class VMClass) SetMemoryStaticRange(sessionID SessionRef, self VMRef, min int, max int) (_err error) {
	_method := "VM.set_memory_static_range"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_minArg, _err := convertIntToXen(fmt.Sprintf("%s(%s)", _method, "min"), min)
	if _err != nil {
		return
	}
	_maxArg, _err := convertIntToXen(fmt.Sprintf("%s(%s)", _method, "max"), max)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _minArg, _maxArg)
	return
}

func (_class VMClass) SetMemoryStaticMin(sessionID SessionRef, self VMRef, value int) (_err error) {
	_method := "VM.set_memory_static_min"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class VMClass) SetMemoryStaticMax(sessionID SessionRef, self VMRef, value int) (_err error) {
	_method := "VM.set_memory_static_max"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class VMClass) SetMemoryDynamicRange(sessionID SessionRef, self VMRef, min int, max int) (_err error) {
	_method := "VM.set_memory_dynamic_range"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_minArg, _err := convertIntToXen(fmt.Sprintf("%s(%s)", _method, "min"), min)
	if _err != nil {
		return
	}
	_maxArg, _err := convertIntToXen(fmt.Sprintf("%s(%s)", _method, "max"), max)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _minArg, _maxArg)
	return
}

func (_class VMClass) SetMemoryDynamicMin(sessionID SessionRef, self VMRef, value int) (_err error) {
	_method := "VM.set_memory_dynamic_min"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class VMClass) SetMemoryDynamicMax(sessionID SessionRef, self VMRef, value int) (_err error) {
	_method := "VM.set_memory_dynamic_max"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class VMClass) ComputeMemoryOverhead(sessionID SessionRef, vm VMRef) (_retval int, _err error) {
	_method := "VM.compute_memory_overhead"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_vmArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "vm"), vm)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _vmArg)
	if _err != nil {
		return
	}
	_retval, _err = convertIntToGo(_method + " -> ", _result.Value)
	return
}

func (_class VMClass) SetHaAlwaysRun(sessionID SessionRef, self VMRef, value bool) (_err error) {
	_method := "VM.set_ha_always_run"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class VMClass) SetHaRestartPriority(sessionID SessionRef, self VMRef, value string) (_err error) {
	_method := "VM.set_ha_restart_priority"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class VMClass) AddToVCPUsParamsLive(sessionID SessionRef, self VMRef, key string, value string) (_err error) {
	_method := "VM.add_to_VCPUs_params_live"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class VMClass) SetVCPUsNumberLive(sessionID SessionRef, self VMRef, nvcpu int) (_err error) {
	_method := "VM.set_VCPUs_number_live"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_nvcpuArg, _err := convertIntToXen(fmt.Sprintf("%s(%s)", _method, "nvcpu"), nvcpu)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _nvcpuArg)
	return
}

func (_class VMClass) PoolMigrate(sessionID SessionRef, vm VMRef, host HostRef, options map[string]string) (_err error) {
	_method := "VM.pool_migrate"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_vmArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "vm"), vm)
	if _err != nil {
		return
	}
	_hostArg, _err := convertHostRefToXen(fmt.Sprintf("%s(%s)", _method, "host"), host)
	if _err != nil {
		return
	}
	_optionsArg, _err := convertStringToStringMapToXen(fmt.Sprintf("%s(%s)", _method, "options"), options)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _vmArg, _hostArg, _optionsArg)
	return
}

func (_class VMClass) ResumeOn(sessionID SessionRef, vm VMRef, host HostRef, startPaused bool, force bool) (_err error) {
	_method := "VM.resume_on"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_vmArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "vm"), vm)
	if _err != nil {
		return
	}
	_hostArg, _err := convertHostRefToXen(fmt.Sprintf("%s(%s)", _method, "host"), host)
	if _err != nil {
		return
	}
	_startPausedArg, _err := convertBoolToXen(fmt.Sprintf("%s(%s)", _method, "start_paused"), startPaused)
	if _err != nil {
		return
	}
	_forceArg, _err := convertBoolToXen(fmt.Sprintf("%s(%s)", _method, "force"), force)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _vmArg, _hostArg, _startPausedArg, _forceArg)
	return
}

func (_class VMClass) Resume(sessionID SessionRef, vm VMRef, startPaused bool, force bool) (_err error) {
	_method := "VM.resume"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_vmArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "vm"), vm)
	if _err != nil {
		return
	}
	_startPausedArg, _err := convertBoolToXen(fmt.Sprintf("%s(%s)", _method, "start_paused"), startPaused)
	if _err != nil {
		return
	}
	_forceArg, _err := convertBoolToXen(fmt.Sprintf("%s(%s)", _method, "force"), force)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _vmArg, _startPausedArg, _forceArg)
	return
}

func (_class VMClass) Suspend(sessionID SessionRef, vm VMRef) (_err error) {
	_method := "VM.suspend"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_vmArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "vm"), vm)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _vmArg)
	return
}

func (_class VMClass) HardReboot(sessionID SessionRef, vm VMRef) (_err error) {
	_method := "VM.hard_reboot"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_vmArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "vm"), vm)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _vmArg)
	return
}

func (_class VMClass) PowerStateReset(sessionID SessionRef, vm VMRef) (_err error) {
	_method := "VM.power_state_reset"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_vmArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "vm"), vm)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _vmArg)
	return
}

func (_class VMClass) HardShutdown(sessionID SessionRef, vm VMRef) (_err error) {
	_method := "VM.hard_shutdown"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_vmArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "vm"), vm)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _vmArg)
	return
}

func (_class VMClass) CleanReboot(sessionID SessionRef, vm VMRef) (_err error) {
	_method := "VM.clean_reboot"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_vmArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "vm"), vm)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _vmArg)
	return
}

func (_class VMClass) Shutdown(sessionID SessionRef, vm VMRef) (_err error) {
	_method := "VM.shutdown"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_vmArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "vm"), vm)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _vmArg)
	return
}

func (_class VMClass) CleanShutdown(sessionID SessionRef, vm VMRef) (_err error) {
	_method := "VM.clean_shutdown"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_vmArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "vm"), vm)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _vmArg)
	return
}

func (_class VMClass) Unpause(sessionID SessionRef, vm VMRef) (_err error) {
	_method := "VM.unpause"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_vmArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "vm"), vm)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _vmArg)
	return
}

func (_class VMClass) Pause(sessionID SessionRef, vm VMRef) (_err error) {
	_method := "VM.pause"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_vmArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "vm"), vm)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _vmArg)
	return
}

func (_class VMClass) StartOn(sessionID SessionRef, vm VMRef, host HostRef, startPaused bool, force bool) (_err error) {
	_method := "VM.start_on"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_vmArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "vm"), vm)
	if _err != nil {
		return
	}
	_hostArg, _err := convertHostRefToXen(fmt.Sprintf("%s(%s)", _method, "host"), host)
	if _err != nil {
		return
	}
	_startPausedArg, _err := convertBoolToXen(fmt.Sprintf("%s(%s)", _method, "start_paused"), startPaused)
	if _err != nil {
		return
	}
	_forceArg, _err := convertBoolToXen(fmt.Sprintf("%s(%s)", _method, "force"), force)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _vmArg, _hostArg, _startPausedArg, _forceArg)
	return
}

func (_class VMClass) Start(sessionID SessionRef, vm VMRef, startPaused bool, force bool) (_err error) {
	_method := "VM.start"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_vmArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "vm"), vm)
	if _err != nil {
		return
	}
	_startPausedArg, _err := convertBoolToXen(fmt.Sprintf("%s(%s)", _method, "start_paused"), startPaused)
	if _err != nil {
		return
	}
	_forceArg, _err := convertBoolToXen(fmt.Sprintf("%s(%s)", _method, "force"), force)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _vmArg, _startPausedArg, _forceArg)
	return
}

func (_class VMClass) Provision(sessionID SessionRef, vm VMRef) (_err error) {
	_method := "VM.provision"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_vmArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "vm"), vm)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _vmArg)
	return
}

func (_class VMClass) Checkpoint(sessionID SessionRef, vm VMRef, newName string) (_retval VMRef, _err error) {
	_method := "VM.checkpoint"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_vmArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "vm"), vm)
	if _err != nil {
		return
	}
	_newNameArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "new_name"), newName)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _vmArg, _newNameArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVMRefToGo(_method + " -> ", _result.Value)
	return
}

func (_class VMClass) Revert(sessionID SessionRef, snapshot VMRef) (_err error) {
	_method := "VM.revert"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_snapshotArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "snapshot"), snapshot)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _snapshotArg)
	return
}

func (_class VMClass) Copy(sessionID SessionRef, vm VMRef, newName string, sr SRRef) (_retval VMRef, _err error) {
	_method := "VM.copy"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_vmArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "vm"), vm)
	if _err != nil {
		return
	}
	_newNameArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "new_name"), newName)
	if _err != nil {
		return
	}
	_srArg, _err := convertSRRefToXen(fmt.Sprintf("%s(%s)", _method, "sr"), sr)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _vmArg, _newNameArg, _srArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVMRefToGo(_method + " -> ", _result.Value)
	return
}

func (_class VMClass) Clone(sessionID SessionRef, vm VMRef, newName string) (_retval VMRef, _err error) {
	_method := "VM.clone"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_vmArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "vm"), vm)
	if _err != nil {
		return
	}
	_newNameArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "new_name"), newName)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _vmArg, _newNameArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVMRefToGo(_method + " -> ", _result.Value)
	return
}

func (_class VMClass) SnapshotWithQuiesce(sessionID SessionRef, vm VMRef, newName string) (_retval VMRef, _err error) {
	_method := "VM.snapshot_with_quiesce"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_vmArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "vm"), vm)
	if _err != nil {
		return
	}
	_newNameArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "new_name"), newName)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _vmArg, _newNameArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVMRefToGo(_method + " -> ", _result.Value)
	return
}

func (_class VMClass) Snapshot(sessionID SessionRef, vm VMRef, newName string) (_retval VMRef, _err error) {
	_method := "VM.snapshot"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_vmArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "vm"), vm)
	if _err != nil {
		return
	}
	_newNameArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "new_name"), newName)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _vmArg, _newNameArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVMRefToGo(_method + " -> ", _result.Value)
	return
}

func (_class VMClass) SetHardwarePlatformVersion(sessionID SessionRef, self VMRef, value int) (_err error) {
	_method := "VM.set_hardware_platform_version"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class VMClass) SetSuspendSR(sessionID SessionRef, self VMRef, value SRRef) (_err error) {
	_method := "VM.set_suspend_SR"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class VMClass) RemoveFromBlockedOperations(sessionID SessionRef, self VMRef, key VMOperations) (_err error) {
	_method := "VM.remove_from_blocked_operations"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_keyArg, _err := convertEnumVMOperationsToXen(fmt.Sprintf("%s(%s)", _method, "key"), key)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _keyArg)
	return
}

func (_class VMClass) AddToBlockedOperations(sessionID SessionRef, self VMRef, key VMOperations, value string) (_err error) {
	_method := "VM.add_to_blocked_operations"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_keyArg, _err := convertEnumVMOperationsToXen(fmt.Sprintf("%s(%s)", _method, "key"), key)
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

func (_class VMClass) SetBlockedOperations(sessionID SessionRef, self VMRef, value map[VMOperations]string) (_err error) {
	_method := "VM.set_blocked_operations"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertEnumVMOperationsToStringMapToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}

func (_class VMClass) RemoveTags(sessionID SessionRef, self VMRef, value string) (_err error) {
	_method := "VM.remove_tags"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class VMClass) AddTags(sessionID SessionRef, self VMRef, value string) (_err error) {
	_method := "VM.add_tags"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class VMClass) SetTags(sessionID SessionRef, self VMRef, value []string) (_err error) {
	_method := "VM.set_tags"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class VMClass) RemoveFromXenstoreData(sessionID SessionRef, self VMRef, key string) (_err error) {
	_method := "VM.remove_from_xenstore_data"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class VMClass) AddToXenstoreData(sessionID SessionRef, self VMRef, key string, value string) (_err error) {
	_method := "VM.add_to_xenstore_data"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class VMClass) SetXenstoreData(sessionID SessionRef, self VMRef, value map[string]string) (_err error) {
	_method := "VM.set_xenstore_data"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class VMClass) SetRecommendations(sessionID SessionRef, self VMRef, value string) (_err error) {
	_method := "VM.set_recommendations"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class VMClass) RemoveFromOtherConfig(sessionID SessionRef, self VMRef, key string) (_err error) {
	_method := "VM.remove_from_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class VMClass) AddToOtherConfig(sessionID SessionRef, self VMRef, key string, value string) (_err error) {
	_method := "VM.add_to_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class VMClass) SetOtherConfig(sessionID SessionRef, self VMRef, value map[string]string) (_err error) {
	_method := "VM.set_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class VMClass) SetPCIBus(sessionID SessionRef, self VMRef, value string) (_err error) {
	_method := "VM.set_PCI_bus"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class VMClass) RemoveFromPlatform(sessionID SessionRef, self VMRef, key string) (_err error) {
	_method := "VM.remove_from_platform"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class VMClass) AddToPlatform(sessionID SessionRef, self VMRef, key string, value string) (_err error) {
	_method := "VM.add_to_platform"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class VMClass) SetPlatform(sessionID SessionRef, self VMRef, value map[string]string) (_err error) {
	_method := "VM.set_platform"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class VMClass) RemoveFromHVMBootParams(sessionID SessionRef, self VMRef, key string) (_err error) {
	_method := "VM.remove_from_HVM_boot_params"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class VMClass) AddToHVMBootParams(sessionID SessionRef, self VMRef, key string, value string) (_err error) {
	_method := "VM.add_to_HVM_boot_params"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class VMClass) SetHVMBootParams(sessionID SessionRef, self VMRef, value map[string]string) (_err error) {
	_method := "VM.set_HVM_boot_params"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class VMClass) SetHVMBootPolicy(sessionID SessionRef, self VMRef, value string) (_err error) {
	_method := "VM.set_HVM_boot_policy"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class VMClass) SetPVLegacyArgs(sessionID SessionRef, self VMRef, value string) (_err error) {
	_method := "VM.set_PV_legacy_args"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class VMClass) SetPVBootloaderArgs(sessionID SessionRef, self VMRef, value string) (_err error) {
	_method := "VM.set_PV_bootloader_args"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class VMClass) SetPVArgs(sessionID SessionRef, self VMRef, value string) (_err error) {
	_method := "VM.set_PV_args"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class VMClass) SetPVRamdisk(sessionID SessionRef, self VMRef, value string) (_err error) {
	_method := "VM.set_PV_ramdisk"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class VMClass) SetPVKernel(sessionID SessionRef, self VMRef, value string) (_err error) {
	_method := "VM.set_PV_kernel"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class VMClass) SetPVBootloader(sessionID SessionRef, self VMRef, value string) (_err error) {
	_method := "VM.set_PV_bootloader"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class VMClass) SetActionsAfterCrash(sessionID SessionRef, self VMRef, value OnCrashBehaviour) (_err error) {
	_method := "VM.set_actions_after_crash"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertEnumOnCrashBehaviourToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}

func (_class VMClass) SetActionsAfterReboot(sessionID SessionRef, self VMRef, value OnNormalExit) (_err error) {
	_method := "VM.set_actions_after_reboot"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertEnumOnNormalExitToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}

func (_class VMClass) SetActionsAfterShutdown(sessionID SessionRef, self VMRef, value OnNormalExit) (_err error) {
	_method := "VM.set_actions_after_shutdown"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertEnumOnNormalExitToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}

func (_class VMClass) RemoveFromVCPUsParams(sessionID SessionRef, self VMRef, key string) (_err error) {
	_method := "VM.remove_from_VCPUs_params"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class VMClass) AddToVCPUsParams(sessionID SessionRef, self VMRef, key string, value string) (_err error) {
	_method := "VM.add_to_VCPUs_params"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class VMClass) SetVCPUsParams(sessionID SessionRef, self VMRef, value map[string]string) (_err error) {
	_method := "VM.set_VCPUs_params"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class VMClass) SetAffinity(sessionID SessionRef, self VMRef, value HostRef) (_err error) {
	_method := "VM.set_affinity"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertHostRefToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}

func (_class VMClass) SetIsATemplate(sessionID SessionRef, self VMRef, value bool) (_err error) {
	_method := "VM.set_is_a_template"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class VMClass) SetUserVersion(sessionID SessionRef, self VMRef, value int) (_err error) {
	_method := "VM.set_user_version"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class VMClass) SetNameDescription(sessionID SessionRef, self VMRef, value string) (_err error) {
	_method := "VM.set_name_description"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class VMClass) SetNameLabel(sessionID SessionRef, self VMRef, value string) (_err error) {
	_method := "VM.set_name_label"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class VMClass) GetAutoUpdateDrivers(sessionID SessionRef, self VMRef) (_retval bool, _err error) {
	_method := "VM.get_auto_update_drivers"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class VMClass) GetHardwarePlatformVersion(sessionID SessionRef, self VMRef) (_retval int, _err error) {
	_method := "VM.get_hardware_platform_version"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class VMClass) GetGenerationID(sessionID SessionRef, self VMRef) (_retval string, _err error) {
	_method := "VM.get_generation_id"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class VMClass) GetVersion(sessionID SessionRef, self VMRef) (_retval int, _err error) {
	_method := "VM.get_version"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class VMClass) GetSuspendSR(sessionID SessionRef, self VMRef) (_retval SRRef, _err error) {
	_method := "VM.get_suspend_SR"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class VMClass) GetAttachedPCIs(sessionID SessionRef, self VMRef) (_retval []PCIRef, _err error) {
	_method := "VM.get_attached_PCIs"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertPCIRefSetToGo(_method + " -> ", _result.Value)
	return
}

func (_class VMClass) GetVGPUs(sessionID SessionRef, self VMRef) (_retval []VGPURef, _err error) {
	_method := "VM.get_VGPUs"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class VMClass) GetOrder(sessionID SessionRef, self VMRef) (_retval int, _err error) {
	_method := "VM.get_order"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class VMClass) GetShutdownDelay(sessionID SessionRef, self VMRef) (_retval int, _err error) {
	_method := "VM.get_shutdown_delay"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class VMClass) GetStartDelay(sessionID SessionRef, self VMRef) (_retval int, _err error) {
	_method := "VM.get_start_delay"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class VMClass) GetAppliance(sessionID SessionRef, self VMRef) (_retval VMApplianceRef, _err error) {
	_method := "VM.get_appliance"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVMApplianceRefToGo(_method + " -> ", _result.Value)
	return
}

func (_class VMClass) GetIsSnapshotFromVmpp(sessionID SessionRef, self VMRef) (_retval bool, _err error) {
	_method := "VM.get_is_snapshot_from_vmpp"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class VMClass) GetProtectionPolicy(sessionID SessionRef, self VMRef) (_retval VMPPRef, _err error) {
	_method := "VM.get_protection_policy"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVMPPRefToGo(_method + " -> ", _result.Value)
	return
}

func (_class VMClass) GetBiosStrings(sessionID SessionRef, self VMRef) (_retval map[string]string, _err error) {
	_method := "VM.get_bios_strings"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class VMClass) GetChildren(sessionID SessionRef, self VMRef) (_retval []VMRef, _err error) {
	_method := "VM.get_children"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVMRefSetToGo(_method + " -> ", _result.Value)
	return
}

func (_class VMClass) GetParent(sessionID SessionRef, self VMRef) (_retval VMRef, _err error) {
	_method := "VM.get_parent"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class VMClass) GetSnapshotMetadata(sessionID SessionRef, self VMRef) (_retval string, _err error) {
	_method := "VM.get_snapshot_metadata"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class VMClass) GetSnapshotInfo(sessionID SessionRef, self VMRef) (_retval map[string]string, _err error) {
	_method := "VM.get_snapshot_info"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class VMClass) GetBlockedOperations(sessionID SessionRef, self VMRef) (_retval map[VMOperations]string, _err error) {
	_method := "VM.get_blocked_operations"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertEnumVMOperationsToStringMapToGo(_method + " -> ", _result.Value)
	return
}

func (_class VMClass) GetTags(sessionID SessionRef, self VMRef) (_retval []string, _err error) {
	_method := "VM.get_tags"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class VMClass) GetBlobs(sessionID SessionRef, self VMRef) (_retval map[string]BlobRef, _err error) {
	_method := "VM.get_blobs"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class VMClass) GetTransportableSnapshotID(sessionID SessionRef, self VMRef) (_retval string, _err error) {
	_method := "VM.get_transportable_snapshot_id"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class VMClass) GetSnapshotTime(sessionID SessionRef, self VMRef) (_retval time.Time, _err error) {
	_method := "VM.get_snapshot_time"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class VMClass) GetSnapshots(sessionID SessionRef, self VMRef) (_retval []VMRef, _err error) {
	_method := "VM.get_snapshots"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVMRefSetToGo(_method + " -> ", _result.Value)
	return
}

func (_class VMClass) GetSnapshotOf(sessionID SessionRef, self VMRef) (_retval VMRef, _err error) {
	_method := "VM.get_snapshot_of"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class VMClass) GetIsASnapshot(sessionID SessionRef, self VMRef) (_retval bool, _err error) {
	_method := "VM.get_is_a_snapshot"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class VMClass) GetHaRestartPriority(sessionID SessionRef, self VMRef) (_retval string, _err error) {
	_method := "VM.get_ha_restart_priority"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class VMClass) GetHaAlwaysRun(sessionID SessionRef, self VMRef) (_retval bool, _err error) {
	_method := "VM.get_ha_always_run"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class VMClass) GetXenstoreData(sessionID SessionRef, self VMRef) (_retval map[string]string, _err error) {
	_method := "VM.get_xenstore_data"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class VMClass) GetRecommendations(sessionID SessionRef, self VMRef) (_retval string, _err error) {
	_method := "VM.get_recommendations"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class VMClass) GetLastBootedRecord(sessionID SessionRef, self VMRef) (_retval string, _err error) {
	_method := "VM.get_last_booted_record"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class VMClass) GetGuestMetrics(sessionID SessionRef, self VMRef) (_retval VMGuestMetricsRef, _err error) {
	_method := "VM.get_guest_metrics"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVMGuestMetricsRefToGo(_method + " -> ", _result.Value)
	return
}

func (_class VMClass) GetMetrics(sessionID SessionRef, self VMRef) (_retval VMMetricsRef, _err error) {
	_method := "VM.get_metrics"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVMMetricsRefToGo(_method + " -> ", _result.Value)
	return
}

func (_class VMClass) GetIsControlDomain(sessionID SessionRef, self VMRef) (_retval bool, _err error) {
	_method := "VM.get_is_control_domain"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class VMClass) GetLastBootCPUFlags(sessionID SessionRef, self VMRef) (_retval map[string]string, _err error) {
	_method := "VM.get_last_boot_CPU_flags"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class VMClass) GetDomarch(sessionID SessionRef, self VMRef) (_retval string, _err error) {
	_method := "VM.get_domarch"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class VMClass) GetDomid(sessionID SessionRef, self VMRef) (_retval int, _err error) {
	_method := "VM.get_domid"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class VMClass) GetOtherConfig(sessionID SessionRef, self VMRef) (_retval map[string]string, _err error) {
	_method := "VM.get_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class VMClass) GetPCIBus(sessionID SessionRef, self VMRef) (_retval string, _err error) {
	_method := "VM.get_PCI_bus"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class VMClass) GetPlatform(sessionID SessionRef, self VMRef) (_retval map[string]string, _err error) {
	_method := "VM.get_platform"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class VMClass) GetHVMShadowMultiplier(sessionID SessionRef, self VMRef) (_retval float64, _err error) {
	_method := "VM.get_HVM_shadow_multiplier"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertFloatToGo(_method + " -> ", _result.Value)
	return
}

func (_class VMClass) GetHVMBootParams(sessionID SessionRef, self VMRef) (_retval map[string]string, _err error) {
	_method := "VM.get_HVM_boot_params"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class VMClass) GetHVMBootPolicy(sessionID SessionRef, self VMRef) (_retval string, _err error) {
	_method := "VM.get_HVM_boot_policy"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class VMClass) GetPVLegacyArgs(sessionID SessionRef, self VMRef) (_retval string, _err error) {
	_method := "VM.get_PV_legacy_args"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class VMClass) GetPVBootloaderArgs(sessionID SessionRef, self VMRef) (_retval string, _err error) {
	_method := "VM.get_PV_bootloader_args"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class VMClass) GetPVArgs(sessionID SessionRef, self VMRef) (_retval string, _err error) {
	_method := "VM.get_PV_args"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class VMClass) GetPVRamdisk(sessionID SessionRef, self VMRef) (_retval string, _err error) {
	_method := "VM.get_PV_ramdisk"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class VMClass) GetPVKernel(sessionID SessionRef, self VMRef) (_retval string, _err error) {
	_method := "VM.get_PV_kernel"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class VMClass) GetPVBootloader(sessionID SessionRef, self VMRef) (_retval string, _err error) {
	_method := "VM.get_PV_bootloader"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class VMClass) GetVTPMs(sessionID SessionRef, self VMRef) (_retval []VTPMRef, _err error) {
	_method := "VM.get_VTPMs"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVTPMRefSetToGo(_method + " -> ", _result.Value)
	return
}

func (_class VMClass) GetCrashDumps(sessionID SessionRef, self VMRef) (_retval []CrashdumpRef, _err error) {
	_method := "VM.get_crash_dumps"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertCrashdumpRefSetToGo(_method + " -> ", _result.Value)
	return
}

func (_class VMClass) GetVBDs(sessionID SessionRef, self VMRef) (_retval []VBDRef, _err error) {
	_method := "VM.get_VBDs"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVBDRefSetToGo(_method + " -> ", _result.Value)
	return
}

func (_class VMClass) GetVIFs(sessionID SessionRef, self VMRef) (_retval []VIFRef, _err error) {
	_method := "VM.get_VIFs"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVIFRefSetToGo(_method + " -> ", _result.Value)
	return
}

func (_class VMClass) GetConsoles(sessionID SessionRef, self VMRef) (_retval []ConsoleRef, _err error) {
	_method := "VM.get_consoles"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertConsoleRefSetToGo(_method + " -> ", _result.Value)
	return
}

func (_class VMClass) GetActionsAfterCrash(sessionID SessionRef, self VMRef) (_retval OnCrashBehaviour, _err error) {
	_method := "VM.get_actions_after_crash"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertEnumOnCrashBehaviourToGo(_method + " -> ", _result.Value)
	return
}

func (_class VMClass) GetActionsAfterReboot(sessionID SessionRef, self VMRef) (_retval OnNormalExit, _err error) {
	_method := "VM.get_actions_after_reboot"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertEnumOnNormalExitToGo(_method + " -> ", _result.Value)
	return
}

func (_class VMClass) GetActionsAfterShutdown(sessionID SessionRef, self VMRef) (_retval OnNormalExit, _err error) {
	_method := "VM.get_actions_after_shutdown"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertEnumOnNormalExitToGo(_method + " -> ", _result.Value)
	return
}

func (_class VMClass) GetVCPUsAtStartup(sessionID SessionRef, self VMRef) (_retval int, _err error) {
	_method := "VM.get_VCPUs_at_startup"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class VMClass) GetVCPUsMax(sessionID SessionRef, self VMRef) (_retval int, _err error) {
	_method := "VM.get_VCPUs_max"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class VMClass) GetVCPUsParams(sessionID SessionRef, self VMRef) (_retval map[string]string, _err error) {
	_method := "VM.get_VCPUs_params"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class VMClass) GetMemoryStaticMin(sessionID SessionRef, self VMRef) (_retval int, _err error) {
	_method := "VM.get_memory_static_min"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class VMClass) GetMemoryDynamicMin(sessionID SessionRef, self VMRef) (_retval int, _err error) {
	_method := "VM.get_memory_dynamic_min"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class VMClass) GetMemoryDynamicMax(sessionID SessionRef, self VMRef) (_retval int, _err error) {
	_method := "VM.get_memory_dynamic_max"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class VMClass) GetMemoryStaticMax(sessionID SessionRef, self VMRef) (_retval int, _err error) {
	_method := "VM.get_memory_static_max"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class VMClass) GetMemoryTarget(sessionID SessionRef, self VMRef) (_retval int, _err error) {
	_method := "VM.get_memory_target"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class VMClass) GetMemoryOverhead(sessionID SessionRef, self VMRef) (_retval int, _err error) {
	_method := "VM.get_memory_overhead"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class VMClass) GetAffinity(sessionID SessionRef, self VMRef) (_retval HostRef, _err error) {
	_method := "VM.get_affinity"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class VMClass) GetResidentOn(sessionID SessionRef, self VMRef) (_retval HostRef, _err error) {
	_method := "VM.get_resident_on"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class VMClass) GetSuspendVDI(sessionID SessionRef, self VMRef) (_retval VDIRef, _err error) {
	_method := "VM.get_suspend_VDI"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class VMClass) GetIsATemplate(sessionID SessionRef, self VMRef) (_retval bool, _err error) {
	_method := "VM.get_is_a_template"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class VMClass) GetUserVersion(sessionID SessionRef, self VMRef) (_retval int, _err error) {
	_method := "VM.get_user_version"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class VMClass) GetNameDescription(sessionID SessionRef, self VMRef) (_retval string, _err error) {
	_method := "VM.get_name_description"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class VMClass) GetNameLabel(sessionID SessionRef, self VMRef) (_retval string, _err error) {
	_method := "VM.get_name_label"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class VMClass) GetPowerState(sessionID SessionRef, self VMRef) (_retval VMPowerState, _err error) {
	_method := "VM.get_power_state"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertEnumVMPowerStateToGo(_method + " -> ", _result.Value)
	return
}

func (_class VMClass) GetCurrentOperations(sessionID SessionRef, self VMRef) (_retval map[string]VMOperations, _err error) {
	_method := "VM.get_current_operations"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertStringToEnumVMOperationsMapToGo(_method + " -> ", _result.Value)
	return
}

func (_class VMClass) GetAllowedOperations(sessionID SessionRef, self VMRef) (_retval []VMOperations, _err error) {
	_method := "VM.get_allowed_operations"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertEnumVMOperationsSetToGo(_method + " -> ", _result.Value)
	return
}

func (_class VMClass) GetUUID(sessionID SessionRef, self VMRef) (_retval string, _err error) {
	_method := "VM.get_uuid"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class VMClass) GetByNameLabel(sessionID SessionRef, label string) (_retval []VMRef, _err error) {
	_method := "VM.get_by_name_label"
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
	_retval, _err = convertVMRefSetToGo(_method + " -> ", _result.Value)
	return
}

func (_class VMClass) Destroy(sessionID SessionRef, self VMRef) (_err error) {
	_method := "VM.destroy"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg)
	return
}

func (_class VMClass) Create(sessionID SessionRef, args VMRecord) (_retval VMRef, _err error) {
	_method := "VM.create"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_argsArg, _err := convertVMRecordToXen(fmt.Sprintf("%s(%s)", _method, "args"), args)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _argsArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVMRefToGo(_method + " -> ", _result.Value)
	return
}

func (_class VMClass) GetByUUID(sessionID SessionRef, uuid string) (_retval VMRef, _err error) {
	_method := "VM.get_by_uuid"
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
	_retval, _err = convertVMRefToGo(_method + " -> ", _result.Value)
	return
}

func (_class VMClass) GetRecord(sessionID SessionRef, self VMRef) (_retval VMRecord, _err error) {
	_method := "VM.get_record"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVMRecordToGo(_method + " -> ", _result.Value)
	return
}
