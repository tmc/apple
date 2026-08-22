package rdma

// Work-completion statuses used to classify provider errors. These values are
// part of the verbs ABI; they do not imply that an unsupported verb is usable.
const (
	IBV_WC_LOC_PROT_ERR   = 4
	IBV_WC_LOC_ACCESS_ERR = 8
	IBV_WC_REM_ACCESS_ERR = 10
)

// CompletionStatusClass describes the result class of a work completion.
type CompletionStatusClass string

const (
	CompletionSuccess    CompletionStatusClass = "success"
	CompletionProtection CompletionStatusClass = "protection error"
	CompletionFailure    CompletionStatusClass = "failure"
)

// ClassifyCompletionStatus classifies status without collapsing protection
// errors into a generic provider failure.
func ClassifyCompletionStatus(status int32) CompletionStatusClass {
	switch status {
	case IBV_WC_SUCCESS:
		return CompletionSuccess
	case IBV_WC_LOC_PROT_ERR, IBV_WC_LOC_ACCESS_ERR, IBV_WC_REM_ACCESS_ERR:
		return CompletionProtection
	default:
		return CompletionFailure
	}
}
