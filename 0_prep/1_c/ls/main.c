#include "ls.h"

int main(int argc, char *argv[]) {
    if (argc == 1) { // default: current directory
        dirwalk(".", print_entry);
    } else {
        while (--argc > 0) {
            dirwalk(*++argv, print_entry);
        }
    }

    return 0;
}

void print_entry(char *dirname, char *filename) {
    char *path;
    sprintf(path, "%s/%s", dirname, filename);

    printf("%8ld %s\n", fsize(path), filename);
}

off_t fsize(char *name) {
    struct stat stbuf;

    if (stat(name, &stbuf) == -1) {
        fprintf(stderr, "fsize: can't access %s\n", name);
        return -1;
    }

    return stbuf.st_size;
}

void dirwalk(char *dir, void (*fcn)(char *, char *)) {
    struct dirent *dp;
    DIR *dfd;

    if ((dfd = opendir(dir)) == NULL) {
        fprintf(stderr, "dirwalk: can't open %s\n", dir);
        return;
    }

    while ((dp = readdir(dfd)) != NULL) {
        if (strcmp(dp->d_name, ".") == 0 || strcmp(dp->d_name, "..") == 0) {
            continue; // skip self and parent
        }

        if (strlen(dir) + strlen(dp->d_name) + 2 > PATH_MAX) {
            fprintf(stderr, "dirwalk: name %s/%s too long\n", dir, dp->d_name);
        } else {
            (*fcn)(dir, dp->d_name);
        }
    }

    closedir(dfd);
}