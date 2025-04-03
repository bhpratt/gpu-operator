/*
Copyright (c) Advanced Micro Devices, Inc. All rights reserved.

Licensed under the Apache License, Version 2.0 (the \"License\");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an \"AS IS\" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by MockGen. DO NOT EDIT.
// Source: upgrademgr.go
//
// Generated by this command:
//
//	mockgen -source=upgrademgr.go -package=controllers -destination=mock_upgrademgr.go upgradeMgrHelperAPI
//
// Package controllers is a generated GoMock package.
package controllers

import (
	context "context"
	reflect "reflect"

	v1alpha1 "github.com/ROCm/gpu-operator/api/v1alpha1"
	gomock "go.uber.org/mock/gomock"
	v1 "k8s.io/api/core/v1"
	controllerruntime "sigs.k8s.io/controller-runtime"
)

// MockupgradeMgrAPI is a mock of upgradeMgrAPI interface.
type MockupgradeMgrAPI struct {
	ctrl     *gomock.Controller
	recorder *MockupgradeMgrAPIMockRecorder
}

// MockupgradeMgrAPIMockRecorder is the mock recorder for MockupgradeMgrAPI.
type MockupgradeMgrAPIMockRecorder struct {
	mock *MockupgradeMgrAPI
}

// NewMockupgradeMgrAPI creates a new mock instance.
func NewMockupgradeMgrAPI(ctrl *gomock.Controller) *MockupgradeMgrAPI {
	mock := &MockupgradeMgrAPI{ctrl: ctrl}
	mock.recorder = &MockupgradeMgrAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockupgradeMgrAPI) EXPECT() *MockupgradeMgrAPIMockRecorder {
	return m.recorder
}

// GetNodeBootId mocks base method.
func (m *MockupgradeMgrAPI) GetNodeBootId(nodeName string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNodeBootId", nodeName)
	ret0, _ := ret[0].(string)
	return ret0
}

// GetNodeBootId indicates an expected call of GetNodeBootId.
func (mr *MockupgradeMgrAPIMockRecorder) GetNodeBootId(nodeName any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNodeBootId", reflect.TypeOf((*MockupgradeMgrAPI)(nil).GetNodeBootId), nodeName)
}

// GetNodeStatus mocks base method.
func (m *MockupgradeMgrAPI) GetNodeStatus(nodeName string) v1alpha1.UpgradeState {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNodeStatus", nodeName)
	ret0, _ := ret[0].(v1alpha1.UpgradeState)
	return ret0
}

// GetNodeStatus indicates an expected call of GetNodeStatus.
func (mr *MockupgradeMgrAPIMockRecorder) GetNodeStatus(nodeName any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNodeStatus", reflect.TypeOf((*MockupgradeMgrAPI)(nil).GetNodeStatus), nodeName)
}

// GetNodeUpgradeStartTime mocks base method.
func (m *MockupgradeMgrAPI) GetNodeUpgradeStartTime(nodeName string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNodeUpgradeStartTime", nodeName)
	ret0, _ := ret[0].(string)
	return ret0
}

// GetNodeUpgradeStartTime indicates an expected call of GetNodeUpgradeStartTime.
func (mr *MockupgradeMgrAPIMockRecorder) GetNodeUpgradeStartTime(nodeName any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNodeUpgradeStartTime", reflect.TypeOf((*MockupgradeMgrAPI)(nil).GetNodeUpgradeStartTime), nodeName)
}

// HandleDelete mocks base method.
func (m *MockupgradeMgrAPI) HandleDelete(ctx context.Context, deviceConfig *v1alpha1.DeviceConfig, nodes *v1.NodeList) (controllerruntime.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HandleDelete", ctx, deviceConfig, nodes)
	ret0, _ := ret[0].(controllerruntime.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HandleDelete indicates an expected call of HandleDelete.
func (mr *MockupgradeMgrAPIMockRecorder) HandleDelete(ctx, deviceConfig, nodes any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleDelete", reflect.TypeOf((*MockupgradeMgrAPI)(nil).HandleDelete), ctx, deviceConfig, nodes)
}

// HandleUpgrade mocks base method.
func (m *MockupgradeMgrAPI) HandleUpgrade(ctx context.Context, deviceConfig *v1alpha1.DeviceConfig, nodes *v1.NodeList) (controllerruntime.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HandleUpgrade", ctx, deviceConfig, nodes)
	ret0, _ := ret[0].(controllerruntime.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HandleUpgrade indicates an expected call of HandleUpgrade.
func (mr *MockupgradeMgrAPIMockRecorder) HandleUpgrade(ctx, deviceConfig, nodes any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleUpgrade", reflect.TypeOf((*MockupgradeMgrAPI)(nil).HandleUpgrade), ctx, deviceConfig, nodes)
}

// MockupgradeMgrHelperAPI is a mock of upgradeMgrHelperAPI interface.
type MockupgradeMgrHelperAPI struct {
	ctrl     *gomock.Controller
	recorder *MockupgradeMgrHelperAPIMockRecorder
}

// MockupgradeMgrHelperAPIMockRecorder is the mock recorder for MockupgradeMgrHelperAPI.
type MockupgradeMgrHelperAPIMockRecorder struct {
	mock *MockupgradeMgrHelperAPI
}

// NewMockupgradeMgrHelperAPI creates a new mock instance.
func NewMockupgradeMgrHelperAPI(ctrl *gomock.Controller) *MockupgradeMgrHelperAPI {
	mock := &MockupgradeMgrHelperAPI{ctrl: ctrl}
	mock.recorder = &MockupgradeMgrHelperAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockupgradeMgrHelperAPI) EXPECT() *MockupgradeMgrHelperAPIMockRecorder {
	return m.recorder
}

// checkUpgradeTimeExceeded mocks base method.
func (m *MockupgradeMgrHelperAPI) checkUpgradeTimeExceeded(ctx context.Context, nodeName string, deviceConfig *v1alpha1.DeviceConfig) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "checkUpgradeTimeExceeded", ctx, nodeName, deviceConfig)
	ret0, _ := ret[0].(bool)
	return ret0
}

// checkUpgradeTimeExceeded indicates an expected call of checkUpgradeTimeExceeded.
func (mr *MockupgradeMgrHelperAPIMockRecorder) checkUpgradeTimeExceeded(ctx, nodeName, deviceConfig any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "checkUpgradeTimeExceeded", reflect.TypeOf((*MockupgradeMgrHelperAPI)(nil).checkUpgradeTimeExceeded), ctx, nodeName, deviceConfig)
}

// clearNodeStatus mocks base method.
func (m *MockupgradeMgrHelperAPI) clearNodeStatus() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "clearNodeStatus")
}

// clearNodeStatus indicates an expected call of clearNodeStatus.
func (mr *MockupgradeMgrHelperAPIMockRecorder) clearNodeStatus() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "clearNodeStatus", reflect.TypeOf((*MockupgradeMgrHelperAPI)(nil).clearNodeStatus))
}

// clearUpgradeStartTime mocks base method.
func (m *MockupgradeMgrHelperAPI) clearUpgradeStartTime(nodeName string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "clearUpgradeStartTime", nodeName)
}

// clearUpgradeStartTime indicates an expected call of clearUpgradeStartTime.
func (mr *MockupgradeMgrHelperAPIMockRecorder) clearUpgradeStartTime(nodeName any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "clearUpgradeStartTime", reflect.TypeOf((*MockupgradeMgrHelperAPI)(nil).clearUpgradeStartTime), nodeName)
}

// cordonOrUncordonNode mocks base method.
func (m *MockupgradeMgrHelperAPI) cordonOrUncordonNode(ctx context.Context, deviceConfig *v1alpha1.DeviceConfig, node *v1.Node, add bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "cordonOrUncordonNode", ctx, deviceConfig, node, add)
	ret0, _ := ret[0].(error)
	return ret0
}

// cordonOrUncordonNode indicates an expected call of cordonOrUncordonNode.
func (mr *MockupgradeMgrHelperAPIMockRecorder) cordonOrUncordonNode(ctx, deviceConfig, node, add any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "cordonOrUncordonNode", reflect.TypeOf((*MockupgradeMgrHelperAPI)(nil).cordonOrUncordonNode), ctx, deviceConfig, node, add)
}

// deleteOrDrainPods mocks base method.
func (m *MockupgradeMgrHelperAPI) deleteOrDrainPods(ctx context.Context, deviceConfig *v1alpha1.DeviceConfig, node *v1.Node) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "deleteOrDrainPods", ctx, deviceConfig, node)
	ret0, _ := ret[0].(error)
	return ret0
}

// deleteOrDrainPods indicates an expected call of deleteOrDrainPods.
func (mr *MockupgradeMgrHelperAPIMockRecorder) deleteOrDrainPods(ctx, deviceConfig, node any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "deleteOrDrainPods", reflect.TypeOf((*MockupgradeMgrHelperAPI)(nil).deleteOrDrainPods), ctx, deviceConfig, node)
}

// deleteRebootPod mocks base method.
func (m *MockupgradeMgrHelperAPI) deleteRebootPod(ctx context.Context, nodeName string, dc *v1alpha1.DeviceConfig, force bool, genId int64) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "deleteRebootPod", ctx, nodeName, dc, force, genId)
}

// deleteRebootPod indicates an expected call of deleteRebootPod.
func (mr *MockupgradeMgrHelperAPIMockRecorder) deleteRebootPod(ctx, nodeName, dc, force, genId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "deleteRebootPod", reflect.TypeOf((*MockupgradeMgrHelperAPI)(nil).deleteRebootPod), ctx, nodeName, dc, force, genId)
}

// getBootID mocks base method.
func (m *MockupgradeMgrHelperAPI) getBootID(nodeName string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "getBootID", nodeName)
	ret0, _ := ret[0].(string)
	return ret0
}

// getBootID indicates an expected call of getBootID.
func (mr *MockupgradeMgrHelperAPIMockRecorder) getBootID(nodeName any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "getBootID", reflect.TypeOf((*MockupgradeMgrHelperAPI)(nil).getBootID), nodeName)
}

// getNode mocks base method.
func (m *MockupgradeMgrHelperAPI) getNode(ctx context.Context, nodeName string) (*v1.Node, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "getNode", ctx, nodeName)
	ret0, _ := ret[0].(*v1.Node)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// getNode indicates an expected call of getNode.
func (mr *MockupgradeMgrHelperAPIMockRecorder) getNode(ctx, nodeName any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "getNode", reflect.TypeOf((*MockupgradeMgrHelperAPI)(nil).getNode), ctx, nodeName)
}

// getNodeStatus mocks base method.
func (m *MockupgradeMgrHelperAPI) getNodeStatus(nodeName string) v1alpha1.UpgradeState {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "getNodeStatus", nodeName)
	ret0, _ := ret[0].(v1alpha1.UpgradeState)
	return ret0
}

// getNodeStatus indicates an expected call of getNodeStatus.
func (mr *MockupgradeMgrHelperAPIMockRecorder) getNodeStatus(nodeName any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "getNodeStatus", reflect.TypeOf((*MockupgradeMgrHelperAPI)(nil).getNodeStatus), nodeName)
}

// getPodsToDrainOrDelete mocks base method.
func (m *MockupgradeMgrHelperAPI) getPodsToDrainOrDelete(ctx context.Context, deviceConfig *v1alpha1.DeviceConfig, node *v1.Node) ([]v1.Pod, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "getPodsToDrainOrDelete", ctx, deviceConfig, node)
	ret0, _ := ret[0].([]v1.Pod)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// getPodsToDrainOrDelete indicates an expected call of getPodsToDrainOrDelete.
func (mr *MockupgradeMgrHelperAPIMockRecorder) getPodsToDrainOrDelete(ctx, deviceConfig, node any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "getPodsToDrainOrDelete", reflect.TypeOf((*MockupgradeMgrHelperAPI)(nil).getPodsToDrainOrDelete), ctx, deviceConfig, node)
}

// getRebootPod mocks base method.
func (m *MockupgradeMgrHelperAPI) getRebootPod(nodeName string, dc *v1alpha1.DeviceConfig) *v1.Pod {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "getRebootPod", nodeName, dc)
	ret0, _ := ret[0].(*v1.Pod)
	return ret0
}

// getRebootPod indicates an expected call of getRebootPod.
func (mr *MockupgradeMgrHelperAPIMockRecorder) getRebootPod(nodeName, dc any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "getRebootPod", reflect.TypeOf((*MockupgradeMgrHelperAPI)(nil).getRebootPod), nodeName, dc)
}

// getUpgradeStartTime mocks base method.
func (m *MockupgradeMgrHelperAPI) getUpgradeStartTime(nodeName string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "getUpgradeStartTime", nodeName)
	ret0, _ := ret[0].(string)
	return ret0
}

// getUpgradeStartTime indicates an expected call of getUpgradeStartTime.
func (mr *MockupgradeMgrHelperAPIMockRecorder) getUpgradeStartTime(nodeName any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "getUpgradeStartTime", reflect.TypeOf((*MockupgradeMgrHelperAPI)(nil).getUpgradeStartTime), nodeName)
}

// handleInitStatus mocks base method.
func (m *MockupgradeMgrHelperAPI) handleInitStatus(ctx context.Context, node *v1.Node) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "handleInitStatus", ctx, node)
}

// handleInitStatus indicates an expected call of handleInitStatus.
func (mr *MockupgradeMgrHelperAPIMockRecorder) handleInitStatus(ctx, node any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "handleInitStatus", reflect.TypeOf((*MockupgradeMgrHelperAPI)(nil).handleInitStatus), ctx, node)
}

// handleNodeReboot mocks base method.
func (m *MockupgradeMgrHelperAPI) handleNodeReboot(ctx context.Context, node *v1.Node, dc *v1alpha1.DeviceConfig) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "handleNodeReboot", ctx, node, dc)
}

// handleNodeReboot indicates an expected call of handleNodeReboot.
func (mr *MockupgradeMgrHelperAPIMockRecorder) handleNodeReboot(ctx, node, dc any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "handleNodeReboot", reflect.TypeOf((*MockupgradeMgrHelperAPI)(nil).handleNodeReboot), ctx, node, dc)
}

// handleNodeUpgrade mocks base method.
func (m *MockupgradeMgrHelperAPI) handleNodeUpgrade(ctx context.Context, deviceConfig v1alpha1.DeviceConfig, node v1.Node) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "handleNodeUpgrade", ctx, deviceConfig, node)
}

// handleNodeUpgrade indicates an expected call of handleNodeUpgrade.
func (mr *MockupgradeMgrHelperAPIMockRecorder) handleNodeUpgrade(ctx, deviceConfig, node any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "handleNodeUpgrade", reflect.TypeOf((*MockupgradeMgrHelperAPI)(nil).handleNodeUpgrade), ctx, deviceConfig, node)
}

// isDeviceConfigValid mocks base method.
func (m *MockupgradeMgrHelperAPI) isDeviceConfigValid(ctx context.Context, deviceConfig *v1alpha1.DeviceConfig) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "isDeviceConfigValid", ctx, deviceConfig)
	ret0, _ := ret[0].(bool)
	return ret0
}

// isDeviceConfigValid indicates an expected call of isDeviceConfigValid.
func (mr *MockupgradeMgrHelperAPIMockRecorder) isDeviceConfigValid(ctx, deviceConfig any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "isDeviceConfigValid", reflect.TypeOf((*MockupgradeMgrHelperAPI)(nil).isDeviceConfigValid), ctx, deviceConfig)
}

// isInit mocks base method.
func (m *MockupgradeMgrHelperAPI) isInit() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "isInit")
	ret0, _ := ret[0].(bool)
	return ret0
}

// isInit indicates an expected call of isInit.
func (mr *MockupgradeMgrHelperAPIMockRecorder) isInit() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "isInit", reflect.TypeOf((*MockupgradeMgrHelperAPI)(nil).isInit))
}

// isNodeNew mocks base method.
func (m *MockupgradeMgrHelperAPI) isNodeNew(ctx context.Context, node *v1.Node, deviceConfig *v1alpha1.DeviceConfig) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "isNodeNew", ctx, node, deviceConfig)
	ret0, _ := ret[0].(bool)
	return ret0
}

// isNodeNew indicates an expected call of isNodeNew.
func (mr *MockupgradeMgrHelperAPIMockRecorder) isNodeNew(ctx, node, deviceConfig any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "isNodeNew", reflect.TypeOf((*MockupgradeMgrHelperAPI)(nil).isNodeNew), ctx, node, deviceConfig)
}

// isNodeReady mocks base method.
func (m *MockupgradeMgrHelperAPI) isNodeReady(ctx context.Context, node *v1.Node, deviceConfig *v1alpha1.DeviceConfig) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "isNodeReady", ctx, node, deviceConfig)
	ret0, _ := ret[0].(bool)
	return ret0
}

// isNodeReady indicates an expected call of isNodeReady.
func (mr *MockupgradeMgrHelperAPIMockRecorder) isNodeReady(ctx, node, deviceConfig any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "isNodeReady", reflect.TypeOf((*MockupgradeMgrHelperAPI)(nil).isNodeReady), ctx, node, deviceConfig)
}

// isNodeReadyForUpgrade mocks base method.
func (m *MockupgradeMgrHelperAPI) isNodeReadyForUpgrade(ctx context.Context, node *v1.Node) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "isNodeReadyForUpgrade", ctx, node)
	ret0, _ := ret[0].(bool)
	return ret0
}

// isNodeReadyForUpgrade indicates an expected call of isNodeReadyForUpgrade.
func (mr *MockupgradeMgrHelperAPIMockRecorder) isNodeReadyForUpgrade(ctx, node any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "isNodeReadyForUpgrade", reflect.TypeOf((*MockupgradeMgrHelperAPI)(nil).isNodeReadyForUpgrade), ctx, node)
}

// isNodeStateInstallInProgress mocks base method.
func (m *MockupgradeMgrHelperAPI) isNodeStateInstallInProgress(ctx context.Context, node *v1.Node, deviceConfig *v1alpha1.DeviceConfig) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "isNodeStateInstallInProgress", ctx, node, deviceConfig)
	ret0, _ := ret[0].(bool)
	return ret0
}

// isNodeStateInstallInProgress indicates an expected call of isNodeStateInstallInProgress.
func (mr *MockupgradeMgrHelperAPIMockRecorder) isNodeStateInstallInProgress(ctx, node, deviceConfig any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "isNodeStateInstallInProgress", reflect.TypeOf((*MockupgradeMgrHelperAPI)(nil).isNodeStateInstallInProgress), ctx, node, deviceConfig)
}

// isNodeStateUpgradeFailed mocks base method.
func (m *MockupgradeMgrHelperAPI) isNodeStateUpgradeFailed(ctx context.Context, node *v1.Node, deviceConfig *v1alpha1.DeviceConfig) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "isNodeStateUpgradeFailed", ctx, node, deviceConfig)
	ret0, _ := ret[0].(bool)
	return ret0
}

// isNodeStateUpgradeFailed indicates an expected call of isNodeStateUpgradeFailed.
func (mr *MockupgradeMgrHelperAPIMockRecorder) isNodeStateUpgradeFailed(ctx, node, deviceConfig any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "isNodeStateUpgradeFailed", reflect.TypeOf((*MockupgradeMgrHelperAPI)(nil).isNodeStateUpgradeFailed), ctx, node, deviceConfig)
}

// isNodeStateUpgradeInProgress mocks base method.
func (m *MockupgradeMgrHelperAPI) isNodeStateUpgradeInProgress(ctx context.Context, node *v1.Node, deviceConfig *v1alpha1.DeviceConfig) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "isNodeStateUpgradeInProgress", ctx, node, deviceConfig)
	ret0, _ := ret[0].(bool)
	return ret0
}

// isNodeStateUpgradeInProgress indicates an expected call of isNodeStateUpgradeInProgress.
func (mr *MockupgradeMgrHelperAPIMockRecorder) isNodeStateUpgradeInProgress(ctx, node, deviceConfig any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "isNodeStateUpgradeInProgress", reflect.TypeOf((*MockupgradeMgrHelperAPI)(nil).isNodeStateUpgradeInProgress), ctx, node, deviceConfig)
}

// isNodeStateUpgradeStarted mocks base method.
func (m *MockupgradeMgrHelperAPI) isNodeStateUpgradeStarted(node *v1.Node) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "isNodeStateUpgradeStarted", node)
	ret0, _ := ret[0].(bool)
	return ret0
}

// isNodeStateUpgradeStarted indicates an expected call of isNodeStateUpgradeStarted.
func (mr *MockupgradeMgrHelperAPIMockRecorder) isNodeStateUpgradeStarted(node any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "isNodeStateUpgradeStarted", reflect.TypeOf((*MockupgradeMgrHelperAPI)(nil).isNodeStateUpgradeStarted), node)
}

// isUpgradePolicyViolated mocks base method.
func (m *MockupgradeMgrHelperAPI) isUpgradePolicyViolated(upgradeInProgress, upgradeFailedState, totalNodes int, deviceConfig *v1alpha1.DeviceConfig) (int, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "isUpgradePolicyViolated", upgradeInProgress, upgradeFailedState, totalNodes, deviceConfig)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// isUpgradePolicyViolated indicates an expected call of isUpgradePolicyViolated.
func (mr *MockupgradeMgrHelperAPIMockRecorder) isUpgradePolicyViolated(upgradeInProgress, upgradeFailedState, totalNodes, deviceConfig any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "isUpgradePolicyViolated", reflect.TypeOf((*MockupgradeMgrHelperAPI)(nil).isUpgradePolicyViolated), upgradeInProgress, upgradeFailedState, totalNodes, deviceConfig)
}

// setBootID mocks base method.
func (m *MockupgradeMgrHelperAPI) setBootID(nodeName, bootID string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "setBootID", nodeName, bootID)
}

// setBootID indicates an expected call of setBootID.
func (mr *MockupgradeMgrHelperAPIMockRecorder) setBootID(nodeName, bootID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "setBootID", reflect.TypeOf((*MockupgradeMgrHelperAPI)(nil).setBootID), nodeName, bootID)
}

// setNodeStatus mocks base method.
func (m *MockupgradeMgrHelperAPI) setNodeStatus(ctx context.Context, nodeName string, status v1alpha1.UpgradeState) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "setNodeStatus", ctx, nodeName, status)
}

// setNodeStatus indicates an expected call of setNodeStatus.
func (mr *MockupgradeMgrHelperAPIMockRecorder) setNodeStatus(ctx, nodeName, status any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "setNodeStatus", reflect.TypeOf((*MockupgradeMgrHelperAPI)(nil).setNodeStatus), ctx, nodeName, status)
}

// setUpgradeStartTime mocks base method.
func (m *MockupgradeMgrHelperAPI) setUpgradeStartTime(nodeName string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "setUpgradeStartTime", nodeName)
}

// setUpgradeStartTime indicates an expected call of setUpgradeStartTime.
func (mr *MockupgradeMgrHelperAPIMockRecorder) setUpgradeStartTime(nodeName any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "setUpgradeStartTime", reflect.TypeOf((*MockupgradeMgrHelperAPI)(nil).setUpgradeStartTime), nodeName)
}

// setcurrentSpec mocks base method.
func (m *MockupgradeMgrHelperAPI) setcurrentSpec(deviceConfig *v1alpha1.DeviceConfig) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "setcurrentSpec", deviceConfig)
}

// setcurrentSpec indicates an expected call of setcurrentSpec.
func (mr *MockupgradeMgrHelperAPIMockRecorder) setcurrentSpec(deviceConfig any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "setcurrentSpec", reflect.TypeOf((*MockupgradeMgrHelperAPI)(nil).setcurrentSpec), deviceConfig)
}

// specChanged mocks base method.
func (m *MockupgradeMgrHelperAPI) specChanged(deviceConfig *v1alpha1.DeviceConfig) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "specChanged", deviceConfig)
	ret0, _ := ret[0].(bool)
	return ret0
}

// specChanged indicates an expected call of specChanged.
func (mr *MockupgradeMgrHelperAPIMockRecorder) specChanged(deviceConfig any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "specChanged", reflect.TypeOf((*MockupgradeMgrHelperAPI)(nil).specChanged), deviceConfig)
}

// updateModuleVersionOnNode mocks base method.
func (m *MockupgradeMgrHelperAPI) updateModuleVersionOnNode(ctx context.Context, deviceConfig *v1alpha1.DeviceConfig, node *v1.Node) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "updateModuleVersionOnNode", ctx, deviceConfig, node)
	ret0, _ := ret[0].(error)
	return ret0
}

// updateModuleVersionOnNode indicates an expected call of updateModuleVersionOnNode.
func (mr *MockupgradeMgrHelperAPIMockRecorder) updateModuleVersionOnNode(ctx, deviceConfig, node any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "updateModuleVersionOnNode", reflect.TypeOf((*MockupgradeMgrHelperAPI)(nil).updateModuleVersionOnNode), ctx, deviceConfig, node)
}
