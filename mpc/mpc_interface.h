#include "mpc.h"

inline void mpc_cleanup_if(mpc_parser_t* parser) {
  mpc_cleanup(1, parser);
}

