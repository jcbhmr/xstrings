#include "callbacktest.h"
#include <stdint.h>
#include <stddef.h>

int call_adder_with_one(adder_fn fn) {
    return fn(1);
}