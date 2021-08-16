#include <string.h>
#include <stdlib.h>
#include <stdio.h>
#define HASHSIZE 101

struct nlist { // table entry
    struct nlist *next;
    char *name; // defined name
    char *defn; // replacement text
};

static struct nlist *hashtab[HASHSIZE]; // pointer table

unsigned hash(char *s) {
    unsigned hashval;

    for (hashval = 0; *s != '\0'; s++) {
        hashval = *s + 31 * hashval;
    }

    return hashval % HASHSIZE;
}

struct nlist *lookup(char *s) {
    struct nlist *np;

    for (np = hashtab[hash(s)]; np != NULL; np = np->next) {
        if (strcmp(s, np->name) == 0) {
            return np; // found
        }
    };

    return NULL;
}

struct nlist *install(char *name, char *defn) {
    struct nlist *np;
    unsigned hashval;

    if ((np = lookup(name)) == NULL) { // not found
        np = (struct nlist *) malloc(sizeof(*np));

        if (np == NULL || (np->name = strdup(name)) == NULL) {
            return NULL;
        }

        hashval = hash(name);
        np->next = hashtab[hashval];
        hashtab[hashval] = np;
    } else { // already there
        free((void *) np->defn); // free previous defn
    }

    if((np->defn = strdup(defn)) == NULL) {
        return NULL;
    }

    return np;
}

int main() {
    struct nlist *entry;

    install("hello", "world");

    entry = lookup("hello");

    printf("key: %s, value: %s\n", entry->name, entry->defn);
}