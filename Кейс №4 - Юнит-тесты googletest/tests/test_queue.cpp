#include <gtest/gtest.h>
#include "Queue.h"

TEST(QueueTest, PushAndPopSingleElement) {
    Queue q;
    EXPECT_TRUE(q.empty());
    q.push(42);
    EXPECT_FALSE(q.empty());
    EXPECT_EQ(q.size(), 1);
    EXPECT_EQ(q.pop(), 42);
    EXPECT_TRUE(q.empty());
}

TEST(QueueTest, FIFOOrder) {
    Queue q;
    q.push(1);
    q.push(2);
    q.push(3);
    EXPECT_EQ(q.pop(), 1);
    EXPECT_EQ(q.pop(), 2);
    EXPECT_EQ(q.pop(), 3);
    EXPECT_TRUE(q.empty());
}

TEST(QueueTest, PopEmptyThrows) {
    Queue q;
    EXPECT_THROW(q.pop(), std::out_of_range);
}

TEST(QueueTest, MultiplePushPopInterleaved) {
    Queue q;
    q.push(10);
    q.push(20);
    EXPECT_EQ(q.pop(), 10);
    q.push(30);
    EXPECT_EQ(q.pop(), 20);
    EXPECT_EQ(q.pop(), 30);
    EXPECT_TRUE(q.empty());
}

TEST(QueueTest, SizeAfterOperations) {
    Queue q;
    EXPECT_EQ(q.size(), 0);
    q.push(1);
    q.push(2);
    EXPECT_EQ(q.size(), 2);
    q.pop();
    EXPECT_EQ(q.size(), 1);
}
