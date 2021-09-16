function reverseWave(arr) {
    const length = arr.length;
    const depth = arr[0].length;
    const reverseWave = [];
    let direction = "down";
    for (let j = length - 1; j > -1; j--) {
        if (direction == "up") {
            for (let i = depth - 1; i > -1; i--) {
                reverseWave.push(arr[i][j]);
            }
        }
        if (direction == "down") {
            for (let i = 0; i < depth; i++) {
                reverseWave.push(arr[i][j]);
            }
        }
        direction = direction == "up" ? "down" : "up";
    }
    return reverseWave;
}

function main() {
    const input = [[1, 2, 3, 4], [5, 6, 7, 8], [9, 10, 11, 12], [13, 14, 15, 16]]
    const result = reverseWave(input);
    console.log(result.join(','))
}

main()