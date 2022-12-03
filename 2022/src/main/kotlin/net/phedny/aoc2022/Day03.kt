package net.phedny.aoc2022

fun main() {
  readInput().split("\n")
    .answerWith { it.map { r -> r.chunked(r.length / 2) } }
    .answerWith { it.chunked(3) }
}

fun List<String>.answerWith(transform: (List<String>) -> List<List<String>>): List<String> {
  transform(this)
    .sumOf { r -> r
      .map(String::toSet)
      .reduce(Iterable<Char>::intersect)
      .let { (it.first() - '&') % 58 }
    }
    .also(::println)
  return this
}
