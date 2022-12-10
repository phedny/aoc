package net.phedny.aoc2022

import kotlin.math.absoluteValue

fun main() {
    val input = readInput().split("\n")

    val signal = sequence {
        var x = 1
        input.map { it.split(" ").drop(1) }
            .forEach {
                if (it.isEmpty()) {
                    yield(x)
                } else {
                    yield(x)
                    yield(x)
                    x += it[0].toInt()
                }
            }
    }.toList()

    listOf(20, 60, 100, 140, 180, 220)
        .sumOf { it * signal[it - 1] }
        .also(::println)

    signal.chunked(40)
        .map { line -> line.mapIndexed { col, sprite -> if ((col - sprite).absoluteValue < 2) '#' else ' ' }.joinToString() }
        .forEach(::println)
}
