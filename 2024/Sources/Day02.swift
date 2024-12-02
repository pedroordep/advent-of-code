import Algorithms

struct Day02: AdventDay {
  // Save your data in a corresponding text file in the `Data` directory.
  var data: String

  // Splits input data into its component parts and convert from string.
  var entities: [[Int]] {
    data.split(separator: "\n").map {
      $0.split(separator: " ").compactMap { Int($0) }
    }
  }
  
  func checkSafety(report: [Int], most: Int, order: Int) -> Bool {
    for i in 1..<report.count {
      let dist = report[i - 1] - report[i]
      if report[i - 1] * order >= report[i] * order || abs(dist) > most {
//        print("\(report) failed: \(report[i-1]) \(report[i]), dist=\(dist)")
        return false
      }
    }
    
//    print("Safety check passed for \(report)")
    return true
  }
  
  func checkSafetyWithFailover(report: [Int], most: Int, order: Int) -> Bool {
    for i in 1..<report.count {
      let dist = report[i - 1] - report[i]
      if report[i - 1] * order >= report[i] * order || abs(dist) > most {
        var a = report, b = report
        a.remove(at: i-1)
        b.remove(at: i)
        
//        print("\(report) failed: \(report[i-1]) \(report[i]), dist=\(dist), testing with \(a) and \(b)")
        return checkSafety(report: a, most: most, order: 1) || checkSafety(report: a, most: most, order: -1) || checkSafety(report: b, most: most, order: 1) || checkSafety(report: b, most: most, order: -1)
      }
    }
    
//    print("Safety check passed for \(report)")
    return true
  }

  // Replace this with your solution for the first part of the day's challenge.
  func part1() -> Any {
    var sum = 0
    
    entities.forEach { report in
      sum += checkSafety(report: report, most: 3, order: report[0] > report[1] ? -1 : 1) ? 1 : 0
    }
    
    return sum
  }

  // Replace this with your solution for the second part of the day's challenge.
  func part2() -> Any {
    var sum = 0
    
    entities.forEach { report in
      sum += checkSafetyWithFailover(report: report, most: 3, order: 1) || checkSafetyWithFailover(report: report, most: 3, order: -1) ? 1 : 0
    }
    
    return sum
  }
}
