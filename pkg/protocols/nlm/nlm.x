package "github.com/buildbarn/go-xdr/pkg/protocols/nlm";

/* Definitions copied from RFC 1813, appendix II. */

/* 6.1.4 Basic Data Types. */

typedef unsigned hyper uint64;
typedef hyper int64;
typedef unsigned int uint32;
typedef int int32;

enum nlm4_stats {
   NLM4_GRANTED = 0,
   NLM4_DENIED = 1,
   NLM4_DENIED_NOLOCKS = 2,
   NLM4_BLOCKED = 3,
   NLM4_DENIED_GRACE_PERIOD = 4,
   NLM4_DEADLCK = 5,
   NLM4_ROFS = 6,
   NLM4_STALE_FH = 7,
   NLM4_FBIG = 8,
   NLM4_FAILED = 9
};

struct nlm4_holder {
     bool     exclusive;
     int32    svid;
     netobj   oh;
     uint64   l_offset;
     uint64   l_len;
};

struct nlm4_lock {
     string   caller_name<LM_MAXSTRLEN>;
     netobj   fh;
     netobj   oh;
     int32    svid;
     uint64   l_offset;
     uint64   l_len;
};

struct nlm4_share {
     string      caller_name<LM_MAXSTRLEN>;
     netobj      fh;
     netobj      oh;
     fsh4_mode   mode;
     fsh4_access access;
};

/* 6.2 NLM Procedures. */

program NLM_PROG {
   version NLM4_VERS {
      void
         NLMPROC4_NULL(void)                  = 0;

      nlm4_testres
         NLMPROC4_TEST(nlm4_testargs)         = 1;

      nlm4_res
         NLMPROC4_LOCK(nlm4_lockargs)         = 2;

      nlm4_res
         NLMPROC4_CANCEL(nlm4_cancargs)       = 3;

      nlm4_res
         NLMPROC4_UNLOCK(nlm4_unlockargs)     = 4;

      nlm4_res
         NLMPROC4_GRANTED(nlm4_testargs)      = 5;

      void
         NLMPROC4_TEST_MSG(nlm4_testargs)     = 6;

      void
         NLMPROC4_LOCK_MSG(nlm4_lockargs)     = 7;

      void
         NLMPROC4_CANCEL_MSG(nlm4_cancargs)   = 8;

      void
         NLMPROC4_UNLOCK_MSG(nlm4_unlockargs) = 9;

      void
         NLMPROC4_GRANTED_MSG(nlm4_testargs) = 10;

      void
         NLMPROC4_TEST_RES(nlm4_testres)     = 11;

      void
         NLMPROC4_LOCK_RES(nlm4_res)         = 12;

      void
         NLMPROC4_CANCEL_RES(nlm4_res)       = 13;

      void
         NLMPROC4_UNLOCK_RES(nlm4_res)       = 14;

      void
         NLMPROC4_GRANTED_RES(nlm4_res)      = 15;

      nlm4_shareres
         NLMPROC4_SHARE(nlm4_shareargs)      = 20;

      nlm4_shareres
         NLMPROC4_UNSHARE(nlm4_shareargs)    = 21;

      nlm4_res
         NLMPROC4_NM_LOCK(nlm4_lockargs)     = 22;

      void
         NLMPROC4_FREE_ALL(nlm4_notify)      = 23;

   } = 4;
} = 100021;

/* https://pubs.opengroup.org/onlinepubs/9629799/chap10.htm */

const LM_MAXSTRLEN = 1024;
const LM_MAXNAMELEN = 1025;
const MAXNETOBJ_SZ = 1024;

typedef opaque netobj<MAXNETOBJ_SZ>;

struct nlm4_stat {
    nlm4_stats stat;
};

struct nlm4_res {
    netobj cookie;
    nlm4_stat stat;
};

union nlm4_testrply switch (nlm4_stats stat) {
    case NLM4_DENIED:
        nlm4_holder holder;    /*  holder of the lock */
    default:
        void;
};

struct nlm4_testres {
    netobj cookie;
    nlm4_testrply test_stat;
};

struct nlm4_lockargs {
    netobj cookie;
    bool block;      /*  Flag to indicate blocking behaviour. */
    bool exclusive;  /*  If exclusive access is desired. */
    nlm4_lock alock; /*  The actual lock data (see above) */
    bool reclaim;    /*  used for recovering locks  */
    int state;       /*  specify local NSM state  */
};

struct nlm4_cancargs {
    netobj cookie;
    bool block;
    bool exclusive;
    nlm4_lock alock;
};

struct nlm4_testargs {
    netobj cookie;
    bool exclusive;
    nlm4_lock alock;
};

struct nlm4_unlockargs {
    netobj cookie;
    nlm4_lock alock;
};

enum fsh4_mode {
    fsm_DN = 0,        /*  deny none  */
    fsm_DR = 1,        /*  deny read  */
    fsm_DW = 2,        /*  deny write  */
    fsm_DRW = 3        /*  deny read/write  */
};

enum fsh4_access {
    fsa_NONE = 0,     /*  for completeness  */
    fsa_R = 1,        /*  read-only  */
    fsa_W = 2,        /*  write-only  */
    fsa_RW = 3        /*  read/write  */
};

struct nlm4_shareargs {
    netobj cookie;
    nlm4_share share;    /*  actual share data  */
    bool reclaim;        /*  used for recovering shares  */
};

struct nlm4_shareres {
    netobj cookie;
    nlm4_stats stat;
    int sequence;
};

struct nlm4_notify {
    string name<LM_MAXNAMELEN>;
    int state;
};
