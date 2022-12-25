package net.phedny.aoc2022

fun main() {
    readInput().split("\n")
        .sumOf { it.fromSnafu() }
        .also { println(it.toSnafu()) }
}

fun String.fromSnafu(): Long =
    this.fold(0L) { acc, c ->
        5 * acc + when (c) {
            '=' -> -2L
            '-' -> -1L
            else -> c.digitToInt().toLong()
        }
    }

fun Long.toSnafu(): String =
    if (this == 0L) {
        ""
    } else {
        when (this % 5) {
            0L, 1L, 2L -> (this / 5).toSnafu() + (this % 5)
            3L -> (this / 5 + 1).toSnafu() + '='
            4L -> (this / 5 + 1).toSnafu() + '-'
            else -> throw AssertionError()
        }
    }
