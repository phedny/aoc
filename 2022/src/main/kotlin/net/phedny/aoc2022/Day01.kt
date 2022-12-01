package net.phedny.aoc2022

fun main() {
    val input = readInput("real")
        .split("\n\n")
        .map { it.split("\n").map(Integer::parseInt).sum() }
        .sortedDescending()

    println(input.take(1).sum())
    println(input.take(3).sum())
}
