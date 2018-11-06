package delay

import (
	"math"
	"math/rand"
	"testing"
	"time"
)

const testSeed = 99

func TestDelaySetAndGet(t *testing.T) {
	initialValue, err := time.ParseDuration("1000ms")

	if err != nil {
		t.Fatal("Parse error during setup")
	}

	modifiedValue, err := time.ParseDuration("2000ms")

	if err != nil {
		t.Fatal("Parse error during setup")
	}

	deviation, err := time.ParseDuration("1000ms")

	if err != nil {
		t.Fatal("Parse error during setup")
	}

	fixed := Fixed(initialValue)
	variableNormal := VariableNormal(initialValue, deviation, rand.New(rand.NewSource(testSeed)))
	variableUniform := VariableUniform(initialValue, deviation, rand.New(rand.NewSource(testSeed)))

	if fixed.Get().Seconds() != 1 {
		t.Fatal("Fixed delay not initialized correctly")
	}

	if variableNormal.Get().Seconds() != 1 {
		t.Fatal("Normalized variable delay not initialized correctly")
	}

	if variableUniform.Get().Seconds() != 1 {
		t.Fatal("Uniform variable delay not initialized correctly")
	}

	fixed.Set(modifiedValue)

	if fixed.Get().Seconds() != 2 {
		t.Fatal("Fixed delay not set correctly")
	}

	variableNormal.Set(modifiedValue)

	if variableNormal.Get().Seconds() != 2 {
		t.Fatal("Normalized variable delay not set correctly")
	}

	variableUniform.Set(modifiedValue)

	if variableUniform.Get().Seconds() != 2 {
		t.Fatal("Uniform variable delay not initialized correctly")
	}

}

func TestDelayNextWaitTime(t *testing.T) {
	initialValue, err := time.ParseDuration("1000ms")

	if err != nil {
		t.Fatal("Parse error during setup")
	}

	deviation, err := time.ParseDuration("1000ms")

	if err != nil {
		t.Fatal("Parse error during setup")
	}

	fixed := Fixed(initialValue)
	firstRandomNormal := rand.New(rand.NewSource(testSeed)).NormFloat64()
	firstRandom := rand.New(rand.NewSource(testSeed)).Float64()
	variableNormal := VariableNormal(initialValue, deviation, rand.New(rand.NewSource(testSeed)))
	variableUniform := VariableUniform(initialValue, deviation, rand.New(rand.NewSource(testSeed)))
	if fixed.NextWaitTime().Seconds() != 1 {
		t.Fatal("Fixed delay output incorrect wait time")
	}

	if math.Abs(variableNormal.NextWaitTime().Seconds()-(firstRandomNormal*deviation.Seconds()+initialValue.Seconds())) > 0.00001 {
		t.Fatal("Normalized variable delay output incorrect wait time")
	}

	if math.Abs(variableUniform.NextWaitTime().Seconds()-(firstRandom*deviation.Seconds()+initialValue.Seconds())) > 0.00001 {
		t.Fatal("Uniform variable delay output incorrect wait time")
	}

}
