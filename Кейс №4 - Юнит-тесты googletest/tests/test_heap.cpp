#include <gtest/gtest.h>
#include "Heap.h"

TEST(HeapTest, PushAndPopSingleElement) {
    Heap h;
    EXPECT_TRUE(h.empty());
    h.push(99);
    EXPECT_FALSE(h.empty());
    EXPECT_EQ(h.size(), 1);
    EXPECT_EQ(h.pop(), 99);
    EXPECT_TRUE(h.empty());
}

TEST(HeapTest, MaxHeapOrder) {
    Heap h;
    h.push(5);
    h.push(10);
    h.push(3);
    h.push(8);
    EXPECT_EQ(h.pop(), 10);
    EXPECT_EQ(h.pop(), 8);
    EXPECT_EQ(h.pop(), 5);
    EXPECT_EQ(h.pop(), 3);
    EXPECT_TRUE(h.empty());
}

TEST(HeapTest, PopEmptyThrows) {
    Heap h;
    EXPECT_THROW(h.pop(), std::out_of_range);
}

TEST(HeapTest, DuplicateValues) {
    Heap h;
    h.push(7);
    h.push(7);
    h.push(7);
    EXPECT_EQ(h.size(), 3);
    EXPECT_EQ(h.pop(), 7);
    EXPECT_EQ(h.pop(), 7);
    EXPECT_EQ(h.pop(), 7);
}

TEST(HeapTest, SizeAfterOperations) {
    Heap h;
    EXPECT_EQ(h.size(), 0);
    h.push(1);
    h.push(2);
    h.push(3);
    EXPECT_EQ(h.size(), 3);
    h.pop();
    EXPECT_EQ(h.size(), 2);
}
