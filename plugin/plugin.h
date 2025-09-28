#ifndef PLUGIN_H
#define PLUGIN_H

typedef int (*read_cb_t)(char* port, char* buf);
typedef int (*write_cb_t)(char* port, char* buf, int size);
typedef void (*init_func_t)(const char* port, read_cb_t r, write_cb_t w);
typedef void (*run_func_t)(const char* port);

#endif