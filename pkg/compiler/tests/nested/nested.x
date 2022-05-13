package "github.com/buildbarn/go-xdr/pkg/compiler/tests/nested";

struct a {
  struct {
    struct {
      enum {
        FOO = 1,
        BAR = 2
      } d;
      union switch(enum {
        ABRA = 3,
        CADABRA = 4
      } i) {
        case ABRA: enum {
          CAR = 1,
          BOAT = 2
        } j;
        case CADABRA: struct {
          int l;
        } k;
      } e;
    } c;
  } b;
};

typedef enum {
  DOG = 0,
  CAT = 1
} *pointer_to_enum;
