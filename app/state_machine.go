package app

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"
)

/*
- Mode sender / receiver
- Both
  - Broadcast and find other device
    - Store data about any other devices
  - Show name of other device
  - Attempt pairing
- Sender
  - Show password
  - Await password
  - Correct password
  - Sending payload
  - Sent payload
- Receiver
  - Send password
  - Receiving payload
*/

type StateMode int
type StateStep int

const (
	Unset StateMode = iota
	Sender
	Receiver
)

func (s StateMode) String() string {
	switch s {
	case Unset:
		return "Unset"
	case Sender:
		return "Sender"
	case Receiver:
		return "Receiver"
	default:
		return "unknown"
	}
}

const (
	GetMode StateStep = iota + 1
	FindingDevices
	ShowFoundDevices
	AttemptPairing
	ShowPassword
	CorrectPassword
	SendingPayload
	SentPayload
	SendingPassword
	ReceivingPayload
	ReceivedPayload
	Complete
)

var (
	StepToModeMap = map[StateStep][]StateMode{
		GetMode:          {Unset},
		FindingDevices:   {Sender, Receiver},
		ShowFoundDevices: {Sender, Receiver},
		AttemptPairing:   {Sender, Receiver},
		ShowPassword:     {Sender},
		CorrectPassword:  {Sender},
		SendingPayload:   {Sender},
		SentPayload:      {Sender},
		SendingPassword:  {Receiver},
		ReceivingPayload: {Receiver},
		ReceivedPayload:  {Receiver},
		Complete:         {Sender, Receiver},
	}

	ErrorNoModeSet       = errors.New("no mode set")
	ErrorCompleteStepMet = errors.New("can not progress the complete step, call Reset()")
)

type State struct {
	Mode           StateMode
	Step           StateStep
	Devices        []Device
	SelectedDevice int
}

func NewState() *State {
	s := new(State)
	s.Reset()
	return s
}

func (s *State) Reset() {
	s.Mode = Unset
	s.Step = GetMode
	s.Devices = nil
	s.SelectedDevice = -1
}

func (s *State) ValidateCurrentStep() error {
	return s.ValidateStep(s.Step)
}

func (s *State) ValidateStep(step StateStep) error {
	found := false
	stepModes := StepToModeMap[step]
	for _, validMode := range stepModes {
		if validMode == s.Mode {
			found = true
		}
	}

	if found {
		return nil
	}

	return fmt.Errorf("state step and mode mismatch: %d is not a valid step of %d", s.Step, s.Mode)
}

func (s *State) ProgressCurrentStep() error {
	nextStep, err := s.ProgressStep(s.Step)
	if err != nil {
		return err
	}

	s.Step = nextStep
	return nil
}

func (s *State) ProgressStep(step StateStep) (StateStep, error) {
	if s.Mode == Unset {
		return 0, ErrorNoModeSet
	}
	nextPotentialStep := step + 1

	if nextPotentialStep > Complete {
		return 0, ErrorCompleteStepMet
	}

	err := s.ValidateStep(nextPotentialStep)
	if err != nil {
		return s.ProgressStep(nextPotentialStep)
	}

	return nextPotentialStep, nil
}

func (s *State) stepToFunc(step StateStep) func() error {
	var stepToFuncMap = map[StateStep]func() error{
		GetMode:          s.getMode,
		FindingDevices:   s.startSearch,
		ShowFoundDevices: s.searchingForDevices,
		AttemptPairing:   s.foundDevices,
		ShowPassword:     s.showPassword,
		CorrectPassword:  s.correctPassword,
		SendingPayload:   s.sendingPayload,
		SentPayload:      s.payloadSent,
		SendingPassword:  s.sendingPassword,
		ReceivingPayload: s.receivingPayload,
		ReceivedPayload:  s.payloadReceived,
		Complete:         s.completed,
	}

	return stepToFuncMap[step]
}

func (s *State) getMode() error {
	s.printTitle("Set Pairing Mode")
	for {
		fmt.Println("Select pairing mode:\n- [1] Sender:\n- [2] Receiver:")
		var mode string
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			mode = scanner.Text()
		}

		// Validate mode data
		im, err := strconv.ParseInt(mode, 10, 64)
		if err != nil && err.(*strconv.NumError).Err == strconv.ErrSyntax {
			fmt.Printf("ERROR - Unrecognised option '%s' : valid options '1' or '2'\n\n", mode)
			continue
		}
		if err != nil {
			return err
		}

		switch im {
		case 1, 2:
			// Set mode in state
			s.Mode = StateMode(im)
			fmt.Printf("Pairing mode set to '%s'\n", s.Mode)
		default:
			fmt.Printf("ERROR - Unrecognised option '%d' : valid options '1' or '2'\n\n", im)
			continue
		}

		return nil
	}
}

func (s *State) startSearch() error {
	s.printTitle("Confirm Device Search")
	for {
		var startSearch string
		fmt.Println("Are you ready to search for devices to pair with? [Y/N]")
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			startSearch = scanner.Text()
		}

		switch startSearch {
		case "Y", "y":
			fmt.Println("Commencing search for other devices...")
		case "N", "n":
			fmt.Println("Ok ... Just let me know when you are.")
			time.Sleep(time.Millisecond * 500)
			fmt.Println("Thanks. :)")
			time.Sleep(time.Millisecond * 500)
			continue
		default:
			fmt.Printf("ERROR - Unrecognised option '%s' : valid options 'Y' or 'N'\n", startSearch)
			continue
		}

		return nil
	}
}

func (s *State) searchingForDevices() error {
	s.printTitle("Device Search Process")
	for {
		fmt.Println("Searching for devices...")

		time.Sleep(time.Second * 2)
		// The code for finding devices

		s.Devices = ds
		fmt.Printf("Found '%d' device(s)\n", len(s.Devices))
		return nil
	}
}

func (s *State) foundDevices() error {
	s.printTitle("Select a Device To Pair With")
	for {
		fmt.Println("The following device(s) were found:")
		for i, d := range s.Devices {
			fmt.Printf("-[%d] Name: '%s' - IP: '%s'\n", i+1, d.Name, d.IP)
		}

		var device string
		fmt.Println("Choose which device to pair with:")
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			device = scanner.Text()
		}

		// Validate option
		id, err := strconv.ParseInt(device, 10, 64)
		if err != nil && err.(*strconv.NumError).Err == strconv.ErrSyntax {
			fmt.Printf("ERROR - Device ID not valid '%s', please give an integer value\n\n", device)
			continue
		}
		if err != nil {
			return err
		}

		if id < 0 || int(id) > len(s.Devices) {
			fmt.Printf("ERROR - Device ID not valid '%d', please give a device ID between '%d' and '%d'\n\n", id, 1, len(s.Devices))
			continue
		}

		// Set device to pair with
		s.SelectedDevice = int(id)
		fmt.Printf("SELECTED - [%d] Name: '%s' - IP: '%s'\n", id, s.Devices[id-1].Name, s.Devices[id-1].IP)
		return nil
	}
}

func (s *State) showPassword() error {
	password := "password"
	s.printTitle("Your Connection Password")
	for {
		fmt.Println("Enter the below password into your receiving device:")
		fmt.Println(password)

		fmt.Println("Press any key to continue")
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			_ = scanner.Text()
		}

		fmt.Println("Awaiting receiving device secure connection with given password")

		return nil
	}
}

func (s *State) correctPassword() error {
	s.printTitle("Secure Connection Established")
	for {
		fmt.Println("Ready to commence key transfer:")

		fmt.Println("Press any key to continue")
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			_ = scanner.Text()
		}

		return nil
	}
}

func (s *State) sendingPayload() error {
	s.printTitle("Sending Payload")
	for {
		fmt.Println("Payload transferring to receiving device:")
		time.Sleep(time.Millisecond * 1500)

		return nil
	}
}

func (s *State) payloadSent() error {
	s.printTitle("Payload Sent")
	for {
		fmt.Println("Payload was successfully sent:")
		return nil
	}
}

func (s *State) sendingPassword() error {
	s.printTitle("Enter Connection Password")
	for {
		fmt.Println("If your other device is in sender mode and is ready to pair, it will have given you a password:")

		var password string
		fmt.Println("Please enter the connection password below:")
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			password = scanner.Text()
		}

		if password != "password" {
			fmt.Println("The password was invalid")
			continue
		}

		return nil
	}
}

func (s *State) receivingPayload() error {
	s.printTitle("Receiving Payload")
	for {
		fmt.Println("Payload transferring from sending device:")
		time.Sleep(time.Millisecond * 1500)

		return nil
	}
}

func (s *State) payloadReceived() error {
	s.printTitle("Payload Received")
	for {
		fmt.Println("Payload was successfully received:")
		return nil
	}
}

func (s *State) completed() error {
	s.printTitle("Initial Pairing Completed")
	for {
		fmt.Println("The initial pairing is completed")
		return nil
	}
}

func (s *State) Perform() error {
	action := s.stepToFunc(s.Step)
	err := action()
	if err != nil {
		return err
	}

	return nil
}

func (s *State) printTitle(title string) {
	fmt.Printf("\n=== %s - %s ===\n", s.Mode, title)
}
