#include "call.h"
#include <stdint.h>
#include <stddef.h>
#include <assert.h>
#include <stdlib.h>

uintptr_t call(void *fn, uintptr_t *args, size_t args_len) {
    assert(fn != NULL);
    assert(args != NULL || args_len == 0);
    assert(args_len <= 15);
    switch (args_len) {
    case 0:
        return ((uintptr_t (*)(void))fn)();
    case 1:
        return ((uintptr_t (*)(uintptr_t))fn)(args[0]);
    case 2:
        return ((uintptr_t (*)(uintptr_t, uintptr_t))fn)(args[0], args[1]);
    case 3:
        return ((uintptr_t (*)(uintptr_t, uintptr_t, uintptr_t))fn)(args[0], args[1], args[2]);
    case 4:
        return ((uintptr_t (*)(uintptr_t, uintptr_t, uintptr_t, uintptr_t))fn)(args[0], args[1], args[2], args[3]);
    case 5:
        return ((uintptr_t (*)(uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t))fn)(args[0], args[1], args[2], args[3], args[4]);
    case 6:
        return ((uintptr_t (*)(uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t))fn)(args[0], args[1], args[2], args[3], args[4], args[5]);
    case 7:
        return ((uintptr_t (*)(uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t))fn)(args[0], args[1], args[2], args[3], args[4], args[5], args[6]);
    case 8:
        return ((uintptr_t (*)(uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t))fn)(args[0], args[1], args[2], args[3], args[4], args[5], args[6], args[7]);
    case 9:
        return ((uintptr_t (*)(uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t))fn)(args[0], args[1], args[2], args[3], args[4], args[5], args[6], args[7], args[8]);
    case 10:
        return ((uintptr_t (*)(uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t))fn)(args[0], args[1], args[2], args[3], args[4], args[5], args[6], args[7], args[8], args[9]);
    case 11:
        return ((uintptr_t (*)(uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t))fn)(args[0], args[1], args[2], args[3], args[4], args[5], args[6], args[7], args[8], args[9], args[10]);
    case 12:
        return ((uintptr_t (*)(uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t))fn)(args[0], args[1], args[2], args[3], args[4], args[5], args[6], args[7], args[8], args[9], args[10], args[11]);
    case 13:
        return ((uintptr_t (*)(uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t))fn)(args[0], args[1], args[2], args[3], args[4], args[5], args[6], args[7], args[8], args[9], args[10], args[11], args[12]);
    case 14:
        return ((uintptr_t (*)(uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t))fn)(args[0], args[1], args[2], args[3], args[4], args[5], args[6], args[7], args[8], args[9], args[10], args[11], args[12], args[13]);
    case 15:
        return ((uintptr_t (*)(uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t))fn)(args[0], args[1], args[2], args[3], args[4], args[5], args[6], args[7], args[8], args[9], args[10], args[11], args[12], args[13], args[14]);
    default:
        abort();
    }
}
