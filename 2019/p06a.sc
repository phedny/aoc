val orbitPattern = "(\\w+)\\)(\\w+)".r

// Didn't want to read from stdin, just paste the input as multi-line string
val objects = """""".split("\n").map(_ match {
  case orbitPattern(o1, o2) => (o1, o2)
}).groupBy(_._1).view.mapValues(_.toList.map(_._2))

def countOrbits(d: Int, o: String): Int = {
  objects.get(o).toList.flatMap(_.map(oo => d + countOrbits(d + 1, oo))).sum
}

println(countOrbits(1, "COM"))
