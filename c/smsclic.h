// smsclic.h

#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <curl/curl.h>
#include <fcntl.h>
#include <errno.h>
#include <unistd.h>
#include "ini.h"

// Defines
// #define  MENUTEXT          (BLACK + (LIGHTGRAY <<4))
// #define  K_ESC             (0x011B)
// #define  lengthof(x)       ((sizeof(x))/(sizeof(x[0]))) 

// Global constants and variables declaration
extern const char prgname[64];
extern const char prgver[16];
extern char cfgfile[64]; 
extern char svcurl[256];
extern char svckey[256];

// Types
// typedef unsigned char       BYTE;     // unsigned 8-bit number
// typedef unsigned int        WORD;     // unsigned 16-bit number
// typedef unsigned long       DWORD;    // unsigned 32-bit number
// typedef enum {FALSE, TRUE}  BOOL;

// Functions prototypes
void PrintUsage(void);
void ProcessConfigFile(void);
void CreateSkeletonConfigFile(void);
void CURLPostData(char * url, char * data);
