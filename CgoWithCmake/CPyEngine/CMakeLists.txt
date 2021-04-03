# CMakeList.txt : CMake project for ExtC, include source and define
# project specific logic here.
#
cmake_minimum_required (VERSION 3.8)

project (
    CPyEngine 
    LANGUAGES CXX
)

add_subdirectory(extern/pybind11)

set (CMAKE_C_FLAGS "${CMAKE_C_FLAGS} -fPIC")
set (CMAKE_CXX_FLAGS "${CMAKE_CXX_FLAGS} -fPIC")

# Add source to this project's library.
# STATIC
add_library(CPyEngine_library SHARED
    src/CPyEngineIntf.cpp)

target_include_directories(CPyEngine_library
    PUBLIC 
       ${PROJECT_SOURCE_DIR}/include
)

target_link_libraries(CPyEngine_library PRIVATE pybind11::embed)

foreach(flag_var
        CMAKE_CXX_FLAGS CMAKE_CXX_FLAGS_DEBUG CMAKE_CXX_FLAGS_RELEASE
        CMAKE_CXX_FLAGS_MINSIZEREL CMAKE_CXX_FLAGS_RELWITHDEBINFO)
    STRING (REGEX REPLACE "/RTC[^ ]*" "" ${flag_var} "${${flag_var}}")
endforeach(flag_var)

# TODO: Add tests and install targets if needed.
