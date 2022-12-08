package net.phedny.aoc2022

fun main() {
    val grid = readInput().split("\n")
        .map { it.map(Char::digitToInt) }

    grid.mapCells(Boolean::or) { v, line -> line.all { it < v } }
        .sumOf { it.count { b -> b } }
        .also(::println)

    grid.mapCells(Int::times) { v, line -> line.indexOfFirst { it >= v }.let { if (it == -1) line.size else it + 1 } }
        .maxOf { it.maxOrNull()!! }
        .also(::println)
}

fun <T, R> List<List<T>>.mapCells(combine: (R, R) -> R, transform: (T, List<T>) -> R): List<List<R>> =
    this.mapIndexed { rowIdx, row ->
        row.mapIndexed { colIdx, value ->
            listOf(
                transform(value, row.take(colIdx).reversed()),
                transform(value, row.drop(colIdx + 1)),
                transform(value, this.take(rowIdx).map { it[colIdx] }.reversed()),
                transform(value, this.drop(rowIdx + 1).map { it[colIdx] })
            ).reduce(combine)
        }
    }
