#include "BinaryTree.h"

BinaryTree::BinaryTree() : root(nullptr) {}

BinaryTree::~BinaryTree() {
    destroyTree(root);
}

void BinaryTree::destroyTree(Node* node) {
    if (node) {
        destroyTree(node->left);
        destroyTree(node->right);
        delete node;
    }
}

Node* BinaryTree::insert(Node* node, int value) {
    if (!node) {
        return new Node(value);
    }
    if (value < node->data) {
        node->left = insert(node->left, value);
    } else if (value > node->data) {
        node->right = insert(node->right, value);
    }
    return node;
}

void BinaryTree::push(int value) {
    root = insert(root, value);
}

bool BinaryTree::search(Node* node, int value) const {
    if (!node) return false;
    if (value == node->data) return true;
    if (value < node->data) return search(node->left, value);
    return search(node->right, value);
}

bool BinaryTree::search(int value) const {
    return search(root, value);
}

Node* BinaryTree::findMin(Node* node) {
    while (node && node->left) {
        node = node->left;
    }
    return node;
}

Node* BinaryTree::remove(Node* node, int value) {
    if (!node) return nullptr;

    if (value < node->data) {
        node->left = remove(node->left, value);
    } else if (value > node->data) {
        node->right = remove(node->right, value);
    } else {
        if (!node->left) {
            Node* temp = node->right;
            delete node;
            return temp;
        } else if (!node->right) {
            Node* temp = node->left;
            delete node;
            return temp;
        } else {
            Node* temp = findMin(node->right);
            node->data = temp->data;
            node->right = remove(node->right, temp->data);
        }
    }
    return node;
}

bool BinaryTree::pop(int value) {
    if (!search(value)) {
        return false;
    }
    root = remove(root, value);
    return true;
}

bool BinaryTree::empty() const {
    return root == nullptr;
}
