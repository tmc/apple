//go:build darwin

package espresso

import (
	"fmt"
	"runtime"
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
	pe "github.com/tmc/apple/private/espresso"
)

var (
	ErrTrainerInit = fmt.Errorf("espresso: trainer initialization failed")
	ErrTrainerFit  = fmt.Errorf("espresso: training fit failed")
	ErrInference   = fmt.Errorf("espresso: inference failed")
)

// Trainer wraps ETTask for Espresso on-device training.
//
// Create a Trainer with a model definition, optimizer, and loss function,
// then call Fit to train or Predict to run inference.
type Trainer struct {
	task pe.ETTask

	mu     sync.Mutex
	closed bool
}

// TrainerConfig configures trainer creation.
type TrainerConfig struct {
	Model     *ModelDef
	Optimizer *OptimizerDef
	Loss      *LossDef
}

// NewTrainer creates a training task from the given configuration.
func NewTrainer(cfg TrainerConfig) (*Trainer, error) {
	var modelObj objectivec.IObject
	if cfg.Model != nil {
		if cfg.Model.mlp.GetID() != 0 {
			modelObj = objectivec.Object{ID: cfg.Model.mlp.GetID()}
		} else {
			modelObj = objectivec.Object{ID: cfg.Model.raw.GetID()}
		}
	}

	var optObj objectivec.IObject
	if cfg.Optimizer != nil {
		optObj = objectivec.Object{ID: cfg.Optimizer.raw.GetID()}
	}

	var task pe.ETTask
	if cfg.Loss != nil {
		task = pe.NewETTaskWithModelDefOptimizerDefLossConfig(
			modelObj,
			optObj,
			objectivec.Object{ID: cfg.Loss.raw.GetID()},
		)
	} else {
		task = pe.NewETTaskWithModelDefOptimizerDefExtractor(
			modelObj,
			optObj,
			nil, // no extractor
		)
	}

	if task.GetID() == 0 {
		return nil, ErrTrainerInit
	}

	t := &Trainer{task: task}
	runtime.SetFinalizer(t, (*Trainer).Close)
	return t, nil
}

// Fit trains the model for the given number of epochs on the data source.
// The callback is invoked after each batch (may be nil).
func (t *Trainer) Fit(epochs int, data *DataSource, callback func()) error {
	t.mu.Lock()
	if t.closed {
		t.mu.Unlock()
		return ErrClosed
	}
	t.mu.Unlock()

	var cb pe.VoidHandler
	if callback != nil {
		cb = callback
	} else {
		cb = func() {}
	}

	ok := t.task.FitNumberOfEpochsOutputNamesBatchCallback(
		data.obj,
		epochs,
		nil, // all output names
		cb,
	)
	if !ok {
		return ErrTrainerFit
	}
	return nil
}

// FitBatches trains for the given number of batches.
func (t *Trainer) FitBatches(batches uint32, data *DataSource, callback func()) error {
	t.mu.Lock()
	if t.closed {
		t.mu.Unlock()
		return ErrClosed
	}
	t.mu.Unlock()

	var cb pe.VoidHandler
	if callback != nil {
		cb = callback
	} else {
		cb = func() {}
	}

	ok := t.task.FitNumberOfBatchesOutputNamesBatchCallback(
		data.obj,
		batches,
		nil,
		cb,
	)
	if !ok {
		return ErrTrainerFit
	}
	return nil
}

// Predict runs inference on the data source.
// The callback is invoked after each batch (may be nil).
func (t *Trainer) Predict(data *DataSource, callback func()) error {
	t.mu.Lock()
	if t.closed {
		t.mu.Unlock()
		return ErrClosed
	}
	t.mu.Unlock()

	var cb pe.VoidHandler
	if callback != nil {
		cb = callback
	} else {
		cb = func() {}
	}

	ok := t.task.RunInferenceOutputNamesBatchCallback(
		data.obj,
		nil,
		cb,
	)
	if !ok {
		return ErrInference
	}
	return nil
}

// MoveToGPU transfers the model to GPU memory for faster execution.
// The gpu parameter selects the GPU device index (typically 0).
func (t *Trainer) MoveToGPU(gpu int) error {
	t.mu.Lock()
	if t.closed {
		t.mu.Unlock()
		return ErrClosed
	}
	t.mu.Unlock()

	_, err := t.task.MoveToGPUError(gpu)
	if err != nil {
		return fmt.Errorf("espresso trainer: %w", err)
	}
	return nil
}

// SaveWeights saves the current model weights to the given path.
func (t *Trainer) SaveWeights(path string) {
	nsPath := objectivec.Object{ID: objc.String(path)}
	t.task.SaveNetwork(nsPath)
}

// SaveWeightsInferenceMode saves the weights and reverts the network
// to inference mode (removes training-only layers).
func (t *Trainer) SaveWeightsInferenceMode(path string) {
	nsPath := objectivec.Object{ID: objc.String(path)}
	t.task.SaveNetworkRevertToInferenceMode(nsPath, true)
}

// ReinitializeVariables reinitializes all trainable variables.
func (t *Trainer) ReinitializeVariables() {
	t.task.ReinitializeVariables()
}

// DumpData enables or disables data dumping for debugging.
func (t *Trainer) SetDumpData(enabled bool) {
	t.task.SetDumpData(enabled)
}

// Close releases the trainer resources.
func (t *Trainer) Close() error {
	t.mu.Lock()
	defer t.mu.Unlock()
	if t.closed {
		return nil
	}
	t.closed = true
	runtime.SetFinalizer(t, nil)
	return nil
}
