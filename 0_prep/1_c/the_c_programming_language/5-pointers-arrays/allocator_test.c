#include <stdio.h>

char *alloc(int n);
void afree(char *p);

int main() {
    char *c1p = alloc(1);
    *c1p = 'A';

    char *c2p = alloc(1);
    *c2p = 'B';

    printf("%c\n", *c1p);
    printf("%c\n", *c2p);

    afree(c2p);

    c2p = alloc(1);
    *c2p = 'C';

    printf("%c\n", *c1p);
    printf("%c\n", *c2p);
}