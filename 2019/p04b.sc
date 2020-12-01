val range = /**/ to /**/

val digitGroupPattern = "(\\d)(?=(?!\\1))(\\d)\\2(?!\\2)|^(\\d)\\3(?!\\3)".r

val passwords = for {
  p <- range.map(_.toString.toSeq)
  if digitGroupPattern.findFirstIn(p).nonEmpty
  if p.sorted == p
} yield p

println(passwords.length)

