#include <stdint.h>
#include <stddef.h>

typedef int (*adder_fn)(int);

int call_adder_with_one(adder_fn fn);