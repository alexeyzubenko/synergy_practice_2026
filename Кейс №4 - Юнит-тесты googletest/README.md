# Кейс-задача №4 — Юнит-тесты с GoogleTest (C++)

## Описание
Реализованы юнит-тесты с использованием фреймворка **GoogleTest** для классов:
- `Queue` (очередь) — `push`, `pop`
- `Heap` (куча/min-heap) — `push`, `pop`
- `BinaryTree` (бинарное дерево поиска) — `push`, `pop`, `search`

**Важно:** Реализации классов сделаны простыми (на базе std::queue, std::priority_queue и собственной реализации BST) **только для демонстрации тестов**. В задании сказано «сами процедуры реализовывать не нужно, только интерфейс» — ты можешь оставить как есть или заменить реализации на свои.

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

## Как собрать и запустить на Windows (рекомендуемый способ)

### Вариант 1 — Visual Studio (самый простой и стабильный)

1. Установи **Visual Studio Community 2022** (бесплатно) с компонентом **"Desktop development with C++"**.

2. Открой **Developer Command Prompt for VS 2022** (или обычный PowerShell после установки VS).

3. Перейди в папку проекта:
   ```powershell
   cd "C:\Users\alexe\Documents\synergy_practice_2026\Кейс №4 - Юнит-тесты googletest"
   ```

4. Создай папку сборки и скомпилируй:
   ```powershell
   mkdir build
   cd build
   cmake .. -G "Visual Studio 17 2022"
   cmake --build . --config Release
   ```

5. Запусти тесты:
   ```powershell
   .\Release\practice_tests.exe
   ```

Или просто открой файл `practice_tests.sln` в Visual Studio и нажми **Build → Build Solution**, потом запусти.

### Вариант 2 — MinGW (если используешь MSYS2 или MinGW-w64)

```powershell
cd "C:\Users\alexe\Documents\synergy_practice_2026\Кейс №4 - Юнит-тесты googletest"
mkdir build
cd build
cmake .. -G "MinGW Makefiles"
cmake --build .
.\practice_tests.exe
```

## Ожидаемый результат
Все тесты должны пройти успешно:
```
[==========] Running 19 tests from 3 test suites.
...
[  PASSED  ] 19 tests.
```

## Как загрузить на GitHub (если нужно)
1. Создай репозиторий `synergy_practice_2026` (или любой другой).
2. Скопируй содержимое этой папки в репозиторий.
3. `git add .`, `git commit`, `git push`.

Ссылка на репозиторий — ответ на задание.

## Примечание
Если у тебя уже есть папка `synergy_practice_2026`, просто помести эту папку `Кейс №4 - Юнит-тесты googletest` внутрь неё.

Удачи с защитой! Если будут ошибки при сборке — пришли текст ошибки, помогу исправить.
