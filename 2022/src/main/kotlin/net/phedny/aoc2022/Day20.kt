package net.phedny.aoc2022

data class Entry(val value: Long) {
    var prev: Entry = this
    var next: Entry = this
    var mod: Int = Int.MAX_VALUE

    val asSequence: Sequence<Entry>
        get() = generateSequence(this) { it.next }

    fun move() {
        if (value == 0L) {
            return
        }

        prev.next = next
        next.prev = prev

        prev = asSequence.drop((value % mod + mod).toInt()).first()
        next = prev.next

        prev.next = this
        next.prev = this
    }
}

fun main() {
    val input = readInput().split("\n")
        .map(Integer::parseInt)

    println(mixWith(input, 1, 1))
    println(mixWith(input, 811589153, 10))
}

fun mixWith(input: List<Int>, decryptionKey: Long, rounds: Int): Long {
    val entries = input.map { Entry(it * decryptionKey) }

    entries.windowed(3, 1)
        .forEach { (prev, it, next) ->
            it.prev = prev
            it.next = next
            it.mod = input.size - 1
        }
    entries.first().let {
        it.prev = entries.last()
        it.next = entries[1]
        it.mod = input.size - 1
    }
    entries.last().let {
        it.prev = entries[entries.size - 2]
        it.next = entries.first()
        it.mod = input.size - 1
    }

    repeat(rounds) { entries.forEach(Entry::move) }

    return entries.first()
        .asSequence
        .dropWhile { it.value != 0L }
        .filterIndexed { index, _ -> index % 1000 == 0 }
        .drop(1)
        .take(3)
        .sumOf(Entry::value)
}
