package net.phedny.aoc2022

fun main() {
    listOf(4, 14).forEach { l ->
        readInput().windowed(l, 1)
            .indexOfFirst { it.toSet().size == l }
            .also { println(it + l) }
    }
}
