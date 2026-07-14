#include <gtest/gtest.h>
#include "BinaryTree.h"

TEST(BinaryTreeTest, PushAndSearch) {
    BinaryTree tree;
    EXPECT_TRUE(tree.empty());
    tree.push(50);
    EXPECT_FALSE(tree.empty());
    EXPECT_TRUE(tree.search(50));
    EXPECT_FALSE(tree.search(99));
}

TEST(BinaryTreeTest, SearchAfterMultipleInserts) {
    BinaryTree tree;
    tree.push(50);
    tree.push(30);
    tree.push(70);
    tree.push(20);
    tree.push(40);
    tree.push(60);
    tree.push(80);

    EXPECT_TRUE(tree.search(20));
    EXPECT_TRUE(tree.search(40));
    EXPECT_TRUE(tree.search(60));
    EXPECT_TRUE(tree.search(80));
    EXPECT_FALSE(tree.search(25));
    EXPECT_FALSE(tree.search(55));
}

TEST(BinaryTreeTest, PopLeafNode) {
    BinaryTree tree;
    tree.push(50);
    tree.push(30);
    tree.push(70);
    EXPECT_TRUE(tree.pop(30));
    EXPECT_FALSE(tree.search(30));
    EXPECT_TRUE(tree.search(50));
    EXPECT_TRUE(tree.search(70));
}

TEST(BinaryTreeTest, PopNodeWithOneChild) {
    BinaryTree tree;
    tree.push(50);
    tree.push(30);
    tree.push(20);
    EXPECT_TRUE(tree.pop(30));
    EXPECT_FALSE(tree.search(30));
    EXPECT_TRUE(tree.search(20));
    EXPECT_TRUE(tree.search(50));
}

TEST(BinaryTreeTest, PopNodeWithTwoChildren) {
    BinaryTree tree;
    tree.push(50);
    tree.push(30);
    tree.push(70);
    tree.push(20);
    tree.push(40);
    EXPECT_TRUE(tree.pop(30));
    EXPECT_FALSE(tree.search(30));
    EXPECT_TRUE(tree.search(20));
    EXPECT_TRUE(tree.search(40));
    EXPECT_TRUE(tree.search(50));
    EXPECT_TRUE(tree.search(70));
}

TEST(BinaryTreeTest, PopRoot) {
    BinaryTree tree;
    tree.push(50);
    tree.push(30);
    tree.push(70);
    EXPECT_TRUE(tree.pop(50));
    EXPECT_FALSE(tree.search(50));
    EXPECT_TRUE(tree.search(30));
    EXPECT_TRUE(tree.search(70));
}

TEST(BinaryTreeTest, PopNonExistent) {
    BinaryTree tree;
    tree.push(10);
    EXPECT_FALSE(tree.pop(99));
    EXPECT_TRUE(tree.search(10));
}

TEST(BinaryTreeTest, PopEmptyTree) {
    BinaryTree tree;
    EXPECT_FALSE(tree.pop(42));
}

TEST(BinaryTreeTest, DuplicateInsertIgnored) {
    BinaryTree tree;
    tree.push(42);
    tree.push(42);
    EXPECT_TRUE(tree.search(42));
}
