package net.phedny.aoc2022

fun main() {
    data class Accumulator(val workingDir: List<String>, val files: Map<String, Int>, val directories: Set<String>)

    val (_, files, directories) = readInput().split("\n")
        .map { "(?:\\$ )?(cd|ls|dir|\\d+)(?: (/|\\.{2}|[a-z.]+))?".toRegex().matchEntire(it)!!.groupValues }
        .fold(Accumulator(emptyList(), emptyMap(), setOf(""))) { acc, (_, cmd, arg) ->
            when (cmd) {
                "cd" -> acc.copy(workingDir = when (arg) {
                    "/" -> listOf("")
                    ".." -> acc.workingDir.dropLast(1)
                    else -> acc.workingDir + arg
                })
                "ls" -> acc
                "dir" -> acc.copy(directories = acc.directories + ((acc.workingDir + arg).joinToString("/")))
                else -> acc.copy(files = acc.files + ((acc.workingDir + arg).joinToString("/") to cmd.toInt()))
            }
        }

    val dSizes = directories
        .map { dName -> files.filterKeys { fName -> fName.startsWith("$dName/") }.values.sum() }

    println(dSizes.filter { it <= 100000 }.sum())

    val needToFree = dSizes.maxOrNull()!! - 40000000
    println(dSizes.filter { it > needToFree }.minOrNull())
}
