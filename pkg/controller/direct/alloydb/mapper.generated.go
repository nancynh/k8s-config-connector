// Copyright 2025 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +generated:mapper
// krm.group: alloydb.cnrm.cloud.google.com
// krm.version: v1beta1
// proto.service: google.cloud.alloydb.v1beta

package alloydb

import (
	pb "cloud.google.com/go/alloydb/apiv1beta/alloydbpb"
	krmalloydbv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/alloydb/v1alpha1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/alloydb/v1beta1"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func AlloyDBClusterObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Cluster) *krmalloydbv1alpha1.AlloyDBClusterObservedState {
	if in == nil {
		return nil
	}
	out := &krmalloydbv1alpha1.AlloyDBClusterObservedState{}
	// MISSING: BackupSource
	// MISSING: MigrationSource
	// MISSING: CloudsqlBackupRunSource
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Uid
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: DeleteTime
	// MISSING: Labels
	// MISSING: State
	// MISSING: ClusterType
	// MISSING: DatabaseVersion
	// MISSING: NetworkConfig
	// MISSING: Network
	// MISSING: Etag
	// MISSING: Annotations
	// MISSING: Reconciling
	// MISSING: InitialUser
	// MISSING: AutomatedBackupPolicy
	// MISSING: SSLConfig
	// MISSING: EncryptionConfig
	// MISSING: EncryptionInfo
	// MISSING: ContinuousBackupConfig
	// MISSING: ContinuousBackupInfo
	// MISSING: SecondaryConfig
	// MISSING: PrimaryConfig
	// MISSING: SatisfiesPzs
	// MISSING: PSCConfig
	// MISSING: MaintenanceUpdatePolicy
	// MISSING: MaintenanceSchedule
	// MISSING: GeminiConfig
	// MISSING: SubscriptionType
	// MISSING: TrialMetadata
	// MISSING: Tags
	// MISSING: ServiceAccountEmail
	return out
}
func AlloyDBClusterObservedState_ToProto(mapCtx *direct.MapContext, in *krmalloydbv1alpha1.AlloyDBClusterObservedState) *pb.Cluster {
	if in == nil {
		return nil
	}
	out := &pb.Cluster{}
	// MISSING: BackupSource
	// MISSING: MigrationSource
	// MISSING: CloudsqlBackupRunSource
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Uid
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: DeleteTime
	// MISSING: Labels
	// MISSING: State
	// MISSING: ClusterType
	// MISSING: DatabaseVersion
	// MISSING: NetworkConfig
	// MISSING: Network
	// MISSING: Etag
	// MISSING: Annotations
	// MISSING: Reconciling
	// MISSING: InitialUser
	// MISSING: AutomatedBackupPolicy
	// MISSING: SSLConfig
	// MISSING: EncryptionConfig
	// MISSING: EncryptionInfo
	// MISSING: ContinuousBackupConfig
	// MISSING: ContinuousBackupInfo
	// MISSING: SecondaryConfig
	// MISSING: PrimaryConfig
	// MISSING: SatisfiesPzs
	// MISSING: PSCConfig
	// MISSING: MaintenanceUpdatePolicy
	// MISSING: MaintenanceSchedule
	// MISSING: GeminiConfig
	// MISSING: SubscriptionType
	// MISSING: TrialMetadata
	// MISSING: Tags
	// MISSING: ServiceAccountEmail
	return out
}
func AlloyDBClusterObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Cluster) *krm.AlloyDBClusterObservedState {
	if in == nil {
		return nil
	}
	out := &krm.AlloyDBClusterObservedState{}
	// MISSING: BackupSource
	// MISSING: MigrationSource
	// MISSING: CloudsqlBackupRunSource
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Uid
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: DeleteTime
	// MISSING: Labels
	// MISSING: State
	out.ClusterType = direct.Enum_FromProto(mapCtx, in.GetClusterType())
	out.DatabaseVersion = direct.Enum_FromProto(mapCtx, in.GetDatabaseVersion())
	// MISSING: NetworkConfig
	// MISSING: Network
	// MISSING: Etag
	// MISSING: Annotations
	// MISSING: Reconciling
	// MISSING: InitialUser
	// MISSING: AutomatedBackupPolicy
	// MISSING: SSLConfig
	// MISSING: EncryptionConfig
	// MISSING: EncryptionInfo
	// MISSING: ContinuousBackupConfig
	// MISSING: ContinuousBackupInfo
	// MISSING: SecondaryConfig
	// MISSING: PrimaryConfig
	// MISSING: SatisfiesPzs
	// MISSING: PSCConfig
	// MISSING: MaintenanceUpdatePolicy
	// MISSING: MaintenanceSchedule
	// MISSING: GeminiConfig
	// MISSING: SubscriptionType
	// MISSING: TrialMetadata
	// MISSING: Tags
	// MISSING: ServiceAccountEmail
	return out
}
func AlloyDBClusterObservedState_ToProto(mapCtx *direct.MapContext, in *krm.AlloyDBClusterObservedState) *pb.Cluster {
	if in == nil {
		return nil
	}
	out := &pb.Cluster{}
	// MISSING: BackupSource
	// MISSING: MigrationSource
	// MISSING: CloudsqlBackupRunSource
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Uid
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: DeleteTime
	// MISSING: Labels
	// MISSING: State
	out.ClusterType = direct.Enum_ToProto[pb.Cluster_ClusterType](mapCtx, in.ClusterType)
	out.DatabaseVersion = direct.Enum_ToProto[pb.DatabaseVersion](mapCtx, in.DatabaseVersion)
	// MISSING: NetworkConfig
	// MISSING: Network
	// MISSING: Etag
	// MISSING: Annotations
	// MISSING: Reconciling
	// MISSING: InitialUser
	// MISSING: AutomatedBackupPolicy
	// MISSING: SSLConfig
	// MISSING: EncryptionConfig
	// MISSING: EncryptionInfo
	// MISSING: ContinuousBackupConfig
	// MISSING: ContinuousBackupInfo
	// MISSING: SecondaryConfig
	// MISSING: PrimaryConfig
	// MISSING: SatisfiesPzs
	// MISSING: PSCConfig
	// MISSING: MaintenanceUpdatePolicy
	// MISSING: MaintenanceSchedule
	// MISSING: GeminiConfig
	// MISSING: SubscriptionType
	// MISSING: TrialMetadata
	// MISSING: Tags
	// MISSING: ServiceAccountEmail
	return out
}
func AlloyDBClusterSpec_FromProto(mapCtx *direct.MapContext, in *pb.Cluster) *krmalloydbv1alpha1.AlloyDBClusterSpec {
	if in == nil {
		return nil
	}
	out := &krmalloydbv1alpha1.AlloyDBClusterSpec{}
	// MISSING: BackupSource
	// MISSING: MigrationSource
	// MISSING: CloudsqlBackupRunSource
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Uid
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: DeleteTime
	// MISSING: Labels
	// MISSING: State
	// MISSING: ClusterType
	// MISSING: DatabaseVersion
	// MISSING: NetworkConfig
	// MISSING: Network
	// MISSING: Etag
	// MISSING: Annotations
	// MISSING: Reconciling
	// MISSING: InitialUser
	// MISSING: AutomatedBackupPolicy
	// MISSING: SSLConfig
	// MISSING: EncryptionConfig
	// MISSING: EncryptionInfo
	// MISSING: ContinuousBackupConfig
	// MISSING: ContinuousBackupInfo
	// MISSING: SecondaryConfig
	// MISSING: PrimaryConfig
	// MISSING: SatisfiesPzs
	// MISSING: PSCConfig
	// MISSING: MaintenanceUpdatePolicy
	// MISSING: MaintenanceSchedule
	// MISSING: GeminiConfig
	// MISSING: SubscriptionType
	// MISSING: TrialMetadata
	// MISSING: Tags
	// MISSING: ServiceAccountEmail
	return out
}
func AlloyDBClusterSpec_ToProto(mapCtx *direct.MapContext, in *krmalloydbv1alpha1.AlloyDBClusterSpec) *pb.Cluster {
	if in == nil {
		return nil
	}
	out := &pb.Cluster{}
	// MISSING: BackupSource
	// MISSING: MigrationSource
	// MISSING: CloudsqlBackupRunSource
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Uid
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: DeleteTime
	// MISSING: Labels
	// MISSING: State
	// MISSING: ClusterType
	// MISSING: DatabaseVersion
	// MISSING: NetworkConfig
	// MISSING: Network
	// MISSING: Etag
	// MISSING: Annotations
	// MISSING: Reconciling
	// MISSING: InitialUser
	// MISSING: AutomatedBackupPolicy
	// MISSING: SSLConfig
	// MISSING: EncryptionConfig
	// MISSING: EncryptionInfo
	// MISSING: ContinuousBackupConfig
	// MISSING: ContinuousBackupInfo
	// MISSING: SecondaryConfig
	// MISSING: PrimaryConfig
	// MISSING: SatisfiesPzs
	// MISSING: PSCConfig
	// MISSING: MaintenanceUpdatePolicy
	// MISSING: MaintenanceSchedule
	// MISSING: GeminiConfig
	// MISSING: SubscriptionType
	// MISSING: TrialMetadata
	// MISSING: Tags
	// MISSING: ServiceAccountEmail
	return out
}
func AlloyDBClusterSpec_FromProto(mapCtx *direct.MapContext, in *pb.Cluster) *krm.AlloyDBClusterSpec {
	if in == nil {
		return nil
	}
	out := &krm.AlloyDBClusterSpec{}
	// MISSING: BackupSource
	// MISSING: MigrationSource
	// MISSING: CloudsqlBackupRunSource
	// MISSING: Name
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	// MISSING: Uid
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: DeleteTime
	// MISSING: Labels
	// MISSING: State
	out.ClusterType = direct.Enum_FromProto(mapCtx, in.GetClusterType())
	out.DatabaseVersion = direct.Enum_FromProto(mapCtx, in.GetDatabaseVersion())
	out.NetworkConfig = Cluster_NetworkConfig_FromProto(mapCtx, in.GetNetworkConfig())
	if in.GetNetwork() != "" {
		out.NetworkRef = &refsv1beta1.ComputeNetworkRef{External: in.GetNetwork()}
	}
	// MISSING: Etag
	// MISSING: Annotations
	// MISSING: Reconciling
	out.InitialUser = UserPassword_FromProto(mapCtx, in.GetInitialUser())
	out.AutomatedBackupPolicy = AutomatedBackupPolicy_FromProto(mapCtx, in.GetAutomatedBackupPolicy())
	// MISSING: SSLConfig
	out.EncryptionConfig = EncryptionConfig_FromProto(mapCtx, in.GetEncryptionConfig())
	// MISSING: EncryptionInfo
	out.ContinuousBackupConfig = ContinuousBackupConfig_FromProto(mapCtx, in.GetContinuousBackupConfig())
	// MISSING: ContinuousBackupInfo
	out.SecondaryConfig = Cluster_SecondaryConfig_FromProto(mapCtx, in.GetSecondaryConfig())
	// MISSING: PrimaryConfig
	// MISSING: SatisfiesPzs
	// MISSING: PSCConfig
	out.MaintenanceUpdatePolicy = MaintenanceUpdatePolicy_FromProto(mapCtx, in.GetMaintenanceUpdatePolicy())
	// MISSING: MaintenanceSchedule
	// MISSING: GeminiConfig
	// MISSING: SubscriptionType
	// MISSING: TrialMetadata
	// MISSING: Tags
	// MISSING: ServiceAccountEmail
	return out
}
func AlloyDBClusterSpec_ToProto(mapCtx *direct.MapContext, in *krm.AlloyDBClusterSpec) *pb.Cluster {
	if in == nil {
		return nil
	}
	out := &pb.Cluster{}
	// MISSING: BackupSource
	// MISSING: MigrationSource
	// MISSING: CloudsqlBackupRunSource
	// MISSING: Name
	out.DisplayName = direct.ValueOf(in.DisplayName)
	// MISSING: Uid
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: DeleteTime
	// MISSING: Labels
	// MISSING: State
	out.ClusterType = direct.Enum_ToProto[pb.Cluster_ClusterType](mapCtx, in.ClusterType)
	out.DatabaseVersion = direct.Enum_ToProto[pb.DatabaseVersion](mapCtx, in.DatabaseVersion)
	out.NetworkConfig = Cluster_NetworkConfig_ToProto(mapCtx, in.NetworkConfig)
	if in.NetworkRef != nil {
		out.Network = in.NetworkRef.External
	}
	// MISSING: Etag
	// MISSING: Annotations
	// MISSING: Reconciling
	out.InitialUser = UserPassword_ToProto(mapCtx, in.InitialUser)
	out.AutomatedBackupPolicy = AutomatedBackupPolicy_ToProto(mapCtx, in.AutomatedBackupPolicy)
	// MISSING: SSLConfig
	out.EncryptionConfig = EncryptionConfig_ToProto(mapCtx, in.EncryptionConfig)
	// MISSING: EncryptionInfo
	out.ContinuousBackupConfig = ContinuousBackupConfig_ToProto(mapCtx, in.ContinuousBackupConfig)
	// MISSING: ContinuousBackupInfo
	out.SecondaryConfig = Cluster_SecondaryConfig_ToProto(mapCtx, in.SecondaryConfig)
	// MISSING: PrimaryConfig
	// MISSING: SatisfiesPzs
	// MISSING: PSCConfig
	out.MaintenanceUpdatePolicy = MaintenanceUpdatePolicy_ToProto(mapCtx, in.MaintenanceUpdatePolicy)
	// MISSING: MaintenanceSchedule
	// MISSING: GeminiConfig
	// MISSING: SubscriptionType
	// MISSING: TrialMetadata
	// MISSING: Tags
	// MISSING: ServiceAccountEmail
	return out
}
func AlloyDBInstanceObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Instance) *krmalloydbv1alpha1.AlloyDBInstanceObservedState {
	if in == nil {
		return nil
	}
	out := &krmalloydbv1alpha1.AlloyDBInstanceObservedState{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Uid
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: DeleteTime
	// MISSING: Labels
	// MISSING: State
	// MISSING: InstanceType
	// MISSING: MachineConfig
	// MISSING: AvailabilityType
	// MISSING: GCEZone
	// MISSING: DatabaseFlags
	// MISSING: WritableNode
	// MISSING: Nodes
	// MISSING: QueryInsightsConfig
	// MISSING: ObservabilityConfig
	// MISSING: ReadPoolConfig
	// MISSING: IPAddress
	// MISSING: PublicIPAddress
	// MISSING: Reconciling
	// MISSING: Etag
	// MISSING: Annotations
	// MISSING: UpdatePolicy
	// MISSING: ClientConnectionConfig
	// MISSING: SatisfiesPzs
	// MISSING: PSCInstanceConfig
	// MISSING: NetworkConfig
	// MISSING: GeminiConfig
	// MISSING: OutboundPublicIPAddresses
	// MISSING: ActivationPolicy
	// MISSING: ConnectionPoolConfig
	// MISSING: GcaConfig
	return out
}
func AlloyDBInstanceObservedState_ToProto(mapCtx *direct.MapContext, in *krmalloydbv1alpha1.AlloyDBInstanceObservedState) *pb.Instance {
	if in == nil {
		return nil
	}
	out := &pb.Instance{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Uid
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: DeleteTime
	// MISSING: Labels
	// MISSING: State
	// MISSING: InstanceType
	// MISSING: MachineConfig
	// MISSING: AvailabilityType
	// MISSING: GCEZone
	// MISSING: DatabaseFlags
	// MISSING: WritableNode
	// MISSING: Nodes
	// MISSING: QueryInsightsConfig
	// MISSING: ObservabilityConfig
	// MISSING: ReadPoolConfig
	// MISSING: IPAddress
	// MISSING: PublicIPAddress
	// MISSING: Reconciling
	// MISSING: Etag
	// MISSING: Annotations
	// MISSING: UpdatePolicy
	// MISSING: ClientConnectionConfig
	// MISSING: SatisfiesPzs
	// MISSING: PSCInstanceConfig
	// MISSING: NetworkConfig
	// MISSING: GeminiConfig
	// MISSING: OutboundPublicIPAddresses
	// MISSING: ActivationPolicy
	// MISSING: ConnectionPoolConfig
	// MISSING: GcaConfig
	return out
}
func AutomatedBackupPolicy_FromProto(mapCtx *direct.MapContext, in *pb.AutomatedBackupPolicy) *krmalloydbv1alpha1.AutomatedBackupPolicy {
	if in == nil {
		return nil
	}
	out := &krmalloydbv1alpha1.AutomatedBackupPolicy{}
	out.WeeklySchedule = AutomatedBackupPolicy_WeeklySchedule_FromProto(mapCtx, in.GetWeeklySchedule())
	out.TimeBasedRetention = AutomatedBackupPolicy_TimeBasedRetention_FromProto(mapCtx, in.GetTimeBasedRetention())
	out.QuantityBasedRetention = AutomatedBackupPolicy_QuantityBasedRetention_FromProto(mapCtx, in.GetQuantityBasedRetention())
	out.Enabled = in.Enabled
	out.BackupWindow = direct.StringDuration_FromProto(mapCtx, in.GetBackupWindow())
	out.EncryptionConfig = EncryptionConfig_FromProto(mapCtx, in.GetEncryptionConfig())
	out.Location = direct.LazyPtr(in.GetLocation())
	out.Labels = in.Labels
	return out
}
func AutomatedBackupPolicy_ToProto(mapCtx *direct.MapContext, in *krmalloydbv1alpha1.AutomatedBackupPolicy) *pb.AutomatedBackupPolicy {
	if in == nil {
		return nil
	}
	out := &pb.AutomatedBackupPolicy{}
	if oneof := AutomatedBackupPolicy_WeeklySchedule_ToProto(mapCtx, in.WeeklySchedule); oneof != nil {
		out.Schedule = &pb.AutomatedBackupPolicy_WeeklySchedule_{WeeklySchedule: oneof}
	}
	if oneof := AutomatedBackupPolicy_TimeBasedRetention_ToProto(mapCtx, in.TimeBasedRetention); oneof != nil {
		out.Retention = &pb.AutomatedBackupPolicy_TimeBasedRetention_{TimeBasedRetention: oneof}
	}
	if oneof := AutomatedBackupPolicy_QuantityBasedRetention_ToProto(mapCtx, in.QuantityBasedRetention); oneof != nil {
		out.Retention = &pb.AutomatedBackupPolicy_QuantityBasedRetention_{QuantityBasedRetention: oneof}
	}
	out.Enabled = in.Enabled
	out.BackupWindow = direct.StringDuration_ToProto(mapCtx, in.BackupWindow)
	out.EncryptionConfig = EncryptionConfig_ToProto(mapCtx, in.EncryptionConfig)
	out.Location = direct.ValueOf(in.Location)
	out.Labels = in.Labels
	return out
}
func AutomatedBackupPolicy_FromProto(mapCtx *direct.MapContext, in *pb.AutomatedBackupPolicy) *krm.AutomatedBackupPolicy {
	if in == nil {
		return nil
	}
	out := &krm.AutomatedBackupPolicy{}
	out.WeeklySchedule = AutomatedBackupPolicy_WeeklySchedule_FromProto(mapCtx, in.GetWeeklySchedule())
	out.TimeBasedRetention = AutomatedBackupPolicy_TimeBasedRetention_FromProto(mapCtx, in.GetTimeBasedRetention())
	out.QuantityBasedRetention = AutomatedBackupPolicy_QuantityBasedRetention_FromProto(mapCtx, in.GetQuantityBasedRetention())
	out.Enabled = in.Enabled
	out.BackupWindow = direct.StringDuration_FromProto(mapCtx, in.GetBackupWindow())
	out.EncryptionConfig = EncryptionConfig_FromProto(mapCtx, in.GetEncryptionConfig())
	out.Location = direct.LazyPtr(in.GetLocation())
	out.Labels = in.Labels
	return out
}
func AutomatedBackupPolicy_ToProto(mapCtx *direct.MapContext, in *krm.AutomatedBackupPolicy) *pb.AutomatedBackupPolicy {
	if in == nil {
		return nil
	}
	out := &pb.AutomatedBackupPolicy{}
	if oneof := AutomatedBackupPolicy_WeeklySchedule_ToProto(mapCtx, in.WeeklySchedule); oneof != nil {
		out.Schedule = &pb.AutomatedBackupPolicy_WeeklySchedule_{WeeklySchedule: oneof}
	}
	if oneof := AutomatedBackupPolicy_TimeBasedRetention_ToProto(mapCtx, in.TimeBasedRetention); oneof != nil {
		out.Retention = &pb.AutomatedBackupPolicy_TimeBasedRetention_{TimeBasedRetention: oneof}
	}
	if oneof := AutomatedBackupPolicy_QuantityBasedRetention_ToProto(mapCtx, in.QuantityBasedRetention); oneof != nil {
		out.Retention = &pb.AutomatedBackupPolicy_QuantityBasedRetention_{QuantityBasedRetention: oneof}
	}
	out.Enabled = in.Enabled
	out.BackupWindow = direct.StringDuration_ToProto(mapCtx, in.BackupWindow)
	out.EncryptionConfig = EncryptionConfig_ToProto(mapCtx, in.EncryptionConfig)
	out.Location = direct.ValueOf(in.Location)
	out.Labels = in.Labels
	return out
}
func AutomatedBackupPolicy_QuantityBasedRetention_FromProto(mapCtx *direct.MapContext, in *pb.AutomatedBackupPolicy_QuantityBasedRetention) *krmalloydbv1alpha1.AutomatedBackupPolicy_QuantityBasedRetention {
	if in == nil {
		return nil
	}
	out := &krmalloydbv1alpha1.AutomatedBackupPolicy_QuantityBasedRetention{}
	out.Count = direct.LazyPtr(in.GetCount())
	return out
}
func AutomatedBackupPolicy_QuantityBasedRetention_ToProto(mapCtx *direct.MapContext, in *krmalloydbv1alpha1.AutomatedBackupPolicy_QuantityBasedRetention) *pb.AutomatedBackupPolicy_QuantityBasedRetention {
	if in == nil {
		return nil
	}
	out := &pb.AutomatedBackupPolicy_QuantityBasedRetention{}
	out.Count = direct.ValueOf(in.Count)
	return out
}
func AutomatedBackupPolicy_QuantityBasedRetention_FromProto(mapCtx *direct.MapContext, in *pb.AutomatedBackupPolicy_QuantityBasedRetention) *krm.AutomatedBackupPolicy_QuantityBasedRetention {
	if in == nil {
		return nil
	}
	out := &krm.AutomatedBackupPolicy_QuantityBasedRetention{}
	out.Count = direct.LazyPtr(in.GetCount())
	return out
}
func AutomatedBackupPolicy_QuantityBasedRetention_ToProto(mapCtx *direct.MapContext, in *krm.AutomatedBackupPolicy_QuantityBasedRetention) *pb.AutomatedBackupPolicy_QuantityBasedRetention {
	if in == nil {
		return nil
	}
	out := &pb.AutomatedBackupPolicy_QuantityBasedRetention{}
	out.Count = direct.ValueOf(in.Count)
	return out
}
func AutomatedBackupPolicy_TimeBasedRetention_FromProto(mapCtx *direct.MapContext, in *pb.AutomatedBackupPolicy_TimeBasedRetention) *krmalloydbv1alpha1.AutomatedBackupPolicy_TimeBasedRetention {
	if in == nil {
		return nil
	}
	out := &krmalloydbv1alpha1.AutomatedBackupPolicy_TimeBasedRetention{}
	out.RetentionPeriod = direct.StringDuration_FromProto(mapCtx, in.GetRetentionPeriod())
	return out
}
func AutomatedBackupPolicy_TimeBasedRetention_ToProto(mapCtx *direct.MapContext, in *krmalloydbv1alpha1.AutomatedBackupPolicy_TimeBasedRetention) *pb.AutomatedBackupPolicy_TimeBasedRetention {
	if in == nil {
		return nil
	}
	out := &pb.AutomatedBackupPolicy_TimeBasedRetention{}
	out.RetentionPeriod = direct.StringDuration_ToProto(mapCtx, in.RetentionPeriod)
	return out
}
func AutomatedBackupPolicy_TimeBasedRetention_FromProto(mapCtx *direct.MapContext, in *pb.AutomatedBackupPolicy_TimeBasedRetention) *krm.AutomatedBackupPolicy_TimeBasedRetention {
	if in == nil {
		return nil
	}
	out := &krm.AutomatedBackupPolicy_TimeBasedRetention{}
	out.RetentionPeriod = direct.StringDuration_FromProto(mapCtx, in.GetRetentionPeriod())
	return out
}
func AutomatedBackupPolicy_TimeBasedRetention_ToProto(mapCtx *direct.MapContext, in *krm.AutomatedBackupPolicy_TimeBasedRetention) *pb.AutomatedBackupPolicy_TimeBasedRetention {
	if in == nil {
		return nil
	}
	out := &pb.AutomatedBackupPolicy_TimeBasedRetention{}
	out.RetentionPeriod = direct.StringDuration_ToProto(mapCtx, in.RetentionPeriod)
	return out
}
func BackupSource_FromProto(mapCtx *direct.MapContext, in *pb.BackupSource) *krmalloydbv1alpha1.BackupSource {
	if in == nil {
		return nil
	}
	out := &krmalloydbv1alpha1.BackupSource{}
	// MISSING: BackupUid
	out.BackupName = direct.LazyPtr(in.GetBackupName())
	return out
}
func BackupSource_ToProto(mapCtx *direct.MapContext, in *krmalloydbv1alpha1.BackupSource) *pb.BackupSource {
	if in == nil {
		return nil
	}
	out := &pb.BackupSource{}
	// MISSING: BackupUid
	out.BackupName = direct.ValueOf(in.BackupName)
	return out
}
func BackupSource_FromProto(mapCtx *direct.MapContext, in *pb.BackupSource) *krm.BackupSource {
	if in == nil {
		return nil
	}
	out := &krm.BackupSource{}
	// MISSING: BackupUid
	if in.GetBackupName() != "" {
		out.BackupNameRef = &refsv1beta1.AlloyDBBackupRef{External: in.GetBackupName()}
	}
	return out
}
func BackupSource_ToProto(mapCtx *direct.MapContext, in *krm.BackupSource) *pb.BackupSource {
	if in == nil {
		return nil
	}
	out := &pb.BackupSource{}
	// MISSING: BackupUid
	if in.BackupNameRef != nil {
		out.BackupName = in.BackupNameRef.External
	}
	return out
}
func BackupSourceObservedState_FromProto(mapCtx *direct.MapContext, in *pb.BackupSource) *krmalloydbv1alpha1.BackupSourceObservedState {
	if in == nil {
		return nil
	}
	out := &krmalloydbv1alpha1.BackupSourceObservedState{}
	out.BackupUid = direct.LazyPtr(in.GetBackupUid())
	// MISSING: BackupName
	return out
}
func BackupSourceObservedState_ToProto(mapCtx *direct.MapContext, in *krmalloydbv1alpha1.BackupSourceObservedState) *pb.BackupSource {
	if in == nil {
		return nil
	}
	out := &pb.BackupSource{}
	out.BackupUid = direct.ValueOf(in.BackupUid)
	// MISSING: BackupName
	return out
}
func BackupSourceObservedState_FromProto(mapCtx *direct.MapContext, in *pb.BackupSource) *krm.BackupSourceObservedState {
	if in == nil {
		return nil
	}
	out := &krm.BackupSourceObservedState{}
	// MISSING: BackupUid
	out.BackupName = direct.LazyPtr(in.GetBackupName())
	return out
}
func BackupSourceObservedState_ToProto(mapCtx *direct.MapContext, in *krm.BackupSourceObservedState) *pb.BackupSource {
	if in == nil {
		return nil
	}
	out := &pb.BackupSource{}
	// MISSING: BackupUid
	out.BackupName = direct.ValueOf(in.BackupName)
	return out
}
func CloudSQLBackupRunSource_FromProto(mapCtx *direct.MapContext, in *pb.CloudSQLBackupRunSource) *krmalloydbv1alpha1.CloudSQLBackupRunSource {
	if in == nil {
		return nil
	}
	out := &krmalloydbv1alpha1.CloudSQLBackupRunSource{}
	out.Project = direct.LazyPtr(in.GetProject())
	out.InstanceID = direct.LazyPtr(in.GetInstanceId())
	out.BackupRunID = direct.LazyPtr(in.GetBackupRunId())
	return out
}
func CloudSQLBackupRunSource_ToProto(mapCtx *direct.MapContext, in *krmalloydbv1alpha1.CloudSQLBackupRunSource) *pb.CloudSQLBackupRunSource {
	if in == nil {
		return nil
	}
	out := &pb.CloudSQLBackupRunSource{}
	out.Project = direct.ValueOf(in.Project)
	out.InstanceId = direct.ValueOf(in.InstanceID)
	out.BackupRunId = direct.ValueOf(in.BackupRunID)
	return out
}
func CloudSQLBackupRunSource_FromProto(mapCtx *direct.MapContext, in *pb.CloudSQLBackupRunSource) *krm.CloudSQLBackupRunSource {
	if in == nil {
		return nil
	}
	out := &krm.CloudSQLBackupRunSource{}
	out.Project = direct.LazyPtr(in.GetProject())
	out.InstanceID = direct.LazyPtr(in.GetInstanceId())
	out.BackupRunID = direct.LazyPtr(in.GetBackupRunId())
	return out
}
func CloudSQLBackupRunSource_ToProto(mapCtx *direct.MapContext, in *krm.CloudSQLBackupRunSource) *pb.CloudSQLBackupRunSource {
	if in == nil {
		return nil
	}
	out := &pb.CloudSQLBackupRunSource{}
	out.Project = direct.ValueOf(in.Project)
	out.InstanceId = direct.ValueOf(in.InstanceID)
	out.BackupRunId = direct.ValueOf(in.BackupRunID)
	return out
}
func Cluster_FromProto(mapCtx *direct.MapContext, in *pb.Cluster) *krmalloydbv1alpha1.Cluster {
	if in == nil {
		return nil
	}
	out := &krmalloydbv1alpha1.Cluster{}
	// MISSING: BackupSource
	// MISSING: MigrationSource
	// MISSING: CloudsqlBackupRunSource
	// MISSING: Name
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	// MISSING: Uid
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: DeleteTime
	out.Labels = in.Labels
	// MISSING: State
	// MISSING: ClusterType
	out.DatabaseVersion = direct.Enum_FromProto(mapCtx, in.GetDatabaseVersion())
	out.NetworkConfig = Cluster_NetworkConfig_FromProto(mapCtx, in.GetNetworkConfig())
	out.Network = direct.LazyPtr(in.GetNetwork())
	out.Etag = direct.LazyPtr(in.GetEtag())
	out.Annotations = in.Annotations
	// MISSING: Reconciling
	out.InitialUser = UserPassword_FromProto(mapCtx, in.GetInitialUser())
	out.AutomatedBackupPolicy = AutomatedBackupPolicy_FromProto(mapCtx, in.GetAutomatedBackupPolicy())
	out.SSLConfig = SSLConfig_FromProto(mapCtx, in.GetSslConfig())
	out.EncryptionConfig = EncryptionConfig_FromProto(mapCtx, in.GetEncryptionConfig())
	// MISSING: EncryptionInfo
	out.ContinuousBackupConfig = ContinuousBackupConfig_FromProto(mapCtx, in.GetContinuousBackupConfig())
	// MISSING: ContinuousBackupInfo
	out.SecondaryConfig = Cluster_SecondaryConfig_FromProto(mapCtx, in.GetSecondaryConfig())
	// MISSING: PrimaryConfig
	// MISSING: SatisfiesPzs
	out.PSCConfig = Cluster_PSCConfig_FromProto(mapCtx, in.GetPscConfig())
	out.MaintenanceUpdatePolicy = MaintenanceUpdatePolicy_FromProto(mapCtx, in.GetMaintenanceUpdatePolicy())
	// MISSING: MaintenanceSchedule
	out.GeminiConfig = GeminiClusterConfig_FromProto(mapCtx, in.GetGeminiConfig())
	out.SubscriptionType = direct.Enum_FromProto(mapCtx, in.GetSubscriptionType())
	// MISSING: TrialMetadata
	out.Tags = in.Tags
	// MISSING: ServiceAccountEmail
	return out
}
func Cluster_ToProto(mapCtx *direct.MapContext, in *krmalloydbv1alpha1.Cluster) *pb.Cluster {
	if in == nil {
		return nil
	}
	out := &pb.Cluster{}
	// MISSING: BackupSource
	// MISSING: MigrationSource
	// MISSING: CloudsqlBackupRunSource
	// MISSING: Name
	out.DisplayName = direct.ValueOf(in.DisplayName)
	// MISSING: Uid
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: DeleteTime
	out.Labels = in.Labels
	// MISSING: State
	// MISSING: ClusterType
	out.DatabaseVersion = direct.Enum_ToProto[pb.DatabaseVersion](mapCtx, in.DatabaseVersion)
	out.NetworkConfig = Cluster_NetworkConfig_ToProto(mapCtx, in.NetworkConfig)
	out.Network = direct.ValueOf(in.Network)
	out.Etag = direct.ValueOf(in.Etag)
	out.Annotations = in.Annotations
	// MISSING: Reconciling
	out.InitialUser = UserPassword_ToProto(mapCtx, in.InitialUser)
	out.AutomatedBackupPolicy = AutomatedBackupPolicy_ToProto(mapCtx, in.AutomatedBackupPolicy)
	out.SslConfig = SSLConfig_ToProto(mapCtx, in.SSLConfig)
	out.EncryptionConfig = EncryptionConfig_ToProto(mapCtx, in.EncryptionConfig)
	// MISSING: EncryptionInfo
	out.ContinuousBackupConfig = ContinuousBackupConfig_ToProto(mapCtx, in.ContinuousBackupConfig)
	// MISSING: ContinuousBackupInfo
	out.SecondaryConfig = Cluster_SecondaryConfig_ToProto(mapCtx, in.SecondaryConfig)
	// MISSING: PrimaryConfig
	// MISSING: SatisfiesPzs
	out.PscConfig = Cluster_PSCConfig_ToProto(mapCtx, in.PSCConfig)
	out.MaintenanceUpdatePolicy = MaintenanceUpdatePolicy_ToProto(mapCtx, in.MaintenanceUpdatePolicy)
	// MISSING: MaintenanceSchedule
	out.GeminiConfig = GeminiClusterConfig_ToProto(mapCtx, in.GeminiConfig)
	out.SubscriptionType = direct.Enum_ToProto[pb.SubscriptionType](mapCtx, in.SubscriptionType)
	// MISSING: TrialMetadata
	out.Tags = in.Tags
	// MISSING: ServiceAccountEmail
	return out
}
func ClusterObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Cluster) *krmalloydbv1alpha1.ClusterObservedState {
	if in == nil {
		return nil
	}
	out := &krmalloydbv1alpha1.ClusterObservedState{}
	out.BackupSource = BackupSource_FromProto(mapCtx, in.GetBackupSource())
	out.MigrationSource = MigrationSource_FromProto(mapCtx, in.GetMigrationSource())
	out.CloudsqlBackupRunSource = CloudSQLBackupRunSource_FromProto(mapCtx, in.GetCloudsqlBackupRunSource())
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: DisplayName
	out.Uid = direct.LazyPtr(in.GetUid())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.DeleteTime = direct.StringTimestamp_FromProto(mapCtx, in.GetDeleteTime())
	// MISSING: Labels
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.ClusterType = direct.Enum_FromProto(mapCtx, in.GetClusterType())
	// MISSING: DatabaseVersion
	// MISSING: NetworkConfig
	// MISSING: Network
	// MISSING: Etag
	// MISSING: Annotations
	out.Reconciling = direct.LazyPtr(in.GetReconciling())
	// MISSING: InitialUser
	// MISSING: AutomatedBackupPolicy
	// MISSING: SSLConfig
	// MISSING: EncryptionConfig
	out.EncryptionInfo = EncryptionInfo_FromProto(mapCtx, in.GetEncryptionInfo())
	// MISSING: ContinuousBackupConfig
	out.ContinuousBackupInfo = ContinuousBackupInfo_FromProto(mapCtx, in.GetContinuousBackupInfo())
	// MISSING: SecondaryConfig
	out.PrimaryConfig = Cluster_PrimaryConfig_FromProto(mapCtx, in.GetPrimaryConfig())
	out.SatisfiesPzs = direct.LazyPtr(in.GetSatisfiesPzs())
	out.PSCConfig = Cluster_PSCConfigObservedState_FromProto(mapCtx, in.GetPscConfig())
	// MISSING: MaintenanceUpdatePolicy
	out.MaintenanceSchedule = MaintenanceSchedule_FromProto(mapCtx, in.GetMaintenanceSchedule())
	out.GeminiConfig = GeminiClusterConfigObservedState_FromProto(mapCtx, in.GetGeminiConfig())
	// MISSING: SubscriptionType
	out.TrialMetadata = Cluster_TrialMetadata_FromProto(mapCtx, in.GetTrialMetadata())
	// MISSING: Tags
	out.ServiceAccountEmail = direct.LazyPtr(in.GetServiceAccountEmail())
	return out
}
func ClusterObservedState_ToProto(mapCtx *direct.MapContext, in *krmalloydbv1alpha1.ClusterObservedState) *pb.Cluster {
	if in == nil {
		return nil
	}
	out := &pb.Cluster{}
	if oneof := BackupSource_ToProto(mapCtx, in.BackupSource); oneof != nil {
		out.Source = &pb.Cluster_BackupSource{BackupSource: oneof}
	}
	if oneof := MigrationSource_ToProto(mapCtx, in.MigrationSource); oneof != nil {
		out.Source = &pb.Cluster_MigrationSource{MigrationSource: oneof}
	}
	if oneof := CloudSQLBackupRunSource_ToProto(mapCtx, in.CloudsqlBackupRunSource); oneof != nil {
		out.Source = &pb.Cluster_CloudsqlBackupRunSource{CloudsqlBackupRunSource: oneof}
	}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: DisplayName
	out.Uid = direct.ValueOf(in.Uid)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.DeleteTime = direct.StringTimestamp_ToProto(mapCtx, in.DeleteTime)
	// MISSING: Labels
	out.State = direct.Enum_ToProto[pb.Cluster_State](mapCtx, in.State)
	out.ClusterType = direct.Enum_ToProto[pb.Cluster_ClusterType](mapCtx, in.ClusterType)
	// MISSING: DatabaseVersion
	// MISSING: NetworkConfig
	// MISSING: Network
	// MISSING: Etag
	// MISSING: Annotations
	out.Reconciling = direct.ValueOf(in.Reconciling)
	// MISSING: InitialUser
	// MISSING: AutomatedBackupPolicy
	// MISSING: SSLConfig
	// MISSING: EncryptionConfig
	out.EncryptionInfo = EncryptionInfo_ToProto(mapCtx, in.EncryptionInfo)
	// MISSING: ContinuousBackupConfig
	out.ContinuousBackupInfo = ContinuousBackupInfo_ToProto(mapCtx, in.ContinuousBackupInfo)
	// MISSING: SecondaryConfig
	out.PrimaryConfig = Cluster_PrimaryConfig_ToProto(mapCtx, in.PrimaryConfig)
	out.SatisfiesPzs = direct.ValueOf(in.SatisfiesPzs)
	out.PscConfig = Cluster_PSCConfigObservedState_ToProto(mapCtx, in.PSCConfig)
	// MISSING: MaintenanceUpdatePolicy
	out.MaintenanceSchedule = MaintenanceSchedule_ToProto(mapCtx, in.MaintenanceSchedule)
	out.GeminiConfig = GeminiClusterConfigObservedState_ToProto(mapCtx, in.GeminiConfig)
	// MISSING: SubscriptionType
	out.TrialMetadata = Cluster_TrialMetadata_ToProto(mapCtx, in.TrialMetadata)
	// MISSING: Tags
	out.ServiceAccountEmail = direct.ValueOf(in.ServiceAccountEmail)
	return out
}
func Cluster_NetworkConfig_FromProto(mapCtx *direct.MapContext, in *pb.Cluster_NetworkConfig) *krmalloydbv1alpha1.Cluster_NetworkConfig {
	if in == nil {
		return nil
	}
	out := &krmalloydbv1alpha1.Cluster_NetworkConfig{}
	out.Network = direct.LazyPtr(in.GetNetwork())
	out.AllocatedIPRange = direct.LazyPtr(in.GetAllocatedIpRange())
	return out
}
func Cluster_NetworkConfig_ToProto(mapCtx *direct.MapContext, in *krmalloydbv1alpha1.Cluster_NetworkConfig) *pb.Cluster_NetworkConfig {
	if in == nil {
		return nil
	}
	out := &pb.Cluster_NetworkConfig{}
	out.Network = direct.ValueOf(in.Network)
	out.AllocatedIpRange = direct.ValueOf(in.AllocatedIPRange)
	return out
}
func Cluster_NetworkConfig_FromProto(mapCtx *direct.MapContext, in *pb.Cluster_NetworkConfig) *krm.Cluster_NetworkConfig {
	if in == nil {
		return nil
	}
	out := &krm.Cluster_NetworkConfig{}
	if in.GetNetwork() != "" {
		out.NetworkRef = &refsv1beta1.ComputeNetworkRef{External: in.GetNetwork()}
	}
	out.AllocatedIPRange = direct.LazyPtr(in.GetAllocatedIpRange())
	return out
}
func Cluster_NetworkConfig_ToProto(mapCtx *direct.MapContext, in *krm.Cluster_NetworkConfig) *pb.Cluster_NetworkConfig {
	if in == nil {
		return nil
	}
	out := &pb.Cluster_NetworkConfig{}
	if in.NetworkRef != nil {
		out.Network = in.NetworkRef.External
	}
	out.AllocatedIpRange = direct.ValueOf(in.AllocatedIPRange)
	return out
}
func Cluster_PSCConfig_FromProto(mapCtx *direct.MapContext, in *pb.Cluster_PscConfig) *krmalloydbv1alpha1.Cluster_PSCConfig {
	if in == nil {
		return nil
	}
	out := &krmalloydbv1alpha1.Cluster_PSCConfig{}
	out.PSCEnabled = direct.LazyPtr(in.GetPscEnabled())
	// MISSING: ServiceOwnedProjectNumber
	return out
}
func Cluster_PSCConfig_ToProto(mapCtx *direct.MapContext, in *krmalloydbv1alpha1.Cluster_PSCConfig) *pb.Cluster_PscConfig {
	if in == nil {
		return nil
	}
	out := &pb.Cluster_PscConfig{}
	out.PscEnabled = direct.ValueOf(in.PSCEnabled)
	// MISSING: ServiceOwnedProjectNumber
	return out
}
func Cluster_PSCConfig_FromProto(mapCtx *direct.MapContext, in *pb.Cluster_PscConfig) *krm.Cluster_PSCConfig {
	if in == nil {
		return nil
	}
	out := &krm.Cluster_PSCConfig{}
	out.PSCEnabled = direct.LazyPtr(in.GetPscEnabled())
	// MISSING: ServiceOwnedProjectNumber
	return out
}
func Cluster_PSCConfig_ToProto(mapCtx *direct.MapContext, in *krm.Cluster_PSCConfig) *pb.Cluster_PscConfig {
	if in == nil {
		return nil
	}
	out := &pb.Cluster_PscConfig{}
	out.PscEnabled = direct.ValueOf(in.PSCEnabled)
	// MISSING: ServiceOwnedProjectNumber
	return out
}
func Cluster_PSCConfigObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Cluster_PscConfig) *krmalloydbv1alpha1.Cluster_PSCConfigObservedState {
	if in == nil {
		return nil
	}
	out := &krmalloydbv1alpha1.Cluster_PSCConfigObservedState{}
	// MISSING: PSCEnabled
	out.ServiceOwnedProjectNumber = direct.LazyPtr(in.GetServiceOwnedProjectNumber())
	return out
}
func Cluster_PSCConfigObservedState_ToProto(mapCtx *direct.MapContext, in *krmalloydbv1alpha1.Cluster_PSCConfigObservedState) *pb.Cluster_PscConfig {
	if in == nil {
		return nil
	}
	out := &pb.Cluster_PscConfig{}
	// MISSING: PSCEnabled
	out.ServiceOwnedProjectNumber = direct.ValueOf(in.ServiceOwnedProjectNumber)
	return out
}
func Cluster_PSCConfigObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Cluster_PscConfig) *krm.Cluster_PSCConfigObservedState {
	if in == nil {
		return nil
	}
	out := &krm.Cluster_PSCConfigObservedState{}
	// MISSING: PSCEnabled
	out.ServiceOwnedProjectNumber = direct.LazyPtr(in.GetServiceOwnedProjectNumber())
	return out
}
func Cluster_PSCConfigObservedState_ToProto(mapCtx *direct.MapContext, in *krm.Cluster_PSCConfigObservedState) *pb.Cluster_PscConfig {
	if in == nil {
		return nil
	}
	out := &pb.Cluster_PscConfig{}
	// MISSING: PSCEnabled
	out.ServiceOwnedProjectNumber = direct.ValueOf(in.ServiceOwnedProjectNumber)
	return out
}
func Cluster_PrimaryConfig_FromProto(mapCtx *direct.MapContext, in *pb.Cluster_PrimaryConfig) *krmalloydbv1alpha1.Cluster_PrimaryConfig {
	if in == nil {
		return nil
	}
	out := &krmalloydbv1alpha1.Cluster_PrimaryConfig{}
	// MISSING: SecondaryClusterNames
	return out
}
func Cluster_PrimaryConfig_ToProto(mapCtx *direct.MapContext, in *krmalloydbv1alpha1.Cluster_PrimaryConfig) *pb.Cluster_PrimaryConfig {
	if in == nil {
		return nil
	}
	out := &pb.Cluster_PrimaryConfig{}
	// MISSING: SecondaryClusterNames
	return out
}
func Cluster_PrimaryConfig_FromProto(mapCtx *direct.MapContext, in *pb.Cluster_PrimaryConfig) *krm.Cluster_PrimaryConfig {
	if in == nil {
		return nil
	}
	out := &krm.Cluster_PrimaryConfig{}
	// MISSING: SecondaryClusterNames
	return out
}
func Cluster_PrimaryConfig_ToProto(mapCtx *direct.MapContext, in *krm.Cluster_PrimaryConfig) *pb.Cluster_PrimaryConfig {
	if in == nil {
		return nil
	}
	out := &pb.Cluster_PrimaryConfig{}
	// MISSING: SecondaryClusterNames
	return out
}
func Cluster_PrimaryConfigObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Cluster_PrimaryConfig) *krmalloydbv1alpha1.Cluster_PrimaryConfigObservedState {
	if in == nil {
		return nil
	}
	out := &krmalloydbv1alpha1.Cluster_PrimaryConfigObservedState{}
	out.SecondaryClusterNames = in.SecondaryClusterNames
	return out
}
func Cluster_PrimaryConfigObservedState_ToProto(mapCtx *direct.MapContext, in *krmalloydbv1alpha1.Cluster_PrimaryConfigObservedState) *pb.Cluster_PrimaryConfig {
	if in == nil {
		return nil
	}
	out := &pb.Cluster_PrimaryConfig{}
	out.SecondaryClusterNames = in.SecondaryClusterNames
	return out
}
func Cluster_PrimaryConfigObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Cluster_PrimaryConfig) *krm.Cluster_PrimaryConfigObservedState {
	if in == nil {
		return nil
	}
	out := &krm.Cluster_PrimaryConfigObservedState{}
	out.SecondaryClusterNames = in.SecondaryClusterNames
	return out
}
func Cluster_PrimaryConfigObservedState_ToProto(mapCtx *direct.MapContext, in *krm.Cluster_PrimaryConfigObservedState) *pb.Cluster_PrimaryConfig {
	if in == nil {
		return nil
	}
	out := &pb.Cluster_PrimaryConfig{}
	out.SecondaryClusterNames = in.SecondaryClusterNames
	return out
}
func ContinuousBackupConfig_FromProto(mapCtx *direct.MapContext, in *pb.ContinuousBackupConfig) *krmalloydbv1alpha1.ContinuousBackupConfig {
	if in == nil {
		return nil
	}
	out := &krmalloydbv1alpha1.ContinuousBackupConfig{}
	out.Enabled = in.Enabled
	out.RecoveryWindowDays = direct.LazyPtr(in.GetRecoveryWindowDays())
	out.EncryptionConfig = EncryptionConfig_FromProto(mapCtx, in.GetEncryptionConfig())
	return out
}
func ContinuousBackupConfig_ToProto(mapCtx *direct.MapContext, in *krmalloydbv1alpha1.ContinuousBackupConfig) *pb.ContinuousBackupConfig {
	if in == nil {
		return nil
	}
	out := &pb.ContinuousBackupConfig{}
	out.Enabled = in.Enabled
	out.RecoveryWindowDays = direct.ValueOf(in.RecoveryWindowDays)
	out.EncryptionConfig = EncryptionConfig_ToProto(mapCtx, in.EncryptionConfig)
	return out
}
func ContinuousBackupConfig_FromProto(mapCtx *direct.MapContext, in *pb.ContinuousBackupConfig) *krm.ContinuousBackupConfig {
	if in == nil {
		return nil
	}
	out := &krm.ContinuousBackupConfig{}
	out.Enabled = in.Enabled
	out.RecoveryWindowDays = direct.LazyPtr(in.GetRecoveryWindowDays())
	out.EncryptionConfig = EncryptionConfig_FromProto(mapCtx, in.GetEncryptionConfig())
	return out
}
func ContinuousBackupConfig_ToProto(mapCtx *direct.MapContext, in *krm.ContinuousBackupConfig) *pb.ContinuousBackupConfig {
	if in == nil {
		return nil
	}
	out := &pb.ContinuousBackupConfig{}
	out.Enabled = in.Enabled
	out.RecoveryWindowDays = direct.ValueOf(in.RecoveryWindowDays)
	out.EncryptionConfig = EncryptionConfig_ToProto(mapCtx, in.EncryptionConfig)
	return out
}
func ContinuousBackupInfo_FromProto(mapCtx *direct.MapContext, in *pb.ContinuousBackupInfo) *krmalloydbv1alpha1.ContinuousBackupInfo {
	if in == nil {
		return nil
	}
	out := &krmalloydbv1alpha1.ContinuousBackupInfo{}
	// MISSING: EncryptionInfo
	// MISSING: EnabledTime
	// MISSING: Schedule
	// MISSING: EarliestRestorableTime
	return out
}
func ContinuousBackupInfo_ToProto(mapCtx *direct.MapContext, in *krmalloydbv1alpha1.ContinuousBackupInfo) *pb.ContinuousBackupInfo {
	if in == nil {
		return nil
	}
	out := &pb.ContinuousBackupInfo{}
	// MISSING: EncryptionInfo
	// MISSING: EnabledTime
	// MISSING: Schedule
	// MISSING: EarliestRestorableTime
	return out
}
func EncryptionConfig_FromProto(mapCtx *direct.MapContext, in *pb.EncryptionConfig) *krmalloydbv1alpha1.EncryptionConfig {
	if in == nil {
		return nil
	}
	out := &krmalloydbv1alpha1.EncryptionConfig{}
	out.KMSKeyName = direct.LazyPtr(in.GetKmsKeyName())
	return out
}
func EncryptionConfig_ToProto(mapCtx *direct.MapContext, in *krmalloydbv1alpha1.EncryptionConfig) *pb.EncryptionConfig {
	if in == nil {
		return nil
	}
	out := &pb.EncryptionConfig{}
	out.KmsKeyName = direct.ValueOf(in.KMSKeyName)
	return out
}
func EncryptionConfig_FromProto(mapCtx *direct.MapContext, in *pb.EncryptionConfig) *krm.EncryptionConfig {
	if in == nil {
		return nil
	}
	out := &krm.EncryptionConfig{}
	if in.GetKmsKeyName() != "" {
		out.KMSKeyNameRef = &refsv1beta1.KMSCryptoKeyRef{External: in.GetKmsKeyName()}
	}
	return out
}
func EncryptionConfig_ToProto(mapCtx *direct.MapContext, in *krm.EncryptionConfig) *pb.EncryptionConfig {
	if in == nil {
		return nil
	}
	out := &pb.EncryptionConfig{}
	if in.KMSKeyNameRef != nil {
		out.KmsKeyName = in.KMSKeyNameRef.External
	}
	return out
}
func EncryptionInfo_FromProto(mapCtx *direct.MapContext, in *pb.EncryptionInfo) *krmalloydbv1alpha1.EncryptionInfo {
	if in == nil {
		return nil
	}
	out := &krmalloydbv1alpha1.EncryptionInfo{}
	// MISSING: EncryptionType
	// MISSING: KMSKeyVersions
	return out
}
func EncryptionInfo_ToProto(mapCtx *direct.MapContext, in *krmalloydbv1alpha1.EncryptionInfo) *pb.EncryptionInfo {
	if in == nil {
		return nil
	}
	out := &pb.EncryptionInfo{}
	// MISSING: EncryptionType
	// MISSING: KMSKeyVersions
	return out
}
func EncryptionInfo_FromProto(mapCtx *direct.MapContext, in *pb.EncryptionInfo) *krm.EncryptionInfo {
	if in == nil {
		return nil
	}
	out := &krm.EncryptionInfo{}
	// MISSING: EncryptionType
	// MISSING: KMSKeyVersions
	return out
}
func EncryptionInfo_ToProto(mapCtx *direct.MapContext, in *krm.EncryptionInfo) *pb.EncryptionInfo {
	if in == nil {
		return nil
	}
	out := &pb.EncryptionInfo{}
	// MISSING: EncryptionType
	// MISSING: KMSKeyVersions
	return out
}
func EncryptionInfoObservedState_FromProto(mapCtx *direct.MapContext, in *pb.EncryptionInfo) *krmalloydbv1alpha1.EncryptionInfoObservedState {
	if in == nil {
		return nil
	}
	out := &krmalloydbv1alpha1.EncryptionInfoObservedState{}
	out.EncryptionType = direct.Enum_FromProto(mapCtx, in.GetEncryptionType())
	out.KMSKeyVersions = in.KmsKeyVersions
	return out
}
func EncryptionInfoObservedState_ToProto(mapCtx *direct.MapContext, in *krmalloydbv1alpha1.EncryptionInfoObservedState) *pb.EncryptionInfo {
	if in == nil {
		return nil
	}
	out := &pb.EncryptionInfo{}
	out.EncryptionType = direct.Enum_ToProto[pb.EncryptionInfo_Type](mapCtx, in.EncryptionType)
	out.KmsKeyVersions = in.KMSKeyVersions
	return out
}
func EncryptionInfoObservedState_FromProto(mapCtx *direct.MapContext, in *pb.EncryptionInfo) *krm.EncryptionInfoObservedState {
	if in == nil {
		return nil
	}
	out := &krm.EncryptionInfoObservedState{}
	out.EncryptionType = direct.Enum_FromProto(mapCtx, in.GetEncryptionType())
	out.KMSKeyVersions = in.KmsKeyVersions
	return out
}
func EncryptionInfoObservedState_ToProto(mapCtx *direct.MapContext, in *krm.EncryptionInfoObservedState) *pb.EncryptionInfo {
	if in == nil {
		return nil
	}
	out := &pb.EncryptionInfo{}
	out.EncryptionType = direct.Enum_ToProto[pb.EncryptionInfo_Type](mapCtx, in.EncryptionType)
	out.KmsKeyVersions = in.KMSKeyVersions
	return out
}
func GcaInstanceConfig_FromProto(mapCtx *direct.MapContext, in *pb.GCAInstanceConfig) *krmalloydbv1alpha1.GcaInstanceConfig {
	if in == nil {
		return nil
	}
	out := &krmalloydbv1alpha1.GcaInstanceConfig{}
	// MISSING: GcaEntitlement
	return out
}
func GcaInstanceConfig_ToProto(mapCtx *direct.MapContext, in *krmalloydbv1alpha1.GcaInstanceConfig) *pb.GCAInstanceConfig {
	if in == nil {
		return nil
	}
	out := &pb.GCAInstanceConfig{}
	// MISSING: GcaEntitlement
	return out
}
func GcaInstanceConfig_FromProto(mapCtx *direct.MapContext, in *pb.GCAInstanceConfig) *krm.GcaInstanceConfig {
	if in == nil {
		return nil
	}
	out := &krm.GcaInstanceConfig{}
	// MISSING: GcaEntitlement
	return out
}
func GcaInstanceConfig_ToProto(mapCtx *direct.MapContext, in *krm.GcaInstanceConfig) *pb.GCAInstanceConfig {
	if in == nil {
		return nil
	}
	out := &pb.GCAInstanceConfig{}
	// MISSING: GcaEntitlement
	return out
}
func GcaInstanceConfigObservedState_FromProto(mapCtx *direct.MapContext, in *pb.GCAInstanceConfig) *krmalloydbv1alpha1.GcaInstanceConfigObservedState {
	if in == nil {
		return nil
	}
	out := &krmalloydbv1alpha1.GcaInstanceConfigObservedState{}
	out.GcaEntitlement = direct.Enum_FromProto(mapCtx, in.GetGcaEntitlement())
	return out
}
func GcaInstanceConfigObservedState_ToProto(mapCtx *direct.MapContext, in *krmalloydbv1alpha1.GcaInstanceConfigObservedState) *pb.GCAInstanceConfig {
	if in == nil {
		return nil
	}
	out := &pb.GCAInstanceConfig{}
	out.GcaEntitlement = direct.Enum_ToProto[pb.GCAEntitlementType](mapCtx, in.GcaEntitlement)
	return out
}
func GcaInstanceConfigObservedState_FromProto(mapCtx *direct.MapContext, in *pb.GCAInstanceConfig) *krm.GcaInstanceConfigObservedState {
	if in == nil {
		return nil
	}
	out := &krm.GcaInstanceConfigObservedState{}
	out.GcaEntitlement = direct.Enum_FromProto(mapCtx, in.GetGcaEntitlement())
	return out
}
func GcaInstanceConfigObservedState_ToProto(mapCtx *direct.MapContext, in *krm.GcaInstanceConfigObservedState) *pb.GCAInstanceConfig {
	if in == nil {
		return nil
	}
	out := &pb.GCAInstanceConfig{}
	out.GcaEntitlement = direct.Enum_ToProto[pb.GCAEntitlementType](mapCtx, in.GcaEntitlement)
	return out
}
func GeminiClusterConfig_FromProto(mapCtx *direct.MapContext, in *pb.GeminiClusterConfig) *krmalloydbv1alpha1.GeminiClusterConfig {
	if in == nil {
		return nil
	}
	out := &krmalloydbv1alpha1.GeminiClusterConfig{}
	// MISSING: Entitled
	return out
}
func GeminiClusterConfig_ToProto(mapCtx *direct.MapContext, in *krmalloydbv1alpha1.GeminiClusterConfig) *pb.GeminiClusterConfig {
	if in == nil {
		return nil
	}
	out := &pb.GeminiClusterConfig{}
	// MISSING: Entitled
	return out
}
func GeminiClusterConfig_FromProto(mapCtx *direct.MapContext, in *pb.GeminiClusterConfig) *krm.GeminiClusterConfig {
	if in == nil {
		return nil
	}
	out := &krm.GeminiClusterConfig{}
	// MISSING: Entitled
	return out
}
func GeminiClusterConfig_ToProto(mapCtx *direct.MapContext, in *krm.GeminiClusterConfig) *pb.GeminiClusterConfig {
	if in == nil {
		return nil
	}
	out := &pb.GeminiClusterConfig{}
	// MISSING: Entitled
	return out
}
func GeminiClusterConfigObservedState_FromProto(mapCtx *direct.MapContext, in *pb.GeminiClusterConfig) *krmalloydbv1alpha1.GeminiClusterConfigObservedState {
	if in == nil {
		return nil
	}
	out := &krmalloydbv1alpha1.GeminiClusterConfigObservedState{}
	out.Entitled = direct.LazyPtr(in.GetEntitled())
	return out
}
func GeminiClusterConfigObservedState_ToProto(mapCtx *direct.MapContext, in *krmalloydbv1alpha1.GeminiClusterConfigObservedState) *pb.GeminiClusterConfig {
	if in == nil {
		return nil
	}
	out := &pb.GeminiClusterConfig{}
	out.Entitled = direct.ValueOf(in.Entitled)
	return out
}
func GeminiClusterConfigObservedState_FromProto(mapCtx *direct.MapContext, in *pb.GeminiClusterConfig) *krm.GeminiClusterConfigObservedState {
	if in == nil {
		return nil
	}
	out := &krm.GeminiClusterConfigObservedState{}
	out.Entitled = direct.LazyPtr(in.GetEntitled())
	return out
}
func GeminiClusterConfigObservedState_ToProto(mapCtx *direct.MapContext, in *krm.GeminiClusterConfigObservedState) *pb.GeminiClusterConfig {
	if in == nil {
		return nil
	}
	out := &pb.GeminiClusterConfig{}
	out.Entitled = direct.ValueOf(in.Entitled)
	return out
}
func GeminiInstanceConfig_FromProto(mapCtx *direct.MapContext, in *pb.GeminiInstanceConfig) *krmalloydbv1alpha1.GeminiInstanceConfig {
	if in == nil {
		return nil
	}
	out := &krmalloydbv1alpha1.GeminiInstanceConfig{}
	// MISSING: Entitled
	return out
}
func GeminiInstanceConfig_ToProto(mapCtx *direct.MapContext, in *krmalloydbv1alpha1.GeminiInstanceConfig) *pb.GeminiInstanceConfig {
	if in == nil {
		return nil
	}
	out := &pb.GeminiInstanceConfig{}
	// MISSING: Entitled
	return out
}
func GeminiInstanceConfig_FromProto(mapCtx *direct.MapContext, in *pb.GeminiInstanceConfig) *krm.GeminiInstanceConfig {
	if in == nil {
		return nil
	}
	out := &krm.GeminiInstanceConfig{}
	// MISSING: Entitled
	return out
}
func GeminiInstanceConfig_ToProto(mapCtx *direct.MapContext, in *krm.GeminiInstanceConfig) *pb.GeminiInstanceConfig {
	if in == nil {
		return nil
	}
	out := &pb.GeminiInstanceConfig{}
	// MISSING: Entitled
	return out
}
func GeminiInstanceConfigObservedState_FromProto(mapCtx *direct.MapContext, in *pb.GeminiInstanceConfig) *krmalloydbv1alpha1.GeminiInstanceConfigObservedState {
	if in == nil {
		return nil
	}
	out := &krmalloydbv1alpha1.GeminiInstanceConfigObservedState{}
	out.Entitled = direct.LazyPtr(in.GetEntitled())
	return out
}
func GeminiInstanceConfigObservedState_ToProto(mapCtx *direct.MapContext, in *krmalloydbv1alpha1.GeminiInstanceConfigObservedState) *pb.GeminiInstanceConfig {
	if in == nil {
		return nil
	}
	out := &pb.GeminiInstanceConfig{}
	out.Entitled = direct.ValueOf(in.Entitled)
	return out
}
func GeminiInstanceConfigObservedState_FromProto(mapCtx *direct.MapContext, in *pb.GeminiInstanceConfig) *krm.GeminiInstanceConfigObservedState {
	if in == nil {
		return nil
	}
	out := &krm.GeminiInstanceConfigObservedState{}
	out.Entitled = direct.LazyPtr(in.GetEntitled())
	return out
}
func GeminiInstanceConfigObservedState_ToProto(mapCtx *direct.MapContext, in *krm.GeminiInstanceConfigObservedState) *pb.GeminiInstanceConfig {
	if in == nil {
		return nil
	}
	out := &pb.GeminiInstanceConfig{}
	out.Entitled = direct.ValueOf(in.Entitled)
	return out
}
func Instance_FromProto(mapCtx *direct.MapContext, in *pb.Instance) *krmalloydbv1alpha1.Instance {
	if in == nil {
		return nil
	}
	out := &krmalloydbv1alpha1.Instance{}
	// MISSING: Name
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	// MISSING: Uid
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: DeleteTime
	out.Labels = in.Labels
	// MISSING: State
	out.InstanceType = direct.Enum_FromProto(mapCtx, in.GetInstanceType())
	out.MachineConfig = Instance_MachineConfig_FromProto(mapCtx, in.GetMachineConfig())
	out.AvailabilityType = direct.Enum_FromProto(mapCtx, in.GetAvailabilityType())
	out.GCEZone = direct.LazyPtr(in.GetGceZone())
	out.DatabaseFlags = in.DatabaseFlags
	// MISSING: WritableNode
	// MISSING: Nodes
	out.QueryInsightsConfig = Instance_QueryInsightsInstanceConfig_FromProto(mapCtx, in.GetQueryInsightsConfig())
	out.ObservabilityConfig = Instance_ObservabilityInstanceConfig_FromProto(mapCtx, in.GetObservabilityConfig())
	out.ReadPoolConfig = Instance_ReadPoolConfig_FromProto(mapCtx, in.GetReadPoolConfig())
	// MISSING: IPAddress
	// MISSING: PublicIPAddress
	// MISSING: Reconciling
	out.Etag = direct.LazyPtr(in.GetEtag())
	out.Annotations = in.Annotations
	out.UpdatePolicy = Instance_UpdatePolicy_FromProto(mapCtx, in.GetUpdatePolicy())
	out.ClientConnectionConfig = Instance_ClientConnectionConfig_FromProto(mapCtx, in.GetClientConnectionConfig())
	// MISSING: SatisfiesPzs
	out.PSCInstanceConfig = Instance_PSCInstanceConfig_FromProto(mapCtx, in.GetPscInstanceConfig())
	out.NetworkConfig = Instance_InstanceNetworkConfig_FromProto(mapCtx, in.GetNetworkConfig())
	out.GeminiConfig = GeminiInstanceConfig_FromProto(mapCtx, in.GetGeminiConfig())
	// MISSING: OutboundPublicIPAddresses
	out.ActivationPolicy = direct.Enum_FromProto(mapCtx, in.GetActivationPolicy())
	out.ConnectionPoolConfig = Instance_ConnectionPoolConfig_FromProto(mapCtx, in.GetConnectionPoolConfig())
	// MISSING: GcaConfig
	return out
}
func Instance_ToProto(mapCtx *direct.MapContext, in *krmalloydbv1alpha1.Instance) *pb.Instance {
	if in == nil {
		return nil
	}
	out := &pb.Instance{}
	// MISSING: Name
	out.DisplayName = direct.ValueOf(in.DisplayName)
	// MISSING: Uid
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: DeleteTime
	out.Labels = in.Labels
	// MISSING: State
	out.InstanceType = direct.Enum_ToProto[pb.Instance_InstanceType](mapCtx, in.InstanceType)
	out.MachineConfig = Instance_MachineConfig_ToProto(mapCtx, in.MachineConfig)
	out.AvailabilityType = direct.Enum_ToProto[pb.Instance_AvailabilityType](mapCtx, in.AvailabilityType)
	out.GceZone = direct.ValueOf(in.GCEZone)
	out.DatabaseFlags = in.DatabaseFlags
	// MISSING: WritableNode
	// MISSING: Nodes
	out.QueryInsightsConfig = Instance_QueryInsightsInstanceConfig_ToProto(mapCtx, in.QueryInsightsConfig)
	out.ObservabilityConfig = Instance_ObservabilityInstanceConfig_ToProto(mapCtx, in.ObservabilityConfig)
	out.ReadPoolConfig = Instance_ReadPoolConfig_ToProto(mapCtx, in.ReadPoolConfig)
	// MISSING: IPAddress
	// MISSING: PublicIPAddress
	// MISSING: Reconciling
	out.Etag = direct.ValueOf(in.Etag)
	out.Annotations = in.Annotations
	out.UpdatePolicy = Instance_UpdatePolicy_ToProto(mapCtx, in.UpdatePolicy)
	out.ClientConnectionConfig = Instance_ClientConnectionConfig_ToProto(mapCtx, in.ClientConnectionConfig)
	// MISSING: SatisfiesPzs
	out.PscInstanceConfig = Instance_PSCInstanceConfig_ToProto(mapCtx, in.PSCInstanceConfig)
	out.NetworkConfig = Instance_InstanceNetworkConfig_ToProto(mapCtx, in.NetworkConfig)
	out.GeminiConfig = GeminiInstanceConfig_ToProto(mapCtx, in.GeminiConfig)
	// MISSING: OutboundPublicIPAddresses
	out.ActivationPolicy = direct.Enum_ToProto[pb.Instance_ActivationPolicy](mapCtx, in.ActivationPolicy)
	out.ConnectionPoolConfig = Instance_ConnectionPoolConfig_ToProto(mapCtx, in.ConnectionPoolConfig)
	// MISSING: GcaConfig
	return out
}
func InstanceObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Instance) *krmalloydbv1alpha1.InstanceObservedState {
	if in == nil {
		return nil
	}
	out := &krmalloydbv1alpha1.InstanceObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: DisplayName
	out.Uid = direct.LazyPtr(in.GetUid())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.DeleteTime = direct.StringTimestamp_FromProto(mapCtx, in.GetDeleteTime())
	// MISSING: Labels
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	// MISSING: InstanceType
	// MISSING: MachineConfig
	// MISSING: AvailabilityType
	// MISSING: GCEZone
	// MISSING: DatabaseFlags
	out.WritableNode = Instance_Node_FromProto(mapCtx, in.GetWritableNode())
	out.Nodes = direct.Slice_FromProto(mapCtx, in.Nodes, Instance_Node_FromProto)
	// MISSING: QueryInsightsConfig
	out.ObservabilityConfig = Instance_ObservabilityInstanceConfigObservedState_FromProto(mapCtx, in.GetObservabilityConfig())
	// MISSING: ReadPoolConfig
	out.IPAddress = direct.LazyPtr(in.GetIpAddress())
	out.PublicIPAddress = direct.LazyPtr(in.GetPublicIpAddress())
	out.Reconciling = direct.LazyPtr(in.GetReconciling())
	// MISSING: Etag
	// MISSING: Annotations
	// MISSING: UpdatePolicy
	// MISSING: ClientConnectionConfig
	out.SatisfiesPzs = direct.LazyPtr(in.GetSatisfiesPzs())
	out.PSCInstanceConfig = Instance_PSCInstanceConfigObservedState_FromProto(mapCtx, in.GetPscInstanceConfig())
	out.NetworkConfig = Instance_InstanceNetworkConfigObservedState_FromProto(mapCtx, in.GetNetworkConfig())
	out.GeminiConfig = GeminiInstanceConfigObservedState_FromProto(mapCtx, in.GetGeminiConfig())
	out.OutboundPublicIPAddresses = in.OutboundPublicIpAddresses
	// MISSING: ActivationPolicy
	// MISSING: ConnectionPoolConfig
	out.GcaConfig = GcaInstanceConfig_FromProto(mapCtx, in.GetGcaConfig())
	return out
}
func InstanceObservedState_ToProto(mapCtx *direct.MapContext, in *krmalloydbv1alpha1.InstanceObservedState) *pb.Instance {
	if in == nil {
		return nil
	}
	out := &pb.Instance{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: DisplayName
	out.Uid = direct.ValueOf(in.Uid)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.DeleteTime = direct.StringTimestamp_ToProto(mapCtx, in.DeleteTime)
	// MISSING: Labels
	out.State = direct.Enum_ToProto[pb.Instance_State](mapCtx, in.State)
	// MISSING: InstanceType
	// MISSING: MachineConfig
	// MISSING: AvailabilityType
	// MISSING: GCEZone
	// MISSING: DatabaseFlags
	out.WritableNode = Instance_Node_ToProto(mapCtx, in.WritableNode)
	out.Nodes = direct.Slice_ToProto(mapCtx, in.Nodes, Instance_Node_ToProto)
	// MISSING: QueryInsightsConfig
	out.ObservabilityConfig = Instance_ObservabilityInstanceConfigObservedState_ToProto(mapCtx, in.ObservabilityConfig)
	// MISSING: ReadPoolConfig
	out.IpAddress = direct.ValueOf(in.IPAddress)
	out.PublicIpAddress = direct.ValueOf(in.PublicIPAddress)
	out.Reconciling = direct.ValueOf(in.Reconciling)
	// MISSING: Etag
	// MISSING: Annotations
	// MISSING: UpdatePolicy
	// MISSING: ClientConnectionConfig
	out.SatisfiesPzs = direct.ValueOf(in.SatisfiesPzs)
	out.PscInstanceConfig = Instance_PSCInstanceConfigObservedState_ToProto(mapCtx, in.PSCInstanceConfig)
	out.NetworkConfig = Instance_InstanceNetworkConfigObservedState_ToProto(mapCtx, in.NetworkConfig)
	out.GeminiConfig = GeminiInstanceConfigObservedState_ToProto(mapCtx, in.GeminiConfig)
	out.OutboundPublicIpAddresses = in.OutboundPublicIPAddresses
	// MISSING: ActivationPolicy
	// MISSING: ConnectionPoolConfig
	out.GcaConfig = GcaInstanceConfig_ToProto(mapCtx, in.GcaConfig)
	return out
}
func Instance_ClientConnectionConfig_FromProto(mapCtx *direct.MapContext, in *pb.Instance_ClientConnectionConfig) *krmalloydbv1alpha1.Instance_ClientConnectionConfig {
	if in == nil {
		return nil
	}
	out := &krmalloydbv1alpha1.Instance_ClientConnectionConfig{}
	out.RequireConnectors = direct.LazyPtr(in.GetRequireConnectors())
	out.SSLConfig = SSLConfig_FromProto(mapCtx, in.GetSslConfig())
	return out
}
func Instance_ClientConnectionConfig_ToProto(mapCtx *direct.MapContext, in *krmalloydbv1alpha1.Instance_ClientConnectionConfig) *pb.Instance_ClientConnectionConfig {
	if in == nil {
		return nil
	}
	out := &pb.Instance_ClientConnectionConfig{}
	out.RequireConnectors = direct.ValueOf(in.RequireConnectors)
	out.SslConfig = SSLConfig_ToProto(mapCtx, in.SSLConfig)
	return out
}
func Instance_ClientConnectionConfig_FromProto(mapCtx *direct.MapContext, in *pb.Instance_ClientConnectionConfig) *krm.Instance_ClientConnectionConfig {
	if in == nil {
		return nil
	}
	out := &krm.Instance_ClientConnectionConfig{}
	out.RequireConnectors = direct.LazyPtr(in.GetRequireConnectors())
	out.SSLConfig = SSLConfig_FromProto(mapCtx, in.GetSslConfig())
	return out
}
func Instance_ClientConnectionConfig_ToProto(mapCtx *direct.MapContext, in *krm.Instance_ClientConnectionConfig) *pb.Instance_ClientConnectionConfig {
	if in == nil {
		return nil
	}
	out := &pb.Instance_ClientConnectionConfig{}
	out.RequireConnectors = direct.ValueOf(in.RequireConnectors)
	out.SslConfig = SSLConfig_ToProto(mapCtx, in.SSLConfig)
	return out
}
func Instance_ConnectionPoolConfig_FromProto(mapCtx *direct.MapContext, in *pb.Instance_ConnectionPoolConfig) *krmalloydbv1alpha1.Instance_ConnectionPoolConfig {
	if in == nil {
		return nil
	}
	out := &krmalloydbv1alpha1.Instance_ConnectionPoolConfig{}
	out.Enabled = direct.LazyPtr(in.GetEnabled())
	out.Flags = in.Flags
	return out
}
func Instance_ConnectionPoolConfig_ToProto(mapCtx *direct.MapContext, in *krmalloydbv1alpha1.Instance_ConnectionPoolConfig) *pb.Instance_ConnectionPoolConfig {
	if in == nil {
		return nil
	}
	out := &pb.Instance_ConnectionPoolConfig{}
	out.Enabled = direct.ValueOf(in.Enabled)
	out.Flags = in.Flags
	return out
}
func Instance_ConnectionPoolConfig_FromProto(mapCtx *direct.MapContext, in *pb.Instance_ConnectionPoolConfig) *krm.Instance_ConnectionPoolConfig {
	if in == nil {
		return nil
	}
	out := &krm.Instance_ConnectionPoolConfig{}
	out.Enabled = direct.LazyPtr(in.GetEnabled())
	out.Flags = in.Flags
	return out
}
func Instance_ConnectionPoolConfig_ToProto(mapCtx *direct.MapContext, in *krm.Instance_ConnectionPoolConfig) *pb.Instance_ConnectionPoolConfig {
	if in == nil {
		return nil
	}
	out := &pb.Instance_ConnectionPoolConfig{}
	out.Enabled = direct.ValueOf(in.Enabled)
	out.Flags = in.Flags
	return out
}
func Instance_InstanceNetworkConfig_FromProto(mapCtx *direct.MapContext, in *pb.Instance_InstanceNetworkConfig) *krmalloydbv1alpha1.Instance_InstanceNetworkConfig {
	if in == nil {
		return nil
	}
	out := &krmalloydbv1alpha1.Instance_InstanceNetworkConfig{}
	out.AuthorizedExternalNetworks = direct.Slice_FromProto(mapCtx, in.AuthorizedExternalNetworks, Instance_InstanceNetworkConfig_AuthorizedNetwork_FromProto)
	out.EnablePublicIP = direct.LazyPtr(in.GetEnablePublicIp())
	out.EnableOutboundPublicIP = direct.LazyPtr(in.GetEnableOutboundPublicIp())
	// MISSING: Network
	out.AllocatedIPRangeOverride = direct.LazyPtr(in.GetAllocatedIpRangeOverride())
	return out
}
func Instance_InstanceNetworkConfig_ToProto(mapCtx *direct.MapContext, in *krmalloydbv1alpha1.Instance_InstanceNetworkConfig) *pb.Instance_InstanceNetworkConfig {
	if in == nil {
		return nil
	}
	out := &pb.Instance_InstanceNetworkConfig{}
	out.AuthorizedExternalNetworks = direct.Slice_ToProto(mapCtx, in.AuthorizedExternalNetworks, Instance_InstanceNetworkConfig_AuthorizedNetwork_ToProto)
	out.EnablePublicIp = direct.ValueOf(in.EnablePublicIP)
	out.EnableOutboundPublicIp = direct.ValueOf(in.EnableOutboundPublicIP)
	// MISSING: Network
	out.AllocatedIpRangeOverride = direct.ValueOf(in.AllocatedIPRangeOverride)
	return out
}
func Instance_InstanceNetworkConfig_FromProto(mapCtx *direct.MapContext, in *pb.Instance_InstanceNetworkConfig) *krm.Instance_InstanceNetworkConfig {
	if in == nil {
		return nil
	}
	out := &krm.Instance_InstanceNetworkConfig{}
	out.AuthorizedExternalNetworks = direct.Slice_FromProto(mapCtx, in.AuthorizedExternalNetworks, Instance_InstanceNetworkConfig_AuthorizedNetwork_FromProto)
	out.EnablePublicIP = direct.LazyPtr(in.GetEnablePublicIp())
	out.EnableOutboundPublicIP = direct.LazyPtr(in.GetEnableOutboundPublicIp())
	// MISSING: Network
	// MISSING: AllocatedIPRangeOverride
	return out
}
func Instance_InstanceNetworkConfig_ToProto(mapCtx *direct.MapContext, in *krm.Instance_InstanceNetworkConfig) *pb.Instance_InstanceNetworkConfig {
	if in == nil {
		return nil
	}
	out := &pb.Instance_InstanceNetworkConfig{}
	out.AuthorizedExternalNetworks = direct.Slice_ToProto(mapCtx, in.AuthorizedExternalNetworks, Instance_InstanceNetworkConfig_AuthorizedNetwork_ToProto)
	out.EnablePublicIp = direct.ValueOf(in.EnablePublicIP)
	out.EnableOutboundPublicIp = direct.ValueOf(in.EnableOutboundPublicIP)
	// MISSING: Network
	// MISSING: AllocatedIPRangeOverride
	return out
}
func Instance_InstanceNetworkConfigObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Instance_InstanceNetworkConfig) *krmalloydbv1alpha1.Instance_InstanceNetworkConfigObservedState {
	if in == nil {
		return nil
	}
	out := &krmalloydbv1alpha1.Instance_InstanceNetworkConfigObservedState{}
	// MISSING: AuthorizedExternalNetworks
	// MISSING: EnablePublicIP
	// MISSING: EnableOutboundPublicIP
	out.Network = direct.LazyPtr(in.GetNetwork())
	// MISSING: AllocatedIPRangeOverride
	return out
}
func Instance_InstanceNetworkConfigObservedState_ToProto(mapCtx *direct.MapContext, in *krmalloydbv1alpha1.Instance_InstanceNetworkConfigObservedState) *pb.Instance_InstanceNetworkConfig {
	if in == nil {
		return nil
	}
	out := &pb.Instance_InstanceNetworkConfig{}
	// MISSING: AuthorizedExternalNetworks
	// MISSING: EnablePublicIP
	// MISSING: EnableOutboundPublicIP
	out.Network = direct.ValueOf(in.Network)
	// MISSING: AllocatedIPRangeOverride
	return out
}
func Instance_InstanceNetworkConfig_AuthorizedNetwork_FromProto(mapCtx *direct.MapContext, in *pb.Instance_InstanceNetworkConfig_AuthorizedNetwork) *krmalloydbv1alpha1.Instance_InstanceNetworkConfig_AuthorizedNetwork {
	if in == nil {
		return nil
	}
	out := &krmalloydbv1alpha1.Instance_InstanceNetworkConfig_AuthorizedNetwork{}
	out.CIDRRange = direct.LazyPtr(in.GetCidrRange())
	return out
}
func Instance_InstanceNetworkConfig_AuthorizedNetwork_ToProto(mapCtx *direct.MapContext, in *krmalloydbv1alpha1.Instance_InstanceNetworkConfig_AuthorizedNetwork) *pb.Instance_InstanceNetworkConfig_AuthorizedNetwork {
	if in == nil {
		return nil
	}
	out := &pb.Instance_InstanceNetworkConfig_AuthorizedNetwork{}
	out.CidrRange = direct.ValueOf(in.CIDRRange)
	return out
}
func Instance_InstanceNetworkConfig_AuthorizedNetwork_FromProto(mapCtx *direct.MapContext, in *pb.Instance_InstanceNetworkConfig_AuthorizedNetwork) *krm.Instance_InstanceNetworkConfig_AuthorizedNetwork {
	if in == nil {
		return nil
	}
	out := &krm.Instance_InstanceNetworkConfig_AuthorizedNetwork{}
	out.CIDRRange = direct.LazyPtr(in.GetCidrRange())
	return out
}
func Instance_InstanceNetworkConfig_AuthorizedNetwork_ToProto(mapCtx *direct.MapContext, in *krm.Instance_InstanceNetworkConfig_AuthorizedNetwork) *pb.Instance_InstanceNetworkConfig_AuthorizedNetwork {
	if in == nil {
		return nil
	}
	out := &pb.Instance_InstanceNetworkConfig_AuthorizedNetwork{}
	out.CidrRange = direct.ValueOf(in.CIDRRange)
	return out
}
func Instance_MachineConfig_FromProto(mapCtx *direct.MapContext, in *pb.Instance_MachineConfig) *krmalloydbv1alpha1.Instance_MachineConfig {
	if in == nil {
		return nil
	}
	out := &krmalloydbv1alpha1.Instance_MachineConfig{}
	out.CPUCount = direct.LazyPtr(in.GetCpuCount())
	out.MachineType = direct.LazyPtr(in.GetMachineType())
	return out
}
func Instance_MachineConfig_ToProto(mapCtx *direct.MapContext, in *krmalloydbv1alpha1.Instance_MachineConfig) *pb.Instance_MachineConfig {
	if in == nil {
		return nil
	}
	out := &pb.Instance_MachineConfig{}
	out.CpuCount = direct.ValueOf(in.CPUCount)
	out.MachineType = direct.ValueOf(in.MachineType)
	return out
}
func Instance_MachineConfig_FromProto(mapCtx *direct.MapContext, in *pb.Instance_MachineConfig) *krm.Instance_MachineConfig {
	if in == nil {
		return nil
	}
	out := &krm.Instance_MachineConfig{}
	out.CPUCount = direct.LazyPtr(in.GetCpuCount())
	out.MachineType = direct.LazyPtr(in.GetMachineType())
	return out
}
func Instance_MachineConfig_ToProto(mapCtx *direct.MapContext, in *krm.Instance_MachineConfig) *pb.Instance_MachineConfig {
	if in == nil {
		return nil
	}
	out := &pb.Instance_MachineConfig{}
	out.CpuCount = direct.ValueOf(in.CPUCount)
	out.MachineType = direct.ValueOf(in.MachineType)
	return out
}
func Instance_Node_FromProto(mapCtx *direct.MapContext, in *pb.Instance_Node) *krmalloydbv1alpha1.Instance_Node {
	if in == nil {
		return nil
	}
	out := &krmalloydbv1alpha1.Instance_Node{}
	// MISSING: ZoneID
	// MISSING: ID
	// MISSING: IP
	// MISSING: State
	return out
}
func Instance_Node_ToProto(mapCtx *direct.MapContext, in *krmalloydbv1alpha1.Instance_Node) *pb.Instance_Node {
	if in == nil {
		return nil
	}
	out := &pb.Instance_Node{}
	// MISSING: ZoneID
	// MISSING: ID
	// MISSING: IP
	// MISSING: State
	return out
}
func Instance_Node_FromProto(mapCtx *direct.MapContext, in *pb.Instance_Node) *krm.Instance_Node {
	if in == nil {
		return nil
	}
	out := &krm.Instance_Node{}
	// MISSING: ZoneID
	// MISSING: ID
	// MISSING: IP
	// MISSING: State
	return out
}
func Instance_Node_ToProto(mapCtx *direct.MapContext, in *krm.Instance_Node) *pb.Instance_Node {
	if in == nil {
		return nil
	}
	out := &pb.Instance_Node{}
	// MISSING: ZoneID
	// MISSING: ID
	// MISSING: IP
	// MISSING: State
	return out
}
func Instance_NodeObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Instance_Node) *krmalloydbv1alpha1.Instance_NodeObservedState {
	if in == nil {
		return nil
	}
	out := &krmalloydbv1alpha1.Instance_NodeObservedState{}
	out.ZoneID = direct.LazyPtr(in.GetZoneId())
	out.ID = direct.LazyPtr(in.GetId())
	out.IP = direct.LazyPtr(in.GetIp())
	out.State = direct.LazyPtr(in.GetState())
	return out
}
func Instance_NodeObservedState_ToProto(mapCtx *direct.MapContext, in *krmalloydbv1alpha1.Instance_NodeObservedState) *pb.Instance_Node {
	if in == nil {
		return nil
	}
	out := &pb.Instance_Node{}
	out.ZoneId = direct.ValueOf(in.ZoneID)
	out.Id = direct.ValueOf(in.ID)
	out.Ip = direct.ValueOf(in.IP)
	out.State = direct.ValueOf(in.State)
	return out
}
func Instance_NodeObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Instance_Node) *krm.Instance_NodeObservedState {
	if in == nil {
		return nil
	}
	out := &krm.Instance_NodeObservedState{}
	out.ZoneID = direct.LazyPtr(in.GetZoneId())
	out.ID = direct.LazyPtr(in.GetId())
	out.IP = direct.LazyPtr(in.GetIp())
	out.State = direct.LazyPtr(in.GetState())
	return out
}
func Instance_NodeObservedState_ToProto(mapCtx *direct.MapContext, in *krm.Instance_NodeObservedState) *pb.Instance_Node {
	if in == nil {
		return nil
	}
	out := &pb.Instance_Node{}
	out.ZoneId = direct.ValueOf(in.ZoneID)
	out.Id = direct.ValueOf(in.ID)
	out.Ip = direct.ValueOf(in.IP)
	out.State = direct.ValueOf(in.State)
	return out
}
func Instance_ObservabilityInstanceConfig_FromProto(mapCtx *direct.MapContext, in *pb.Instance_ObservabilityInstanceConfig) *krmalloydbv1alpha1.Instance_ObservabilityInstanceConfig {
	if in == nil {
		return nil
	}
	out := &krmalloydbv1alpha1.Instance_ObservabilityInstanceConfig{}
	out.Enabled = in.Enabled
	out.PreserveComments = in.PreserveComments
	out.TrackWaitEvents = in.TrackWaitEvents
	// MISSING: TrackWaitEventTypes
	out.MaxQueryStringLength = in.MaxQueryStringLength
	out.RecordApplicationTags = in.RecordApplicationTags
	out.QueryPlansPerMinute = in.QueryPlansPerMinute
	out.TrackActiveQueries = in.TrackActiveQueries
	out.TrackClientAddress = in.TrackClientAddress
	out.AssistiveExperiencesEnabled = in.AssistiveExperiencesEnabled
	return out
}
func Instance_ObservabilityInstanceConfig_ToProto(mapCtx *direct.MapContext, in *krmalloydbv1alpha1.Instance_ObservabilityInstanceConfig) *pb.Instance_ObservabilityInstanceConfig {
	if in == nil {
		return nil
	}
	out := &pb.Instance_ObservabilityInstanceConfig{}
	out.Enabled = in.Enabled
	out.PreserveComments = in.PreserveComments
	out.TrackWaitEvents = in.TrackWaitEvents
	// MISSING: TrackWaitEventTypes
	out.MaxQueryStringLength = in.MaxQueryStringLength
	out.RecordApplicationTags = in.RecordApplicationTags
	out.QueryPlansPerMinute = in.QueryPlansPerMinute
	out.TrackActiveQueries = in.TrackActiveQueries
	out.TrackClientAddress = in.TrackClientAddress
	out.AssistiveExperiencesEnabled = in.AssistiveExperiencesEnabled
	return out
}
func Instance_ObservabilityInstanceConfig_FromProto(mapCtx *direct.MapContext, in *pb.Instance_ObservabilityInstanceConfig) *krm.Instance_ObservabilityInstanceConfig {
	if in == nil {
		return nil
	}
	out := &krm.Instance_ObservabilityInstanceConfig{}
	out.Enabled = in.Enabled
	out.PreserveComments = in.PreserveComments
	out.TrackWaitEvents = in.TrackWaitEvents
	// MISSING: TrackWaitEventTypes
	out.MaxQueryStringLength = in.MaxQueryStringLength
	out.RecordApplicationTags = in.RecordApplicationTags
	out.QueryPlansPerMinute = in.QueryPlansPerMinute
	out.TrackActiveQueries = in.TrackActiveQueries
	out.TrackClientAddress = in.TrackClientAddress
	out.AssistiveExperiencesEnabled = in.AssistiveExperiencesEnabled
	return out
}
func Instance_ObservabilityInstanceConfig_ToProto(mapCtx *direct.MapContext, in *krm.Instance_ObservabilityInstanceConfig) *pb.Instance_ObservabilityInstanceConfig {
	if in == nil {
		return nil
	}
	out := &pb.Instance_ObservabilityInstanceConfig{}
	out.Enabled = in.Enabled
	out.PreserveComments = in.PreserveComments
	out.TrackWaitEvents = in.TrackWaitEvents
	// MISSING: TrackWaitEventTypes
	out.MaxQueryStringLength = in.MaxQueryStringLength
	out.RecordApplicationTags = in.RecordApplicationTags
	out.QueryPlansPerMinute = in.QueryPlansPerMinute
	out.TrackActiveQueries = in.TrackActiveQueries
	out.TrackClientAddress = in.TrackClientAddress
	out.AssistiveExperiencesEnabled = in.AssistiveExperiencesEnabled
	return out
}
func Instance_ObservabilityInstanceConfigObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Instance_ObservabilityInstanceConfig) *krmalloydbv1alpha1.Instance_ObservabilityInstanceConfigObservedState {
	if in == nil {
		return nil
	}
	out := &krmalloydbv1alpha1.Instance_ObservabilityInstanceConfigObservedState{}
	// MISSING: Enabled
	// MISSING: PreserveComments
	// MISSING: TrackWaitEvents
	out.TrackWaitEventTypes = in.TrackWaitEventTypes
	// MISSING: MaxQueryStringLength
	// MISSING: RecordApplicationTags
	// MISSING: QueryPlansPerMinute
	// MISSING: TrackActiveQueries
	// MISSING: TrackClientAddress
	// MISSING: AssistiveExperiencesEnabled
	return out
}
func Instance_ObservabilityInstanceConfigObservedState_ToProto(mapCtx *direct.MapContext, in *krmalloydbv1alpha1.Instance_ObservabilityInstanceConfigObservedState) *pb.Instance_ObservabilityInstanceConfig {
	if in == nil {
		return nil
	}
	out := &pb.Instance_ObservabilityInstanceConfig{}
	// MISSING: Enabled
	// MISSING: PreserveComments
	// MISSING: TrackWaitEvents
	out.TrackWaitEventTypes = in.TrackWaitEventTypes
	// MISSING: MaxQueryStringLength
	// MISSING: RecordApplicationTags
	// MISSING: QueryPlansPerMinute
	// MISSING: TrackActiveQueries
	// MISSING: TrackClientAddress
	// MISSING: AssistiveExperiencesEnabled
	return out
}
func Instance_ObservabilityInstanceConfigObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Instance_ObservabilityInstanceConfig) *krm.Instance_ObservabilityInstanceConfigObservedState {
	if in == nil {
		return nil
	}
	out := &krm.Instance_ObservabilityInstanceConfigObservedState{}
	// MISSING: Enabled
	// MISSING: PreserveComments
	// MISSING: TrackWaitEvents
	out.TrackWaitEventTypes = in.TrackWaitEventTypes
	// MISSING: MaxQueryStringLength
	// MISSING: RecordApplicationTags
	// MISSING: QueryPlansPerMinute
	// MISSING: TrackActiveQueries
	// MISSING: TrackClientAddress
	// MISSING: AssistiveExperiencesEnabled
	return out
}
func Instance_ObservabilityInstanceConfigObservedState_ToProto(mapCtx *direct.MapContext, in *krm.Instance_ObservabilityInstanceConfigObservedState) *pb.Instance_ObservabilityInstanceConfig {
	if in == nil {
		return nil
	}
	out := &pb.Instance_ObservabilityInstanceConfig{}
	// MISSING: Enabled
	// MISSING: PreserveComments
	// MISSING: TrackWaitEvents
	out.TrackWaitEventTypes = in.TrackWaitEventTypes
	// MISSING: MaxQueryStringLength
	// MISSING: RecordApplicationTags
	// MISSING: QueryPlansPerMinute
	// MISSING: TrackActiveQueries
	// MISSING: TrackClientAddress
	// MISSING: AssistiveExperiencesEnabled
	return out
}
func Instance_PSCAutoConnectionConfig_FromProto(mapCtx *direct.MapContext, in *pb.Instance_PscAutoConnectionConfig) *krmalloydbv1alpha1.Instance_PSCAutoConnectionConfig {
	if in == nil {
		return nil
	}
	out := &krmalloydbv1alpha1.Instance_PSCAutoConnectionConfig{}
	out.ConsumerProject = direct.LazyPtr(in.GetConsumerProject())
	out.ConsumerNetwork = direct.LazyPtr(in.GetConsumerNetwork())
	// MISSING: IPAddress
	// MISSING: Status
	// MISSING: ConsumerNetworkStatus
	return out
}
func Instance_PSCAutoConnectionConfig_ToProto(mapCtx *direct.MapContext, in *krmalloydbv1alpha1.Instance_PSCAutoConnectionConfig) *pb.Instance_PscAutoConnectionConfig {
	if in == nil {
		return nil
	}
	out := &pb.Instance_PscAutoConnectionConfig{}
	out.ConsumerProject = direct.ValueOf(in.ConsumerProject)
	out.ConsumerNetwork = direct.ValueOf(in.ConsumerNetwork)
	// MISSING: IPAddress
	// MISSING: Status
	// MISSING: ConsumerNetworkStatus
	return out
}
func Instance_PSCAutoConnectionConfig_FromProto(mapCtx *direct.MapContext, in *pb.Instance_PscAutoConnectionConfig) *krm.Instance_PSCAutoConnectionConfig {
	if in == nil {
		return nil
	}
	out := &krm.Instance_PSCAutoConnectionConfig{}
	out.ConsumerProject = direct.LazyPtr(in.GetConsumerProject())
	out.ConsumerNetwork = direct.LazyPtr(in.GetConsumerNetwork())
	// MISSING: IPAddress
	// MISSING: Status
	// MISSING: ConsumerNetworkStatus
	return out
}
func Instance_PSCAutoConnectionConfig_ToProto(mapCtx *direct.MapContext, in *krm.Instance_PSCAutoConnectionConfig) *pb.Instance_PscAutoConnectionConfig {
	if in == nil {
		return nil
	}
	out := &pb.Instance_PscAutoConnectionConfig{}
	out.ConsumerProject = direct.ValueOf(in.ConsumerProject)
	out.ConsumerNetwork = direct.ValueOf(in.ConsumerNetwork)
	// MISSING: IPAddress
	// MISSING: Status
	// MISSING: ConsumerNetworkStatus
	return out
}
func Instance_PSCAutoConnectionConfigObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Instance_PscAutoConnectionConfig) *krmalloydbv1alpha1.Instance_PSCAutoConnectionConfigObservedState {
	if in == nil {
		return nil
	}
	out := &krmalloydbv1alpha1.Instance_PSCAutoConnectionConfigObservedState{}
	// MISSING: ConsumerProject
	// MISSING: ConsumerNetwork
	out.IPAddress = direct.LazyPtr(in.GetIpAddress())
	out.Status = direct.LazyPtr(in.GetStatus())
	out.ConsumerNetworkStatus = direct.LazyPtr(in.GetConsumerNetworkStatus())
	return out
}
func Instance_PSCAutoConnectionConfigObservedState_ToProto(mapCtx *direct.MapContext, in *krmalloydbv1alpha1.Instance_PSCAutoConnectionConfigObservedState) *pb.Instance_PscAutoConnectionConfig {
	if in == nil {
		return nil
	}
	out := &pb.Instance_PscAutoConnectionConfig{}
	// MISSING: ConsumerProject
	// MISSING: ConsumerNetwork
	out.IpAddress = direct.ValueOf(in.IPAddress)
	out.Status = direct.ValueOf(in.Status)
	out.ConsumerNetworkStatus = direct.ValueOf(in.ConsumerNetworkStatus)
	return out
}
func Instance_PSCAutoConnectionConfigObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Instance_PscAutoConnectionConfig) *krm.Instance_PSCAutoConnectionConfigObservedState {
	if in == nil {
		return nil
	}
	out := &krm.Instance_PSCAutoConnectionConfigObservedState{}
	// MISSING: ConsumerProject
	// MISSING: ConsumerNetwork
	out.IPAddress = direct.LazyPtr(in.GetIpAddress())
	out.Status = direct.LazyPtr(in.GetStatus())
	out.ConsumerNetworkStatus = direct.LazyPtr(in.GetConsumerNetworkStatus())
	return out
}
func Instance_PSCAutoConnectionConfigObservedState_ToProto(mapCtx *direct.MapContext, in *krm.Instance_PSCAutoConnectionConfigObservedState) *pb.Instance_PscAutoConnectionConfig {
	if in == nil {
		return nil
	}
	out := &pb.Instance_PscAutoConnectionConfig{}
	// MISSING: ConsumerProject
	// MISSING: ConsumerNetwork
	out.IpAddress = direct.ValueOf(in.IPAddress)
	out.Status = direct.ValueOf(in.Status)
	out.ConsumerNetworkStatus = direct.ValueOf(in.ConsumerNetworkStatus)
	return out
}
func Instance_PSCInstanceConfig_FromProto(mapCtx *direct.MapContext, in *pb.Instance_PscInstanceConfig) *krmalloydbv1alpha1.Instance_PSCInstanceConfig {
	if in == nil {
		return nil
	}
	out := &krmalloydbv1alpha1.Instance_PSCInstanceConfig{}
	// MISSING: ServiceAttachmentLink
	out.AllowedConsumerProjects = in.AllowedConsumerProjects
	// MISSING: PSCDNSName
	out.PSCInterfaceConfigs = direct.Slice_FromProto(mapCtx, in.PSCInterfaceConfigs, Instance_PSCInterfaceConfig_FromProto)
	out.PSCAutoConnections = direct.Slice_FromProto(mapCtx, in.PSCAutoConnections, Instance_PSCAutoConnectionConfig_FromProto)
	return out
}
func Instance_PSCInstanceConfig_ToProto(mapCtx *direct.MapContext, in *krmalloydbv1alpha1.Instance_PSCInstanceConfig) *pb.Instance_PscInstanceConfig {
	if in == nil {
		return nil
	}
	out := &pb.Instance_PscInstanceConfig{}
	// MISSING: ServiceAttachmentLink
	out.AllowedConsumerProjects = in.AllowedConsumerProjects
	// MISSING: PSCDNSName
	out.PscInterfaceConfigs = direct.Slice_ToProto(mapCtx, in.PSCInterfaceConfigs, Instance_PSCInterfaceConfig_ToProto)
	out.PscAutoConnections = direct.Slice_ToProto(mapCtx, in.PSCAutoConnections, Instance_PSCAutoConnectionConfig_ToProto)
	return out
}
func Instance_PSCInstanceConfig_FromProto(mapCtx *direct.MapContext, in *pb.Instance_PscInstanceConfig) *krm.Instance_PSCInstanceConfig {
	if in == nil {
		return nil
	}
	out := &krm.Instance_PSCInstanceConfig{}
	// MISSING: ServiceAttachmentLink
	out.AllowedConsumerProjects = in.AllowedConsumerProjects
	// MISSING: PSCDNSName
	out.PSCInterfaceConfigs = direct.Slice_FromProto(mapCtx, in.PSCInterfaceConfigs, Instance_PSCInterfaceConfig_FromProto)
	out.PSCAutoConnections = direct.Slice_FromProto(mapCtx, in.PSCAutoConnections, Instance_PSCAutoConnectionConfig_FromProto)
	return out
}
func Instance_PSCInstanceConfig_ToProto(mapCtx *direct.MapContext, in *krm.Instance_PSCInstanceConfig) *pb.Instance_PscInstanceConfig {
	if in == nil {
		return nil
	}
	out := &pb.Instance_PscInstanceConfig{}
	// MISSING: ServiceAttachmentLink
	out.AllowedConsumerProjects = in.AllowedConsumerProjects
	// MISSING: PSCDNSName
	out.PscInterfaceConfigs = direct.Slice_ToProto(mapCtx, in.PSCInterfaceConfigs, Instance_PSCInterfaceConfig_ToProto)
	out.PscAutoConnections = direct.Slice_ToProto(mapCtx, in.PSCAutoConnections, Instance_PSCAutoConnectionConfig_ToProto)
	return out
}
func Instance_PSCInstanceConfigObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Instance_PscInstanceConfig) *krmalloydbv1alpha1.Instance_PSCInstanceConfigObservedState {
	if in == nil {
		return nil
	}
	out := &krmalloydbv1alpha1.Instance_PSCInstanceConfigObservedState{}
	out.ServiceAttachmentLink = direct.LazyPtr(in.GetServiceAttachmentLink())
	// MISSING: AllowedConsumerProjects
	out.PSCDNSName = direct.LazyPtr(in.GetPscDnsName())
	// MISSING: PSCInterfaceConfigs
	out.PSCAutoConnections = direct.Slice_FromProto(mapCtx, in.PSCAutoConnections, Instance_PSCAutoConnectionConfigObservedState_FromProto)
	return out
}
func Instance_PSCInstanceConfigObservedState_ToProto(mapCtx *direct.MapContext, in *krmalloydbv1alpha1.Instance_PSCInstanceConfigObservedState) *pb.Instance_PscInstanceConfig {
	if in == nil {
		return nil
	}
	out := &pb.Instance_PscInstanceConfig{}
	out.ServiceAttachmentLink = direct.ValueOf(in.ServiceAttachmentLink)
	// MISSING: AllowedConsumerProjects
	out.PscDnsName = direct.ValueOf(in.PSCDNSName)
	// MISSING: PSCInterfaceConfigs
	out.PscAutoConnections = direct.Slice_ToProto(mapCtx, in.PSCAutoConnections, Instance_PSCAutoConnectionConfigObservedState_ToProto)
	return out
}
func Instance_PSCInstanceConfigObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Instance_PscInstanceConfig) *krm.Instance_PSCInstanceConfigObservedState {
	if in == nil {
		return nil
	}
	out := &krm.Instance_PSCInstanceConfigObservedState{}
	out.ServiceAttachmentLink = direct.LazyPtr(in.GetServiceAttachmentLink())
	// MISSING: AllowedConsumerProjects
	out.PSCDNSName = direct.LazyPtr(in.GetPscDnsName())
	// MISSING: PSCInterfaceConfigs
	out.PSCAutoConnections = direct.Slice_FromProto(mapCtx, in.PSCAutoConnections, Instance_PSCAutoConnectionConfigObservedState_FromProto)
	return out
}
func Instance_PSCInstanceConfigObservedState_ToProto(mapCtx *direct.MapContext, in *krm.Instance_PSCInstanceConfigObservedState) *pb.Instance_PscInstanceConfig {
	if in == nil {
		return nil
	}
	out := &pb.Instance_PscInstanceConfig{}
	out.ServiceAttachmentLink = direct.ValueOf(in.ServiceAttachmentLink)
	// MISSING: AllowedConsumerProjects
	out.PscDnsName = direct.ValueOf(in.PSCDNSName)
	// MISSING: PSCInterfaceConfigs
	out.PscAutoConnections = direct.Slice_ToProto(mapCtx, in.PSCAutoConnections, Instance_PSCAutoConnectionConfigObservedState_ToProto)
	return out
}
func Instance_PSCInterfaceConfig_FromProto(mapCtx *direct.MapContext, in *pb.Instance_PscInterfaceConfig) *krmalloydbv1alpha1.Instance_PSCInterfaceConfig {
	if in == nil {
		return nil
	}
	out := &krmalloydbv1alpha1.Instance_PSCInterfaceConfig{}
	out.NetworkAttachmentResource = direct.LazyPtr(in.GetNetworkAttachmentResource())
	return out
}
func Instance_PSCInterfaceConfig_ToProto(mapCtx *direct.MapContext, in *krmalloydbv1alpha1.Instance_PSCInterfaceConfig) *pb.Instance_PscInterfaceConfig {
	if in == nil {
		return nil
	}
	out := &pb.Instance_PscInterfaceConfig{}
	out.NetworkAttachmentResource = direct.ValueOf(in.NetworkAttachmentResource)
	return out
}
func Instance_PSCInterfaceConfig_FromProto(mapCtx *direct.MapContext, in *pb.Instance_PscInterfaceConfig) *krm.Instance_PSCInterfaceConfig {
	if in == nil {
		return nil
	}
	out := &krm.Instance_PSCInterfaceConfig{}
	out.NetworkAttachmentResource = direct.LazyPtr(in.GetNetworkAttachmentResource())
	return out
}
func Instance_PSCInterfaceConfig_ToProto(mapCtx *direct.MapContext, in *krm.Instance_PSCInterfaceConfig) *pb.Instance_PscInterfaceConfig {
	if in == nil {
		return nil
	}
	out := &pb.Instance_PscInterfaceConfig{}
	out.NetworkAttachmentResource = direct.ValueOf(in.NetworkAttachmentResource)
	return out
}
func Instance_QueryInsightsInstanceConfig_FromProto(mapCtx *direct.MapContext, in *pb.Instance_QueryInsightsInstanceConfig) *krmalloydbv1alpha1.Instance_QueryInsightsInstanceConfig {
	if in == nil {
		return nil
	}
	out := &krmalloydbv1alpha1.Instance_QueryInsightsInstanceConfig{}
	out.RecordApplicationTags = in.RecordApplicationTags
	out.RecordClientAddress = in.RecordClientAddress
	out.QueryStringLength = direct.LazyPtr(in.GetQueryStringLength())
	out.QueryPlansPerMinute = in.QueryPlansPerMinute
	return out
}
func Instance_QueryInsightsInstanceConfig_ToProto(mapCtx *direct.MapContext, in *krmalloydbv1alpha1.Instance_QueryInsightsInstanceConfig) *pb.Instance_QueryInsightsInstanceConfig {
	if in == nil {
		return nil
	}
	out := &pb.Instance_QueryInsightsInstanceConfig{}
	out.RecordApplicationTags = in.RecordApplicationTags
	out.RecordClientAddress = in.RecordClientAddress
	out.QueryStringLength = direct.ValueOf(in.QueryStringLength)
	out.QueryPlansPerMinute = in.QueryPlansPerMinute
	return out
}
func Instance_QueryInsightsInstanceConfig_FromProto(mapCtx *direct.MapContext, in *pb.Instance_QueryInsightsInstanceConfig) *krm.Instance_QueryInsightsInstanceConfig {
	if in == nil {
		return nil
	}
	out := &krm.Instance_QueryInsightsInstanceConfig{}
	out.RecordApplicationTags = in.RecordApplicationTags
	out.RecordClientAddress = in.RecordClientAddress
	out.QueryStringLength = direct.LazyPtr(in.GetQueryStringLength())
	out.QueryPlansPerMinute = in.QueryPlansPerMinute
	return out
}
func Instance_QueryInsightsInstanceConfig_ToProto(mapCtx *direct.MapContext, in *krm.Instance_QueryInsightsInstanceConfig) *pb.Instance_QueryInsightsInstanceConfig {
	if in == nil {
		return nil
	}
	out := &pb.Instance_QueryInsightsInstanceConfig{}
	out.RecordApplicationTags = in.RecordApplicationTags
	out.RecordClientAddress = in.RecordClientAddress
	out.QueryStringLength = direct.ValueOf(in.QueryStringLength)
	out.QueryPlansPerMinute = in.QueryPlansPerMinute
	return out
}
func Instance_ReadPoolConfig_FromProto(mapCtx *direct.MapContext, in *pb.Instance_ReadPoolConfig) *krmalloydbv1alpha1.Instance_ReadPoolConfig {
	if in == nil {
		return nil
	}
	out := &krmalloydbv1alpha1.Instance_ReadPoolConfig{}
	out.NodeCount = direct.LazyPtr(in.GetNodeCount())
	return out
}
func Instance_ReadPoolConfig_ToProto(mapCtx *direct.MapContext, in *krmalloydbv1alpha1.Instance_ReadPoolConfig) *pb.Instance_ReadPoolConfig {
	if in == nil {
		return nil
	}
	out := &pb.Instance_ReadPoolConfig{}
	out.NodeCount = direct.ValueOf(in.NodeCount)
	return out
}
func Instance_ReadPoolConfig_FromProto(mapCtx *direct.MapContext, in *pb.Instance_ReadPoolConfig) *krm.Instance_ReadPoolConfig {
	if in == nil {
		return nil
	}
	out := &krm.Instance_ReadPoolConfig{}
	out.NodeCount = direct.LazyPtr(in.GetNodeCount())
	return out
}
func Instance_ReadPoolConfig_ToProto(mapCtx *direct.MapContext, in *krm.Instance_ReadPoolConfig) *pb.Instance_ReadPoolConfig {
	if in == nil {
		return nil
	}
	out := &pb.Instance_ReadPoolConfig{}
	out.NodeCount = direct.ValueOf(in.NodeCount)
	return out
}
func Instance_UpdatePolicy_FromProto(mapCtx *direct.MapContext, in *pb.Instance_UpdatePolicy) *krmalloydbv1alpha1.Instance_UpdatePolicy {
	if in == nil {
		return nil
	}
	out := &krmalloydbv1alpha1.Instance_UpdatePolicy{}
	out.Mode = direct.Enum_FromProto(mapCtx, in.GetMode())
	return out
}
func Instance_UpdatePolicy_ToProto(mapCtx *direct.MapContext, in *krmalloydbv1alpha1.Instance_UpdatePolicy) *pb.Instance_UpdatePolicy {
	if in == nil {
		return nil
	}
	out := &pb.Instance_UpdatePolicy{}
	out.Mode = direct.Enum_ToProto[pb.Instance_UpdatePolicy_Mode](mapCtx, in.Mode)
	return out
}
func Instance_UpdatePolicy_FromProto(mapCtx *direct.MapContext, in *pb.Instance_UpdatePolicy) *krm.Instance_UpdatePolicy {
	if in == nil {
		return nil
	}
	out := &krm.Instance_UpdatePolicy{}
	out.Mode = direct.Enum_FromProto(mapCtx, in.GetMode())
	return out
}
func Instance_UpdatePolicy_ToProto(mapCtx *direct.MapContext, in *krm.Instance_UpdatePolicy) *pb.Instance_UpdatePolicy {
	if in == nil {
		return nil
	}
	out := &pb.Instance_UpdatePolicy{}
	out.Mode = direct.Enum_ToProto[pb.Instance_UpdatePolicy_Mode](mapCtx, in.Mode)
	return out
}
func MaintenanceSchedule_FromProto(mapCtx *direct.MapContext, in *pb.MaintenanceSchedule) *krmalloydbv1alpha1.MaintenanceSchedule {
	if in == nil {
		return nil
	}
	out := &krmalloydbv1alpha1.MaintenanceSchedule{}
	// MISSING: StartTime
	return out
}
func MaintenanceSchedule_ToProto(mapCtx *direct.MapContext, in *krmalloydbv1alpha1.MaintenanceSchedule) *pb.MaintenanceSchedule {
	if in == nil {
		return nil
	}
	out := &pb.MaintenanceSchedule{}
	// MISSING: StartTime
	return out
}
func MaintenanceSchedule_FromProto(mapCtx *direct.MapContext, in *pb.MaintenanceSchedule) *krm.MaintenanceSchedule {
	if in == nil {
		return nil
	}
	out := &krm.MaintenanceSchedule{}
	// MISSING: StartTime
	return out
}
func MaintenanceSchedule_ToProto(mapCtx *direct.MapContext, in *krm.MaintenanceSchedule) *pb.MaintenanceSchedule {
	if in == nil {
		return nil
	}
	out := &pb.MaintenanceSchedule{}
	// MISSING: StartTime
	return out
}
func MaintenanceScheduleObservedState_FromProto(mapCtx *direct.MapContext, in *pb.MaintenanceSchedule) *krmalloydbv1alpha1.MaintenanceScheduleObservedState {
	if in == nil {
		return nil
	}
	out := &krmalloydbv1alpha1.MaintenanceScheduleObservedState{}
	out.StartTime = direct.StringTimestamp_FromProto(mapCtx, in.GetStartTime())
	return out
}
func MaintenanceScheduleObservedState_ToProto(mapCtx *direct.MapContext, in *krmalloydbv1alpha1.MaintenanceScheduleObservedState) *pb.MaintenanceSchedule {
	if in == nil {
		return nil
	}
	out := &pb.MaintenanceSchedule{}
	out.StartTime = direct.StringTimestamp_ToProto(mapCtx, in.StartTime)
	return out
}
func MaintenanceScheduleObservedState_FromProto(mapCtx *direct.MapContext, in *pb.MaintenanceSchedule) *krm.MaintenanceScheduleObservedState {
	if in == nil {
		return nil
	}
	out := &krm.MaintenanceScheduleObservedState{}
	out.StartTime = direct.StringTimestamp_FromProto(mapCtx, in.GetStartTime())
	return out
}
func MaintenanceScheduleObservedState_ToProto(mapCtx *direct.MapContext, in *krm.MaintenanceScheduleObservedState) *pb.MaintenanceSchedule {
	if in == nil {
		return nil
	}
	out := &pb.MaintenanceSchedule{}
	out.StartTime = direct.StringTimestamp_ToProto(mapCtx, in.StartTime)
	return out
}
func MaintenanceUpdatePolicy_FromProto(mapCtx *direct.MapContext, in *pb.MaintenanceUpdatePolicy) *krmalloydbv1alpha1.MaintenanceUpdatePolicy {
	if in == nil {
		return nil
	}
	out := &krmalloydbv1alpha1.MaintenanceUpdatePolicy{}
	out.MaintenanceWindows = direct.Slice_FromProto(mapCtx, in.MaintenanceWindows, MaintenanceUpdatePolicy_MaintenanceWindow_FromProto)
	out.DenyMaintenancePeriods = direct.Slice_FromProto(mapCtx, in.DenyMaintenancePeriods, MaintenanceUpdatePolicy_DenyMaintenancePeriod_FromProto)
	return out
}
func MaintenanceUpdatePolicy_ToProto(mapCtx *direct.MapContext, in *krmalloydbv1alpha1.MaintenanceUpdatePolicy) *pb.MaintenanceUpdatePolicy {
	if in == nil {
		return nil
	}
	out := &pb.MaintenanceUpdatePolicy{}
	out.MaintenanceWindows = direct.Slice_ToProto(mapCtx, in.MaintenanceWindows, MaintenanceUpdatePolicy_MaintenanceWindow_ToProto)
	out.DenyMaintenancePeriods = direct.Slice_ToProto(mapCtx, in.DenyMaintenancePeriods, MaintenanceUpdatePolicy_DenyMaintenancePeriod_ToProto)
	return out
}
func MaintenanceUpdatePolicy_FromProto(mapCtx *direct.MapContext, in *pb.MaintenanceUpdatePolicy) *krm.MaintenanceUpdatePolicy {
	if in == nil {
		return nil
	}
	out := &krm.MaintenanceUpdatePolicy{}
	out.MaintenanceWindows = direct.Slice_FromProto(mapCtx, in.MaintenanceWindows, MaintenanceUpdatePolicy_MaintenanceWindow_FromProto)
	out.DenyMaintenancePeriods = direct.Slice_FromProto(mapCtx, in.DenyMaintenancePeriods, MaintenanceUpdatePolicy_DenyMaintenancePeriod_FromProto)
	return out
}
func MaintenanceUpdatePolicy_ToProto(mapCtx *direct.MapContext, in *krm.MaintenanceUpdatePolicy) *pb.MaintenanceUpdatePolicy {
	if in == nil {
		return nil
	}
	out := &pb.MaintenanceUpdatePolicy{}
	out.MaintenanceWindows = direct.Slice_ToProto(mapCtx, in.MaintenanceWindows, MaintenanceUpdatePolicy_MaintenanceWindow_ToProto)
	out.DenyMaintenancePeriods = direct.Slice_ToProto(mapCtx, in.DenyMaintenancePeriods, MaintenanceUpdatePolicy_DenyMaintenancePeriod_ToProto)
	return out
}
func MaintenanceUpdatePolicy_DenyMaintenancePeriod_FromProto(mapCtx *direct.MapContext, in *pb.MaintenanceUpdatePolicy_DenyMaintenancePeriod) *krmalloydbv1alpha1.MaintenanceUpdatePolicy_DenyMaintenancePeriod {
	if in == nil {
		return nil
	}
	out := &krmalloydbv1alpha1.MaintenanceUpdatePolicy_DenyMaintenancePeriod{}
	out.StartDate = Date_FromProto(mapCtx, in.GetStartDate())
	out.EndDate = Date_FromProto(mapCtx, in.GetEndDate())
	out.Time = TimeOfDay_FromProto(mapCtx, in.GetTime())
	return out
}
func MaintenanceUpdatePolicy_DenyMaintenancePeriod_ToProto(mapCtx *direct.MapContext, in *krmalloydbv1alpha1.MaintenanceUpdatePolicy_DenyMaintenancePeriod) *pb.MaintenanceUpdatePolicy_DenyMaintenancePeriod {
	if in == nil {
		return nil
	}
	out := &pb.MaintenanceUpdatePolicy_DenyMaintenancePeriod{}
	out.StartDate = Date_ToProto(mapCtx, in.StartDate)
	out.EndDate = Date_ToProto(mapCtx, in.EndDate)
	out.Time = TimeOfDay_ToProto(mapCtx, in.Time)
	return out
}
func MaintenanceUpdatePolicy_DenyMaintenancePeriod_FromProto(mapCtx *direct.MapContext, in *pb.MaintenanceUpdatePolicy_DenyMaintenancePeriod) *krm.MaintenanceUpdatePolicy_DenyMaintenancePeriod {
	if in == nil {
		return nil
	}
	out := &krm.MaintenanceUpdatePolicy_DenyMaintenancePeriod{}
	out.StartDate = Date_FromProto(mapCtx, in.GetStartDate())
	out.EndDate = Date_FromProto(mapCtx, in.GetEndDate())
	out.Time = TimeOfDay_FromProto(mapCtx, in.GetTime())
	return out
}
func MaintenanceUpdatePolicy_DenyMaintenancePeriod_ToProto(mapCtx *direct.MapContext, in *krm.MaintenanceUpdatePolicy_DenyMaintenancePeriod) *pb.MaintenanceUpdatePolicy_DenyMaintenancePeriod {
	if in == nil {
		return nil
	}
	out := &pb.MaintenanceUpdatePolicy_DenyMaintenancePeriod{}
	out.StartDate = Date_ToProto(mapCtx, in.StartDate)
	out.EndDate = Date_ToProto(mapCtx, in.EndDate)
	out.Time = TimeOfDay_ToProto(mapCtx, in.Time)
	return out
}
func MigrationSource_FromProto(mapCtx *direct.MapContext, in *pb.MigrationSource) *krmalloydbv1alpha1.MigrationSource {
	if in == nil {
		return nil
	}
	out := &krmalloydbv1alpha1.MigrationSource{}
	// MISSING: HostPort
	// MISSING: ReferenceID
	// MISSING: SourceType
	return out
}
func MigrationSource_ToProto(mapCtx *direct.MapContext, in *krmalloydbv1alpha1.MigrationSource) *pb.MigrationSource {
	if in == nil {
		return nil
	}
	out := &pb.MigrationSource{}
	// MISSING: HostPort
	// MISSING: ReferenceID
	// MISSING: SourceType
	return out
}
func MigrationSourceObservedState_FromProto(mapCtx *direct.MapContext, in *pb.MigrationSource) *krmalloydbv1alpha1.MigrationSourceObservedState {
	if in == nil {
		return nil
	}
	out := &krmalloydbv1alpha1.MigrationSourceObservedState{}
	out.HostPort = direct.LazyPtr(in.GetHostPort())
	out.ReferenceID = direct.LazyPtr(in.GetReferenceId())
	out.SourceType = direct.Enum_FromProto(mapCtx, in.GetSourceType())
	return out
}
func MigrationSourceObservedState_ToProto(mapCtx *direct.MapContext, in *krmalloydbv1alpha1.MigrationSourceObservedState) *pb.MigrationSource {
	if in == nil {
		return nil
	}
	out := &pb.MigrationSource{}
	out.HostPort = direct.ValueOf(in.HostPort)
	out.ReferenceId = direct.ValueOf(in.ReferenceID)
	out.SourceType = direct.Enum_ToProto[pb.MigrationSource_MigrationSourceType](mapCtx, in.SourceType)
	return out
}
func MigrationSourceObservedState_FromProto(mapCtx *direct.MapContext, in *pb.MigrationSource) *krm.MigrationSourceObservedState {
	if in == nil {
		return nil
	}
	out := &krm.MigrationSourceObservedState{}
	out.HostPort = direct.LazyPtr(in.GetHostPort())
	out.ReferenceID = direct.LazyPtr(in.GetReferenceId())
	out.SourceType = direct.Enum_FromProto(mapCtx, in.GetSourceType())
	return out
}
func MigrationSourceObservedState_ToProto(mapCtx *direct.MapContext, in *krm.MigrationSourceObservedState) *pb.MigrationSource {
	if in == nil {
		return nil
	}
	out := &pb.MigrationSource{}
	out.HostPort = direct.ValueOf(in.HostPort)
	out.ReferenceId = direct.ValueOf(in.ReferenceID)
	out.SourceType = direct.Enum_ToProto[pb.MigrationSource_MigrationSourceType](mapCtx, in.SourceType)
	return out
}
func SSLConfig_FromProto(mapCtx *direct.MapContext, in *pb.SslConfig) *krmalloydbv1alpha1.SSLConfig {
	if in == nil {
		return nil
	}
	out := &krmalloydbv1alpha1.SSLConfig{}
	out.SSLMode = direct.Enum_FromProto(mapCtx, in.GetSslMode())
	out.CASource = direct.Enum_FromProto(mapCtx, in.GetCaSource())
	return out
}
func SSLConfig_ToProto(mapCtx *direct.MapContext, in *krmalloydbv1alpha1.SSLConfig) *pb.SslConfig {
	if in == nil {
		return nil
	}
	out := &pb.SslConfig{}
	out.SslMode = direct.Enum_ToProto[pb.SslConfig_SslMode](mapCtx, in.SSLMode)
	out.CaSource = direct.Enum_ToProto[pb.SslConfig_CaSource](mapCtx, in.CASource)
	return out
}
func SSLConfig_FromProto(mapCtx *direct.MapContext, in *pb.SslConfig) *krm.SSLConfig {
	if in == nil {
		return nil
	}
	out := &krm.SSLConfig{}
	out.SSLMode = direct.Enum_FromProto(mapCtx, in.GetSslMode())
	out.CASource = direct.Enum_FromProto(mapCtx, in.GetCaSource())
	return out
}
func SSLConfig_ToProto(mapCtx *direct.MapContext, in *krm.SSLConfig) *pb.SslConfig {
	if in == nil {
		return nil
	}
	out := &pb.SslConfig{}
	out.SslMode = direct.Enum_ToProto[pb.SslConfig_SslMode](mapCtx, in.SSLMode)
	out.CaSource = direct.Enum_ToProto[pb.SslConfig_CaSource](mapCtx, in.CASource)
	return out
}
