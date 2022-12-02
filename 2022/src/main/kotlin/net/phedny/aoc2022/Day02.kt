package net.phedny.aoc2022

fun main() {
    val input = readInput()
        .split("\n")
        .map { Pair(it[0] - 'A', it[2] - 'X') }

    println(input.sumOf { (theirs, my) -> 1 + my + (4 + my - theirs) % 3 * 3 })
    println(input.sumOf { (theirs, result) -> 1 + 3 * result + (theirs + result + 2) % 3 })
}
