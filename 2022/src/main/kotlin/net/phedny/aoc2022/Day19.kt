package net.phedny.aoc2022

import kotlin.math.ceil

enum class Resource { ORE, CLAY, OBSIDIAN, GEODE }

data class Blueprint(val id: Int, val robotCosts: Map<Resource, Map<Resource, Int>>) {
    val maxRobots = Resource.values().associateWith { resource -> if (resource == Resource.GEODE) Integer.MAX_VALUE else robotCosts.maxOf { (_, costs) -> costs[resource] ?: 0 } }
}

fun main() {
    val blueprints = readInput().split("\n")
        .map { "\\d+".toRegex().findAll(it).map(MatchResult::value).map(Integer::parseInt).toList() }
        .map { Blueprint(it.first(), mapOf(
            Resource.ORE to mapOf(Resource.ORE to it[1]),
            Resource.CLAY to mapOf(Resource.ORE to it[2]),
            Resource.OBSIDIAN to mapOf(Resource.ORE to it[3], Resource.CLAY to it[4]),
            Resource.GEODE to mapOf(Resource.ORE to it[5], Resource.OBSIDIAN to it[6])
        )) }

    println(blueprints.sumOf { it.id * findMaxGeodes(it, 24) })
    println(blueprints.take(3).fold(1) { acc, it -> acc * findMaxGeodes(it, 32) })
}

fun findMaxGeodes(blueprint: Blueprint, minutesRemaining: Int, resources: Map<Resource, Int> = Resource.values().associateWith { 0 }, robots: Map<Resource, Int> = Resource.values().associateWith { if (it == Resource.ORE) 1 else 0 }, ifAtLeast: Int = 0): Int =
    if (resources.getValue(Resource.GEODE) + minutesRemaining * robots.getValue(Resource.GEODE) + (minutesRemaining * (minutesRemaining - 1) / 2) <= ifAtLeast) {
        ifAtLeast
    } else {
        Resource.values()
            .filter { blueprint.maxRobots.getValue(it) > robots.getValue(it) }
            .filter { blueprint.robotCosts.getValue(it).all { (resource) -> robots.getValue(resource) > 0 } }
            .associateWith { blueprint.robotCosts.getValue(it).maxOf { (resource, amount) -> ceil(1.0 * (amount - resources.getValue(resource)) / robots.getValue(resource)).toInt() } }
            .filterValues { it in 0 until minutesRemaining - 1 }
            .entries
            .fold(ifAtLeast.coerceAtLeast(resources.getValue(Resource.GEODE) + minutesRemaining * robots.getValue(Resource.GEODE))) { acc, (newRobot, collectionTime) -> findMaxGeodes(
                blueprint,
                minutesRemaining - collectionTime - 1,
                resources.mapValues { (resource, amount) -> amount + (collectionTime + 1) * robots.getValue(resource) - (blueprint.robotCosts.getValue(newRobot)[resource] ?: 0) },
                robots + (newRobot to robots.getValue(newRobot) + 1),
                acc
            ).coerceAtLeast(acc) }
    }
