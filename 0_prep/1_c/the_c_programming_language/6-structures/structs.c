#include <stdio.h>

struct point {
    int x;
    int y;
};

struct rect {
    struct point pt1;
    struct point pt2;
};

int main() {
    struct point maxpt = {320, 200};
    printf("%d,%d\n", maxpt.x, maxpt.y);

    struct rect screen;
    struct point middle;
    struct point makepoint(int, int);

    screen.pt1 = makepoint(0, 0);
    screen.pt2 = makepoint(maxpt.x, maxpt.y);
    middle = makepoint((screen.pt1.x + screen.pt2.x) / 2, (screen.pt1.y + screen.pt2.y) / 2);

    struct point origin, *pp;

    origin = (struct point) {.x = 1, .y = 2};

    pp = &origin;
    printf("origin is (%d,%d)\n", (*pp).x, (*pp).y);
    printf("origin is (%d,%d)\n", pp->x, pp->y);
}

struct point makepoint(int x, int y) {
    struct point temp;

    temp.x = x;
    temp.y = y;

    return temp;
}

struct point addpoint(struct point p1, struct point p2) {
    p1.x += p2.x;
    p1.y += p2.y;

    return p1;
}