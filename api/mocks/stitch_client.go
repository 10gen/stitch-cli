// Code generated by MockGen. DO NOT EDIT.
// Source: ./api/stitch_client.go

// Package mock_api is a generated GoMock package.
package mock_api

import (
	io "io"
	reflect "reflect"

	api "github.com/10gen/stitch-cli/api"
	auth "github.com/10gen/stitch-cli/auth"
	hosting "github.com/10gen/stitch-cli/hosting"
	models "github.com/10gen/stitch-cli/models"
	secrets "github.com/10gen/stitch-cli/secrets"
	gomock "github.com/golang/mock/gomock"
)

// MockStitchClient is a mock of StitchClient interface
type MockStitchClient struct {
	ctrl     *gomock.Controller
	recorder *MockStitchClientMockRecorder
}

// MockStitchClientMockRecorder is the mock recorder for MockStitchClient
type MockStitchClientMockRecorder struct {
	mock *MockStitchClient
}

// NewMockStitchClient creates a new mock instance
func NewMockStitchClient(ctrl *gomock.Controller) *MockStitchClient {
	mock := &MockStitchClient{ctrl: ctrl}
	mock.recorder = &MockStitchClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockStitchClient) EXPECT() *MockStitchClientMockRecorder {
	return m.recorder
}

// AddSecret mocks base method
func (m *MockStitchClient) AddSecret(groupID, appID string, secret secrets.Secret) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddSecret", groupID, appID, secret)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddSecret indicates an expected call of AddSecret
func (mr *MockStitchClientMockRecorder) AddSecret(groupID, appID, secret interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddSecret", reflect.TypeOf((*MockStitchClient)(nil).AddSecret), groupID, appID, secret)
}

// Authenticate mocks base method
func (m *MockStitchClient) Authenticate(authProvider auth.AuthenticationProvider) (*auth.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Authenticate", authProvider)
	ret0, _ := ret[0].(*auth.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Authenticate indicates an expected call of Authenticate
func (mr *MockStitchClientMockRecorder) Authenticate(authProvider interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Authenticate", reflect.TypeOf((*MockStitchClient)(nil).Authenticate), authProvider)
}

// CopyAsset mocks base method
func (m *MockStitchClient) CopyAsset(groupID, appID, fromPath, toPath string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CopyAsset", groupID, appID, fromPath, toPath)
	ret0, _ := ret[0].(error)
	return ret0
}

// CopyAsset indicates an expected call of CopyAsset
func (mr *MockStitchClientMockRecorder) CopyAsset(groupID, appID, fromPath, toPath interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CopyAsset", reflect.TypeOf((*MockStitchClient)(nil).CopyAsset), groupID, appID, fromPath, toPath)
}

// CreateDraft mocks base method
func (m *MockStitchClient) CreateDraft(groupID, appID string) (*models.AppDraft, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateDraft", groupID, appID)
	ret0, _ := ret[0].(*models.AppDraft)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateDraft indicates an expected call of CreateDraft
func (mr *MockStitchClientMockRecorder) CreateDraft(groupID, appID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateDraft", reflect.TypeOf((*MockStitchClient)(nil).CreateDraft), groupID, appID)
}

// CreateEmptyApp mocks base method
func (m *MockStitchClient) CreateEmptyApp(groupID, appName, location, deploymentModel string) (*models.App, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateEmptyApp", groupID, appName, location, deploymentModel)
	ret0, _ := ret[0].(*models.App)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateEmptyApp indicates an expected call of CreateEmptyApp
func (mr *MockStitchClientMockRecorder) CreateEmptyApp(groupID, appName, location, deploymentModel interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateEmptyApp", reflect.TypeOf((*MockStitchClient)(nil).CreateEmptyApp), groupID, appName, location, deploymentModel)
}

// DeleteAsset mocks base method
func (m *MockStitchClient) DeleteAsset(groupID, appID, path string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAsset", groupID, appID, path)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAsset indicates an expected call of DeleteAsset
func (mr *MockStitchClientMockRecorder) DeleteAsset(groupID, appID, path interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAsset", reflect.TypeOf((*MockStitchClient)(nil).DeleteAsset), groupID, appID, path)
}

// DeployDraft mocks base method
func (m *MockStitchClient) DeployDraft(groupID, appID, draftID string) (*models.Deployment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeployDraft", groupID, appID, draftID)
	ret0, _ := ret[0].(*models.Deployment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeployDraft indicates an expected call of DeployDraft
func (mr *MockStitchClientMockRecorder) DeployDraft(groupID, appID, draftID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeployDraft", reflect.TypeOf((*MockStitchClient)(nil).DeployDraft), groupID, appID, draftID)
}

// Diff mocks base method
func (m *MockStitchClient) Diff(groupID, appID string, appData []byte, strategy string) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Diff", groupID, appID, appData, strategy)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Diff indicates an expected call of Diff
func (mr *MockStitchClientMockRecorder) Diff(groupID, appID, appData, strategy interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Diff", reflect.TypeOf((*MockStitchClient)(nil).Diff), groupID, appID, appData, strategy)
}

// DiscardDraft mocks base method
func (m *MockStitchClient) DiscardDraft(groupID, appID, draftID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DiscardDraft", groupID, appID, draftID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DiscardDraft indicates an expected call of DiscardDraft
func (mr *MockStitchClientMockRecorder) DiscardDraft(groupID, appID, draftID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DiscardDraft", reflect.TypeOf((*MockStitchClient)(nil).DiscardDraft), groupID, appID, draftID)
}

// DraftDiff mocks base method
func (m *MockStitchClient) DraftDiff(groupID, appID, draftID string) (*models.DraftDiff, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DraftDiff", groupID, appID, draftID)
	ret0, _ := ret[0].(*models.DraftDiff)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DraftDiff indicates an expected call of DraftDiff
func (mr *MockStitchClientMockRecorder) DraftDiff(groupID, appID, draftID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DraftDiff", reflect.TypeOf((*MockStitchClient)(nil).DraftDiff), groupID, appID, draftID)
}

// Export mocks base method
func (m *MockStitchClient) Export(groupID, appID string, strategy api.ExportStrategy) (string, io.ReadCloser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Export", groupID, appID, strategy)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(io.ReadCloser)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Export indicates an expected call of Export
func (mr *MockStitchClientMockRecorder) Export(groupID, appID, strategy interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Export", reflect.TypeOf((*MockStitchClient)(nil).Export), groupID, appID, strategy)
}

// FetchAppByClientAppID mocks base method
func (m *MockStitchClient) FetchAppByClientAppID(clientAppID string) (*models.App, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchAppByClientAppID", clientAppID)
	ret0, _ := ret[0].(*models.App)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchAppByClientAppID indicates an expected call of FetchAppByClientAppID
func (mr *MockStitchClientMockRecorder) FetchAppByClientAppID(clientAppID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchAppByClientAppID", reflect.TypeOf((*MockStitchClient)(nil).FetchAppByClientAppID), clientAppID)
}

// FetchAppByGroupIDAndClientAppID mocks base method
func (m *MockStitchClient) FetchAppByGroupIDAndClientAppID(groupID, clientAppID string) (*models.App, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchAppByGroupIDAndClientAppID", groupID, clientAppID)
	ret0, _ := ret[0].(*models.App)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchAppByGroupIDAndClientAppID indicates an expected call of FetchAppByGroupIDAndClientAppID
func (mr *MockStitchClientMockRecorder) FetchAppByGroupIDAndClientAppID(groupID, clientAppID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchAppByGroupIDAndClientAppID", reflect.TypeOf((*MockStitchClient)(nil).FetchAppByGroupIDAndClientAppID), groupID, clientAppID)
}

// FetchAppsByGroupID mocks base method
func (m *MockStitchClient) FetchAppsByGroupID(groupID string) ([]*models.App, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchAppsByGroupID", groupID)
	ret0, _ := ret[0].([]*models.App)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchAppsByGroupID indicates an expected call of FetchAppsByGroupID
func (mr *MockStitchClientMockRecorder) FetchAppsByGroupID(groupID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchAppsByGroupID", reflect.TypeOf((*MockStitchClient)(nil).FetchAppsByGroupID), groupID)
}

// GetDeployment mocks base method
func (m *MockStitchClient) GetDeployment(groupID, appID, deploymentID string) (*models.Deployment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDeployment", groupID, appID, deploymentID)
	ret0, _ := ret[0].(*models.Deployment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDeployment indicates an expected call of GetDeployment
func (mr *MockStitchClientMockRecorder) GetDeployment(groupID, appID, deploymentID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDeployment", reflect.TypeOf((*MockStitchClient)(nil).GetDeployment), groupID, appID, deploymentID)
}

// GetDrafts mocks base method
func (m *MockStitchClient) GetDrafts(groupID, appID string) ([]models.AppDraft, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDrafts", groupID, appID)
	ret0, _ := ret[0].([]models.AppDraft)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDrafts indicates an expected call of GetDrafts
func (mr *MockStitchClientMockRecorder) GetDrafts(groupID, appID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDrafts", reflect.TypeOf((*MockStitchClient)(nil).GetDrafts), groupID, appID)
}

// Import mocks base method
func (m *MockStitchClient) Import(groupID, appID string, appData []byte, strategy string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Import", groupID, appID, appData, strategy)
	ret0, _ := ret[0].(error)
	return ret0
}

// Import indicates an expected call of Import
func (mr *MockStitchClientMockRecorder) Import(groupID, appID, appData, strategy interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Import", reflect.TypeOf((*MockStitchClient)(nil).Import), groupID, appID, appData, strategy)
}

// InvalidateCache mocks base method
func (m *MockStitchClient) InvalidateCache(groupID, appID, path string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InvalidateCache", groupID, appID, path)
	ret0, _ := ret[0].(error)
	return ret0
}

// InvalidateCache indicates an expected call of InvalidateCache
func (mr *MockStitchClientMockRecorder) InvalidateCache(groupID, appID, path interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InvalidateCache", reflect.TypeOf((*MockStitchClient)(nil).InvalidateCache), groupID, appID, path)
}

// ListAssetsForAppID mocks base method
func (m *MockStitchClient) ListAssetsForAppID(groupID, appID string) ([]hosting.AssetMetadata, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAssetsForAppID", groupID, appID)
	ret0, _ := ret[0].([]hosting.AssetMetadata)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAssetsForAppID indicates an expected call of ListAssetsForAppID
func (mr *MockStitchClientMockRecorder) ListAssetsForAppID(groupID, appID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAssetsForAppID", reflect.TypeOf((*MockStitchClient)(nil).ListAssetsForAppID), groupID, appID)
}

// ListSecrets mocks base method
func (m *MockStitchClient) ListSecrets(groupID, appID string) ([]secrets.Secret, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListSecrets", groupID, appID)
	ret0, _ := ret[0].([]secrets.Secret)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListSecrets indicates an expected call of ListSecrets
func (mr *MockStitchClientMockRecorder) ListSecrets(groupID, appID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListSecrets", reflect.TypeOf((*MockStitchClient)(nil).ListSecrets), groupID, appID)
}

// MoveAsset mocks base method
func (m *MockStitchClient) MoveAsset(groupID, appID, fromPath, toPath string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MoveAsset", groupID, appID, fromPath, toPath)
	ret0, _ := ret[0].(error)
	return ret0
}

// MoveAsset indicates an expected call of MoveAsset
func (mr *MockStitchClientMockRecorder) MoveAsset(groupID, appID, fromPath, toPath interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MoveAsset", reflect.TypeOf((*MockStitchClient)(nil).MoveAsset), groupID, appID, fromPath, toPath)
}

// RemoveSecretByID mocks base method
func (m *MockStitchClient) RemoveSecretByID(groupID, appID, secretID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveSecretByID", groupID, appID, secretID)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveSecretByID indicates an expected call of RemoveSecretByID
func (mr *MockStitchClientMockRecorder) RemoveSecretByID(groupID, appID, secretID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveSecretByID", reflect.TypeOf((*MockStitchClient)(nil).RemoveSecretByID), groupID, appID, secretID)
}

// RemoveSecretByName mocks base method
func (m *MockStitchClient) RemoveSecretByName(groupID, appID, secretName string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveSecretByName", groupID, appID, secretName)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveSecretByName indicates an expected call of RemoveSecretByName
func (mr *MockStitchClientMockRecorder) RemoveSecretByName(groupID, appID, secretName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveSecretByName", reflect.TypeOf((*MockStitchClient)(nil).RemoveSecretByName), groupID, appID, secretName)
}

// SetAssetAttributes mocks base method
func (m *MockStitchClient) SetAssetAttributes(groupID, appID, path string, attributes ...hosting.AssetAttribute) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{groupID, appID, path}
	for _, a := range attributes {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SetAssetAttributes", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetAssetAttributes indicates an expected call of SetAssetAttributes
func (mr *MockStitchClientMockRecorder) SetAssetAttributes(groupID, appID, path interface{}, attributes ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{groupID, appID, path}, attributes...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetAssetAttributes", reflect.TypeOf((*MockStitchClient)(nil).SetAssetAttributes), varargs...)
}

// UpdateSecretByID mocks base method
func (m *MockStitchClient) UpdateSecretByID(groupID, appID, secretID, secretValue string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateSecretByID", groupID, appID, secretID, secretValue)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateSecretByID indicates an expected call of UpdateSecretByID
func (mr *MockStitchClientMockRecorder) UpdateSecretByID(groupID, appID, secretID, secretValue interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateSecretByID", reflect.TypeOf((*MockStitchClient)(nil).UpdateSecretByID), groupID, appID, secretID, secretValue)
}

// UpdateSecretByName mocks base method
func (m *MockStitchClient) UpdateSecretByName(groupID, appID, secretName, secretValue string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateSecretByName", groupID, appID, secretName, secretValue)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateSecretByName indicates an expected call of UpdateSecretByName
func (mr *MockStitchClientMockRecorder) UpdateSecretByName(groupID, appID, secretName, secretValue interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateSecretByName", reflect.TypeOf((*MockStitchClient)(nil).UpdateSecretByName), groupID, appID, secretName, secretValue)
}

// UploadAsset mocks base method
func (m *MockStitchClient) UploadAsset(groupID, appID, path, hash string, size int64, body io.Reader, attributes ...hosting.AssetAttribute) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{groupID, appID, path, hash, size, body}
	for _, a := range attributes {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UploadAsset", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// UploadAsset indicates an expected call of UploadAsset
func (mr *MockStitchClientMockRecorder) UploadAsset(groupID, appID, path, hash, size, body interface{}, attributes ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{groupID, appID, path, hash, size, body}, attributes...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UploadAsset", reflect.TypeOf((*MockStitchClient)(nil).UploadAsset), varargs...)
}

// UploadDependencies mocks base method
func (m *MockStitchClient) UploadDependencies(groupID, appID, fullPath string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UploadDependencies", groupID, appID, fullPath)
	ret0, _ := ret[0].(error)
	return ret0
}

// UploadDependencies indicates an expected call of UploadDependencies
func (mr *MockStitchClientMockRecorder) UploadDependencies(groupID, appID, fullPath interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UploadDependencies", reflect.TypeOf((*MockStitchClient)(nil).UploadDependencies), groupID, appID, fullPath)
}
