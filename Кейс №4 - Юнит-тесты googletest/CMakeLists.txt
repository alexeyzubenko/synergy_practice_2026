cmake_minimum_required(VERSION 3.14)
project(GoogleTestPractice LANGUAGES CXX)

set(CMAKE_CXX_STANDARD 17)
set(CMAKE_CXX_STANDARD_REQUIRED ON)

include(FetchContent)
FetchContent_Declare(
  googletest
  URL https://github.com/google/googletest/archive/03597a01ee50ed33e9dfd640b249b4be3799d395.zip
  DOWNLOAD_EXTRACT_TIMESTAMP TRUE
)
set(gtest_force_shared_crt ON CACHE BOOL "" FORCE)
FetchContent_MakeAvailable(googletest)

enable_testing()

add_executable(
  practice_tests
  tests/test_queue.cpp
  tests/test_heap.cpp
  tests/test_binary_tree.cpp
  src/Queue.cpp
  src/Heap.cpp
  src/BinaryTree.cpp
)

target_include_directories(practice_tests PRIVATE ${CMAKE_CURRENT_SOURCE_DIR}/include)

target_link_libraries(
  practice_tests
  GTest::gtest_main
)

include(GoogleTest)
gtest_discover_tests(practice_tests)
