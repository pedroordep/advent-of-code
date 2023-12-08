import java.io.File

fun Pair<String, String>.goTo(to: Char): String {
    return when (to) {
        'L' -> {
            this.first
        }

        'R' -> {
            this.second
        }

        else -> {
            ""
        }
    }
}

var maze = mutableMapOf<String, Pair<String, String>>()
var instructions = ""

fun setMazeAndInstructions(lines: List<String>) {
    instructions = lines[0]
    maze = mutableMapOf()
    val regex = """(\w\w\w) = \((\w\w\w), (\w\w\w)\)""".toRegex()
    lines.drop(2).forEach {
        val (from, toLeft, toRight) = regex.find(it)!!.groupValues.drop(1)
        maze[from] = Pair(toLeft, toRight)
    }
}

fun findPath(from: String, to: String, step: Int): Int {
    if (from == to) {
        return step
    }
    // println("at $from, going ${instructions[step % instructions.length]} to ${maze[from]!!.goTo(instructions[step % instructions.length])}")
    return findPath(maze[from]!!.goTo(instructions[step % instructions.length]), to, step + 1)
}

fun part1(filename: String): Int {
    val lines = File(filename).readLines()
    setMazeAndInstructions(lines)
    return findPath("AAA", "ZZZ", 0)
}

println(part1("input.txt"))

fun String.isEnd(): Boolean {
    return this[2] == 'Z'
}

fun String.isStart(): Boolean {
    return this[2] == 'A'
}

// https://www.baeldung.com/kotlin/lcm#finding-the-lcm-of-two-numbers
fun findLCM(a: Long, b: Long): Long {
    val larger = if (a > b) a else b
    val maxLcm = a * b
    var lcm = larger
    while (lcm <= maxLcm) {
        if (lcm % a == 0L && lcm % b == 0L) {
            return lcm
        }
        lcm += larger
    }
    return maxLcm
}

// https://www.baeldung.com/kotlin/lcm#finding-the-lcm-in-a-list-of-numbers
fun findLCMOfListOfNumbers(numbers: List<Long>): Long {
    var result = numbers[0]
    for (i in 1 until numbers.size) {
        result = findLCM(result, numbers[i])
    }
    return result
}

fun part2(filename: String): Long {
    val lines = File(filename).readLines()
    setMazeAndInstructions(lines)

    val cur = maze.keys.filter { it.isStart() }.toMutableList()
    val endsCount = maze.keys.count { it.isEnd() }
    val occurrenceAtIndex = mutableMapOf<String, Int>()
    var index = 0
    while (endsCount != occurrenceAtIndex.keys.size) {
        cur.forEachIndexed { curIndex, it ->
            if (it.isEnd()) {
                occurrenceAtIndex[it] = index
            }
            cur[curIndex] = maze[it]!!.goTo(instructions[index % instructions.length])
        }
        index++
    }
    return findLCMOfListOfNumbers(occurrenceAtIndex.values.map { it.toLong() })
}

println(part2("input_example_3.txt"))
println(part2("input.txt"))