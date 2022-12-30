package net.phedny.aoc2022

fun main() {
    val input = readInput().split("\n")
        .flatMapIndexed { rowId, row -> row.mapIndexed { colId, cell -> if (cell == '#') Pair(rowId, colId) else null } }
        .filterNotNull()

    val startChecks: List<Pair<Collection<Pair<Int, Int>>, Pair<Int, Int>>> = listOf(
        listOf(Pair(-1, -1), Pair(-1, 0), Pair(-1, +1)) to Pair(-1, 0),
        listOf(Pair(+1, -1), Pair(+1, 0), Pair(+1, +1)) to Pair(+1, 0),
        listOf(Pair(-1, -1), Pair(0, -1), Pair(+1, -1)) to Pair(0, -1),
        listOf(Pair(-1, +1), Pair(0, +1), Pair(+1, +1)) to Pair(0, +1)
    )

    val s = generateSequence(Pair(input.toSet(), startChecks)) { (elves, checks) ->
        val newPositions = elves.map { (x, y) ->
            if (checks.all { check -> (check.first.map { x + it.first to y + it.second } intersect elves).isEmpty() }) {
                Pair(x to y, x to y)
            } else {
                checks.firstOrNull { check -> (check.first.map { x + it.first to y + it.second } intersect elves).isEmpty() }
                    ?.second?.let { Pair(x to y, x + it.first to y + it.second) }
                    ?: Pair(x to y, x to y)
            }
        }
            .groupBy(Pair<Pair<Int, Int>, Pair<Int, Int>>::second)
            .flatMap { (key, value) -> if (value.size == 1) listOf(key) else value.map(Pair<Pair<Int, Int>, Pair<Int, Int>>::first) }

        Pair(newPositions.toSet(), checks.drop(1) + checks.first())
    }
        .map(Pair<Set<Pair<Int, Int>>, List<Pair<Collection<Pair<Int, Int>>, Pair<Int, Int>>>>::first)

    s.drop(10)
        .first()
        .let { (it.maxOf(Pair<Int, Int>::first) - it.minOf(Pair<Int, Int>::first) + 1) * (it.maxOf(Pair<Int, Int>::second) - it.minOf(Pair<Int, Int>::second) + 1) - it.size }
        .also(::println)

    s.windowed(2, 1)
        .indexOfFirst { (a, b) -> a == b }
        .also { println(it + 1) }
}
