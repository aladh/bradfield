#include <stdio.h>
#include <string.h>
#include <sys/stat.h>
#include <sys/dir.h>
#include <fcntl.h>
#include <linux/limits.h>

#define PROGRAM_NAME "ls"

off_t file_size(char *);
void traverse_dir(char *, void (*fcn)(char *, char *));
void print_entry(char *, char *);