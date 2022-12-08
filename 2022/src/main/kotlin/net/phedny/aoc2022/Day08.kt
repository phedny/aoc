package net.phedny.aoc2022

fun main() {
    val grid = readInput().split("\n")
        .map { it.map(Char::digitToInt) }

    grid
        .mapCells { v, left, right, up, down ->
            if (left.all { it < v } || right.all { it < v } || up.all { it < v } || down.all { it < v }) 1 else 0
        }
        .sumOf { it.sum() }
        .also(::println)

    grid
        .mapCells { v, left, right, up, down ->
            left.visibleAt(v) * right.visibleAt(v) * up.visibleAt(v) * down.visibleAt(v)
        }
        .maxOf { it.maxOrNull()!! }
        .also(::println)
}

fun <T, R> List<List<T>>.mapCells(transform: (value: T, left: List<T>, right: List<T>, up: List<T>, down: List<T>) -> R): List<List<R>> =
    this.indices.map { row ->
        this.first().indices.map { col ->
            transform(
                this[row][col],
                this[row].subList(0, col).reversed(),
                this[row].let { it.subList(col + 1, it.size) },
                this.subList(0, row).map { it[col] }.reversed(),
                this.subList(row + 1, this.size).map { it[col] }
            )
        }
    }

fun List<Int>.visibleAt(height: Int): Int =
    this.indexOfFirst { it >= height }
        .let { if (it == -1) this.size else it + 1 }
