#include<stdio.h>
#include<unistd.h>
#include<stdlib.h>

int main(int argc, char *argv[]) {
    int opt;
    while((opt=getopt(argc, argv, ":x::if:lr"))!=-1) {
        switch(opt) {
            case 'i':
            case 'l':
            case 'r':
                printf("option : %c\n", opt);
                break;
            case 'x':
            case 'f':
                printf("option:%c, name:%s\n", opt, optarg);
                break;
            case ':':
                printf("option needs a value\n");
                break;
            case '?':
                printf("unknown option: %c\n", optopt);
                break;
        }
    }

    for(; optind < argc; optind++) {
        printf("argument: %s\n", argv[optind]);
    }

    return 0;
}