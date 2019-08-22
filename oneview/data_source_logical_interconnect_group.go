// (C) Copyright 2018 Hewlett Packard Enterprise Development LP
//
// Licensed under the Apache License, Version 2.0 (the "License");
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software distributed
// under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR
// CONDITIONS OF ANY KIND, either express or implied. See the License for the
// specific language governing permissions and limitations under the License.

package oneview

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"strconv"
)

func dataSourceLogicalInterconnectGroup() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceLogicalInterconnectGroupRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"interconnect_bay_set": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"redundancy_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"enclosure_indexes": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeInt},
				Set: func(a interface{}) int {
					return a.(int)
				},
			},
			"interconnect_map_entry_template": {
				Computed: true,
				Type:     schema.TypeSet,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"bay_number": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"interconnect_type_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"enclosure_index": {
							Type:     schema.TypeInt,
							Computed: true,
						},
					},
				},
			},
			"uplink_set": {
				Computed: true,
				Type:     schema.TypeList,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"network_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"ethernet_network_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"logical_port_config": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"desired_speed": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"port_num": {
										Type:     schema.TypeSet,
										Computed: true,
										Elem:     &schema.Schema{Type: schema.TypeInt},
										Set: func(a interface{}) int {
											return a.(int)
										},
									},
									"bay_num": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"enclosure_num": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"primary_port": {
										Type:     schema.TypeBool,
										Computed: true,
									},
								},
							},
						},
						"mode": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"network_uris": {
							Type:     schema.TypeSet,
							Computed: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Set:      schema.HashString,
						},
						"lacp_timer": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"native_network_uri": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"internal_network_uris": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Set:      schema.HashString,
			},
			"telemetry_configuration": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"enabled": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"sample_count": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"sample_interval": {
							Type:     schema.TypeInt,
							Computed: true,
						},
					},
				},
			},
			"snmp_configuration": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enabled": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"v3_enabled": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"read_community": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"system_contact": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"snmp_access": {
							Type:     schema.TypeSet,
							Computed: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Set:      schema.HashString,
						},
						"trap_destination": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"community_string": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"enet_trap_categories": {
										Type:     schema.TypeSet,
										Computed: true,
										Elem:     &schema.Schema{Type: schema.TypeString},
										Set:      schema.HashString,
									},
									"fc_trap_categories": {
										Type:     schema.TypeSet,
										Computed: true,
										Elem:     &schema.Schema{Type: schema.TypeString},
										Set:      schema.HashString,
									},
									"vcm_trap_categories": {
										Type:     schema.TypeSet,
										Computed: true,
										Elem:     &schema.Schema{Type: schema.TypeString},
										Set:      schema.HashString,
									},
									"trap_destination": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"trap_format": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"trap_severities": {
										Type:     schema.TypeSet,
										Computed: true,
										Elem:     &schema.Schema{Type: schema.TypeString},
										Set:      schema.HashString,
									},
								},
							},
						},
					},
				},
			},
			"interconnect_settings": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"fast_mac_cache_failover": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"igmp_snooping": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"network_loop_protection": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"pause_flood_protection": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"rich_tlv": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"igmp_timeout_interval": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"mac_refresh_interval": {
							Type:     schema.TypeInt,
							Computed: true,
						},
					},
				},
			},
			"quality_of_service": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"active_qos_config_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"config_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"uplink_classification_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"downlink_classification_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"qos_traffic_class": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"enabled": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"egress_dot1p_value": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"real_time": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"bandwidth_share": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"max_bandwidth": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"qos_classification_map": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"dot1p_class_map": {
													Type:     schema.TypeSet,
													Computed: true,
													Elem:     &schema.Schema{Type: schema.TypeInt},
													Set: func(a interface{}) int {
														return a.(int)
													},
												},
												"dscp_class_map": {
													Type:     schema.TypeSet,
													Computed: true,
													Elem:     &schema.Schema{Type: schema.TypeString},
													Set:      schema.HashString,
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
			"created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"modified": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"uri": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"category": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"fabric_uri": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"eTag": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceLogicalInterconnectGroupRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	id := d.Get("name").(string)

	logicalInterconnectGroup, err := config.ovClient.GetLogicalInterconnectGroupByName(id)
	if err != nil || logicalInterconnectGroup.URI.IsNil() {
		d.SetId("")
		return nil
	}
	d.SetId(id)
	d.Set("name", logicalInterconnectGroup.Name)
	d.Set("type", logicalInterconnectGroup.Type)
	d.Set("created", logicalInterconnectGroup.Created)
	d.Set("modified", logicalInterconnectGroup.Modified)
	d.Set("uri", logicalInterconnectGroup.URI.String())
	d.Set("status", logicalInterconnectGroup.Status)
	d.Set("category", logicalInterconnectGroup.Category)
	d.Set("state", logicalInterconnectGroup.State)
	d.Set("fabric_uri", logicalInterconnectGroup.FabricUri.String())
	d.Set("eTag", logicalInterconnectGroup.ETAG)
	d.Set("description", logicalInterconnectGroup.Description)
	d.Set("interconnect_settings.0.igmp_snooping", logicalInterconnectGroup.EthernetSettings.EnableIgmpSnooping)
	d.Set("interconnect_bay_set", logicalInterconnectGroup.InterconnectBaySet)
	d.Set("redundancy_type", logicalInterconnectGroup.RedundancyType)

	enclosureIndexes := make([]interface{}, len(logicalInterconnectGroup.EnclosureIndexes))
	for i, enclosureIndexVal := range logicalInterconnectGroup.EnclosureIndexes {
		enclosureIndexes[i] = enclosureIndexVal
	}
	d.Set("enclosure_indexes", schema.NewSet(func(a interface{}) int { return a.(int) }, enclosureIndexes))

	interconnectMapEntryTemplates := make([]map[string]interface{}, 0, len(logicalInterconnectGroup.InterconnectMapTemplate.InterconnectMapEntryTemplates))
	for _, interconnectMapEntryTemplate := range logicalInterconnectGroup.InterconnectMapTemplate.InterconnectMapEntryTemplates {
		interconnectType, err := config.ovClient.GetInterconnectTypeByUri(interconnectMapEntryTemplate.PermittedInterconnectTypeUri)
		if err != nil {
			return err
		}
		if interconnectType.Name == "" {
			return fmt.Errorf("Could not find interconnectType with URI %s", interconnectMapEntryTemplate.PermittedInterconnectTypeUri.String())
		}
		var bayNum int
		var enclosureIndex int
		if interconnectMapEntryTemplate.LogicalLocation.LocationEntries[0].Type == "Bay" {
			bayNum = interconnectMapEntryTemplate.LogicalLocation.LocationEntries[0].RelativeValue
			enclosureIndex = interconnectMapEntryTemplate.LogicalLocation.LocationEntries[1].RelativeValue
		} else {
			bayNum = interconnectMapEntryTemplate.LogicalLocation.LocationEntries[1].RelativeValue
			enclosureIndex = interconnectMapEntryTemplate.LogicalLocation.LocationEntries[0].RelativeValue
		}

		interconnectMapEntryTemplates = append(interconnectMapEntryTemplates, map[string]interface{}{
			"interconnect_type_name": interconnectType.Name,
			"bay_number":             bayNum,
			"enclosure_index":        enclosureIndex,
		})
	}

	d.Set("interconnect_map_entry_template", interconnectMapEntryTemplates)

	uplinkSets := make([]map[string]interface{}, 0, len(logicalInterconnectGroup.UplinkSets))
	for i, uplinkSet := range logicalInterconnectGroup.UplinkSets {

		primaryPortEnclosure := 0
		primaryPortBay := 0
		primaryPortPort := 0

		if uplinkSet.PrimaryPort != nil {
			for _, primaryPortLocation := range uplinkSet.PrimaryPort.LocationEntries {
				if primaryPortLocation.Type == "Bay" {
					primaryPortBay = primaryPortLocation.RelativeValue
				}
				if primaryPortLocation.Type == "Enclosure" {
					primaryPortEnclosure = primaryPortLocation.RelativeValue
				}
				if primaryPortLocation.Type == "Port" {
					primaryPortPort = primaryPortLocation.RelativeValue
				}
			}
		}

		logicalPortConfigs := make([]map[string]interface{}, 0, len(uplinkSet.LogicalPortConfigInfos))
		for _, logicalPortConfigInfo := range uplinkSet.LogicalPortConfigInfos {
			portEnclosure := 0
			portBay := 0
			portPort := 0
			primaryPort := false
			for _, portLocation := range logicalPortConfigInfo.LogicalLocation.LocationEntries {
				if portLocation.Type == "Bay" {
					portBay = portLocation.RelativeValue
				}
				if portLocation.Type == "Enclosure" {
					portEnclosure = portLocation.RelativeValue
				}
				if portLocation.Type == "Port" {
					portPort = portLocation.RelativeValue
				}
			}
			if primaryPortEnclosure == portEnclosure && primaryPortBay == portBay && primaryPortPort == portPort {
				primaryPort = true
			}

			portPorts := make([]interface{}, 0)
			portPorts = append(portPorts, portPort)

			included := false
			for j, portConfig := range logicalPortConfigs {
				if portConfig["bay_num"] == portBay && portConfig["enclosure_num"] == portEnclosure {
					included = true
					portSet := logicalPortConfigs[j]["port_num"].(*schema.Set)
					portSet.Add(portPort)
				}
			}

			if included == false {
				logicalPortConfigs = append(logicalPortConfigs, map[string]interface{}{
					"desired_speed": logicalPortConfigInfo.DesiredSpeed,
					"primary_port":  primaryPort,
					"port_num":      schema.NewSet(func(a interface{}) int { return a.(int) }, portPorts),
					"bay_num":       portBay,
					"enclosure_num": portEnclosure,
				})
			}
		}

		//Oneview returns an unordered list so order it to match the configuration file
		logicalPortCount := d.Get("uplink_set." + strconv.Itoa(i) + ".logical_port_config.#").(int)
		oneviewLogicalPortCount := len(logicalPortConfigs)
		for j := 0; j < logicalPortCount; j++ {
			currBay := d.Get("uplink_set." + strconv.Itoa(i) + ".logical_port_config." + strconv.Itoa(j) + ".bay_num").(int)
			for k := 0; k < oneviewLogicalPortCount; k++ {
				if currBay == logicalPortConfigs[k]["bay_num"] && j <= k {
					logicalPortConfigs[j], logicalPortConfigs[k] = logicalPortConfigs[k], logicalPortConfigs[j]
				}
			}
		}

		networkUris := make([]interface{}, len(uplinkSet.NetworkUris))
		for i, networkUri := range uplinkSet.NetworkUris {
			networkUris[i] = networkUri.String()
		}

		uplinkSets = append(uplinkSets, map[string]interface{}{
			"network_type":          uplinkSet.NetworkType,
			"ethernet_network_type": uplinkSet.EthernetNetworkType,
			"name":                  uplinkSet.Name,
			"mode":                  uplinkSet.Mode,
			"lacp_timer":            uplinkSet.LacpTimer,
			"native_network_uri":    uplinkSet.NativeNetworkUri,
			"logical_port_config":   logicalPortConfigs,
			"network_uris":          schema.NewSet(schema.HashString, networkUris),
		})
	}
	uplinkCount := d.Get("uplink_set.#").(int)
	oneviewUplinkCount := len(uplinkSets)
	for i := 0; i < uplinkCount; i++ {
		currUplinkName := d.Get("uplink_set." + strconv.Itoa(i) + ".name").(string)
		for j := 0; j < oneviewUplinkCount; j++ {
			if currUplinkName == uplinkSets[j]["name"] && i <= j {
				uplinkSets[i], uplinkSets[j] = uplinkSets[j], uplinkSets[i]
			}
		}
	}
	d.Set("uplink_set", uplinkSets)

	internalNetworkUris := make([]interface{}, len(logicalInterconnectGroup.InternalNetworkUris))
	for i, internalNetworkUri := range logicalInterconnectGroup.InternalNetworkUris {
		internalNetworkUris[i] = internalNetworkUri
	}
	d.Set("internal_network_uris", internalNetworkUris)

	telemetryConfigurations := make([]map[string]interface{}, 0, 1)
	telemetryConfigurations = append(telemetryConfigurations, map[string]interface{}{
		"enabled":         *logicalInterconnectGroup.TelemetryConfiguration.EnableTelemetry,
		"sample_count":    logicalInterconnectGroup.TelemetryConfiguration.SampleCount,
		"sample_interval": logicalInterconnectGroup.TelemetryConfiguration.SampleInterval,
		"type":            logicalInterconnectGroup.TelemetryConfiguration.Type,
	})
	d.Set("telemetry_configuration", telemetryConfigurations)

	trapDestinations := make([]map[string]interface{}, 0, 1)
	for _, trapDestination := range logicalInterconnectGroup.SnmpConfiguration.TrapDestinations {

		enetTrapCategories := make([]interface{}, len(trapDestination.EnetTrapCategories))
		for i, enetTrapCategory := range trapDestination.EnetTrapCategories {
			enetTrapCategories[i] = enetTrapCategory
		}

		fcTrapCategories := make([]interface{}, len(trapDestination.FcTrapCategories))
		for i, fcTrapCategory := range trapDestination.FcTrapCategories {
			fcTrapCategories[i] = fcTrapCategory
		}

		vcmTrapCategories := make([]interface{}, len(trapDestination.VcmTrapCategories))
		for i, vcmTrapCategory := range trapDestination.VcmTrapCategories {
			vcmTrapCategories[i] = vcmTrapCategory
		}

		trapSeverities := make([]interface{}, len(trapDestination.TrapSeverities))
		for i, trapSeverity := range trapDestination.TrapSeverities {
			trapSeverities[i] = trapSeverity
		}

		trapDestinations = append(trapDestinations, map[string]interface{}{
			"trap_destination":     trapDestination.TrapDestination,
			"community_string":     trapDestination.CommunityString,
			"trap_format":          trapDestination.TrapFormat,
			"enet_trap_categories": schema.NewSet(schema.HashString, enetTrapCategories),
			"fc_trap_categories":   schema.NewSet(schema.HashString, fcTrapCategories),
			"vcm_trap_categories":  schema.NewSet(schema.HashString, vcmTrapCategories),
			"trap_severities":      schema.NewSet(schema.HashString, trapSeverities),
		})
	}

	//Oneview returns an unordered list so order it to match the configuration file
	trapDestinationCount := d.Get("snmp_configuration.0.trap_destination.#").(int)
	oneviewTrapDestinationCount := len(trapDestinations)
	for i := 0; i < trapDestinationCount; i++ {
		currDest := d.Get("snmp_configuration.0.trap_destination." + strconv.Itoa(i) + ".trap_destination").(string)
		for j := 0; j < oneviewTrapDestinationCount; j++ {
			if currDest == trapDestinations[j]["trap_destination"] && i <= j {
				trapDestinations[i], trapDestinations[j] = trapDestinations[j], trapDestinations[i]
			}
		}
	}

	snmpAccess := make([]interface{}, len(logicalInterconnectGroup.SnmpConfiguration.SnmpAccess))
	for i, snmpAccessIP := range logicalInterconnectGroup.SnmpConfiguration.SnmpAccess {
		snmpAccess[i] = snmpAccessIP
	}

	snmpConfiguration := make([]map[string]interface{}, 0, 1)
	snmpConfiguration = append(snmpConfiguration, map[string]interface{}{
		"enabled":          *logicalInterconnectGroup.SnmpConfiguration.Enabled,
		"v3_enabled":       *logicalInterconnectGroup.SnmpConfiguration.V3Enabled,
		"read_community":   logicalInterconnectGroup.SnmpConfiguration.ReadCommunity,
		"snmp_access":      schema.NewSet(schema.HashString, snmpAccess),
		"system_contact":   logicalInterconnectGroup.SnmpConfiguration.SystemContact,
		"type":             logicalInterconnectGroup.SnmpConfiguration.Type,
		"trap_destination": trapDestinations,
	})
	d.Set("snmp_configuration", snmpConfiguration)

	interconnectSettings := make([]map[string]interface{}, 0, 1)
	interconnectSettings = append(interconnectSettings, map[string]interface{}{
		"type": logicalInterconnectGroup.EthernetSettings.Type,
		"fast_mac_cache_failover": *logicalInterconnectGroup.EthernetSettings.EnableFastMacCacheFailover,
		"igmp_snooping":           *logicalInterconnectGroup.EthernetSettings.EnableIgmpSnooping,
		"network_loop_protection": *logicalInterconnectGroup.EthernetSettings.EnableNetworkLoopProtection,
		"pause_flood_protection":  *logicalInterconnectGroup.EthernetSettings.EnablePauseFloodProtection,
		"rich_tlv":                *logicalInterconnectGroup.EthernetSettings.EnableRichTLV,
		"igmp_timeout_interval":   logicalInterconnectGroup.EthernetSettings.IgmpIdleTimeoutInterval,
		"mac_refresh_interval":    logicalInterconnectGroup.EthernetSettings.MacRefreshInterval,
	})
	d.Set("interconnect_settings", interconnectSettings)

	qosTrafficClasses := make([]map[string]interface{}, 0, 1)
	for _, qosTrafficClass := range logicalInterconnectGroup.QosConfiguration.ActiveQosConfig.QosTrafficClassifiers {

		dscpClassMap := make([]interface{}, len(qosTrafficClass.QosClassificationMapping.DscpClassMapping))
		for i, dscpValue := range qosTrafficClass.QosClassificationMapping.DscpClassMapping {
			dscpClassMap[i] = dscpValue
		}

		dot1pClassMap := make([]interface{}, len(qosTrafficClass.QosClassificationMapping.Dot1pClassMapping))
		for i, dot1pValue := range qosTrafficClass.QosClassificationMapping.Dot1pClassMapping {
			dot1pClassMap[i] = dot1pValue
		}
		qosClassificationMap := make([]map[string]interface{}, 0, 1)
		qosClassificationMap = append(qosClassificationMap, map[string]interface{}{
			"dot1p_class_map": schema.NewSet(func(a interface{}) int { return a.(int) }, dot1pClassMap),
			"dscp_class_map":  schema.NewSet(schema.HashString, dscpClassMap),
		})

		qosTrafficClasses = append(qosTrafficClasses, map[string]interface{}{
			"name":                   qosTrafficClass.QosTrafficClass.ClassName,
			"enabled":                *qosTrafficClass.QosTrafficClass.Enabled,
			"egress_dot1p_value":     qosTrafficClass.QosTrafficClass.EgressDot1pValue,
			"real_time":              *qosTrafficClass.QosTrafficClass.RealTime,
			"bandwidth_share":        qosTrafficClass.QosTrafficClass.BandwidthShare,
			"max_bandwidth":          qosTrafficClass.QosTrafficClass.MaxBandwidth,
			"qos_classification_map": qosClassificationMap,
		})
	}
	qosTrafficClassCount := d.Get("quality_of_service.0.qos_traffic_class.#").(int)
	oneviewTrafficClassCount := len(qosTrafficClasses)
	for i := 0; i < qosTrafficClassCount; i++ {
		currName := d.Get("quality_of_service.0.qos_traffic_class." + strconv.Itoa(i) + ".name").(string)
		for j := 0; j < oneviewTrafficClassCount; j++ {
			if currName == qosTrafficClasses[j]["name"] && i <= j {
				qosTrafficClasses[i], qosTrafficClasses[j] = qosTrafficClasses[j], qosTrafficClasses[i]
			}
		}
	}

	qualityOfService := make([]map[string]interface{}, 0, 1)
	qualityOfService = append(qualityOfService, map[string]interface{}{
		"type": logicalInterconnectGroup.QosConfiguration.Type,
		"active_qos_config_type":       logicalInterconnectGroup.QosConfiguration.ActiveQosConfig.Type,
		"config_type":                  logicalInterconnectGroup.QosConfiguration.ActiveQosConfig.ConfigType,
		"uplink_classification_type":   logicalInterconnectGroup.QosConfiguration.ActiveQosConfig.UplinkClassificationType,
		"downlink_classification_type": logicalInterconnectGroup.QosConfiguration.ActiveQosConfig.DownlinkClassificationType,
		"qos_traffic_class":            qosTrafficClasses,
	})

	d.Set("quality_of_service", qualityOfService)

	return nil
}