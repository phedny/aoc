package net.phedny.aoc2022

fun main() {
    val input = readInput().split("\n")
        .map { "(\\d+)-(\\d+),(\\d+)-(\\d+)".toRegex().matchEntire(it)!!.groupValues.drop(1) }
        .map { it.map(Integer::parseInt) }

    val fullOverlap = listOf(listOf(0, 2, 3, 1), listOf(2, 0, 1, 3))
    val partialOverlap = fullOverlap + listOf(listOf(0, 2, 1, 3), listOf(2, 0, 3, 1))

    println(input.count { pair -> fullOverlap.any { isAscending(pair, it) } })
    println(input.count { pair -> partialOverlap.any { isAscending(pair, it) } })
}

fun isAscending(list: List<Int>, order: List<Int>) = order.windowed(2, 1).all { list[it[0]] <= list[it[1]] }
