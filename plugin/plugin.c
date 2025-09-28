#include <stdio.h>
#include "plugin.h"

static read_cb_t go_read = NULL;
static write_cb_t go_write = NULL;

void init(const char* port, read_cb_t r, write_cb_t w) {
    go_read = r;
    go_write = w;
}

void run(const char* port) {
    if (!go_read || !go_write) {
        printf("handlers not registered\n");
        return;
    }

    char buf[256];
    int n = go_read(port, buf);
    if (n > 0) {
        printf("plugin read (len=%d): %.*s\n", n, n, buf);
        go_write(port, buf, n); // echo back
    } else {
        printf("plugin read returned %d\n", n);
    }
}