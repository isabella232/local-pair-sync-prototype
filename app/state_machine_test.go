package app

import "testing"

func TestState_ProgressCurrentStep(t *testing.T) {
	cs := []struct{
		Input State
		Result StateStep
		Error error
	}{
		{
			State{Step: GetMode, Mode: Sender},
			FindingDevices,
			nil,
		},
		{
			State{Step: GetMode, Mode: Unset},
			GetMode,
			ErrorNoModeSet,
		},
		{
			State{Step: ShowFoundDevices, Mode: Sender},
			AttemptPairing,
			nil,
		},
		{
			State{Step: AttemptPairing, Mode: Receiver},
			SendingPassword,
			nil,
		},
		{
			State{Step: SentPayload, Mode: Sender},
			Complete,
			nil,
		},
		{
			State{Step: Complete, Mode: Sender},
			Complete,
			ErrorCompleteStepMet,
		},
	}

	for _, c := range cs {
		err := c.Input.ProgressCurrentStep()
		if err != c.Error{
			t.Error(err)
		}

		if c.Input.Step != c.Result {
			t.Logf("unexpected result, expected '%d', received '%d'", c.Result, c.Input.Step)
			t.Fail()
		}
	}
}