# CMakeList.txt : CMake project for ExtC, include source and define
# project specific logic here.
#
cmake_minimum_required (VERSION 3.8)

project ("ExtC")

# Add source to this project's library.
add_library(ExtC_library STATIC
    src/ExtC.cpp
)

target_include_directories(ExtC_library
    PUBLIC 
        ${PROJECT_SOURCE_DIR}/include
)

foreach(flag_var
        CMAKE_CXX_FLAGS CMAKE_CXX_FLAGS_DEBUG CMAKE_CXX_FLAGS_RELEASE
        CMAKE_CXX_FLAGS_MINSIZEREL CMAKE_CXX_FLAGS_RELWITHDEBINFO)
    STRING (REGEX REPLACE "/RTC[^ ]*" "" ${flag_var} "${${flag_var}}")
endforeach(flag_var)

# TODO: Add tests and install targets if needed.
