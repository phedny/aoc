val input = StdIn.readLine()

type State = (Option[Int], Vector[Int])
val program = input.split(",").toVector.map(_.toInt)

def step: State => State = _ match {
  case (Some(pc), m) if m(pc) == 1  => (Some(pc + 4), m.updated(m(pc + 3), m(m(pc + 1)) + m(m(pc + 2))))
  case (Some(pc), m) if m(pc) == 2  => (Some(pc + 4), m.updated(m(pc + 3), m(m(pc + 1)) * m(m(pc + 2))))
  case (Some(pc), m) if m(pc) == 99 => (None, m)
}

def runProgram(noun: Int, verb: Int): Int = {
  lazy val stepStream: LazyList[State] = (Some(0), program.updated(1, noun).updated(2, verb)) #:: stepStream.map(step)
  stepStream.takeWhile(_ match { case (maybePc, _) => maybePc.nonEmpty }).last._2(0)
}

(for {
  noun <- 0 to 99
  verb <- 0 to 99
  if runProgram(noun, verb) == 19690720
} yield 100 * noun + verb) foreach println

