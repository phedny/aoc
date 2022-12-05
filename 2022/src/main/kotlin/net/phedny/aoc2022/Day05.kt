package net.phedny.aoc2022

fun main() {
    val (stackInput, moves) = readInput().split("\n\n")
        .map { it.split("\n") }

    val initialStacks = stackInput.reversed().drop(1)
        .map { it.chunked(4).map { chunk -> chunk[1] } }
        .let { (0 until (stackInput.last().length + 2) / 4)
            .map { stackIdx -> it
                .mapNotNull { row -> row.getOrNull(stackIdx) }
                .filterNot { it == ' ' } } }

    moves
        .map { "move (\\d+) from (\\d+) to (\\d+)".toRegex().matchEntire(it)!!.groupValues.drop(1) }
        .map { it.map(Integer::parseInt) }
        .map { (count, from, to) -> Triple(count, from - 1, to - 1) }
        .also { println(doFold(initialStacks, it, true).map(List<Char>::last).joinToString("")) }
        .also { println(doFold(initialStacks, it, false).map(List<Char>::last).joinToString("")) }
}

fun doFold(initialStacks: List<List<Char>>, moves: List<Triple<Int, Int, Int>>, reverse: Boolean): List<List<Char>> =
    moves.fold(initialStacks) { stacks, (count, from, to) ->
        stacks
            .mapIndexed { stackIdx, stack ->
                when (stackIdx) {
                    from -> stack.dropLast(count)
                    to -> stack + stacks[from].takeLast(count).let { if (reverse) it.reversed() else it }
                    else -> stack
                }
            }
    }
