import java.io.File

fun part1(filename: String): Long {
    val sections = File(filename).readText().split("\n\n")
    val seeds = sections[0].removePrefix("seeds: ").split(" ").map { it.toLong() }
    val transformations = seeds.toMutableList()
//    println(transformations)
    sections.drop(1).forEach {
        val hasTransformed = transformations.map { false }.toMutableList()
        it.split("\n").drop(1).forEach {
            val (to, from, range) = it.split(" ").map { it.toLong() }
            transformations.forEachIndexed { index, curTransformation ->
                if (!hasTransformed[index] && from <= curTransformation && curTransformation < from + range) {
                    hasTransformed[index] = true
                    transformations[index] = to + transformations[index] - from
                }
            }
        }
//        println(transformations)
    }
    return transformations.min()
}

println(part1("input_example.txt"))
println(part1("input.txt"))

fun part2(filename: String): Long {
    val sections = File(filename).readText().split("\n\n")
    val seeds = sections[0].removePrefix("seeds: ").split(" ").chunked(2).map { Pair(it[0].toLong(), it[0].toLong() + it[1].toLong() - 1) }
    val transformations = seeds.toMutableList()
//    println(transformations)
    sections.drop(1).forEach {
        val hasTransformed = transformations.map { false }.toMutableList()
        val transformationsToConsider = transformations.size
        it.split("\n").drop(1).forEach {
            val (to, start, range) = it.split(" ").map { it.toLong() }
            val end = start + range - 1
            val transformationValue = to - start
            for (index in 0..<transformationsToConsider) {
                val curTransformation = transformations[index]
                if (!hasTransformed[index] && curTransformation.first >= start && curTransformation.second <= end) {
                    hasTransformed[index] = true
                    transformations[index] = Pair(transformations[index].first + transformationValue, transformations[index].second + transformationValue)
//                    println("[$to $start $range] converted $curTransformation to ${transformations[index]}")
                } else if (!hasTransformed[index] && curTransformation.second > start && curTransformation.second < end) {
                    transformations.add(Pair(curTransformation.first, start - 1))
                    hasTransformed[index] = true
                    transformations[index] = Pair(start + transformationValue, curTransformation.second + transformationValue)
//                    println("[$to $start $range] converted $curTransformation to ${Pair(curTransformation.first, start - 1)} ${transformations[index]}")
                } else if (!hasTransformed[index] && curTransformation.first > start && curTransformation.first < end) {
                    transformations.add(Pair(curTransformation.first + transformationValue, to + range - 1))
                    hasTransformed[index] = true
                    transformations[index] = Pair(start + range, curTransformation.second)
//                    println("[$to $start $range] converted $curTransformation to ${Pair(curTransformation.first + transformationValue, to + range - 1)} ${transformations[index]}")
                } else if (!hasTransformed[index] && curTransformation.first < start && curTransformation.second > end) {
                    transformations.add(Pair(curTransformation.first, start - 1))
                    transformations.add(Pair(start + range, curTransformation.second))
                    hasTransformed[index] = true
                    transformations[index] = Pair(to, to + range - 1)
//                    println("[$to $start $range] converted $curTransformation to ${Pair(curTransformation.first, start - 1)} ${Pair(start + range, curTransformation.second)} ${transformations[index]}")
                }
            }
        }
//        println(transformations)
    }
    return transformations.minOf { it.first }
}

println(part2("input_example.txt"))
println(part2("input.txt"))