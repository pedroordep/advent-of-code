import java.io.File
import kotlin.math.pow

fun <K> increment(map: MutableMap<K, Int>, key: K) {
    when (val count = map[key]) {
        null -> map[key] = 1
        else -> map[key] = count + 1
    }
}

val cardScore = mapOf<Char, Long>('2' to 2, '3' to 3, '4' to 4, '5' to 5, '6' to 6, '7' to 7, '8' to 8, '9' to 9, 'T' to 10, 'J' to 11, 'Q' to 12, 'K' to 13, 'A' to 14)
val cardJokerScore = mapOf<Char, Long>('2' to 2, '3' to 3, '4' to 4, '5' to 5, '6' to 6, '7' to 7, '8' to 8, '9' to 9, 'T' to 10, 'J' to 1, 'Q' to 12, 'K' to 13, 'A' to 14)

data class Play(val hand: String, val bid: Int, val withJoker: Boolean = false) {
    var score: Long = 0

    private fun getScoreForLetter(letter: Char, at: Int): Long {
        val scoreMap = if (withJoker) {
            cardJokerScore
        } else cardScore
        return scoreMap[letter]!!.times(10.toDouble().pow(8 - at * 2).toLong())
    }

    private fun getHandScore(): Long {
        val occurrences = mutableMapOf<Char, Int>()
        hand.forEach { increment(occurrences, it) }

        if (withJoker) {
            if (occurrences.getOrDefault('J', 0) == 5) {
                occurrences.remove('J')
                occurrences['A'] = 5
            } else if (occurrences.getOrDefault('J', 0) > 0) {
                val numJokers = occurrences['J']!!
                occurrences.remove('J')
                val bestCard = occurrences.toList().sortedBy { -it.second * 100 + cardJokerScore[it.first]!! }[0].first
                occurrences[bestCard] = occurrences[bestCard]!! + numJokers
            }
        }

        val type = if (occurrences.size == 1) {
            7
        } else if (occurrences.size == 2 && occurrences.values.any { it == 4 }) {
            6
        } else if (occurrences.size == 2 && occurrences.values.any { it == 3 }) {
            5
        } else if (occurrences.size == 3 && occurrences.values.any { it == 3 }) {
            4
        } else if (occurrences.size == 3) {
            3
        } else if (occurrences.size == 4) {
            2
        } else {
            1
        }
        return type.times(10.toDouble().pow(10).toLong())
    }

    init {
        score = getHandScore() + hand.foldIndexed(0L) { index, acc, it -> acc + getScoreForLetter(it, index) }
    }
}

fun part1(filename: String): Int {
    return File(filename).readLines().map { Play(it.split(" ")[0], it.split(" ")[1].toInt()) }.sortedBy { it.score }.foldIndexed(0) { index, acc, it -> acc + (it.bid * (index + 1)) }
}

fun part2(filename: String): Int {
//    File(filename).readLines().map { Play(it.split(" ")[0], it.split(" ")[1].toInt(), true) }.sortedBy { it.score }.map { println("${it.hand} ${it.score}") }
    return File(filename).readLines().map { Play(it.split(" ")[0], it.split(" ")[1].toInt(), true) }.sortedBy { it.score }.foldIndexed(0) { index, acc, it -> acc + (it.bid * (index + 1)) }
}

assert(6440 == part1("input_example.txt"))
println("part 1: ${part1("input.txt")}")

assert(5905 == part2("input_example.txt"))
println("part 2: ${part2("input.txt")}")
