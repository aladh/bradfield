#include <stdio.h>

int main(int argc, char *argv[], char* envp[]) {
  char *env;
  while((env = *envp++)) {
    printf("%s\n", env);
  }
}
