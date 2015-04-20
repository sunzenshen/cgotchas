#ifndef MPC_INTERFACE_H
#define MPC_INTERFACE_H

#include "mpc.h"

inline void mpc_cleanup_if
(
  mpc_parser_t* p1,
  mpc_parser_t* p2,
  mpc_parser_t* p3,
  mpc_parser_t* p4
)
{
  mpc_cleanup(4, p1, p2, p3, p4);
}

#endif

