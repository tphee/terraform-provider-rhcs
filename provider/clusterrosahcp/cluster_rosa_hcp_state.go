/*
Copyright (c) 2021 Red Hat, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package clusterrosahcp

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/terraform-redhat/terraform-provider-rhcs/provider/proxy"
)

type ClusterRosaHcpState struct {
	APIURL                               types.String       `tfsdk:"api_url"`
	AWSAccountID                         types.String       `tfsdk:"aws_account_id"`
	BillingAccountID                     types.String       `tfsdk:"billing_account_id"`
	AWSSubnetIDs                         types.List         `tfsdk:"aws_subnet_ids"`
	AWSAdditionalComputeSecurityGroupIds types.List         `tfsdk:"aws_additional_compute_security_group_ids"`
	AWSPrivateLink                       types.Bool         `tfsdk:"aws_private_link"`
	Private                              types.Bool         `tfsdk:"private"`
	Sts                                  *Sts               `tfsdk:"sts"`
	CCSEnabled                           types.Bool         `tfsdk:"ccs_enabled"`
	HypershiftEnabled                    types.Bool         `tfsdk:"hypershift_enabled"`
	EtcdEncryption                       types.Bool         `tfsdk:"etcd_encryption"`
	ChannelGroup                         types.String       `tfsdk:"channel_group"`
	CloudRegion                          types.String       `tfsdk:"cloud_region"`
	ComputeMachineType                   types.String       `tfsdk:"compute_machine_type"`
	WorkerDiskSize                       types.Int64        `tfsdk:"worker_disk_size"`
	DefaultMPLabels                      types.Map          `tfsdk:"default_mp_labels"`
	Replicas                             types.Int64        `tfsdk:"replicas"`
	ConsoleURL                           types.String       `tfsdk:"console_url"`
	Domain                               types.String       `tfsdk:"domain"`
	InfraID                              types.String       `tfsdk:"infra_id"`
	HostPrefix                           types.Int64        `tfsdk:"host_prefix"`
	ID                                   types.String       `tfsdk:"id"`
	FIPS                                 types.Bool         `tfsdk:"fips"`
	KMSKeyArn                            types.String       `tfsdk:"kms_key_arn"`
	ExternalID                           types.String       `tfsdk:"external_id"`
	MachineCIDR                          types.String       `tfsdk:"machine_cidr"`
	MultiAZ                              types.Bool         `tfsdk:"multi_az"`
	DisableWorkloadMonitoring            types.Bool         `tfsdk:"disable_workload_monitoring"`
	DisableSCPChecks                     types.Bool         `tfsdk:"disable_scp_checks"`
	AvailabilityZones                    types.List         `tfsdk:"availability_zones"`
	Name                                 types.String       `tfsdk:"name"`
	PodCIDR                              types.String       `tfsdk:"pod_cidr"`
	Properties                           types.Map          `tfsdk:"properties"`
	OCMProperties                        types.Map          `tfsdk:"ocm_properties"`
	Tags                                 types.Map          `tfsdk:"tags"`
	ServiceCIDR                          types.String       `tfsdk:"service_cidr"`
	Proxy                                *proxy.Proxy       `tfsdk:"proxy"`
	State                                types.String       `tfsdk:"state"`
	Version                              types.String       `tfsdk:"version"`
	CurrentVersion                       types.String       `tfsdk:"current_version"`
	DisableWaitingInDestroy              types.Bool         `tfsdk:"disable_waiting_in_destroy"`
	DestroyTimeout                       types.Int64        `tfsdk:"destroy_timeout"`
	UpgradeAcksFor                       types.String       `tfsdk:"upgrade_acknowledgements_for"`
	AdminCredentials                     *AdminCredentials  `tfsdk:"admin_credentials"`
	PrivateHostedZone                    *PrivateHostedZone `tfsdk:"private_hosted_zone"`
	BaseDNSDomain                        types.String       `tfsdk:"base_dns_domain"`
	WaitForCreateComplete                types.Bool         `tfsdk:"wait_for_create_complete"`
}

type Sts struct {
	OIDCEndpointURL    types.String    `tfsdk:"oidc_endpoint_url"`
	OIDCConfigID       types.String    `tfsdk:"oidc_config_id"`
	Thumbprint         types.String    `tfsdk:"thumbprint"`
	RoleARN            types.String    `tfsdk:"role_arn"`
	SupportRoleArn     types.String    `tfsdk:"support_role_arn"`
	InstanceIAMRoles   InstanceIAMRole `tfsdk:"instance_iam_roles"`
	OperatorRolePrefix types.String    `tfsdk:"operator_role_prefix"`
}

type InstanceIAMRole struct {
	MasterRoleARN types.String `tfsdk:"master_role_arn"`
	WorkerRoleARN types.String `tfsdk:"worker_role_arn"`
}

type AdminCredentials struct {
	Username types.String `tfsdk:"username"`
	Password types.String `tfsdk:"password"`
}

type PrivateHostedZone struct {
	ID      types.String `tfsdk:"id"`
	RoleARN types.String `tfsdk:"role_arn"`
}
