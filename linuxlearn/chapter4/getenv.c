#include<stdio.h>
#include<stdlib.h>
#include<string.h>


int main(int argc, char * argv[]) {
    char *var, *value;

    if (argc == 1 || argc > 3) {
        fprintf(stderr, "usage: environ var [value]\n");
        exit(1);
    }

    putenv("a=b");
    var = argv[1];
    value = getenv(var);
    if (value)
        printf("Variable %s has value %s\n", var, value);
    else
        printf("Variable %s has no value\n", var);

    return 0;
}
