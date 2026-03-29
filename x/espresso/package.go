//go:build darwin

package espresso

import (
	"fmt"

	"github.com/tmc/apple/x/ane"
)

// CompiledPackage owns the ANE client and loaded package-backed model returned
// by CompileAndLoadPackage.
type CompiledPackage struct {
	Client *ane.Client
	Model  *ane.Model
}

// CompileAndLoadPackage uses the practical ANE package-loading path and returns
// a ready-to-evaluate ANE model. This avoids exposing callers to the raw
// Espresso IR constructor path, which requires additional CoreML bundle state.
func CompileAndLoadPackage(path string, opts ane.CompileOptions) (*CompiledPackage, error) {
	client, err := ane.Open()
	if err != nil {
		return nil, fmt.Errorf("espresso compile package: %w", err)
	}
	opts.ModelType = ane.ModelTypePackage
	opts.PackagePath = path
	model, err := client.Compile(opts)
	if err != nil {
		client.Close()
		return nil, fmt.Errorf("espresso compile package: %w", err)
	}
	return &CompiledPackage{
		Client: client,
		Model:  model,
	}, nil
}

// Close releases the model and its client.
func (p *CompiledPackage) Close() error {
	if p == nil {
		return nil
	}
	var err error
	if p.Model != nil {
		err = p.Model.Close()
		p.Model = nil
	}
	if p.Client != nil {
		if closeErr := p.Client.Close(); err == nil {
			err = closeErr
		}
		p.Client = nil
	}
	return err
}
