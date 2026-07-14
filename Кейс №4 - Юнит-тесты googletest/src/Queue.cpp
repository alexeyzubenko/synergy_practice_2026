#include "Queue.h"

void Queue::push(int value) {
    data.push(value);
}

int Queue::pop() {
    if (data.empty()) {
        throw std::out_of_range("Queue is empty");
    }
    int front = data.front();
    data.pop();
    return front;
}

bool Queue::empty() const {
    return data.empty();
}

size_t Queue::size() const {
    return data.size();
}
