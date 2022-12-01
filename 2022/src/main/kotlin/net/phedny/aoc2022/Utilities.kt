package net.phedny.aoc2022

fun readInput(fileName: String? = null): String {
    val fullClassName = Thread.currentThread()
        .stackTrace[if (fileName == null) 3 else 2]
        .className

    val input = ".*Day(\\d+)Kt".toRegex()
        .matchEntire(fullClassName)
        ?.let { object {}.javaClass.getResource("/${it.groupValues[1]}/${fileName ?: "real"}.txt") }
        ?.readText()

    return input!!
}
