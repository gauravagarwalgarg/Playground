cmake_minimum_required(VERSION 3.16)
project(InterviewPrep LANGUAGES CXX)

set(CMAKE_CXX_STANDARD 20)
set(CMAKE_CXX_STANDARD_REQUIRED ON)
set(CMAKE_RUNTIME_OUTPUT_DIRECTORY ${CMAKE_BINARY_DIR}/bin)

# Auto-discover all .cpp files with main() and build each as a standalone executable
file(GLOB_RECURSE ALL_SOURCES
    "${CMAKE_SOURCE_DIR}/algorithms-dsa/**/*.cpp"
    "${CMAKE_SOURCE_DIR}/low-level-design-lld/**/*.cpp"
    "${CMAKE_SOURCE_DIR}/computer-core/**/*.cpp"
    "${CMAKE_SOURCE_DIR}/application-core/**/*.cpp"
)

foreach(SOURCE_FILE ${ALL_SOURCES})
    # Derive executable name from relative path
    file(RELATIVE_PATH REL_PATH ${CMAKE_SOURCE_DIR} ${SOURCE_FILE})
    string(REPLACE "/" "_" EXEC_NAME ${REL_PATH})
    string(REPLACE ".cpp" "" EXEC_NAME ${EXEC_NAME})

    add_executable(${EXEC_NAME} ${SOURCE_FILE})
endforeach()
