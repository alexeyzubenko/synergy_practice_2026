#ifndef QUEUE_H
#define QUEUE_H

#include <queue>
#include <stdexcept>

class Queue {
private:
    std::queue<int> data;

public:
    void push(int value);
    int pop();
    bool empty() const;
    size_t size() const;
};

#endif // QUEUE_H
