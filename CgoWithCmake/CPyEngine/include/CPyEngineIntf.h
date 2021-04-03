// ExtC.h : Include file for standard system include files,
// or project specific include files.

#ifdef __cplusplus
extern "C" {
#endif
  typedef const char cchar_t;

  extern char * GetCmd();
  extern void GoPrint(cchar_t *);

  void StartPyEngine();
#ifdef __cplusplus
}
#endif

#ifndef DEBUG
#define DEBUG
#endif
