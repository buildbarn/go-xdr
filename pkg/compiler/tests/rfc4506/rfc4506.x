package "github.com/buildbarn/go-xdr/pkg/compiler/tests/rfc4506";

/*
 * Code from RFC 4506.
 */

/* Section 4.17. */

const DOZEN = 12;

/* Section 4.18. */

typedef int egg;
typedef egg eggbox[DOZEN];

struct eggs {
   eggbox  fresheggs1;
   egg     fresheggs2[DOZEN];
};

/* Section 4.19. */

struct stringentry1 {
   string item<>;
   stringentry1 *next;
};

typedef stringentry1 *stringlist1;

union stringlist2 switch (bool opted) {
case TRUE:
   struct {
      string item<>;
      stringlist2 next;
   } element;
case FALSE:
   void;
};

struct stringentry3 {
   string item<>;
   stringentry3 next<1>;
};

typedef stringentry3 stringlist3<1>;

/* Chapter 7. */

const MAXUSERNAME = 32;     /* max length of a user name */
const MAXFILELEN = 65535;   /* max length of a file      */
const MAXNAMELEN = 255;     /* max length of a file name */

/*
 * Types of files:
 */
enum filekind {
   TEXT = 0,       /* ascii data */
   DATA = 1,       /* raw data   */
   EXEC = 2        /* executable */
};

/*
 * File information, per kind of file:
 */
union filetype switch (filekind kind) {
case TEXT:
   void;                           /* no extra information */
case DATA:
   string creator<MAXNAMELEN>;     /* data creator         */
case EXEC:
   string interpretor<MAXNAMELEN>; /* program interpretor  */
};

/*
 * A complete file:
 */
struct file {
   string filename<MAXNAMELEN>; /* name of file    */
   filetype type;               /* info about file */
   string owner<MAXUSERNAME>;   /* owner of file   */
   opaque data<MAXFILELEN>;     /* file data       */
};
