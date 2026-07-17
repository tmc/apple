#include <infiniband/verbs.h>
#include <stddef.h>
#include <stdio.h>

#define SIZE(type) printf(#type ".size=%zu\n", sizeof(type))
#define OFFSET(type, field) printf(#type "." #field "=%zu\n", offsetof(type, field))

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
	OFFSET(struct ibv_port_attr, gid_tbl_len);
	OFFSET(struct ibv_port_attr, lid);
	OFFSET(struct ibv_port_attr, link_layer);
	OFFSET(struct ibv_port_attr, port_cap_flags2);
	return 0;
}
