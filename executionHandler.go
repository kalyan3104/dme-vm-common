package vmcommon

// VMExecutionHandler interface for any Kalyan VM endpoint
type VMExecutionHandler interface {
	// Computes how a smart contract creation should be performed
	RunSmartContractCreate(input *ContractCreateInput) (*VMOutput, error)

	// Computes the result of a smart contract call and how the system must change after the execution
	RunSmartContractCall(input *ContractCallInput) (*VMOutput, error)
}
