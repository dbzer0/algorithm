# Pop Count

Алгоритм для расчета популяции: сколько в байте содержится единиц.

Наибольший интерес имеет алгоритм [SWAR](https://graphics.stanford.edu/~seander/bithacks.html#CountBitsSetParallel), выполняющий подсчет со скоростью O(1).