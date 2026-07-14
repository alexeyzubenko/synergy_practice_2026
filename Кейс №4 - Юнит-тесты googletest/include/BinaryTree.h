#ifndef BINARY_TREE_H
#define BINARY_TREE_H

#include <stdexcept>

struct Node {
    int data;
    Node* left;
    Node* right;
    Node(int val) : data(val), left(nullptr), right(nullptr) {}
};

class BinaryTree {
private:
    Node* root;

    Node* insert(Node* node, int value);
    bool search(Node* node, int value) const;
    Node* findMin(Node* node);
    Node* remove(Node* node, int value);
    void destroyTree(Node* node);

public:
    BinaryTree();
    ~BinaryTree();

    void push(int value);
    bool pop(int value);
    bool search(int value) const;
    bool empty() const;
};

#endif // BINARY_TREE_H
