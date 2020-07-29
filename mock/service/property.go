// Code generated by MockGen. DO NOT EDIT.
// Source: ./service/property.go

// Package plugin is a generated GoMock package.
package plugin

import (
	models "github.com/baetyl/baetyl-cloud/models"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockPropertyService is a mock of PropertyService interface.
type MockPropertyService struct {
	ctrl     *gomock.Controller
	recorder *MockPropertyServiceMockRecorder
}

// MockPropertyServiceMockRecorder is the mock recorder for MockPropertyService.
type MockPropertyServiceMockRecorder struct {
	mock *MockPropertyService
}

// NewMockPropertyService creates a new mock instance.
func NewMockPropertyService(ctrl *gomock.Controller) *MockPropertyService {
	mock := &MockPropertyService{ctrl: ctrl}
	mock.recorder = &MockPropertyServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPropertyService) EXPECT() *MockPropertyServiceMockRecorder {
	return m.recorder
}

// CreateProperty mocks base method.
func (m *MockPropertyService) CreateProperty(property *models.Property) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateProperty", property)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateProperty indicates an expected call of CreateProperty.
func (mr *MockPropertyServiceMockRecorder) CreateProperty(property interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateProperty", reflect.TypeOf((*MockPropertyService)(nil).CreateProperty), property)
}

// DeleteProperty mocks base method.
func (m *MockPropertyService) DeleteProperty(key string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteProperty", key)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteProperty indicates an expected call of DeleteProperty.
func (mr *MockPropertyServiceMockRecorder) DeleteProperty(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteProperty", reflect.TypeOf((*MockPropertyService)(nil).DeleteProperty), key)
}

// ListProperty mocks base method.
func (m *MockPropertyService) ListProperty(page *models.Filter) ([]models.Property, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListProperty", page)
	ret0, _ := ret[0].([]models.Property)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListProperty indicates an expected call of ListProperty.
func (mr *MockPropertyServiceMockRecorder) ListProperty(page interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListProperty", reflect.TypeOf((*MockPropertyService)(nil).ListProperty), page)
}

// CountProperty mocks base method.
func (m *MockPropertyService) CountProperty(key string) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CountProperty", key)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CountProperty indicates an expected call of CountProperty.
func (mr *MockPropertyServiceMockRecorder) CountProperty(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CountProperty", reflect.TypeOf((*MockPropertyService)(nil).CountProperty), key)
}

// UpdateProperty mocks base method.
func (m *MockPropertyService) UpdateProperty(property *models.Property) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateProperty", property)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateProperty indicates an expected call of UpdateProperty.
func (mr *MockPropertyServiceMockRecorder) UpdateProperty(property interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateProperty", reflect.TypeOf((*MockPropertyService)(nil).UpdateProperty), property)
}

// GetPropertyValue mocks base method.
func (m *MockPropertyService) GetPropertyValue(key string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPropertyValue", key)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPropertyValue indicates an expected call of GetPropertyValue.
func (mr *MockPropertyServiceMockRecorder) GetPropertyValue(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPropertyValue", reflect.TypeOf((*MockPropertyService)(nil).GetPropertyValue), key)
}
