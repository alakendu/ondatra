/*
Package lacp is a generated package which contains definitions
of structs which generate gNMI paths for a YANG schema. The generated paths are
based on a compressed form of the schema.

This package was generated by /usr/local/google/home/gdennis/go/pkg/mod/github.com/openconfig/ygot@v0.23.1/genutil/names.go
using the following YANG input files:
  - gnmi-collector-metadata.yang
  - gnsi/authz/gnsi-authz.yang
  - gnsi/cert/gnsi-cert.yang
  - gnsi/console/gnsi-console.yang
  - gnsi/pathz/gnsi-pathz.yang
  - gnsi/ssh/gnsi-ssh.yang
  - public/release/models/acl/openconfig-acl.yang
  - public/release/models/acl/openconfig-packet-match.yang
  - public/release/models/aft/openconfig-aft.yang
  - public/release/models/aft/openconfig-aft-network-instance.yang
  - public/release/models/ate/openconfig-ate-flow.yang
  - public/release/models/ate/openconfig-ate-intf.yang
  - public/release/models/bfd/openconfig-bfd.yang
  - public/release/models/bgp/openconfig-bgp-policy.yang
  - public/release/models/bgp/openconfig-bgp-types.yang
  - public/release/models/interfaces/openconfig-if-aggregate.yang
  - public/release/models/interfaces/openconfig-if-ethernet.yang
  - public/release/models/interfaces/openconfig-if-ethernet-ext.yang
  - public/release/models/interfaces/openconfig-if-ip-ext.yang
  - public/release/models/interfaces/openconfig-if-ip.yang
  - public/release/models/interfaces/openconfig-interfaces.yang
  - public/release/models/isis/openconfig-isis.yang
  - public/release/models/lacp/openconfig-lacp.yang
  - public/release/models/lldp/openconfig-lldp-types.yang
  - public/release/models/lldp/openconfig-lldp.yang
  - public/release/models/local-routing/openconfig-local-routing.yang
  - public/release/models/mpls/openconfig-mpls-types.yang
  - public/release/models/multicast/openconfig-pim.yang
  - public/release/models/network-instance/openconfig-network-instance.yang
  - public/release/models/openconfig-extensions.yang
  - public/release/models/optical-transport/openconfig-transport-types.yang
  - public/release/models/ospf/openconfig-ospfv2.yang
  - public/release/models/p4rt/openconfig-p4rt.yang
  - public/release/models/platform/openconfig-platform-cpu.yang
  - public/release/models/platform/openconfig-platform-fan.yang
  - public/release/models/platform/openconfig-platform-integrated-circuit.yang
  - public/release/models/platform/openconfig-platform-software.yang
  - public/release/models/platform/openconfig-platform-transceiver.yang
  - public/release/models/platform/openconfig-platform.yang
  - public/release/models/policy-forwarding/openconfig-policy-forwarding.yang
  - public/release/models/policy/openconfig-policy-types.yang
  - public/release/models/qos/openconfig-qos-elements.yang
  - public/release/models/qos/openconfig-qos-interfaces.yang
  - public/release/models/qos/openconfig-qos-types.yang
  - public/release/models/qos/openconfig-qos.yang
  - public/release/models/rib/openconfig-rib-bgp.yang
  - public/release/models/sampling/openconfig-sampling-sflow.yang
  - public/release/models/segment-routing/openconfig-segment-routing-types.yang
  - public/release/models/system/openconfig-system.yang
  - public/release/models/types/openconfig-inet-types.yang
  - public/release/models/types/openconfig-types.yang
  - public/release/models/types/openconfig-yang-types.yang
  - public/release/models/vlan/openconfig-vlan.yang
  - public/third_party/ietf/iana-if-type.yang
  - public/third_party/ietf/ietf-inet-types.yang
  - public/third_party/ietf/ietf-interfaces.yang
  - public/third_party/ietf/ietf-yang-types.yang

Imported modules were sourced from:
  - public/release/models/...
  - public/third_party/ietf/...
*/
package lacp

import (
	"github.com/openconfig/ygot/ygot"
)

// LacpPath represents the /openconfig-lacp/lacp YANG schema element.
type LacpPath struct {
	*ygot.NodePath
}

// LacpPathAny represents the wildcard version of the /openconfig-lacp/lacp YANG schema element.
type LacpPathAny struct {
	*ygot.NodePath
}

// Lacp_SystemPriorityPath represents the /openconfig-lacp/lacp/config/system-priority YANG schema element.
type Lacp_SystemPriorityPath struct {
	*ygot.NodePath
}

// Lacp_SystemPriorityPathAny represents the wildcard version of the /openconfig-lacp/lacp/config/system-priority YANG schema element.
type Lacp_SystemPriorityPathAny struct {
	*ygot.NodePath
}

// InterfaceAny (list): List of aggregate interfaces managed by LACP
// ----------------------------------------
// Defining module: "openconfig-lacp"
// Instantiating module: "openconfig-lacp"
// Path from parent: "interfaces/interface"
// Path from root: "/lacp/interfaces/interface"
// Name (wildcarded): string
func (n *LacpPath) InterfaceAny() *Lacp_InterfacePathAny {
	return &Lacp_InterfacePathAny{
		NodePath: ygot.NewNodePath(
			[]string{"interfaces", "interface"},
			map[string]interface{}{"name": "*"},
			n,
		),
	}
}

// InterfaceAny (list): List of aggregate interfaces managed by LACP
// ----------------------------------------
// Defining module: "openconfig-lacp"
// Instantiating module: "openconfig-lacp"
// Path from parent: "interfaces/interface"
// Path from root: "/lacp/interfaces/interface"
// Name (wildcarded): string
func (n *LacpPathAny) InterfaceAny() *Lacp_InterfacePathAny {
	return &Lacp_InterfacePathAny{
		NodePath: ygot.NewNodePath(
			[]string{"interfaces", "interface"},
			map[string]interface{}{"name": "*"},
			n,
		),
	}
}

// Interface (list): List of aggregate interfaces managed by LACP
// ----------------------------------------
// Defining module: "openconfig-lacp"
// Instantiating module: "openconfig-lacp"
// Path from parent: "interfaces/interface"
// Path from root: "/lacp/interfaces/interface"
// Name: string
func (n *LacpPath) Interface(Name string) *Lacp_InterfacePath {
	return &Lacp_InterfacePath{
		NodePath: ygot.NewNodePath(
			[]string{"interfaces", "interface"},
			map[string]interface{}{"name": Name},
			n,
		),
	}
}

// Interface (list): List of aggregate interfaces managed by LACP
// ----------------------------------------
// Defining module: "openconfig-lacp"
// Instantiating module: "openconfig-lacp"
// Path from parent: "interfaces/interface"
// Path from root: "/lacp/interfaces/interface"
// Name: string
func (n *LacpPathAny) Interface(Name string) *Lacp_InterfacePathAny {
	return &Lacp_InterfacePathAny{
		NodePath: ygot.NewNodePath(
			[]string{"interfaces", "interface"},
			map[string]interface{}{"name": Name},
			n,
		),
	}
}

// SystemPriority (leaf): Sytem priority used by the node on this LAG interface.
// Lower value is higher priority for determining which node
// is the controlling system.
// ----------------------------------------
// Defining module: "openconfig-lacp"
// Instantiating module: "openconfig-lacp"
// Path from parent: "config/system-priority"
// Path from root: "/lacp/config/system-priority"
func (n *LacpPath) SystemPriority() *Lacp_SystemPriorityPath {
	return &Lacp_SystemPriorityPath{
		NodePath: ygot.NewNodePath(
			[]string{"config", "system-priority"},
			map[string]interface{}{},
			n,
		),
	}
}

// SystemPriority (leaf): Sytem priority used by the node on this LAG interface.
// Lower value is higher priority for determining which node
// is the controlling system.
// ----------------------------------------
// Defining module: "openconfig-lacp"
// Instantiating module: "openconfig-lacp"
// Path from parent: "config/system-priority"
// Path from root: "/lacp/config/system-priority"
func (n *LacpPathAny) SystemPriority() *Lacp_SystemPriorityPathAny {
	return &Lacp_SystemPriorityPathAny{
		NodePath: ygot.NewNodePath(
			[]string{"config", "system-priority"},
			map[string]interface{}{},
			n,
		),
	}
}

// Lacp_InterfacePath represents the /openconfig-lacp/lacp/interfaces/interface YANG schema element.
type Lacp_InterfacePath struct {
	*ygot.NodePath
}

// Lacp_InterfacePathAny represents the wildcard version of the /openconfig-lacp/lacp/interfaces/interface YANG schema element.
type Lacp_InterfacePathAny struct {
	*ygot.NodePath
}

// Lacp_Interface_IntervalPath represents the /openconfig-lacp/lacp/interfaces/interface/config/interval YANG schema element.
type Lacp_Interface_IntervalPath struct {
	*ygot.NodePath
}

// Lacp_Interface_IntervalPathAny represents the wildcard version of the /openconfig-lacp/lacp/interfaces/interface/config/interval YANG schema element.
type Lacp_Interface_IntervalPathAny struct {
	*ygot.NodePath
}

// Lacp_Interface_LacpModePath represents the /openconfig-lacp/lacp/interfaces/interface/config/lacp-mode YANG schema element.
type Lacp_Interface_LacpModePath struct {
	*ygot.NodePath
}

// Lacp_Interface_LacpModePathAny represents the wildcard version of the /openconfig-lacp/lacp/interfaces/interface/config/lacp-mode YANG schema element.
type Lacp_Interface_LacpModePathAny struct {
	*ygot.NodePath
}

// Lacp_Interface_NamePath represents the /openconfig-lacp/lacp/interfaces/interface/config/name YANG schema element.
type Lacp_Interface_NamePath struct {
	*ygot.NodePath
}

// Lacp_Interface_NamePathAny represents the wildcard version of the /openconfig-lacp/lacp/interfaces/interface/config/name YANG schema element.
type Lacp_Interface_NamePathAny struct {
	*ygot.NodePath
}

// Lacp_Interface_SystemIdMacPath represents the /openconfig-lacp/lacp/interfaces/interface/config/system-id-mac YANG schema element.
type Lacp_Interface_SystemIdMacPath struct {
	*ygot.NodePath
}

// Lacp_Interface_SystemIdMacPathAny represents the wildcard version of the /openconfig-lacp/lacp/interfaces/interface/config/system-id-mac YANG schema element.
type Lacp_Interface_SystemIdMacPathAny struct {
	*ygot.NodePath
}

// Lacp_Interface_SystemPriorityPath represents the /openconfig-lacp/lacp/interfaces/interface/config/system-priority YANG schema element.
type Lacp_Interface_SystemPriorityPath struct {
	*ygot.NodePath
}

// Lacp_Interface_SystemPriorityPathAny represents the wildcard version of the /openconfig-lacp/lacp/interfaces/interface/config/system-priority YANG schema element.
type Lacp_Interface_SystemPriorityPathAny struct {
	*ygot.NodePath
}

// Interval (leaf): Set the period between LACP messages -- uses
// the lacp-period-type enumeration.
// ----------------------------------------
// Defining module: "openconfig-lacp"
// Instantiating module: "openconfig-lacp"
// Path from parent: "config/interval"
// Path from root: "/lacp/interfaces/interface/config/interval"
func (n *Lacp_InterfacePath) Interval() *Lacp_Interface_IntervalPath {
	return &Lacp_Interface_IntervalPath{
		NodePath: ygot.NewNodePath(
			[]string{"config", "interval"},
			map[string]interface{}{},
			n,
		),
	}
}

// Interval (leaf): Set the period between LACP messages -- uses
// the lacp-period-type enumeration.
// ----------------------------------------
// Defining module: "openconfig-lacp"
// Instantiating module: "openconfig-lacp"
// Path from parent: "config/interval"
// Path from root: "/lacp/interfaces/interface/config/interval"
func (n *Lacp_InterfacePathAny) Interval() *Lacp_Interface_IntervalPathAny {
	return &Lacp_Interface_IntervalPathAny{
		NodePath: ygot.NewNodePath(
			[]string{"config", "interval"},
			map[string]interface{}{},
			n,
		),
	}
}

// LacpMode (leaf): ACTIVE is to initiate the transmission of LACP packets.
// PASSIVE is to wait for peer to initiate the transmission of
// LACP packets.
// ----------------------------------------
// Defining module: "openconfig-lacp"
// Instantiating module: "openconfig-lacp"
// Path from parent: "config/lacp-mode"
// Path from root: "/lacp/interfaces/interface/config/lacp-mode"
func (n *Lacp_InterfacePath) LacpMode() *Lacp_Interface_LacpModePath {
	return &Lacp_Interface_LacpModePath{
		NodePath: ygot.NewNodePath(
			[]string{"config", "lacp-mode"},
			map[string]interface{}{},
			n,
		),
	}
}

// LacpMode (leaf): ACTIVE is to initiate the transmission of LACP packets.
// PASSIVE is to wait for peer to initiate the transmission of
// LACP packets.
// ----------------------------------------
// Defining module: "openconfig-lacp"
// Instantiating module: "openconfig-lacp"
// Path from parent: "config/lacp-mode"
// Path from root: "/lacp/interfaces/interface/config/lacp-mode"
func (n *Lacp_InterfacePathAny) LacpMode() *Lacp_Interface_LacpModePathAny {
	return &Lacp_Interface_LacpModePathAny{
		NodePath: ygot.NewNodePath(
			[]string{"config", "lacp-mode"},
			map[string]interface{}{},
			n,
		),
	}
}

// Name (leaf): Reference to the interface on which LACP should be
// configured.   The type of the target interface must be
// ieee8023adLag
// ----------------------------------------
// Defining module: "openconfig-lacp"
// Instantiating module: "openconfig-lacp"
// Path from parent: "config/name"
// Path from root: "/lacp/interfaces/interface/config/name"
func (n *Lacp_InterfacePath) Name() *Lacp_Interface_NamePath {
	return &Lacp_Interface_NamePath{
		NodePath: ygot.NewNodePath(
			[]string{"config", "name"},
			map[string]interface{}{},
			n,
		),
	}
}

// Name (leaf): Reference to the interface on which LACP should be
// configured.   The type of the target interface must be
// ieee8023adLag
// ----------------------------------------
// Defining module: "openconfig-lacp"
// Instantiating module: "openconfig-lacp"
// Path from parent: "config/name"
// Path from root: "/lacp/interfaces/interface/config/name"
func (n *Lacp_InterfacePathAny) Name() *Lacp_Interface_NamePathAny {
	return &Lacp_Interface_NamePathAny{
		NodePath: ygot.NewNodePath(
			[]string{"config", "name"},
			map[string]interface{}{},
			n,
		),
	}
}

// SystemIdMac (leaf): The MAC address portion of the node's System ID. This is
// combined with the system priority to construct the 8-octet
// system-id
// ----------------------------------------
// Defining module: "openconfig-lacp"
// Instantiating module: "openconfig-lacp"
// Path from parent: "config/system-id-mac"
// Path from root: "/lacp/interfaces/interface/config/system-id-mac"
func (n *Lacp_InterfacePath) SystemIdMac() *Lacp_Interface_SystemIdMacPath {
	return &Lacp_Interface_SystemIdMacPath{
		NodePath: ygot.NewNodePath(
			[]string{"config", "system-id-mac"},
			map[string]interface{}{},
			n,
		),
	}
}

// SystemIdMac (leaf): The MAC address portion of the node's System ID. This is
// combined with the system priority to construct the 8-octet
// system-id
// ----------------------------------------
// Defining module: "openconfig-lacp"
// Instantiating module: "openconfig-lacp"
// Path from parent: "config/system-id-mac"
// Path from root: "/lacp/interfaces/interface/config/system-id-mac"
func (n *Lacp_InterfacePathAny) SystemIdMac() *Lacp_Interface_SystemIdMacPathAny {
	return &Lacp_Interface_SystemIdMacPathAny{
		NodePath: ygot.NewNodePath(
			[]string{"config", "system-id-mac"},
			map[string]interface{}{},
			n,
		),
	}
}

// SystemPriority (leaf): Sytem priority used by the node on this LAG interface.
// Lower value is higher priority for determining which node
// is the controlling system.
// ----------------------------------------
// Defining module: "openconfig-lacp"
// Instantiating module: "openconfig-lacp"
// Path from parent: "config/system-priority"
// Path from root: "/lacp/interfaces/interface/config/system-priority"
func (n *Lacp_InterfacePath) SystemPriority() *Lacp_Interface_SystemPriorityPath {
	return &Lacp_Interface_SystemPriorityPath{
		NodePath: ygot.NewNodePath(
			[]string{"config", "system-priority"},
			map[string]interface{}{},
			n,
		),
	}
}

// SystemPriority (leaf): Sytem priority used by the node on this LAG interface.
// Lower value is higher priority for determining which node
// is the controlling system.
// ----------------------------------------
// Defining module: "openconfig-lacp"
// Instantiating module: "openconfig-lacp"
// Path from parent: "config/system-priority"
// Path from root: "/lacp/interfaces/interface/config/system-priority"
func (n *Lacp_InterfacePathAny) SystemPriority() *Lacp_Interface_SystemPriorityPathAny {
	return &Lacp_Interface_SystemPriorityPathAny{
		NodePath: ygot.NewNodePath(
			[]string{"config", "system-priority"},
			map[string]interface{}{},
			n,
		),
	}
}
