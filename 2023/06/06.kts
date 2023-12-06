import java.io.File

fun calculateRace(time: Long, distance: Long): Long {
    return (0..time).sumOf { it -> if (it * (time - it) > distance) 1L else 0 }
}

fun part1(filename: String): Long {
    val lines = File(filename).readLines()
    val times = lines[0].removePrefix("Time:").trim().split("""\W+""".toRegex()).map { it.toLong() }
    val distances = lines[1].removePrefix("Distance:").trim().split("""\W+""".toRegex()).map { it.toLong() }
    var sum = 1L
    times.indices.forEach { index ->
        sum *= calculateRace(times[index], distances[index])
    }
    return sum
}

assert(part1("input_example.txt") == 288L)
println(part1("input.txt"))

fun part2(filename: String): Long {
    val lines = File(filename).readLines()
    val time = lines[0].removePrefix("Time:").replace(" ", "").toLong()
    val distance = lines[1].removePrefix("Distance:").replace(" ", "").toLong()
    return calculateRace(time, distance)
}

assert(part2("input_example.txt") == 71503L)
println(part2("input.txt"))