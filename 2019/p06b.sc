val orbitPattern = "(\\w+)\\)(\\w+)".r

// Didn't want to read from stdin, just paste the input as multi-line string
val objects = """""".split("\n").map(_ match {
  case orbitPattern(o1, o2) => (o1, o2)
}).groupBy(_._1).view.mapValues(_.toList.map(_._2))

def pathToObject(from: String, to: String): List[String] = {
  if (from == to) {
    from :: Nil
  } else {
    val path = objects.get(from).toList.flatten.flatMap(pathToObject(_, to))
    if (path.isEmpty) {
      Nil
    } else {
      from :: path
    }
  }
}

def distance(l1: List[String], l2: List[String]): Int = (l1, l2) match {
  case (x :: xs, y :: ys) if x == y => distance(xs, ys)
  case (xs, ys)                     => xs.length + ys.length
}

println(distance(pathToObject("COM", "YOU"), pathToObject("COM", "SAN")) - 2)
