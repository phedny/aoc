trait State {
  def mem: Vector[Int]
  def output: Vector[Int]
}
case class HaltedState(val mem: Vector[Int], val output: Vector[Int]) extends State {}
case class RunnableState(val ip: Int, val mem: Vector[Int], val input: Vector[Int], val output: Vector[Int]) extends State {
  val instruction = mem(ip) % 100
  def pMode(i: Int) = (mem(ip) / (math.pow(10, i + 1)) % 10).toInt
  def pRead(i: Int) = pMode(i) match {
    case 0 => mem(mem(ip + i))
    case 1 => mem(ip + i)
  }
  def pWrite(i: Int, v: Int) = pMode(i) match {
    case 0 => mem.updated(mem(ip + i), v)
  }
}

val program = StdIn.readLine().split(",").toVector.map(_.toInt)

def step: State => State = _ match {
  case s @ RunnableState(ip, _, _, _) if s.instruction == 1  => s.copy(ip = ip + 4, mem = s.pWrite(3, s.pRead(1) + s.pRead(2)))
  case s @ RunnableState(ip, _, _, _) if s.instruction == 2  => s.copy(ip = ip + 4, mem = s.pWrite(3, s.pRead(1) * s.pRead(2)))
  case s @ RunnableState(ip, _, i, _) if s.instruction == 3  => s.copy(ip = ip + 2, mem = s.pWrite(1, i.head), input = i.tail)
  case s @ RunnableState(ip, _, _, o) if s.instruction == 4  => s.copy(ip = ip + 2, output = o :+ s.pRead(1))
  case s @ RunnableState(ip, _, _, _) if s.instruction == 5  => s.copy(ip = if (s.pRead(1) != 0) s.pRead(2) else ip + 3)
  case s @ RunnableState(ip, _, _, _) if s.instruction == 6  => s.copy(ip = if (s.pRead(1) == 0) s.pRead(2) else ip + 3)
  case s @ RunnableState(ip, _, _, _) if s.instruction == 7  => s.copy(ip = ip + 4, mem = s.pWrite(3, if (s.pRead(1) < s.pRead(2)) 1 else 0))
  case s @ RunnableState(ip, _, _, _) if s.instruction == 8  => s.copy(ip = ip + 4, mem = s.pWrite(3, if (s.pRead(1) == s.pRead(2)) 1 else 0))
  case s @ RunnableState(_, m, _, o) if s.instruction == 99 => HaltedState(m, o)
}

def runProgram(signal: Int, phase: Int): Int = {
  lazy val stepStream: LazyList[State] = RunnableState(0, program, Vector(phase, signal), Vector.empty) #:: stepStream.map(step)
  val finalState = stepStream.takeWhile(_.isInstanceOf[RunnableState]).last
  finalState.output(0)
}

println(List(4, 3, 2, 1, 0).permutations.map(_.foldLeft(0)(runProgram(_, _))).max)

