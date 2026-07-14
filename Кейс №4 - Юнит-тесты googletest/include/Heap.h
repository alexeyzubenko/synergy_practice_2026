#ifndef HEAP_H
#define HEAP_H

#include <queue>
#include <stdexcept>

class Heap {
private:
    std::priority_queue<int> data; // max-heap

public:
    void push(int value);
    int pop();
    bool empty() const;
    size_t size() const;
};

#endif // HEAP_H
