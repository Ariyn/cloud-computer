package gate_test

import (
	"testing"

	"github.com/ariyn/cloud-computer/gate"
	"github.com/ariyn/cloud-computer" // For cloud_computer.Element
	"github.com/stretchr/testify/assert"
)

// Helper to initialize FlipFlop and set its initial Q state
func initializeFlipFlopState(ff *gate.FlipFlop, initialQ bool) {
	// Ensure PreviousOutputs is initialized, similar to what Gate.Init would do.
	// A FlipFlop has 2 outputs (Q and Q').
	// The Handler method assigns to PreviousOutputs at the end,
	// but the Changed method reads it first.
	if ff.PreviousOutputs == nil || len(ff.PreviousOutputs) != len(ff.Outputs) {
		ff.PreviousOutputs = make([]bool, len(ff.Outputs))
	}

	if initialQ { // Set Q to 1
		// To set Q=1: S=1, R=0, Enable=true
		ff.PreviousInputs = []bool{true, false, true}
		// Set a PreviousOutputs state that will make 'changed' true if Q flips from 0 to 1
		ff.PreviousOutputs[0] = false // Assume Q was 0
		ff.PreviousOutputs[1] = true  // Assume Q' was 1
		ff.Handler(0, true) // Index 0 (S) becomes true. Q is now 1.
	} else { // Set Q to 0
		// To set Q=0: S=0, R=1, Enable=true
		ff.PreviousInputs = []bool{false, true, true}
		// Set a PreviousOutputs state that will make 'changed' true if Q flips from 1 to 0
		ff.PreviousOutputs[0] = true  // Assume Q was 1
		ff.PreviousOutputs[1] = false // Assume Q' was 0
		ff.Handler(1, true) // Index 1 (R) becomes true. Q is now 0.
	}
	// After setting, inputs for next test should be S=0,R=0,Enable=true to maintain state before actual test input
	// This ensures the previous Handler call that set the state has its PreviousOutputs updated correctly
	// and the 'changed' status for the actual test case is accurate.
	// The Handler call above has already updated ff.PreviousOutputs.
	// Now, set inputs for the actual test condition, usually S=0,R=0 to check for "maintain state".
	ff.PreviousInputs = []bool{false, false, true} // S, R, Enable (for next step, usually maintain)
	// We need to call Handler once more to stabilize PreviousOutputs to the state we just set (Q or !Q)
	// and ensure 'changed' is false for the next "no change" test if state is indeed maintained.
	results, _ := ff.Handler(2, true) // Index 2 (Enable) is 'set' to true. S=0,R=0 should maintain.
	ff.PreviousOutputs = results // Solidify PreviousOutputs to the initialized state.
}

func TestFlipFlop_SR_Latch_Logic(t *testing.T) {
	createTestFF := func() *gate.FlipFlop {
		// Inputs: S, R, Enable. Outputs: Q, Q'
		inputs := make([]cloud_computer.Element, 3)
		for i := 0; i < 3; i++ {
			inputs[i] = cloud_computer.Element{GateName: string(rune('A' + i))}
		}
		outputs := make([]cloud_computer.Element, 2)
		for i := 0; i < 2; i++ {
			outputs[i] = cloud_computer.Element{GateName: string(rune('X' + i))}
		}
		return gate.NewFlipFlopGate("testFF", inputs, outputs)
	}

	t.Run("Enable=false_maintains_Q0", func(t *testing.T) {
		ff := createTestFF()
		initializeFlipFlopState(ff, false) // Initial Q=0

		ff.PreviousInputs = []bool{true, false, false} // S=1, R=0, Enable=false
		results, changed := ff.Handler(2, false) // Enable input at index 2 changes to false
		assert.False(t, results[0], "Q should be 0")
		assert.True(t, results[1], "Q' should be 1")
		assert.False(t, changed, "Enable=false, Q=0, S=1, R=0. Expected no change.")
	})

	t.Run("Enable=false_maintains_Q1", func(t *testing.T) {
		ff := createTestFF()
		initializeFlipFlopState(ff, true) // Initial Q=1

		ff.PreviousInputs = []bool{false, true, false} // S=0, R=1, Enable=false
		results, changed := ff.Handler(2, false) // Enable input at index 2 changes to false
		assert.True(t, results[0], "Q should be 1")
		assert.False(t, results[1], "Q' should be 0")
		assert.False(t, changed, "Enable=false, Q=1, S=0, R=1. Expected no change.")
	})

	t.Run("S=0_R=0_Enable=true_maintains_Q0", func(t *testing.T) {
		ff := createTestFF()
		initializeFlipFlopState(ff, false) // Initial Q=0

		ff.PreviousInputs = []bool{false, false, true} // S=0, R=0, Enable=true
		results, changed := ff.Handler(0, false) // S input (at index 0) is 'set' to false (no actual change in value)
		assert.False(t, results[0], "Q should be 0")
		assert.False(t, changed, "S=0,R=0,Enable=true, Q=0. Expected no change.")
	})

	t.Run("S=0_R=0_Enable=true_maintains_Q1", func(t *testing.T) {
		ff := createTestFF()
		initializeFlipFlopState(ff, true) // Initial Q=1

		ff.PreviousInputs = []bool{false, false, true} // S=0, R=0, Enable=true
		results, changed := ff.Handler(0, false) // S input (at index 0) is 'set' to false (no actual change in value)
		assert.True(t, results[0], "Q should be 1")
		assert.False(t, changed, "S=0,R=0,Enable=true, Q=1. Expected no change.")
	})

	t.Run("S=1_R=0_Enable=true_sets_Q_to_1_from_Q0", func(t *testing.T) {
		ff := createTestFF()
		initializeFlipFlopState(ff, false) // Initial Q=0

		ff.PreviousInputs = []bool{true, false, true} // S=1, R=0, Enable=true
		results, changed := ff.Handler(0, true) // S input (index 0) becomes true
		assert.True(t, results[0], "Q should be 1")
		assert.True(t, changed, "S=1,R=0,Enable=true, Q=0. Expected Q to change to 1.")
	})

    t.Run("S=1_R=0_Enable=true_maintains_Q_when_Q1", func(t *testing.T) {
		ff := createTestFF()
		initializeFlipFlopState(ff, true) // Initial Q=1

		ff.PreviousInputs = []bool{true, false, true} // S=1, R=0, Enable=true
		results, changed := ff.Handler(0, true) // S input (index 0) becomes true
		assert.True(t, results[0], "Q should be 1")
		// 'changed' is true if ff.PreviousOutputs (from init) is different from current results.
		// initializeFlipFlopState sets Q to 1 (S=1,R=0,E=1 -> Q=1,Q'=0) then does S=0,R=0,E=1 -> Q=1,Q'=0.
		// So PreviousOutputs is {true, false}.
		// Current operation is S=1,R=0,E=1 -> Q=1,Q'=0. Results are {true, false}.
		// So, changed should be false.
		assert.False(t, changed, "S=1,R=0,Enable=true, Q=1. Expected no change.")
	})

	t.Run("S=0_R=1_Enable=true_sets_Q_to_0_from_Q1", func(t *testing.T) {
		ff := createTestFF()
		initializeFlipFlopState(ff, true) // Initial Q=1

		ff.PreviousInputs = []bool{false, true, true} // S=0, R=1, Enable=true
		results, changed := ff.Handler(1, true) // R input (index 1) becomes true
		assert.False(t, results[0], "Q should be 0")
		assert.True(t, changed, "S=0,R=1,Enable=true, Q=1. Expected Q to change to 0.")
	})

    t.Run("S=0_R=1_Enable=true_maintains_Q_when_Q0", func(t *testing.T) {
		ff := createTestFF()
		initializeFlipFlopState(ff, false) // Initial Q=0

		ff.PreviousInputs = []bool{false, true, true} // S=0, R=1, Enable=true
		results, changed := ff.Handler(1, true) // R input (index 1) becomes true
		assert.False(t, results[0], "Q should be 0")
		assert.False(t, changed, "S=0,R=1,Enable=true, Q=0. Expected no change.")
	})

	t.Run("S=1_R=1_Enable=true_sets_Q_to_1_from_Q0_documented", func(t *testing.T) {
		ff := createTestFF()
		initializeFlipFlopState(ff, false) // Initial Q=0

		ff.PreviousInputs = []bool{true, true, true} // S=1, R=1, Enable=true
		results, changed := ff.Handler(0, true) // S (index 0) becomes true
		assert.True(t, results[0], "Q should be 1")
		assert.True(t, changed, "S=1,R=1,Enable=true, Q=0. Expected Q to change to 1 (documented behavior).")
	})

    t.Run("S=1_R=1_Enable=true_maintains_Q_when_Q1_documented", func(t *testing.T) {
		ff := createTestFF()
		initializeFlipFlopState(ff, true) // Initial Q=1

		ff.PreviousInputs = []bool{true, true, true} // S=1, R=1, Enable=true
		results, changed := ff.Handler(0, true)
		assert.True(t, results[0], "Q should be 1")
		assert.False(t, changed, "S=1,R=1,Enable=true, Q=1. Expected no change (documented behavior).")
	})

	t.Run("Enable_transitions_false_to_true_S1_R0_from_Q0", func(t *testing.T) {
		ff := createTestFF()
		initializeFlipFlopState(ff, false) // Initial Q=0

		// Set S=1, R=0, but Enable=false initially
		ff.PreviousInputs = []bool{true, false, false}
		results, changed := ff.Handler(2, false) // Enable remains false, or is set to false
		assert.False(t, results[0], "Q should initially be 0 (Enable false)")
		assert.False(t, changed, "Should not change with Enable false")

		// Now, transition Enable to true
		ff.PreviousInputs[2] = true // Enable becomes true
		results, changed = ff.Handler(2, true) // Enable (index 2) changes to true
		assert.True(t, results[0], "Q should be 1 after Enable becomes true with S=1,R=0")
		assert.True(t, changed, "Should change when Enable becomes true and inputs dictate change")
	})

	t.Run("Enable_transitions_false_to_true_S0_R1_from_Q1", func(t *testing.T) {
		ff := createTestFF()
		initializeFlipFlopState(ff, true) // Initial Q=1

		// Set S=0, R=1, but Enable=false initially
		ff.PreviousInputs = []bool{false, true, false}
		results, changed := ff.Handler(2, false) // Enable remains false
		assert.True(t, results[0], "Q should initially be 1 (Enable false)")
		assert.False(t, changed, "Should not change with Enable false")

		// Now, transition Enable to true
		ff.PreviousInputs[2] = true // Enable becomes true
		results, changed = ff.Handler(2, true) // Enable (index 2) changes to true
		assert.False(t, results[0], "Q should be 0 after Enable becomes true with S=0,R=1")
		assert.True(t, changed, "Should change when Enable becomes true and inputs dictate change")
	})
}
