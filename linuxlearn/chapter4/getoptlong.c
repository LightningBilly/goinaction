#include<stdio.h>
#include<unistd.h>
#include<stdlib.h>
#include<getopt.h>

int main(int argc, char *argv[]) {
    int opt;
    struct option longopts[] = {
        {"initialize", 0, NULL, 'i'},
        {"file", 1, NULL, 'f'},
        {"list", 0, NULL, 'l'},
        {"restart", 0, NULL, 'r'},
        {0,0,0,0}
    };
    while((opt=getopt_long(argc, argv, ":x::if:lr", longopts, NULL))!=-1) {
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