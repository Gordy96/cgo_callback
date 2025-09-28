#ifndef ARRAY_H
#define ARRAY_H

#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define ARRAY_INIT_CAP 4

#define DEFINE_ARRAY(T, Name)                                \
typedef struct {                                             \
    T* data;                                                 \
    size_t len;                                              \
    size_t cap;                                              \
} Name;                                                      \
                                                             \
static inline void Name##_init(Name* a) {                    \
    a->data = NULL;                                          \
    a->len = 0;                                              \
    a->cap = 0;                                              \
}                                                            \
                                                             \
static inline void Name##_free(Name* a) {                    \
    free(a->data);                                           \
    a->data = NULL;                                          \
    a->len = a->cap = 0;                                     \
}                                                            \
                                                             \
static inline void Name##_append(Name* a, T value) {         \
    if (a->len == a->cap) {                                  \
        a->cap = a->cap ? a->cap * 2 : ARRAY_INIT_CAP;       \
        a->data = realloc(a->data, a->cap * sizeof(T));      \
        if (!a->data) { perror("realloc"); exit(1); }        \
    }                                                        \
    a->data[a->len++] = value;                               \
}                                                            \
                                                             \
static inline T Name##_get(Name* a, size_t i) {              \
    return a->data[i];                                       \
}

#endif