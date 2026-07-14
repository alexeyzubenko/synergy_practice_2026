#include "Heap.h"

void Heap::push(int value) {
    data.push(value);
}

int Heap::pop() {
    if (data.empty()) {
        throw std::out_of_range("Heap is empty");
    }
    int top = data.top();
    data.pop();
    return top;
}

bool Heap::empty() const {
    return data.empty();
}

size_t Heap::size() const {
    return data.size();
}
