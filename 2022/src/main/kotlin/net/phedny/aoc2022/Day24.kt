package net.phedny.aoc2022

data class Blizzard(val x: Int, val y: Int, val dx: Int, val dy: Int, val height: Int, val width: Int) {
    fun move(): Blizzard = copy(x = (x + dx) % height, y = (y + dy) % width)
}

fun main() {
    val input = readInput().split("\n")
        .drop(1)
        .dropLast(1)

    val height = input.size
    val width = input.first().length - 2

    val blizzards = input.flatMapIndexed { x, row ->
        row.drop(1)
            .dropLast(1)
            .mapIndexed { y, cell ->
                when (cell) {
                    '>' -> Blizzard(x, y, 0, 1, height, width)
                    '<' -> Blizzard(x, y, 0, width - 1, height, width)
                    '^' -> Blizzard(x, y, height - 1, 0, height, width)
                    'v' -> Blizzard(x, y, 1, 0, height, width)
                    else -> null
                }
            }
    }.filterNotNull()

    val goals = listOf(Pair(height, width - 1), Pair(-1, 0), Pair(height, width - 1))
    val s = generateSequence(Triple(listOf(Pair(-1, 0)), blizzards, goals)) { (positions, blizzards, goals) ->
        val newBlizzards = blizzards.map(Blizzard::move)
        val newPositions = positions.flatMap { (x, y) -> listOf(Pair(x, y), Pair(x + 1, y), Pair(x - 1, y), Pair(x, y + 1), Pair(x, y - 1)) }
            .distinct()
            .filter { (x, y) -> (x in 0 until height && y in 0 until width) || (x == -1 && y == 0) || (x == height && y == width - 1) }
            .filter { (x, y) -> !newBlizzards.any { it.x == x && it.y == y} }

        when {
            !newPositions.contains(goals.first()) -> Triple(newPositions, newBlizzards, goals)
            goals.size > 1 -> Triple(listOf(goals.first()), newBlizzards, goals.drop(1))
            else -> null
        }
    }

    println(s.indexOfFirst { (positions) -> positions.any { it.first == height && it.second == width - 1 } })
    println(s.count())
}
