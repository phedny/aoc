package net.phedny.aoc2022

import kotlin.math.absoluteValue
import kotlin.math.sign

enum class Direction(val dx: Int, val dy: Int) {
    U(0, 1), D(0, -1), L(-1, 0), R(1, 0)
}

interface Node {
    val x: Int
    val y: Int
    fun step(dx: Int, dy: Int): Node
    fun finalNode(): Node
}

data class ExpanderNode(val expandTo: Node) : Node {
    override val x: Int
        get() = expandTo.x
    override val y: Int
        get() = expandTo.y
    override fun step(dx: Int, dy: Int): ExpanderNode =
        if (dx == 0 && dy == 0) {
            this
        } else {
            ExpanderNode(expandTo.step(dx.sign, dy.sign)).step(dx - dx.sign, dy - dy.sign)
        }
    override fun finalNode(): Node =
        expandTo.finalNode()
}

data class DamperNode(override val x: Int = 0, override val y: Int = 0, val dampTo: Node) : Node {
    override fun step(dx: Int, dy: Int): DamperNode =
        when {
            dx.absoluteValue > 1 || dy.absoluteValue > 1 ->
                throw IllegalArgumentException("can't step faster than 1")
            (x + dx - dampTo.x).absoluteValue > 1 || (y + dy - dampTo.y).absoluteValue > 1 ->
                DamperNode(x + dx, y + dy, dampTo.step((x + dx - dampTo.x).sign, (y + dy - dampTo.y).sign))
            else ->
                DamperNode(x + dx, y + dy, dampTo)
        }
    override fun finalNode(): Node =
        dampTo.finalNode()
}

data class TrackingNode(override val x: Int = 0, override val y: Int = 0, val visited: Set<Pair<Int, Int>> = emptySet()) : Node {
    override fun step(dx: Int, dy: Int): TrackingNode =
        TrackingNode(x + dx, y + dy, visited + Pair(x, y))
    override fun finalNode(): Node =
        this
}

fun main() {
    val input = readInput().split("\n")
        .map { it.split(" ").let { (dir, amount) -> Pair(Direction.valueOf(dir), amount.toInt()) } }

    listOf(1, 9).forEach { ropeCount ->
        val rope = ExpanderNode((0 until ropeCount).fold(TrackingNode(0, 0, emptySet())) { node: Node, _ -> DamperNode(0, 0, node) })
        input
            .fold(rope) { node, (dir, amount) -> node.step(amount * dir.dx, amount * dir.dy) }
            .let { it.finalNode() as TrackingNode }
            .also { println((it.visited + Pair(it.x, it.y)).size) }
    }
}
