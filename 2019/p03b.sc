case class Point(val x: Int, val y: Int) {
  val distanceToOrigin = x.abs + y.abs
  def +(p: Point) = Point(x + p.x, y + p.y)
}

val increments: Map[String, Point] = Map("U" -> Point(0, 1), "D" -> Point(0, -1), "L" -> Point(-1, 0), "R" -> Point(1, 0))
val stepPattern = "(\\w)(\\d+)".r

def wirePoints(s: String) = {
  s.split(",").flatMap(_ match {
    case stepPattern(dir, dist) => List.fill(dist.toInt)(increments(dir))
  }).scanLeft(Point(0, 0))(_ + _)
}

val wire1 = wirePoints(StdIn.readLine())
val wire2 = wirePoints(StdIn.readLine())
println((wire1 intersect wire2).map(p => wire1.indexOf(p) + wire2.indexOf(p)).filter(_ > 0).min)

