package channelconfig

import "github.com/hyperledger/fabric-protos-go/common"

type attestationCheckingParameters struct {
	enableChecking           bool
	requiredNumberOfMessages uint32
	frequency                uint32
}

func (p *attestationCheckingParameters) EnableChecking() bool {
	return p.enableChecking
}

func (p *attestationCheckingParameters) RequiredNumberOfMessages() uint32 {
	return p.requiredNumberOfMessages
}

func (p *attestationCheckingParameters) Frequency() uint32 {
	return p.frequency
}

func newAttestationCheckingParameters(parameters *common.AttestationCheckingParameters) AttestationCheckingParameters {
	return &attestationCheckingParameters{
		enableChecking:           parameters.EnableChecking,
		requiredNumberOfMessages: parameters.RequiredNumberOfMessages,
		frequency:                parameters.Frequency,
	}
}
