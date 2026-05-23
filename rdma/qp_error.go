package rdma

import (
	"errors"
	"fmt"
	"strings"
)

// ModifyQPError describes a failed ibv_modify_qp transition.
type ModifyQPError struct {
	Transition string
	Device     string
	QPN        uint32
	Mask       int
	Attr       IbvQPAttr
	Return     int
	Cause      error
}

// NewModifyQPError formats a failed ibv_modify_qp operation.
func NewModifyQPError(qp RDMAQP, attr *IbvQPAttr, attrMask int, rc int, err error) error {
	if err == nil && rc == 0 {
		return nil
	}
	context := rdmaContextFromQP(qp)
	cause := err
	if cause == nil {
		cause = rdmaProviderStatusError("ibv_modify_qp", rc, context, true)
	}
	e := &ModifyQPError{
		Transition: inferQPTransition(attr),
		Device:     rdmaContextDevice(context),
		QPN:        Ibv_qp_num(qp),
		Mask:       attrMask,
		Return:     rc,
		Cause:      cause,
	}
	if attr != nil {
		e.Attr = *attr
	}
	return e
}

func (e *ModifyQPError) Error() string {
	if e == nil {
		return ""
	}
	var b strings.Builder
	b.WriteString("ibv_modify_qp")
	if e.Transition != "" {
		b.WriteString(" ")
		b.WriteString(e.Transition)
	}
	if e.Device != "" {
		b.WriteString(" on ")
		b.WriteString(e.Device)
	}
	if e.QPN != 0 {
		fmt.Fprintf(&b, " qpn=%d", e.QPN)
	}
	b.WriteString(" failed")
	if e.Return != 0 {
		b.WriteString(": ")
		b.WriteString(ErrnoText(e.Return))
	} else if e.Cause != nil {
		b.WriteString(": ")
		b.WriteString(e.Cause.Error())
	}
	fmt.Fprintf(&b, ", mask=%#x", e.Mask)
	if names := QPAttrMaskNames(e.Mask); len(names) != 0 {
		b.WriteString(" [")
		b.WriteString(strings.Join(names, "|"))
		b.WriteString("]")
	}
	b.WriteString("; attrs: ")
	b.WriteString(formatQPAttr(e.Attr))
	if e.Return == 60 || errors.Is(e.Cause, ErrProviderTimeout) {
		b.WriteString("; hint: provider timed out applying ")
		if e.Transition != "" {
			b.WriteString(e.Transition)
		} else {
			b.WriteString("QP transition")
		}
		b.WriteString("; check peer LID/GID and Thunderbolt path")
	}
	return b.String()
}

func (e *ModifyQPError) Unwrap() error {
	if e == nil {
		return nil
	}
	return e.Cause
}

func QPAttrMaskNames(mask int) []string {
	all := []struct {
		bit  int
		name string
	}{
		{IBV_QP_STATE, "IBV_QP_STATE"},
		{IBV_QP_ACCESS_FLAGS, "IBV_QP_ACCESS_FLAGS"},
		{IBV_QP_PKEY_INDEX, "IBV_QP_PKEY_INDEX"},
		{IBV_QP_PORT, "IBV_QP_PORT"},
		{IBV_QP_AV, "IBV_QP_AV"},
		{IBV_QP_PATH_MTU, "IBV_QP_PATH_MTU"},
		{IBV_QP_RQ_PSN, "IBV_QP_RQ_PSN"},
		{IBV_QP_SQ_PSN, "IBV_QP_SQ_PSN"},
		{IBV_QP_DEST_QPN, "IBV_QP_DEST_QPN"},
	}
	names := make([]string, 0, len(all))
	remaining := mask
	for _, entry := range all {
		if mask&entry.bit != 0 {
			names = append(names, entry.name)
			remaining &^= entry.bit
		}
	}
	if remaining != 0 {
		names = append(names, fmt.Sprintf("unknown(%#x)", remaining))
	}
	return names
}

func QPStateName(state int32) string {
	switch state {
	case IBV_QPS_INIT:
		return "INIT"
	case IBV_QPS_RTR:
		return "RTR"
	case IBV_QPS_RTS:
		return "RTS"
	default:
		return fmt.Sprintf("state(%d)", state)
	}
}

func inferQPTransition(attr *IbvQPAttr) string {
	if attr == nil {
		return ""
	}
	switch attr.QPState {
	case IBV_QPS_INIT:
		return "RESET->INIT"
	case IBV_QPS_RTR:
		return "INIT->RTR"
	case IBV_QPS_RTS:
		return "RTR->RTS"
	default:
		return "->" + QPStateName(attr.QPState)
	}
}

func formatQPAttr(attr IbvQPAttr) string {
	parts := []string{
		"state=" + QPStateName(attr.QPState),
		fmt.Sprintf("port=%d", attr.AHAttr.PortNum),
		fmt.Sprintf("dlid=%d", attr.AHAttr.DLID),
		fmt.Sprintf("dest_qpn=%d", attr.DestQPNum),
		fmt.Sprintf("rq_psn=%d", attr.RQPSN),
		fmt.Sprintf("sq_psn=%d", attr.SQPSN),
		fmt.Sprintf("gid_index=%d", attr.AHAttr.GRH.SGIDIndex),
		fmt.Sprintf("gid=%x", attr.AHAttr.GRH.DGID[:]),
		fmt.Sprintf("path_mtu=%s", mtuText(attr.PathMTU)),
		fmt.Sprintf("timeout=%d", attr.Timeout),
		fmt.Sprintf("retry_cnt=%d", attr.RetryCnt),
		fmt.Sprintf("rnr_retry=%d", attr.RNRetry),
		fmt.Sprintf("pkey_index=%d", attr.PKeyIndex),
		fmt.Sprintf("access_flags=%#x", attr.QPAccessFlags),
	}
	return strings.Join(parts, " ")
}

func mtuText(mtu int32) string {
	switch mtu {
	case 1:
		return "256"
	case 2:
		return "512"
	case IBV_MTU_1024:
		return "1024"
	case 4:
		return "2048"
	case 5:
		return "4096"
	default:
		return fmt.Sprintf("%d", mtu)
	}
}
