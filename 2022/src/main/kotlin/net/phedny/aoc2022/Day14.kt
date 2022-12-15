package net.phedny.aoc2022

fun main() {
    val map = readInput().split("\n")
        .flatMap { line -> line.split(" -> ")
            .map { it.split(",").map(String::toInt) }
            .map { (first, second) -> Pair(first, second) }
            .windowed(2, 1)
            .flatMap { (start, end) ->
                when {
                    start.first < end.first -> (start.first .. end.first).map { Pair(it, start.second) }
                    start.first > end.first -> (start.first downTo end.first).map { Pair(it, start.second) }
                    start.second > end.second -> (start.second downTo end.second).map { Pair(start.first, it) }
                    start.second < end.second -> (start.second..end.second).map { Pair(start.first, it) }
                    else -> setOf(start)
                }
            }
        }.toSet()

    val spawningPoint = Pair(500, 0)
    val lowestRock = map.maxOf(Pair<Int, Int>::second)

    val seq = generateSequence(Pair(map, spawningPoint)) { (map, sand) ->
        if (sand.second > lowestRock)
            Pair(map + sand, spawningPoint)
        else
            listOf(Pair(sand.first, sand.second + 1), Pair(sand.first - 1, sand.second + 1), Pair(sand.first + 1, sand.second + 1))
                .find { !map.contains(it) }
                ?.let { Pair(map, it) }
                ?: Pair(map + sand, spawningPoint)
    }

    println(seq.first { (_, sand) -> sand.second > lowestRock }.first.size - map.size)
    println(seq.first { (map) -> map.contains(spawningPoint) }.first.size - map.size)
}
