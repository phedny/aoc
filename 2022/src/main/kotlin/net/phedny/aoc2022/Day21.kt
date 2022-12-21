package net.phedny.aoc2022

sealed interface Monkey21Override {
    fun create(monkeys: Map<String, List<String>>, name: String, overrides: Map<String, Monkey21Override>): Monkey21
}

sealed interface Monkey21 {
    val value: Long?
    fun solve(result: Long): Long

    companion object {
        fun create(monkeys: Map<String, List<String>>, name: String, overrides: Map<String, Monkey21Override>): Monkey21 =
            overrides[name]?.create(monkeys, name, overrides)
                ?: monkeys[name]!!.let { (_, literal, left, operator, right) ->
                    when (operator) {
                        "+" -> AdditionMonkey(create(monkeys, left, overrides), create(monkeys, right, overrides))
                        "-" -> SubtractionMonkey(create(monkeys, left, overrides), create(monkeys, right, overrides))
                        "*" -> MultiplicationMonkey(create(monkeys, left, overrides), create(monkeys, right, overrides))
                        "/" -> DivisionMonkey(create(monkeys, left, overrides), create(monkeys, right, overrides))
                        else -> LiteralMonkey(literal.toLong())
                        }
                    }
    }
}

data class LiteralMonkey(override val value: Long) : Monkey21 {
    override fun solve(result: Long): Long = throw AssertionError()
}

data class AdditionMonkey(val left: Monkey21, val right: Monkey21) : Monkey21 {
    override val value: Long? = if (left.value == null || right.value == null) null else left.value!! + right.value!!
    override fun solve(result: Long): Long =
        when {
            left.value == null -> left.solve(result - right.value!!)
            right.value == null -> right.solve(result - left.value!!)
            else -> throw AssertionError()
        }
}

data class SubtractionMonkey(val left: Monkey21, val right: Monkey21) : Monkey21 {
    override val value: Long? = if (left.value == null || right.value == null) null else left.value!! - right.value!!
    override fun solve(result: Long): Long =
        when {
            left.value == null -> left.solve(right.value!! + result)
            right.value == null -> right.solve(left.value!! - result)
            else -> throw AssertionError()
        }
}

data class MultiplicationMonkey(val left: Monkey21, val right: Monkey21) : Monkey21 {
    override val value: Long? = if (left.value == null || right.value == null) null else left.value!! * right.value!!
    override fun solve(result: Long): Long =
        when {
            left.value == null -> left.solve(result / right.value!!)
            right.value == null -> right.solve(result / left.value!!)
            else -> throw AssertionError()
        }
}

data class DivisionMonkey(val left: Monkey21, val right: Monkey21) : Monkey21 {
    override val value: Long? = if (left.value == null || right.value == null) null else left.value!! / right.value!!
    override fun solve(result: Long): Long =
        when {
            left.value == null -> left.solve(right.value!! * result)
            right.value == null -> right.solve(left.value!! / result)
            else -> throw AssertionError()
        }
}

data class RootMonkey(val left: Monkey21, val right: Monkey21) : Monkey21 {
    override val value: Long? = null
    override fun solve(result: Long): Long =
        if (left.value == null) {
            left.solve(right.value!!)
        } else {
            right.solve(left.value!!)
        }

    companion object : Monkey21Override {
        override fun create(monkeys: Map<String, List<String>>, name: String, overrides: Map<String, Monkey21Override>): RootMonkey =
            monkeys[name]!!.let { (_, _, left, _, right) ->
                RootMonkey(Monkey21.create(monkeys, left, overrides), Monkey21.create(monkeys, right, overrides))
            }
    }
}

object HumanMonkey : Monkey21, Monkey21Override {
    override val value: Long? = null
    override fun solve(result: Long): Long = result
    override fun create(monkeys: Map<String, List<String>>, name: String, overrides: Map<String, Monkey21Override>): HumanMonkey =
        HumanMonkey
}

fun main() {
    val monkeys = readInput().split("\n")
        .map { "(\\w+): (?:(\\d+)|(\\w+) ([+\\-*/]) (\\w+))".toRegex().matchEntire(it)!!.groupValues.drop(1) }
        .associateBy(List<String>::first)

    println(Monkey21.create(monkeys, "root", emptyMap()).value)
    println(Monkey21.create(monkeys, "root", mapOf("humn" to HumanMonkey, "root" to RootMonkey)).solve(0))
}
