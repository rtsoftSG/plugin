package toplevel

import (
	"calculator/toolbox/cache"
	"time"
)

//CalculationResult result of single calculation.
type CalculationResult interface {
	//Type return measurement type.
	Type() string
	//Time return time.Time associated with calculation result.
	Time() time.Time
	//CalculatorId return calculator instance unique id (it will be used as ActorName in measurement).
	CalculatorId() string
	//Fields return calculation data.
	Fields() interface{}
}

// Common interface for algorithms.
type Calculator interface {
	Calculate() ([]CalculationResult, error)
}

// CalculatorFactoryFunc creates calculator instance. This function must be realized by every plugin that provides calculator functionality.
//
// NewCalculator function with such signature has to be implemented in order to lookup calculator.
type CalculatorFactoryFunc func(id string, cfg string, cache cache.Storage, db *pgxpool.Pool) (Calculator, error)
