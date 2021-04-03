// ExtC.cpp : Defines the entry point for the application.
//

#include "ExtC.h"

void ProcCmd(char* cMsg) {
    GoPrint(cMsg);
}

void StartProc() {
    GoPrint("Enter StartProc");
    while (true) {
        char* cmdStr;
        cmdStr = GetCmd();
        ProcCmd(cmdStr);
        GoPrint(cmdStr);
    }
}
