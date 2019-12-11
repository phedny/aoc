val range = /**/ to /**/

val passwords = for {
  p <- range.map(_.toString.toSeq)
  if (p zip p.tail).exists(cs => cs._1 == cs._2)
  if p.sorted == p
} yield p

println(passwords.length)

