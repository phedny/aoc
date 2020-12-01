val input = StdIn.readLine()

type State = (Option[Int], Vector[Int])
val init: State = (Some(0), input.split(",").toVector.map(_.toInt).updated(1, 12).updated(2, 2))

def step: State => State = _ match {
  case (Some(pc), m) if m(pc) == 1  => (Some(pc + 4), m.updated(m(pc + 3), m(m(pc + 1)) + m(m(pc + 2))))
  case (Some(pc), m) if m(pc) == 2  => (Some(pc + 4), m.updated(m(pc + 3), m(m(pc + 1)) * m(m(pc + 2))))
  case (Some(pc), m) if m(pc) == 99 => (None, m)
}

lazy val stepStream: LazyList[State] = init #:: stepStream.map(step)
val steps = stepStream.takeWhile(_ match { case (maybePc, _) => maybePc.nonEmpty }).toList

println(steps.last._2(0))

