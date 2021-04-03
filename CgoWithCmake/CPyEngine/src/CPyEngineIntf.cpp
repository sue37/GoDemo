#include <cstddef>
#include <exception>
#include "CPyEngineIntf.h"
//#include "PyEngine.h"

#include <pybind11/embed.h> // everything needed for embedding
namespace py = pybind11;
using namespace std;

void InitEngine() {
#ifdef DEBUG
    GoPrint("Hello from PyEngine");
#endif // DEBUG
    try {
        //py::exec("import PCellBase as pcb");
        GoPrint("InitEngine");
    }
    catch (exception e) {
        GoPrint(e.what());
    }
}

const char* ProcCmd(char* cMsg)  {
#ifdef DEBUG
    GoPrint(cMsg);
#endif
    try {
        //py::exec(cMsg);
        GoPrint(cMsg);
    }
    catch (const exception& e) {
        return e.what();
    }
    return NULL;
}

void StartPyEngine() {
#ifdef DEBUG
   GoPrint("Enter StartPyEngine");
#endif
   py::scoped_interpreter guard{};
   InitEngine();
   while (1) {
      char *cmdStr;
      cmdStr = GetCmd();
      ProcCmd(cmdStr);
   }
}
