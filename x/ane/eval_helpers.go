//go:build darwin

package ane

import (
	"errors"
	"fmt"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

func (m *Model) qosValue() uint32 {
	if m.qos != 0 {
		return m.qos
	}
	return 21
}

func (m *Model) mapRequestWithFallback(req objectivec.IObject) (func(), error) {
	unmap, err := m.mapRequest(req, true)
	if err == nil {
		return unmap, nil
	}
	return m.mapRequest(req, false)
}

func (m *Model) mapRequest(req objectivec.IObject, cacheInference bool) (func(), error) {
	switch m.modelType {
	case ModelTypeMIL:
		if shared := m.inMemModel.SharedConnection(); shared != nil && shared.GetID() != 0 {
			if base := m.inMemModel.Model(); base != nil && base.GetID() != 0 &&
				respondsToSelector(shared, "mapIOSurfacesWithModel:request:cacheInference:error:") &&
				respondsToSelector(shared, "unmapIOSurfacesWithModel:request:") {
				ok, err := shared.MapIOSurfacesWithModelRequestCacheInferenceError(base, req, cacheInference)
				if err == nil && ok {
					return func() {
						shared.UnmapIOSurfacesWithModelRequest(base, req)
					}, nil
				}
			}
		}
		ok, err := m.inMemModel.MapIOSurfacesWithRequestCacheInferenceError(req, cacheInference)
		if err != nil || !ok {
			return nil, boolNSError("map IOSurfaces failed", ok, err)
		}
		return func() {
			m.inMemModel.UnmapIOSurfacesWithRequest(req)
		}, nil
	case ModelTypePackage:
		ok, err := m.aneClient.MapIOSurfacesWithModelRequestCacheInferenceError(m.aneModel, req, cacheInference)
		if err != nil || !ok {
			return nil, boolNSError("map IOSurfaces failed", ok, err)
		}
		return func() {
			m.aneClient.UnmapIOSurfacesWithModelRequest(m.aneModel, req)
		}, nil
	default:
		return nil, fmt.Errorf("unknown model type %d", m.modelType)
	}
}

func (m *Model) evaluateRequestWithOptions(req objectivec.IObject, opts objectivec.IObject, preferDirect bool) error {
	return m.evaluateRequestWithQoS(req, opts, m.qosValue(), preferDirect)
}

func (m *Model) evaluateRequestWithQoS(req objectivec.IObject, opts objectivec.IObject, qos uint32, preferDirect bool) error {
	switch m.modelType {
	case ModelTypeMIL:
		if preferDirect {
			if shared := m.inMemModel.SharedConnection(); shared != nil && shared.GetID() != 0 {
				if base := m.inMemModel.Model(); base != nil && base.GetID() != 0 &&
					respondsToSelector(shared, "doEvaluateDirectWithModel:options:request:qos:error:") {
					ok, err := shared.DoEvaluateDirectWithModelOptionsRequestQosError(base, opts, req, qos)
					if err == nil && ok {
						return nil
					}
				}
			}
		}
		ok, err := m.inMemModel.EvaluateWithQoSOptionsRequestError(qos, opts, req)
		if err == nil && ok {
			return nil
		}
		return boolNSError("evaluate failed", ok, err)
	case ModelTypePackage:
		if preferDirect && respondsToSelector(m.aneClient, "doEvaluateDirectWithModel:options:request:qos:error:") {
			ok, err := m.aneClient.DoEvaluateDirectWithModelOptionsRequestQosError(m.aneModel, opts, req, qos)
			if err == nil && ok {
				return nil
			}
		}
		ok, err := m.aneClient.EvaluateWithModelOptionsRequestQosError(m.aneModel, opts, req, qos)
		if err == nil && ok {
			return nil
		}
		return boolNSError("evaluate failed", ok, err)
	default:
		return fmt.Errorf("unknown model type %d", m.modelType)
	}
}

func respondsToSelector(obj objectivec.IObject, selector string) bool {
	if obj == nil || obj.GetID() == 0 {
		return false
	}
	return objc.Send[bool](obj.GetID(), objc.Sel("respondsToSelector:"), objc.Sel(selector))
}

func boolNSError(msg string, ok bool, err error) error {
	if err != nil {
		return err
	}
	if ok {
		return nil
	}
	return errors.New(msg)
}
