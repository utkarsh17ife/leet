function flattenArray(arr) {
    let finalArray = [];
    arr.forEach(e => {
        if (Array.isArray(e)) {
            finalArray = finalArray.concat(flattenArray(e))
        } else {
            finalArray.push(e)
        }
    });
    return finalArray;
}

function main() {
    const input = [1, 2, 3, [4, [5], [6, [7, [8, 9]]]], 10]
    const result = flattenArray(input);
    console.log(`Flattened Array:`, result.join(','))
}

main()