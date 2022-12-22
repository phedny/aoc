package net.phedny.aoc2022

enum class Orientation(val step: Pair<Int, Int>, val transform: (Pair<Int, Int>, Int) -> Pair<Int, Int>) {
    UP(Pair(-1, 0), { position, _ -> position }),
    RIGHT(Pair(0, 1), { (x, y), edgeLength -> Pair(y, edgeLength - x - 1) }),
    DOWN(Pair(1, 0), { (x, y), edgeLength -> Pair(edgeLength - x - 1, edgeLength - y - 1) }),
    LEFT(Pair(0, -1), { (x, y), edgeLength -> Pair(edgeLength - y - 1, x)});

    fun transform(orientation: Orientation): Orientation =
        Orientation.values()[(orientation.ordinal + ordinal) % 4]
}

val nextFaces: Map<Orientation, List<Int>> = mapOf(
    Orientation.RIGHT to listOf(2, 0, 1, 5, 3, 4),
    Orientation.DOWN to listOf(1, 2, 0, 4, 5, 3),
    Orientation.LEFT to listOf(3, 5, 4, 0, 2, 1),
    Orientation.UP to listOf(4, 3, 5, 1, 0, 2)
)

val nextOrientation: Map<Orientation, Orientation> = mapOf(
    Orientation.RIGHT to Orientation.RIGHT,
    Orientation.DOWN to Orientation.LEFT,
    Orientation.LEFT to Orientation.DOWN,
    Orientation.UP to Orientation.DOWN
)

data class Face(val base: Pair<Int, Int>, val orientation: Orientation)

sealed interface FaceMapper {
    fun toPlane(position: Triple<Int, Int, Int>): Pair<Int, Int>
    fun toPlane(position: Triple<Int, Int, Int>, orientation: Orientation): Orientation
    fun fix(position: Triple<Int, Int, Int>, orientation: Orientation): Pair<Triple<Int, Int, Int>, Orientation>
}

class FlatFaceMapper(private val edgeLength: Int, private val faces: List<Face>) : FaceMapper {
    override fun toPlane(position: Triple<Int, Int, Int>): Pair<Int, Int> =
        faces[position.first].base.let {
            Pair(position.second + edgeLength * it.first, position.third + edgeLength * it.second)
        }

    override fun toPlane(position: Triple<Int, Int, Int>, orientation: Orientation): Orientation =
        orientation

    override fun fix(position: Triple<Int, Int, Int>, orientation: Orientation): Pair<Triple<Int, Int, Int>, Orientation> =
       when {
            position.second < 0 -> {
                val base = faces[position.first].base
                faces.indexOfFirst { it.base == Pair(base.first - 1, base.second) }.let { faceOneUp ->
                    if (faceOneUp == -1) {
                        val baseFirst = faces.filter { it.base.second == base.second }.maxOf { it.base.first }
                        val faceAllDown = faces.indexOfFirst { it.base == Pair(baseFirst, base.second) }
                        fix(Triple(faceAllDown, position.second + edgeLength, position.third), orientation)
                    } else {
                        fix(Triple(faceOneUp, position.second + edgeLength, position.third), orientation)
                    }
                }
            }
            position.second >= edgeLength -> {
                val base = faces[position.first].base
                faces.indexOfFirst { it.base == Pair(base.first + 1, base.second) }.let { faceOneDown ->
                    if (faceOneDown == -1) {
                        val baseFirst = faces.filter { it.base.second == base.second }.minOf { it.base.first }
                        val faceAllUp = faces.indexOfFirst { it.base == Pair(baseFirst, base.second) }
                        fix(Triple(faceAllUp, position.second - edgeLength, position.third), orientation)
                    } else {
                        fix(Triple(faceOneDown, position.second - edgeLength, position.third), orientation)
                    }
                }
            }
            position.third < 0 -> {
                val base = faces[position.first].base
                faces.indexOfFirst { it.base == Pair(base.first, base.second - 1) }.let { faceOneLeft ->
                    if (faceOneLeft == -1) {
                        val baseSecond = faces.filter { it.base.first == base.first }.maxOf { it.base.second }
                        val faceAllRight = faces.indexOfFirst { it.base == Pair(base.first, baseSecond) }
                        fix(Triple(faceAllRight, position.second, position.third + edgeLength), orientation)
                    } else {
                        fix(Triple(faceOneLeft, position.second, position.third + edgeLength), orientation)
                    }
                }
            }
            position.third >= edgeLength -> {
                val base = faces[position.first].base
                faces.indexOfFirst { it.base == Pair(base.first, base.second + 1) }.let { faceOneRight ->
                    if (faceOneRight == -1) {
                        val baseSecond = faces.filter { it.base.first == base.first }.minOf { it.base.second }
                        val faceAllLeft = faces.indexOfFirst { it.base == Pair(base.first, baseSecond) }
                        fix(Triple(faceAllLeft, position.second, position.third - edgeLength), orientation)
                    } else {
                        fix(Triple(faceOneRight, position.second, position.third - edgeLength), orientation)
                    }
                }
            }
            else -> Pair(position, orientation)
        }
}

class DiceFaceMapper(private val edgeLength: Int, private val faces: List<Face>) : FaceMapper {
    override fun toPlane(position: Triple<Int, Int, Int>): Pair<Int, Int> =
        faces[position.first].let { face ->
            face.orientation.transform(Pair(position.second, position.third), edgeLength).let {
                Pair(it.first + edgeLength * face.base.first, it.second + edgeLength * face.base.second)
            }
        }

    override fun toPlane(position: Triple<Int, Int, Int>, orientation: Orientation): Orientation =
        faces[position.first].orientation.transform(orientation)

    override fun fix(position: Triple<Int, Int, Int>, orientation: Orientation): Pair<Triple<Int, Int, Int>, Orientation> =
        when {
            position.second < 0 ->
                Orientation.DOWN.transform(Pair(position.second + edgeLength, position.third), edgeLength).let {
                    Pair(
                        Triple(nextFaces.getValue(Orientation.UP)[position.first], it.first, it.second),
                        Orientation.DOWN.transform(orientation)
                    )
                }
            position.second >= edgeLength ->
                Orientation.RIGHT.transform(Pair(position.second - edgeLength, position.third), edgeLength).let {
                    Pair(
                        Triple(nextFaces.getValue(Orientation.DOWN)[position.first], it.first, it.second),
                        Orientation.RIGHT.transform(orientation)
                    )
                }
            position.third < 0 ->
                Orientation.DOWN.transform(Pair(position.second, position.third + edgeLength), edgeLength).let {
                    Pair(
                        Triple(nextFaces.getValue(Orientation.LEFT)[position.first], it.first, it.second),
                        Orientation.DOWN.transform(orientation)
                    )
                }
            position.third >= edgeLength ->
                Orientation.LEFT.transform(Pair(position.second, position.third - edgeLength), edgeLength).let {
                    Pair(
                        Triple(nextFaces.getValue(Orientation.RIGHT)[position.first], it.first, it.second),
                        Orientation.LEFT.transform(orientation)
                    )
                }
            else -> Pair(position, orientation)
        }
}

fun main() {
    val (map, route) = readInput().split("\n")
        .let { it.chunked(it.size - 2) }

    val instructions = sequence {
        var n = 0
        route[1].forEach {
            n = when (it) {
                'L' -> { yield(Pair(n, Orientation.LEFT)); 0 }
                'R' -> { yield(Pair(n, Orientation.RIGHT)); 0 }
                else -> 10 * n + it.digitToInt()
            }
        }
        yield(Pair(n, null))
    }.toList()

    val (edgeLength, faces) = findFaces(map)
    println(execute(FlatFaceMapper(edgeLength, faces), map, instructions))
    println(execute(DiceFaceMapper(edgeLength, faces), map, instructions))
}

fun findFaces(map: List<String>): Pair<Int, List<Face>> {
    val edgeLength = map.size / if (map.size > map.first().length) 4 else 3
    val present = map.map { it.chunked(edgeLength).map(String::isNotBlank) }
        .chunked(edgeLength).map(List<List<Boolean>>::first)
        .flatMapIndexed { row, it -> it.mapIndexed { col, b -> if (b) Pair(row, col) else null } }
        .filterNotNull()

    val face0 = Face(present.filter { it.first == 0 }.minByOrNull(Pair<Int, Int>::second)!!, Orientation.UP)
    val faces = generateSequence(Pair(listOf(face0, null, null, null, null, null), mapOf(0 to face0))) { (faces, front) ->
        val newFaces = front.flatMap { (faceId, face) ->
            Orientation.values().map { orientation -> Pair(
                nextFaces.getValue(orientation)[faceId],
                Face(
                    orientation.transform(face.orientation).let { Pair(face.base.first + it.step.first, face.base.second + it.step.second) },
                    nextOrientation.getValue(orientation).transform(face.orientation)
                )
            ) } }
            .filter { present.contains(it.second.base) }
            .filter { faces[it.first] == null }
            .toMap()

        Pair(faces.mapIndexed { faceId, face -> if (newFaces.containsKey(faceId)) newFaces[faceId] else face }, newFaces)
    }
        .map(Pair<List<Face?>, Map<Int, Face>>::first)
        .first { it.none { b -> b == null } }
        .filterNotNull()

    return Pair(edgeLength, faces)
}

fun execute(mapper: FaceMapper, map: List<String>, instructions: List<Pair<Int, Orientation?>>): Int =
    instructions.fold(Pair(Triple(0, 0, 0), Orientation.RIGHT)) { posAndOr, (steps, turn) ->
        generateSequence(posAndOr) { (position, orientation) ->
            Triple(position.first, position.second + orientation.step.first, position.third + orientation.step.second)
                .let { mapper.fix(it, orientation) }
        }
            .take(steps + 1)
            .takeWhile { (position) -> mapper.toPlane(position).let { map[it.first][it.second] } == '.' }
            .last()
            .let { (position, orientation) -> Pair(position, orientation.transform(turn ?: Orientation.UP)) }
    }.let { (position, orientation) ->
        mapper.toPlane(position).let { 1000 * it.first + 4 * it.second } + (mapper.toPlane(position, orientation).ordinal + 3) % 4 + 1004
    }
