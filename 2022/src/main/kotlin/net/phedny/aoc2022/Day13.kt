package net.phedny.aoc2022

sealed interface IntOrList : Comparable<IntOrList> {}

data class IolInt(val value: Int) : IntOrList {
    override fun compareTo(other: IntOrList): Int =
        when (other) {
            is IolInt -> value.compareTo(other.value)
            is IolList -> IolList(listOf(this)).compareTo(other)
        }
}

data class IolList(val list: List<IntOrList>) : IntOrList {
    override fun compareTo(other: IntOrList): Int =
        when (other) {
            is IolInt -> compareTo(IolList(listOf(other)))
            is IolList -> when {
                list == other.list -> 0
                list.isEmpty() -> -1
                other.list.isEmpty() -> 1
                else -> list[0].compareTo(other.list[0]).let {
                    if (it == 0) IolList(list.drop(1)).compareTo(IolList(other.list.drop(1))) else it
                }
            }
        }
}

fun main() {
    val input = readInput().split("\n")
        .filter(String::isNotEmpty)
        .map(::parseIntOrList)
        .map(Pair<CharSequence, IntOrList>::second)

    input
        .chunked(2)
        .mapIndexed { i, (left, right) -> if (left < right) i + 1 else 0 }
        .sum()
        .also(::println)

    val divider1 = IolList(listOf(IolList(listOf(IolInt(2)))))
    val divider2 = IolList(listOf(IolList(listOf(IolInt(6)))))
    (input + divider1 + divider2)
        .sorted()
        .let { (it.indexOf(divider1) + 1) * (it.indexOf(divider2) + 1) }
        .also(::println)
}

fun parseIntOrList(line: CharSequence): Pair<CharSequence, IntOrList> =
    if (line[0] == '[') parseList(line.drop(1)) else parseInt(line)

fun parseInt(line: CharSequence, value: Int = 0): Pair<CharSequence, IolInt> =
    if (line[0].isDigit()) parseInt(line.drop(1), 10 * value + line[0].digitToInt()) else Pair(line, IolInt(value))

fun parseList(line: CharSequence, list: List<IntOrList> = emptyList()): Pair<CharSequence, IolList> =
    when (line[0]) {
        ']' -> Pair(line.drop(1), IolList(list))
        ',' -> parseList(line.drop(1), list)
        else -> {
            val (rest, value) = parseIntOrList(line)
            parseList(rest, list + value)
        }
    }
