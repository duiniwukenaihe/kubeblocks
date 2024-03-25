// /*
// Copyright (C) 2022-2023 ApeCloud Co., Ltd
//
// This file is part of KubeBlocks project
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.
// */
//
//

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/apecloud/kubeblocks/pkg/lorry/engines (interfaces: DBManager)

// Package engines is a generated GoMock package.
package engines

import (
	context "context"
	reflect "reflect"

	dcs "github.com/apecloud/kubeblocks/pkg/lorry/dcs"
	models "github.com/apecloud/kubeblocks/pkg/lorry/engines/models"
	logr "github.com/go-logr/logr"
	gomock "github.com/golang/mock/gomock"
)

// MockDBManager is a mock of DBManager interface.
type MockDBManager struct {
	ctrl     *gomock.Controller
	recorder *MockDBManagerMockRecorder
}

// MockDBManagerMockRecorder is the mock recorder for MockDBManager.
type MockDBManagerMockRecorder struct {
	mock *MockDBManager
}

// NewMockDBManager creates a new mock instance.
func NewMockDBManager(ctrl *gomock.Controller) *MockDBManager {
	mock := &MockDBManager{ctrl: ctrl}
	mock.recorder = &MockDBManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDBManager) EXPECT() *MockDBManagerMockRecorder {
	return m.recorder
}

// CreateRoot mocks base method.
func (m *MockDBManager) CreateRoot(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateRoot", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateRoot indicates an expected call of CreateRoot.
func (mr *MockDBManagerMockRecorder) CreateRoot(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateRoot", reflect.TypeOf((*MockDBManager)(nil).CreateRoot), arg0)
}

// CreateUser mocks base method.
func (m *MockDBManager) CreateUser(arg0 context.Context, arg1, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockDBManagerMockRecorder) CreateUser(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockDBManager)(nil).CreateUser), arg0, arg1, arg2)
}

// CurrentMemberHealthyCheck mocks base method.
func (m *MockDBManager) CurrentMemberHealthyCheck(arg0 context.Context, arg1 *dcs.Cluster) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CurrentMemberHealthyCheck", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CurrentMemberHealthyCheck indicates an expected call of CurrentMemberHealthyCheck.
func (mr *MockDBManagerMockRecorder) CurrentMemberHealthyCheck(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CurrentMemberHealthyCheck", reflect.TypeOf((*MockDBManager)(nil).CurrentMemberHealthyCheck), arg0, arg1)
}

// DeleteUser mocks base method.
func (m *MockDBManager) DeleteUser(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUser", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteUser indicates an expected call of DeleteUser.
func (mr *MockDBManagerMockRecorder) DeleteUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUser", reflect.TypeOf((*MockDBManager)(nil).DeleteUser), arg0, arg1)
}

// Demote mocks base method.
func (m *MockDBManager) Demote(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Demote", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Demote indicates an expected call of Demote.
func (mr *MockDBManagerMockRecorder) Demote(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Demote", reflect.TypeOf((*MockDBManager)(nil).Demote), arg0)
}

// DescribeUser mocks base method.
func (m *MockDBManager) DescribeUser(arg0 context.Context, arg1 string) (*models.UserInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeUser", arg0, arg1)
	ret0, _ := ret[0].(*models.UserInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeUser indicates an expected call of DescribeUser.
func (mr *MockDBManagerMockRecorder) DescribeUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeUser", reflect.TypeOf((*MockDBManager)(nil).DescribeUser), arg0, arg1)
}

// Exec mocks base method.
func (m *MockDBManager) Exec(arg0 context.Context, arg1 string) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Exec", arg0, arg1)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Exec indicates an expected call of Exec.
func (mr *MockDBManagerMockRecorder) Exec(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exec", reflect.TypeOf((*MockDBManager)(nil).Exec), arg0, arg1)
}

// Follow mocks base method.
func (m *MockDBManager) Follow(arg0 context.Context, arg1 *dcs.Cluster) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Follow", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Follow indicates an expected call of Follow.
func (mr *MockDBManagerMockRecorder) Follow(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Follow", reflect.TypeOf((*MockDBManager)(nil).Follow), arg0, arg1)
}

// GetCurrentMemberName mocks base method.
func (m *MockDBManager) GetCurrentMemberName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCurrentMemberName")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetCurrentMemberName indicates an expected call of GetCurrentMemberName.
func (mr *MockDBManagerMockRecorder) GetCurrentMemberName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCurrentMemberName", reflect.TypeOf((*MockDBManager)(nil).GetCurrentMemberName))
}

// GetDBState mocks base method.
func (m *MockDBManager) GetDBState(arg0 context.Context, arg1 *dcs.Cluster) *dcs.DBState {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDBState", arg0, arg1)
	ret0, _ := ret[0].(*dcs.DBState)
	return ret0
}

// GetDBState indicates an expected call of GetDBState.
func (mr *MockDBManagerMockRecorder) GetDBState(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDBState", reflect.TypeOf((*MockDBManager)(nil).GetDBState), arg0, arg1)
}

// GetLag mocks base method.
func (m *MockDBManager) GetLag(arg0 context.Context, arg1 *dcs.Cluster) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLag", arg0, arg1)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLag indicates an expected call of GetLag.
func (mr *MockDBManagerMockRecorder) GetLag(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLag", reflect.TypeOf((*MockDBManager)(nil).GetLag), arg0, arg1)
}

// GetLogger mocks base method.
func (m *MockDBManager) GetLogger() logr.Logger {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLogger")
	ret0, _ := ret[0].(logr.Logger)
	return ret0
}

// GetLogger indicates an expected call of GetLogger.
func (mr *MockDBManagerMockRecorder) GetLogger() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLogger", reflect.TypeOf((*MockDBManager)(nil).GetLogger))
}

// GetMemberAddrs mocks base method.
func (m *MockDBManager) GetMemberAddrs(arg0 context.Context, arg1 *dcs.Cluster) []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMemberAddrs", arg0, arg1)
	ret0, _ := ret[0].([]string)
	return ret0
}

// GetMemberAddrs indicates an expected call of GetMemberAddrs.
func (mr *MockDBManagerMockRecorder) GetMemberAddrs(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMemberAddrs", reflect.TypeOf((*MockDBManager)(nil).GetMemberAddrs), arg0, arg1)
}

// GetPort mocks base method.
func (m *MockDBManager) GetPort() (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPort")
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPort indicates an expected call of GetPort.
func (mr *MockDBManagerMockRecorder) GetPort() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPort", reflect.TypeOf((*MockDBManager)(nil).GetPort))
}

// GetReplicaRole mocks base method.
func (m *MockDBManager) GetReplicaRole(arg0 context.Context, arg1 *dcs.Cluster) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetReplicaRole", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetReplicaRole indicates an expected call of GetReplicaRole.
func (mr *MockDBManagerMockRecorder) GetReplicaRole(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetReplicaRole", reflect.TypeOf((*MockDBManager)(nil).GetReplicaRole), arg0, arg1)
}

// GrantUserRole mocks base method.
func (m *MockDBManager) GrantUserRole(arg0 context.Context, arg1, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GrantUserRole", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// GrantUserRole indicates an expected call of GrantUserRole.
func (mr *MockDBManagerMockRecorder) GrantUserRole(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GrantUserRole", reflect.TypeOf((*MockDBManager)(nil).GrantUserRole), arg0, arg1, arg2)
}

// HasOtherHealthyLeader mocks base method.
func (m *MockDBManager) HasOtherHealthyLeader(arg0 context.Context, arg1 *dcs.Cluster) *dcs.Member {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasOtherHealthyLeader", arg0, arg1)
	ret0, _ := ret[0].(*dcs.Member)
	return ret0
}

// HasOtherHealthyLeader indicates an expected call of HasOtherHealthyLeader.
func (mr *MockDBManagerMockRecorder) HasOtherHealthyLeader(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasOtherHealthyLeader", reflect.TypeOf((*MockDBManager)(nil).HasOtherHealthyLeader), arg0, arg1)
}

// HasOtherHealthyMembers mocks base method.
func (m *MockDBManager) HasOtherHealthyMembers(arg0 context.Context, arg1 *dcs.Cluster, arg2 string) []*dcs.Member {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasOtherHealthyMembers", arg0, arg1, arg2)
	ret0, _ := ret[0].([]*dcs.Member)
	return ret0
}

// HasOtherHealthyMembers indicates an expected call of HasOtherHealthyMembers.
func (mr *MockDBManagerMockRecorder) HasOtherHealthyMembers(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasOtherHealthyMembers", reflect.TypeOf((*MockDBManager)(nil).HasOtherHealthyMembers), arg0, arg1, arg2)
}

// InitializeCluster mocks base method.
func (m *MockDBManager) InitializeCluster(arg0 context.Context, arg1 *dcs.Cluster) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InitializeCluster", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// InitializeCluster indicates an expected call of InitializeCluster.
func (mr *MockDBManagerMockRecorder) InitializeCluster(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InitializeCluster", reflect.TypeOf((*MockDBManager)(nil).InitializeCluster), arg0, arg1)
}

// IsClusterHealthy mocks base method.
func (m *MockDBManager) IsClusterHealthy(arg0 context.Context, arg1 *dcs.Cluster) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsClusterHealthy", arg0, arg1)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsClusterHealthy indicates an expected call of IsClusterHealthy.
func (mr *MockDBManagerMockRecorder) IsClusterHealthy(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsClusterHealthy", reflect.TypeOf((*MockDBManager)(nil).IsClusterHealthy), arg0, arg1)
}

// IsClusterInitialized mocks base method.
func (m *MockDBManager) IsClusterInitialized(arg0 context.Context, arg1 *dcs.Cluster) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsClusterInitialized", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsClusterInitialized indicates an expected call of IsClusterInitialized.
func (mr *MockDBManagerMockRecorder) IsClusterInitialized(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsClusterInitialized", reflect.TypeOf((*MockDBManager)(nil).IsClusterInitialized), arg0, arg1)
}

// IsCurrentMemberHealthy mocks base method.
func (m *MockDBManager) IsCurrentMemberHealthy(arg0 context.Context, arg1 *dcs.Cluster) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsCurrentMemberHealthy", arg0, arg1)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsCurrentMemberHealthy indicates an expected call of IsCurrentMemberHealthy.
func (mr *MockDBManagerMockRecorder) IsCurrentMemberHealthy(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsCurrentMemberHealthy", reflect.TypeOf((*MockDBManager)(nil).IsCurrentMemberHealthy), arg0, arg1)
}

// IsCurrentMemberInCluster mocks base method.
func (m *MockDBManager) IsCurrentMemberInCluster(arg0 context.Context, arg1 *dcs.Cluster) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsCurrentMemberInCluster", arg0, arg1)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsCurrentMemberInCluster indicates an expected call of IsCurrentMemberInCluster.
func (mr *MockDBManagerMockRecorder) IsCurrentMemberInCluster(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsCurrentMemberInCluster", reflect.TypeOf((*MockDBManager)(nil).IsCurrentMemberInCluster), arg0, arg1)
}

// IsDBStartupReady mocks base method.
func (m *MockDBManager) IsDBStartupReady() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsDBStartupReady")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsDBStartupReady indicates an expected call of IsDBStartupReady.
func (mr *MockDBManagerMockRecorder) IsDBStartupReady() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsDBStartupReady", reflect.TypeOf((*MockDBManager)(nil).IsDBStartupReady))
}

// IsFirstMember mocks base method.
func (m *MockDBManager) IsFirstMember() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsFirstMember")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsFirstMember indicates an expected call of IsFirstMember.
func (mr *MockDBManagerMockRecorder) IsFirstMember() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsFirstMember", reflect.TypeOf((*MockDBManager)(nil).IsFirstMember))
}

// IsLeader mocks base method.
func (m *MockDBManager) IsLeader(arg0 context.Context, arg1 *dcs.Cluster) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsLeader", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsLeader indicates an expected call of IsLeader.
func (mr *MockDBManagerMockRecorder) IsLeader(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsLeader", reflect.TypeOf((*MockDBManager)(nil).IsLeader), arg0, arg1)
}

// IsLeaderMember mocks base method.
func (m *MockDBManager) IsLeaderMember(arg0 context.Context, arg1 *dcs.Cluster, arg2 *dcs.Member) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsLeaderMember", arg0, arg1, arg2)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsLeaderMember indicates an expected call of IsLeaderMember.
func (mr *MockDBManagerMockRecorder) IsLeaderMember(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsLeaderMember", reflect.TypeOf((*MockDBManager)(nil).IsLeaderMember), arg0, arg1, arg2)
}

// IsMemberHealthy mocks base method.
func (m *MockDBManager) IsMemberHealthy(arg0 context.Context, arg1 *dcs.Cluster, arg2 *dcs.Member) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsMemberHealthy", arg0, arg1, arg2)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsMemberHealthy indicates an expected call of IsMemberHealthy.
func (mr *MockDBManagerMockRecorder) IsMemberHealthy(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsMemberHealthy", reflect.TypeOf((*MockDBManager)(nil).IsMemberHealthy), arg0, arg1, arg2)
}

// IsMemberLagging mocks base method.
func (m *MockDBManager) IsMemberLagging(arg0 context.Context, arg1 *dcs.Cluster, arg2 *dcs.Member) (bool, int64) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsMemberLagging", arg0, arg1, arg2)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(int64)
	return ret0, ret1
}

// IsMemberLagging indicates an expected call of IsMemberLagging.
func (mr *MockDBManagerMockRecorder) IsMemberLagging(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsMemberLagging", reflect.TypeOf((*MockDBManager)(nil).IsMemberLagging), arg0, arg1, arg2)
}

// IsPromoted mocks base method.
func (m *MockDBManager) IsPromoted(arg0 context.Context) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsPromoted", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsPromoted indicates an expected call of IsPromoted.
func (mr *MockDBManagerMockRecorder) IsPromoted(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsPromoted", reflect.TypeOf((*MockDBManager)(nil).IsPromoted), arg0)
}

// IsRootCreated mocks base method.
func (m *MockDBManager) IsRootCreated(arg0 context.Context) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsRootCreated", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsRootCreated indicates an expected call of IsRootCreated.
func (mr *MockDBManagerMockRecorder) IsRootCreated(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsRootCreated", reflect.TypeOf((*MockDBManager)(nil).IsRootCreated), arg0)
}

// IsRunning mocks base method.
func (m *MockDBManager) IsRunning() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsRunning")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsRunning indicates an expected call of IsRunning.
func (mr *MockDBManagerMockRecorder) IsRunning() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsRunning", reflect.TypeOf((*MockDBManager)(nil).IsRunning))
}

// JoinCurrentMemberToCluster mocks base method.
func (m *MockDBManager) JoinCurrentMemberToCluster(arg0 context.Context, arg1 *dcs.Cluster) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "JoinCurrentMemberToCluster", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// JoinCurrentMemberToCluster indicates an expected call of JoinCurrentMemberToCluster.
func (mr *MockDBManagerMockRecorder) JoinCurrentMemberToCluster(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "JoinCurrentMemberToCluster", reflect.TypeOf((*MockDBManager)(nil).JoinCurrentMemberToCluster), arg0, arg1)
}

// LeaderHealthyCheck mocks base method.
func (m *MockDBManager) LeaderHealthyCheck(arg0 context.Context, arg1 *dcs.Cluster) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LeaderHealthyCheck", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// LeaderHealthyCheck indicates an expected call of LeaderHealthyCheck.
func (mr *MockDBManagerMockRecorder) LeaderHealthyCheck(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LeaderHealthyCheck", reflect.TypeOf((*MockDBManager)(nil).LeaderHealthyCheck), arg0, arg1)
}

// LeaveMemberFromCluster mocks base method.
func (m *MockDBManager) LeaveMemberFromCluster(arg0 context.Context, arg1 *dcs.Cluster, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LeaveMemberFromCluster", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// LeaveMemberFromCluster indicates an expected call of LeaveMemberFromCluster.
func (mr *MockDBManagerMockRecorder) LeaveMemberFromCluster(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LeaveMemberFromCluster", reflect.TypeOf((*MockDBManager)(nil).LeaveMemberFromCluster), arg0, arg1, arg2)
}

// ListSystemAccounts mocks base method.
func (m *MockDBManager) ListSystemAccounts(arg0 context.Context) ([]models.UserInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListSystemAccounts", arg0)
	ret0, _ := ret[0].([]models.UserInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListSystemAccounts indicates an expected call of ListSystemAccounts.
func (mr *MockDBManagerMockRecorder) ListSystemAccounts(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListSystemAccounts", reflect.TypeOf((*MockDBManager)(nil).ListSystemAccounts), arg0)
}

// ListUsers mocks base method.
func (m *MockDBManager) ListUsers(arg0 context.Context) ([]models.UserInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListUsers", arg0)
	ret0, _ := ret[0].([]models.UserInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListUsers indicates an expected call of ListUsers.
func (mr *MockDBManagerMockRecorder) ListUsers(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListUsers", reflect.TypeOf((*MockDBManager)(nil).ListUsers), arg0)
}

// Lock mocks base method.
func (m *MockDBManager) Lock(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Lock", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Lock indicates an expected call of Lock.
func (mr *MockDBManagerMockRecorder) Lock(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Lock", reflect.TypeOf((*MockDBManager)(nil).Lock), arg0, arg1)
}

// MemberHealthyCheck mocks base method.
func (m *MockDBManager) MemberHealthyCheck(arg0 context.Context, arg1 *dcs.Cluster, arg2 *dcs.Member) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MemberHealthyCheck", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// MemberHealthyCheck indicates an expected call of MemberHealthyCheck.
func (mr *MockDBManagerMockRecorder) MemberHealthyCheck(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MemberHealthyCheck", reflect.TypeOf((*MockDBManager)(nil).MemberHealthyCheck), arg0, arg1, arg2)
}

// MoveData mocks base method.
func (m *MockDBManager) MoveData(arg0 context.Context, arg1 *dcs.Cluster) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MoveData", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// MoveData indicates an expected call of MoveData.
func (mr *MockDBManagerMockRecorder) MoveData(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MoveData", reflect.TypeOf((*MockDBManager)(nil).MoveData), arg0, arg1)
}

// Promote mocks base method.
func (m *MockDBManager) Promote(arg0 context.Context, arg1 *dcs.Cluster) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Promote", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Promote indicates an expected call of Promote.
func (mr *MockDBManagerMockRecorder) Promote(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Promote", reflect.TypeOf((*MockDBManager)(nil).Promote), arg0, arg1)
}

// Query mocks base method.
func (m *MockDBManager) Query(arg0 context.Context, arg1 string) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Query", arg0, arg1)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Query indicates an expected call of Query.
func (mr *MockDBManagerMockRecorder) Query(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Query", reflect.TypeOf((*MockDBManager)(nil).Query), arg0, arg1)
}

// Recover mocks base method.
func (m *MockDBManager) Recover(arg0 context.Context, arg1 *dcs.Cluster) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Recover", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Recover indicates an expected call of Recover.
func (mr *MockDBManagerMockRecorder) Recover(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Recover", reflect.TypeOf((*MockDBManager)(nil).Recover), arg0, arg1)
}

// RevokeUserRole mocks base method.
func (m *MockDBManager) RevokeUserRole(arg0 context.Context, arg1, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RevokeUserRole", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// RevokeUserRole indicates an expected call of RevokeUserRole.
func (mr *MockDBManagerMockRecorder) RevokeUserRole(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RevokeUserRole", reflect.TypeOf((*MockDBManager)(nil).RevokeUserRole), arg0, arg1, arg2)
}

// ShutDownWithWait mocks base method.
func (m *MockDBManager) ShutDownWithWait() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ShutDownWithWait")
}

// ShutDownWithWait indicates an expected call of ShutDownWithWait.
func (mr *MockDBManagerMockRecorder) ShutDownWithWait() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ShutDownWithWait", reflect.TypeOf((*MockDBManager)(nil).ShutDownWithWait))
}

// Start mocks base method.
func (m *MockDBManager) Start(arg0 context.Context, arg1 *dcs.Cluster) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Start", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Start indicates an expected call of Start.
func (mr *MockDBManagerMockRecorder) Start(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockDBManager)(nil).Start), arg0, arg1)
}

// Stop mocks base method.
func (m *MockDBManager) Stop() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Stop")
	ret0, _ := ret[0].(error)
	return ret0
}

// Stop indicates an expected call of Stop.
func (mr *MockDBManagerMockRecorder) Stop() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockDBManager)(nil).Stop))
}

// Unlock mocks base method.
func (m *MockDBManager) Unlock(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Unlock", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Unlock indicates an expected call of Unlock.
func (mr *MockDBManagerMockRecorder) Unlock(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unlock", reflect.TypeOf((*MockDBManager)(nil).Unlock), arg0)
}
