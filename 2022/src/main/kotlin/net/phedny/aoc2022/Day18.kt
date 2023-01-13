package net.phedny.aoc2022

data class Cube(val x: Int, val y: Int, val z: Int) {
    fun neighbors() = sequenceOf(
        Cube(x - 1, y, z), Cube(x + 1, y, z),
        Cube(x, y - 1, z), Cube(x, y + 1, z),
        Cube(x, y, z - 1), Cube(x, y, z + 1)
    )
}

fun main() {
    val lava = readInput().split("\n")
        .map { it.split(",").map(Integer::parseInt) }
        .map { (x, y, z) -> Cube(x, y, z) }
        .toSet()

    val surface = lava.flatMap(Cube::neighbors) - lava
    println(surface.size)

    val surfaces = findSurfaces(lava, surface.toSet())
    val maxX = surface.maxOf(Cube::x)
    val exteriorSurface = surfaces.first { it.any { cube -> cube.x == maxX } }
    println(surface.count(exteriorSurface::contains))
}

fun findSurfaces(lava: Set<Cube>, surface: Set<Cube>): Set<Set<Cube>> =
    generateSequence(Pair(surface, emptySet<Set<Cube>>())) { (todo, surfaces) ->
        todo.firstOrNull()?.let {
            val collectedSurface = collectSurface(lava, todo, it)
            Pair(todo - collectedSurface, surfaces.plusElement(collectedSurface))
        }
    }
        .last()
        .second

fun collectSurface(lava: Set<Cube>, candidates: Set<Cube>, seed: Cube): Set<Cube> =
    generateSequence(setOf(seed)) { surface ->
        val matches = surface.flatMap(Cube::neighbors).toSet() - lava
        surface + candidates.filter { ((matches + surface) intersect it.neighbors().toSet()).isNotEmpty() }
    }
        .windowed(2, 1)
        .first { (a, b) -> a == b }
        .first()
