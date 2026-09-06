// Code generated from internal/generator/templates/runtime/objc.txtar by applegen. DO NOT EDIT.

package objcinspect

import "sync"

// Registry stores known struct definitions indexed by struct tag name,
// allowing opaque struct references (e.g. "{CGRect=}") to be resolved to their
// fully defined field list and natural size when known.
type Registry struct {
	mu      sync.RWMutex
	structs map[string]Type
}

// NewRegistry creates an empty struct definition registry.
func NewRegistry() *Registry {
	return &Registry{
		structs: make(map[string]Type),
	}
}

// Register stores a struct definition. If t is a KindStruct with fields,
// it is saved under t.Name.
func (r *Registry) Register(t Type) {
	if t.Kind != KindStruct || t.Name == "" || len(t.Fields) == 0 {
		return
	}
	r.mu.Lock()
	defer r.mu.Unlock()
	r.structs[t.Name] = t
}

// RegisterEncoding parses encoding and registers any defined struct within it.
func (r *Registry) RegisterEncoding(encoding string) error {
	p := &parser{s: encoding}
	t, err := p.parseType()
	if err != nil {
		return err
	}
	r.registerRecursive(t)
	return nil
}

func (r *Registry) registerRecursive(t Type) {
	if t.Kind == KindStruct && t.Name != "" && len(t.Fields) > 0 {
		r.mu.Lock()
		r.structs[t.Name] = t
		r.mu.Unlock()
	}
	for _, f := range t.Fields {
		r.registerRecursive(f)
	}
	if t.Elem != nil {
		r.registerRecursive(*t.Elem)
	}
}

// Len returns the number of struct definitions held.
func (r *Registry) Len() int {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return len(r.structs)
}

// Resolve returns a copy of t with any opaque struct references resolved
// using registered definitions.
func (r *Registry) Resolve(t Type) Type {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return r.resolveType(t)
}

func (r *Registry) resolveType(t Type) Type {
	if t.Kind == KindStruct {
		if len(t.Fields) == 0 && t.Name != "" {
			if def, ok := r.structs[t.Name]; ok {
				resolved := def
				resolved.FieldName = t.FieldName
				return resolved
			}
		} else if len(t.Fields) > 0 {
			newFields := make([]Type, len(t.Fields))
			for i, f := range t.Fields {
				newFields[i] = r.resolveType(f)
			}
			t.Fields = newFields
		}
	} else if t.Kind == KindUnion && len(t.Fields) > 0 {
		newFields := make([]Type, len(t.Fields))
		for i, f := range t.Fields {
			newFields[i] = r.resolveType(f)
		}
		t.Fields = newFields
	} else if t.Elem != nil {
		elemCopy := r.resolveType(*t.Elem)
		t.Elem = &elemCopy
	}
	return t
}
