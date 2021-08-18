#include "ls.h"

int main(int argc, char *argv[]) {
    if (argc == 1) { // default: current directory
        traverse_dir(".", print_entry);
    } else {
        while (--argc > 0) {
            traverse_dir(*++argv, print_entry);
        }
    }

    return 0;
}

void print_entry(char *dir_path, char *filename) {
    char file_path[PATH_MAX];
    sprintf(file_path, "%s/%s", dir_path, filename);

    printf("%8ld %s\n", file_size(file_path), filename);
}

off_t file_size(char *file_path) {
    struct stat stat_buf;

    if (stat(file_path, &stat_buf) == -1) {
        fprintf(stderr, "%s: can't access %s\n", PROGRAM_NAME, file_path);
        return -1;
    }

    return stat_buf.st_size;
}

void traverse_dir(char *path, void (*func)(char *, char *)) {
    struct dirent *dir_entry;
    DIR *dir_fd;

    if ((dir_fd = opendir(path)) == NULL) {
        fprintf(stderr, "%s: can't open %s\n", PROGRAM_NAME, path);
        return;
    }

    while ((dir_entry = readdir(dir_fd)) != NULL) {
        char *filename = dir_entry->d_name;

        if (strcmp(filename, ".") == 0 || strcmp(filename, "..") == 0) {
            continue; // skip self and parent
        }

        (*func)(path, filename);
    }

    closedir(dir_fd);
}