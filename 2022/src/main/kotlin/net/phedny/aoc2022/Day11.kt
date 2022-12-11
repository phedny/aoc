package net.phedny.aoc2022

data class Monkey(val id: Int, val items: List<Long>, val operation: (Long) -> Long, val test: (Long) -> Boolean, val trueTo: Int, val falseTo: Int, val inspections: Long) {
    fun catch(item: Long): Monkey = copy(items = items + item)
    fun throwAll(): Monkey = copy(items = emptyList(), inspections = inspections + items.size)
    fun distribute(mod: Int, relieve: Boolean): List<Pair<Long, Int>> = items.map { if (relieve) operation(it) / 3 else operation(it) }.map { Pair(it % mod, if (test(it)) trueTo else falseTo) }
}

fun main() {
    val (monkeys, divs) = readInput().split("\n\n")
        .map { "Monkey .*(\\d):.*items: ([\\d, ]+).*new = old ([+*]) (\\d+|old).*by (\\d+).*monkey (\\d+).*monkey (\\d+)".toRegex(RegexOption.DOT_MATCHES_ALL).matchEntire(it)!!.groupValues }
        .map { groups ->
            Pair(Monkey(
                groups[1].toInt(),
                groups[2].split(", ").map { it.toLong() },
                operation(groups[3], if (groups[4] == "old") null else groups[4].toLong()),
                divisibleBy(groups[5].toLong()),
                groups[6].toInt(),
                groups[7].toInt(),
                0
            ), groups[5].toInt()) }
        .unzip()

    val mod = divs.reduce(Int::times)
    listOf(Pair(20, true), Pair(10000, false)).forEach {(rounds, relieve) ->
        generateSequence(monkeys) { doRound(it, mod, relieve) }
            .drop(rounds)
            .first()
            .map(Monkey::inspections)
            .sortedDescending()
            .also { println(it[0] * it[1]) }
    }
}

fun operation(operator: String, operand: Long?): (Long) -> Long =
    when (operator) {
        "+" -> fun(it): Long = it + (operand ?: it)
        "*" -> fun(it): Long = it * (operand ?: it)
        else -> throw IllegalArgumentException("unsupported operator")
    }

fun divisibleBy(divBy: Long): (Long) -> Boolean =
    fun(it): Boolean = (it % divBy) == 0L

fun doRound(monkeys: List<Monkey>, mod: Int, relieve: Boolean): List<Monkey> =
    monkeys.indices.fold(monkeys) { monkeyList, current ->
        val thrown = monkeyList[current].distribute(mod, relieve)
        monkeyList.mapIndexed { index, monkey ->
            if (index == current) {
                monkey.throwAll()
            } else {
                thrown
                    .filter { it.second == index }
                    .fold(monkey) { it, item -> it.catch(item.first) }
            }
        }
    }
