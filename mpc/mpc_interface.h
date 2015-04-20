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

inline mpc_err_t* mpca_lang_if
(
  int flags,
  const char *language,
  mpc_parser_t* parser1,
  mpc_parser_t* parser2,
  mpc_parser_t* parser3,
  mpc_parser_t* parser4
)
{
  return mpca_lang(flags, language, parser1, parser2, parser3, parser4);
}

#endif

