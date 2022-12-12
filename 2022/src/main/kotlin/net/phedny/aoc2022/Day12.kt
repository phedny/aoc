package net.phedny.aoc2022

fun main() {
    val input = readInput().split("\n")

    val height = input.size
    val width = input.first().length
    val start = input.indexOfFirst { it.contains('S') }.let { Pair(it, input[it].indexOf('S')) }
    val end = input.indexOfFirst { it.contains('E') }.let { Pair(it, input[it].indexOf('E')) }

    val grid = input.map { line ->
        line.map { char ->
            when (char) {
                'S' -> 0
                'E' -> 25
                else -> char - 'a'
            }
        }
    }

    val s = generateSequence(Triple(mapOf(end to 0), listOf(end), 1)) { (known, front, step) ->
        front.flatMap { (x, y) ->
            listOf(Pair(x - 1, y), Pair(x, y - 1), Pair(x + 1, y), Pair(x, y + 1))
                .filterNot(known::contains)
                .filterNot { (newX, newY) -> newX < 0 || newX >= height || newY < 0 || newY >= width }
                .filter { (newX, newY) -> grid[newX][newY] >= grid[x][y] - 1 }
        }.let { newFront -> Triple(known + newFront.map { it to step }, newFront.distinct(), step + 1) }
    }

    s.dropWhile { !it.first.containsKey(start) }
        .first()
        .also { println(it.third - 1) }

    s.dropWhile { it.first.keys.none { (x, y) -> grid[x][y] == 0 } }
        .first()
        .also { println(it.third - 1) }
}
