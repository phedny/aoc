package net.phedny.aoc2022

import kotlin.math.absoluteValue

data class Sensor(val sensorX: Int, val sensorY: Int, val beaconX: Int, val beaconY: Int) {
    private val distance = (sensorX - beaconX).absoluteValue + (sensorY - beaconY).absoluteValue
    fun rangeOfXAtY(y: Int): IntRange = (distance - (sensorY - y).absoluteValue).let { sensorX - it .. sensorX + it }
}

interface IntRanges {
    operator fun minus(range: IntRange): IntRanges

    companion object {
        fun createFrom(togglePoints: List<Int>): IntRanges =
            togglePoints.chunked(2)
                .filter { (a, b) -> a != b }
                .flatten()
                .let { if (it.isEmpty()) EmptyRanges else TogglePointsIntRanges(it) }
    }
}

fun IntRange.toIntRanges(): IntRanges = IntRanges.createFrom(listOf(this.first, this.last + 1))

object EmptyRanges : IntRanges {
    override fun minus(range: IntRange): IntRanges = this
}

data class TogglePointsIntRanges(val togglePoints: List<Int>) : IntRanges {
    override operator fun minus(range: IntRange): IntRanges {
        if (range.isEmpty()) {
            return this
        }

        val prefix = togglePoints.takeWhile { it < range.first }
        val postfix = togglePoints.dropWhile { it <= range.last }
        val togglePoints = listOf(
            prefix,
            if (prefix.size % 2 == 0) emptyList() else listOf(range.first),
            if (postfix.size % 2 == 0) emptyList() else listOf(range.last + 1),
            postfix,
        ).flatten()

        return IntRanges.createFrom(togglePoints)
    }
}

fun main() {
    val sensors = readInput().split("\n")
        .map { ".*x=(-?\\d+), y=(-?\\d+).*x=(-?\\d+), y=(-?\\d+).*".toRegex().matchEntire(it)!!.groupValues }
        .map { (_, sensorX, sensorY, beaconX, beaconY) -> Sensor(sensorX.toInt(), sensorY.toInt(), beaconX.toInt(), beaconY.toInt()) }

    val beacons = sensors.map { Pair(it.beaconX, it.beaconY) }.toSet()

    val depth = 2000000
    sensors.flatMap { it.rangeOfXAtY(depth) }
        .toSet()
        .also { println(it.size - beacons.count { (x, y) -> it.contains(x) && y == depth } ) }

    val ranges = (0 .. 2 * depth).toIntRanges()
    val result = sensors.fold(List(2 * depth + 1) { ranges }) { rs, sensor -> rs.mapIndexed { d, r -> r - sensor.rangeOfXAtY(d) } }
    val resultY = result.indexOfFirst { it is TogglePointsIntRanges }
    val resultX = (result[resultY] as TogglePointsIntRanges).togglePoints.first()
    println(4000000L * resultX + resultY)
}
