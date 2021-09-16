function findWithKeyAndValueInObjArray(key, value, array) {
    let foundObj = null;
    for (let i = 0; i < array.length; i++) {
        let currObj = array[i];
        if (currObj[key] == value) {
            foundObj = currObj;
            break;
        }
    }
    return foundObj;
}

function mergeObjectArray(arr1, arr2) {
    const resultArr = [];
    arr1.forEach(e => {
        const objFrmSecondArr = findWithKeyAndValueInObjArray("id", e.id, arr2);
        resultArr.push({
            ...e, ...objFrmSecondArr
        })
    })
    return resultArr
}

function main() {
    const arr1 = [{ id: 1, name: "akshay" }, { id: 2, name: "gaurav" }]
    const arr2 = [{ id: 2, salary: "1000" }, { id: 1, salary: "2000" }]
    const result = mergeObjectArray(arr1, arr2)
    console.log(result);
}

main()