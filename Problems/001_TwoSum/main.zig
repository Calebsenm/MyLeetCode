const std = @import("std");

fn twoSum(num: []const u8, target: u8) [2]u8 {
    var indNum: [2]u8 = [2]u8{ 0, 0 };

    var i: u8 = 0;
    while (i < num.len) : (i += 1) {
        var j: u8 = i;

        while (j < num.len) : (j += 1) {
            if (i == j) continue;

            if (num[i] + num[j] == target) {
                indNum[0] = i;
                indNum[1] = j;
            }
        }
    }

    return indNum;
}

pub fn main() void {
    const numbers = [_]u8{ 2, 7, 11, 15, 1, 5, 6, 2 };
    const target = 9;
    const result = twoSum(numbers[0..numbers.len], target);
    std.debug.print("{any}", .{result});
}
