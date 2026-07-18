#include <infiniband/verbs.h>
#include <stddef.h>
#include <stdio.h>

#define SIZE(type) printf(#type ".size=%zu\n", sizeof(type))
#define OFFSET(type, field) printf(#type "." #field "=%zu\n", offsetof(type, field))
#define VALUE(value) printf(#value "=%d\n", value)

int
main(void)
{
	SIZE(union ibv_gid);
	SIZE(struct ibv_qp_cap);
	SIZE(struct ibv_qp_init_attr);
	SIZE(struct ibv_global_route);
	SIZE(struct ibv_ah_attr);
	SIZE(struct ibv_qp_attr);
	SIZE(struct ibv_sge);
	SIZE(struct ibv_send_wr);
	SIZE(struct ibv_recv_wr);
	SIZE(struct ibv_wc);
	SIZE(struct ibv_port_attr);

	OFFSET(struct ibv_qp_init_attr, send_cq);
	OFFSET(struct ibv_qp_init_attr, cap);
	OFFSET(struct ibv_qp_init_attr, qp_type);
	OFFSET(struct ibv_ah_attr, dlid);
	OFFSET(struct ibv_qp_attr, cap);
	OFFSET(struct ibv_qp_attr, ah_attr);
	OFFSET(struct ibv_qp_attr, alt_ah_attr);
	OFFSET(struct ibv_qp_attr, pkey_index);
	OFFSET(struct ibv_qp_attr, port_num);
	OFFSET(struct ibv_qp_attr, rate_limit);
	OFFSET(struct ibv_send_wr, wr);
	OFFSET(struct ibv_context, ops);
	OFFSET(struct ibv_context_ops, poll_cq);
	OFFSET(struct ibv_context_ops, post_send);
	OFFSET(struct ibv_context_ops, post_recv);
	OFFSET(struct ibv_qp, qp_num);
	OFFSET(struct ibv_mr, lkey);
	OFFSET(struct ibv_mr, rkey);
	OFFSET(struct ibv_comp_channel, fd);
	OFFSET(struct ibv_cq, channel);
	OFFSET(struct ibv_port_attr, gid_tbl_len);
	OFFSET(struct ibv_port_attr, lid);
	OFFSET(struct ibv_port_attr, link_layer);
	OFFSET(struct ibv_port_attr, port_cap_flags2);

	VALUE(IBV_QPT_RC);
	VALUE(IBV_QPT_UC);
	VALUE(IBV_WR_RDMA_WRITE);
	VALUE(IBV_WR_SEND);
	VALUE(IBV_WR_RDMA_READ);
	VALUE(IBV_ACCESS_LOCAL_WRITE);
	VALUE(IBV_ACCESS_REMOTE_WRITE);
	VALUE(IBV_ACCESS_REMOTE_READ);
	VALUE(IBV_WC_SUCCESS);
	VALUE(IBV_WC_LOC_PROT_ERR);
	VALUE(IBV_WC_LOC_ACCESS_ERR);
	VALUE(IBV_WC_REM_ACCESS_ERR);
	VALUE(IBV_QP_STATE);
	VALUE(IBV_QP_ACCESS_FLAGS);
	VALUE(IBV_QP_PKEY_INDEX);
	VALUE(IBV_QP_PORT);
	VALUE(IBV_QP_AV);
	VALUE(IBV_QP_PATH_MTU);
	VALUE(IBV_QP_RQ_PSN);
	VALUE(IBV_QP_SQ_PSN);
	VALUE(IBV_QP_DEST_QPN);
	VALUE(IBV_SEND_SIGNALED);
	VALUE(ENOTSUP);
	return 0;
}
