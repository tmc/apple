package espresso

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
)

// LayerShape describes the tensor shape for a layer blob as parsed from
// model.espresso.shape JSON.
type LayerShape struct {
	Name string // blob name
	N    int    // batch
	C    int    // channels
	H    int    // height
	W    int    // width
}

// Elements returns the total number of elements in this shape.
func (s LayerShape) Elements() int {
	n := s.C * s.H * s.W
	if s.N > 0 {
		n *= s.N
	}
	return n
}

// ShapeMap holds parsed shape information for all blobs in a model.
type ShapeMap struct {
	Shapes map[string]LayerShape
}

// Get returns the shape for a named blob.
func (m *ShapeMap) Get(name string) (LayerShape, bool) {
	s, ok := m.Shapes[name]
	return s, ok
}

// Names returns sorted blob names.
func (m *ShapeMap) Names() []string {
	names := make([]string, 0, len(m.Shapes))
	for k := range m.Shapes {
		names = append(names, k)
	}
	sort.Strings(names)
	return names
}

// ReadEspressoShape parses a model.espresso.shape JSON file into a ShapeMap.
//
// The shape JSON has the structure:
//
//	{
//	  "layer_shapes": {
//	    "blob_name": { "n": 1, "k": 64, "h": 1, "w": 128 },
//	    ...
//	  }
//	}
//
// where k=channels, h=height, w=width, n=batch.
func ReadEspressoShape(data []byte) (*ShapeMap, error) {
	var raw struct {
		LayerShapes map[string]json.RawMessage `json:"layer_shapes"`
	}
	if err := json.Unmarshal(data, &raw); err != nil {
		return nil, fmt.Errorf("espresso shape: %w", err)
	}

	m := &ShapeMap{Shapes: make(map[string]LayerShape, len(raw.LayerShapes))}
	for name, entry := range raw.LayerShapes {
		var dims struct {
			N int `json:"n"`
			K int `json:"k"`
			H int `json:"h"`
			W int `json:"w"`
		}
		if err := json.Unmarshal(entry, &dims); err != nil {
			// Empty shape entries are valid (shapes computed at runtime).
			m.Shapes[name] = LayerShape{Name: name}
			continue
		}
		m.Shapes[name] = LayerShape{
			Name: name,
			N:    dims.N,
			C:    dims.K,
			H:    dims.H,
			W:    dims.W,
		}
	}
	return m, nil
}

// ReadEspressoShapeFile reads and parses a model.espresso.shape JSON file.
func ReadEspressoShapeFile(path string) (*ShapeMap, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("espresso shape: %w", err)
	}
	return ReadEspressoShape(data)
}
