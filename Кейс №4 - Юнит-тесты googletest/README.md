# Кейс-задача №4 — Юнит-тесты с GoogleTest (C++)

## Описание задания
Реализовать ряд юнит-тестов с использованием фреймворка googletest для каждого из следующих классов (сами процедуры реализовывать не нужно, только интерфейс): 
- Класса очереди (queue) с функциями pop, push
- Класса кучи (heap) с функциями pop, push
- Класса бинарного дерева с функциями pop, push, search

## Реализация

Реализованы юнит-тесты с использованием фреймворка **GoogleTest** для классов:
- `Queue` (очередь) — `push`, `pop`
- `Heap` (куча/min-heap) — `push`, `pop`
- `BinaryTree` (бинарное дерево поиска) — `push`, `pop`, `search`


## Структура проекта
```
Кейс №4 - Юнит-тесты googletest/
├── CMakeLists.txt
├── include/
│   ├── Queue.h
│   ├── Heap.h
│   └── BinaryTree.h
├── src/
│   ├── Queue.cpp
│   ├── Heap.cpp
│   └── BinaryTree.cpp
├── tests/
│   ├── test_queue.cpp
│   ├── test_heap.cpp
│   └── test_binary_tree.cpp
└── README.md
```

## Как запустить тесты

Visual Studio 2022/2026 с поддержкой CMake:

1. Открыть **Visual Studio 2026**.
2. Выбрать **File → Open → Folder...**
3. Указать папку `Кейс №4 - Юнит-тесты googletest`.
4. Дождаться окончания настройки CMake.
6. Запустить `practice_tests.exe`.
