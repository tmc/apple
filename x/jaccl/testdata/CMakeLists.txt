cmake_minimum_required(VERSION 3.24)
project(jaccl_hardware_oracle LANGUAGES CXX)

set(JACCL_SOURCE "" CACHE PATH "Path to the upstream distributed/jaccl/lib source directory")
if(NOT JACCL_SOURCE)
  message(FATAL_ERROR "set JACCL_SOURCE to upstream distributed/jaccl/lib")
endif()

add_subdirectory("${JACCL_SOURCE}" jaccl)

add_executable(upstream_hardware_oracle upstream_hardware_oracle.cpp)
target_link_libraries(upstream_hardware_oracle PRIVATE jaccl)
