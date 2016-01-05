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

type IPConfigurationMode string

const (
	IPConfigurationModeNone IPConfigurationMode = "None"
	IPConfigurationModeDHCP IPConfigurationMode = "DHCP"
	IPConfigurationModeStatic IPConfigurationMode = "Static"
)

type Ipv6ConfigurationMode string

const (
	Ipv6ConfigurationModeNone Ipv6ConfigurationMode = "None"
	Ipv6ConfigurationModeDHCP Ipv6ConfigurationMode = "DHCP"
	Ipv6ConfigurationModeStatic Ipv6ConfigurationMode = "Static"
	Ipv6ConfigurationModeAutoconf Ipv6ConfigurationMode = "Autoconf"
)

type PrimaryAddressType string

const (
	PrimaryAddressTypeIPv4 PrimaryAddressType = "IPv4"
	PrimaryAddressTypeIPv6 PrimaryAddressType = "IPv6"
)

type PIFRecord struct {
	UUID string
	Device string
	Network NetworkRef
	Host HostRef
	MAC string
	MTU int
	VLAN int
	Metrics PIFMetricsRef
	Physical bool
	CurrentlyAttached bool
	IPConfigurationMode IPConfigurationMode
	IP string
	Netmask string
	Gateway string
	DNS string
	BondSlaveOf BondRef
	BondMasterOf []BondRef
	VLANMasterOf VLANRef
	VLANSlaveOf []VLANRef
	Management bool
	OtherConfig map[string]string
	DisallowUnplug bool
	TunnelAccessPIFOf []TunnelRef
	TunnelTransportPIFOf []TunnelRef
	Ipv6ConfigurationMode Ipv6ConfigurationMode
	IPv6 []string
	Ipv6Gateway string
	PrimaryAddressType PrimaryAddressType
	Managed bool
	Properties map[string]string
	Capabilities []string
}

type PIFRef string

type PIFClass struct {
	client *Client
}

func (client *Client) PIF() PIFClass {
	return PIFClass{client}
}

func (_class PIFClass) GetAllRecords(sessionID SessionRef) (_retval map[PIFRef]PIFRecord, _err error) {
	_method := "PIF.get_all_records"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg)
	if _err != nil {
		return
	}
	_retval, _err = convertPIFRefToPIFRecordMapToGo(_method + " -> ", _result.Value)
	return
}

func (_class PIFClass) GetAll(sessionID SessionRef) (_retval []PIFRef, _err error) {
	_method := "PIF.get_all"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg)
	if _err != nil {
		return
	}
	_retval, _err = convertPIFRefSetToGo(_method + " -> ", _result.Value)
	return
}

func (_class PIFClass) SetProperty(sessionID SessionRef, self PIFRef, name string, value string) (_err error) {
	_method := "PIF.set_property"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPIFRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class PIFClass) DbForget(sessionID SessionRef, self PIFRef) (_err error) {
	_method := "PIF.db_forget"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPIFRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg)
	return
}

func (_class PIFClass) DbIntroduce(sessionID SessionRef, device string, network NetworkRef, host HostRef, mac string, mtu int, vlan int, physical bool, ipConfigurationMode IPConfigurationMode, ip string, netmask string, gateway string, dns string, bondSlaveOf BondRef, vlanMasterOf VLANRef, management bool, otherConfig map[string]string, disallowUnplug bool, ipv6ConfigurationMode Ipv6ConfigurationMode, ipv6 []string, ipv6Gateway string, primaryAddressType PrimaryAddressType, managed bool, properties map[string]string) (_retval PIFRef, _err error) {
	_method := "PIF.db_introduce"
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
	_hostArg, _err := convertHostRefToXen(fmt.Sprintf("%s(%s)", _method, "host"), host)
	if _err != nil {
		return
	}
	_macArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "MAC"), mac)
	if _err != nil {
		return
	}
	_mtuArg, _err := convertIntToXen(fmt.Sprintf("%s(%s)", _method, "MTU"), mtu)
	if _err != nil {
		return
	}
	_vlanArg, _err := convertIntToXen(fmt.Sprintf("%s(%s)", _method, "VLAN"), vlan)
	if _err != nil {
		return
	}
	_physicalArg, _err := convertBoolToXen(fmt.Sprintf("%s(%s)", _method, "physical"), physical)
	if _err != nil {
		return
	}
	_ipConfigurationModeArg, _err := convertEnumIPConfigurationModeToXen(fmt.Sprintf("%s(%s)", _method, "ip_configuration_mode"), ipConfigurationMode)
	if _err != nil {
		return
	}
	_ipArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "IP"), ip)
	if _err != nil {
		return
	}
	_netmaskArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "netmask"), netmask)
	if _err != nil {
		return
	}
	_gatewayArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "gateway"), gateway)
	if _err != nil {
		return
	}
	_dnsArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "DNS"), dns)
	if _err != nil {
		return
	}
	_bondSlaveOfArg, _err := convertBondRefToXen(fmt.Sprintf("%s(%s)", _method, "bond_slave_of"), bondSlaveOf)
	if _err != nil {
		return
	}
	_vlanMasterOfArg, _err := convertVLANRefToXen(fmt.Sprintf("%s(%s)", _method, "VLAN_master_of"), vlanMasterOf)
	if _err != nil {
		return
	}
	_managementArg, _err := convertBoolToXen(fmt.Sprintf("%s(%s)", _method, "management"), management)
	if _err != nil {
		return
	}
	_otherConfigArg, _err := convertStringToStringMapToXen(fmt.Sprintf("%s(%s)", _method, "other_config"), otherConfig)
	if _err != nil {
		return
	}
	_disallowUnplugArg, _err := convertBoolToXen(fmt.Sprintf("%s(%s)", _method, "disallow_unplug"), disallowUnplug)
	if _err != nil {
		return
	}
	_ipv6ConfigurationModeArg, _err := convertEnumIpv6ConfigurationModeToXen(fmt.Sprintf("%s(%s)", _method, "ipv6_configuration_mode"), ipv6ConfigurationMode)
	if _err != nil {
		return
	}
	_ipv6Arg, _err := convertStringSetToXen(fmt.Sprintf("%s(%s)", _method, "IPv6"), ipv6)
	if _err != nil {
		return
	}
	_ipv6GatewayArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "ipv6_gateway"), ipv6Gateway)
	if _err != nil {
		return
	}
	_primaryAddressTypeArg, _err := convertEnumPrimaryAddressTypeToXen(fmt.Sprintf("%s(%s)", _method, "primary_address_type"), primaryAddressType)
	if _err != nil {
		return
	}
	_managedArg, _err := convertBoolToXen(fmt.Sprintf("%s(%s)", _method, "managed"), managed)
	if _err != nil {
		return
	}
	_propertiesArg, _err := convertStringToStringMapToXen(fmt.Sprintf("%s(%s)", _method, "properties"), properties)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _deviceArg, _networkArg, _hostArg, _macArg, _mtuArg, _vlanArg, _physicalArg, _ipConfigurationModeArg, _ipArg, _netmaskArg, _gatewayArg, _dnsArg, _bondSlaveOfArg, _vlanMasterOfArg, _managementArg, _otherConfigArg, _disallowUnplugArg, _ipv6ConfigurationModeArg, _ipv6Arg, _ipv6GatewayArg, _primaryAddressTypeArg, _managedArg, _propertiesArg)
	if _err != nil {
		return
	}
	_retval, _err = convertPIFRefToGo(_method + " -> ", _result.Value)
	return
}

func (_class PIFClass) Plug(sessionID SessionRef, self PIFRef) (_err error) {
	_method := "PIF.plug"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPIFRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg)
	return
}

func (_class PIFClass) Unplug(sessionID SessionRef, self PIFRef) (_err error) {
	_method := "PIF.unplug"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPIFRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg)
	return
}

func (_class PIFClass) Forget(sessionID SessionRef, self PIFRef) (_err error) {
	_method := "PIF.forget"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPIFRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg)
	return
}

func (_class PIFClass) Introduce(sessionID SessionRef, host HostRef, mac string, device string, managed bool) (_retval PIFRef, _err error) {
	_method := "PIF.introduce"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_hostArg, _err := convertHostRefToXen(fmt.Sprintf("%s(%s)", _method, "host"), host)
	if _err != nil {
		return
	}
	_macArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "MAC"), mac)
	if _err != nil {
		return
	}
	_deviceArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "device"), device)
	if _err != nil {
		return
	}
	_managedArg, _err := convertBoolToXen(fmt.Sprintf("%s(%s)", _method, "managed"), managed)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _hostArg, _macArg, _deviceArg, _managedArg)
	if _err != nil {
		return
	}
	_retval, _err = convertPIFRefToGo(_method + " -> ", _result.Value)
	return
}

func (_class PIFClass) Scan(sessionID SessionRef, host HostRef) (_err error) {
	_method := "PIF.scan"
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

func (_class PIFClass) SetPrimaryAddressType(sessionID SessionRef, self PIFRef, primaryAddressType PrimaryAddressType) (_err error) {
	_method := "PIF.set_primary_address_type"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPIFRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_primaryAddressTypeArg, _err := convertEnumPrimaryAddressTypeToXen(fmt.Sprintf("%s(%s)", _method, "primary_address_type"), primaryAddressType)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _primaryAddressTypeArg)
	return
}

func (_class PIFClass) ReconfigureIpv6(sessionID SessionRef, self PIFRef, mode Ipv6ConfigurationMode, ipv6 string, gateway string, dns string) (_err error) {
	_method := "PIF.reconfigure_ipv6"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPIFRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_modeArg, _err := convertEnumIpv6ConfigurationModeToXen(fmt.Sprintf("%s(%s)", _method, "mode"), mode)
	if _err != nil {
		return
	}
	_ipv6Arg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "IPv6"), ipv6)
	if _err != nil {
		return
	}
	_gatewayArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "gateway"), gateway)
	if _err != nil {
		return
	}
	_dnsArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "DNS"), dns)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _modeArg, _ipv6Arg, _gatewayArg, _dnsArg)
	return
}

func (_class PIFClass) ReconfigureIP(sessionID SessionRef, self PIFRef, mode IPConfigurationMode, ip string, netmask string, gateway string, dns string) (_err error) {
	_method := "PIF.reconfigure_ip"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPIFRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_modeArg, _err := convertEnumIPConfigurationModeToXen(fmt.Sprintf("%s(%s)", _method, "mode"), mode)
	if _err != nil {
		return
	}
	_ipArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "IP"), ip)
	if _err != nil {
		return
	}
	_netmaskArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "netmask"), netmask)
	if _err != nil {
		return
	}
	_gatewayArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "gateway"), gateway)
	if _err != nil {
		return
	}
	_dnsArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "DNS"), dns)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _modeArg, _ipArg, _netmaskArg, _gatewayArg, _dnsArg)
	return
}

func (_class PIFClass) Destroy(sessionID SessionRef, self PIFRef) (_err error) {
	_method := "PIF.destroy"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPIFRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg)
	return
}

func (_class PIFClass) CreateVLAN(sessionID SessionRef, device string, network NetworkRef, host HostRef, vlan int) (_retval PIFRef, _err error) {
	_method := "PIF.create_VLAN"
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
	_hostArg, _err := convertHostRefToXen(fmt.Sprintf("%s(%s)", _method, "host"), host)
	if _err != nil {
		return
	}
	_vlanArg, _err := convertIntToXen(fmt.Sprintf("%s(%s)", _method, "VLAN"), vlan)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _deviceArg, _networkArg, _hostArg, _vlanArg)
	if _err != nil {
		return
	}
	_retval, _err = convertPIFRefToGo(_method + " -> ", _result.Value)
	return
}

func (_class PIFClass) SetDisallowUnplug(sessionID SessionRef, self PIFRef, value bool) (_err error) {
	_method := "PIF.set_disallow_unplug"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPIFRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class PIFClass) RemoveFromOtherConfig(sessionID SessionRef, self PIFRef, key string) (_err error) {
	_method := "PIF.remove_from_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPIFRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class PIFClass) AddToOtherConfig(sessionID SessionRef, self PIFRef, key string, value string) (_err error) {
	_method := "PIF.add_to_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPIFRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class PIFClass) SetOtherConfig(sessionID SessionRef, self PIFRef, value map[string]string) (_err error) {
	_method := "PIF.set_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPIFRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class PIFClass) GetCapabilities(sessionID SessionRef, self PIFRef) (_retval []string, _err error) {
	_method := "PIF.get_capabilities"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPIFRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class PIFClass) GetProperties(sessionID SessionRef, self PIFRef) (_retval map[string]string, _err error) {
	_method := "PIF.get_properties"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPIFRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class PIFClass) GetManaged(sessionID SessionRef, self PIFRef) (_retval bool, _err error) {
	_method := "PIF.get_managed"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPIFRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class PIFClass) GetPrimaryAddressType(sessionID SessionRef, self PIFRef) (_retval PrimaryAddressType, _err error) {
	_method := "PIF.get_primary_address_type"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPIFRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertEnumPrimaryAddressTypeToGo(_method + " -> ", _result.Value)
	return
}

func (_class PIFClass) GetIpv6Gateway(sessionID SessionRef, self PIFRef) (_retval string, _err error) {
	_method := "PIF.get_ipv6_gateway"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPIFRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class PIFClass) GetIPv6(sessionID SessionRef, self PIFRef) (_retval []string, _err error) {
	_method := "PIF.get_IPv6"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPIFRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class PIFClass) GetIpv6ConfigurationMode(sessionID SessionRef, self PIFRef) (_retval Ipv6ConfigurationMode, _err error) {
	_method := "PIF.get_ipv6_configuration_mode"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPIFRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertEnumIpv6ConfigurationModeToGo(_method + " -> ", _result.Value)
	return
}

func (_class PIFClass) GetTunnelTransportPIFOf(sessionID SessionRef, self PIFRef) (_retval []TunnelRef, _err error) {
	_method := "PIF.get_tunnel_transport_PIF_of"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPIFRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTunnelRefSetToGo(_method + " -> ", _result.Value)
	return
}

func (_class PIFClass) GetTunnelAccessPIFOf(sessionID SessionRef, self PIFRef) (_retval []TunnelRef, _err error) {
	_method := "PIF.get_tunnel_access_PIF_of"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPIFRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTunnelRefSetToGo(_method + " -> ", _result.Value)
	return
}

func (_class PIFClass) GetDisallowUnplug(sessionID SessionRef, self PIFRef) (_retval bool, _err error) {
	_method := "PIF.get_disallow_unplug"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPIFRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class PIFClass) GetOtherConfig(sessionID SessionRef, self PIFRef) (_retval map[string]string, _err error) {
	_method := "PIF.get_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPIFRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class PIFClass) GetManagement(sessionID SessionRef, self PIFRef) (_retval bool, _err error) {
	_method := "PIF.get_management"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPIFRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class PIFClass) GetVLANSlaveOf(sessionID SessionRef, self PIFRef) (_retval []VLANRef, _err error) {
	_method := "PIF.get_VLAN_slave_of"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPIFRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVLANRefSetToGo(_method + " -> ", _result.Value)
	return
}

func (_class PIFClass) GetVLANMasterOf(sessionID SessionRef, self PIFRef) (_retval VLANRef, _err error) {
	_method := "PIF.get_VLAN_master_of"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPIFRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVLANRefToGo(_method + " -> ", _result.Value)
	return
}

func (_class PIFClass) GetBondMasterOf(sessionID SessionRef, self PIFRef) (_retval []BondRef, _err error) {
	_method := "PIF.get_bond_master_of"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPIFRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertBondRefSetToGo(_method + " -> ", _result.Value)
	return
}

func (_class PIFClass) GetBondSlaveOf(sessionID SessionRef, self PIFRef) (_retval BondRef, _err error) {
	_method := "PIF.get_bond_slave_of"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPIFRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertBondRefToGo(_method + " -> ", _result.Value)
	return
}

func (_class PIFClass) GetDNS(sessionID SessionRef, self PIFRef) (_retval string, _err error) {
	_method := "PIF.get_DNS"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPIFRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class PIFClass) GetGateway(sessionID SessionRef, self PIFRef) (_retval string, _err error) {
	_method := "PIF.get_gateway"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPIFRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class PIFClass) GetNetmask(sessionID SessionRef, self PIFRef) (_retval string, _err error) {
	_method := "PIF.get_netmask"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPIFRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class PIFClass) GetIP(sessionID SessionRef, self PIFRef) (_retval string, _err error) {
	_method := "PIF.get_IP"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPIFRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class PIFClass) GetIPConfigurationMode(sessionID SessionRef, self PIFRef) (_retval IPConfigurationMode, _err error) {
	_method := "PIF.get_ip_configuration_mode"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPIFRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertEnumIPConfigurationModeToGo(_method + " -> ", _result.Value)
	return
}

func (_class PIFClass) GetCurrentlyAttached(sessionID SessionRef, self PIFRef) (_retval bool, _err error) {
	_method := "PIF.get_currently_attached"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPIFRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class PIFClass) GetPhysical(sessionID SessionRef, self PIFRef) (_retval bool, _err error) {
	_method := "PIF.get_physical"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPIFRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class PIFClass) GetMetrics(sessionID SessionRef, self PIFRef) (_retval PIFMetricsRef, _err error) {
	_method := "PIF.get_metrics"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPIFRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertPIFMetricsRefToGo(_method + " -> ", _result.Value)
	return
}

func (_class PIFClass) GetVLAN(sessionID SessionRef, self PIFRef) (_retval int, _err error) {
	_method := "PIF.get_VLAN"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPIFRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class PIFClass) GetMTU(sessionID SessionRef, self PIFRef) (_retval int, _err error) {
	_method := "PIF.get_MTU"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPIFRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class PIFClass) GetMAC(sessionID SessionRef, self PIFRef) (_retval string, _err error) {
	_method := "PIF.get_MAC"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPIFRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class PIFClass) GetHost(sessionID SessionRef, self PIFRef) (_retval HostRef, _err error) {
	_method := "PIF.get_host"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPIFRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class PIFClass) GetNetwork(sessionID SessionRef, self PIFRef) (_retval NetworkRef, _err error) {
	_method := "PIF.get_network"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPIFRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertNetworkRefToGo(_method + " -> ", _result.Value)
	return
}

func (_class PIFClass) GetDevice(sessionID SessionRef, self PIFRef) (_retval string, _err error) {
	_method := "PIF.get_device"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPIFRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class PIFClass) GetUUID(sessionID SessionRef, self PIFRef) (_retval string, _err error) {
	_method := "PIF.get_uuid"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPIFRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

func (_class PIFClass) GetByUUID(sessionID SessionRef, uuid string) (_retval PIFRef, _err error) {
	_method := "PIF.get_by_uuid"
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
	_retval, _err = convertPIFRefToGo(_method + " -> ", _result.Value)
	return
}

func (_class PIFClass) GetRecord(sessionID SessionRef, self PIFRef) (_retval PIFRecord, _err error) {
	_method := "PIF.get_record"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPIFRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertPIFRecordToGo(_method + " -> ", _result.Value)
	return
}
