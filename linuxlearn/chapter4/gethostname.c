#include<sys/utsname.h>
#include<unistd.h>
#include<stdlib.h>
#include<stdio.h>

int main()
{
    char computer[256];
    struct utsname uts;

    if(gethostname(computer, 255) !=0 || uname(&uts)<0) {
        fprintf(stderr, "Could not get host information\n");
    }

    printf("Computer host name %s\n", computer);
    printf("System is %s on %s hardware\n", uts.sysname, uts.machine);
    printf("Nodename is %s\n", uts.nodename);
    printf("Version is %s, %s\n", uts. release, uts.version);

    printf("%ld\n", gethostid());
    return 0;
}