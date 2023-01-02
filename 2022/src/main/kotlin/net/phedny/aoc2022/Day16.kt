package net.phedny.aoc2022

data class Valve(val id: String, val flowRate: Int, val tunnelsTo: Set<String>)

data class Network(val valves: Map<String, Valve>) {
    val distances = valves.keys.associateWith { from ->
        generateSequence(Triple(mapOf(from to 0), listOf(from), 1)) { (distances, front, distance) ->
            front.flatMap { valves.getValue(it).tunnelsTo }
                .filter { !distances.containsKey(it) }
                .let { Triple(distances + it.map { d -> d to distance }, it, distance + 1) }
        }
            .first { it.second.isEmpty() }
            .first
    }
}

fun main() {
    val network = readInput().split("\n")
        .map { "Valve (\\w+) has flow rate=(\\d+);.* valves? ([\\w, ]+)".toRegex().matchEntire(it)!!.groupValues }
        .map { Valve(it[1], it[2].toInt(), it[3].split(", ").toSet()) }
        .associateBy(Valve::id)
        .let(::Network)

    val closeableValves = network.valves.values.filter { it.flowRate > 0 }.map(Valve::id).toSet()

    findSolutions(network, "AA", 30, 0, closeableValves)
        .let { it.maxOf(Pair<Set<String>, Int>::second) }
        .also(::println)

    findSolutions(network, "AA", 26, 0, closeableValves)
        .groupBy(Pair<Set<String>, Int>::first).map { (k, v) -> k to v.maxOf(Pair<Set<String>, Int>::second) }
        .let { it.maxOf { a -> it.filter { b -> a.first.none(b.first::contains) }.maxOf { b -> a.second + b.second } } }
        .also(::println)
}

fun findSolutions(network: Network, location: String, minutesRemaining: Int, pressureReleased: Int, closeableValves: Set<String>): List<Pair<Set<String>, Int>> =
    if (minutesRemaining <= 0 || closeableValves.isEmpty()) {
        listOf(emptySet<String>() to pressureReleased)
    } else {
        closeableValves.flatMap { valve ->
            val minutesAfterStep = minutesRemaining - network.distances.getValue(location).getValue(valve) - 1
            findSolutions(
                network,
                valve,
                minutesAfterStep,
                pressureReleased + minutesAfterStep * network.valves.getValue(valve).flowRate,
                closeableValves - valve
            ).map { (vs, p) -> vs + valve to p } + (emptySet<String>() to pressureReleased)
        }
    }
